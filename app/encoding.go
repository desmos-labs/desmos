package app

import (
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/desmos-labs/desmos/v4/app/params"
)

// MakeTestEncodingConfig creates an EncodingConfig for testing
func MakeTestEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
