package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcodec "github.com/cosmos/cosmos-sdk/x/gov/codec"

	"github.com/desmos-labs/desmos/v5/types/crypto/ethsecp256k1"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// Register custom key types
	cdc.RegisterConcrete(&ethsecp256k1.PubKey{}, ethsecp256k1.PubKeyName, nil)
	cdc.RegisterConcrete(&ethsecp256k1.PrivKey{}, ethsecp256k1.PrivKeyName, nil)

	cdc.RegisterConcrete(&MsgSaveProfile{}, "desmos/MsgSaveProfile", nil)
	cdc.RegisterConcrete(&MsgDeleteProfile{}, "desmos/MsgDeleteProfile", nil)
	cdc.RegisterConcrete(&MsgRequestDTagTransfer{}, "desmos/MsgRequestDTagTransfer", nil)
	cdc.RegisterConcrete(&MsgCancelDTagTransferRequest{}, "desmos/MsgCancelDTagTransferRequest", nil)
	cdc.RegisterConcrete(&MsgAcceptDTagTransferRequest{}, "desmos/MsgAcceptDTagTransferRequest", nil)
	cdc.RegisterConcrete(&MsgRefuseDTagTransferRequest{}, "desmos/MsgRefuseDTagTransferRequest", nil)
	cdc.RegisterConcrete(&MsgLinkChainAccount{}, "desmos/MsgLinkChainAccount", nil)
	cdc.RegisterConcrete(&MsgUnlinkChainAccount{}, "desmos/MsgUnlinkChainAccount", nil)
	cdc.RegisterConcrete(&MsgSetDefaultExternalAddress{}, "desmos/MsgSetDefaultExternalAddress", nil)
	cdc.RegisterConcrete(&MsgLinkApplication{}, "desmos/MsgLinkApplication", nil)
	cdc.RegisterConcrete(&MsgUnlinkApplication{}, "desmos/MsgUnlinkApplication", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "desmos/x/profiles/MsgUpdateParams", nil)

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
	registry.RegisterImplementations((*authtypes.AccountI)(nil), &Profile{})
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

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)

	// Register all Amino interfaces and concrete types on the gov Amino codec so that this can later be
	// used to properly serialize MsgSubmitProposal instances
	RegisterLegacyAminoCodec(govcodec.Amino)
}
