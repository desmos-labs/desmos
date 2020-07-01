package msgs

import "github.com/cosmos/cosmos-sdk/codec"

// MsgsCodec is the codec
var MsgsCodec = codec.New()

func init() {
	RegisterMessagesCodec(MsgsCodec)
}

// RegisterMessagesCodec registers concrete types on the Amino codec
func RegisterMessagesCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgReportPost{}, "desmos/MsgReportPost", nil)
}
