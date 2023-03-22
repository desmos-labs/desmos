package v4

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	v4types "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v4/types"
	v5types "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v5/types"
	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v4 to v5.
// The migration includes:
//
// - migrating all the profiles to have the proper Protobuf type
// - delete all the relationships
// - delete all the user blocks
//
// NOTE: This method must be called AFTER the migration from v0 to v1 of the x/relationships module.
//
//	If this order is not preserved, all relationships and blocks WILL BE DELETED.
func MigrateStore(ctx sdk.Context, ak authkeeper.AccountKeeper, storeKey sdk.StoreKey, amino *codec.LegacyAmino, cdc codec.BinaryCodec) error {
	legacyKeeper := NewKeeper(storeKey, cdc)

	// Migrate the profiles
	err := migrateProfiles(ctx, ak)
	if err != nil {
		return err
	}

	// Migrate all the data to use the new keys
	migrateDTags(ctx, legacyKeeper, storeKey)
	migrateDTagTransferRequests(ctx, legacyKeeper, storeKey, cdc)
	migrateApplicationLinks(ctx, legacyKeeper, storeKey, cdc)

	// Migrate the chain links
	err = migrateChainLinks(ctx, legacyKeeper, storeKey, amino, cdc)
	if err != nil {
		return err
	}

	legacyKeeper.DeleteRelationships(ctx)
	legacyKeeper.DeleteBlocks(ctx)
	return nil
}

func migrateProfiles(ctx sdk.Context, ak authkeeper.AccountKeeper) error {
	var profiles []*v4types.Profile
	ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {
		if profile, ok := account.(*v4types.Profile); ok {
			profiles = append(profiles, profile)
		}
		return false
	})

	for _, profile := range profiles {
		// Convert the profile
		v5Profile, err := v5types.NewProfile(
			profile.DTag,
			profile.Nickname,
			profile.Bio,
			v5types.NewPictures(profile.Pictures.Profile, profile.Pictures.Cover),
			profile.CreationDate,
			profile.GetAccount(),
		)
		if err != nil {
			return err
		}

		// Set the account
		ak.SetAccount(ctx, v5Profile)
	}

	return nil
}

func migrateDTags(ctx sdk.Context, k Keeper, storeKey sdk.StoreKey) {
	dTags := map[string][]byte{}
	k.IterateDTags(ctx, func(index int64, dTag string, value []byte) (stop bool) {
		dTags[dTag] = value
		return false
	})

	store := ctx.KVStore(storeKey)
	for dTag, address := range dTags {
		// Delete the old key
		store.Delete(v4types.DTagStoreKey(dTag))

		// Store the DTag using the new key
		store.Set(types.DTagStoreKey(dTag), address)
	}
}

func migrateDTagTransferRequests(ctx sdk.Context, k Keeper, storeKey sdk.StoreKey, cdc codec.BinaryCodec) {
	var requests []v4types.DTagTransferRequest
	k.IterateDTagTransferRequests(ctx, func(index int64, request v4types.DTagTransferRequest) (stop bool) {
		requests = append(requests, request)
		return false
	})

	store := ctx.KVStore(storeKey)
	for i, request := range requests {
		// Delete the old key
		store.Delete(v4types.DTagTransferRequestStoreKey(request.Sender, request.Receiver))

		// Store the request using the new key
		store.Set(
			types.DTagTransferRequestStoreKey(request.Sender, request.Receiver),
			cdc.MustMarshal(&requests[i]),
		)
	}
}

func migrateApplicationLinks(ctx sdk.Context, k Keeper, storeKey sdk.StoreKey, cdc codec.BinaryCodec) {
	store := ctx.KVStore(storeKey)

	var applicationLinks []v4types.ApplicationLink
	k.IterateApplicationLinks(ctx, func(index int64, applicationLink v4types.ApplicationLink) (stop bool) {
		applicationLinks = append(applicationLinks, applicationLink)
		return false
	})

	var clientIDKeys [][]byte
	k.IterateApplicationLinkClientIDKeys(ctx, func(index int64, key []byte, value []byte) (stop bool) {
		clientIDKeys = append(clientIDKeys, key)
		return false
	})

	// Delete all the client ID keys to make sure we remove leftover ones
	// The new client ID keys will be set when storing the application links anyway later
	for _, key := range clientIDKeys {
		store.Delete(key)
	}

	for i, link := range applicationLinks {
		// Delete the old keys
		store.Delete(v4types.UserApplicationLinkKey(link.User, link.Data.Application, link.Data.Username))

		// Store the link with the new key
		linkKey := types.UserApplicationLinkKey(link.User, link.Data.Application, link.Data.Username)
		store.Set(linkKey, cdc.MustMarshal(&applicationLinks[i]))
		store.Set(types.ApplicationLinkClientIDKey(link.OracleRequest.ClientID), linkKey)
	}
}

func migrateChainLinks(ctx sdk.Context, k Keeper, storeKey sdk.StoreKey, amino *codec.LegacyAmino, cdc codec.BinaryCodec) error {
	var chainLinks []v4types.ChainLink
	k.IterateChainLinks(ctx, func(index int64, chainLink v4types.ChainLink) (stop bool) {
		chainLinks = append(chainLinks, chainLink)
		return false
	})

	store := ctx.KVStore(storeKey)
	for _, link := range chainLinks {
		var address v4types.AddressData
		err := cdc.UnpackAny(link.Address, &address)
		if err != nil {
			return err
		}

		// Convert the address data
		var addressData types.AddressData
		switch address := address.(type) {
		case *v4types.Bech32Address:
			addressData = v5types.NewBech32Address(address.Value, address.Prefix)
		case *v4types.Base58Address:
			addressData = v5types.NewBase58Address(address.Value)
		case *v4types.HexAddress:
			addressData = v5types.NewHexAddress(address.Value, address.Prefix)
		default:
			panic(fmt.Errorf("unsupported AddressData type: %T", link.Address))
		}

		var pubKey cryptotypes.PubKey
		err = cdc.UnpackAny(link.Proof.PubKey, &pubKey)
		if err != nil {
			return err
		}

		plainText, err := hex.DecodeString(link.Proof.PlainText)
		if err != nil {
			return err
		}

		signature, err := hex.DecodeString(link.Proof.Signature)
		if err != nil {
			return err
		}

		// Get the proper signing method used
		var directSignBytes tx.SignDoc
		var aminoSignBytes legacytx.StdSignDoc

		signMode := signing.SignMode_SIGN_MODE_TEXTUAL
		err = amino.UnmarshalJSON(plainText, &aminoSignBytes)
		if err == nil {
			signMode = signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
		}

		err = cdc.Unmarshal(plainText, &directSignBytes)
		if err == nil {
			signMode = signing.SignMode_SIGN_MODE_DIRECT
		}

		// Build the new chain link
		v5Link := v5types.NewChainLink(
			link.User,
			addressData,
			v5types.NewProof(
				pubKey,
				&v5types.SingleSignatureData{
					Signature: signature,
					Mode:      signMode,
				},
				link.Proof.PlainText,
			),
			v5types.NewChainConfig(
				link.ChainConfig.Name,
			),
			link.CreationTime,
		)

		// Delete the old key
		store.Delete(v4types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, addressData.GetValue()))

		// Store the chain link using the new key
		store.Set(types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, addressData.GetValue()), cdc.MustMarshal(&v5Link))
	}

	return nil
}
