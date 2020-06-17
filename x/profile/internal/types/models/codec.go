package models

import "github.com/cosmos/cosmos-sdk/codec"

// ModelsCdc is the codec
var ModelsCdc = codec.New()

func init() {
	RegisterModelsCodec(ModelsCdc)
}

// RegisterModelsCodec registers concrete types on the Amino codec
func RegisterModelsCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(NameSurnameLenParams{}, "desmos/NameSurnameLenParams", nil)
	cdc.RegisterConcrete(MonikerLenParams{}, "desmos/MonikerLenParams", nil)
	cdc.RegisterConcrete(BioLenParams{}, "desmos/BioLenParams", nil)
	cdc.RegisterConcrete(NameSurnameParamsEditProposal{}, "desmos/NameSurnameParamsEditProposal", nil)
	cdc.RegisterConcrete(MonikerParamsEditProposal{}, "desmos/MonikerParamsEditProposal", nil)
	cdc.RegisterConcrete(BioParamsEditProposal{}, "desmos/BioParamsEditProposal", nil)
}
