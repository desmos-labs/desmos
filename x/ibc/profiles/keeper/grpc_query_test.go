package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
)

func (suite *KeeperTestSuite) TestQueryLink() {
	tests := []struct {
		name        string
		storedLinks []types.Link
		request     *types.QueryLinkRequest
		expError    bool
		expResponse *types.QueryLinkResponse
	}{
		{
			name: "Return link properly",
			storedLinks: []types.Link{
				types.NewLink("cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn", "cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn"),
			},
			request:     &types.QueryLinkRequest{SourceAddress: "cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn"},
			expError:    false,
			expResponse: &types.QueryLinkResponse{Link: types.NewLink("cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn", "cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn")},
		},
		{
			name:        "Empty links returns error",
			storedLinks: []types.Link{},
			request:     &types.QueryLinkRequest{SourceAddress: "cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn"},
			expError:    true,
		},
		{
			name:        "Empty source address",
			storedLinks: []types.Link{},
			request:     &types.QueryLinkRequest{SourceAddress: ""},
			expError:    true,
		},
		{
			name:        "Invalid source address",
			storedLinks: []types.Link{},
			request:     &types.QueryLinkRequest{SourceAddress: "hahaha"},
			expError:    true,
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

			ctx := sdk.WrapSDKContext(suite.ctx)
			response, err := suite.k.Link(ctx, test.request)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expResponse, response)
			}
		})
	}
}
