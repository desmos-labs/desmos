package types

import (
	"encoding/json"
)

type ReactionsMsg struct {
	AddReaction              *json.RawMessage `json:"add_reaction"`
	RemoveReaction           *json.RawMessage `json:"remove_reaction"`
	AddRegisteredReaction    *json.RawMessage `json:"add_registered_reaction"`
	EditRegisteredReaction   *json.RawMessage `json:"edit_registered_reaction"`
	RemoveRegisteredReaction *json.RawMessage `json:"remove_registered_reaction"`
	SetReactionsParams       *json.RawMessage `json:"set_reactions_params"`
}

type ReactionsQuery struct {
	Reactions           *json.RawMessage `json:"reactions"`
	Reaction            *json.RawMessage `json:"reaction"`
	RegisteredReactions *json.RawMessage `json:"registered_reactions"`
	RegisteredReaction  *json.RawMessage `json:"registered_reaction"`
	ReactionsParams     *json.RawMessage `json:"reactions_params"`
}
