package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgSaveProfile{}, "desmos/MsgSaveProfile", nil)
	cdc.RegisterConcrete(MsgDeleteProfile{}, "desmos/MsgDeleteProfile", nil)
	cdc.RegisterConcrete(MsgRequestDTagTransfer{}, "desmos/MsgRequestDTagTransfer", nil)
	cdc.RegisterConcrete(MsgCancelDTagTransfer{}, "desmos/MsgCancelDTagTransfer", nil)
	cdc.RegisterConcrete(MsgAcceptDTagTransfer{}, "desmos/MsgAcceptDTagTransfer", nil)
	cdc.RegisterConcrete(MsgRefuseDTagTransfer{}, "desmos/MsgRefuseDTagTransfer", nil)
	cdc.RegisterConcrete(MsgCreateRelationship{}, "desmos/MsgCreateRelationship", nil)
	cdc.RegisterConcrete(MsgDeleteRelationship{}, "desmos/MsgDeleteRelationship", nil)
	cdc.RegisterConcrete(MsgBlockUser{}, "desmos/MsgBlockUser", nil)
	cdc.RegisterConcrete(MsgUnblockUser{}, "desmos/MsgUnblockUser", nil)
	cdc.RegisterConcrete(MsgLinkChainAccount{}, "desmos/MsgLinkChainAccount", nil)
	cdc.RegisterConcrete(MsgUnlinkChainAccount{}, "desmos/MsgUnlinkChainAccount", nil)

	cdc.RegisterConcrete(&Profile{}, "desmos/Profile", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*authtypes.AccountI)(nil), &Profile{})
	registry.RegisterImplementations((*authtypes.GenesisAccount)(nil), &Profile{})
	registry.RegisterImplementations((*AddressData)(nil), &Bech32Address{}, &Base58Address{})

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSaveProfile{},
		&MsgDeleteProfile{},
		&MsgRequestDTagTransfer{},
		&MsgCancelDTagTransfer{},
		&MsgAcceptDTagTransfer{},
		&MsgRefuseDTagTransfer{},
		&MsgCreateRelationship{},
		&MsgDeleteRelationship{},
		&MsgBlockUser{},
		&MsgUnblockUser{},
		&MsgLinkChainAccount{},
		&MsgUnlinkChainAccount{},
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

	ProtoCdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
