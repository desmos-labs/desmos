package v9

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	v9types "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v9/types"
	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v9 to v10.
// The migration includes:
//
// - migrating all the profiles to have the proper Protobuf type
// - migrating all the chain links to the new version chain links
func MigrateStore(ctx sdk.Context, ak authkeeper.AccountKeeper, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	// Migrate the profiles
	err := migrateProfiles(ctx, ak)
	if err != nil {
		return err
	}

	// Migrate the chain links
	err = migrateChainLinks(ctx.KVStore(storeKey), cdc)
	if err != nil {
		return err
	}

	return nil
}

func migrateProfiles(ctx sdk.Context, ak authkeeper.AccountKeeper) error {
	var profiles []*v9types.Profile
	ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {
		if profile, ok := account.(*v9types.Profile); ok {
			profiles = append(profiles, profile)
		}
		return false
	})

	for _, profile := range profiles {
		// Convert the profile
		v10Profile, err := types.NewProfile(
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
		ak.SetAccount(ctx, v10Profile)
	}

	return nil
}

// migrateChainLinks migrates the chain links from v9 to v10 by updating their address field accordingly.
// The migration includes:
//
// - migrating all the chain link to new version structure
// - migrating all the proof signature to have the proper Protobuf type
func migrateChainLinks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	chainLinksStore := prefix.NewStore(store, types.ChainLinksPrefix)
	iterator := chainLinksStore.Iterator(nil, nil)
	defer iterator.Close()

	var chainLinks []v9types.ChainLink
	for ; iterator.Valid(); iterator.Next() {
		var chainLink v9types.ChainLink
		err := cdc.Unmarshal(iterator.Value(), &chainLink)
		if err != nil {
			return err
		}
		chainLinks = append(chainLinks, chainLink)
	}

	for _, v9Link := range chainLinks {
		// Migrate the link
		v10Link := types.NewChainLink(
			v9Link.User,
			convertChainLinkAddress(v9Link.GetAddressData()),
			convertChainLinkProof(v9Link.Proof, cdc),
			types.NewChainConfig(v9Link.ChainConfig.Name),
			v9Link.CreationTime,
		)

		// Store the chain link
		chainLinkKey := types.ChainLinksStoreKey(v10Link.User, v10Link.ChainConfig.Name, v10Link.Address.Value)
		store.Set(chainLinkKey, cdc.MustMarshal(&v10Link))
	}
	return nil
}

func convertChainLinkAddress(v9AddressData v9types.AddressData) types.Address {
	switch address := v9AddressData.(type) {
	case *v9types.Bech32Address:
		return types.NewAddress(address.Value, types.GENERATION_ALGORITHM_COSMOS, types.NewBech32Encoding(address.Prefix))
	case *v9types.Base58Address:
		return types.NewAddress(address.Value, types.GENERATION_ALGORITHM_DO_NOTHING, types.NewBase58Encoding(""))
	case *v9types.HexAddress:
		return types.NewAddress(address.Value, types.GENERATION_ALGORITHM_EVM, types.NewHexEncoding(address.Prefix, false))
	default:
		panic(fmt.Errorf("invalid address data type: %T", v9AddressData))
	}
}

func convertChainLinkProof(v9Proof v9types.Proof, cdc codec.BinaryCodec) types.Proof {
	var pubKey cryptotypes.PubKey
	err := cdc.UnpackAny(v9Proof.PubKey, &pubKey)
	if err != nil {
		panic(err)
	}

	var v10Signature types.Signature
	v10SignatureAny := convertChainLinkSignatureData(v9Proof.Signature, cdc)
	err = cdc.UnpackAny(v10SignatureAny, &v10Signature)
	if err != nil {
		panic(err)
	}

	return types.NewProof(pubKey, v10Signature, v9Proof.PlainText)
}

func convertChainLinkSignatureData(data *codectypes.Any, cdc codec.BinaryCodec) *codectypes.Any {
	var v9Signature v9types.Signature
	err := cdc.UnpackAny(data, &v9Signature)
	if err != nil {
		panic(err)
	}

	var signatureAny *codectypes.Any
	switch signature := v9Signature.(type) {
	case *v9types.SingleSignature:
		v10Signature := types.NewSingleSignature(types.SignatureValueType(signature.ValueType), signature.Signature)
		signatureAny, err = codectypes.NewAnyWithValue(v10Signature)
		if err != nil {
			panic(err)
		}

	case *v9types.CosmosMultiSignature:
		signatures := make([]types.Signature, len(signature.Signatures))
		for i, sig := range signature.Signatures {
			// Recursively convert the signature any
			sigAny := convertChainLinkSignatureData(sig, cdc)

			// Unpack the signature
			var cosmosSig types.Signature
			err = cdc.UnpackAny(sigAny, &cosmosSig)
			if err != nil {
				panic(err)
			}

			signatures[i] = cosmosSig
		}

		// Build the signature
		v10Signature := types.NewCosmosMultiSignature(signature.BitArray, signatures)

		// Convert it as an Any
		signatureAny, err = codectypes.NewAnyWithValue(v10Signature)
		if err != nil {
			panic(err)
		}
	}

	return signatureAny
}
