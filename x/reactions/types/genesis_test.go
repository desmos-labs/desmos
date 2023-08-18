package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/x/reactions/types"
)

func TestValidateGenesis(t *testing.T) {
	testCases := []struct {
		name      string
		genesis   *types.GenesisState
		shouldErr bool
	}{
		{
			name: "invalid subspace data returns error",
			genesis: types.NewGenesisState([]types.SubspaceDataEntry{
				types.NewSubspaceDataEntry(0, 1),
			}, nil, nil, nil, nil),
			shouldErr: true,
		},
		{
			name: "duplicated subspace data returns error",
			genesis: types.NewGenesisState([]types.SubspaceDataEntry{
				types.NewSubspaceDataEntry(1, 1),
				types.NewSubspaceDataEntry(1, 1),
			}, nil, nil, nil, nil),
			shouldErr: true,
		},
		{
			name: "invalid registered reaction returns error",
			genesis: types.NewGenesisState(nil, []types.RegisteredReaction{
				types.NewRegisteredReaction(
					0,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				),
			}, nil, nil, nil),
			shouldErr: true,
		},
		{
			name: "duplicated registered reaction returns error",
			genesis: types.NewGenesisState(nil, []types.RegisteredReaction{
				types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				),
				types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				),
			}, nil, nil, nil),
			shouldErr: true,
		},
		{
			name: "invalid post data entry returns error",
			genesis: types.NewGenesisState(nil, nil, []types.PostDataEntry{
				types.NewPostDataEntry(0, 1, 1),
			}, nil, nil),
			shouldErr: true,
		},
		{
			name: "duplicated post data entry returns error",
			genesis: types.NewGenesisState(nil, nil, []types.PostDataEntry{
				types.NewPostDataEntry(1, 1, 1),
				types.NewPostDataEntry(1, 1, 1),
			}, nil, nil),
			shouldErr: true,
		},
		{
			name: "invalid reaction returns error",
			genesis: types.NewGenesisState(nil, nil, nil, []types.Reaction{
				types.NewReaction(
					0,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				),
			}, nil),
			shouldErr: true,
		},
		{
			name: "duplicated reaction returns error",
			genesis: types.NewGenesisState(nil, nil, nil, []types.Reaction{
				types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				),
				types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				),
			}, nil),
			shouldErr: true,
		},
		{
			name: "invalid subspace params returns error",
			genesis: types.NewGenesisState(nil, nil, nil, nil, []types.SubspaceReactionsParams{
				types.NewSubspaceReactionsParams(
					0,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 100, ""),
				),
			}),
			shouldErr: true,
		},
		{
			name: "duplicated subspace params returns error",
			genesis: types.NewGenesisState(nil, nil, nil, nil, []types.SubspaceReactionsParams{
				types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 100, ""),
				),
				types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 100, ""),
				),
			}),
			shouldErr: true,
		},
		{
			name:      "default genesis returns no error",
			genesis:   types.DefaultGenesisState(),
			shouldErr: false,
		},
		{
			name: "valid data returns no error",
			genesis: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 1),
				},
				[]types.RegisteredReaction{
					types.NewRegisteredReaction(
						1,
						1,
						":hello:",
						"https://example.com?image=hello.png",
					),
				},
				[]types.PostDataEntry{
					types.NewPostDataEntry(1, 1, 2),
				},
				[]types.Reaction{
					types.NewReaction(
						1,
						1,
						1,
						types.NewRegisteredReactionValue(1),
						"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
					),
				},
				[]types.SubspaceReactionsParams{
					types.NewSubspaceReactionsParams(
						1,
						types.NewRegisteredReactionValueParams(true),
						types.NewFreeTextValueParams(true, 100, ""),
					),
				},
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateGenesis(tc.genesis)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestSubspaceDataEntry_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		entry     types.SubspaceDataEntry
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			entry:     types.NewSubspaceDataEntry(0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid registered reaction id returns error",
			entry:     types.NewSubspaceDataEntry(1, 0),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			entry:     types.NewSubspaceDataEntry(1, 1),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.entry.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPostDataEntry_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		entry     types.PostDataEntry
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			entry:     types.NewPostDataEntry(0, 1, 1),
			shouldErr: true,
		},
		{
			name:      "invalid post id id returns error",
			entry:     types.NewPostDataEntry(1, 0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid reaction id returns error",
			entry:     types.NewPostDataEntry(1, 1, 0),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			entry:     types.NewPostDataEntry(1, 1, 1),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.entry.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
