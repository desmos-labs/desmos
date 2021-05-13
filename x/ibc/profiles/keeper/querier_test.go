package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/ibc/profiles/keeper"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (suite *KeeperTestSuite) Test_queryLink() {
	tests := []struct {
		name      string
		path      []string
		storeLink types.Link
		expErr    error
	}{
		{
			name:      "Link doesnt exist (address given)",
			path:      []string{types.QueryLink, suite.testData.otherUser},
			storeLink: suite.testData.link,
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"Link with address %s doesn't exists", suite.testData.otherUser,
			),
		},
		{
			name:      "Link doesnt exist (blank path given)",
			path:      []string{types.QueryLink, ""},
			storeLink: suite.testData.link,
			expErr:    sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Address cannot be empty or blank"),
		},
		{
			name:      "Link returned correctly (address given)",
			path:      []string{types.QueryLink, suite.testData.user},
			storeLink: suite.testData.link,
			expErr:    nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			err := suite.k.StoreLink(suite.ctx, test.storeLink)
			suite.Require().Nil(err)

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Require().Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.legacyAminoCdc, &test.storeLink)
				suite.Require().NoError(err)
				suite.Require().Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.Require().Error(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
				suite.Require().Nil(result)
			}
		})
	}
}
