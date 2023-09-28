package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgCreateDenom{}, "desmos/MsgCreateDenom")
	legacy.RegisterAminoMsg(cdc, &MsgMint{}, "desmos/MsgMint")
	legacy.RegisterAminoMsg(cdc, &MsgBurn{}, "desmos/MsgBurn")
	legacy.RegisterAminoMsg(cdc, &MsgSetDenomMetadata{}, "desmos/MsgSetDenomMetadata")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "desmos/x/tokenfactory/MsgUpdateParams")

	cdc.RegisterConcrete(&Params{}, "desmos/x/tokenfactory/Params", nil)

}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgMint{},
		&MsgBurn{},
		&MsgSetDenomMetadata{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
