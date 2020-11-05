package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgSaveProfile{}, "desmos/MsgSaveProfile", nil)
	cdc.RegisterConcrete(MsgDeleteProfile{}, "desmos/MsgDeleteProfile", nil)
	cdc.RegisterConcrete(MsgRequestDTagTransfer{}, "desmos/MsgRequestDTagTransfer", nil)
	cdc.RegisterConcrete(MsgCancelDTagTransfer{}, "desmos/MsgCancelDTagTransfer", nil)
	cdc.RegisterConcrete(MsgAcceptDTagTransfer{}, "desmos/MsgAcceptDTagTransfer", nil)
	cdc.RegisterConcrete(MsgRefuseDTagTransfer{}, "desmos/MsgRefuseDTagTransfer", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSaveProfile{},
		&MsgDeleteProfile{},
		&MsgRequestDTagTransfer{},
		&MsgCancelDTagTransfer{},
		&MsgAcceptDTagTransfer{},
		&MsgRefuseDTagTransfer{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/relationships module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/relationships and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
