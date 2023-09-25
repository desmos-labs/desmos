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

	"github.com/desmos-labs/desmos/v6/testutil/simtesting"
	subspacessim "github.com/desmos-labs/desmos/v6/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
	"github.com/desmos-labs/desmos/v6/x/tokenfactory/keeper"
	"github.com/desmos-labs/desmos/v6/x/tokenfactory/types"
)

// DONTCOVER

// Simulation operation weights constants
//
//nolint:gosec // This is a false positive
const (
	OpWeightMsgCreateDenom      = "op_weight_msg_create_denom"
	OpWeightMsgMint             = "op_weight_msg_mint"
	OpWeightMsgBurn             = "op_weight_msg_burn"
	OpWeightMsgSetDenomMetadata = "op_weight_msg_set_denom_metadata"

	DefaultWeightMsgCreateDenom      = 30
	DefaultWeightMsgMint             = 70
	DefaultWeightMsgBurn             = 40
	DefaultWeightMsgSetDenomMetadata = 20
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	k keeper.Keeper, sk types.SubspacesKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
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
			SimulateMsgCreateDenom(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgMint,
			SimulateMsgMint(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgBurn,
			SimulateMsgBurn(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgSetDenomMetadata,
			SimulateMsgSetDenomMetadata(k, sk, ak, bk),
		),
	}
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgCreateDenom tests and runs a single MsgCreateDenom
func SimulateMsgCreateDenom(
	k keeper.Keeper, sk types.SubspacesKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, subdenom, signer, skip := randomCreateDenomFields(r, ctx, accs, k, sk, bk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgCreateDenom", "skip"), nil, nil
		}

		msg := types.NewMsgCreateDenom(subspaceID, signer.Address.String(), subdenom)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomCreateDenomFields returns the data used to build a random MsgCreateDenom
func randomCreateDenomFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk types.SubspacesKeeper, bk bankkeeper.Keeper,
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

	// Check treasury balances
	balances := bk.SpendableCoins(ctx, sdk.MustAccAddressFromBech32(subspace.Treasury))
	creationFees := k.GetParams(ctx).DenomCreationFee
	if !balances.IsAllGT(creationFees) {
		// Skip because treasury does not have enough coins
		skip = true
		return
	}

	// Get a denom
	subdenom = simtypes.RandStringOfLength(r, 6)
	denom, _ := types.GetTokenDenom(subspace.Treasury, subdenom)
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
	k keeper.Keeper, sk types.SubspacesKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, amount, signer, skip := randomMintFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgMint", "skip"), nil, nil
		}

		msg := types.NewMsgMint(subspaceID, signer.Address.String(), amount)

		// Send the message
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomMintFields returns the data used to build a random MsgMint
func randomMintFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk types.SubspacesKeeper,
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
	denoms := k.GetDenomsFromCreator(ctx, subspace.Treasury)
	if len(denoms) == 0 {
		// Skip because there are no denoms
		skip = true
		return
	}
	denom := RandomDenom(r, denoms)
	amount = sdk.NewCoin(denom, simtypes.RandomAmount(r, math.NewInt(1000)))
	if amount.Amount.Equal(math.NewInt(0)) {
		// Skip because amount with zero is invalid
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

	return subspaceID, amount, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgBurn tests and runs a single MsgBurn
func SimulateMsgBurn(
	k keeper.Keeper, sk types.SubspacesKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, amount, signer, skip := randomBurnFields(r, ctx, accs, k, sk, bk)
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
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk types.SubspacesKeeper, bk bankkeeper.ViewKeeper,
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

	// Get a denom
	denoms := k.GetDenomsFromCreator(ctx, subspace.Treasury)
	if len(denoms) == 0 {
		// Skip because there are no denoms
		skip = true
		return
	}
	denom := RandomDenom(r, denoms)

	// Get a amount to burn
	balance := bk.SpendableCoin(ctx, sdk.MustAccAddressFromBech32(subspace.Treasury), denom)
	amount = sdk.NewCoin(denom, simtypes.RandomAmount(r, balance.Amount))
	if amount.Amount.Equal(math.NewInt(0)) {
		// Skip because amount with zero is invalid
		skip = true
		return
	}

	// Get a signer
	admins := sk.GetUsersWithRootPermissions(ctx, subspace.ID, subspacestypes.NewPermissions(types.PermissionManageSubspaceTokens))
	acc := subspacessim.GetAccount(subspacessim.RandomAddress(r, admins), accs)
	if acc == nil {
		// Skip because the account is invalid
		skip = true
		return
	}
	signer = *acc

	return subspaceID, amount, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgSetDenomMetadata tests and runs a single MsgSetDenomMetadata
func SimulateMsgSetDenomMetadata(
	k keeper.Keeper, sk types.SubspacesKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, _ string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, metadata, signer, skip := randomSetDenomMetadataFields(r, ctx, accs, k, sk)
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
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk types.SubspacesKeeper,
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
	denoms := k.GetDenomsFromCreator(ctx, subspace.Treasury)
	if len(denoms) == 0 {
		// Skip because there are no denoms
		skip = true
		return
	}
	denom := RandomDenom(r, denoms)
	display := simtypes.RandStringOfLength(r, 3)
	metadata = banktypes.Metadata{
		Description: simtypes.RandStringOfLength(r, 15),
		DenomUnits: []*banktypes.DenomUnit{
			{Denom: denom, Exponent: 0, Aliases: nil},
			{Denom: display, Exponent: 3, Aliases: nil},
		},
		Base:    denom,
		Display: display,
		Name:    simtypes.RandStringOfLength(r, 5),
		Symbol:  simtypes.RandStringOfLength(r, 5),
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
