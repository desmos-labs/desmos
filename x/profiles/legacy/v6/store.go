package v6

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	v5types "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v5/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v6 to v7.
// The migration includes:
//
// - migrating all the profiles to have the proper Protobuf type
// - add the expiration date to all application links
func MigrateStore(ctx sdk.Context, ak authkeeper.AccountKeeper, storeKey sdk.StoreKey, amino *codec.LegacyAmino, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	// Migrate the profiles
	err := migrateProfiles(ctx, ak)
	if err != nil {
		return err
	}

	// Migrate all the application links
	err = migrateApplicationLinks(store, cdc)
	if err != nil {
		return err
	}

	// Migrate all the chain links
	err = migrateChainLinks(store, cdc)
	if err != nil {
		return err
	}

	return nil
}

// migrateProfiles migrates the profiles from v5 to v7 to properly update the Protobuf name.
// The migration from v5 to v6 is skipped because the two types are identical (from v5 to v6 no changes were made).
func migrateProfiles(ctx sdk.Context, ak authkeeper.AccountKeeper) error {
	var profiles []*v5types.Profile
	ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {
		if profile, ok := account.(*v5types.Profile); ok {
			profiles = append(profiles, profile)
		}
		return false
	})

	for _, profile := range profiles {
		// Convert the profile
		v7Profile, err := types.NewProfile(
			profile.DTag,
			profile.Nickname,
			profile.Bio,
			types.NewPictures(profile.Pictures.Profile, profile.Pictures.Cover),
			profile.CreationDate,
			profile.GetAccount(),
		)
		if err != nil {
			return err
		}

		// Set the account
		ak.SetAccount(ctx, v7Profile)
	}

	return nil
}

// migrateApplicationLinks migrates the application links from v5 to v7 adding the expiration date properly.
// The migration from v5 to v6 is skipped because the two types are identical (from v5 to v6 no changes were made).
func migrateApplicationLinks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	appLinksStore := prefix.NewStore(store, types.ApplicationLinkPrefix)
	iterator := appLinksStore.Iterator(nil, nil)

	var applicationLinks []v5types.ApplicationLink
	for ; iterator.Valid(); iterator.Next() {
		var applicationLink v5types.ApplicationLink
		err := cdc.Unmarshal(iterator.Value(), &applicationLink)
		if err != nil {
			return err
		}
		applicationLinks = append(applicationLinks, applicationLink)
	}

	for _, v5Link := range applicationLinks {
		// Migrate the link
		v7Link := types.NewApplicationLink(
			v5Link.User,
			convertApplicationLinkData(v5Link.Data),
			convertApplicationLinkState(v5Link.State),
			convertApplicationLinkOracleRequest(v5Link.OracleRequest),
			convertApplicationLinkResult(v5Link.Result),
			v5Link.CreationTime,
			v5Link.CreationTime.Add(types.DefaultAppLinksValidityDuration),
		)

		// Store the application link
		userApplicationLinkKey := types.UserApplicationLinkKey(v7Link.User, v7Link.Data.Application, v7Link.Data.Username)
		store.Set(userApplicationLinkKey, cdc.MustMarshal(&v7Link))

		// Store the expiration time
		applicationLinkExpiringTimeKey := types.ApplicationLinkExpiringTimeKey(v7Link.ExpirationTime, v7Link.OracleRequest.ClientID)
		store.Set(applicationLinkExpiringTimeKey, []byte(v7Link.OracleRequest.ClientID))
	}

	return nil
}

func convertApplicationLinkData(v5Data v5types.Data) types.Data {
	return types.NewData(v5Data.Application, v5Data.Username)
}

func convertApplicationLinkState(v5State v5types.ApplicationLinkState) types.ApplicationLinkState {
	switch v5State {
	case v5types.ApplicationLinkStateInitialized:
		return types.ApplicationLinkStateInitialized
	case v5types.AppLinkStateVerificationStarted:
		return types.AppLinkStateVerificationStarted
	case v5types.AppLinkStateVerificationError:
		return types.AppLinkStateVerificationError
	case v5types.AppLinkStateVerificationSuccess:
		return types.AppLinkStateVerificationSuccess
	case v5types.AppLinkStateVerificationTimedOut:
		return types.AppLinkStateVerificationTimedOut
	default:
		panic(fmt.Errorf("invalid application link state: %s", v5State))
	}
}

