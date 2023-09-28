package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
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
	legacy.RegisterAminoMsg(cdc, &MsgCancelPostOwnerTransferRequest{}, "desmos/MsgCancelPostOwnerTransfer")
	legacy.RegisterAminoMsg(cdc, &MsgAcceptPostOwnerTransferRequest{}, "desmos/MsgAcceptPostOwnerTransfer")
	legacy.RegisterAminoMsg(cdc, &MsgRefusePostOwnerTransferRequest{}, "desmos/MsgRefusePostOwnerTransfer")

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
