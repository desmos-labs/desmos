package models

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModelsCdc is the codec
var ModelsCdc = codec.New()

func init() {
	RegisterModelsCodec(ModelsCdc)
}

// RegisterModelsCodec registers concrete types on the Amino codec
func RegisterModelsCodec(cdc *codec.Codec) {}
