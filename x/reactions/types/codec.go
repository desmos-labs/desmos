package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*ReactionValue)(nil), nil)
	cdc.RegisterConcrete(&RegisteredReactionValue{}, "desmos/RegisteredReactionValue", nil)
	cdc.RegisterConcrete(&FreeTextValue{}, "desmos/FreeTextValue", nil)

	legacy.RegisterAminoMsg(cdc, &MsgAddReaction{}, "desmos/MsgAddReaction")
	legacy.RegisterAminoMsg(cdc, &MsgRemoveReaction{}, "desmos/MsgRemoveReaction")
	legacy.RegisterAminoMsg(cdc, &MsgAddRegisteredReaction{}, "desmos/MsgAddRegisteredReaction")
	legacy.RegisterAminoMsg(cdc, &MsgEditRegisteredReaction{}, "desmos/MsgEditRegisteredReaction")
	legacy.RegisterAminoMsg(cdc, &MsgRemoveRegisteredReaction{}, "desmos/MsgRemoveRegisteredReaction")
	legacy.RegisterAminoMsg(cdc, &MsgSetReactionsParams{}, "desmos/MsgSetReactionsParams")
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"desmos.reactions.v1.ReactionValue",
		(*ReactionValue)(nil),
		&RegisteredReactionValue{},
		&FreeTextValue{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddReaction{},
		&MsgRemoveReaction{},
		&MsgAddRegisteredReaction{},
		&MsgEditRegisteredReaction{},
		&MsgRemoveRegisteredReaction{},
		&MsgSetReactionsParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// AminoCdc references the global x/reactions module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/reactions and
	// defined at the application level.
	AminoCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
