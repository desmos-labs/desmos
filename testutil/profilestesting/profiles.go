package profilestesting

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32/legacybech32"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/gogo/protobuf/proto"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func NewAny(value proto.Message) *codectypes.Any {
	any, err := codectypes.NewAnyWithValue(value)
	if err != nil {
		panic(err)
	}
	return any
}

func AssertNoProfileError(profile *types.Profile, err error) *types.Profile {
	if err != nil {
		panic(err)
	}
	return profile
}

func AccountFromAddr(addr string) sdk.AccountI {
	address, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		panic(err)
	}
	return authtypes.NewBaseAccountWithAddress(address)
}

func PubKeyFromBech32(pubKey string) cryptotypes.PubKey {
	publicKey, err := legacybech32.UnmarshalPubKey(legacybech32.AccPK, pubKey)
	if err != nil {
		panic(err)
	}
	return publicKey
}

func PubKeyFromJSON(cdc codec.Codec, pubKey string) cryptotypes.PubKey {
	var publicKey cryptotypes.PubKey
	err := cdc.UnmarshalInterfaceJSON([]byte(pubKey), &publicKey)
	if err != nil {
		panic(err)
	}
	return publicKey
}

type ProfileOption func(*types.Profile) *types.Profile

func ProfileFromAddr(address string, options ...func(*types.Profile) *types.Profile) *types.Profile {
	profile, err := types.NewProfile(
		fmt.Sprintf("%s-dtag", address),
		"",
		"",
		types.NewPictures("", ""),
		time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		AccountFromAddr(address),
	)
	if err != nil {
		panic(err)
	}

	for _, option := range options {
		profile = option(profile)
	}

	return profile
}

func WithNextAccountNumber(ctx context.Context, ak authkeeper.AccountKeeper) ProfileOption {
	return func(profile *types.Profile) *types.Profile {
		profile.SetAccountNumber(ak.NextAccountNumber(ctx))
		return profile
	}
}

// SingleSignatureFromHex convert the hex-encoded string of the single signature to CosmosSignatureData
func SingleSignatureFromHex(hexEncodedSignature string) types.Signature {
	sig, err := hex.DecodeString(hexEncodedSignature)
	if err != nil {
		panic(err)
	}
	return types.NewSingleSignature(types.SIGNATURE_VALUE_TYPE_RAW, sig)
}

// MultiCosmosSignatureFromHex convert the hex-encoded string of the MultiSignature Any value to CosmosSignatureData
func MultiCosmosSignatureFromHex(unpacker codectypes.AnyUnpacker, hexEncodedSignatureData string) types.Signature {
	sig, err := hex.DecodeString(hexEncodedSignatureData)
	if err != nil {
		panic(err)
	}

	var multisigAny codectypes.Any
	err = multisigAny.Unmarshal(sig)
	if err != nil {
		panic(err)
	}

	var signature types.Signature
	if err = unpacker.UnpackAny(&multisigAny, &signature); err != nil {
		panic(err)
	}
	return signature
}
