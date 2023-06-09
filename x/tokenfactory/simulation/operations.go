package simulation

// DONTCOVER

import (
	"math/rand"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	tokenfactorytypes "github.com/osmosis-labs/osmosis/v15/x/tokenfactory/types"

	"github.com/desmos-labs/desmos/v5/testutil/simtesting"
	subspacessim "github.com/desmos-labs/desmos/v5/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

// DONTCOVER

// Simulation operation weights constants
// #nosec G101 -- This is a false positive
const (
	OpWeightMsgCreateDenom      = "op_weight_msg_create_denom"
	OpWeightMsgMint             = "op_weight_msg_mint"
	OpWeightMsgBurn             = "op_weight_msg_burn"
	OpWeightMsgSetDenomMetadata = "op_weight_msg_set_denom_metadata"

	DefaultWeightMsgCreateDenom      = 30
	DefaultWeightMsgMint             = 70
	DefaultWeightMsgBurn             = 40
	DefaultWeightMsgSetDenomMetadata = 10
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	sk types.SubspacesKeeper, tfk types.TokenFactoryKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) sim.WeightedOperations {

	var weightMsgCreateDenom int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreateDenom, &weightMsgCreateDenom, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDenom = DefaultWeightMsgCreateDenom
		},
	)

	var weightMsgMint int
	appParams.GetOrGenerate(cdc, OpWeightMsgMint, &weightMsgMint, nil,
		func(_ *rand.Rand) {
			weightMsgMint = DefaultWeightMsgMint
		},
	)

	var weightMsgBurn int
	appParams.GetOrGenerate(cdc, OpWeightMsgBurn, &weightMsgBurn, nil,
		func(_ *rand.Rand) {
			weightMsgBurn = DefaultWeightMsgBurn
		},
	)

	var weightMsgSetDenomMetadata int
	appParams.GetOrGenerate(cdc, OpWeightMsgSetDenomMetadata, &weightMsgSetDenomMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgSetDenomMetadata = DefaultWeightMsgSetDenomMetadata
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgCreateDenom,
			SimulateMsgCreateDenom(sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgMint,
			SimulateMsgMint(sk, tfk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgBurn,
			SimulateMsgBurn(sk, tfk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgSetDenomMetadata,
			SimulateMsgSetDenomMetadata(sk, tfk, ak, bk),
		),
	}
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgCreateDenom tests and runs a single MsgCreateDenom
func SimulateMsgCreateDenom(
	sk types.SubspacesKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, subdenom, signer, skip := randomCreateDenomFields(r, ctx, accs, sk, bk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgCreateDenom", "skip"), nil, nil
		}

		msg := types.NewMsgCreateDenom(subspaceID, subdenom, signer.Address.String())

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomCreateDenomFields returns the data used to build a random MsgCreateDenom
func randomCreateDenomFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, sk types.SubspacesKeeper, bk types.BankKeeper,
) (subspaceID uint64, subdenom string, signer simtypes.Account, skip bool) {

	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a denom
	subdenom = simtypes.RandStringOfLength(r, 6)
	denom, _ := tokenfactorytypes.GetTokenDenom(subspace.Treasury, subdenom)
	_, exists := bk.GetDenomMetaData(ctx, denom)
	if exists {
		// Skip because denom has already existed
		skip = true
		return
	}

	// Get a signer
	admins := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionManageSubspaceTokens))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, admins), accs)
	if acc == nil {
		// Skip because the account is not valid
		skip = true
		return
	}
	signer = *acc

	return subspaceID, subdenom, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgMint tests and runs a single MsgMint
