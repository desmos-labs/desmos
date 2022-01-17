package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgCreateSubspace{}, "desmos/MsgCreateSubspace", nil)
	cdc.RegisterConcrete(MsgEditSubspace{}, "desmos/MsgEditSubspace", nil)
	cdc.RegisterConcrete(MsgCreateUserGroup{}, "desmos/MsgCreateUserGroup", nil)
	cdc.RegisterConcrete(MsgDeleteUserGroup{}, "desmos/MsgDeleteUserGroup", nil)
	cdc.RegisterConcrete(MsgAddUserToUserGroup{}, "desmos/MsgAddUserToUserGroup", nil)
	cdc.RegisterConcrete(MsgRemoveUserFromUserGroup{}, "desmos/MsgRemoveUserFromUserGroup", nil)
	cdc.RegisterConcrete(MsgSetPermissions{}, "desmos/MsgSetPermissions", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSubspace{},
		&MsgEditSubspace{},
		&MsgCreateUserGroup{},
		&MsgDeleteUserGroup{},
		&MsgAddUserToUserGroup{},
		&MsgRemoveUserFromUserGroup{},
		&MsgSetPermissions{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// AminoCodec references the global x/subspaces module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/subspaces and
	// defined at the application level.
	AminoCodec = codec.NewAminoCodec(amino)
)
