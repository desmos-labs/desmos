package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgCreateRelationship{}, "desmos/MsgCreateRelationship")
	legacy.RegisterAminoMsg(cdc, &MsgDeleteRelationship{}, "desmos/MsgDeleteRelationship")
	legacy.RegisterAminoMsg(cdc, &MsgBlockUser{}, "desmos/MsgBlockUser")
	legacy.RegisterAminoMsg(cdc, &MsgUnblockUser{}, "desmos/MsgUnblockUser")
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRelationship{},
		&MsgDeleteRelationship{},
		&MsgBlockUser{},
		&MsgUnblockUser{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
