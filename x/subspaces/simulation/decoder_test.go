package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/v2/x/subspaces/simulation"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/app"
	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	decoder := simulation.NewDecodeStore(cdc)

	sdkAddr, err := sdk.AccAddressFromBech32("cosmos19r59nc7wfgc5gjnu5ga5yztkvr5qssj24krx2f")
	require.NoError(t, err)

	subspace := types.NewSubspace(
		1,
		"Test subspace",
		"This is a test subspace",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
	)
	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.SubspaceIDKey,
			Value: types.GetSubspaceIDBytes(1),
		},
		{
			Key:   types.SubspaceKey(subspace.ID),
			Value: cdc.MustMarshal(&subspace),
		},
		{
			Key:   types.GroupStoreKey(1, "group"),
			Value: []byte{0x01},
		},
		{
			Key:   types.GroupMemberStoreKey(1, "group", sdkAddr),
			Value: []byte{0x01},
		},
		{
			Key:   types.PermissionStoreKey(1, "group"),
			Value: types.MarshalPermission(types.PermissionWrite),
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
		{"Subspace ID", fmt.Sprintf("SubspaceIDA: %d\nSubspaceIDB: %d\n",
			1, 1)},
		{"Subspace", fmt.Sprintf("SubspaceA: %s\nSubspaceB: %s\n",
			subspace.String(), subspace.String())},
		{"Group", fmt.Sprintf("GroupKeyA: %s\nGroupKeyB: %s\n",
			types.GroupStoreKey(1, "group"), types.GroupStoreKey(1, "group"))},
		{"Group member", fmt.Sprintf("GroupMemberKeyA: %s\nGroupMemberKeyB: %s\n",
			types.GroupMemberStoreKey(1, "group", sdkAddr), types.GroupMemberStoreKey(1, "group", sdkAddr))},
		{"Permission", fmt.Sprintf("PermissionKeyA: %d\nPermissionKeyB: %d\n",
			types.PermissionWrite, types.PermissionWrite)},
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
