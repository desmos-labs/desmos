package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/relationships/keeper"
	"github.com/desmos-labs/desmos/x/relationships/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (suite *KeeperTestSuite) Test_queryUserRelationships() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	expRelationships := types.NewRelationshipResponse([]sdk.AccAddress{addr1, addr2})

	tests := []struct {
		name          string
		path          []string
		relationships []sdk.AccAddress
		expResult     *types.RelationshipsResponse
		expErr        error
	}{
		{
			name:          "Invalid bech32 address returns error",
			path:          []string{types.QueryRelationships, "invalidAddress"},
			relationships: nil,
			expResult:     nil,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid bech32 address: invalidAddress"),
		},
		{
			name:          "Relationships returned correctly",
			path:          []string{types.QueryRelationships, suite.testData.user.String()},
			relationships: []sdk.AccAddress{addr1, addr2},
			expResult:     &expRelationships,
			expErr:        nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			for _, rel := range test.relationships {
				_ = suite.keeper.StoreRelationship(suite.ctx, suite.testData.user, rel)
			}

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expResult != nil {
				suite.Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
				suite.NoError(err)
				suite.Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
				suite.Nil(result)
			}
		})
	}

}
