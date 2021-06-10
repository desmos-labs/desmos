package types_test

import (
	"encoding/hex"
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
		shouldErr   bool
	}{
		{
			name:        "Empty chain name returns error",
			chainConfig: types.NewChainConfig(""),
			shouldErr:   true,
		},
		{
			name:        "Blank chain name returns error",
			chainConfig: types.NewChainConfig("    "),
			shouldErr:   true,
		},
		{
			name:        "Correct chain config returns no error",
			chainConfig: types.NewChainConfig("test"),
			shouldErr:   false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.chainConfig.Validate()
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestProof_Validate(t *testing.T) {
	tests := []struct {
		name      string
		proof     types.Proof
		shouldErr bool
	}{
		{
			name:      "Null public key returns error",
			proof:     types.Proof{Signature: "74657874", PlainText: "text"},
			shouldErr: true,
		},
		{
			name:      "Invalid signature format returns error",
			proof:     types.NewProof(secp256k1.GenPrivKey().PubKey(), "=", "text"),
			shouldErr: true,
		},
		{
			name:      "Empty plain text returns error",
			proof:     types.NewProof(secp256k1.GenPrivKey().PubKey(), "74657874", ""),
			shouldErr: true,
		},
		{
			name:      "Valid proof returns no error",
			proof:     types.NewProof(secp256k1.GenPrivKey().PubKey(), "74657874", "text"),
			shouldErr: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.proof.Validate()
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
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
		name      string
		proof     types.Proof
		shouldErr bool
	}{
		{
			name:      "Invalid public key value returns error",
			proof:     types.Proof{PubKey: invalidAny, Signature: sigHex, PlainText: plainText},
			shouldErr: true,
		},
		{
			name:      "Wrong plain text returns error",
			proof:     types.NewProof(pubKey, sigHex, "wrong"),
			shouldErr: true,
		},
		{
			name:      "Wrong signature returns error",
			proof:     types.NewProof(pubKey, "74657874", plainText),
			shouldErr: true,
		},
		{
			name:      "Correct proof returns no error",
			proof:     types.NewProof(pubKey, sigHex, plainText),
			shouldErr: false,
		},
	}
	for _, test := range tests {
		cdc, _ := app.MakeCodecs()
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.proof.Verify(cdc)
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func Test_Bech32AddressValidate(t *testing.T) {
	tests := []struct {
		name      string
		address   *types.Bech32Address
		shouldErr bool
	}{
		{
			name:      "Empty address returns error",
			address:   types.NewBech32Address("", ""),
			shouldErr: true,
		},
		{
			name:      "Empty prefix returns error",
			address:   types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", ""),
			shouldErr: true,
		},
		{
			name:      "Wrong prefix returns error",
			address:   types.NewBech32Address("desmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos"),
			shouldErr: true,
		},
		{
			name:      "Invalid address returns error",
			address:   types.NewBech32Address("desmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0", "desmos"),
			shouldErr: true,
		},
		{
			name:      "Valid address returns no error",
			address:   types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos"),
			shouldErr: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.address.Validate()
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_Bech32AddressGetAddress(t *testing.T) {
	addr := types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos")
	require.Equal(t, "cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", addr.GetAddress())
}

// --------------------------------------------------------------------------------------------------------------------

func Test_Base58AddressValidate(t *testing.T) {
	tests := []struct {
		name      string
		address   *types.Base58Address
		shouldErr bool
	}{
		{
			name:      "Empty address returns error",
			address:   types.NewBase58Address(""),
			shouldErr: true,
		},
		{
			name:      "Invalid address returns error",
			address:   types.NewBase58Address("0OiIjJ"),
			shouldErr: true,
		},
		{
			name:      "Valid address returns no error",
			address:   types.NewBase58Address("5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN"),
			shouldErr: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.address.Validate()
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_Base58AddressGetAddress(t *testing.T) {
	addr := types.NewBase58Address("5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN")
	require.Equal(t, "5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN", addr.GetAddress())
}

// --------------------------------------------------------------------------------------------------------------------

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
			name:        "Valid Bech32 address returns no error",
			address:     bech32AddrAny,
			shouldError: false,
		},
		{
			name:        "Valid Base58 address returns no error",
			address:     base58AddrAny,
			shouldError: false,
		},
	}

	for _, test := range tests {
		test := test
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

// --------------------------------------------------------------------------------------------------------------------

func TestChainLink_Validate(t *testing.T) {
	tests := []struct {
		name      string
		chainLink types.ChainLink
		shouldErr bool
	}{
		{
			name: "Empty address returns error",
			chainLink: types.ChainLink{
				Proof:        types.NewProof(secp256k1.GenPrivKey().PubKey(), "=", "text"),
				ChainConfig:  types.NewChainConfig("cosmos"),
				CreationTime: time.Now(),
			},
			shouldErr: true,
		},
		{
			name: "Invalid proof returns error",
			chainLink: types.NewChainLink(
				types.NewBech32Address(addr.String(), "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), "=", "text"),
				types.NewChainConfig("cosmos"),
				time.Now(),
			),
			shouldErr: true,
		},
		{
			name: "Invalid chain config returns error",
			chainLink: types.NewChainLink(
				types.NewBech32Address(addr.String(), "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), "74657874", "text"),
				types.NewChainConfig(""),
				time.Now(),
			),
			shouldErr: true,
		},
		{
			name: "Invalid time returns error",
			chainLink: types.NewChainLink(
				types.NewBech32Address(addr.String(), "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), "74657874", "text"),
				types.NewChainConfig("cosmos"),
				time.Time{},
			),
			shouldErr: true,
		},
		{
			name: "Valid chain link returns no error",
			chainLink: types.NewChainLink(
				types.NewBech32Address(addr.String(), "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), "74657874", "text"),
				types.NewChainConfig("cosmos"),
				time.Now(),
			),
			shouldErr: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.chainLink.Validate()
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_ChainLinkMarshaling(t *testing.T) {
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()

	addr, err := bech32.ConvertAndEncode("cosmos", pubKey.Address().Bytes())
	require.NoError(t, err)

	cdc, _ := app.MakeCodecs()
	chainLink := types.NewChainLink(
		types.NewBech32Address(addr, "cosmos"),
		types.NewProof(pubKey, "sig-hex", "plain-text"),
		types.NewChainConfig("cosmos"),
		time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
	)
	marshalled := types.MustMarshalChainLink(cdc, chainLink)
	unmarshalled := types.MustUnmarshalChainLink(cdc, marshalled)
	require.Equal(t, chainLink, unmarshalled)
}
