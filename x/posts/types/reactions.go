package types

import (
	"fmt"
	"strings"

	"github.com/desmos-labs/desmos/x/commons"

	emoji "github.com/desmos-labs/Go-Emoji-Utils"
)

// NewRegisteredReaction returns a new RegisteredReaction
func NewRegisteredReaction(creator string, shortCode, value, subspace string) RegisteredReaction {
	return RegisteredReaction{
		ShortCode: shortCode,
		Value:     value,
		Subspace:  subspace,
		Creator:   creator,
	}
}

// Validate implements validator
func (reaction RegisteredReaction) Validate() error {
	if reaction.Creator == "" {
		return fmt.Errorf("invalid reaction creator: %s", reaction.Creator)
	}

	if !IsValidReactionCode(reaction.ShortCode) {
		return fmt.Errorf("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'")
	}

	if !commons.IsURIValid(reaction.Value) {
		return fmt.Errorf("reaction value should be a URL")
	}

	if !commons.IsValidSubspace(reaction.Subspace) {
		return fmt.Errorf("reaction subspace must be a valid sha-256 hash")
	}

	if _, found := GetEmojiByShortCodeOrValue(reaction.ShortCode); found {
		return fmt.Errorf("reaction has emoji shortcode: %s", reaction.ShortCode)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// NewPostReaction returns a new PostReaction
func NewPostReaction(shortcode, value string, owner string) PostReaction {
	return PostReaction{
		ShortCode: shortcode,
		Value:     value,
		Owner:     owner,
	}
}

// Validate implements validator
func (reaction PostReaction) Validate() error {
	if reaction.Owner == "" {
		return fmt.Errorf("invalid reaction owner: %s", reaction.Owner)
	}

	if len(strings.TrimSpace(reaction.Value)) == 0 {
		return fmt.Errorf("reaction value cannot be empty or blank")
	}

	if !IsValidReactionCode(reaction.ShortCode) {
		return fmt.Errorf("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'")
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// NewPostReactions allows to create a new PostReactions object from the given reactions
func NewPostReactions(reactions ...PostReaction) PostReactions {
	return PostReactions{Reactions: reactions}
}

// ContainsReactionFrom returns true if the reactions slice contain
// a reaction from the given user having the given value, false otherwise.
// NOTE: The value can be either an emoji or a shortcode.
func (reactions PostReactions) ContainsReactionFrom(user string, value string) bool {
	return reactions.IndexOfByUserAndValue(user, value) != -1
}

// IndexOfByUserAndValue returns the index of the reaction from the
// given user with the specified code inside the reactions slice.
// NOTE: The value can be either an emoji or a shortcode.
func (reactions PostReactions) IndexOfByUserAndValue(owner string, value string) int {
	var reactEmoji *emoji.Emoji
	if ej, found := GetEmojiByShortCodeOrValue(value); found {
		reactEmoji = ej
	}

	for index, reaction := range reactions.Reactions {
		if reaction.Owner == owner {
			if reactEmoji != nil {
				// Check the emoji value
				if reaction.Value == reactEmoji.Value {
					return index
				}

				// Check the emoji shortcodes
				for _, code := range reactEmoji.Shortcodes {
					if reaction.ShortCode == code {
						return index
					}
				}
			}

			if reactEmoji == nil {
				if value == reaction.ShortCode {
					return index
				}
			}
		}
	}
	return -1
}

// DeletePostReaction returns a new PostReactions slice not containing the
// reaction of the given user with the given value.
// If the reaction was removed properly, true is also returned. Otherwise,
// if no reaction was found, false is returned instead.
func (reactions PostReactions) RemoveReaction(user string, value string) (PostReactions, bool) {
	index := reactions.IndexOfByUserAndValue(user, value)
	if index == -1 {
		return reactions, false
	}

	return PostReactions{
		Reactions: append(reactions.Reactions[:index], reactions.Reactions[index+1:]...),
	}, true
}
