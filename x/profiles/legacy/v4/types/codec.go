package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
)

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.AccountI)(nil), &Profile{})
	registry.RegisterImplementations((*exported.VestingAccount)(nil), &Profile{})
	registry.RegisterInterface(
		"desmos.profiles.v1beta1.AddressData",
		(*AddressData)(nil),
		&Bech32Address{},
		&Base58Address{},
		&HexAddress{},
	)
}
