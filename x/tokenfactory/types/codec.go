package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcodec "github.com/cosmos/cosmos-sdk/x/gov/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgCreateDenom{}, "desmos/x/tokenfactory/MsgCreateDenom")
	legacy.RegisterAminoMsg(cdc, &MsgMint{}, "desmos/x/tokenfactory/MsgMint")
	legacy.RegisterAminoMsg(cdc, &MsgBurn{}, "desmos/x/tokenfactory/MsgBurn")
	legacy.RegisterAminoMsg(cdc, &MsgSetDenomMetadata{}, "desmos/x/tokenfactory/MsgSetDenomMetadata")
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

var (
	amino = codec.NewLegacyAmino()

	// AminoCdc references the global x/tokenfactory module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/tokenfactory and
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

	// Register all Amino interfaces and concrete types on the gov Amino codec so that this can later be
	// used to properly serialize MsgSubmitProposal instances
	RegisterLegacyAminoCodec(govcodec.Amino)
}
