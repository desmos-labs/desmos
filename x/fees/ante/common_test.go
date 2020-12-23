package ante_test

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	cosmosante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	desmos "github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/fees/ante"
	feesTypes "github.com/desmos-labs/desmos/x/fees/types"
)

// TestAccount represents an account used in the ante handler tests
type TestAccount struct {
	acc     authtypes.AccountI
	privKey cryptotypes.PrivKey
}

// AnteTestSuite is a test suite to be used with ante handler tests
type AnteTestSuite struct {
	suite.Suite

	app         *desmos.DesmosApp
	anteHandler sdk.AnteHandler
	ctx         sdk.Context
	clientCtx   client.Context
	txBuilder   client.TxBuilder
}

// createTestApp returns context and app with params set on account keeper
func createTestApp(isCheckTx bool) (*desmos.DesmosApp, sdk.Context) {
	db := dbm.NewMemDB()
	app := desmos.NewDesmosApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{},
		desmos.DefaultNodeHome, 0, desmos.MakeTestEncodingConfig(), simapp.EmptyAppOptions{},
	)

	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.FeesKeeper.SetParams(ctx, feesTypes.DefaultParams())

	return app, ctx
}

func (suite *AnteTestSuite) createTestAccount() TestAccount {
	privKey, _, addr := testdata.KeyTestPubAddr()

	// Set the accounts
	acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr)
	err := acc.SetAccountNumber(uint64(0))
	suite.Require().NoError(err)

	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)
	suite.app.BankKeeper.SetBalances(
		suite.ctx,
		addr,
		sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 100000000000)),
	)

	return TestAccount{
		acc:     acc,
		privKey: privKey,
	}
}

// SetupTest setups a new test, with new app, context, and anteHandler.
func (suite *AnteTestSuite) SetupTest(isCheckTx bool) {
	suite.app, suite.ctx = createTestApp(isCheckTx)
	suite.ctx = suite.ctx.WithBlockHeight(1)

	// Setup TxConfig
	encodingConfig := desmos.MakeTestEncodingConfig()

	suite.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	suite.anteHandler = ante.NewAnteHandler(
		suite.app.AccountKeeper,
		suite.app.BankKeeper,
		cosmosante.DefaultSigVerificationGasConsumer,
		suite.app.FeesKeeper,
		encodingConfig.TxConfig.SignModeHandler(),
	)
}

// CreateTestTx is a helper function to create a tx given multiple inputs.
func (suite *AnteTestSuite) CreateTestTx(privs []cryptotypes.PrivKey, accNums []uint64, accSeqs []uint64, chainID string) (xauthsigning.Tx, error) {
	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	var sigsV2 []signing.SignatureV2
	for i, priv := range privs {
		sigV2 := signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  suite.clientCtx.TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: accSeqs[i],
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err := suite.txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return nil, err
	}

	// Second round: all signer infos are set, so each signer can sign.
	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
		signerData := xauthsigning.SignerData{
			ChainID:       chainID,
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
		}
		sigV2, err := tx.SignWithPrivKey(
			suite.clientCtx.TxConfig.SignModeHandler().DefaultMode(), signerData,
			suite.txBuilder, priv, suite.clientCtx.TxConfig, accSeqs[i])
		if err != nil {
			return nil, err
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err = suite.txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return nil, err
	}

	return suite.txBuilder.GetTx(), nil
}

func TestAnteTestSuite(t *testing.T) {
	suite.Run(t, new(AnteTestSuite))
}
