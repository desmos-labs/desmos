package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*ReactionValue)(nil), nil)
	cdc.RegisterConcrete(&RegisteredReactionValue{}, "desmos/RegisteredReactionValue", nil)
	cdc.RegisterConcrete(&FreeTextValue{}, "desmos/FreeTextValue", nil)

	cdc.RegisterConcrete(MsgAddReaction{}, "desmos/MsgAddReaction", nil)
	cdc.RegisterConcrete(MsgRemoveReaction{}, "desmos/MsgRemoveReaction", nil)
	cdc.RegisterConcrete(MsgAddRegisteredReaction{}, "desmos/MsgAddRegisteredReaction", nil)
	cdc.RegisterConcrete(MsgEditRegisteredReaction{}, "desmos/MsgEditRegisteredReaction", nil)
	cdc.RegisterConcrete(MsgRemoveRegisteredReaction{}, "desmos/MsgRemoveRegisteredReaction", nil)
	cdc.RegisterConcrete(MsgSetReactionsParams{}, "desmos/MsgSetReactionsParams", nil)
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
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
