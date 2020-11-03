package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgCreatePost{}, "desmos/MsgCreatePost", nil)
	cdc.RegisterConcrete(MsgEditPost{}, "desmos/MsgEditPost", nil)
	cdc.RegisterConcrete(MsgAddPostReaction{}, "desmos/MsgAddPostReaction", nil)
	cdc.RegisterConcrete(MsgRemovePostReaction{}, "desmos/MsgRemovePostReaction", nil)
	cdc.RegisterConcrete(MsgAnswerPoll{}, "desmos/MsgAnswerPoll", nil)
	cdc.RegisterConcrete(MsgRegisterReaction{}, "desmos/MsgRegisterReaction", nil)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/posts module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/posts and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
