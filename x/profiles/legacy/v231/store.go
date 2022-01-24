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
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, subspace paramstypes.Subspace, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	migrateParams(ctx, subspace)

	err := migrateAppLinks(store, cdc)
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
func migrateParams(ctx sdk.Context, subspace paramstypes.Subspace) {
	var params v200.Params
	subspace.GetParamSet(ctx, &params)

	nicknameParams := types.NewNicknameParams(params.Nickname.MinLength, params.Nickname.MaxLength)
	dtagParams := types.NewDTagParams(params.DTag.RegEx, params.DTag.MinLength, params.DTag.MaxLength)
	bioParams := types.NewBioParams(params.Bio.MaxLength)
	oracleParams := types.NewOracleParams(
		params.Oracle.ScriptID,
		params.Oracle.AskCount,
		params.Oracle.MinCount,
		params.Oracle.PrepareGas,
		params.Oracle.ExecuteGas,
		params.Oracle.FeeAmount...,
	)

	subspace.Set(ctx, types.NicknameParamsKey, &nicknameParams)
	subspace.Set(ctx, types.DTagParamsKey, &dtagParams)
	subspace.Set(ctx, types.BioParamsKey, &bioParams)
	subspace.Set(ctx, types.OracleParamsKey, &oracleParams)
}
