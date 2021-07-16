package types_test

import (
	"errors"
	"testing"

	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/staging/posts/types"

	"github.com/stretchr/testify/require"
)

func TestReaction_Validate(t *testing.T) {
	tests := []struct {
		name     string
		reaction types.RegisteredReaction
		error    error
	}{
		{
			name: "Valid reaction returns no error (url on value)",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile-jpg:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
		{
			name: "Invalid reaction returns error (unicode on value)",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smiles:",
				"U+1F600",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
		},
		{
			name: "Missing creator returns error",
			reaction: types.NewRegisteredReaction(
				"",
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("invalid reaction creator: "),
		},
		{
			name: "Empty short code returns error",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			//nolint - errcheck
			error: errors.New("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'"),
		},
		{
			name: "Invalid short code returns error",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			//nolint - errcheck
			error: errors.New("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'"),
		},
		{
			name: "Empty value returns error",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile:",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
		},
		{
			name: "invalid value returns error (url)",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile:",
				"smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
		},
		{
			name: "invalid value returns error (unicode)",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile:",
				"U+1",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
		},
		{
			name: "invalid subspace returns no error",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile:",
				"https://smile.jpg",
				"1234",
			),
			error: errors.New("reaction subspace must be a valid sha-256 hash"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.error, test.reaction.Validate())
		})
	}
}

func TestRegisteredReactionsMarshaling(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	reaction := types.NewRegisteredReaction(
		"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
		":smile-jpg:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	marshaled := types.MustMarshalRegisteredReaction(cdc, reaction)
	unmarshaled := types.MustUnmarshalRegisteredReaction(cdc, marshaled)
	require.Equal(t, reaction, unmarshaled)
}

// ___________________________________________________________________________________________________________________

func TestPostReaction_Validate(t *testing.T) {
	tests := []struct {
		name     string
		reaction types.PostReaction
		error    error
	}{
		{
			name: "Valid reaction returns no error",
			reaction: types.NewPostReaction(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				":smile:",
				"reaction",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			error: nil,
		},
		{
			name: "Invalid post id returns error",
			reaction: types.NewPostReaction(
				"",
				":smile:",
				"reaction",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			error: errors.New("invalid post id: "),
		},
		{
			name: "Missing owner returns error",
			reaction: types.NewPostReaction(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				":smile:",
				"reaction",
				"",
			),
			error: errors.New("invalid reaction owner: "),
		},
		{
			name: "Missing value returns error",
			reaction: types.NewPostReaction(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				":smile:",
				"",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			error: errors.New("reaction value cannot be empty or blank"),
		},
		{
			name: "Invalid shortcode returns error",
			reaction: types.NewPostReaction(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"invalid",
				"reaction",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			error: errors.New("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.error, test.reaction.Validate())
		})
	}
}

func TestPostReactionsMarshaling(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	reaction := types.NewPostReaction(
		"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		":smile:",
		"reaction",
		"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
	)
	marshaled := types.MustMarshalPostReaction(cdc, reaction)
	unmarshaled := types.MustUnmarshalPostReaction(cdc, marshaled)
	require.Equal(t, reaction, unmarshaled)
}