package keeper

import (
	"fmt"
	"regexp"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.BinaryCodec
	legacyAmino   *codec.LegacyAmino
	paramSubspace paramstypes.Subspace
	hooks         types.ProfilesHooks

	ak authkeeper.AccountKeeper
	rk types.RelationshipsKeeper

	channelKeeper types.ChannelKeeper
	portKeeper    types.PortKeeper
	scopedKeeper  types.ScopedKeeper

	wasmKeeper *wasmkeeper.Keeper
}

// NewKeeper creates new instances of the Profiles Keeper.
// This k stores the profile data using two different associations:
// 1. Address -> Profile
//    This is used to easily retrieve the profile of a user based on an address
// 2. DTag -> Address
//    This is used to get the address of a user based on a DTag
func NewKeeper(
	cdc codec.BinaryCodec,
	legacyAmino *codec.LegacyAmino,
	storeKey sdk.StoreKey,
	paramSpace paramstypes.Subspace,
	ak authkeeper.AccountKeeper,
	rk types.RelationshipsKeeper,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper types.ScopedKeeper,
	wasmKeeper *wasmkeeper.Keeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		legacyAmino:   legacyAmino,
		paramSubspace: paramSpace,
		ak:            ak,
		rk:            rk,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
		wasmKeeper:    wasmKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// SetHooks allows to set the profiles hooks
func (k *Keeper) SetHooks(ph types.ProfilesHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set profiles hooks twice")
	}

	k.hooks = ph
	return k
}

// IsUserBlocked returns true if the provided blocker has blocked the given user for the given subspace.
// If the provided subspace is empty, all subspaces will be checked
func (k Keeper) IsUserBlocked(ctx sdk.Context, user, blocker string) bool {
	return k.rk.HasUserBlocked(ctx, user, blocker, 0)
}

// storeProfileWithoutDTagCheck stores the given profile inside the current context
// without checking if another profile with the same DTag already exists.
// It assumes that the given profile has already been validated.
func (k Keeper) storeProfileWithoutDTagCheck(ctx sdk.Context, profile *types.Profile) error {
	store := ctx.KVStore(k.storeKey)

	oldProfile, found, err := k.GetProfile(ctx, profile.GetAddress().String())
	if err != nil {
		return err
	}
	if found && oldProfile.DTag != profile.DTag {
		// Remove the previous DTag association (if the DTag has changed)
		store.Delete(types.DTagStoreKey(oldProfile.DTag))

		// Remove all incoming DTag transfer requests if the DTag has changed since these will be invalid now
		k.DeleteAllUserIncomingDTagTransferRequests(ctx, profile.GetAddress().String())
	}

	// Store the DTag -> Address association
	store.Set(types.DTagStoreKey(profile.DTag), profile.GetAddress())

	// Store the account inside the auth keeper
	k.ak.SetAccount(ctx, profile)

	k.Logger(ctx).Info("saved profile", "DTag", profile.DTag, "from", profile.GetAddress())

	k.AfterProfileSaved(ctx, profile)

	return nil
}

// SaveProfile stores the given profile inside the current context.
// It assumes that the given profile has already been validated.
// It returns an error if a profile with the same DTag from a different creator already exists
func (k Keeper) SaveProfile(ctx sdk.Context, profile *types.Profile) error {
	addr := k.GetAddressFromDTag(ctx, profile.DTag)
	if addr != "" && addr != profile.GetAddress().String() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"a profile with DTag %s has already been created", profile.DTag)
	}
	return k.storeProfileWithoutDTagCheck(ctx, profile)
}

// GetProfile returns the profile corresponding to the given address inside the current context.
func (k Keeper) GetProfile(ctx sdk.Context, address string) (profile *types.Profile, found bool, err error) {
	sdkAcc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return nil, false, err
	}

	stored, ok := k.ak.GetAccount(ctx, sdkAcc).(*types.Profile)
	if !ok {
		return nil, false, nil
	}

	return stored, true, nil
}

// GetAddressFromDTag returns the address associated to the given DTag or an empty string if it does not exists
func (k Keeper) GetAddressFromDTag(ctx sdk.Context, dTag string) (addr string) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.DTagStoreKey(dTag))
	if bz == nil {
		return ""
	}

	return sdk.AccAddress(bz).String()
}

// RemoveProfile allows to delete a profile associated with the given address inside the current context.
// It assumes that the address-related profile exists.
func (k Keeper) RemoveProfile(ctx sdk.Context, address string) error {
	profile, found, err := k.GetProfile(ctx, address)
	if err != nil {
		return err
	}

	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"no profile associated with the following address found: %s", address)
	}

	// Delete the DTag -> Address association
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.DTagStoreKey(profile.DTag))

	// Delete all DTag transfer requests made towards this account
	k.DeleteAllUserIncomingDTagTransferRequests(ctx, address)

	// Delete all chains links
	k.DeleteAllUserChainLinks(ctx, address)

	// Delete all the application links
	k.DeleteAllUserApplicationLinks(ctx, address)

	// Delete the profile data by replacing the stored account
	k.ak.SetAccount(ctx, profile.GetAccount())

	k.AfterProfileDeleted(ctx, profile)

	return nil
}

// ValidateProfile checks if the given profile is valid according to the current profile's module params
func (k Keeper) ValidateProfile(ctx sdk.Context, profile *types.Profile) error {
	params := k.GetParams(ctx)

	minNicknameLen := params.Nickname.MinLength.Int64()
	maxNicknameLen := params.Nickname.MaxLength.Int64()

	if profile.Nickname != "" {
		nameLen := int64(len(profile.Nickname))
		if nameLen < minNicknameLen {
			return fmt.Errorf("profile nickname cannot be less than %d characters", minNicknameLen)
		}
		if nameLen > maxNicknameLen {
			return fmt.Errorf("profile nickname cannot exceed %d characters", maxNicknameLen)
		}
	}

	dTagRegEx := regexp.MustCompile(params.DTag.RegEx)
	minDTagLen := params.DTag.MinLength.Int64()
	maxDTagLen := params.DTag.MaxLength.Int64()
	dTagLen := int64(len(profile.DTag))

	if !dTagRegEx.MatchString(profile.DTag) {
		return fmt.Errorf("invalid profile dtag, it should match the following regEx %s", dTagRegEx)
	}

	if dTagLen < minDTagLen {
		return fmt.Errorf("profile dtag cannot be less than %d characters", minDTagLen)
	}

	if dTagLen > maxDTagLen {
		return fmt.Errorf("profile dtag cannot exceed %d characters", maxDTagLen)
	}

	maxBioLen := params.Bio.MaxLength.Int64()
	if profile.Bio != "" && int64(len(profile.Bio)) > maxBioLen {
		return fmt.Errorf("profile biography cannot exceed %d characters", maxBioLen)
	}

	return profile.Validate()
}
