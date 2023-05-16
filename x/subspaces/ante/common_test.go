package ante_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v5/app"
	"github.com/desmos-labs/desmos/v5/x/subspaces/ante"
	"github.com/desmos-labs/desmos/v5/x/subspaces/ante/testutil"

	"github.com/cosmos/cosmos-sdk/client"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type AnteTestSuite struct {
	suite.Suite

	ctx       sdk.Context
	clientCtx client.Context

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
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	suite.ctx = sdktestutil.DefaultContext(sdk.NewKVStoreKey("kv_test"), sdk.NewTransientStoreKey("transient_test"))
	suite.bk = testutil.NewMockBankKeeper(ctrl)
	suite.sk = testutil.NewMockSubspacesKeeper(ctrl)
	suite.ak = testutil.NewMockAccountKeeper(ctrl)
	suite.authDeductFeeDecorator = testutil.NewMockAuthDeductFeeDecorator(ctrl)

	encodingConfig := app.MakeEncodingConfig()
	suite.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)

	suite.ante = ante.NewDeductFeeDecorator(suite.authDeductFeeDecorator, suite.ak, suite.bk, suite.sk, MockTxFeeChecker)
}
