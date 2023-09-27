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
	legacy.RegisterAminoMsg(cdc, &MsgCreateSubspace{}, "desmos/MsgCreateSubspace")
	legacy.RegisterAminoMsg(cdc, &MsgEditSubspace{}, "desmos/MsgEditSubspace")
	legacy.RegisterAminoMsg(cdc, &MsgDeleteSubspace{}, "desmos/MsgDeleteSubspace")

	legacy.RegisterAminoMsg(cdc, &MsgCreateSection{}, "desmos/MsgCreateSection")
	legacy.RegisterAminoMsg(cdc, &MsgEditSection{}, "desmos/MsgEditSection")
	legacy.RegisterAminoMsg(cdc, &MsgMoveSection{}, "desmos/MsgMoveSection")
	legacy.RegisterAminoMsg(cdc, &MsgDeleteSection{}, "desmos/MsgDeleteSection")

	legacy.RegisterAminoMsg(cdc, &MsgCreateUserGroup{}, "desmos/MsgCreateUserGroup")
	legacy.RegisterAminoMsg(cdc, &MsgEditUserGroup{}, "desmos/MsgEditUserGroup")
	legacy.RegisterAminoMsg(cdc, &MsgMoveUserGroup{}, "desmos/MsgMoveUserGroup")
	legacy.RegisterAminoMsg(cdc, &MsgSetUserGroupPermissions{}, "desmos/MsgSetUserGroupPermissions")
	legacy.RegisterAminoMsg(cdc, &MsgDeleteUserGroup{}, "desmos/MsgDeleteUserGroup")

	legacy.RegisterAminoMsg(cdc, &MsgAddUserToUserGroup{}, "desmos/MsgAddUserToUserGroup")
	legacy.RegisterAminoMsg(cdc, &MsgRemoveUserFromUserGroup{}, "desmos/MsgRemoveUserFromUserGroup")

	legacy.RegisterAminoMsg(cdc, &MsgSetUserPermissions{}, "desmos/MsgSetUserPermissions")

	legacy.RegisterAminoMsg(cdc, &MsgGrantTreasuryAuthorization{}, "desmos/MsgGrantTreasuryAuthorization")
	legacy.RegisterAminoMsg(cdc, &MsgRevokeTreasuryAuthorization{}, "desmos/MsgRevokeTreasuryAuthorization")

	legacy.RegisterAminoMsg(cdc, &MsgGrantAllowance{}, "desmos/MsgGrantAllowance")
	legacy.RegisterAminoMsg(cdc, &MsgRevokeAllowance{}, "desmos/MsgRevokeAllowance")

	cdc.RegisterInterface((*Grantee)(nil), nil)
	cdc.RegisterConcrete(&UserGrantee{}, "desmos/UserGrantee", nil)
	cdc.RegisterConcrete(&GroupGrantee{}, "desmos/GroupGrantee", nil)

	legacy.RegisterAminoMsg(cdc, &MsgUpdateSubspaceFeeTokens{}, "desmos/MsgUpdateSubspaceFeeTokens")
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"desmos.subspaces.v3.Grantee",
		(*Grantee)(nil),
		&UserGrantee{},
		&GroupGrantee{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSubspace{},
		&MsgEditSubspace{},
		&MsgCreateUserGroup{},
		&MsgCreateSection{},
		&MsgEditSection{},
		&MsgMoveSection{},
		&MsgDeleteSection{},
		&MsgEditUserGroup{},
		&MsgMoveUserGroup{},
		&MsgSetUserGroupPermissions{},
		&MsgDeleteUserGroup{},
		&MsgAddUserToUserGroup{},
		&MsgRemoveUserFromUserGroup{},
		&MsgSetUserPermissions{},
		&MsgGrantTreasuryAuthorization{},
		&MsgRevokeTreasuryAuthorization{},
		&MsgGrantAllowance{},
		&MsgRevokeAllowance{},
		&MsgUpdateSubspaceFeeTokens{},
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

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)
}
