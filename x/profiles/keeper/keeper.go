package keeper

import (
	"fmt"
	"regexp"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	ibctypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.BinaryMarshaler
	paramSubspace paramstypes.Subspace

	ak authkeeper.AccountKeeper

	channelKeeper ibctypes.ChannelKeeper
	portKeeper    ibctypes.PortKeeper
	scopedKeeper  capabilitykeeper.ScopedKeeper
}

// NewKeeper creates new instances of the Profiles Keeper.
// This k stores the profile data using two different associations:
// 1. Address -> Profile
//    This is used to easily retrieve the profile of a user based on an address
// 2. DTag -> Address
//    This is used to get the address of a user based on a DTag
func NewKeeper(
	cdc codec.BinaryMarshaler,
	storeKey sdk.StoreKey,
	paramSpace paramstypes.Subspace,
	ak authkeeper.AccountKeeper,
	channelKeeper ibctypes.ChannelKeeper,
	portKeeper ibctypes.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramSubspace: paramSpace,
		ak:            ak,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
	}
}

// IsUserBlocked tells if the given blocker has blocked the given blocked user
func (k Keeper) IsUserBlocked(ctx sdk.Context, blocker, blocked string) bool {
	return k.HasUserBlocked(ctx, blocker, blocked, "")
}

// StoreProfile stores the given profile inside the current context.
// It assumes that the given profile has already been validated.
// It returns an error if a profile with the same DTag from a different creator already exists
func (k Keeper) StoreProfile(ctx sdk.Context, profile *types.Profile) error {
	addr := k.GetAddressFromDTag(ctx, profile.DTag)
	if addr != "" && addr != profile.GetAddress().String() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"a profile with dtag %s has already been created", profile.DTag)
	}

	store := ctx.KVStore(k.storeKey)

	// Remove the previous DTag association (if the DTag has changed)
	oldProfile, found, err := k.GetProfile(ctx, profile.GetAddress().String())
	if err != nil {
		return err
	}
	if found && oldProfile.DTag != profile.DTag {
		store.Delete(types.DTagStoreKey(oldProfile.DTag))
	}

	// Store the DTag -> Address association
	store.Set(types.DTagStoreKey(profile.DTag), profile.GetAddress())

	// Store the account inside the auth keeper
	k.ak.SetAccount(ctx, profile)
	return nil
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

	// Delete the profile data by replacing the stored account
	k.ak.SetAccount(ctx, profile.GetAccount())
	return nil
}

// ValidateProfile checks if the given profile is valid according to the current profile's module params
func (k Keeper) ValidateProfile(ctx sdk.Context, profile *types.Profile) error {
	params := k.GetParams(ctx)

	minNicknameLen := params.NicknameParams.MinNicknameLength.Int64()
	maxNicknameLen := params.NicknameParams.MaxNicknameLength.Int64()

	if profile.Nickname != "" {
		nameLen := int64(len(profile.Nickname))
		if nameLen < minNicknameLen {
			return fmt.Errorf("profile nickname cannot be less than %d characters", minNicknameLen)
		}
		if nameLen > maxNicknameLen {
			return fmt.Errorf("profile nickname cannot exceed %d characters", maxNicknameLen)
		}
	}

	dTagRegEx := regexp.MustCompile(params.DTagParams.RegEx)
	minDTagLen := params.DTagParams.MinDTagLength.Int64()
	maxDTagLen := params.DTagParams.MaxDTagLength.Int64()
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

	maxBioLen := params.MaxBioLength.Int64()
	if profile.Bio != "" && int64(len(profile.Bio)) > maxBioLen {
		return fmt.Errorf("profile biography cannot exceed %d characters", maxBioLen)
	}

	return profile.Validate()
}

// ___________________________________________________________________________________________________________________

// SaveDTagTransferRequest save the given request associating it to the request recipient.
// It returns an error if the same request already exists.
func (k Keeper) SaveDTagTransferRequest(ctx sdk.Context, request types.DTagTransferRequest) error {
	store := ctx.KVStore(k.storeKey)
	key := types.DTagTransferRequestStoreKey(request.Receiver)

	var requests types.DTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &requests)
	for _, req := range requests.Requests {
		if req.Sender == request.Sender && req.Receiver == request.Receiver {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"the transfer request from %s to %s has already been made",
				request.Sender, request.Receiver)
		}
	}

	requests = types.NewDTagTransferRequests(append(requests.Requests, request))
	store.Set(key, k.cdc.MustMarshalBinaryBare(&requests))
	return nil
}

// GetUserIncomingDTagTransferRequests returns all the request made to the given user inside the current context.
func (k Keeper) GetUserIncomingDTagTransferRequests(ctx sdk.Context, user string) []types.DTagTransferRequest {
	store := ctx.KVStore(k.storeKey)
	key := types.DTagTransferRequestStoreKey(user)

	var requests types.DTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &requests)
	return requests.Requests
}

// GetDTagTransferRequests returns all the requests inside the given context
func (k Keeper) GetDTagTransferRequests(ctx sdk.Context) (requests []types.DTagTransferRequest) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DTagTransferRequestsPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var userRequests types.DTagTransferRequests
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &userRequests)
		requests = append(requests, userRequests.Requests...)
	}

	return requests
}

// DeleteAllDTagTransferRequests delete all the requests made to the given user
func (k Keeper) DeleteAllDTagTransferRequests(ctx sdk.Context, user string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.DTagTransferRequestStoreKey(user))
}

