package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (r *QueryProfileResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if r.Profile != nil {
		var profile authtypes.AccountI
		return unpacker.UnpackAny(r.Profile, &profile)
	}
	return nil
}
