package v0150_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v0130reports "github.com/desmos-labs/desmos/x/reports/legacy/v0.13.0"
	v0150reports "github.com/desmos-labs/desmos/x/reports/legacy/v0.15.0"
)

func TestMigrate(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	v0130GenState := v0130reports.GenesisState{
		Reports: map[string][]v0130reports.Report{
			"postID": {
				{
					Type:    "test",
					Message: "test",
					User:    user,
				},
				{
					Type:    "txt",
					Message: "txt",
					User:    user,
				},
			},
		},
	}

	expGenState := v0150reports.GenesisState{
		Reports: []v0150reports.Report{
			{
				PostID:  "postID",
				Type:    "test",
				Message: "test",
				User:    user.String(),
			},
			{
				PostID:  "postID",
				Type:    "txt",
				Message: "txt",
				User:    user.String(),
			},
		},
	}

	migrated := v0150reports.Migrate(v0130GenState)

	require.Len(t, expGenState.Reports, len(migrated.Reports))
	for index, report := range migrated.Reports {
		require.Equal(t, expGenState.Reports[index], report)
	}
}
