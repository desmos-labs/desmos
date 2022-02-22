package v300

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	subspacestypes "github.com/desmos-labs/desmos/v2/x/subspaces/types"

	v2 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v2"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v2.3 to v3.0.
// The migration includes:
//
// - replace all relationship subspace id from string to uint64
// - replace all user blocks subspace id from string to uint64
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, subspace paramstypes.Subspace, cdc codec.BinaryCodec,
	legacyAmino *codec.LegacyAmino) error {
	store := ctx.KVStore(storeKey)

	err := migrateUserBlocks(store, cdc)
	if err != nil {
		return err
	}

	err = migrateRelationships(store, cdc)
	if err != nil {
		return err
	}

	err = migrateParams(ctx, subspace, legacyAmino)
	if err != nil {
		return err
	}

	err = migrateAppLinks(store, cdc)
	if err != nil {
		return err
	}

	return nil
}

// migrateUserBlocks migrates the user blocks stored to the new type, converting the subspace from string to uint64
func migrateUserBlocks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	var values []v2.UserBlock

	userBlocksStore := prefix.NewStore(store, v2.UsersBlocksStorePrefix)
	iterator := userBlocksStore.Iterator(nil, nil)

	for ; iterator.Valid(); iterator.Next() {
		var block v2.UserBlock
		err := cdc.Unmarshal(iterator.Value(), &block)
		if err != nil {
			return err
		}
		values = append(values, block)
	}

	// Close the iterator
	err := iterator.Close()
	if err != nil {
		return err
	}

	for _, v230Block := range values {
		// Delete the previous key
		store.Delete(v2.UserBlockStoreKey(v230Block.Blocker, v230Block.Subspace, v230Block.Blocked))

		// Get the subspace id
		subspaceID, err := subspacestypes.ParseSubspaceID(v230Block.Subspace)
		if err != nil {
			return err
		}

		// Serialize the block as the new type
		v300Block := types.NewUserBlock(v230Block.Blocker, v230Block.Blocked, v230Block.Reason, subspaceID)
		blockBz, err := cdc.Marshal(&v300Block)
		if err != nil {
			return err
		}

		// Store the new value inside the store
		store.Set(types.UserBlockStoreKey(v300Block.Blocker, v300Block.SubspaceID, v300Block.Blocked), blockBz)
	}

	return nil
}

// migrateRelationships migrates the relationships stored to the new type, converting the subspace from string to uint64
func migrateRelationships(store sdk.KVStore, cdc codec.BinaryCodec) error {
	var values []v2.Relationship

	relationshipsStore := prefix.NewStore(store, types.RelationshipsStorePrefix)
	iterator := relationshipsStore.Iterator(nil, nil)

	for ; iterator.Valid(); iterator.Next() {
		var relationship v2.Relationship
		err := cdc.Unmarshal(iterator.Value(), &relationship)
		if err != nil {
			return err
		}
		values = append(values, relationship)
	}

	// Close the iterator
	err := iterator.Close()
	if err != nil {
		return err
	}

	for _, v230Relationship := range values {
		// Delete the previous key
		store.Delete(v2.RelationshipsStoreKey(v230Relationship.Creator, v230Relationship.Subspace, v230Relationship.Recipient))

		// Get the subspace id
		subspaceID, err := subspacestypes.ParseSubspaceID(v230Relationship.Subspace)
		if err != nil {
			return err
		}

		// Serialize the relationship as the new type
		v300Relationship := types.NewRelationship(v230Relationship.Creator, v230Relationship.Recipient, subspaceID)
		relationshipBz, err := cdc.Marshal(&v300Relationship)
		if err != nil {
			return err
		}

		// Store the new relationship inside the store
		store.Set(types.RelationshipsStoreKey(v300Relationship.Creator, v300Relationship.SubspaceID, v300Relationship.Recipient), relationshipBz)
	}

	return nil
}

