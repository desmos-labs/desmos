package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/x/reports/types"
)

func TestValidateGenesis(t *testing.T) {
	testCases := []struct {
		name      string
		data      *types.GenesisState
		shouldErr bool
	}{
		{
			name: "duplicated subspaces data returns error",
			data: types.NewGenesisState([]types.SubspaceDataEntry{
				types.NewSubspacesDataEntry(1, 1, 1),
				types.NewSubspacesDataEntry(1, 1, 1),
			}, nil, nil, types.DefaultParams()),
			shouldErr: true,
		},
		{
			name: "invalid subspaces data returns error",
			data: types.NewGenesisState([]types.SubspaceDataEntry{
				types.NewSubspacesDataEntry(0, 1, 1),
			}, nil, nil, types.DefaultParams()),
			shouldErr: true,
		},
		{
			name: "duplicated reason returns error",
			data: types.NewGenesisState(nil, []types.Reason{
				types.NewReason(1, 1, "Spam", ""),
				types.NewReason(1, 1, "Spam", ""),
			}, nil, types.DefaultParams()),
			shouldErr: true,
		},
		{
			name: "duplicated report returns error",
			data: types.NewGenesisState(nil, nil, []types.Report{
				types.NewReport(
					1,
					1,
					[]uint32{1},
					"",
					types.NewPostTarget(1),
					"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
				types.NewReport(
					1,
					1,
					[]uint32{1},
					"",
					types.NewPostTarget(1),
					"cosmos1x85xq5m2ehkjzw928j9zfv3awdy0hqtnhrp9r6",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
			}, types.DefaultParams()),
			shouldErr: true,
		},
		{
			name: "invalid report returns error",
			data: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspacesDataEntry(1, 1, 2),
				},
				nil,
				[]types.Report{
					types.NewReport(
						0,
						1,
						[]uint32{1},
						"",
						types.NewPostTarget(1),
						"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
				},
				types.DefaultParams(),
			),
			shouldErr: true,
		},
		{
			name: "invalid params returns error",
			data: types.NewGenesisState(nil, nil, nil, types.NewParams(types.NewStandardReasons(
				types.NewStandardReason(0, "", ""),
			))),
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			data: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspacesDataEntry(1, 2, 2),
				},
				[]types.Reason{
					types.NewReason(1, 1, "Spam", ""),
				},
				[]types.Report{
					types.NewReport(
						1,
						1,
						[]uint32{1},
						"",
						types.NewPostTarget(1),
						"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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

func TestSubspaceDataEntry_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		data      types.SubspaceDataEntry
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			data:      types.NewSubspacesDataEntry(0, 1, 1),
			shouldErr: true,
		},
		{
			name:      "invalid reason id returns error",
			data:      types.NewSubspacesDataEntry(1, 0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid report id returns error",
			data:      types.NewSubspacesDataEntry(1, 1, 0),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			data:      types.NewSubspacesDataEntry(1, 1, 2),
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
