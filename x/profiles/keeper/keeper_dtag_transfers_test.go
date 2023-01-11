package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/testutil/profilestesting"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveDTagTransferRequest() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		transferReq types.DTagTransferRequest
		shouldErr   bool
		check       func(ctx sdk.Context)
	}{
		{
			name: "already present request returns error",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			transferReq: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "request with a different DTag returns an error",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			transferReq: types.NewDTagTransferRequest(
				"dtag1",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "different sender request is saved properly",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn")))
			},
			transferReq: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				expected := []types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
					types.NewDTagTransferRequest(
						"dtag",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					),
				}
				suite.Require().Equal(expected, suite.k.GetDTagTransferRequests(ctx))
			},
		},
		{
			name: "different receiver request is saved correctly",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn")))
			},
			transferReq: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				expected := []types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
					types.NewDTagTransferRequest(
						"dtag",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					),
				}
				suite.Require().Equal(expected, suite.k.GetDTagTransferRequests(ctx))
			},
		},
		{
			name: "not already present request is saved correctly",
			store: func(ctx sdk.Context) {
				receiver := "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(receiver)))
			},
			transferReq: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				expected := []types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				}
				suite.Require().Equal(expected, suite.k.GetDTagTransferRequests(ctx))
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

			err := suite.k.SaveDTagTransferRequest(ctx, tc.transferReq)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetDTagTransferRequest() {
	testCases := []struct {
		name     string
		store    func(ctx sdk.Context)
		sender   string
		receiver string
		expReq   types.DTagTransferRequest
		expFound bool
	}{
		{
			name: "non-empty list is returned properly",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			sender:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			receiver: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expFound: true,
			expReq: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
		},
		{
			name:     "empty list is returned properly",
			sender:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			receiver: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expFound: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			req, found, err := suite.k.GetDTagTransferRequest(ctx, tc.sender, tc.receiver)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expFound, found)
			if found {
				suite.Require().Equal(tc.expReq, req)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetDTagTransferRequests() {
	testCases := []struct {
		name    string
		store   func(ctx sdk.Context)
		expReqs []types.DTagTransferRequest
	}{
		{
			name: "non-empty list is returned properly",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			expReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				),
			},
		},
		{
			name:    "empty list is returned properly",
			expReqs: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.Require().Equal(tc.expReqs, suite.k.GetDTagTransferRequests(ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteDTagTransferRequest() {
	testCases := []struct {
		name     string
		store    func(ctx sdk.Context)
		sender   string
		receiver string
		check    func(ctx sdk.Context)
	}{
		{
			name: "deleting non existent request works properly",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			sender:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			receiver: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		},
		{
			name: "existing request is removed properly",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				)
				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			sender:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			receiver: "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			check: func(ctx sdk.Context) {
				suite.Require().Empty(suite.k.GetDTagTransferRequests(ctx))
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

			suite.k.DeleteDTagTransferRequest(ctx, tc.sender, tc.receiver)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteAllUserIncomingDTagTransferRequests() {
	tests := []struct {
		name  string
		store func(ctx sdk.Context)
		user  string
		check func(ctx sdk.Context)
	}{
		{
			name: "DTag requests are deleted properly",
			store: func(ctx sdk.Context) {
				profile1 := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile1))

				profile2 := profilestesting.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile2))

				request := types.NewDTagTransferRequest(
					profile1.DTag,
					profile2.GetAddress().String(),
					profile1.GetAddress().String(),
				)
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			user: "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			check: func(ctx sdk.Context) {
				user := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"

				var iterations = 0
				suite.k.IterateUserIncomingDTagTransferRequests(ctx, user, func(_ types.DTagTransferRequest) (stop bool) {
					iterations += 1
					return false
				})
				suite.Require().Zero(iterations)
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if test.store != nil {
				test.store(ctx)
			}

			suite.k.DeleteAllUserIncomingDTagTransferRequests(ctx, test.user)
			if test.check != nil {
				test.check(ctx)
			}
		})
	}
}
