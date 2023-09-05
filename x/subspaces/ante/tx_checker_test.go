package ante_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	relationshipstypes "github.com/desmos-labs/desmos/v6/x/relationships/types"
	subspacesante "github.com/desmos-labs/desmos/v6/x/subspaces/ante"
	"github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func (suite *AnteTestSuite) TestCheckTxFeeWithSubspaceMinPrices() {
	signer := sdk.MustAccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
	msg := relationshipstypes.NewMsgCreateRelationship("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", 1)
	manageSubspaceMsg := types.NewMsgAddUserToUserGroup(1, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
	nonSubspaceMsg := testdata.NewTestMsg(signer)

	testCases := []struct {
		name      string
		setup     func()
		fees      sdk.Coins
		buildTx   func(fees sdk.Coins) sdk.Tx
		check     func(ctx sdk.Context)
		shouldErr bool
	}{
		{
			name: "standard tx returns no error",
			fees: sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))),
			buildTx: func(fees sdk.Coins) sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(nonSubspaceMsg)
				txBuilder.SetFeeAmount(fees)
				txBuilder.SetGasLimit(100)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
		{
			name: "standard tx with insufficient fees returns error",
			fees: sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10))),
			buildTx: func(fees sdk.Coins) sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(nonSubspaceMsg)
				txBuilder.SetFeeAmount(fees)
				txBuilder.SetGasLimit(100)
				return txBuilder.GetTx()
			},
			shouldErr: true,
		},
		{
			name: "standard tx with subspace fees returns error",
			fees: sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(100))),
			buildTx: func(fees sdk.Coins) sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(nonSubspaceMsg)
				txBuilder.SetFeeAmount(fees)
				txBuilder.SetGasLimit(1000)
				return txBuilder.GetTx()
			},
			shouldErr: true,
		},
		{
			name: "subspace not found returns no error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(types.Subspace{}, false)
			},
			fees: sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))),
			buildTx: func(fees sdk.Coins) sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(msg)
				txBuilder.SetFeeAmount(fees)
				txBuilder.SetGasLimit(100)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
		{
			name: "subspace not found returns error - with subspace fee token",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(types.Subspace{}, false)
			},
			fees: sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(100))),
			buildTx: func(fees sdk.Coins) sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(msg)
				txBuilder.SetFeeAmount(fees)
				txBuilder.SetGasLimit(100)
				return txBuilder.GetTx()
			},
			shouldErr: true,
		},
		{
			name: "manage subspace tx with subspace fees returns error",
			fees: sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(100))),
			buildTx: func(fees sdk.Coins) sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(manageSubspaceMsg)
				txBuilder.SetFeeAmount(fees)
				txBuilder.SetGasLimit(100)
				return txBuilder.GetTx()
			},
			shouldErr: true,
		},
		{
			name: "manage subspace tx returns no error",
			fees: sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))),
			buildTx: func(fees sdk.Coins) sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(manageSubspaceMsg)
				txBuilder.SetFeeAmount(fees)
				txBuilder.SetGasLimit(100)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
		{
			name: "non manage subspace tx with insufficient subspace fees returns error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(
						types.Subspace{
							AdditionalFeeTokens: sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(1))),
						},
						true,
					)
			},
			fees: sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(100))),
			buildTx: func(fees sdk.Coins) sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(msg)
				txBuilder.SetFeeAmount(fees)
				txBuilder.SetGasLimit(1000)
				return txBuilder.GetTx()
			},
			shouldErr: true,
		},
		{
			name: "non manage subspace tx returns no error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(
						types.Subspace{
							AdditionalFeeTokens: sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(1))),
						},
						true,
					)
			},
			fees: sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))),
			buildTx: func(fees sdk.Coins) sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(msg)
				txBuilder.SetFeeAmount(fees)
				txBuilder.SetGasLimit(100)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
		{
			name: "non manage subspace tx with subspace fees returns no error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(
						types.Subspace{
							AdditionalFeeTokens: sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(1))),
						},
						true,
					)
			},
			fees: sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(100))),
			buildTx: func(fees sdk.Coins) sdk.Tx {
				txBuilder := suite.clientCtx.TxConfig.NewTxBuilder()
				txBuilder.SetMsgs(msg)
				txBuilder.SetFeeAmount(fees)
				txBuilder.SetGasLimit(1)
				return txBuilder.GetTx()
			},
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			if tc.setup != nil {
				tc.setup()
			}

			tx := tc.buildTx(tc.fees)
			coins, _, err := subspacesante.CheckTxFeeWithSubspaceMinPrices(ante.CheckTxFeeWithValidatorMinGasPrices, suite.sk)(suite.ctx, tx)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(coins, tc.fees)
			}
		})
	}
}

func TestAnte_MergeMinPrices(t *testing.T) {
	testCases := []struct {
		name     string
		original sdk.DecCoins
		others   sdk.DecCoins
		result   sdk.DecCoins
	}{
		{
			name:     "merge existing coin does not update amount properly",
			original: sdk.NewDecCoins(sdk.NewDecCoin("udsm", sdk.NewInt(1000))),
			others:   sdk.NewDecCoins(sdk.NewDecCoin("udsm", sdk.NewInt(5000))),
			result:   sdk.NewDecCoins(sdk.NewDecCoin("udsm", sdk.NewInt(1000))),
		},
		{
			name:     "merge non-existing coin properly",
			original: sdk.NewDecCoins(sdk.NewDecCoin("udsm", sdk.NewInt(1000))),
			others:   sdk.NewDecCoins(sdk.NewDecCoin("minttoken", sdk.NewInt(5000))),
			result: sdk.NewDecCoins(
				sdk.NewDecCoin("udsm", sdk.NewInt(1000)),
				sdk.NewDecCoin("minttoken", sdk.NewInt(5000)),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.result, subspacesante.MergeMinPrices(tc.original, tc.others))
		})
	}
}
