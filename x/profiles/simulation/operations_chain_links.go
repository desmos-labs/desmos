package simulation

import (
	"math/rand"
	"time"

	feeskeeper "github.com/desmos-labs/desmos/v3/x/fees/keeper"

	"github.com/desmos-labs/desmos/v3/testutil/profilestesting"
	"github.com/desmos-labs/desmos/v3/testutil/simtesting"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"
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
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, nil
		}

		// Build the message
		msg := types.NewMsgLinkChainAccount(
			link.GetAddressData(),
			link.Proof,
			link.ChainConfig,
			signer.Address.String(),
		)

		// Send the message
		err = simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgRequestDTagTransfer"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
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
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, ""), nil, nil
		}

		// Build the message
		msg := types.NewMsgUnlinkChainAccount(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue())

		// Send the message
		err = simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{signer.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgUnlinkChainAccount"), nil, err
		}
		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
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