func convertApplicationLinkOracleRequest(v5Request v5types.OracleRequest) types.OracleRequest {
	return types.NewOracleRequest(
		v5Request.ID,
		v5Request.OracleScriptID,
		types.NewOracleRequestCallData(v5Request.CallData.Application, v5Request.CallData.CallData),
		v5Request.ClientID,
	)
}

func convertApplicationLinkResult(v5Result *v5types.Result) *types.Result {
	if v5Result == nil {
		return nil
	}

	switch result := v5Result.Sum.(type) {
	case *v5types.Result_Success_:
		return types.NewSuccessResult(result.Success.Value, result.Success.Signature)

	case *v5types.Result_Failed_:
		return types.NewErrorResult(result.Failed.Error)

	default:
		panic(fmt.Errorf("invalid result type: %T", v5Result.Sum))
	}
}

// migrateChainLinks migrates the chain links from v5 to v7 by changing the various Protobuf interface types.
// The migration from v5 to v6 is skipped because the two types are identical (from v5 to v6 no changes were made).
func migrateChainLinks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	appLinksStore := prefix.NewStore(store, types.ChainLinksPrefix)
	iterator := appLinksStore.Iterator(nil, nil)

	var applicationLinks []v5types.ChainLink
	for ; iterator.Valid(); iterator.Next() {
		var applicationLink v5types.ChainLink
		err := cdc.Unmarshal(iterator.Value(), &applicationLink)
		if err != nil {
			return err
		}
		applicationLinks = append(applicationLinks, applicationLink)
	}

	for _, v5Link := range applicationLinks {
		// Migrate the link
		v7Link := types.NewChainLink(
			v5Link.User,
			convertChainLinkAddressData(v5Link.GetAddressData()),
			convertChainLinkProof(v5Link.Proof, cdc),
			types.NewChainConfig(v5Link.ChainConfig.Name),
			v5Link.CreationTime,
		)

		// Store the chain link
		userApplicationLinkKey := types.ChainLinksStoreKey(v7Link.User, v7Link.ChainConfig.Name, v7Link.GetAddressData().GetValue())
		store.Set(userApplicationLinkKey, cdc.MustMarshal(&v7Link))
	}

	return nil
}

func convertChainLinkAddressData(v5Signature v5types.AddressData) types.AddressData {
	switch address := v5Signature.(type) {
	case *v5types.Bech32Address:
		return types.NewBech32Address(address.Value, address.Prefix)
	case *v5types.Base58Address:
		return types.NewBase58Address(address.Value)
	case *v5types.HexAddress:
		return types.NewHexAddress(address.Value, address.Prefix)
	default:
		panic(fmt.Errorf("invalid signature type: %T", v5Signature))
	}
}

func convertChainLinkProof(v5Proof v5types.Proof, cdc codec.BinaryCodec) types.Proof {
	var pubKey cryptotypes.PubKey
	err := cdc.UnpackAny(v5Proof.PubKey, &pubKey)
	if err != nil {
		panic(err)
	}

	var v6Signature types.SignatureData
	v6SignatureAny := convertChainLinkSignatureData(v5Proof.Signature, cdc)
	err = cdc.UnpackAny(v6SignatureAny, &v6Signature)
	if err != nil {
		panic(err)
	}

	return types.NewProof(pubKey, v6Signature, v5Proof.PlainText)

}

func convertChainLinkSignatureData(data *codectypes.Any, cdc codec.BinaryCodec) *codectypes.Any {
	var v5Signature v5types.SignatureData
	err := cdc.UnpackAny(data, &v5Signature)
	if err != nil {
		panic(err)
	}

	var signatureAny *codectypes.Any
	switch signature := v5Signature.(type) {
	case *v5types.SingleSignatureData:
		v6Signature := &types.SingleSignatureData{Signature: signature.Signature, Mode: signature.Mode}
		signatureAny, err = codectypes.NewAnyWithValue(v6Signature)
		if err != nil {
			panic(err)
		}

	case *v5types.MultiSignatureData:
		sigsAnys := make([]*codectypes.Any, len(signature.Signatures))
		for i, signature := range signature.Signatures {
			sigsAnys[i] = convertChainLinkSignatureData(signature, cdc)
		}

		v6Signature := &types.MultiSignatureData{BitArray: signature.BitArray, Signatures: sigsAnys}
		signatureAny, err = codectypes.NewAnyWithValue(v6Signature)
		if err != nil {
			panic(err)
		}

	}

	return signatureAny

}
