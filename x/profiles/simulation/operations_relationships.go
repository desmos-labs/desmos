package simulation

// DONTCOVER

import (
	"math/rand"

	subspacessim "github.com/desmos-labs/desmos/v2/x/subspaces/simulation"

	"github.com/desmos-labs/desmos/v2/testutil/simtesting"

	"github.com/desmos-labs/desmos/v2/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SimulateMsgCreateRelationship tests and runs a single msg create relationships
func SimulateMsgCreateRelationship(
	k keeper.Keeper, sk keeper.SubspacesKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		acc, relationship, skip := randomRelationshipFields(r, ctx, accs, k, sk)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateRelationship"), nil, nil
		}

		msg := types.NewMsgCreateRelationship(relationship.Creator, relationship.Recipient, relationship.SubspaceID)
		err = simtesting.SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{acc.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateRelationship"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgCreateRelationship", nil), nil, nil
	}
}

// randomRelationshipFields returns random relationships fields
func randomRelationshipFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, sk keeper.SubspacesKeeper,
) (simtypes.Account, types.Relationship, bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, types.Relationship{}, true
	}

	// Get random accounts
	sender, _ := simtypes.RandomAcc(r, accs)
	receiver, _ := simtypes.RandomAcc(r, accs)

	subspaces := sk.GetAllSubspaces(ctx)
	subspace, _ := subspacessim.RandomSubspace(r, subspaces)

	// Skip if the send and receiver are equals
	if sender.Equals(receiver) {
		return simtypes.Account{}, types.Relationship{}, true
	}

	// Skip if the creator does not have a profile
	if !k.HasProfile(ctx, sender.Address.String()) {
		return simtypes.Account{}, types.Relationship{}, true
	}

	// Skip if the receiver does not have a profile
	if !k.HasProfile(ctx, receiver.Address.String()) {
		return simtypes.Account{}, types.Relationship{}, true
	}

	// Skip if the receiver has block the sender
	if k.HasUserBlocked(ctx, receiver.Address.String(), sender.Address.String(), subspace.ID) {
		return simtypes.Account{}, types.Relationship{}, true
	}

	rel := types.NewRelationship(sender.Address.String(), receiver.Address.String(), subspace.ID)

	// Skip if relationships already exists
	relationships := k.GetUserRelationships(ctx, sender.Address.String())
	for _, relationship := range relationships {
		if relationship.Equal(rel) {
			return simtypes.Account{}, types.Relationship{}, true
		}
	}

	return sender, rel, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgDeleteRelationship tests and runs a single msg delete relationships
func SimulateMsgDeleteRelationship(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		acc, counterparty, subspace, skip := randomDeleteRelationshipFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteRelationship"), nil, nil
		}

		msg := types.NewMsgDeleteRelationship(acc.Address.String(), counterparty, subspace)
		err = simtesting.SendMsg(r, app, ak, bk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{acc.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteRelationship"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgDeleteRelationship", nil), nil, nil
	}
}

// randomDeleteRelationshipFields returns random delete relationships fields
func randomDeleteRelationshipFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (user simtypes.Account, counterparty string, subspace uint64, skip bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, "", 0, true
	}

	// Get a random account
	user, _ = simtypes.RandomAcc(r, accs)

	// Get the user relationships
	relationships := k.GetUserRelationships(ctx, user.Address.String())

	// Remove all the relationships where the user is not the creator
	var outgoingRelationships []types.Relationship
	for _, relationship := range relationships {
		if user.Address.String() == relationship.Creator {
			outgoingRelationships = append(outgoingRelationships, relationship)
		}
	}

	// Skip the test if the user has no relationships
	if len(outgoingRelationships) == 0 {
		return simtypes.Account{}, "", 0, true
	}

	// Get a random relationship
	relationship := RandomRelationship(r, outgoingRelationships)
	return user, relationship.Recipient, relationship.SubspaceID, false
}
