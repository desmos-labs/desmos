package sign_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	cmd "github.com/desmos-labs/desmos/app/desmos/cmd/sign"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/app"
)

func TestGetSignCmd(t *testing.T) {
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

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd.GetSignCmd(), []string{
		"This is my signed value",
		fmt.Sprintf("--%s=%s", flags.FlagFrom, keyName),
	})
	require.NoError(t, err)

	var data cmd.SignatureData
	err = json.Unmarshal(out.Bytes(), &data)
	require.NoError(t, err)

	expected := cmd.SignatureData{
		Address:   "d133e902b523aa568f0086609c958c83ac0c4fc1",
		PubKey:    "03f3fa3e31c3c2833c92b83f26ef29397991acfeb0fbaad8864d047cdd6a0cc155",
		Signature: "5d492db8df4a6f912188610395574d36cdbaadec00e64bc6341722e81fe38cae7d8b7e6d2896f1fba03e45dd2a9534cab205347f00e177de6796a46811cfa35f",
		Value:     "This is my signed value",
	}
	require.Equal(t, expected, data)

}
