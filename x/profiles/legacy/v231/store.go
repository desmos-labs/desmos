package v231

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	v200 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v200"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v2.3.1 to v2.4.0 The
// migration includes:
// - Add the AppLinkParams to the params set
// - Added expiration time to all app links
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, subspace paramstypes.Subspace, cdc codec.BinaryCodec,
	legacyAmino *codec.LegacyAmino) error {
	store := ctx.KVStore(storeKey)

	err := migrateParams(ctx, subspace, legacyAmino)
	if err != nil {
		return err
	}

	err = migrateAppLinks(store, cdc)
	if err != nil {
		return err
	}

	return nil
}

func migrateAppLinks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	iterator := sdk.KVStorePrefixIterator(store, types.UserApplicationLinkPrefix)
	defer iterator.Close()

	var keys [][]byte
	var newLinks []types.ApplicationLink
	for ; iterator.Valid(); iterator.Next() {
		var legacyAppLink v200.ApplicationLink
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

func migrateAppLinkResult(r *v200.Result) *types.Result {
	if r == nil {
		return nil
	}

	switch result := (r.Sum).(type) {
	case *v200.Result_Success_:
		return types.NewSuccessResult(result.Success.Value, result.Success.Signature)
	case *v200.Result_Failed_:
		return types.NewErrorResult(result.Failed.Error)
	default:
		panic(fmt.Errorf("invalid result type"))
	}
}

// migrateParams add the AppLinksParams to the params set
func migrateParams(ctx sdk.Context, subspace paramstypes.Subspace, legacyAmino *codec.LegacyAmino) error {
	var v2NicknameParams v200.NicknameParams
	err := legacyAmino.UnmarshalJSON(subspace.GetRaw(ctx, types.NicknameParamsKey), &v2NicknameParams)
	if err != nil {
		return err
	}

	var v2DTagParams v200.DTagParams
	err = legacyAmino.UnmarshalJSON(subspace.GetRaw(ctx, types.DTagParamsKey), &v2DTagParams)
	if err != nil {
		return err
	}

	var v2OracleParams v200.OracleParams
	err = legacyAmino.UnmarshalJSON(subspace.GetRaw(ctx, types.OracleParamsKey), &v2OracleParams)
	if err != nil {
		return err
	}

	var v2BioParams v200.BioParams
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
