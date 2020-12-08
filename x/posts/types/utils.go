package types

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
	"unicode"

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

// GetTags matches tags and returns them as an array of strings
//
// The hashtag (#) itself is NOT included as part of the tag string
//
// The function should match the javascript regex: '/([^\S]|^)#([^\s#.,!)]+)(?![^\s.,!)])/g'.
// Since golang re2 engine does not have positive lookahead, the end of the tag is matched by splitting the input string.
// The 'tagsSplitter' function defines the end of a tag, and the 'matchTags' regex has a requirement that it must match the end of a string.
func GetTags(s string) []string {
	res := make([]string, 0)

	// tagsSplitter returns true if the current rune is a tag ending
	// Tags MUST end with whitespace, '.' ',' '!' or ')'
	fields := strings.FieldsFunc(s, func(c rune) bool {
		if unicode.IsSpace(c) {
			return true
		}
		switch c {
		case '.', ',', '!', ')':
			return true
		}
		return false
	})

	for _, v := range fields {
		sub := hashtagRegEx.FindStringSubmatch(v)
		if len(sub) > 1 {
			res = append(res, sub[1])
		}
	}
	return res
}

// IsNumeric returns whether the given string represents a numeric value or not
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsEmoji checks whether the value is an emoji or an emoji unicode
func IsEmoji(value string) bool {

	_, err := emoji.LookupEmoji(value)
	if err == nil {
		return true
	}

	trimmed := strings.TrimPrefix(value, "U+")
	emo := emoji.Emojis[trimmed]
	return len(emo.Key) != 0
}

// ComputeID computes a post ID based on the content of the given post.
func ComputeID(post Post) string {
	jsonPost, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(jsonPost)
	return hex.EncodeToString(hash[:])
}
