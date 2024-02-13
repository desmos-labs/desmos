package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v7/x/reports/types"
)

func (suite *KeeperTestSuite) TestQueryServer_Reports() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		request    *types.QueryReportsRequest
		shouldErr  bool
		expReports []types.Report
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryReportsRequest(0, nil, "", nil),
			shouldErr: true,
		},
		{
			name: "user reports request returns correct data - with reporter",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					2,
					[]uint32{1},
					"This post is spam",
					types.NewPostTarget(1),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			request: types.NewQueryReportsRequest(
				1,
				types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
				"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
				nil,
			),
			shouldErr: false,
			expReports: []types.Report{
				types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
			},
		},
		{
			name: "user reports request returns correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					2,
					[]uint32{1},
					"This post is spam",
					types.NewPostTarget(1),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			request: types.NewQueryReportsRequest(
				1,
				types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
				"",
				&query.PageRequest{
					Limit: 1,
				},
			),
			shouldErr: false,
			expReports: []types.Report{
				types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
			},
		},
		{
			name: "post reports request returns correct data - with reporter",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					2,
					[]uint32{1},
					"This post is spam",
					types.NewPostTarget(1),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			request: types.NewQueryReportsRequest(
				1,
				types.NewPostTarget(1),
				"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
				nil,
			),
			shouldErr: false,
			expReports: []types.Report{
				types.NewReport(
					1,
					2,
					[]uint32{1},
					"This post is spam",
					types.NewPostTarget(1),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
			},
		},
		{
			name: "post reports request returns correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					2,
					[]uint32{1},
					"This post is spam",
					types.NewPostTarget(1),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			request: types.NewQueryReportsRequest(
				1,
				types.NewPostTarget(1),
				"",
				&query.PageRequest{
					Limit: 1,
				},
			),
			shouldErr: false,
			expReports: []types.Report{
				types.NewReport(
					1,
					2,
					[]uint32{1},
					"This post is spam",
					types.NewPostTarget(1),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
			},
		},
		{
			name: "generic reports request returns correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					2,
					[]uint32{1},
					"This post is spam",
					types.NewPostTarget(1),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			request: types.NewQueryReportsRequest(
				1,
				nil,
				"",
				&query.PageRequest{
					Limit: 1,
				},
			),
			shouldErr: false,
			expReports: []types.Report{
				types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
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

			res, err := suite.k.Reports(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expReports, res.Reports)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Report() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QueryReportRequest
		shouldErr bool
		expReport types.Report
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryReportRequest(0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid report id returns error",
			request:   types.NewQueryReportRequest(1, 0),
			shouldErr: true,
		},
		{
			name:      "not found report returns error",
			request:   types.NewQueryReportRequest(1, 1),
			shouldErr: true,
		},
		{
			name: "valid request returns correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			request:   types.NewQueryReportRequest(1, 1),
			shouldErr: false,
			expReport: types.NewReport(
				1,
				1,
				[]uint32{1},
				"This user is spamming",
				types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
				"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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

			res, err := suite.k.Report(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expReport, res.Report)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Reasons() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		request    *types.QueryReasonsRequest
		shouldErr  bool
		expReasons []types.Reason
	}{
		{
			name: "paginated request returns the correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					2,
					"Self harm",
					"This content contains self harm",
				))
			},
			request: types.NewQueryReasonsRequest(
				1,
				&query.PageRequest{Limit: 1},
			),
			shouldErr: false,
			expReasons: []types.Reason{
				types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				),
			},
		},
		{
			name: "non paginated request returns the correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					2,
					"Self harm",
					"This content contains self harm",
				))
			},
			request:   types.NewQueryReasonsRequest(1, nil),
			shouldErr: false,
			expReasons: []types.Reason{
				types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				),
				types.NewReason(
					1,
					2,
					"Self harm",
					"This content contains self harm",
				),
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

			res, err := suite.k.Reasons(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expReasons, res.Reasons)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Reason() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QueryReasonRequest
		shouldErr bool
		expReason types.Reason
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryReasonRequest(0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid reason id returns error",
			request:   types.NewQueryReasonRequest(1, 0),
			shouldErr: true,
		},
		{
			name:      "not found reason returns error",
			request:   types.NewQueryReasonRequest(1, 1),
			shouldErr: true,
		},
		{
			name: "valid request returns correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))
			},
			request:   types.NewQueryReasonRequest(1, 1),
			shouldErr: false,
			expReason: types.NewReason(
				1,
				1,
				"Spam",
				"This content is spam",
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

			res, err := suite.k.Reason(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expReason, res.Reason)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Params() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QueryParamsRequest
		shouldErr bool
		expParams types.Params
	}{
		{
			name: "default params are returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			request:   types.NewQueryParamsRequest(),
			shouldErr: false,
			expParams: types.DefaultParams(),
		},
		{
			name: "custom params return error",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams([]types.StandardReason{
					types.NewStandardReason(1, "Spam", "This content is spam"),
				}))
			},
			request:   types.NewQueryParamsRequest(),
			shouldErr: false,
			expParams: types.NewParams([]types.StandardReason{
				types.NewStandardReason(1, "Spam", "This content is spam"),
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.Params(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expParams, res.Params)
			}
		})
	}
}
