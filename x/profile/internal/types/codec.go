package types

import "github.com/cosmos/cosmos-sdk/codec"

// ModuleCdc is the codec
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateAccount{}, "desmos/MsgCreateAccount", nil)
	cdc.RegisterConcrete(MsgEditAccount{}, "desmos/MsgEditAccount", nil)
	cdc.RegisterConcrete(MsgDeleteAccount{}, "desmos/MsgDeleteAccount", nil)
}
