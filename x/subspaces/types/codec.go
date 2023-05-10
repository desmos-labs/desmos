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
	cdc.RegisterConcrete(&MsgCreateSubspace{}, "desmos/MsgCreateSubspace", nil)
	cdc.RegisterConcrete(&MsgEditSubspace{}, "desmos/MsgEditSubspace", nil)
	cdc.RegisterConcrete(&MsgDeleteSubspace{}, "desmos/MsgDeleteSubspace", nil)

	cdc.RegisterConcrete(&MsgCreateSection{}, "desmos/MsgCreateSection", nil)
	cdc.RegisterConcrete(&MsgEditSection{}, "desmos/MsgEditSection", nil)
	cdc.RegisterConcrete(&MsgMoveSection{}, "desmos/MsgMoveSection", nil)
	cdc.RegisterConcrete(&MsgDeleteSection{}, "desmos/MsgDeleteSection", nil)

	cdc.RegisterConcrete(&MsgCreateUserGroup{}, "desmos/MsgCreateUserGroup", nil)
	cdc.RegisterConcrete(&MsgEditUserGroup{}, "desmos/MsgEditUserGroup", nil)
	cdc.RegisterConcrete(&MsgMoveUserGroup{}, "desmos/MsgMoveUserGroup", nil)
	cdc.RegisterConcrete(&MsgSetUserGroupPermissions{}, "desmos/MsgSetUserGroupPermissions", nil)
	cdc.RegisterConcrete(&MsgDeleteUserGroup{}, "desmos/MsgDeleteUserGroup", nil)

	cdc.RegisterConcrete(&MsgAddUserToUserGroup{}, "desmos/MsgAddUserToUserGroup", nil)
	cdc.RegisterConcrete(&MsgRemoveUserFromUserGroup{}, "desmos/MsgRemoveUserFromUserGroup", nil)

	cdc.RegisterConcrete(&MsgSetUserPermissions{}, "desmos/MsgSetUserPermissions", nil)

	cdc.RegisterConcrete(&MsgGrantTreasuryAuthorization{}, "desmos/MsgGrantTreasuryAuthorization", nil)
	cdc.RegisterConcrete(&MsgRevokeTreasuryAuthorization{}, "desmos/MsgRevokeTreasuryAuthorization", nil)

	cdc.RegisterConcrete(&MsgGrantAllowance{}, "desmos/MsgGrantAllowance", nil)
	cdc.RegisterConcrete(&MsgRevokeAllowance{}, "desmos/MsgRevokeAllowance", nil)

	cdc.RegisterInterface((*Grantee)(nil), nil)
	cdc.RegisterConcrete(&UserGrantee{}, "desmos/UserGrantee", nil)
	cdc.RegisterConcrete(&GroupGrantee{}, "desmos/GroupGrantee", nil)

	cdc.RegisterConcrete(&MsgUpdateSubspaceFeeTokens{}, "desmos/MsgUpdateSubspaceFeeTokens", nil)
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

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
