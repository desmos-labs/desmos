package chainlink_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v3/app"
)

func TestCreateJSONChainLinkSuite(t *testing.T) {
	suite.Run(t, new(CreateJSONChainLinkTestSuite))
}

type CreateJSONChainLinkTestSuite struct {
	suite.Suite

	Codec       codec.Codec
	LegacyAmino *codec.LegacyAmino
	ClientCtx   client.Context
	Owner       string
}

func (suite *CreateJSONChainLinkTestSuite) SetupSuite() {
	cfg := sdk.GetConfig()
	app.SetupConfig(cfg)

	encodingConfig := app.MakeTestEncodingConfig()
	suite.Codec = encodingConfig.Marshaler
	suite.LegacyAmino = encodingConfig.Amino
	suite.ClientCtx = client.Context{}.WithOutput(os.Stdout).WithTxConfig(encodingConfig.TxConfig)
	suite.Owner = "desmos1n8345tvzkg3jumkm859r2qz0v6xsc3henzddcj"
}

func (suite *CreateJSONChainLinkTestSuite) TempFile() string {
	info, err := ioutil.TempFile(suite.T().TempDir(), "Test_")
	suite.Require().NoError(err)
	return info.Name()
}

func (suite *CreateJSONChainLinkTestSuite) GetPubKeyFromTxFile(txFile string) cryptotypes.PubKey {
	parsedTx, err := authclient.ReadTxFromFile(suite.ClientCtx, txFile)
	suite.Require().NoError(err)

	txBuilder, err := suite.ClientCtx.TxConfig.WrapTxBuilder(parsedTx)
	suite.Require().NoError(err)

	sigs, err := txBuilder.GetTx().GetSignaturesV2()
	suite.Require().NoError(err)
	suite.Require().NotEmpty(sigs)

	return sigs[0].PubKey
}

// --------------------------------------------------------------------------------------------------------------------

// MockGetter represents a mock implementation of ChainLinkReferenceGetter
type MockGetter struct {
	FileName string
	IsSingle bool
	TxFile   string
}

func NewMockGetter(fileName string, isSingle bool, txFile string) MockGetter {
	return MockGetter{
		FileName: fileName,
		IsSingle: isSingle,
		TxFile:   txFile,
	}
}

// GetIsSingleSignature implements ChainLinkReferenceGetter
func (mock MockGetter) IsSingleSignatureAccount() (bool, error) {
	return mock.IsSingle, nil
}

// GetChain implements ChainLinkReferenceGetter
func (mock MockGetter) GetChain() (types.Chain, error) {
	return types.NewChain("Cosmos", "cosmos", "cosmos", "m/44'/118'/0'/0/0"), nil
}

// GetFilename implements ChainLinkReferenceGetter
func (mock MockGetter) GetFilename() (string, error) {
	return mock.FileName, nil
}

// GetOwner implements ChainLinkReferenceGetter
func (mock MockGetter) GetOwner() (string, error) {
	return "desmos1n8345tvzkg3jumkm859r2qz0v6xsc3henzddcj", nil
}

// GetMnemonic implements SingleSignatureAccountReferenceGetter
func (mock MockGetter) GetMnemonic() (string, error) {
	return "clip toilet stairs jaguar baby over mosquito capital speed mule adjust eye print voyage verify smart open crack imitate auto gauge museum planet rebel", nil
}

// GetMultiSignedTxFilePath implements MultiSignatureAccountReferenceGetter
func (mock MockGetter) GetMultiSignedTxFilePath() (string, error) {
	return mock.TxFile, nil
}

// GetSignedChainID implements MultiSignatureAccountReferenceGetter
func (mock MockGetter) GetSignedChainID() (string, error) {
	return "cosmos", nil
}
