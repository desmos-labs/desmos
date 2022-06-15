package wasm_test

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/desmos-labs/desmos/v3/testutil/profilestesting"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"

	"github.com/desmos-labs/desmos/v3/x/profiles/wasm"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *TestSuite) TestProfilesWasmQuerier_QueryCustom() {
	subspacesQuery := subspacestypes.QuerySubspacesRequest{Pagination: nil}
	subspacesQueryBz, err := subspacesQuery.Marshal()
	suite.NoError(err)
	wrongQueryBz, err := json.Marshal(subspacesQueryBz)
	suite.NoError(err)

	chainAccount := profilestesting.GetChainLinkAccount("cosmos", "cosmos")

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
			name: "profiles request is parsed correctly",
			request: buildProfileQueryRequest(
				suite.cdc,
				types.NewQueryProfileRequest("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"),
			),
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				err := suite.k.SaveProfile(ctx, profile)
				suite.Require().NoError(err)
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryProfileResponse{
					Profile: profilestesting.NewAny(profilestesting.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"))},
			),
		},
		{
			name: "incoming dtag transfer requests request is parsed correctly",
			request: buildIncomingDtagTransferQueryRequest(
				suite.cdc,
				types.NewQueryIncomingDTagTransferRequestsRequest("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x", nil),
			),
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryIncomingDTagTransferRequestsResponse{
					Requests: []types.DTagTransferRequest{
						types.NewDTagTransferRequest(
							"dtag",
							"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
							"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
						),
					},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1}},
			),
		},
		{
			name: "chain links request request is parsed correctly",
			request: buildChainLinksQueryRequest(suite.cdc, types.NewQueryChainLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "", "", nil)),
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				link := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						profilestesting.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						profilestesting.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryChainLinksResponse{
					Links: []types.ChainLink{
						types.NewChainLink(
							"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
							types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
							types.NewProof(
								profilestesting.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
								profilestesting.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
								"74657874",
							),
							types.NewChainConfig("cosmos"),
							time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
						),
					},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1}},
			),
		},
		{
			name: "chain link owners request request is parsed correctly",
			request: buildChainLinkOwnersQueryRequest(suite.cdc, types.NewQueryChainLinkOwnersRequest(
				"cosmos", chainAccount.Bech32Address().GetValue(), nil)),
			store: func(ctx sdk.Context) {
				suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"))
				suite.Require().NoError(suite.k.SaveChainLink(ctx, chainAccount.GetBech32ChainLink("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", time.Now())))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryChainLinkOwnersResponse{
					Owners: []types.QueryChainLinkOwnersResponse_ChainLinkOwnerDetails{{
						User:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						ChainName: "cosmos",
						Target:    chainAccount.Bech32Address().GetValue(),
					}},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1}},
			),
		},
		{
			name: "app links request is parsed properly",
			request: buildAppLinksQueryRequest(suite.cdc, types.NewQueryApplicationLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"twitter",
				"",
				nil,
			)),
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					)),
				)
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(&types.QueryApplicationLinksResponse{
				Links: []types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				Pagination: &query.PageResponse{NextKey: nil, Total: 1}},
			),
		},
		{
			name:    "application link by client ID request is parsed correctly",
			request: buildApplicationLinkByClientIDQueryRequest(suite.cdc, types.NewQueryApplicationLinkByClientIDRequest("client_id")),
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					)),
				)
			},
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryApplicationLinkByClientIDResponse{
					Link: types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					)},
			),
		},
		{
			name: "app link owners request is parsed properly",
			request: buildApplicationLinkOwnersQueryRequest(suite.cdc, types.NewQueryApplicationLinkOwnersRequest(
				"twitter",
				"twitteruser",
				nil,
			)),
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					)),
				)
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(&types.QueryApplicationLinkOwnersResponse{
				Owners: []types.QueryApplicationLinkOwnersResponse_ApplicationLinkOwnerDetails{{
					User:        "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					Application: "twitter",
					Username:    "twitteruser",
				}},
				Pagination: &query.PageResponse{NextKey: nil, Total: 1}},
			),
		},
	}

	querier := wasm.NewProfilesWasmQuerier(suite.k, suite.cdc)

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
