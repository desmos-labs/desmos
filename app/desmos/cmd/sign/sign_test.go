//go:build norace
// +build norace

package sign_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	cmd "github.com/desmos-labs/desmos/v4/app/desmos/cmd/sign"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
)

func TestGetSignCmd(t *testing.T) {
	cfg := sdk.GetConfig()
	app.SetupConfig(cfg)

	cdc, _ := app.MakeCodecs()
	keyBase := keyring.NewInMemory(cdc)
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

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd.GetSignCmd(), []string{
		"This is my signed value",
		fmt.Sprintf("--%s=%s", flags.FlagFrom, keyName),
	})
	require.NoError(t, err)

	var data cmd.SignatureData
	err = json.Unmarshal(out.Bytes(), &data)
	require.NoError(t, err)

	expected := cmd.SignatureData{
		Address:   "8cbc1d27bd2c8675935f6018f08fed08c7add0d9",
		PubKey:    "0235088ee6a12267eda6410028706c4ec192b1ce04a298def8dad1c146257012eb",
		Signature: "39cc04208d25e445a46bd6853cfa7b1b295961bc43d6facf387f4bd9c0c4647163ef2e296f49e307cdb657a7acdf4e2ee56e9bbf1f9723113385f4dd3601061e",
		Value:     "54686973206973206d79207369676e65642076616c7565",
	}
	require.Equal(t, expected, data)

}
