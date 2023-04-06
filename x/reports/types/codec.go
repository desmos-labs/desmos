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
	cdc.RegisterInterface((*ReportTarget)(nil), nil)
	cdc.RegisterConcrete(&UserTarget{}, "desmos/UserTarget", nil)
	cdc.RegisterConcrete(&PostTarget{}, "desmos/PostTarget", nil)

	cdc.RegisterConcrete(&MsgCreateReport{}, "desmos/MsgCreateReport", nil)
	cdc.RegisterConcrete(&MsgDeleteReport{}, "desmos/MsgDeleteReport", nil)
	cdc.RegisterConcrete(&MsgSupportStandardReason{}, "desmos/MsgSupportStandardReason", nil)
	cdc.RegisterConcrete(&MsgAddReason{}, "desmos/MsgAddReason", nil)
	cdc.RegisterConcrete(&MsgRemoveReason{}, "desmos/MsgRemoveReason", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "desmos/x/reports/MsgUpdateParams", nil)

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

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
