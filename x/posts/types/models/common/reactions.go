package common

import (
	emoji "github.com/desmos-labs/Go-Emoji-Utils"
)

// GetEmojiByShortCodeOrValue returns the emoji that has either one of their shortcode equals to
// the given string value, or its UNIX value equals to it.
// If such emoji is found, returns it along side with `true`. Otherwise, `false` is returned instead.
func GetEmojiByShortCodeOrValue(shortCodeOrValue string) (*emoji.Emoji, bool) {
	if emojiValue, err := emoji.LookupEmojiByCode(shortCodeOrValue); err == nil {
		return &emojiValue, true
	} else if emojiValue, err := emoji.LookupEmoji(shortCodeOrValue); err == nil {
		return &emojiValue, true
	}
	return nil, false
}
