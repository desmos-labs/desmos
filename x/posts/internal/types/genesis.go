package types

import "fmt"

// GenesisState contains the data of the genesis state for the posts module
type GenesisState struct {
	Posts       Posts                          `json:"posts"`
	PollAnswers map[string]UsersAnswersDetails `json:"poll_answers_details"`
	Reactions   map[string]Reactions           `json:"reactions"`
	Hashtags    map[string]PostIDs             `json:"hashtags"`
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Posts {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	for _, pollAnswers := range data.PollAnswers {
		for _, pollAnswer := range pollAnswers {
			if err := pollAnswer.Validate(); err != nil {
				return err
			}
		}
	}

	for _, reactions := range data.Reactions {
		for _, record := range reactions {
			if err := record.Validate(); err != nil {
				return err
			}
		}
	}

	for _, IDs := range data.Hashtags {
		for _, id := range IDs {
			if !id.Valid() {
				return fmt.Errorf("invalid post ID, %s", id)
			}
		}
	}

	return nil
}
