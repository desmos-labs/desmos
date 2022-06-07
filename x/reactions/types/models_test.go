package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/reactions/types"
)

func TestReaction_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		reaction  types.Reaction
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			reaction: types.NewReaction(
				0,
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
			),
			shouldErr: true,
		},
		{
			name: "invalid id returns error",
			reaction: types.NewReaction(
				1,
				0,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			reaction: types.NewReaction(
				1,
				1,
				0,
				types.NewRegisteredReactionValue(1),
				"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
			),
			shouldErr: true,
		},
		{
			name: "invalid value returns error",
			reaction: types.NewReaction(
				1,
				1,
				1,
				types.NewRegisteredReactionValue(0),
				"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
			),
			shouldErr: true,
		},
		{
			name: "invalid author returns error",
			reaction: types.NewReaction(
				0,
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			reaction: types.NewReaction(
				1,
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.reaction.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestRegisteredReactionValue_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		value     *types.RegisteredReactionValue
		shouldErr bool
	}{
		{
			name:      "invalid registered reaction id returns error",
			value:     types.NewRegisteredReactionValue(0),
			shouldErr: true,
		},
		{
			name:      "valid registered reaction id returns no error",
			value:     types.NewRegisteredReactionValue(1),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.value.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestFreeTextValue_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		value     *types.FreeTextValue
		shouldErr bool
	}{
		{
			name:      "empty text returns error",
			value:     types.NewFreeTextValue(""),
			shouldErr: true,
		},
		{
			name:      "blank text returns error",
			value:     types.NewFreeTextValue(" "),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			value:     types.NewFreeTextValue("Wow!"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.value.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestRegisteredReaction_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		reaction  types.RegisteredReaction
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			reaction: types.NewRegisteredReaction(
				0,
				1,
				":hello:",
				"https://example.com?image=hello.png",
			),
			shouldErr: true,
		},
		{
			name: "invalid id returns error",
			reaction: types.NewRegisteredReaction(
				1,
				0,
				":hello:",
				"https://example.com?image=hello.png",
			),
			shouldErr: true,
		},
		{
			name: "invalid shorthand code returns error",
			reaction: types.NewRegisteredReaction(
				1,
				1,
				" ",
				"https://example.com?image=hello.png",
			),
			shouldErr: true,
		},
		{
			name: "invalid display value returns error",
			reaction: types.NewRegisteredReaction(
				1,
				1,
				":hello:",
				" ",
			),
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			reaction: types.NewRegisteredReaction(
				1,
				1,
				":hello:",
				"https://example.com?image=hello.png",
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.reaction.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestSubspaceReactionsParams_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.SubspaceReactionsParams
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			params: types.NewSubspaceReactionsParams(
				0,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 100, ""),
			),
			shouldErr: true,
		},
		{
			name: "invalid free text value params returns error",
			params: types.NewSubspaceReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 0, ""),
			),
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			params: types.NewSubspaceReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 100, ".{1,3}"),
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestFreeTextValueParams_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.FreeTextValueParams
		shouldErr bool
	}{
		{
			name:      "invalid max length returns error",
			params:    types.NewFreeTextValueParams(true, 0, ""),
			shouldErr: true,
		},
		{
			name:      "invalid regex returns error",
			params:    types.NewFreeTextValueParams(true, 10, ".*{1,2}"),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			params:    types.NewFreeTextValueParams(true, 10, "[a-zA-Z]"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
