package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetReportID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reportID   uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing report id is set properly",
			subspaceID: 1,
			reportID:   1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetReportIDFromBytes(store.Get(types.ReportIDStoreKey(1)))
				suite.Require().Equal(uint64(1), stored)
			},
		},
		{
			name: "existing report id is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SetReportID(ctx, 1, 1)
			},
			subspaceID: 1,
			reportID:   2,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetReportIDFromBytes(store.Get(types.ReportIDStoreKey(1)))
				suite.Require().Equal(uint64(2), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.SetReportID(ctx, tc.subspaceID, tc.reportID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetReportID() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		shouldErr   bool
		expReportID uint64
	}{
		{
			name:       "non existing report id returns error",
			subspaceID: 1,
			shouldErr:  true,
		},
		{
			name: "existing report id is returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetReportID(ctx, 1, 1)
			},
			subspaceID:  1,
			shouldErr:   false,
			expReportID: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			reportID, err := suite.k.GetReportID(ctx, tc.subspaceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expReportID, reportID)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteReportID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name: "existing report id is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SetReportID(ctx, 1, 1)
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.ReportIDStoreKey(1)))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.DeleteReportID(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestsuite) TestKeeper_SaveReport() {
	testCases := []struct {
		name   string
		store  func(ctx sdk.Context)
		report types.Report
		check  func(ctx sdk.Context)
	}{
		{
			name: "non existing report is stored properly",
			report: types.NewReport(
				1,
				1,
				1,
				"This content is spam",
				"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
				types.NewPostData(1),
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetReport(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReport(
					1,
					1,
					1,
					"This content is spam",
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					types.NewPostData(1),
				), stored)
			},
		},
		{
			name: "existing report is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					1,
					"This content is spam",
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					types.NewPostData(1),
				))
			},
			report: types.NewReport(
				1,
				1,
				2,
				"This content contains self harm",
				"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
				types.NewPostData(5),
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetReport(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReport(
					1,
					1,
					2,
					"This content contains self harm",
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					types.NewPostData(5),
				), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.SaveReport(ctx, tc.report)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasReport() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reportID   uint64
		expResult  bool
	}{
		{
			name:       "non existing report returns false",
			subspaceID: 1,
			reportID:   1,
			expResult:  false,
		},
		{
			name: "existing report returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					1,
					"This content is spam",
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					types.NewPostData(1),
				))
			},
			subspaceID: 1,
			reportID:   1,
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasReport(ctx, tc.subspaceID, tc.reportID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetReport() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reportID   uint64
		expFound   bool
		expReport  types.Report
	}{
		{
			name:       "non existing report returns false and empty report",
			subspaceID: 1,
			reportID:   1,
			expFound:   false,
			expReport:  types.Report{},
		},
		{
			name: "existing report returns true and correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					1,
					"This content is spam",
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					types.NewPostData(1),
				))
			},
			subspaceID: 1,
			reportID:   1,
			expFound:   true,
			expReport: types.NewReport(
				1,
				1,
				1,
				"This content is spam",
				"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
				types.NewPostData(1),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			report, found := suite.k.GetReport(ctx, tc.subspaceID, tc.reportID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expReport, report)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteReport() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reportID   uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing report is deleted properly",
			subspaceID: 1,
			reportID:   1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasReport(ctx, 1, 1))
			},
		},
		{
			name: "existing report is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					1,
					"This content is spam",
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					types.NewPostData(1),
				))
			},
			subspaceID: 1,
			reportID:   1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasReport(ctx, 1, 1))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.DeleteReport(ctx, tc.subspaceID, tc.reportID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
