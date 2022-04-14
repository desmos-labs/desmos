package wasm_test

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v3/x/relationships/types"
)

func buildCreateRelationshipRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.RelationshipsMsg{CreateRelationship: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildDeleteRelationshipRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.RelationshipsMsg{DeleteRelationship: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildBlockUserRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.RelationshipsMsg{BlockUser: cdc.MustMarshalJSON(msg)})
	return bz
}

func buildUnblockUserRequest(cdc codec.Codec, msg sdk.Msg) json.RawMessage {
	bz, _ := json.Marshal(types.RelationshipsMsg{UnblockUser: cdc.MustMarshalJSON(msg)})
	return bz
}
