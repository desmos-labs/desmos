package utils

import (
	"strings"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAccAddressFromExternal is a utitity in Magpie to get a Desmo address from external namespace
func GetAccAddressFromExternal(address string, bech32Prefix string) (sdk.AccAddress, error){
	if len(strings.TrimSpace(address)) == 0 {
		return sdk.AccAddress{}, nil
	}

	bz, err := sdk.GetFromBech32(address, bech32Prefix)
	if err != nil{
		return nil, err
	}

	err = sdk.VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return sdk.AccAddress(bz), nil
}