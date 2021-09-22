package v200

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	v100 "github.com/desmos-labs/desmos/x/profiles/legacy/v100"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v1.0 to v2.0 The
// migration includes:
//
// - Change OracleRequest to have uint64 ID instead of int64
// - Change OracleRequest to have uint64 OracleScriptID instead of int64
// - Change OracleParams to have uint64 ScriptID instead of int64
// - Change OracleParams to remove FeePayer
func MigrateStore(
	ctx sdk.Context, storeKey sdk.StoreKey, subspace paramstypes.Subspace,
	cdc codec.BinaryCodec, legacyAmino *codec.LegacyAmino,
) error {
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

	var newLinks []types.ApplicationLink
	for ; iterator.Valid(); iterator.Next() {
		var v1ApplicationLink v100.ApplicationLink
		err := cdc.Unmarshal(iterator.Value(), &v1ApplicationLink)
		if err != nil {
			return err
		}

		newLinks = append(newLinks, types.NewApplicationLink(
			v1ApplicationLink.User,
			types.NewData(v1ApplicationLink.Data.Application, v1ApplicationLink.Data.Username),
			types.ApplicationLinkState(v1ApplicationLink.State),
			types.NewOracleRequest(
				uint64(v1ApplicationLink.OracleRequest.ID),
				uint64(v1ApplicationLink.OracleRequest.OracleScriptID),
				types.NewOracleRequestCallData(
					v1ApplicationLink.OracleRequest.CallData.Application,
					v1ApplicationLink.OracleRequest.CallData.CallData,
				),
				v1ApplicationLink.OracleRequest.ClientID,
			),
			migrateAppLinkResult(v1ApplicationLink.Result),
			v1ApplicationLink.CreationTime,
		))
		store.Delete(iterator.Key())
	}

	for _, link := range newLinks {
		store.Set(
			types.UserApplicationLinkKey(link.User, link.Data.Application, link.Data.Username),
			types.MustMarshalApplicationLink(cdc, link),
		)
	}

	return nil
}

func migrateAppLinkResult(r *v100.Result) *types.Result {
	if r == nil {
		return nil
	}

	switch result := (r.Sum).(type) {
	case *v100.Result_Success_:
		return types.NewSuccessResult(result.Success.Value, result.Success.Signature)
	case *v100.Result_Failed_:
		return types.NewErrorResult(result.Failed.Error)
	default:
		panic(fmt.Errorf("invalid result type"))
	}
}

// migrateParams migrates the OracleParams by removing the FeePayer field
// and converting the ScriptID from int64 to uint64
func migrateParams(ctx sdk.Context, subspace paramstypes.Subspace, legacyAmino *codec.LegacyAmino) error {
	var v1Params v100.OracleParams
	err := legacyAmino.UnmarshalJSON(subspace.GetRaw(ctx, types.OracleParamsKey), &v1Params)
	if err != nil {
		return err
	}

	var v2Params = types.NewOracleParams(
		uint64(v1Params.ScriptID),
		v1Params.AskCount,
		v1Params.MinCount,
		v1Params.PrepareGas,
		v1Params.ExecuteGas,
		v1Params.FeeAmount...,
	)
	subspace.Set(ctx, types.OracleParamsKey, &v2Params)
	return nil
}
