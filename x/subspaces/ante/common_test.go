package ante_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/subspaces/ante"
	"github.com/desmos-labs/desmos/v6/x/subspaces/ante/testutil"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type AnteTestSuite struct {
	suite.Suite

	ctx       sdk.Context
	clientCtx client.Context

	ctrl                   *gomock.Controller
	authDeductFeeDecorator *testutil.MockAuthDeductFeeDecorator
	ak                     *testutil.MockAccountKeeper
	bk                     *testutil.MockBankKeeper
	sk                     *testutil.MockSubspacesKeeper

	ante ante.DeductFeeDecorator
}

func MockTxFeeChecker(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
	feeTx := tx.(sdk.FeeTx)
	return feeTx.GetFee(), 10, nil
}

func TestAnteTestSuite(t *testing.T) {
	suite.Run(t, new(AnteTestSuite))
}

func (suite *AnteTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())

	suite.ctx = sdktestutil.
		DefaultContext(storetypes.NewKVStoreKey("kv_test"), storetypes.NewTransientStoreKey("transient_test")).
		WithMinGasPrices(sdk.NewDecCoins(sdk.NewDecCoin("stake", math.NewInt(1)))).
		WithIsCheckTx(true)

	suite.bk = testutil.NewMockBankKeeper(suite.ctrl)
	suite.sk = testutil.NewMockSubspacesKeeper(suite.ctrl)
	suite.ak = testutil.NewMockAccountKeeper(suite.ctrl)
	suite.authDeductFeeDecorator = testutil.NewMockAuthDeductFeeDecorator(suite.ctrl)

	encodingConfig := app.MakeEncodingConfig()
	suite.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)

	suite.ante = ante.NewDeductFeeDecorator(suite.authDeductFeeDecorator, suite.ak, suite.bk, suite.sk, MockTxFeeChecker)
}

func (suite *AnteTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}
