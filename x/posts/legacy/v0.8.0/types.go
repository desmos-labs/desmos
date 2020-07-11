package v080

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v060 "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
	"regexp"
)

var URIRegEx = regexp.MustCompile(
	`^(?:http(s)?://)[\w.-]+(?:\.[\w.-]+)+[\w\-._~:/?#[\]@!$&'()*+,;=.]+$`)

// GenesisState contains the data of a v0.6.0 genesis state for the posts module
type GenesisState struct {
	Posts               []v040posts.Post                  `json:"posts"`
	UsersPollAnswers    map[string][]v040posts.UserAnswer `json:"users_poll_answers"`
	PostReactions       map[string][]v060.PostReaction    `json:"post_reactions"`
	RegisteredReactions []v040posts.Reaction              `json:"registered_reactions"`
	Params              Params                            `json:"params"`
}

type Params struct {
	MaxPostMessageLength            sdk.Int `json:"max_post_message_length" yaml:"max_post_message_length"`
	MaxOptionalDataFieldsNumber     sdk.Int `json:"max_optional_data_fields_number" yaml:"max_optional_data_fields_number"`
	MaxOptionalDataFieldValueLength sdk.Int `json:"max_optional_data_field_value_length" yaml:"max_optional_data_field_value_length"`
}