func migrateAppLinks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	iterator := sdk.KVStorePrefixIterator(store, types.UserApplicationLinkPrefix)
	defer iterator.Close()

	var keys [][]byte
	var newLinks []types.ApplicationLink
	for ; iterator.Valid(); iterator.Next() {
		var legacyAppLink v2.ApplicationLink
		err := cdc.Unmarshal(iterator.Value(), &legacyAppLink)
		if err != nil {
			return err
		}

		link := types.NewApplicationLink(
			legacyAppLink.User,
			types.NewData(legacyAppLink.Data.Application, legacyAppLink.Data.Username),
			types.ApplicationLinkState(legacyAppLink.State),
			types.NewOracleRequest(
				legacyAppLink.OracleRequest.ID,
				legacyAppLink.OracleRequest.OracleScriptID,
				types.NewOracleRequestCallData(
					legacyAppLink.OracleRequest.CallData.Application,
					legacyAppLink.OracleRequest.CallData.CallData,
				),
				legacyAppLink.OracleRequest.ClientID,
			),
			migrateAppLinkResult(legacyAppLink.Result),
			legacyAppLink.CreationTime,
			legacyAppLink.CreationTime.Add(types.DefaultAppLinksParams().ExpirationTime),
		)

		keys = append(keys, iterator.Key())
		newLinks = append(newLinks, link)

		store.Delete(iterator.Key())
	}

	for index, link := range newLinks {
		store.Set(keys[index], types.MustMarshalApplicationLink(cdc, link))
		store.Set(
			types.ApplicationLinkExpiringTimeKey(link.ExpirationTime, link.OracleRequest.ClientID),
			[]byte(link.OracleRequest.ClientID),
		)
	}

	return nil
}

func migrateAppLinkResult(r *v2.Result) *types.Result {
	if r == nil {
		return nil
	}

	switch result := (r.Sum).(type) {
	case *v2.Result_Success_:
		return types.NewSuccessResult(result.Success.Value, result.Success.Signature)
	case *v2.Result_Failed_:
		return types.NewErrorResult(result.Failed.Error)
	default:
		panic(fmt.Errorf("invalid result type"))
	}
}

// migrateParams add the AppLinksParams to the params set
func migrateParams(ctx sdk.Context, subspace paramstypes.Subspace, legacyAmino *codec.LegacyAmino) error {
	var v2NicknameParams v2.NicknameParams
	err := legacyAmino.UnmarshalJSON(subspace.GetRaw(ctx, types.NicknameParamsKey), &v2NicknameParams)
	if err != nil {
		return err
	}

	var v2DTagParams v2.DTagParams
	err = legacyAmino.UnmarshalJSON(subspace.GetRaw(ctx, types.DTagParamsKey), &v2DTagParams)
	if err != nil {
		return err
	}

	var v2OracleParams v2.OracleParams
	err = legacyAmino.UnmarshalJSON(subspace.GetRaw(ctx, types.OracleParamsKey), &v2OracleParams)
	if err != nil {
		return err
	}

	var v2BioParams v2.BioParams
	err = legacyAmino.UnmarshalJSON(subspace.GetRaw(ctx, types.BioParamsKey), &v2BioParams)
	if err != nil {
		return err
	}

	nicknameParams := types.NewNicknameParams(v2NicknameParams.MinLength, v2NicknameParams.MaxLength)
	dtagParams := types.NewDTagParams(v2DTagParams.RegEx, v2DTagParams.MinLength, v2DTagParams.MaxLength)
	bioParams := types.NewBioParams(v2BioParams.MaxLength)
	oracleParams := types.NewOracleParams(
		v2OracleParams.ScriptID,
		v2OracleParams.AskCount,
		v2OracleParams.MinCount,
		v2OracleParams.PrepareGas,
		v2OracleParams.ExecuteGas,
		v2OracleParams.FeeAmount...,
	)

	subspace.Set(ctx, types.NicknameParamsKey, &nicknameParams)
	subspace.Set(ctx, types.DTagParamsKey, &dtagParams)
	subspace.Set(ctx, types.BioParamsKey, &bioParams)
	subspace.Set(ctx, types.OracleParamsKey, &oracleParams)

	return nil
}
