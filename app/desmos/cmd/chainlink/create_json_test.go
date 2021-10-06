package chainlink_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	cmd "github.com/desmos-labs/desmos/v2/app/desmos/cmd/chainlink"
	"github.com/desmos-labs/desmos/v2/app/desmos/cmd/chainlink/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/app"
	profilescliutils "github.com/desmos-labs/desmos/v2/x/profiles/client/utils"
	profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"
)

// MockGetter represents a mock implementation of ChainLinkReferenceGetter
type MockGetter struct {
	FileName string
}

func NewMockGetter(fileName string) MockGetter {
	return MockGetter{
		FileName: fileName,
	}
}

// GetMnemonic implements ChainLinkReferenceGetter
func (mock MockGetter) GetMnemonic() (string, error) {
	return "clip toilet stairs jaguar baby over mosquito capital speed mule adjust eye print voyage verify smart open crack imitate auto gauge museum planet rebel", nil
}

// GetChain implements ChainLinkReferenceGetter
func (mock MockGetter) GetChain() (types.Chain, error) {
	return types.NewChain("Cosmos", "cosmos", "cosmos", "m/44'/118'/0'/0/0"), nil
}

// GetFilename implements ChainLinkReferenceGetter
func (mock MockGetter) GetFilename() (string, error) {
	return mock.FileName, nil
}

func TestGetCreateChainLinkJSON(t *testing.T) {
	cfg := sdk.GetConfig()
	app.SetupConfig(cfg)

	// Define where the data will be saved
	fileName := path.Join(t.TempDir(), "out.json")

	clientCtx := client.Context{}.WithOutput(os.Stdout)
	_, err := clitestutil.ExecTestCLICmd(clientCtx, cmd.GetCreateChainLinkJSON(NewMockGetter(fileName)), []string{})
	require.NoError(t, err)

	out, err := ioutil.ReadFile(fileName)
	require.NoError(t, err)

	cdc, _ := app.MakeCodecs()
	var data profilescliutils.ChainLinkJSON
	err = cdc.UnmarshalJSON(out, &data)
	require.NoError(t, err)

	// Create an account inside the inmemory keybase
	keyBase := keyring.NewInMemory()
	keyName := "chainlink"
	_, err = keyBase.NewAccount(
		"chainlink",
		"clip toilet stairs jaguar baby over mosquito capital speed mule adjust eye print voyage verify smart open crack imitate auto gauge museum planet rebel",
		"",
		"m/44'/118'/0'/0/0",
		hd.Secp256k1,
	)
	require.NoError(t, err)

	// Get the key from the keybase
	key, err := keyBase.Key(keyName)
	require.NoError(t, err)

	expected := profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address("cosmos13j7p6faa9jr8ty6lvqv0prldprr6m5xenmafnt", "cosmos"),
		profilestypes.NewProof(
			key.GetPubKey(),
			"c3bd014b2178d63d94b9c28e628bfcf56736de28f352841b0bb27d6fff2968d62c13a10aeddd1ebfe3b13f3f8e61f79a2c63ae6ff5cb78cb0d64e6b0a70fae57",
			"cosmos13j7p6faa9jr8ty6lvqv0prldprr6m5xenmafnt"),
		profilestypes.NewChainConfig("cosmos"),
	)

	require.Equal(t, expected, data)
}
