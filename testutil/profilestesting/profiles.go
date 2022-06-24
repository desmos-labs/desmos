package profilestesting

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32/legacybech32"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/gogo/protobuf/proto"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
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

func AccountFromAddr(addr string) authtypes.AccountI {
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

func ProfileFromAddr(address string) *types.Profile {
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

	return profile
}

// SingleSignatureProtoFromHex convert the hex-encoded string of the single signature to SignatureData
func SingleSignatureProtoFromHex(s string) types.SignatureData {
	sig, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return &types.SingleSignatureData{
		Mode:      signing.SignMode_SIGN_MODE_DIRECT,
		Signature: sig,
	}
}

// MultiSignatureProtoFromAnyHex convert the hex-encoded string of the MultiSignature Any value to SignatureData
func MultiSignatureProtoFromAnyHex(unpacker codectypes.AnyUnpacker, hexEncodedSignatureData string) types.SignatureData {
	sig, err := hex.DecodeString(hexEncodedSignatureData)
	if err != nil {
		panic(err)
	}

	var multisigAny codectypes.Any
	err = multisigAny.Unmarshal(sig)
	if err != nil {
		panic(err)
	}

	var sigData types.SignatureData
	if err = unpacker.UnpackAny(&multisigAny, &sigData); err != nil {
		panic(err)
	}
	return sigData
}
