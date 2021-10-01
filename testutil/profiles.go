package testutil

import (
	"fmt"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32/legacybech32"
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
