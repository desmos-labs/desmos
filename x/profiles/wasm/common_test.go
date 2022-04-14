package wasm_test

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func buildSaveProfileRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.ProfilesMsg{SaveProfile: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildDeleteProfileRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.ProfilesMsg{DeleteProfile: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildRequestDTagTransferRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.ProfilesMsg{RequestDtagTransfer: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildAcceptDTagTransferRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.ProfilesMsg{AcceptDtagTransferRequest: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildRefuseDTagTransferRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.ProfilesMsg{RefuseDtagTransferRequest: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildCancelDTagTransferRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.ProfilesMsg{CancelDtagTransferRequest: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildLinkChainAccountRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.ProfilesMsg{LinkChainAccount: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildLinkApplicationRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.ProfilesMsg{LinkApplication: cdc.MustMarshalJSON(msg)})
	return bz
}
