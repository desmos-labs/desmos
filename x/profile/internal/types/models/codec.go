package models

import "github.com/cosmos/cosmos-sdk/codec"

// ModelsCdc is the codec
var ModelsCdc = codec.New()

func init() {
	RegisterModelsCodec(ModelsCdc)
}

// RegisterModelsCodec registers concrete types on the Amino codec
func RegisterModelsCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(NameSurnameLengths{}, "desmos/NameSurnameLengths", nil)
	cdc.RegisterConcrete(MonikerLengths{}, "desmos/MonikerLengths", nil)
	cdc.RegisterConcrete(BiographyLengths{}, "desmos/BiographyLengths", nil)
	cdc.RegisterConcrete(EditNameSurnameParamsProposal{}, "desmos/EditNameSurnameParamsProposal", nil)
	cdc.RegisterConcrete(EditMonikerParamsProposal{}, "desmos/EditMonikerParamsProposal", nil)
	cdc.RegisterConcrete(EditBioParamsProposal{}, "desmos/EditBioParamsProposal", nil)
}
