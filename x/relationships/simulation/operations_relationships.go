package simulation

// DONTCOVER

import (
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"

	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"

	subspacessim "github.com/desmos-labs/desmos/v4/x/subspaces/simulation"

	"github.com/desmos-labs/desmos/v4/testutil/simtesting"

	"github.com/desmos-labs/desmos/v4/x/relationships/keeper"
	"github.com/desmos-labs/desmos/v4/x/relationships/types"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SimulateMsgCreateRelationship tests and runs a single MsgCreateRelationship
func SimulateMsgCreateRelationship(
	k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		acc, relationship, skip := randomCreateRelationshipFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateRelationship"), nil, nil
		}

		msg := types.NewMsgCreateRelationship(relationship.Creator, relationship.Counterparty, relationship.SubspaceID)
		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, acc)
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateRelationship"), nil, err
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomCreateRelationshipFields returns the data used to build a random MsgCreateRelationship
func randomCreateRelationshipFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk subspaceskeeper.Keeper,
) (sender simtypes.Account, relationship types.Relationship, skip bool) {
	// Get a sender
	if len(accs) == 0 {
		skip = true
		return
	}
	sender, _ = simtypes.RandomAcc(r, accs)
	senderAddr := sender.Address.String()

	// Get a receiver
	receiver, _ := simtypes.RandomAcc(r, accs)
	receiverAddr := receiver.Address.String()

	// Get a subspace
	subspaces := sk.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := subspacessim.RandomSubspace(r, subspaces)
	subspaceID := subspace.ID

	// Skip if the sender and receiver are equals
	if senderAddr == receiverAddr {
		skip = true
		return
	}

	// Skip if the receiver has blocked the sender
	if k.HasUserBlocked(ctx, receiverAddr, senderAddr, subspaceID) {
		skip = true
		return
	}

	// Skip if relationships already exists
	if k.HasRelationship(ctx, senderAddr, receiverAddr, subspaceID) {
		skip = true
		return
	}

	return sender, types.NewRelationship(senderAddr, receiverAddr, subspaceID), false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgDeleteRelationship tests and runs a single MsgDeleteRelationship
func SimulateMsgDeleteRelationship(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		acc, counterparty, subspace, skip := randomDeleteRelationshipFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteRelationship"), nil, nil
		}

		msg := types.NewMsgDeleteRelationship(acc.Address.String(), counterparty, subspace)
		txCtx, err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, acc)
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteRelationship"), nil, err
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// randomDeleteRelationshipFields returns the data used to build a random MsgDeleteRelationship
func randomDeleteRelationshipFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (user simtypes.Account, counterparty string, subspace uint64, skip bool) {
	// Get a user
	if len(accs) == 0 {
		skip = true
		return
	}
	user, _ = simtypes.RandomAcc(r, accs)
	userAddr := user.Address.String()

	// Get the user relationships
	var relationships []types.Relationship
	k.IterateRelationships(ctx, func(_ int64, relationship types.Relationship) (stop bool) {
		if relationship.Creator == userAddr {
			relationships = append(relationships, relationship)
		}
		return false
	})
	if len(relationships) == 0 {
		// Skip because there are no relationships
		skip = true
		return
	}

	// Get a random relationship
	relationship := RandomRelationship(r, relationships)
	return user, relationship.Counterparty, relationship.SubspaceID, false
}
