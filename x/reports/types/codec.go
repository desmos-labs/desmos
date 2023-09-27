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
	cdc.RegisterInterface((*ReportTarget)(nil), nil)
	cdc.RegisterConcrete(&UserTarget{}, "desmos/UserTarget", nil)
	cdc.RegisterConcrete(&PostTarget{}, "desmos/PostTarget", nil)

	legacy.RegisterAminoMsg(cdc, &MsgCreateReport{}, "desmos/MsgCreateReport")
	legacy.RegisterAminoMsg(cdc, &MsgDeleteReport{}, "desmos/MsgDeleteReport")
	legacy.RegisterAminoMsg(cdc, &MsgSupportStandardReason{}, "desmos/MsgSupportStandardReason")
	legacy.RegisterAminoMsg(cdc, &MsgAddReason{}, "desmos/MsgAddReason")
	legacy.RegisterAminoMsg(cdc, &MsgRemoveReason{}, "desmos/MsgRemoveReason")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "desmos/x/reports/MsgUpdateParams")

	cdc.RegisterConcrete(&Params{}, "desmos/x/reports/Params", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"desmos.reports.v1.ReportTarget",
		(*ReportTarget)(nil),
		&UserTarget{},
		&PostTarget{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateReport{},
		&MsgDeleteReport{},
		&MsgSupportStandardReason{},
		&MsgAddReason{},
		&MsgRemoveReason{},
		&MsgUpdateParams{},
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
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

}
