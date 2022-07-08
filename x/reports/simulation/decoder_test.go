package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/v4/x/reports/simulation"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/x/reports/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	decoder := simulation.NewDecodeStore(cdc)

	report := types.NewReport(
		1,
		1,
		[]uint32{1},
		"This user is spamming",
		types.NewUserTarget("cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e"),
		"cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79",
		time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
	)
	reason := types.NewReason(
		1,
		1,
		"Spam",
		"This content is spam",
	)

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.NextReportIDStoreKey(1),
			Value: types.GetReportIDBytes(1),
		},
		{
			Key:   types.ReportStoreKey(1, 1),
			Value: cdc.MustMarshal(&report),
		},
		{
			Key:   types.PostReportStoreKey(1, 1, "cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
			Value: types.GetReportIDBytes(1),
		},
		{
			Key:   types.UserReportStoreKey(1, "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e", "cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
			Value: types.GetReportIDBytes(1),
		},
		{
			Key:   types.NextReasonIDStoreKey(1),
			Value: types.GetReasonIDBytes(1),
		},
		{
			Key:   types.ReasonStoreKey(1, 1),
			Value: cdc.MustMarshal(&reason),
		},
		{
			Key:   []byte("Unknown key"),
			Value: nil,
		},
	}}

	testCases := []struct {
		name        string
		expectedLog string
	}{
		{"Next Report ID", fmt.Sprintf("NextReportIDA: %d\nNextReportIDB: %d\n",
			1, 1)},
		{"Report", fmt.Sprintf("ReportA: %s\nReportB: %s\n",
			&report, &report)},
		{"Post Report ID", fmt.Sprintf("PostReportIDA: %d\nPostReportIDB: %d\n",
			1, 1)},
		{"User Report ID", fmt.Sprintf("UserReportIDA: %d\nUserReportIDB: %d\n",
			1, 1)},
		{"Next Reason ID", fmt.Sprintf("NextReasonIDA: %d\nNextReasonIDB: %d\n",
			1, 1)},
		{"Reason", fmt.Sprintf("ReasonA: %s\nReasonB: %s\n",
			&reason, &reason)},
		{"other", ""},
	}

	for i, tc := range testCases {
		i, tc := i, tc
		t.Run(tc.name, func(t *testing.T) {
			switch i {
			case len(testCases) - 1:
				require.Panics(t, func() { decoder(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tc.name)
			default:
				require.Equal(t, tc.expectedLog, decoder(kvPairs.Pairs[i], kvPairs.Pairs[i]), tc.name)
			}
		})
	}
}
