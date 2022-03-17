package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (suite *KeeperTestSuite) TestKeeper_CalculateCirculatingSupply() {
	testCases := []struct {
		name                      string
		coinDenom                 string
		expectedCirculatingSupply sdk.Coin
		store                     func(ctx sdk.Context)
	}{
		{
			name:                      "circulating supply calculated correctly",
			coinDenom:                 "udsm",
			expectedCirculatingSupply: sdk.NewCoin("udsm", sdk.NewInt(600_000)),
			store: func(ctx sdk.Context) {
				accAddr, err := sdk.AccAddressFromBech32("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				if err != nil {
					panic(err)
				}
				// Create a non-vesting account
				account := suite.ak.NewAccountWithAddress(ctx, accAddr)
				suite.ak.SetAccount(ctx, account)

				// Create a vesting account
				vestingAccount := vestingtypes.NewBaseVestingAccount(
					suite.CreateBaseAccount(),
					sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(200000))),
					12324125423,
				)
				suite.ak.SetAccount(ctx, vestingAccount)

				moduleAcc := authtypes.NewModuleAccount(
					suite.CreateBaseAccount(),
					banktypes.ModuleName,
					"minter",
				)

				suite.ak.SetModuleAccount(ctx, moduleAcc)

				// Mint coins from bank modules
				//err = suite.bk.MintCoins(
				//	ctx,
				//	moduleAcc.Name,
				//	sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(1_000_000))),
				//)
				//if err != nil {
				//	panic(err)
				//}

				suite.SetSupply(sdk.NewCoin("udsm", sdk.NewInt(1_000_000)))

				// Send out coins to accounts
				err = suite.bk.SendCoins(
					ctx,
					moduleAcc.GetAddress(),
					account.GetAddress(),
					sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(500000))),
				)
				if err != nil {
					panic(err)
				}

				err = suite.bk.SendCoinsFromModuleToAccount(
					ctx,
					banktypes.ModuleName,
					vestingAccount.GetAddress(),
					sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(200000))),
				)
				if err != nil {
					panic(err)
				}

				err = suite.dk.FundCommunityPool(
					ctx,
					sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(200000))),
					account.GetAddress(),
				)

				// Expected circulating supply
				// ECS = 1_000_000 - 200_000 - 200_000 = 600_000
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

			circulatingSupply := suite.k.CalculateCirculatingSupply(ctx, tc.coinDenom)
			suite.Require().Equal(tc.expectedCirculatingSupply, circulatingSupply)
		})
	}
}
