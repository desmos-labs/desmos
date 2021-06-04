package types_test

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/desmos-labs/desmos/app"

	"github.com/desmos-labs/desmos/x/profiles/types"

	"github.com/stretchr/testify/require"
)

func TestProof_Validate(t *testing.T) {
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	plainText := "test"
	sig, err := privKey.Sign([]byte(plainText))
	require.NoError(t, err)
	sigHex := hex.EncodeToString(sig)

	tests := []struct {
		name     string
		proof    types.Proof
		expError error
	}{
		{
			name: "Invalid signature format returns error",
			proof: types.NewProof(
				pubKey,
				"=",
				plainText,
			),
			expError: fmt.Errorf("failed to decode hex string of signature"),
		},
		{
			name: "Empty plain text returns error",
			proof: types.NewProof(
				pubKey,
				sigHex,
				"",
			),
			expError: fmt.Errorf("plain text cannot be empty or blank"),
		},
		{
			name: "Correct proof returns no error",
			proof: types.NewProof(
				pubKey,
				sigHex,
				plainText,
			),
			expError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expError, test.proof.Validate())
		})
	}
}

func TestProof_Verify(t *testing.T) {
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	plainText := "test"
	sig, err := privKey.Sign([]byte(plainText))
	require.NoError(t, err)
	sigHex := hex.EncodeToString(sig)

	tests := []struct {
		name     string
		proof    types.Proof
		expError error
	}{
		{
			name: "Verify proof fails returns error",
			proof: types.NewProof(
				pubKey,
				sigHex,
				"wrong",
			),
			expError: fmt.Errorf("failed to verify the signature"),
		},
		{
			name: "Correct proof returns no error",
			proof: types.NewProof(
				pubKey,
				sigHex,
				plainText,
			),
			expError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cdc, _ := app.MakeCodecs()
			require.Equal(t, test.expError, test.proof.Verify(cdc))
		})
	}
}

func TestChainConfig_Validate(t *testing.T) {
	tests := []struct {
		name        string
		chainConfig types.ChainConfig
		expError    error
	}{
		{
			name: "Empty chain name returns error",
			chainConfig: types.NewChainConfig(
				"",
			),
			expError: fmt.Errorf("chain name cannot be empty or blank"),
		},
		{
			name: "Correct chain config returns no error",
			chainConfig: types.NewChainConfig(
				"test",
			),
			expError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expError, test.chainConfig.Validate())
		})
	}
}

func TestChainLink_Validate(t *testing.T) {
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	addr, err := bech32.ConvertAndEncode("cosmos", pubKey.Address().Bytes())
	require.NoError(t, err)

	plainText := addr
	sig, err := privKey.Sign([]byte(plainText))
	require.NoError(t, err)
	sigHex := hex.EncodeToString(sig)

	tests := []struct {
		name      string
		chainLink types.ChainLink
		expError  error
	}{
		{
			name: "Invalid Proof returns error",
			chainLink: types.NewChainLink(
				types.NewBech32Address(addr, "cosmos"),
				types.NewProof(pubKey, "=", plainText),
				types.NewChainConfig("cosmos"),
				time.Time{},
			),
			expError: fmt.Errorf("failed to decode hex string of signature"),
		},
		{
			name: "Invalid chain config returns error",
			chainLink: types.NewChainLink(
				types.NewBech32Address(addr, "cosmos"),
				types.NewProof(pubKey, sigHex, plainText),
				types.NewChainConfig(""),
				time.Time{},
			),
			expError: fmt.Errorf("chain name cannot be empty or blank"),
		},
		{
			name: "invalid time returns no error",
			chainLink: types.NewChainLink(
				types.NewBech32Address(addr, "cosmos"),
				types.NewProof(pubKey, sigHex, plainText),
				types.NewChainConfig("cosmos"),
				time.Time{},
			),
			expError: fmt.Errorf("createion time cannot be zero"),
		},
		{
			name: "Correct chain link returns no error",
			chainLink: types.NewChainLink(
				types.NewBech32Address(addr, "cosmos"),
				types.NewProof(pubKey, sigHex, plainText),
				types.NewChainConfig("cosmos"),
				time.Now(),
			),
			expError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.chainLink)
			require.Equal(t, test.expError, test.chainLink.Validate())
		})
	}
}

func Test_ChainLinkMarshaling(t *testing.T) {
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	addr, err := bech32.ConvertAndEncode("cosmos", pubKey.Address().Bytes())
	require.NoError(t, err)

	plainText := addr
	sig, err := privKey.Sign([]byte(plainText))
	require.NoError(t, err)
	sigHex := hex.EncodeToString(sig)

	cdc, _ := app.MakeCodecs()
	chainLink := types.NewChainLink(
		types.NewBech32Address(addr, "cosmos"),
		types.NewProof(pubKey, sigHex, plainText),
		types.NewChainConfig("cosmos"),
		time.Time{},
	)
	marshalled := types.MustMarshalChainLink(cdc, chainLink)
	unmarshalled := types.MustUnmarshalChainLink(cdc, marshalled)
	require.Equal(t, chainLink, unmarshalled)
}
