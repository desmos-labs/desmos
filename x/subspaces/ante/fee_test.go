package ante_test

import (
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func (suite *AnteTestSuite) TestAnte_Ante() {
	feeAmount := sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100)))
	signer := sdk.MustAccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
	granter := sdk.MustAccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
	module := sdk.MustAccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
	nonSubspaceMsg := testdata.NewTestMsg(signer)
	subspaceID, otherSubspaceID := uint64(1), uint64(2)
	subspaceMsg := types.NewMsgAddUserToUserGroup(subspaceID, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
	otherSubspaceMsg := types.NewMsgAddUserToUserGroup(otherSubspaceID, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")

	testCases := []struct {
		name      string
		malleate  func(ctx sdk.Context)
		buildTx   func() sdk.Tx
		shouldErr bool
		expEvents sdk.Events
	}{

		{
			name: "not set module fee collector returns error",
			malleate: func(ctx sdk.Context) {
				suite.ak.EXPECT().GetModuleAddress(authtypes.FeeCollectorName).Return(nil)
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
			name: "non existing granter account returns error",
			malleate: func(ctx sdk.Context) {
				suite.sk.EXPECT().UseGrantedFees(ctx, subspaceID, granter, signer, feeAmount, []sdk.Msg{subspaceMsg}).Return(true)
				suite.ak.EXPECT().GetModuleAddress(authtypes.FeeCollectorName).Return(module)
				suite.ak.EXPECT().GetAccount(ctx, granter).Return(nil)
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
			malleate: func(ctx sdk.Context) {
				suite.sk.EXPECT().UseGrantedFees(ctx, subspaceID, granter, signer, feeAmount, []sdk.Msg{subspaceMsg}).Return(true)
				suite.ak.EXPECT().GetModuleAddress(authtypes.FeeCollectorName).Return(module)
				suite.ak.EXPECT().GetAccount(ctx, granter).Return(authtypes.NewBaseAccountWithAddress(granter))
				suite.bk.EXPECT().SendCoinsFromAccountToModule(ctx, granter, authtypes.FeeCollectorName, feeAmount).Return(sdkerrors.ErrInsufficientFunds)
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
			name: "non-zero fees valid tx with returns no error",
			malleate: func(ctx sdk.Context) {
				suite.sk.EXPECT().UseGrantedFees(ctx, subspaceID, granter, signer, feeAmount, []sdk.Msg{subspaceMsg}).Return(true)
				suite.ak.EXPECT().GetModuleAddress(authtypes.FeeCollectorName).Return(module)
				suite.ak.EXPECT().GetAccount(ctx, granter).Return(authtypes.NewBaseAccountWithAddress(granter))
				suite.bk.EXPECT().SendCoinsFromAccountToModule(ctx, granter, authtypes.FeeCollectorName, feeAmount).Return(nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeGranter(granter)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
		{
			name: "0 fees subspace tx with the granter returns no error",
			malleate: func(ctx sdk.Context) {
				suite.sk.EXPECT().UseGrantedFees(ctx, subspaceID, granter, signer, nil, []sdk.Msg{subspaceMsg}).Return(true)
				suite.ak.EXPECT().GetModuleAddress(authtypes.FeeCollectorName).Return(module)
				suite.ak.EXPECT().GetAccount(ctx, granter).Return(authtypes.NewBaseAccountWithAddress(granter))
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeGranter(granter)
				txBuilder.SetFeeAmount(nil)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
		{
			name: "subspace tx without the granter using auth.DeductFeeDecorator returns no error",
			malleate: func(ctx sdk.Context) {
				suite.ak.EXPECT().GetModuleAddress(authtypes.FeeCollectorName).Return(module)
				suite.ak.EXPECT().GetAccount(ctx, signer).Return(authtypes.NewBaseAccountWithAddress(signer))
				suite.bk.EXPECT().SendCoinsFromAccountToModule(ctx, signer, authtypes.FeeCollectorName, feeAmount).Return(nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
		{
			name: "valid tx with valid feegrant allowance but no subspace allowance returns no error",
			malleate: func(ctx sdk.Context) {
				suite.sk.EXPECT().UseGrantedFees(ctx, subspaceID, granter, signer, feeAmount, []sdk.Msg{subspaceMsg}).Return(false)
				suite.ak.EXPECT().GetModuleAddress(authtypes.FeeCollectorName).Return(module).Times(2)
				suite.ak.EXPECT().GetAccount(ctx, granter).Return(authtypes.NewBaseAccountWithAddress(granter))
				suite.fk.EXPECT().UseGrantedFees(ctx, granter, signer, feeAmount, []sdk.Msg{subspaceMsg}).Return(nil)
				suite.bk.EXPECT().SendCoinsFromAccountToModule(ctx, granter, authtypes.FeeCollectorName, feeAmount).Return(nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg)
				txBuilder.SetFeeGranter(granter)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
		{
			name: "non subspace tx using auth.DeductFeeDecorator returns no error",
			malleate: func(ctx sdk.Context) {
				suite.ak.EXPECT().GetModuleAddress(authtypes.FeeCollectorName).Return(module)
				suite.ak.EXPECT().GetAccount(ctx, signer).Return(authtypes.NewBaseAccountWithAddress(signer))
				suite.bk.EXPECT().SendCoinsFromAccountToModule(ctx, signer, authtypes.FeeCollectorName, feeAmount).Return(nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(nonSubspaceMsg)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
		{
			name: "subspace tx with different subspaces msgs using auth.DeductFeeDecorator returns no error",
			malleate: func(ctx sdk.Context) {
				suite.ak.EXPECT().GetModuleAddress(authtypes.FeeCollectorName).Return(module)
				suite.ak.EXPECT().GetAccount(ctx, signer).Return(authtypes.NewBaseAccountWithAddress(signer))
				suite.bk.EXPECT().SendCoinsFromAccountToModule(ctx, signer, authtypes.FeeCollectorName, feeAmount).Return(nil)
			},
			buildTx: func() sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(subspaceMsg, otherSubspaceMsg)
				txBuilder.SetFeeAmount(feeAmount)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.malleate != nil {
				tc.malleate(ctx)
			}
			tx := tc.buildTx()
			ctx, err := suite.ante.AnteHandle(ctx, tx, false, func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
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
