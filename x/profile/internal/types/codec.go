package types

import "github.com/cosmos/cosmos-sdk/codec"

// ModuleCdc is the codec
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateProfile{}, "desmos/MsgCreateProfile", nil)
	cdc.RegisterConcrete(MsgEditProfile{}, "desmos/MsgEditProfile", nil)
	cdc.RegisterConcrete(MsgDeleteProfile{}, "desmos/MsgDeleteProfile", nil)
}
