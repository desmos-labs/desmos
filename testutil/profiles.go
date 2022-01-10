package testutil

import (
	"encoding/hex"
	"fmt"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32/legacybech32"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gogo/protobuf/proto"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
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

func GeneratePubKeyAndMultiSignatureData(n int, msg []byte) (cryptotypes.PubKey, *types.MultiSignatureData) {
	pubKeys := make([]cryptotypes.PubKey, n)
	cosmosMultisig := multisig.NewMultisig(n)
	for i := 0; i < n; i++ {
		privkey := secp256k1.GenPrivKey()
		pubKeys[i] = privkey.PubKey()
		sig, _ := privkey.Sign(msg)
		sigData := &signing.SingleSignatureData{Signature: sig}
		multisig.AddSignatureFromPubKey(cosmosMultisig, sigData, pubKeys[i], pubKeys)
	}
	sigData := CosmosSignatureDataToDesmosSignatureData(cosmosMultisig)
	return kmultisig.NewLegacyAminoPubKey(n, pubKeys), sigData.(*types.MultiSignatureData)
}

func CosmosSignatureDataToDesmosSignatureData(data signing.SignatureData) types.SignatureData {
	switch data := data.(type) {
	case *signing.SingleSignatureData:
		return &types.SingleSignatureData{
			Mode:      data.SignMode,
			Signature: data.Signature,
		}
	case *signing.MultiSignatureData:
		sigAnys := make([]*codectypes.Any, len(data.Signatures))
		for i, data := range data.Signatures {
			sigAny, _ := codectypes.NewAnyWithValue(CosmosSignatureDataToDesmosSignatureData(data))
			sigAnys[i] = sigAny
		}
		return &types.MultiSignatureData{
			BitArray:   data.BitArray,
			Signatures: sigAnys,
		}
	default:
		panic(fmt.Errorf("unexpected case %+v", data))
	}
}
