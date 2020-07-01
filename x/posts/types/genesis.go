package types

// GenesisState contains the data of the genesis state for the posts module
type GenesisState struct {
	Posts               Posts                    `json:"posts"`
	UsersPollAnswers    map[string]UserAnswers   `json:"users_poll_answers"`
	PostReactions       map[string]PostReactions `json:"post_reactions"`
	RegisteredReactions Reactions                `json:"registered_reactions"`
	Params              Params                   `json:"params"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(posts Posts, postReactions map[string]PostReactions, registeredR Reactions, params Params) GenesisState {
	return GenesisState{
		Posts:               posts,
		PostReactions:       postReactions,
		RegisteredReactions: registeredR,
		Params:              params,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params: DefaultParams(),
	}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, reaction := range data.RegisteredReactions {
		if err := reaction.Validate(); err != nil {
			return err
		}
	}

	for _, record := range data.Posts {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	for _, pollAnswers := range data.UsersPollAnswers {
		for _, pollAnswer := range pollAnswers {
			if err := pollAnswer.Validate(); err != nil {
				return err
			}
		}
	}

	for _, postReaction := range data.PostReactions {
		for _, record := range postReaction {
			if err := record.Validate(); err != nil {
				return err
			}
		}
	}

	if err := data.Params.Validate(); err != nil {
		return err
	}

	return nil
}
