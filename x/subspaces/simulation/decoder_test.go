package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/v4/x/subspaces/simulation"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	decoder := simulation.NewDecodeStore(cdc)

	subspace := types.NewSubspace(
		1,
		"Test subspace",
		"This is a test subspace",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
	)
	group := types.NewUserGroup(
		1,
		0,
		1,
		"Test group",
		"This is a test group",
		types.NewPermissions(types.PermissionEditSubspace),
	)

	section := types.NewSection(
		1,
		1,
		0,
		"Test section",
		"This is a test section",
	)

	permission := types.NewUserPermission(
		1,
		1,
		"cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
		types.NewPermissions(types.PermissionEverything),
	)

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.SubspaceIDKey,
			Value: types.GetSubspaceIDBytes(1),
		},
		{
			Key:   types.SubspaceStoreKey(subspace.ID),
			Value: cdc.MustMarshal(&subspace),
		},
		{
			Key:   types.NextGroupIDStoreKey(1),
			Value: types.GetGroupIDBytes(1),
		},
		{
			Key:   types.GroupStoreKey(1, 0, 1),
			Value: cdc.MustMarshal(&group),
		},
		{
			Key:   types.GroupMemberStoreKey(1, 1, "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e"),
			Value: []byte{0x01},
		},
		{
			Key:   types.UserPermissionStoreKey(1, 0, "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e"),
			Value: cdc.MustMarshal(&permission),
		},
		{
			Key:   types.NextSectionIDStoreKey(1),
			Value: types.GetSectionIDBytes(1),
		},
		{
			Key:   types.SectionStoreKey(1, 1),
			Value: cdc.MustMarshal(&section),
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
		{"Group ID", fmt.Sprintf("GroupIDA: %d\nGroupIDB: %d\n",
			1, 1)},
		{"Group", fmt.Sprintf("GroupA: %s\nGroupB: %s\n",
			group.String(), group.String())},
		{"Group member", fmt.Sprintf("GroupMemberKeyA: %s\nGroupMemberKeyB: %s\n",
			types.GroupMemberStoreKey(1, 1, "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e"), types.GroupMemberStoreKey(1, 1, "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e"))},
		{"Permission", fmt.Sprintf("PermissionA: %s\nPermissionB: %s\n",
			&permission, &permission)},
		{"Section ID", fmt.Sprintf("SectionIDA: %d\nSectionIDB: %d\n", 1, 1)},
		{"Section", fmt.Sprintf("SectionA: %s\nSectionB: %s\n", &section, &section)},
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
