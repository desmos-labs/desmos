package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreatePost{}, "dwitter/MsgCreatePost", nil)
	cdc.RegisterConcrete(MsgEditPost{}, "dwitter/MsgEditPost", nil)
	cdc.RegisterConcrete(MsgLike{}, "dwitter/MsgLike", nil)
	cdc.RegisterConcrete(MsgUnlike{}, "dwitter/MsgUnlike", nil)
}
