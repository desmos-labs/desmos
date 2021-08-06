package cmd_test

import (
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/app"
	cmd "github.com/desmos-labs/desmos/app/desmos/cmd"
	profilescliutils "github.com/desmos-labs/desmos/x/profiles/client/utils"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestGetGenerateChainlinkJsonCmd(t *testing.T) {
	cfg := sdk.GetConfig()
	app.SetupConfig(cfg)

	keyBase := keyring.NewInMemory()
	algo := hd.Secp256k1
	hdPath := sdk.GetConfig().GetFullFundraiserPath()

	keyName := "test"
	mnemonic := "clip toilet stairs jaguar baby over mosquito capital speed mule adjust eye print voyage verify smart open crack imitate auto gauge museum planet rebel"
	_, err := keyBase.NewAccount(keyName, mnemonic, "", hdPath, algo)
	require.NoError(t, err)

	output := os.Stdout
	clientCtx := client.Context{}.
		WithKeyring(keyBase).
		WithOutput(output)

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd.GetGenerateChainlinkJsonCmd(), []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, keyName),
	})
	require.NoError(t, err)

	key, err := keyBase.Key(keyName)
	addr, _ := sdk.Bech32ifyAddressBytes(app.Bech32MainPrefix, key.GetAddress())
	sig, pubkey, err := clientCtx.Keyring.Sign(keyName, []byte(addr))
	require.NoError(t, err)

	cdc, _ := app.MakeCodecs()
	var data profilescliutils.ChainLinkJSON
	err = cdc.UnmarshalJSON(out.Bytes(), &data)
	require.NoError(t, err)

	expected := profilescliutils.NewChainLinkJSON(
		types.NewBech32Address(addr, app.Bech32MainPrefix),
		types.NewProof(pubkey, hex.EncodeToString(sig), addr),
		types.NewChainConfig(app.Bech32MainPrefix),
	)

	require.Equal(t, expected, data)

}
