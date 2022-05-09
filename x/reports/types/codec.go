package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*ReportData)(nil), nil)
	cdc.RegisterConcrete(&UserData{}, "desmos/UserData", nil)
	cdc.RegisterConcrete(&PostData{}, "desmos/PostData", nil)

	cdc.RegisterConcrete(MsgCreateReport{}, "desmos/MsgCreateReport", nil)
	cdc.RegisterConcrete(MsgDeleteReport{}, "desmos/MsgDeleteReport", nil)
	cdc.RegisterConcrete(MsgSupportReasons{}, "desmos/MsgSupportReasons", nil)
	cdc.RegisterConcrete(MsgAddReason{}, "desmos/MsgAddReason", nil)
	cdc.RegisterConcrete(MsgRemoveReason{}, "desmos/MsgRemoveReason", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"desmos.reports.v1.ReportData",
		(*ReportData)(nil),
		&UserData{},
		&PostData{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateReport{},
		&MsgDeleteReport{},
		&MsgSupportReasons{},
		&MsgAddReason{},
		&MsgRemoveReason{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// AminoCdc references the global x/reports module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/reports and
	// defined at the application level.
	AminoCdc = codec.NewAminoCodec(amino)

	ModuleCdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
