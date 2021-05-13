package keeper_test

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_StoreLink() {
	tests := []struct {
		name        string
		link        types.Link
		storedLinks []types.Link
		expError    error
	}{
		{
			name:        "Non existent Link saved correctly",
			link:        suite.testData.link,
			storedLinks: nil,
			expError:    nil,
		},
		{
			name:        "Link already exists returns error",
			link:        suite.testData.link,
			storedLinks: []types.Link{suite.testData.link},
			expError:    sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "source address already exists"),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, link := range test.storedLinks {
				err := suite.k.StoreLink(suite.ctx, link)
				suite.Require().NoError(err)
			}

			err := suite.k.StoreLink(suite.ctx, test.link)
			suite.RequireErrorsEqual(test.expError, err)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetLink() {
	tests := []struct {
		name        string
		storedLinks []types.Link
		address     string
		expFound    bool
		expLink     *types.Link
	}{
		{
			name: "Link founded",
			storedLinks: []types.Link{
				suite.testData.link,
			},
			address:  suite.testData.user,
			expFound: true,
			expLink:  &suite.testData.link,
		},
		{
			name:        "Link not found",
			storedLinks: []types.Link{},
			address:     suite.testData.user,
			expFound:    false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, link := range test.storedLinks {
				err := suite.k.StoreLink(suite.ctx, link)
				suite.Require().NoError(err)
			}

			res, found := suite.k.GetLink(suite.ctx, test.address)
			suite.Require().Equal(test.expFound, found)

			if found {
				suite.Require().True(res.Equal(test.expLink))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAllLinks() {
	tests := []struct {
		name  string
		links []types.Link
	}{
		{
			name:  "Non empty Links list returned",
			links: []types.Link{suite.testData.link},
		},
		{
			name:  "Link not found",
			links: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if len(test.links) != 0 {
				store := suite.ctx.KVStore(suite.storeKey)
				key := types.LinkStoreKey(test.links[0].GetSourceAddress())
				store.Set(key, suite.cdc.MustMarshalBinaryBare(&test.links[0]))
			}

			res := suite.k.GetAllLinks(suite.ctx)
			suite.Require().Equal(test.links, res)
		})
	}
}
