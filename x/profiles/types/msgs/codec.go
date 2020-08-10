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
	cdc.RegisterConcrete(MsgCreateMonoDirectionalRelationship{}, "desmos/MsgCreateMonoDirectionalRelationship", nil)
	cdc.RegisterConcrete(MsgRequestBidirectionalRelationship{}, "desmos/MsgRequestBidirectionalRelationship", nil)
	cdc.RegisterConcrete(MsgAcceptBidirectionalRelationship{}, "desmos/MsgAcceptBidirectionalRelationship", nil)
	cdc.RegisterConcrete(MsgDenyBidirectionalRelationship{}, "desmos/MsgDenyBidirectionalRelationship", nil)
	cdc.RegisterConcrete(MsgDeleteRelationships{}, "desmos/MsgDeleteRelationships", nil)
}
