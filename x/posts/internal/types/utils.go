package types

import emoji "github.com/tmdvs/Go-Emoji-Utils"

// convertToUnicode returns the Unicode value of the given string or the given string if it is not an actual emoji
func convertToUnicode(emo string) string {
	result, err := emoji.LookupEmoji(emo)
	if err != nil {
		return emo
	}

	return "U+" + result.Key
}
