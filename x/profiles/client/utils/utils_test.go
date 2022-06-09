package utils_test

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/app"
	"github.com/desmos-labs/desmos/v3/testutil/profilestesting"
	"github.com/desmos-labs/desmos/v3/x/profiles/client/utils"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func TestParseLinkAccountJSON(t *testing.T) {
	mnemonic := "foam panther lottery brisk lesson fan february process fish invest slow argue lonely surround tragic prize grab common conduct coil own vote leopard glow"
	derivedPriv, err := hd.Secp256k1.Derive()(mnemonic, "", sdk.GetConfig().GetFullFundraiserPath())
	require.NoError(t, err)

	privKey := hd.Secp256k1.Generate()(derivedPriv)
	pubKey := privKey.PubKey()

	addStr, err := sdk.Bech32ifyAddressBytes("cosmos", pubKey.Address())
	require.NoError(t, err)

	plainText := addStr
	sigBz, err := privKey.Sign([]byte(plainText))
	require.NoError(t, err)

	jsonData := utils.NewChainLinkJSON(
		types.NewBech32Address(addStr, "cosmos"),
		types.NewProof(pubKey, profilestesting.SingleSignatureProtoFromHex(hex.EncodeToString(sigBz)), plainText),
		types.NewChainConfig("cosmos"),
	)

	params := app.MakeTestEncodingConfig()
	jsonBz := params.Marshaler.MustMarshalJSON(&jsonData)

	// Write the JSON to a temp file
	filePath := path.Join(t.TempDir(), t.Name())
	require.NoError(t, ioutil.WriteFile(filePath, jsonBz, 0666))

	// Read the temp file and check for equality
	data, err := utils.ParseChainLinkJSON(params.Marshaler, filePath)
	require.NoError(t, err)
	require.True(t, jsonData.Equal(data))

	// Delete the temp folder
	require.NoError(t, os.RemoveAll(t.TempDir()))
}
