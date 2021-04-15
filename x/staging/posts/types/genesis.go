package types

import "fmt"

func NewPostReactionsEntry(postID string, reactions []PostReaction) PostReactionsEntry {
	return PostReactionsEntry{
		PostId:    postID,
		Reactions: reactions,
	}
}

func NewUserAnswersEntry(postID string, answers []UserAnswer) UserAnswersEntry {
	return UserAnswersEntry{
		PostId:      postID,
		UserAnswers: answers,
	}
}

// ___________________________________________________________________________________________________________________

// NewGenesisState creates a new genesis state
func NewGenesisState(
	posts []Post, userPollAnswers []UserAnswersEntry,
	postReactions []PostReactionsEntry, registeredReactions []RegisteredReaction, params Params,
) *GenesisState {
	return &GenesisState{
		Posts:               posts,
		UsersPollAnswers:    userPollAnswers,
		PostsReactions:      postReactions,
		RegisteredReactions: registeredReactions,
		Params:              params,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, nil, nil, DefaultParams())
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, reaction := range data.RegisteredReactions {
		err := reaction.Validate()
		if err != nil {
			return err
		}
	}

	for _, record := range data.Posts {
		err := record.Validate()
		if err != nil {
			return err
		}
	}

	for _, pollAnswers := range data.UsersPollAnswers {
		if !containsPostWithID(data.Posts, pollAnswers.PostId) {
			return fmt.Errorf("invalid poll answers; post with id %s does not exist", pollAnswers.PostId)
		}

		for _, pollAnswer := range pollAnswers.UserAnswers {
			err := pollAnswer.Validate()
			if err != nil {
				return err
			}
		}
	}

	for _, postReaction := range data.PostsReactions {
		if !containsPostWithID(data.Posts, postReaction.PostId) {
			return fmt.Errorf("invalid reactions; post with id %s does not exist", postReaction.PostId)
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
		if p.PostId == id {
			return true
		}
	}
	return false
}
