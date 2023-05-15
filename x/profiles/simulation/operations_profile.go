package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/desmos-labs/desmos/v5/testutil/simtesting"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"

	"github.com/desmos-labs/desmos/v5/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v5/x/profiles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SimulateMsgSaveProfile tests and runs a single msg save profile where the creator already exists
func SimulateMsgSaveProfile(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {

		acc, data, skip := randomProfileSaveFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "save profile", "skip"), nil, nil
		}

		msg := types.NewMsgSaveProfile(
			data.DTag,
			data.Nickname,
			data.Bio,
			data.Pictures.Profile,
			data.Pictures.Cover,
			acc.Address.String(),
		)
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, acc)
	}
}

// randomProfileSaveFields returns random profile data
func randomProfileSaveFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (account simtypes.Account, profile *types.Profile, skip bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, nil, true
	}

	// Get a random account
	account, _ = simtypes.RandomAcc(r, accs)
	acc := ak.GetAccount(ctx, account.Address)

	// See if there is already the profile, otherwise create it from scratch
	existing, found, err := k.GetProfile(ctx, account.Address.String())
	if err != nil {
		return simtypes.Account{}, nil, true
	}

	if found {
		profile = existing
	} else {
		profile = NewRandomProfile(r, acc)
	}

	// Skip if another profile with the same DTag exists
	address := k.GetAddressFromDTag(ctx, profile.DTag)
	if address != acc.GetAddress().String() {
		return simtypes.Account{}, nil, true
	}

	// 50% chance of changing something
	if r.Intn(101) <= 50 {
		profile, _ = profile.Update(types.NewProfileUpdate(
			RandomDTag(r),
			RandomNickname(r),
			RandomBio(r),
			types.NewPictures(RandomProfilePic(r), RandomProfileCover(r)),
		))
	}

	return account, profile, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgDeleteProfile tests and runs a single msg delete profile where the creator already exists
func SimulateMsgDeleteProfile(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		acc, skip := randomProfileDeleteFields(r, ctx, accs, k, ak)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "delete profile", "skip"), nil, nil
		}

		msg := types.NewMsgDeleteProfile(acc.Address.String())

		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, acc)
	}
}

// randomProfileDeleteFields returns random profile data
func randomProfileDeleteFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper, ak authkeeper.AccountKeeper,
) (account simtypes.Account, skip bool) {
	if len(accs) == 0 {
		return simtypes.Account{}, true
	}

	// Get a random account
	account, _ = simtypes.RandomAcc(r, accs)
	acc := ak.GetAccount(ctx, account.Address)

	// See if the account has a profile, and skip if he does not
	_, found, err := k.GetProfile(ctx, acc.GetAddress().String())
	if !found || err != nil {
		return simtypes.Account{}, true
	}

	return account, false
}
