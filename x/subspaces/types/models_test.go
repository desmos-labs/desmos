package types_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v6/x/subspaces/types"

	"github.com/stretchr/testify/require"
)

func TestParseSubspaceID(t *testing.T) {
	testCases := []struct {
		name      string
		value     string
		shouldErr bool
		expID     uint64
	}{
		{
			name:      "invalid id returns error",
			value:     "id",
			shouldErr: true,
		},
		{
			name:      "empty value returns zero",
			value:     "",
			shouldErr: false,
			expID:     0,
		},
		{
			name:      "valid id returns correct value",
			value:     "2",
			shouldErr: false,
			expID:     2,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			id, err := types.ParseSubspaceID(tc.value)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expID, id)
			}
		})
	}
}

func TestSubspace_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		subspace  types.Subspace
		shouldErr bool
	}{
		{
			name: "invalid id returns error",
			subspace: types.NewSubspace(
				0,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid name returns error",
			subspace: types.NewSubspace(
				1,
				"",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid treasury returns error",
			subspace: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsx",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid owner returns error",
			subspace: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid creator returns error",
			subspace: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zw",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid creation time returns error",
			subspace: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Time{},
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid subspace fee tokens returns no error",
			subspace: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				sdk.Coins{{Denom: "minttoken", Amount: sdk.NewInt(-10)}},
			),
			shouldErr: true,
		},
		{
			name: "valid subspace returns no error",
			subspace: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.subspace.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSubspace_Update(t *testing.T) {
	testCases := []struct {
		name      string
		subspace  types.Subspace
		update    types.SubspaceUpdate
		expResult types.Subspace
	}{
		{
			name: "nothing is updated when using DoNotModify",
			subspace: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			update: types.NewSubspaceUpdate(
				types.DoNotModify,
				types.DoNotModify,
				types.DoNotModify,
			),
			expResult: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
		},
		{
			name: "each field is updated when edited",
			subspace: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			update: types.NewSubspaceUpdate(
				"New subspace name",
				"New subspace description",
				"cosmos10ya9y35qkf4puaklx5fs07sxfxqncx9usgsnz6",
			),
			expResult: types.NewSubspace(
				1,
				"New subspace name",
				"New subspace description",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos10ya9y35qkf4puaklx5fs07sxfxqncx9usgsnz6",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.subspace.Update(tc.update)
			require.Equal(t, tc.expResult, result)
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestParseSectionID(t *testing.T) {
	testCases := []struct {
		name      string
		value     string
		shouldErr bool
		expID     uint32
	}{
		{
			name:      "invalid id returns error",
			value:     "id",
			shouldErr: true,
		},
		{
			name:      "empty value returns zero",
			value:     "",
			shouldErr: false,
			expID:     0,
		},
		{
			name:      "valid id returns correct value",
			value:     "2",
			shouldErr: false,
			expID:     2,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			id, err := types.ParseSectionID(tc.value)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expID, id)
			}
		})
	}
}

func TestSection_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		section   types.Section
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			section: types.NewSection(
				0,
				0,
				1,
				"Test section",
				"This is a test section",
			),
			shouldErr: true,
		},
		{
			name: "invalid parent id returns error",
			section: types.NewSection(
				1,
				1,
				1,
				"Test section",
				"This is a test section",
			),
			shouldErr: true,
		},
		{
			name: "invalid section name returns error - empty",
			section: types.NewSection(
				1,
				1,
				0,
				"",
				"This is a test section",
			),
			shouldErr: true,
		},
		{
			name: "invalid section name returns error - blank",
			section: types.NewSection(
				1,
				1,
				0,
				"   ",
				"This is a test section",
			),
			shouldErr: true,
		},
		{
			name:      "default section does not return error",
			section:   types.DefaultSection(1),
			shouldErr: false,
		},
		{
			name:      "default section returns no error",
			section:   types.DefaultSection(1),
			shouldErr: false,
		},
		{
			name: "valid data returns no error",
			section: types.NewSection(
				1,
				1,
				0,
				"Test section",
				"This is a test section",
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.section.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSection_Update(t *testing.T) {
	testCases := []struct {
		name      string
		section   types.Section
		update    types.SectionUpdate
		expResult types.Section
	}{
		{
			name: "nothing is updated when using DoNotModify",
			section: types.NewSection(
				1,
				1,
				0,
				"Test section",
				"This is a test section",
			),
			update: types.NewSectionUpdate(
				types.DoNotModify,
				types.DoNotModify,
			),
			expResult: types.NewSection(
				1,
				1,
				0,
				"Test section",
				"This is a test section",
			),
		},
		{
			name: "each field is updated when edited",
			section: types.NewSection(
				1,
				1,
				0,
				"Test section",
				"This is a test section",
			),
			update: types.NewSectionUpdate(
				"New section name",
				"New section description",
			),
			expResult: types.NewSection(
				1,
				1,
				0,
				"New section name",
				"New section description",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.section.Update(tc.update)
			require.Equal(t, tc.expResult, result)
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestParseUserGroupID(t *testing.T) {
	testCases := []struct {
		name      string
		value     string
		shouldErr bool
		expID     uint32
	}{
		{
			name:      "invalid id returns error",
			value:     "id",
			shouldErr: true,
		},
		{
			name:      "empty value returns zero",
			value:     "",
			shouldErr: false,
			expID:     0,
		},
		{
			name:      "valid id returns correct value",
			value:     "2",
			shouldErr: false,
			expID:     2,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			id, err := types.ParseGroupID(tc.value)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expID, id)
			}
		})
	}
}

func TestUserGroup_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		group     types.UserGroup
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			group: types.NewUserGroup(
				0,
				0,
				1,
				"Test group",
				"This is a test group",
				types.NewPermissions(types.PermissionEditSubspace),
			),
			shouldErr: true,
		},
		{
			name: "invalid group name returns error - empty",
			group: types.NewUserGroup(
				1,
				0,
				1,
				"",
				"This is a test group",
				types.NewPermissions(types.PermissionEditSubspace),
			),
			shouldErr: true,
		},
		{
			name: "invalid group name returns error - blank",
			group: types.NewUserGroup(
				1,
				0,
				1,
				"  ",
				"This is a test group",
				types.NewPermissions(types.PermissionEditSubspace),
			),
			shouldErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.group.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserGroup_Update(t *testing.T) {
	testCases := []struct {
		name      string
		subspace  types.UserGroup
		update    types.GroupUpdate
		expResult types.UserGroup
	}{
		{
			name: "nothing is updated when using DoNotModify",
			subspace: types.NewUserGroup(
				1,
				0,
				1,
				"Test group",
				"This is a test group",
				types.NewPermissions(types.PermissionEditSubspace),
			),
			update: types.NewGroupUpdate(
				types.DoNotModify,
				types.DoNotModify,
			),
			expResult: types.NewUserGroup(
				1,
				0,
				1,
				"Test group",
				"This is a test group",
				types.NewPermissions(types.PermissionEditSubspace),
			),
		},
		{
			name: "each field is updated when edited",
			subspace: types.NewUserGroup(
				1,
				0,
				1,
				"Test group",
				"This is a test group",
				types.NewPermissions(types.PermissionEditSubspace),
			),
			update: types.NewGroupUpdate(
				"New group name",
				"New group description",
			),
			expResult: types.NewUserGroup(
				1,
				0,
				1,
				"New group name",
				"New group description",
				types.NewPermissions(types.PermissionEditSubspace),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := tc.subspace.Update(tc.update)
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestGetTreasuryAddress(t *testing.T) {
	subspaceIDs := getRandomSubspaceIDs(1000)
	treasuryAddrSet := make(map[string]bool, 1000)
	for _, subspaceID := range subspaceIDs {
		treasuryAddr := types.GetTreasuryAddress(subspaceID).String()

		// Ensure that address is generated using `authtypes.NewModuleAddress` and each address is unique
		require.Equal(t, authtypes.NewModuleAddress(fmt.Sprintf("subspace-%d", subspaceID)).String(), treasuryAddr)
		require.False(t, treasuryAddrSet[treasuryAddr])
		treasuryAddrSet[treasuryAddr] = true
	}
}

func getRandomSubspaceIDs(size int) []uint64 {
	subspaceIDs := make([]uint64, size)
	subspaceIDsSet := make(map[uint64]bool, size)
	for i := 0; i < size; {
		newID := rand.Uint64()
		// Add it if new ID is not found in the set
		if found := subspaceIDsSet[newID]; !found {
			subspaceIDs[i] = newID
			subspaceIDsSet[newID] = true
			i++
		}
	}
	return subspaceIDs
}

func TestAdditionalFeeTokensUpdate_Update(t *testing.T) {
	subspace := types.NewSubspace(
		1,
		"Test subspace",
		"This is a test subspace",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
		nil,
	)

	updated := types.NewAdditionalFeeTokensUpdate(
		sdk.NewCoin("minttoken", sdk.NewInt(10)),
		sdk.NewCoin("ustar", sdk.NewInt(10)),
	).Update(subspace)

	require.Equal(t, types.NewSubspace(
		1,
		"Test subspace",
		"This is a test subspace",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
		time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
		sdk.NewCoins(
			sdk.NewCoin("minttoken", sdk.NewInt(10)),
			sdk.NewCoin("ustar", sdk.NewInt(10)),
		),
	), updated)
}
