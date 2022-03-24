package ante_test

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	feestypes "github.com/desmos-labs/desmos/v3/x/fees/types"
)

func (suite *AnteTestSuite) TestAnteHandlerFees_MsgCreatePost() {
	suite.SetupTest(true)
	account := suite.createTestAccount()

	tests := []struct {
		name     string
		givenFee sdk.Coins
		params   feestypes.Params
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
			params: feestypes.NewParams([]feestypes.MinFee{
				feestypes.NewMinFee("desmos.profiles.v2.MsgSaveProfile", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
			msgs:     []sdk.Msg{},
			privs:    []cryptotypes.PrivKey{account.privKey},
			accNums:  []uint64{0},
			accSeqs:  []uint64{0},
			expError: true,
			simulate: false,
		},
		{
			name:     "Signer has not specified enough fees",
			givenFee: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 9999)),
			params: feestypes.NewParams([]feestypes.MinFee{
				feestypes.NewMinFee("desmos.profiles.v2.MsgSaveProfile", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
			msgs:     []sdk.Msg{},
			privs:    []cryptotypes.PrivKey{account.privKey},
			accNums:  []uint64{0},
			accSeqs:  []uint64{0},
			expError: true,
			simulate: false,
		},
		{
			name:     "Signer has specified enough fees",
			givenFee: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)),
			params: feestypes.NewParams([]feestypes.MinFee{
				feestypes.NewMinFee("desmos.profiles.v2.MsgSaveProfile", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
			msgs:     []sdk.Msg{},
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
