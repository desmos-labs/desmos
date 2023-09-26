package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"

	"github.com/desmos-labs/desmos/v6/types/crypto/ethsecp256k1"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// Register custom key types
	cdc.RegisterConcrete(&ethsecp256k1.PubKey{}, ethsecp256k1.PubKeyName, nil)
	cdc.RegisterConcrete(&ethsecp256k1.PrivKey{}, ethsecp256k1.PrivKeyName, nil)

	legacy.RegisterAminoMsg(cdc, &MsgSaveProfile{}, "desmos/MsgSaveProfile")
	legacy.RegisterAminoMsg(cdc, &MsgDeleteProfile{}, "desmos/MsgDeleteProfile")
	legacy.RegisterAminoMsg(cdc, &MsgRequestDTagTransfer{}, "desmos/MsgRequestDTagTransfer")
	legacy.RegisterAminoMsg(cdc, &MsgCancelDTagTransferRequest{}, "desmos/MsgCancelDTagTransferRequest")
	legacy.RegisterAminoMsg(cdc, &MsgAcceptDTagTransferRequest{}, "desmos/MsgAcceptDTagTransferRequest")
	legacy.RegisterAminoMsg(cdc, &MsgRefuseDTagTransferRequest{}, "desmos/MsgRefuseDTagTransferRequest")
	legacy.RegisterAminoMsg(cdc, &MsgLinkChainAccount{}, "desmos/MsgLinkChainAccount")
	legacy.RegisterAminoMsg(cdc, &MsgUnlinkChainAccount{}, "desmos/MsgUnlinkChainAccount")
	legacy.RegisterAminoMsg(cdc, &MsgSetDefaultExternalAddress{}, "desmos/MsgSetDefaultExternalAddress")
	legacy.RegisterAminoMsg(cdc, &MsgLinkApplication{}, "desmos/MsgLinkApplication")
	legacy.RegisterAminoMsg(cdc, &MsgUnlinkApplication{}, "desmos/MsgUnlinkApplication")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "desmos/x/profiles/MsgUpdateParams")

	cdc.RegisterConcrete(&Params{}, "desmos/x/profiles/Params", nil)

	cdc.RegisterInterface((*AddressData)(nil), nil)
	cdc.RegisterConcrete(&Bech32Address{}, "desmos/Bech32Address", nil)
	cdc.RegisterConcrete(&Base58Address{}, "desmos/Base58Address", nil)
	cdc.RegisterConcrete(&HexAddress{}, "desmos/HexAddress", nil)

	cdc.RegisterInterface((*Signature)(nil), nil)
	cdc.RegisterConcrete(&SingleSignature{}, "desmos/SingleSignature", nil)
	cdc.RegisterConcrete(&CosmosMultiSignature{}, "desmos/CosmosMultiSignature", nil)

	cdc.RegisterConcrete(&Profile{}, "desmos/Profile", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.AccountI)(nil), &Profile{})
	registry.RegisterImplementations((*exported.VestingAccount)(nil), &Profile{})
	registry.RegisterImplementations((*authtypes.GenesisAccount)(nil), &Profile{})
	registry.RegisterInterface(
		"desmos.profiles.v3.AddressData",
		(*AddressData)(nil),
		&Bech32Address{},
		&Base58Address{},
		&HexAddress{},
	)
	registry.RegisterInterface(
		"desmos.profiles.v3.Signature",
		(*Signature)(nil),
		&SingleSignature{},
		&CosmosMultiSignature{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSaveProfile{},
		&MsgDeleteProfile{},
		&MsgRequestDTagTransfer{},
		&MsgCancelDTagTransferRequest{},
		&MsgAcceptDTagTransferRequest{},
		&MsgRefuseDTagTransferRequest{},
		&MsgLinkChainAccount{},
		&MsgUnlinkChainAccount{},
		&MsgLinkApplication{},
		&MsgUnlinkApplication{},
		&MsgSetDefaultExternalAddress{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// AminoCdc references the global x/relationships module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/relationships and
	// defined at the application level.
	AminoCdc = codec.NewAminoCodec(amino)

	ModuleCdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)
}
