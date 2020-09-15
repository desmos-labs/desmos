package msgs

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// MsgsCodec is the codec
var MsgsCodec = codec.New()

func init() {
	RegisterMessagesCodec(MsgsCodec)
}

// RegisterMessagesCodec registers concrete types on the Amino codec
func RegisterMessagesCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSaveProfile{}, "desmos/MsgSaveProfile", nil)
	cdc.RegisterConcrete(MsgDeleteProfile{}, "desmos/MsgDeleteProfile", nil)
	cdc.RegisterConcrete(MsgRequestDTagTransfer{}, "desmos/MsgRequestDTagTransfer", nil)
	cdc.RegisterConcrete(MsgAcceptDTagTransfer{}, "desmos/MsgAcceptDTagTransfer", nil)
}
