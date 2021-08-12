package cmd_test

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/app"
	cmd "github.com/desmos-labs/desmos/app/desmos/cmd"
	"github.com/desmos-labs/desmos/app/desmos/types"

	profilescliutils "github.com/desmos-labs/desmos/x/profiles/client/utils"
	profilestypes "github.com/desmos-labs/desmos/x/profiles/types"
)

// Create mock chain type generator instead of prompt
type MockChainTypeGenerator struct{}

func (mock MockChainTypeGenerator) GetReference() (string, types.ChainType, error) {
	return "clip toilet stairs jaguar baby over mosquito capital speed mule adjust eye print voyage verify smart open crack imitate auto gauge museum planet rebel",
		types.ChainType{ID: "Cosmos", Name: "cosmos", Prefix: "cosmos", DerivationPath: "m/44'/118'/0'/0/0"},
		nil
}

func TestGetGenerateChainlinkJsonCmd(t *testing.T) {
	cfg := sdk.GetConfig()
	app.SetupConfig(cfg)

	output := os.Stdout
	clientCtx := client.Context{}.
		WithOutput(output)

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd.GetCreateChainlinkJSON(MockChainTypeGenerator{}), []string{})
	require.NoError(t, err)

	cdc, _ := app.MakeCodecs()
	var data profilescliutils.ChainLinkJSON
	err = cdc.UnmarshalJSON(out.Bytes(), &data)
	require.NoError(t, err)

	// create a account to inmemory keybase
	keyBase := keyring.NewInMemory()
	algo := hd.Secp256k1
	hdPath := "m/44'/118'/0'/0/0"
	keyName := "chainlink"
	_, err = keyBase.NewAccount(
		"chainlink",
		"clip toilet stairs jaguar baby over mosquito capital speed mule adjust eye print voyage verify smart open crack imitate auto gauge museum planet rebel",
		"",
		hdPath,
		algo,
	)
	require.NoError(t, err)

	// get key from keybase
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
