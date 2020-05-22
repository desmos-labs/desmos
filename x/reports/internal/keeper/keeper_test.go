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
		test := test
		t.Run(test.name, func(t *testing.T) {
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
		})
	}
}

func TestKeeper_RegisterReportTypes(t *testing.T) {
	tests := []struct {
		name             string
		existentRepTypes types.ReportTypes
		repType          types.ReportType
		expBool          bool
		expRepTypes      types.ReportTypes
	}{
		{
			name: "non existent rep type inserted correctly returns true",
			existentRepTypes: types.ReportTypes{
				"spam",
				"offense",
			},
			repType: "nudity",
			expBool: true,
			expRepTypes: types.ReportTypes{
				"spam",
				"offense",
				"nudity",
			},
		},
		{
			name: "existent rep type non inserted returns false",
			existentRepTypes: types.ReportTypes{
				"spam",
				"nudity",
			},
			repType: "nudity",
			expBool: false,
			expRepTypes: types.ReportTypes{
				"spam",
				"nudity",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			if test.existentRepTypes != nil {
				store.Set(types.ReportsTypeStorePrefix, k.Cdc.MustMarshalBinaryBare(&test.existentRepTypes))
			}

			actual := k.RegisterReportsTypes(ctx, test.repType)

			require.Equal(t, test.expBool, actual)

			var reportTypes types.ReportTypes
			bz := store.Get(types.ReportsTypeStorePrefix)
			k.Cdc.MustUnmarshalBinaryBare(bz, &reportTypes)

			require.Equal(t, test.expRepTypes, reportTypes)
		})
	}
}

func TestKeeper_GetRegisteredReportsTypes(t *testing.T) {
	tests := []struct {
		name             string
		existentRepTypes types.ReportTypes
	}{
		{
			name: "non empty reports types array is returned correctly",
			existentRepTypes: types.ReportTypes{
				"spam",
				"offense",
				"nudity",
			},
		},
		{
			name:             "empty reports types array is returned correctly",
			existentRepTypes: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			if test.existentRepTypes != nil {
				store.Set(types.ReportsTypeStorePrefix, k.Cdc.MustMarshalBinaryBare(&test.existentRepTypes))
			}

			actual := k.GetRegisteredReportsTypes(ctx)

			require.Equal(t, test.existentRepTypes, actual)
		})
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
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			if test.expReports != nil {
				store.Set(types.ReportStoreKey(postID), k.Cdc.MustMarshalBinaryBare(&test.expReports))
			}

			actualRep := k.GetPostReports(ctx, postID)
			require.Equal(t, test.expReports, actualRep)
		})
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
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			if test.existingReports != nil {
				store.Set(types.ReportStoreKey(postID), k.Cdc.MustMarshalBinaryBare(&test.existingReports))
			}

			actualRep := k.GetReportsMap(ctx)
			require.Equal(t, test.expReportsMap, actualRep)
		})
	}
}
