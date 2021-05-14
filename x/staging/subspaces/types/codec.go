package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgCreateSubspace{}, "desmos/MsgCreateSubspace", nil)
	cdc.RegisterConcrete(MsgAddAdmin{}, "desmos/MsgAddAdmin", nil)
	cdc.RegisterConcrete(MsgRemoveAdmin{}, "desmos/MsgRemoveAdmin", nil)
	cdc.RegisterConcrete(MsgEnableUserPosts{}, "desmos/MsgAllowUserPosts", nil)
	cdc.RegisterConcrete(MsgDisableUserPosts{}, "desmos/MsgBlockUserPosts", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSubspace{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/subspaces module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/subspaces and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)
