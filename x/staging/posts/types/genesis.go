package types

import "fmt"

// NewGenesisState creates a new genesis state
func NewGenesisState(
	posts []Post, userAnswers []UserAnswer,
	postReactions []PostReaction, registeredReactions []RegisteredReaction, reports []Report, params Params,
) *GenesisState {
	return &GenesisState{
		Posts:               posts,
		UsersPollAnswers:    userAnswers,
		PostsReactions:      postReactions,
		RegisteredReactions: registeredReactions,
		Reports:             reports,
		Params:              params,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, nil, nil, nil, DefaultParams())
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, reaction := range data.RegisteredReactions {
		err := reaction.Validate()
		if err != nil {
			return err
		}
	}

	// The map for checking if post id exists or not
	postMap := make(map[string]bool)
	for _, record := range data.Posts {
		err := record.Validate()
		if err != nil {
			return err
		}
		postMap[record.PostID] = true
	}

	for _, answer := range data.UsersPollAnswers {
		if _, ok := postMap[answer.PostID]; !ok {
			return fmt.Errorf("invalid user answers; post with id %s does not exist", answer.PostID)
		}
		err := answer.Validate()
		if err != nil {
			return err
		}
	}

	for _, reaction := range data.PostsReactions {
		if _, ok := postMap[reaction.PostID]; !ok {
			return fmt.Errorf("invalid reactions; post with id %s does not exist", reaction.PostID)
		}
		err := reaction.Validate()
		if err != nil {
			return err
		}
	}

	for _, report := range data.Reports {
		err := report.Validate()
		if err != nil {
			return err
		}
	}

	return data.Params.Validate()
}
