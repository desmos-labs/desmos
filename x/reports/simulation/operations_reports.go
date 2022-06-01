package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	postskeeper "github.com/desmos-labs/desmos/v3/x/posts/keeper"
	postssim "github.com/desmos-labs/desmos/v3/x/posts/simulation"
	subspaceskeeper "github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	subspacessim "github.com/desmos-labs/desmos/v3/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	feeskeeper "github.com/desmos-labs/desmos/v3/x/fees/keeper"

	"github.com/desmos-labs/desmos/v3/testutil/simtesting"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v3/x/reports/keeper"
	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

// SimulateMsgCreateReport tests and runs a single MsgCreateReport
func SimulateMsgCreateReport(
	k keeper.Keeper, sk subspaceskeeper.Keeper, pk postskeeper.Keeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		data, creator, skip := randomCreateReportFields(r, ctx, accs, sk, pk, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateReport"), nil, nil
		}

		// Build the message
		msg := types.NewMsgCreateReport(
			data.SubspaceID,
			data.ReasonID,
			data.Message,
			data.Target.GetCachedValue().(types.ReportTarget),
			creator.Address.String(),
		)

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateReport"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgCreateReport", nil), nil, nil
	}
}

// randomCreateReportFields returns the data used to build a random MsgCreateReport
func randomCreateReportFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, sk subspaceskeeper.Keeper, pk postskeeper.Keeper, k keeper.Keeper,
) (report types.Report, creator simtypes.Account, skip bool) {
	// Get the creator
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}
	creator, _ = simtypes.RandomAcc(r, accs)

	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID := subspace.ID

	// Get a reason
	reasons := k.GetSubspaceReasons(ctx, subspaceID)
	if len(reasons) == 0 {
		// Skip because there are no reasons to report the data for
		skip = true
		return
	}
	reason := RandomReason(r, reasons)

	// Get a report target
	var data types.ReportTarget
	if r.Intn(101) < 50 {
		// 50% of having a user report
		user, _ := simtypes.RandomAcc(r, accs)
		data = types.NewUserTarget(user.Address.String())
	} else {
		posts := pk.GetSubspacePosts(ctx, subspaceID)
		if len(posts) == 0 {
			// Skip because there are no posts to be reported
			skip = true
			return
		}
		post := postssim.RandomPost(r, posts)
		data = types.NewPostTarget(post.ID)
	}

	// Get a reporter
	reporters, _ := sk.GetUsersWithRootPermission(ctx, subspace.ID, subspacestypes.PermissionReportContent)
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, reporters), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	creator = *acc

	// Get the report target
	report = types.NewReport(
		subspaceID,
		0,
		reason.ID,
		GetRandomMessage(r),
		data,
		creator.Address.String(),
		time.Time{},
	)

	return report, creator, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgDeleteReport tests and runs a single msg delete subspace
func SimulateMsgDeleteReport(
	k keeper.Keeper, sk subspaceskeeper.Keeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, reportID, editor, skip := randomDeleteReportFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteReport"), nil, nil
		}

		// Build the message
		msg := types.NewMsgDeleteReport(subspaceID, reportID, editor.Address.String())

		// Send the data
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, 1_500_000, []cryptotypes.PrivKey{editor.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteReport"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgDeleteReport", nil), nil, nil
	}
}

// randomDeleteReportFields returns the data needed to delete a subspace
func randomDeleteReportFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk subspaceskeeper.Keeper,
) (subspaceID uint64, reportID uint64, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a report
	reports := k.GetSubspaceReports(ctx, subspaceID)
	if len(reports) == 0 {
		// Skip because there are no reports
		skip = true
		return
	}
	report := RandomReport(r, reports)
	reportID = report.ID

	// Get an editor
	editors, _ := sk.GetUsersWithRootPermission(ctx, subspace.ID, subspacestypes.PermissionManageReports)
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, editors), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, reportID, account, false
}
