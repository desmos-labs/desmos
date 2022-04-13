package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v3/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func (suite *KeeperTestSuite) TestQueryServer_Profile() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryProfileRequest
		shouldErr   bool
		expResponse *types.QueryProfileResponse
	}{
		{
			name:      "empty user returns error",
			req:       types.NewQueryProfileRequest(""),
			shouldErr: true,
		},
		{
			name:      "non existing DTag returns error",
			req:       types.NewQueryProfileRequest("invalid-dtag"),
			shouldErr: true,
		},
		{
			name:        "profile not found",
			req:         types.NewQueryProfileRequest("cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa"),
			shouldErr:   false,
			expResponse: &types.QueryProfileResponse{Profile: nil},
		},
		{
			name: "found profile using DTag",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				err := suite.k.SaveProfile(ctx, profile)
				suite.Require().NoError(err)
			},
			req:       types.NewQueryProfileRequest("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x-dtag"),
			shouldErr: false,
			expResponse: &types.QueryProfileResponse{
				Profile: testutil.NewAny(testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
			},
		},
		{
			name: "found profile using address",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				err := suite.k.SaveProfile(ctx, profile)
				suite.Require().NoError(err)
			},
			req:       types.NewQueryProfileRequest("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"),
			shouldErr: false,
			expResponse: &types.QueryProfileResponse{
				Profile: testutil.NewAny(testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
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

			res, err := suite.k.Profile(sdk.WrapSDKContext(ctx), tc.req)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(tc.expResponse, res)
				if tc.expResponse.Profile != nil {
					suite.Require().Equal(tc.expResponse, res)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_IncomingDTagTransferRequests() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryIncomingDTagTransferRequestsRequest
		shouldErr   bool
		expRequests []types.DTagTransferRequest
	}{
		{
			name: "valid request without pagination",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

				request = types.NewDTagTransferRequest(
					"dtag",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			req: types.NewQueryIncomingDTagTransferRequestsRequest(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				nil,
			),
			shouldErr: false,
			expRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				),
			},
		},
		{
			name: "valid request with pagination",
			store: func(ctx sdk.Context) {
				receiver := "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr(receiver)))

				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					receiver,
				)
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

				request = types.NewDTagTransferRequest(
					"dtag",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					receiver,
				)
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			req: types.NewQueryIncomingDTagTransferRequestsRequest(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				&query.PageRequest{Limit: 1},
			),
			shouldErr: false,
			expRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				),
			},
		},
		{
			name: "valid request without user",
			store: func(ctx sdk.Context) {
				receiver1 := "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr(receiver1)))

				receiver2 := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr(receiver2)))

				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					receiver1,
				)
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

				request = types.NewDTagTransferRequest(
					"dtag",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					receiver2,
				)
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			req:       types.NewQueryIncomingDTagTransferRequestsRequest("", nil),
			shouldErr: false,
			expRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				),
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
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

			res, err := suite.k.IncomingDTagTransferRequests(sdk.WrapSDKContext(ctx), tc.req)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(tc.expRequests, res.Requests)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_ChainLinks() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryChainLinksRequest
		shouldErr bool
		expLinks  []types.ChainLink
	}{
		{
			name: "valid request without pagination",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				link := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)

				link = types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
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
			req: types.NewQueryChainLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"", "", nil,
			),
			shouldErr: false,
			expLinks: []types.ChainLink{
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
		{
			name: "valid request with chain name",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				link := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)

				link = types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)

				link = types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("likecoin"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)
			},
			req: types.NewQueryChainLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos",
				"",
				nil,
			),
			shouldErr: false,
			expLinks: []types.ChainLink{
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
		{
			name: "valid request with target",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				link := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)

				link = types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
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
			req: types.NewQueryChainLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos",
				"cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw",
				nil,
			),
			shouldErr: false,
			expLinks: []types.ChainLink{
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
		{
			name: "valid request with pagination",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				link := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)

				link = types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
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
			req: types.NewQueryChainLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"", "",
				&query.PageRequest{Limit: 1},
			),
			shouldErr: false,
			expLinks: []types.ChainLink{
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
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

			res, err := suite.k.ChainLinks(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(tc.expLinks, res.Links)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_ChainLinkOwners() {
	firstAccount := testutil.GetChainLinkAccount("cosmos", "cosmos")
	secondAccount := testutil.GetChainLinkAccount("likecoin", "cosmos")

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QueryChainLinkOwnersRequest
		shouldErr bool
		expOwners []types.QueryChainLinkOwnersResponse_ChainLinkOwnerDetails
	}{
		{
			name: "query without any data returns everything",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
				suite.Require().NoError(suite.k.SaveChainLink(ctx, firstAccount.GetBech32ChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
				)))
				suite.Require().NoError(suite.k.SaveChainLink(ctx, secondAccount.GetBech32ChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
				)))
			},
			request:   types.NewQueryChainLinkOwnersRequest("", "", nil),
			shouldErr: false,
			expOwners: []types.QueryChainLinkOwnersResponse_ChainLinkOwnerDetails{
				{
					User:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					ChainName: firstAccount.ChainName(),
					Target:    firstAccount.Bech32Address().GetValue(),
				},
				{
					User:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					ChainName: secondAccount.ChainName(),
					Target:    secondAccount.Bech32Address().GetValue(),
				},
			},
		},
		{
			name: "query with chain name returns the correct data",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
				suite.Require().NoError(suite.k.SaveChainLink(ctx, firstAccount.GetBech32ChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
				)))
				suite.Require().NoError(suite.k.SaveChainLink(ctx, secondAccount.GetBech32ChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
				)))
			},
			request:   types.NewQueryChainLinkOwnersRequest("cosmos", "", nil),
			shouldErr: false,
			expOwners: []types.QueryChainLinkOwnersResponse_ChainLinkOwnerDetails{
				{
					Target:    firstAccount.Bech32Address().GetValue(),
					ChainName: firstAccount.ChainName(),
					User:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
		},
		{
			name: "query with chain name and target returns the correct data",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
				suite.Require().NoError(suite.k.SaveChainLink(ctx, firstAccount.GetBech32ChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
				)))
				suite.Require().NoError(suite.k.SaveChainLink(ctx, secondAccount.GetBech32ChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
				)))
			},
			request:   types.NewQueryChainLinkOwnersRequest("cosmos", firstAccount.Bech32Address().GetValue(), nil),
			shouldErr: false,
			expOwners: []types.QueryChainLinkOwnersResponse_ChainLinkOwnerDetails{
				{
					User:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					ChainName: firstAccount.ChainName(),
					Target:    firstAccount.Bech32Address().GetValue(),
				},
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

			res, err := suite.k.ChainLinkOwners(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expOwners, res.Owners)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_ApplicationLinks() {
	testCases := []struct {
		name                string
		store               func(ctx sdk.Context)
		req                 *types.QueryApplicationLinksRequest
		shouldErr           bool
		expApplicationLinks []types.ApplicationLink
	}{
		{
			name:                "empty requests return empty result",
			req:                 types.NewQueryApplicationLinksRequest("user", "", "", nil),
			shouldErr:           false,
			expApplicationLinks: nil,
		},
		{
			name: "valid request with application",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
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
				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("github", "githubuser"),
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
				))
			},
			req: types.NewQueryApplicationLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"twitter",
				"",
				nil,
			),
			shouldErr: false,
			expApplicationLinks: []types.ApplicationLink{
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
		},
		{
			name: "valid request with application and username",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
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
				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "githubuser"),
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
				))
			},
			req: types.NewQueryApplicationLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"twitter",
				"twitteruser",
				nil,
			),
			shouldErr: false,
			expApplicationLinks: []types.ApplicationLink{
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
		},
		{
			name: "valid request without pagination",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
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
				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("github", "githubuser"),
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
				))
			},
			req: types.NewQueryApplicationLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"", "", nil,
			),
			shouldErr: false,
			expApplicationLinks: []types.ApplicationLink{
				types.NewApplicationLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewData("github", "githubuser"),
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
		},
		{
			name: "valid paginated request returns proper response",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
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
				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("github", "githubuser"),
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
				))
			},
			req: types.NewQueryApplicationLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
				"",
				&query.PageRequest{Limit: 1, Offset: 1, CountTotal: true},
			),
			shouldErr: false,
			expApplicationLinks: []types.ApplicationLink{
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
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.ApplicationLinks(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expApplicationLinks, res.Links)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_ApplicationLinkByClientID() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryApplicationLinkByClientIDRequest
		shouldErr bool
		expLink   types.ApplicationLink
	}{
		{
			name:      "not found link returns error",
			req:       types.NewQueryApplicationLinkByClientIDRequest("client_id"),
			shouldErr: true,
		},
		{
			name: "valid request returns proper response",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
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
			req:       types.NewQueryApplicationLinkByClientIDRequest("client_id"),
			shouldErr: false,
			expLink: types.NewApplicationLink(
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
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.ApplicationLinkByClientID(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expLink, res.Link)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_ApplicationLinkOwners() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryApplicationLinkOwnersRequest
		shouldErr bool
		expOwners []types.QueryApplicationLinkOwnersResponse_ApplicationLinkOwnerDetails
	}{
		{
			name: "query without any data returns everything",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitter_user"),
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

				profile = testutil.ProfileFromAddr("cosmos1ngzeux3j0vfkps0779y0c8pnrmszlg0hekp5um")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1ngzeux3j0vfkps0779y0c8pnrmszlg0hekp5um",
						types.NewData("github", "github_user"),
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
			req:       types.NewQueryApplicationLinkOwnersRequest("", "", nil),
			shouldErr: false,
			expOwners: []types.QueryApplicationLinkOwnersResponse_ApplicationLinkOwnerDetails{
				{
					User:        "cosmos1ngzeux3j0vfkps0779y0c8pnrmszlg0hekp5um",
					Application: "github",
					Username:    "github_user",
				},
				{
					User:        "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					Application: "twitter",
					Username:    "twitter_user",
				},
			},
		},
		{
			name: "query with application returns the correct data",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "first_user"),
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

				profile = testutil.ProfileFromAddr("cosmos1ngzeux3j0vfkps0779y0c8pnrmszlg0hekp5um")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1ngzeux3j0vfkps0779y0c8pnrmszlg0hekp5um",
						types.NewData("twitter", "second_user"),
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

				profile = testutil.ProfileFromAddr("cosmos1pxsak5c7ke5tz3d8alawuzu3cayr9s65ce7njr")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1pxsak5c7ke5tz3d8alawuzu3cayr9s65ce7njr",
						types.NewData("github", "second_user"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData(
								"github",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					)),
				)
			},
			req:       types.NewQueryApplicationLinkOwnersRequest("twitter", "", nil),
			shouldErr: false,
			expOwners: []types.QueryApplicationLinkOwnersResponse_ApplicationLinkOwnerDetails{
				{
					User:        "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					Application: "twitter",
					Username:    "first_user",
				},
				{
					User:        "cosmos1ngzeux3j0vfkps0779y0c8pnrmszlg0hekp5um",
					Application: "twitter",
					Username:    "second_user",
				},
			},
		},
		{
			name: "query with application and username returns the correct data",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "user"),
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

				profile = testutil.ProfileFromAddr("cosmos1ngzeux3j0vfkps0779y0c8pnrmszlg0hekp5um")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1ngzeux3j0vfkps0779y0c8pnrmszlg0hekp5um",
						types.NewData("twitter", "user"),
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

				profile = testutil.ProfileFromAddr("cosmos1mrmyggajlv0k3mlrhergjjnt75srn5y5u5a83x")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1mrmyggajlv0k3mlrhergjjnt75srn5y5u5a83x",
						types.NewData("twitter", "second_user"),
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
			req:       types.NewQueryApplicationLinkOwnersRequest("twitter", "user", nil),
			shouldErr: false,
			expOwners: []types.QueryApplicationLinkOwnersResponse_ApplicationLinkOwnerDetails{
				{
					User:        "cosmos1ngzeux3j0vfkps0779y0c8pnrmszlg0hekp5um",
					Application: "twitter",
					Username:    "user",
				},
				{
					User:        "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					Application: "twitter",
					Username:    "user",
				},
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

			res, err := suite.k.ApplicationLinkOwners(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expOwners, res.Owners)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Params() {
	ctx, _ := suite.ctx.CacheContext()

	suite.k.SetParams(ctx, types.DefaultParams())

	res, err := suite.k.Params(sdk.WrapSDKContext(ctx), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().NotNil(res)

	suite.Require().Equal(types.DefaultParams(), res.Params)
}
