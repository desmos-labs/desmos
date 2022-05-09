package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

func TestValidateGenesis(t *testing.T) {
	testCases := []struct {
		name      string
		data      *types.GenesisState
		shouldErr bool
	}{
		{
			name: "duplicated subspaces data returns error",
			data: types.NewGenesisState([]types.SubspaceData{
				types.NewSubspacesData(1, 1, 1, nil),
				types.NewSubspacesData(1, 1, 1, nil),
			}, nil, types.DefaultParams()),
			shouldErr: true,
		},
		{
			name: "invalid subspaces data returns error",
			data: types.NewGenesisState([]types.SubspaceData{
				types.NewSubspacesData(0, 1, 1, nil),
			}, nil, types.DefaultParams()),
			shouldErr: true,
		},
		{
			name: "duplicated report returns error",
			data: types.NewGenesisState(nil, []types.Report{
				types.NewReport(
					1,
					1,
					1,
					"",
					"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
					types.NewPostData(1),
				),
				types.NewReport(
					1,
					1,
					1,
					"",
					"cosmos1x85xq5m2ehkjzw928j9zfv3awdy0hqtnhrp9r6",
					types.NewPostData(1),
				),
			}, types.DefaultParams()),
			shouldErr: true,
		},
		{
			name: "invalid report id returns error",
			data: types.NewGenesisState(
				[]types.SubspaceData{
					types.NewSubspacesData(1, 1, 1, nil),
				},
				[]types.Report{
					types.NewReport(
						1,
						1,
						1,
						"",
						"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
						types.NewPostData(1),
					),
				},
				types.DefaultParams(),
			),
			shouldErr: true,
		},
		{
			name: "invalid report returns error",
			data: types.NewGenesisState(
				[]types.SubspaceData{
					types.NewSubspacesData(1, 2, 1, nil),
				},
				[]types.Report{
					types.NewReport(
						0,
						1,
						1,
						"",
						"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
						types.NewPostData(1),
					),
				},
				types.DefaultParams(),
			),
			shouldErr: true,
		},
		{
			name: "invalid params returns error",
			data: types.NewGenesisState(nil, nil, types.NewParams([]types.Reason{
				types.NewReason(0, "", ""),
			})),
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			data: types.NewGenesisState(
				[]types.SubspaceData{
					types.NewSubspacesData(1, 2, 1, nil),
				},
				[]types.Report{
					types.NewReport(
						1,
						1,
						1,
						"",
						"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
						types.NewPostData(1),
					),
				},
				types.DefaultParams(),
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateGenesis(tc.data)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestSubspaceData_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		data      types.SubspaceData
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			data:      types.NewSubspacesData(0, 1, 1, nil),
			shouldErr: true,
		},
		{
			name:      "invalid report id returns error",
			data:      types.NewSubspacesData(1, 0, 1, nil),
			shouldErr: true,
		},
		{
			name:      "invalid reason id returns error",
			data:      types.NewSubspacesData(1, 1, 0, nil),
			shouldErr: true,
		},
		{
			name: "too high reason id returns error",
			data: types.NewSubspacesData(1, 1, 1, []types.Reason{
				types.NewReason(1, "Spam", "This content is spam or the poster is a spammer"),
			}),
			shouldErr: true,
		},
		{
			name: "invalid reason returns error",
			data: types.NewSubspacesData(1, 1, 2, []types.Reason{
				types.NewReason(0, "Spam", "This content is spam or the poster is a spammer"),
			}),
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			data: types.NewSubspacesData(1, 1, 2, []types.Reason{
				types.NewReason(1, "Spam", "This content is spam or the poster is a spammer"),
			}),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.data.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
