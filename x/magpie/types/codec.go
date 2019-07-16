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
	cdc.RegisterConcrete(MsgCreatePost{}, "magpie/MsgCreatePost", nil)
	cdc.RegisterConcrete(MsgEditPost{}, "magpie/MsgEditPost", nil)
	cdc.RegisterConcrete(MsgLike{}, "magpie/MsgLike", nil)
	cdc.RegisterConcrete(MsgUnlike{}, "magpie/MsgUnlike", nil)
	cdc.RegisterConcrete(MsgCreateSession{}, "magpie/MsgCreateSession", nil)
}
