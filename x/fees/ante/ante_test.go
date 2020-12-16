package ante_test

import (
	"time"

	feesTypes "github.com/desmos-labs/desmos/x/fees/types"
	"github.com/desmos-labs/desmos/x/posts/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *AnteTestSuite) TestAnteHandlerFees_MsgCreatePost() {
	suite.SetupTest(true)
	account := suite.createTestAccount()

	tests := []struct {
		name     string
		givenFee sdk.Coins
		params   feesTypes.Params
		msgs     []sdk.Msg
		privs    []cryptotypes.PrivKey
		accNums  []uint64
		accSeqs  []uint64
		expError bool
		simulate bool
	}{
		{
			name:     "Signer has not specified the fees",
			givenFee: sdk.NewCoins(),
			params: feesTypes.NewParams([]feesTypes.MinFee{
				feesTypes.NewMinFee("create_post", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
			msgs: []sdk.Msg{
				types.NewMsgCreatePost(
					"My new post",
					"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					nil,
					account.acc.GetAddress().String(),
					types.NewAttachments(types.NewAttachment("https://uri.com", "text/plain", nil)),
					types.NewPollData(
						"poll?",
						time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
						types.NewPollAnswers(
							types.NewPollAnswer("1", "Yes"),
							types.NewPollAnswer("2", "No"),
						),
						false,
						true,
					),
				),
			},
			privs:    []cryptotypes.PrivKey{account.privKey},
			accNums:  []uint64{0},
			accSeqs:  []uint64{0},
			expError: true,
			simulate: false,
		},
		{
			name:     "Signer has not specified enough fees",
			givenFee: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 9999)),
			params: feesTypes.NewParams([]feesTypes.MinFee{
				feesTypes.NewMinFee("create_post", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
			msgs: []sdk.Msg{
				types.NewMsgCreatePost(
					"My new post",
					"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					nil,
					account.acc.GetAddress().String(),
					types.NewAttachments(types.NewAttachment("https://uri.com", "text/plain", nil)),
					types.NewPollData(
						"poll?",
						time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
						types.NewPollAnswers(
							types.NewPollAnswer("1", "Yes"),
							types.NewPollAnswer("2", "No"),
						),
						false,
						true,
					),
				),
			},
			privs:    []cryptotypes.PrivKey{account.privKey},
			accNums:  []uint64{0},
			accSeqs:  []uint64{0},
			expError: true,
			simulate: false,
		},
		{
			name:     "Signer has specified enough fees",
			givenFee: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)),
			params: feesTypes.NewParams([]feesTypes.MinFee{
				feesTypes.NewMinFee("create_post", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
			msgs: []sdk.Msg{
				types.NewMsgCreatePost(
					"My new post",
					"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					nil,
					account.acc.GetAddress().String(),
					types.NewAttachments(types.NewAttachment("https://uri.com", "text/plain", nil)),
					types.NewPollData(
						"poll?",
						time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
						types.NewPollAnswers(
							types.NewPollAnswer("1", "Yes"),
							types.NewPollAnswer("2", "No"),
						),
						false,
						true,
					),
				),
			},
			privs:    []cryptotypes.PrivKey{account.privKey},
			accNums:  []uint64{0},
			accSeqs:  []uint64{0},
			expError: false,
			simulate: true,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.app.FeesKeeper.SetParams(suite.ctx, test.params)

			suite.Require().NoError(suite.txBuilder.SetMsgs(test.msgs...))
			suite.txBuilder.SetFeeAmount(test.givenFee)
			suite.txBuilder.SetGasLimit(200000)

			tx, txErr := suite.CreateTestTx(test.privs, test.accNums, test.accSeqs, suite.ctx.ChainID())
			newCtx, anteErr := suite.anteHandler(suite.ctx, tx, test.simulate)

			if !test.expError {
				suite.Require().NoError(txErr)
				suite.Require().NoError(anteErr)
				suite.Require().NotNil(newCtx)

				suite.ctx = newCtx
			} else {
				switch {
				case txErr != nil:
					suite.Require().Error(txErr)
				case anteErr != nil:
					suite.Require().Error(anteErr)
				default:
					suite.Fail("expected one of txErr, anteErr to be an error")
				}
			}
		})
	}
}
