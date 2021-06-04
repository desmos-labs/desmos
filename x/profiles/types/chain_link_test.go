package types_test

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/desmos-labs/desmos/app"

	"github.com/desmos-labs/desmos/x/profiles/types"

	"github.com/stretchr/testify/require"
)

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

// ___________________________________________________________________________________________________________________

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
			name: "Null public key returns error",
			proof: types.Proof{
				nil,
				sigHex,
				plainText,
			},
			expError: fmt.Errorf("public key field cannot be nil"),
		},
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
	invalidAny, err := codectypes.NewAnyWithValue(privKey)
	require.NoError(t, err)

	tests := []struct {
		name     string
		proof    types.Proof
		expError error
	}{
		{
			name: "Unpack public key failed returns error",
			proof: types.Proof{
				invalidAny,
				sigHex,
				plainText,
			},
			expError: fmt.Errorf("failed to unpack the pubkey"),
		},
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

// ___________________________________________________________________________________________________________________

func Test_Bech32AddressValidate(t *testing.T) {
	tests := []struct {
		name     string
		address  *types.Bech32Address
		expError error
	}{
		{
			name:     "Empty address returns error",
			address:  types.NewBech32Address("", ""),
			expError: fmt.Errorf("address cannot be empty or blank"),
		},
		{
			name:     "Empty prefix returns error",
			address:  types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", ""),
			expError: fmt.Errorf("prefix cannot be empty or blank"),
		},
		{
			name:     "Invalid address returns error",
			address:  types.NewBech32Address("desmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos"),
			expError: fmt.Errorf("invalid address encoded type or wrong prefix"),
		},
		{
			name:     "Valid address returns no error",
			address:  types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos"),
			expError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expError, test.address.Validate())
		})
	}
}

func Test_Bech32AddressGetAddress(t *testing.T) {
	addr := types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos")
	require.Equal(t, "cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", addr.GetAddress())
}

func Test_Base58AddressValidate(t *testing.T) {
	tests := []struct {
		name     string
		address  *types.Base58Address
		expError error
	}{
		{
			name:     "Empty address returns error",
			address:  types.NewBase58Address(""),
			expError: fmt.Errorf("address cannot be empty or blank"),
		},
		{
			name:     "Invalid address returns error",
			address:  types.NewBase58Address("0OiIjJ"),
			expError: fmt.Errorf("invalid address encoded type"),
		},
		{
			name:     "Valid address returns no error",
			address:  types.NewBase58Address("5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN"),
			expError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expError, test.address.Validate())
		})
	}
}

func Test_Base58AddressGetAddress(t *testing.T) {
	addr := types.NewBase58Address("5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN")
	require.Equal(t, "5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN", addr.GetAddress())
}

func Test_UnpackAddressData(t *testing.T) {
	bech32Addr := types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos")
	base58Addr := types.NewBase58Address("5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN")
	privKey := secp256k1.GenPrivKey()

	bech32AddrAny, err := codectypes.NewAnyWithValue(bech32Addr)
	require.NoError(t, err)

	base58AddrAny, err := codectypes.NewAnyWithValue(base58Addr)
	require.NoError(t, err)

	invalidAny, err := codectypes.NewAnyWithValue(privKey)
	require.NoError(t, err)

	tests := []struct {
		name        string
		address     *codectypes.Any
		shouldError bool
	}{
		{
			name:        "Invalid address returns error",
			address:     invalidAny,
			shouldError: true,
		},
		{
			name:        "Valid bech32 address returns no error",
			address:     bech32AddrAny,
			shouldError: false,
		},
		{
			name:        "Valid base58 address returns no error",
			address:     base58AddrAny,
			shouldError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cdc, _ := app.MakeCodecs()
			_, err := types.UnpackAddressData(cdc, test.address)
			if test.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

		})
	}
}

// ___________________________________________________________________________________________________________________

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
			name: "null address returns error",
			chainLink: types.ChainLink{
				nil,
				types.NewProof(pubKey, "=", plainText),
				types.NewChainConfig("cosmos"),
				time.Time{},
			},
			expError: fmt.Errorf("address cannot be nil"),
		},
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
