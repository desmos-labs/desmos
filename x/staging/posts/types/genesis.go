package types

import "fmt"

func NewPostReactionsEntry(postID string, reactions []PostReaction) PostReactionsEntry {
	return PostReactionsEntry{
		PostID:    postID,
		Reactions: reactions,
	}
}

// ___________________________________________________________________________________________________________________

// NewGenesisState creates a new genesis state
func NewGenesisState(
	posts []Post, userAnswers []UserAnswer,
	postReactions []PostReactionsEntry, registeredReactions []RegisteredReaction, reports []Report, params Params,
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
			return fmt.Errorf("invalid poll answers; post with id %s does not exist", answer.PostID)
		}
		err := answer.Validate()
		if err != nil {
			return err
		}
	}

	for _, postReaction := range data.PostsReactions {
		if _, ok := postMap[postReaction.PostID]; !ok {
			return fmt.Errorf("invalid reactions; post with id %s does not exist", postReaction.PostID)
		}

		for _, record := range postReaction.Reactions {
			err := record.Validate()
			if err != nil {
				return err
			}
		}
	}

	return data.Params.Validate()
}

// containsPostWithID tells whether or not the given posts contain one having the provided id
func containsPostWithID(posts []Post, id string) bool {
	for _, p := range posts {
		if p.PostID == id {
			return true
		}
	}
	return false
}
