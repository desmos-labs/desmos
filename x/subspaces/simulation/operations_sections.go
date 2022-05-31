package simulation

// DONTCOVER

import (
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v3/x/fees/keeper"

	"github.com/desmos-labs/desmos/v3/testutil/simtesting"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// SimulateMsgCreateSection tests and runs a single MsgCreateSection
func SimulateMsgCreateSection(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, update, parentID, creator, skip := randomCreateSectionFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateSection"), nil, nil
		}

		// Build the message
		msg := types.NewMsgCreateSection(
			subspaceID,
			update.Name,
			update.Description,
			parentID,
			creator.Address.String(),
		)

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgCreateSection"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgCreateSection", nil), nil, nil
	}
}

// randomCreateSectionFields returns the data used to build a random MsgCreateSection
func randomCreateSectionFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, update types.SectionUpdate, parentID uint32, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a parent section
	sections := k.GetSubspaceSections(ctx, subspaceID)
	parentSection := RandomSection(r, sections)
	parentID = parentSection.ID

	// Get a random update
	update = types.NewSectionUpdate(
		RandomSectionName(r),
		RandomSectionDescription(r),
	)

	// Get a signer
	signers, _ := k.GetUsersWithRootPermission(ctx, subspace.ID, types.PermissionManageSections)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, update, parentID, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgEditSection tests and runs a single MsgEditSection
func SimulateMsgEditSection(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, sectionID, update, creator, skip := randomEditSectionFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgEditSection"), nil, nil
		}

		// Build the message
		msg := types.NewMsgEditSection(
			subspaceID,
			sectionID,
			update.Name,
			update.Description,
			creator.Address.String(),
		)

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgEditSection"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgEditSection", nil), nil, nil
	}
}

// randomEditSectionFields returns the data used to build a random MsgEditSection
func randomEditSectionFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, sectionID uint32, update types.SectionUpdate, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a random section
	sections := k.GetSubspaceSections(ctx, subspaceID)
	section := RandomSection(r, sections)
	sectionID = section.ID

	// Get a random update
	update = types.NewSectionUpdate(
		RandomSectionName(r),
		RandomSectionDescription(r),
	)

	// Get a signer
	signers, _ := k.GetUsersWithRootPermission(ctx, subspace.ID, types.PermissionManageSections)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, sectionID, update, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgMoveSection tests and runs a single MsgMoveSection
func SimulateMsgMoveSection(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, sectionID, newParentID, creator, skip := randomMoveSectionFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgMoveSection"), nil, nil
		}

		// Build the message
		msg := types.NewMsgMoveSection(
			subspaceID,
			sectionID,
			newParentID,
			creator.Address.String(),
		)

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgMoveSection"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgMoveSection", nil), nil, nil
	}
}

// randomMoveSectionFields returns the data used to build a random MsgMoveSection
func randomMoveSectionFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, sectionID uint32, newParentID uint32, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a random section
	sections := k.GetSubspaceSections(ctx, subspaceID)
	section := RandomSection(r, sections)
	sectionID = section.ID
	if sectionID == types.RootSectionID {
		// Skip because we can't move the default section
		skip = true
		return
	}

	// Get a random new parent
	parent := RandomSection(r, sections)
	newParentID = parent.ID

	// Get a signer
	signers, _ := k.GetUsersWithRootPermission(ctx, subspace.ID, types.PermissionManageSections)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, sectionID, newParentID, account, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgDeleteSection tests and runs a single MsgDeleteSection
func SimulateMsgDeleteSection(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		// Get the data
		subspaceID, sectionID, creator, skip := randomDeleteFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteSection"), nil, nil
		}

		// Build the message
		msg := types.NewMsgDeleteSection(
			subspaceID,
			sectionID,
			creator.Address.String(),
		)

		// Send the message
		err := simtesting.SendMsg(r, app, ak, bk, fk, msg, ctx, chainID, DefaultGasValue, []cryptotypes.PrivKey{creator.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.RouterKey, types.ModuleName, "MsgDeleteSection"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "MsgDeleteSection", nil), nil, nil
	}
}

// randomDeleteFields returns the data used to build a random MsgDeleteSection
func randomDeleteFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (subspaceID uint64, sectionID uint32, account simtypes.Account, skip bool) {
	// Get a subspace id
	subspaces := k.GetAllSubspaces(ctx)
	if len(subspaces) == 0 {
		// Skip because there are no subspaces
		skip = true
		return
	}
	subspace := RandomSubspace(r, subspaces)
	subspaceID = subspace.ID

	// Get a random section
	sections := k.GetSubspaceSections(ctx, subspaceID)
	section := RandomSection(r, sections)
	sectionID = section.ID
	if sectionID == 0 {
		// Skip because we can't delete the default section
		skip = true
		return
	}

	// Get a signer
	signers, _ := k.GetUsersWithRootPermission(ctx, subspace.ID, types.PermissionManageSections)
	acc := GetAccount(RandomAddress(r, signers), accs)
	if acc == nil {
		// Skip the operation without error as the account is not valid
		skip = true
		return
	}
	account = *acc

	return subspaceID, sectionID, account, false
}
