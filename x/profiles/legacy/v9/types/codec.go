package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*AddressData)(nil), nil)
	cdc.RegisterConcrete(&Bech32Address{}, "desmos/Bech32Address", nil)
	cdc.RegisterConcrete(&Base58Address{}, "desmos/Base58Address", nil)
	cdc.RegisterConcrete(&HexAddress{}, "desmos/HexAddress", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*authtypes.AccountI)(nil), &Profile{})
	registry.RegisterImplementations((*exported.VestingAccount)(nil), &Profile{})
	registry.RegisterInterface(
		"desmos.profiles.v3.AddressData",
		(*AddressData)(nil),
		&Bech32Address{},
		&Base58Address{},
		&HexAddress{},
	)
	registry.RegisterInterface(
		"desmos.profiles.v3.Signature",
		(*Signature)(nil),
		&SingleSignature{},
		&CosmosMultiSignature{},
	)
}
