package keeper_test

import "github.com/desmos-labs/desmos/x/profiles/types"

func (suite *KeeperTestSuite) TestKeeper_SaveDTagTransferRequest() {
	tests := []struct {
		name                  string
		storedTransferReqs    []types.DTagTransferRequest
		transferReq           types.DTagTransferRequest
		shouldErr             bool
		expStoredTransferReqs []types.DTagTransferRequest
	}{
		{
			name: "Already present request returns error",
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			shouldErr:   true,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name: "Different sender request is saved properly",
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.otherUser, suite.testData.user),
			shouldErr:   false,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
				types.NewDTagTransferRequest("dtag", suite.testData.otherUser, suite.testData.user),
			},
		},
		{
			name: "Different receiver request is saved correctly",
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			shouldErr:   false,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name: "Different DTag request returns an error",
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag1", suite.testData.user, suite.testData.otherUser),
			shouldErr:   true,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name:               "Not already present request is saved correctly",
			storedTransferReqs: nil,
			transferReq:        types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			shouldErr:          false,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedTransferReqs {
				err := suite.k.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			err := suite.k.SaveDTagTransferRequest(suite.ctx, test.transferReq)

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			stored := suite.k.GetDTagTransferRequests(suite.ctx)
			suite.Require().Len(stored, len(test.expStoredTransferReqs))
			for _, req := range stored {
				suite.Require().Contains(test.expStoredTransferReqs, req)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserDTagTransferRequests() {
	tests := []struct {
		name       string
		storedReqs []types.DTagTransferRequest
		expReqs    []types.DTagTransferRequest
	}{
		{
			name: "returns a non-empty array of dTag requests",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			expReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name:       "returns an empty array of dTag requests",
			storedReqs: nil,
			expReqs:    nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			if test.storedReqs != nil {
				reqs := types.NewDTagTransferRequests(test.storedReqs)
				store.Set(
					types.DTagTransferRequestStoreKey(suite.testData.user),
					suite.cdc.MustMarshalBinaryBare(&reqs),
				)
			}

			suite.Require().Equal(test.expReqs, suite.k.IterateUserIncomingDTagTransferRequests(suite.ctx, suite.testData.user))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetDTagTransferRequests() {
	tests := []struct {
		name       string
		storedReqs []types.DTagTransferRequest
		expReqs    []types.DTagTransferRequest
	}{
		{
			name: "returns a non-empty array of dTag requests",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.otherUser),
			},
			expReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name:       "returns an empty array of dTag requests",
			storedReqs: nil,
			expReqs:    nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			if test.storedReqs != nil {
				reqs := types.NewDTagTransferRequests(test.storedReqs)
				store.Set(
					types.DTagTransferRequestStoreKey(suite.testData.user),
					suite.cdc.MustMarshalBinaryBare(&reqs),
				)
			}

			suite.Require().Equal(test.expReqs, suite.k.GetDTagTransferRequests(suite.ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteAllDTagTransferRequests() {
	tests := []struct {
		name       string
		storedReqs []types.DTagTransferRequest
		expReqs    []types.DTagTransferRequest
	}{
		{
			name: "returns a non-empty array of dTag requests",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.otherUser),
			},
			expReqs: nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			if test.storedReqs != nil {
				reqs := types.NewDTagTransferRequests(test.storedReqs)
				store.Set(
					types.DTagTransferRequestStoreKey(suite.testData.user),
					suite.cdc.MustMarshalBinaryBare(&reqs),
				)
			}

			suite.k.DeleteAllDTagTransferRequests(suite.ctx, suite.testData.user)
			suite.Require().Equal(test.expReqs, suite.k.GetDTagTransferRequests(suite.ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteDTagTransferRequest() {
	tests := []struct {
		name            string
		storedReqs      []types.DTagTransferRequest
		sender          string
		receiver        string
		shouldErr       bool
		storedReqsAfter []types.DTagTransferRequest
	}{
		{
			name:       "Empty requests array returns error",
			storedReqs: nil,
			sender:     suite.testData.user,
			receiver:   suite.testData.otherUser,
			shouldErr:  true,
		},
		{
			name: "Deleting non existent request returns an error and doesn't change the store",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
			sender:    suite.testData.user,
			receiver:  suite.testData.otherUser,
			shouldErr: true,
			storedReqsAfter: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
		},
		{
			name: "Existing request gets removed properly and leaves an array",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
			sender:   suite.testData.user,
			receiver: suite.testData.otherUser,
			storedReqsAfter: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
		},
		{
			name: "Existing request gets removed properly and doesn't leave anything",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			sender:          suite.testData.user,
			receiver:        suite.testData.otherUser,
			storedReqsAfter: nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedReqs {
				err := suite.k.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			err := suite.k.DeleteDTagTransferRequest(suite.ctx, test.sender, test.receiver)

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			reqs := suite.k.GetDTagTransferRequests(suite.ctx)
			suite.Require().Equal(test.storedReqsAfter, reqs)
		})
	}
}
