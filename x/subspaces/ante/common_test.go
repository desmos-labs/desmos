package ante_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/x/subspaces/ante"
	"github.com/desmos-labs/desmos/v4/x/subspaces/ante/testutil"

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

func TestAnteTestSuite(t *testing.T) {
	suite.Run(t, new(AnteTestSuite))
}

func (suite *AnteTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.ctx = sdktestutil.DefaultContext(sdk.NewKVStoreKey("kv_test"), sdk.NewTransientStoreKey("transient_test"))
	suite.bk = testutil.NewMockBankKeeper(ctrl)
	suite.sk = testutil.NewMockSubspacesKeeper(ctrl)
	suite.ak = testutil.NewMockAccountKeeper(ctrl)
	suite.authDeductFeeDecorator = testutil.NewMockAuthDeductFeeDecorator(ctrl)

	encodingConfig := app.MakeTestEncodingConfig()
	suite.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)

	suite.ante = ante.NewDeductFeeDecorator(suite.authDeductFeeDecorator, suite.ak, suite.bk, suite.sk)
}
