package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcodec "github.com/cosmos/cosmos-sdk/x/gov/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*AttachmentContent)(nil), nil)
	cdc.RegisterConcrete(&Poll{}, "desmos/Poll", nil)
	cdc.RegisterConcrete(&Media{}, "desmos/Media", nil)

	cdc.RegisterConcrete(&MsgCreatePost{}, "desmos/MsgCreatePost", nil)
	cdc.RegisterConcrete(&MsgEditPost{}, "desmos/MsgEditPost", nil)
	cdc.RegisterConcrete(&MsgAddPostAttachment{}, "desmos/MsgAddPostAttachment", nil)
	cdc.RegisterConcrete(&MsgRemovePostAttachment{}, "desmos/MsgRemovePostAttachment", nil)
	cdc.RegisterConcrete(&MsgDeletePost{}, "desmos/MsgDeletePost", nil)
	cdc.RegisterConcrete(&MsgAnswerPoll{}, "desmos/MsgAnswerPoll", nil)

	cdc.RegisterConcrete(&Params{}, "desmos/x/posts/Params", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "desmos/x/posts/MsgUpdateParams", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"desmos.posts.v3.AttachmentContent",
		(*AttachmentContent)(nil),
		&Poll{},
		&Media{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePost{},
		&MsgEditPost{},
		&MsgAddPostAttachment{},
		&MsgRemovePostAttachment{},
		&MsgDeletePost{},
		&MsgAnswerPoll{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// AminoCodec references the global x/posts module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/posts and
	// defined at the application level.
	AminoCodec = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)

	// Register all Amino interfaces and concrete types on the gov Amino codec so that this can later be
	// used to properly serialize MsgSubmitProposal instances
	RegisterLegacyAminoCodec(govcodec.Amino)
}
