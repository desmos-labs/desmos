package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
)

// NewGenesisState returns a new GenesisState instance
func NewGenesisState(
	subspacesData []SubspaceDataEntry,
	registeredReactions []RegisteredReaction,
	postsData []PostDataEntry,
	reactions []Reaction,
	subspacesParams []SubspaceReactionsParams,
) *GenesisState {
	return &GenesisState{
		SubspacesData:       subspacesData,
		RegisteredReactions: registeredReactions,
		PostsData:           postsData,
		Reactions:           reactions,
		SubspacesParams:     subspacesParams,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, nil, nil, nil)
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (e GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, report := range e.Reactions {
		err := report.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, entry := range data.SubspacesData {
		if containsDuplicatedSubspaceDataEntry(data.SubspacesData, entry) {
			return fmt.Errorf("duplicated subspace data entry: %d", entry.SubspaceID)
		}

		err := entry.Validate()
		if err != nil {
			return fmt.Errorf("invalid subspace data entry (subspace %d): %s", entry.SubspaceID, err)
		}
	}

	for _, reaction := range data.RegisteredReactions {
		if containsDuplicatedRegisteredReactions(data.RegisteredReactions, reaction) {
			return fmt.Errorf("duplicated registered reaction: subspace id %d, reaction id %d",
				reaction.SubspaceID, reaction.ID)
		}

		err := reaction.Validate()
		if err != nil {
			return fmt.Errorf("invalid registered reaction: %s", err)
		}
	}

	for _, entry := range data.PostsData {
		if containsDuplicatedPostDataEntry(data.PostsData, entry) {
			return fmt.Errorf("duplicated post data entry: subspace id %d, post id %d", entry.SubspaceID, entry.PostID)
		}

		err := entry.Validate()
		if err != nil {
			return fmt.Errorf("invalid post data entry (subspace id %d, post id %d): %s",
				entry.SubspaceID, entry.PostID, err)
		}
	}

	for _, reaction := range data.Reactions {
		if containsDuplicatedReaction(data.Reactions, reaction) {
			return fmt.Errorf("duplicated reaction: subspace id %d, reaction id %d",
				reaction.SubspaceID, reaction.ID)
		}

		err := reaction.Validate()
		if err != nil {
			return fmt.Errorf("invalid reaction: %s", err)
		}
	}

	for _, params := range data.SubspacesParams {
		if containsDuplicatedSubspaceParams(data.SubspacesParams, params) {
			return fmt.Errorf("duplicated params: subspace id %d", params.SubspaceID)
		}

		err := params.Validate()
		if err != nil {
			return fmt.Errorf("invalid params: %s", err)
		}
	}

	return nil
}

// containsDuplicatedSubspaceDataEntry tells whether the given entries slice contains
// two or more entries for the same subspace
func containsDuplicatedSubspaceDataEntry(entries []SubspaceDataEntry, entry SubspaceDataEntry) bool {
	var count = 0
	for _, s := range entries {
		if s.SubspaceID == entry.SubspaceID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedRegisteredReactions tells whether the given registered reactions slice contains
// two or more reactions having the same id of the given one
func containsDuplicatedRegisteredReactions(reactions []RegisteredReaction, reaction RegisteredReaction) bool {
	var count = 0
	for _, s := range reactions {
		if s.SubspaceID == reaction.SubspaceID && s.ID == reaction.ID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedPostDataEntry tells whether the given entries slice contains
// two or more entries for the same subspace
func containsDuplicatedPostDataEntry(entries []PostDataEntry, entry PostDataEntry) bool {
	var count = 0
	for _, s := range entries {
		if s.SubspaceID == entry.SubspaceID && s.PostID == entry.PostID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedReaction tells whether the given reactions slice contains
// two or more reactions having the same id of the given one
func containsDuplicatedReaction(reactions []Reaction, reaction Reaction) bool {
	var count = 0
	for _, s := range reactions {
		if s.SubspaceID == reaction.SubspaceID && s.PostID == reaction.PostID && s.ID == reaction.ID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedSubspaceParams tells whether the given params slice contains
// two or more params for the same subspace of the given one
func containsDuplicatedSubspaceParams(paramsSlice []SubspaceReactionsParams, params SubspaceReactionsParams) bool {
	var count = 0
	for _, s := range paramsSlice {
		if s.SubspaceID == params.SubspaceID {
			count++
		}
	}
	return count > 1
}

// --------------------------------------------------------------------------------------------------------------------

// NewSubspaceDataEntry returns a new SubspaceDataEntry instance
func NewSubspaceDataEntry(subspaceID uint64, registeredReactionID uint32) SubspaceDataEntry {
	return SubspaceDataEntry{
		SubspaceID:           subspaceID,
		RegisteredReactionID: registeredReactionID,
	}
}

// Validate returns an error if something is wrong within the entry data
func (e SubspaceDataEntry) Validate() error {
	if e.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", e.SubspaceID)
	}

	if e.RegisteredReactionID == 0 {
		return fmt.Errorf("invalid initial registered reaction id: %d", e.RegisteredReactionID)
	}

	return nil
}

// NewPostDataEntry returns a new PostDataEntry instance
func NewPostDataEntry(subspaceID uint64, postID uint64, reactionID uint32) PostDataEntry {
	return PostDataEntry{
		SubspaceID: subspaceID,
		PostID:     postID,
		ReactionID: reactionID,
	}
}

// Validate returns an error if something is wrong within the entry data
func (e PostDataEntry) Validate() error {
	if e.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", e.SubspaceID)
	}

	if e.PostID == 0 {
		return fmt.Errorf("invalid post id: %d", e.PostID)
	}

	if e.ReactionID == 0 {
		return fmt.Errorf("invalid initial reaction id: %d", e.ReactionID)
	}

	return nil
}
