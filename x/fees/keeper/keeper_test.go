package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/desmos-labs/desmos/x/fees/types"
	posts "github.com/desmos-labs/desmos/x/posts/types"
	"time"
)

func (suite *KeeperTestSuite) TestKeeper_SetParams() {
	params := types.DefaultParams()
	suite.keeper.SetParams(suite.ctx, params)

	actualParams := suite.keeper.GetParams(suite.ctx)

	suite.Equal(params, actualParams)
}

func (suite *KeeperTestSuite) TestKeeper_GetParams() {
	params := types.DefaultParams()
	suite.keeper.SetParams(suite.ctx, params)

	actualParams := suite.keeper.GetParams(suite.ctx)

	suite.Equal(params, actualParams)

	tests := []struct {
		name      string
		params    *types.Params
		expParams *types.Params
	}{
		{
			name:      "Returning previously set params",
			params:    &params,
			expParams: &params,
		},
		{
			name:      "Returning nothing",
			params:    nil,
			expParams: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			if test.params != nil {
				suite.keeper.SetParams(suite.ctx, *test.params)
			}

			if test.expParams != nil {
				suite.Equal(*test.expParams, suite.keeper.GetParams(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_CheckFees() {
	// variables for later usage
	timeZone, _ := time.LoadLocation("UTC")
	pollData := posts.NewPollData(
		"poll?",
		time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
		posts.NewPollAnswers(
			posts.NewPollAnswer(posts.AnswerID(1), "Yes"),
			posts.NewPollAnswer(posts.AnswerID(2), "No"),
		),
		false,
		true,
	)
	attachments := posts.NewAttachments(posts.NewAttachment("https://uri.com", "text/plain", nil))
	id := posts.PostID("dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1")

	var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")

	tests := []struct {
		name      string
		params    types.Params
		givenFees authtypes.StdFee
		msgs      []sdk.Msg
		expError  error
	}{
		{
			name: "Not enough fees returns error",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("create_post", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
			givenFees: authtypes.NewStdFee(100000,
				sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 150)),
			),
			msgs: []sdk.Msg{
				posts.NewMsgCreatePost(
					"My new post",
					id,
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					nil,
					testOwner,
					attachments,
					&pollData,
				),
			},
			expError: sdkerrors.Wrap(sdkerrors.ErrInsufficientFee,
				"Expected at least 10000stake, got 150stake"),
		},
		{
			name: "Enough fees works properly",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("create_post", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
			givenFees: authtypes.NewStdFee(100000,
				sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)),
			),
			msgs: []sdk.Msg{
				posts.NewMsgCreatePost(
					"My new post",
					id,
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					nil,
					testOwner,
					attachments,
					&pollData,
				),
			},
			expError: nil,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.keeper.SetParams(suite.ctx, test.params)
			if err := suite.keeper.CheckFees(suite.ctx, test.givenFees, test.msgs); err != nil {
				suite.Equal(test.expError.Error(), err.Error())
			} else {
				suite.Nil(err)
			}
		})
	}
}
