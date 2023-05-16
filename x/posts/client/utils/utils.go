package utils

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

// ParseCreatePostJSON reads and parses a CreatePostJSON from file.
func ParseCreatePostJSON(cdc codec.Codec, dataFile string) (CreatePostJSON, error) {
	var data CreatePostJSON

	contents, err := os.ReadFile(dataFile)
	if err != nil {
		return data, err
	}

	err = cdc.UnmarshalJSON(contents, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// ParseEditPostJSON reads and parses a EditPostJSON from file.
func ParseEditPostJSON(cdc codec.Codec, dataFile string) (EditPostJSON, error) {
	var data EditPostJSON

	contents, err := os.ReadFile(dataFile)
	if err != nil {
		return data, err
	}

	err = cdc.UnmarshalJSON(contents, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// ParseAttachmentContent reads and parses a AttachmentContent from file.
func ParseAttachmentContent(cdc codec.Codec, dataFile string) (types.AttachmentContent, error) {
	var data types.AttachmentContent

	contents, err := os.ReadFile(dataFile)
	if err != nil {
		return data, err
	}

	err = cdc.UnmarshalInterfaceJSON(contents, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
