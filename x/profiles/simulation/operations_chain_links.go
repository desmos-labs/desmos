package simulation

import (
	"math/rand"
	"time"

	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"

	"github.com/desmos-labs/desmos/v4/testutil/profilestesting"
	"github.com/desmos-labs/desmos/v4/testutil/simtesting"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

// DONTCOVER

// SimulateMsgLinkChainAccount tests and runs a single MsgLinkChainAccount
func SimulateMsgLinkChainAccount(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		// Get the data
		link, signer, skip := randomLinkChainAccountFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.ModuleName, "", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgLinkChainAccount(
			link.GetAddressData(),
			link.Proof,
			link.ChainConfig,
			signer.Address.String(),
		)

		// Send the message
		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, signer)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "MsgLinkChainAccount", "invalid"), nil, nil
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomLinkChainAccountFields returns the data used to build a random MsgLinkChainAccount
func randomLinkChainAccountFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (link types.ChainLink, signer simtypes.Account, skip bool) {
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}

	// Get random signer
	signer, _ = simtypes.RandomAcc(r, accs)

	chainAccount := profilestesting.GetChainLinkAccount("cosmos", "cosmos")
	link = chainAccount.GetBech32ChainLink(signer.Address.String(), time.Now())

	// Skip if link already exists
	_, found := k.GetChainLink(ctx, signer.Address.String(), link.ChainConfig.Name, link.GetAddressData().GetValue())
	if found {
		skip = true
		return
	}

	return link, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgUnlinkChainAccount tests and runs a single MsgUnlinkChainAccount
func SimulateMsgUnlinkChainAccount(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		// Get the data
		link, signer, skip := randomUnlinkChainAccountFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.ModuleName, "MsgUnlinkChainAccount", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgUnlinkChainAccount(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue())

		// Send the message
		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, signer)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "MsgUnlinkChainAccount", "invalid"), nil, nil
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomUnlinkChainAccountFields returns the data used to build a random MsgUnlinkChainAccount
func randomUnlinkChainAccountFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (link types.ChainLink, signer simtypes.Account, skip bool) {
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}

	// Get a random chain link
	links := k.GetChainLinks(ctx)
	if len(links) == 0 {
		skip = true
		return
	}
	link = RandomChainLink(r, links)

	// Get the signer
	addr, _ := sdk.AccAddressFromBech32(link.User)
	signerAcc := GetSimAccount(addr, accs)
	if signerAcc == nil {
		skip = true
		return
	}

	signer = *signerAcc
	return link, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgSetDefaultExternalAddress tests and runs a single MsgSetDefaultExternalAddress
func SimulateMsgSetDefaultExternalAddress(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		// Get the data
		link, signer, skip := randomSetDefaultExternalAddressFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.ModuleName, "MsgSetDefaultExternalAddress", "skip"), nil, nil
		}

		// Build the message
		msg := types.NewMsgSetDefaultExternalAddress(link.ChainConfig.Name, link.GetAddressData().GetValue(), link.User)

		// Send the message
		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, signer)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, "MsgSetDefaultExternalAddress", "invalid"), nil, nil
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomSetDefaultExternalAddressFields returns the data used to build a random MsgSetDefaultExternalAddress
func randomSetDefaultExternalAddressFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (link types.ChainLink, signer simtypes.Account, skip bool) {
	if len(accs) == 0 {
		// Skip because there are no accounts
		skip = true
		return
	}

	// Get a random chain link
	links := k.GetChainLinks(ctx)
	if len(links) == 0 {
		skip = true
		return
	}
	link = RandomChainLink(r, links)

	// Get the signer
	addr, _ := sdk.AccAddressFromBech32(link.User)
	signerAcc := GetSimAccount(addr, accs)
	if signerAcc == nil {
		skip = true
		return
	}

	signer = *signerAcc
	return link, signer, false
}