// DeleteDTagTransferRequest deletes the transfer requests made from the sender towards the recipient
func (k Keeper) DeleteDTagTransferRequest(ctx sdk.Context, sender, recipient string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.DTagTransferRequestStoreKey(recipient)

	var wrapped types.DTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &wrapped)

	for index, request := range wrapped.Requests {
		if request.Sender == sender {
			requests := append(wrapped.Requests[:index], wrapped.Requests[index+1:]...)
			if len(requests) == 0 {
				store.Delete(key)
			} else {
				store.Set(key, k.cdc.MustMarshalBinaryBare(&types.DTagTransferRequests{Requests: requests}))
			}
			return nil
		}
	}

	return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "request from %s to %s not found", sender, recipient)
}

// ____________________________________________________________________________________________________________________

// SaveRelationship allows to store the given relationship returning an error if he's already present.
func (k Keeper) SaveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator)

	relationships := types.MustUnmarshalRelationships(k.cdc, store.Get(key))
	for _, rel := range relationships {
		if rel.Equal(relationship) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"relationship already exists with %s", relationship.Recipient)
		}
	}

	relationships = append(relationships, relationship)
	store.Set(key, types.MustMarshalRelationships(k.cdc, relationships))
	return nil
}

// GetUserRelationships allows to list all the stored relationships that involve the given user.
func (k Keeper) GetUserRelationships(ctx sdk.Context, user string) []types.Relationship {
	var relationships []types.Relationship
	k.IterateRelationships(ctx, func(index int64, relationship types.Relationship) (stop bool) {
		if relationship.Creator == user || relationship.Recipient == user {
			relationships = append(relationships, relationship)
		}
		return false
	})
	return relationships
}

// GetAllRelationships allows to returns the list of all stored relationships
func (k Keeper) GetAllRelationships(ctx sdk.Context) []types.Relationship {
	var relationships []types.Relationship
	k.IterateRelationships(ctx, func(index int64, relationship types.Relationship) (stop bool) {
		relationships = append(relationships, relationship)
		return false
	})
	return relationships
}

// RemoveRelationship allows to delete the relationship between the given user and his counterparty
func (k Keeper) RemoveRelationship(ctx sdk.Context, relationship types.Relationship) error {
	store := ctx.KVStore(k.storeKey)
	key := types.RelationshipsStoreKey(relationship.Creator)

	relationships := types.MustUnmarshalRelationships(k.cdc, store.Get(key))
	relationships, found := types.RemoveRelationship(relationships, relationship)

	// The relationship didn't exist, so return an error
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"relationship between %s and %s for subspace %s not found",
			relationship.Creator, relationship.Recipient, relationship.Subspace)
	}

	// Delete the key if no relationships are left.
	// This cleans up the store avoiding export/import tests to fail due to a different number of keys present.
	if len(relationships) == 0 {
		store.Delete(key)
	} else {
		store.Set(key, types.MustMarshalRelationships(k.cdc, relationships))
	}
	return nil
}

// ___________________________________________________________________________________________________________________

// SaveUserBlock allows to store the given block inside the store, returning an error if
// something goes wrong.
func (k Keeper) SaveUserBlock(ctx sdk.Context, userBlock types.UserBlock) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UsersBlocksStoreKey(userBlock.Blocker)

	blocks := types.MustUnmarshalUserBlocks(k.cdc, store.Get(key))
	for _, ub := range blocks {
		if ub == userBlock {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"the user with address %s has already been blocked", userBlock.Blocked)
		}
	}

	store.Set(key, types.MustMarshalUserBlocks(k.cdc, append(blocks, userBlock)))
	return nil
}

// DeleteUserBlock allows to the specified blocker to unblock the given blocked user.
func (k Keeper) DeleteUserBlock(ctx sdk.Context, blocker, blocked string, subspace string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UsersBlocksStoreKey(blocker)

	blocks := types.MustUnmarshalUserBlocks(k.cdc, store.Get(key))

	blocks, found := types.RemoveUserBlock(blocks, blocker, blocked, subspace)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"block from %s towards %s for subspace %s not found", blocker, blocked, subspace)
	}

	// Delete the key if no blocks are left.
	// This cleans up the store avoiding export/import tests to fail due to a different number of keys present.
	if len(blocks) == 0 {
		store.Delete(key)
	} else {
		store.Set(key, types.MustMarshalUserBlocks(k.cdc, blocks))
	}
	return nil
}

// GetUserBlocks returns the list of users that the specified user has blocked.
func (k Keeper) GetUserBlocks(ctx sdk.Context, user string) []types.UserBlock {
	store := ctx.KVStore(k.storeKey)
	return types.MustUnmarshalUserBlocks(k.cdc, store.Get(types.UsersBlocksStoreKey(user)))
}

// GetAllUsersBlocks returns a list of all the users blocks inside the given context.
func (k Keeper) GetAllUsersBlocks(ctx sdk.Context) []types.UserBlock {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.UsersBlocksStorePrefix)
	defer iterator.Close()

	var usersBlocks []types.UserBlock
	for ; iterator.Valid(); iterator.Next() {
		blocks := types.MustUnmarshalUserBlocks(k.cdc, iterator.Value())
		usersBlocks = append(usersBlocks, blocks...)
	}

	return usersBlocks
}

// HasUserBlocked returns true if the provided blocker has blocked the given user for the given subspace.
// If the provided subspace is empty, all subspaces will be checked
func (k Keeper) HasUserBlocked(ctx sdk.Context, blocker, user, subspace string) bool {
	blocks := k.GetUserBlocks(ctx, blocker)

	for _, block := range blocks {
		if block.Blocked == user {
			return subspace == "" || block.Subspace == subspace
		}
	}

	return false
}

// ___________________________________________________________________________________________________________________

// GetLinkPubkey returns the pubkey corresponding to the given account
func (k Keeper) GetAccountPubKey(ctx sdk.Context, acc sdk.AccAddress) (cryptotypes.PubKey, error) {
	return k.ak.GetPubKey(ctx, acc)
}
