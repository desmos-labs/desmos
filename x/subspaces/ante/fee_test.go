package ante_test

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func (suite *AnteTestSuite) TestAnte_Ante() {
	feeAmount := sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100)))
	signer := sdk.MustAccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
	granter := types.GetTreasuryAddress(1)
	nonTreasuryGranter := sdk.MustAccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
	module := sdk.MustAccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
	nonSubspaceMsg := testdata.NewTestMsg(signer)
	subspaceID, otherSubspaceID := uint64(1), uint64(2)
	subspaceMsg := types.NewMsgAddUserToUserGroup(subspaceID, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
	otherSubspaceMsg := types.NewMsgAddUserToUserGroup(otherSubspaceID, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")

	testCases := []struct {
		name      string
		setup     func()
		buildTx   func() sdk.Tx
		shouldErr bool
		check     func(ctx sdk.Context)
		expEvents sdk.Events
	}{
		{
			name: "standard tx returns no error",
			setup: func() {
				suite.authDeductFeeDecorator.EXPECT().
					AnteHandle(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(suite.ctx, nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(nonSubspaceMsg)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: false,
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "standard tx that fails in auth fee deduction phase returns error",
			setup: func() {
				suite.authDeductFeeDecorator.EXPECT().
					AnteHandle(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(suite.ctx, fmt.Errorf("error"))
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(nonSubspaceMsg)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: true,
		},
		{
			name: "valid tx with different subspaces msgs returns no error",
			setup: func() {
				suite.authDeductFeeDecorator.EXPECT().
					AnteHandle(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(suite.ctx, nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg, otherSubspaceMsg)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: false,
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "valid tx with non treasury account granter returns no error",
			setup: func() {
				suite.authDeductFeeDecorator.EXPECT().
					AnteHandle(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(suite.ctx, nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeGranter(nonTreasuryGranter)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: false,
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "non existing granter account returns error",
			setup: func() {
				suite.ak.EXPECT().
					GetModuleAddress(authtypes.FeeCollectorName).
					Return(module)

				suite.sk.EXPECT().
					UseGrantedFees(gomock.Any(), subspaceID, signer, feeAmount, []sdk.Msg{subspaceMsg}).
					Return(true)

				suite.ak.EXPECT().
					GetAccount(gomock.Any(), granter).
					Return(nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeGranter(granter)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: true,
		},
		{
			name: "not set module fee collector returns error",
			setup: func() {
				suite.ak.EXPECT().
					GetModuleAddress(authtypes.FeeCollectorName).
					Return(nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeGranter(granter)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: true,
		},
		{
			name: "not enough funds returns error",
			setup: func() {
				suite.ak.EXPECT().
					GetModuleAddress(authtypes.FeeCollectorName).
					Return(module)

				suite.sk.EXPECT().
					UseGrantedFees(gomock.Any(), subspaceID, signer, feeAmount, []sdk.Msg{subspaceMsg}).
					Return(true)

				suite.ak.EXPECT().
					GetAccount(gomock.Any(), granter).
					Return(authtypes.NewBaseAccountWithAddress(granter))

				suite.bk.EXPECT().
					SendCoinsFromAccountToModule(gomock.Any(), granter, authtypes.FeeCollectorName, feeAmount).
					Return(sdkerrors.ErrInsufficientFunds)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeGranter(granter)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: true,
		},
		{
			name: "non-zero fees valid tx returns no error",
			setup: func() {
				suite.ak.EXPECT().
					GetModuleAddress(authtypes.FeeCollectorName).
					Return(module)

				suite.sk.EXPECT().
					UseGrantedFees(gomock.Any(), subspaceID, signer, feeAmount, []sdk.Msg{subspaceMsg}).
					Return(true)

				suite.ak.EXPECT().
					GetAccount(gomock.Any(), granter).
					Return(authtypes.NewBaseAccountWithAddress(granter))

				suite.bk.EXPECT().
					SendCoinsFromAccountToModule(gomock.Any(), granter, authtypes.FeeCollectorName, feeAmount).
					Return(nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeGranter(granter)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			check: func(ctx sdk.Context) {
				suite.Require().Equal(int64(10), ctx.Priority())
			},
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeTx,
					sdk.NewAttribute(sdk.AttributeKeyFee, feeAmount.String()),
					sdk.NewAttribute(sdk.AttributeKeyFeePayer, granter.String()),
				),
			},
		},
		{
			name: "zero fees valid tx returns no error",
			setup: func() {
				suite.ak.EXPECT().
					GetModuleAddress(authtypes.FeeCollectorName).
					Return(module)

				suite.sk.EXPECT().
					UseGrantedFees(gomock.Any(), subspaceID, signer, nil, []sdk.Msg{subspaceMsg}).
					Return(true)

				suite.ak.EXPECT().
					GetAccount(gomock.Any(), granter).
					Return(authtypes.NewBaseAccountWithAddress(granter))
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeGranter(granter)
				txBuilder.SetFeeAmount(nil)
				return txBuilder.GetTx()
			},
			check: func(ctx sdk.Context) {
				suite.Require().Equal(int64(10), ctx.Priority())
			},
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeTx,
					sdk.NewAttribute(sdk.AttributeKeyFee, ""),
					sdk.NewAttribute(sdk.AttributeKeyFeePayer, granter.String()),
				),
			},
		},
		{
			name: "valid tx without granter returns no error",
			setup: func() {
				suite.authDeductFeeDecorator.EXPECT().
					AnteHandle(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(suite.ctx, nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: false,
			expEvents: sdk.EmptyEvents(),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}

			tx := tc.buildTx()
			ctx, err := suite.ante.AnteHandle(ctx, tx, false, func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
				return ctx, nil
			})
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *AnteTestSuite) TestAnte_Ante_Simulate() {
	feeAmount := sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100)))
	subspaceMsg := types.NewMsgAddUserToUserGroup(1, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")

	testCases := []struct {
		name      string
		setup     func()
		buildTx   func() sdk.Tx
		simulate  bool
		shouldErr bool
	}{
		{
			name: "non simulate mode without gas returns error",
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeAmount(feeAmount)
				txBuilder.SetGasLimit(0)
				return txBuilder.GetTx()
			},
			shouldErr: true,
		},
		{
			name: "non simulate mode with gas returns no error",
			setup: func() {
				suite.authDeductFeeDecorator.EXPECT().
					AnteHandle(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(suite.ctx, nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeAmount(feeAmount)
				txBuilder.SetGasLimit(100)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
		{
			name: "simulate mode without gas returns no error",
			setup: func() {
				suite.authDeductFeeDecorator.EXPECT().
					AnteHandle(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(suite.ctx, nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeAmount(feeAmount)
				txBuilder.SetGasLimit(0)
				return txBuilder.GetTx()
			},
			simulate: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}

			ctx = ctx.WithBlockHeight(10)

			tx := tc.buildTx()
			ctx, err := suite.ante.AnteHandle(ctx, tx, tc.simulate, func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
				return ctx, nil
			})
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
