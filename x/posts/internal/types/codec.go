package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec used inside the whole posts module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreatePost{}, "desmos/MsgCreatePost", nil)
	cdc.RegisterConcrete(MsgEditPost{}, "desmos/MsgEditPost", nil)
	cdc.RegisterConcrete(MsgAddPostReaction{}, "desmos/MsgAddPostReaction", nil)
	cdc.RegisterConcrete(MsgRemovePostReaction{}, "desmos/MsgRemovePostReaction", nil)
	cdc.RegisterConcrete(MsgAnswerPoll{}, "desmos/MsgAnswerPoll", nil)
	cdc.RegisterConcrete(MsgRegisterReaction{}, "desmos/MsgRegisterReaction", nil)
}
