package types

import "encoding/json"

type RelationshipsMsg struct {
	CreateRelationship *json.RawMessage `json:"create_relationship"`
	DeleteRelationship *json.RawMessage `json:"delete_relationship"`
	BlockUser          *json.RawMessage `json:"block_user"`
	UnblockUser        *json.RawMessage `json:"unblock_user"`
}

type RelationshipsQuery struct {
	Relationships *json.RawMessage `json:"relationships"`
	Blocks        *json.RawMessage `json:"blocks"`
}
