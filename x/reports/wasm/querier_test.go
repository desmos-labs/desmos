package wasm_test

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
	"github.com/desmos-labs/desmos/v3/x/reports/wasm"
)

func (suite *Testsuite) TestReportsWasmQuerier_QueryCustom() {
	profilesQuery := profilestypes.QueryProfileRequest{User: ""}
	profilesQueryBz, err := profilesQuery.Marshal()
	suite.NoError(err)
	wrongQueryBz, err := json.Marshal(profilesQueryBz)
	suite.NoError(err)

	testCases := []struct {
		name        string
		request     json.RawMessage
		store       func(ctx sdk.Context)
		shouldErr   bool
		expResponse []byte
	}{
		{
			name:        "wrong request type returns error",
			request:     wrongQueryBz,
			shouldErr:   true,
			expResponse: nil,
		},
		{
			name:    "reports request is parsed correctly",
			request: buildReportsQueryRequest(suite.cdc, types.NewQueryReportsRequest(1, nil, "", nil)),
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"test",
					types.NewPostTarget(1),
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC),
				))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryReportsResponse{
					Reports: []types.Report{types.NewReport(
						1,
						1,
						[]uint32{1},
						"test",
						types.NewPostTarget(1),
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2022, 6, 15, 0, 0, 0, 0, time.UTC),
					)},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1},
				},
			),
		},
		{
			name:    "reasons request is parsed correctly",
			request: buildReasonsQueryRequest(suite.cdc, types.NewQueryReasonsRequest(1, nil)),
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(1, 1, "test", "test"))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryReasonsResponse{Reasons: []types.Reason{types.NewReason(1, 1, "test", "test")},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1},
				},
			),
		},
	}

	querier := wasm.NewReportsWasmQuerier(suite.k, suite.cdc)

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}
			query, err := querier.QueryCustom(ctx, tc.request)
			if tc.shouldErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
			}
			suite.Equal(tc.expResponse, query)
		})
	}

}