func SimulateMsgMint(
	sk types.SubspacesKeeper, tfk types.TokenFactoryKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, toAddress, amount, signer, skip := randomMintFields(r, ctx, accs, sk, tfk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgMint", "skip"), nil, nil
		}

		msg := types.NewMsgMint(subspaceID, signer.Address.String(), toAddress, amount)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomMintFields returns the data used to build a random MsgMint
func randomMintFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, sk types.SubspacesKeeper, tfk types.TokenFactoryKeeper,
) (subspaceID uint64, toAddress string, amount sdk.Coin, signer simtypes.Account, skip bool) {

	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get an amount
	denoms := tfk.GetDenomsFromCreator(ctx, subspace.Treasury)
	if len(denoms) == 0 {
		// Skip because there are no denoms
		skip = true
		return
	}
	denom := RandomDenom(r, denoms)
	amount = sdk.NewCoin(denom, simtypes.RandomAmount(r, math.NewInt(1000)))

	// 50% mint to subspace treasury
	if r.Intn(2)%2 == 0 {
		toAddress = subspace.Treasury
	} else {
		toAccount, _ := simtypes.RandomAcc(r, accs)
		toAddress = toAccount.Address.String()
	}

	// Get a signer
	admins := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionManageSubspaceTokens))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, admins), accs)
	if acc == nil {
		// Skip because the account is not valid
		skip = true
		return
	}
	signer = *acc

	return subspaceID, toAddress, amount, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgBurn tests and runs a single MsgBurn
func SimulateMsgBurn(
	sk types.SubspacesKeeper, tfk types.TokenFactoryKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, amount, signer, skip := randomBurnFields(r, ctx, accs, sk, tfk, bk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgBurn", "skip"), nil, nil
		}

		msg := types.NewMsgBurn(subspaceID, signer.Address.String(), amount)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomBurnFields returns the data used to build a random MsgBurn
func randomBurnFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, sk types.SubspacesKeeper, tfk types.TokenFactoryKeeper, bk bankkeeper.ViewKeeper,
) (subspaceID uint64, amount sdk.Coin, signer simtypes.Account, skip bool) {

	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get an amount
	denoms := tfk.GetDenomsFromCreator(ctx, subspace.Treasury)
	if len(denoms) == 0 {
		// Skip because there are no denoms
		skip = true
		return
	}
	denom := RandomDenom(r, denoms)
	balance := bk.GetBalance(ctx, sdk.MustAccAddressFromBech32(subspace.Treasury), denom)
	amount = sdk.NewCoin(denom, simtypes.RandomAmount(r, balance.Amount))

	// Get a signer
	admins := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionManageSubspaceTokens))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, admins), accs)
	if acc == nil {
		// Skip because the account is not valid
		skip = true
		return
	}
	signer = *acc

	return subspaceID, amount, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgSetDenomMetadata tests and runs a single MsgSetDenomMetadata
func SimulateMsgSetDenomMetadata(
	sk types.SubspacesKeeper, tfk types.TokenFactoryKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, metadata, signer, skip := randomSetDenomMetadataFields(r, ctx, accs, sk, tfk, bk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgSetDenomMetadata", "skip"), nil, nil
		}

		msg := types.NewMsgSetDenomMetadata(subspaceID, signer.Address.String(), metadata)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomSetDenomMetadataFields returns the data used to build a random MsgSetDenomMetadata
func randomSetDenomMetadataFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, sk types.SubspacesKeeper, tfk types.TokenFactoryKeeper, bk bankkeeper.ViewKeeper,
) (subspaceID uint64, metadata banktypes.Metadata, signer simtypes.Account, skip bool) {

	// Get a subspace id
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get an metadata
	denoms := tfk.GetDenomsFromCreator(ctx, subspace.Treasury)
	if len(denoms) == 0 {
		// Skip because there are no denoms
		skip = true
		return
	}
	denom := RandomDenom(r, denoms)
	display := simtypes.RandStringOfLength(r, 5)
	metadata = banktypes.Metadata{
		Description: simtypes.RandStringOfLength(r, 15),
		DenomUnits: []*banktypes.DenomUnit{
			{Denom: denom, Exponent: 0, Aliases: nil},
			{Denom: display, Exponent: 3, Aliases: nil},
		},
		Base:    denom,
		Display: simtypes.RandStringOfLength(r, 5),
		Name:    simtypes.RandStringOfLength(r, 5),
	}

	// Get a signer
	admins := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionManageSubspaceTokens))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, admins), accs)
	if acc == nil {
		// Skip because the account is not valid
		skip = true
		return
	}
	signer = *acc

	return subspaceID, metadata, signer, false
}
