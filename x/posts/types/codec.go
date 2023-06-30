package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
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

	legacy.RegisterAminoMsg(cdc, &MsgCreatePost{}, "desmos/MsgCreatePost")
	legacy.RegisterAminoMsg(cdc, &MsgEditPost{}, "desmos/MsgEditPost")
	legacy.RegisterAminoMsg(cdc, &MsgAddPostAttachment{}, "desmos/MsgAddPostAttachment")
	legacy.RegisterAminoMsg(cdc, &MsgRemovePostAttachment{}, "desmos/MsgRemovePostAttachment")
	legacy.RegisterAminoMsg(cdc, &MsgDeletePost{}, "desmos/MsgDeletePost")
	legacy.RegisterAminoMsg(cdc, &MsgAnswerPoll{}, "desmos/MsgAnswerPoll")
	legacy.RegisterAminoMsg(cdc, &MsgMovePost{}, "desmos/MsgMovePost")

	legacy.RegisterAminoMsg(cdc, &MsgRequestPostOwnerTransfer{}, "desmos/MsgRequestPostOwnerTransfer")
	legacy.RegisterAminoMsg(cdc, &MsgCancelPostOwnerTransferRequest{}, "desmos/MsgCancelPostOwnerTransferRequest")
	legacy.RegisterAminoMsg(cdc, &MsgAcceptPostOwnerTransferRequest{}, "desmos/MsgAcceptPostOwnerTransferRequest")
	legacy.RegisterAminoMsg(cdc, &MsgRefusePostOwnerTransferRequest{}, "desmos/MsgRefusePostOwnerTransferRequest")

	cdc.RegisterConcrete(&Params{}, "desmos/x/posts/Params", nil)
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "desmos/x/posts/MsgUpdateParams")
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
		&MsgMovePost{},
		&MsgRequestPostOwnerTransfer{},
		&MsgCancelPostOwnerTransferRequest{},
		&MsgAcceptPostOwnerTransferRequest{},
		&MsgRefusePostOwnerTransferRequest{},
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
