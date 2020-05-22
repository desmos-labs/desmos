package keeper_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/desmos-labs/desmos/x/reports/internal/types/models"
	"github.com/stretchr/testify/require"
)

func TestKeeper_SaveReport(t *testing.T) {
	tests := []struct {
		name            string
		report          models.Report
		existentReports models.Reports
		expAdd          bool
		expReports      models.Reports
	}{
		{
			name:            "New reports added correctly",
			report:          models.NewReport("type", "message", creator),
			existentReports: nil,
			expAdd:          true,
			expReports:      models.Reports{models.NewReport("type", "message", creator)},
		},
		{
			name:   "Existent reports not added",
			report: models.NewReport("type", "message", creator),
			existentReports: models.Reports{
				{Type: "type", Message: "message", User: creator},
			},
			expAdd:     false,
			expReports: models.Reports{models.NewReport("type", "message", creator)},
		},
	}

	for _, test := range tests {
		ctx, k := SetupTestInput()
		store := ctx.KVStore(k.StoreKey)
		if test.existentReports != nil {
			store.Set(types.ReportStoreKey(postID), k.Cdc.MustMarshalBinaryBare(&test.existentReports))
		}

		actual := k.SaveReport(ctx, postID, test.report)

		require.Equal(t, test.expAdd, actual)

		var reports models.Reports
		bz := store.Get(types.ReportStoreKey(postID))
		k.Cdc.MustUnmarshalBinaryBare(bz, &reports)

		require.Equal(t, test.expReports, reports)
	}
}

func TestKeeper_GetPostReports(t *testing.T) {
	tests := []struct {
		name       string
		expReports models.Reports
	}{
		{
			name: "Returns a non-empty reports array",
			expReports: models.Reports{
				{Type: "type", Message: "message", User: creator},
			},
		},
		{
			name:       "Returns an empty reports array",
			expReports: nil,
		},
	}

	for _, test := range tests {
		ctx, k := SetupTestInput()
		store := ctx.KVStore(k.StoreKey)
		if test.expReports != nil {
			store.Set(types.ReportStoreKey(postID), k.Cdc.MustMarshalBinaryBare(&test.expReports))
		}

		actualRep := k.GetPostReports(ctx, postID)
		require.Equal(t, test.expReports, actualRep)
	}
}

func TestKeeper_GetReportsMap(t *testing.T) {
	reports := models.Reports{
		{Type: "type", Message: "message", User: creator},
	}
	tests := []struct {
		name            string
		existingReports models.Reports
		expReportsMap   map[string]models.Reports
	}{
		{
			name:            "Returns a non-empty reports map",
			existingReports: reports,
			expReportsMap: map[string]models.Reports{
				postID.String(): reports,
			},
		},
		{
			name:            "Returns an empty reports map",
			existingReports: nil,
			expReportsMap:   map[string]models.Reports{},
		},
	}

	for _, test := range tests {
		ctx, k := SetupTestInput()
		store := ctx.KVStore(k.StoreKey)
		if test.existingReports != nil {
			store.Set(types.ReportStoreKey(postID), k.Cdc.MustMarshalBinaryBare(&test.existingReports))
		}

		actualRep := k.GetReportsMap(ctx)
		require.Equal(t, test.expReportsMap, actualRep)
	}
}
