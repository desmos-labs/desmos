package types_test

import (
	"encoding/hex"

	"github.com/desmos-labs/desmos/v2/testutil"

	"github.com/mr-tron/base58"

	"testing"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	"github.com/desmos-labs/desmos/v2/app"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func TestChainConfig_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		chainConfig types.ChainConfig
		shouldErr   bool
	}{
		{
			name:        "empty chain name returns error",
			chainConfig: types.NewChainConfig(""),
			shouldErr:   true,
		},
		{
			name:        "blank chain name returns error",
			chainConfig: types.NewChainConfig("    "),
			shouldErr:   true,
		},
		{
			name:        "uppercase chain name returns error",
			chainConfig: types.NewChainConfig("TC"),
			shouldErr:   true,
		},
		{
			name:        "correct chain config returns no error",
			chainConfig: types.NewChainConfig("tc"),
			shouldErr:   false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.chainConfig.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestProof_Validate(t *testing.T) {

	testCases := []struct {
		name      string
		proof     types.Proof
		shouldErr bool
	}{
		{
			name:      "null public key returns error",
			proof:     types.Proof{Signature: &codectypes.Any{}, PlainText: "74657874"},
			shouldErr: true,
		},
		{
			name:      "null signature returns error",
			proof:     types.NewProof(secp256k1.GenPrivKey().PubKey(), nil, "74657874"),
			shouldErr: true,
		},
		{
			name:      "empty plain text returns error",
			proof:     types.NewProof(secp256k1.GenPrivKey().PubKey(), testutil.SingleSignatureProtoFromHex("74657874"), ""),
			shouldErr: true,
		},
		{
			name:      "invalid plain text format returns error",
			proof:     types.NewProof(secp256k1.GenPrivKey().PubKey(), testutil.SingleSignatureProtoFromHex("74657874"), "="),
			shouldErr: true,
		},
		{
			name: "valid proof returns no error",
			proof: types.NewProof(
				secp256k1.GenPrivKey().PubKey(),
				testutil.SingleSignatureProtoFromHex("74657874"),
				"74657874",
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.proof.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestProof_Verify(t *testing.T) {
	plainText := "tc"

	// Bech32
	bech32PrivKey := secp256k1.GenPrivKey()
	bech32PubKey := bech32PrivKey.PubKey()
	bech32Addr, err := sdk.Bech32ifyAddressBytes("cosmos", bech32PubKey.Address())
	require.NoError(t, err)

	bech32Sig, err := bech32PrivKey.Sign([]byte(plainText))
	require.NoError(t, err)
	bech32SigData := &types.SingleSignatureData{
		Mode:      signing.SignMode_SIGN_MODE_DIRECT,
		Signature: bech32Sig,
	}
	anySigData, err := codectypes.NewAnyWithValue(bech32SigData)
	require.NoError(t, err)

	// Base58
	base58PrivKeyBz, err := hex.DecodeString("bb98111da675930d32f79451fa8d05f188289699558c17148a5d9c82cdb31d1fe04fb0a0d9e689b436b59eff9676d7f2d788244cc4ccfc5768fe117efbd0f9d3")
	require.NoError(t, err)
	base58PrivKey := ed25519.PrivKey{Key: base58PrivKeyBz}
	base58PubKey := base58PrivKey.PubKey()
	base58Addr := base58.Encode(base58PubKey.Bytes())

	base58Sig, err := base58PrivKey.Sign([]byte(plainText))
	require.NoError(t, err)
	base58SigData := &types.SingleSignatureData{
		Mode:      signing.SignMode_SIGN_MODE_DIRECT,
		Signature: base58Sig,
	}

	// Hex
	hexPrivKeyBz, err := hex.DecodeString("2842d8f3701d16711b9ee320f32efe38e6b0891e243eaf6515250e7b006de53e")
	require.NoError(t, err)
	hexPrivKey := secp256k1.PrivKey{Key: hexPrivKeyBz}
	hexPubKey := hexPrivKey.PubKey()

	hexAddr := "0x941991947B6eC9F5537bcaC30C1295E8154Df4cC"
	hexSig, err := hexPrivKey.Sign([]byte(plainText))
	require.NoError(t, err)
	hexSigData := &types.SingleSignatureData{
		Mode:      signing.SignMode_SIGN_MODE_DIRECT,
		Signature: hexSig,
	}

	invalidAny, err := codectypes.NewAnyWithValue(bech32PrivKey)
	require.NoError(t, err)

	testCases := []struct {
		name        string
		proof       types.Proof
		addressData types.AddressData
		shouldErr   bool
	}{
		{
			name:        "Invalid public key value returns error",
			proof:       types.Proof{PubKey: invalidAny, Signature: anySigData, PlainText: hex.EncodeToString([]byte(plainText))},
			addressData: types.NewBech32Address(bech32Addr, "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "wrong plain text returns error",
			proof:       types.NewProof(bech32PubKey, bech32SigData, "wrong"),
			addressData: types.NewBech32Address(bech32Addr, "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "wrong signature returns error",
			proof:       types.NewProof(bech32PubKey, testutil.SingleSignatureProtoFromHex("74657874"), hex.EncodeToString([]byte(plainText))),
			addressData: types.NewBech32Address(bech32Addr, "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "wrong Bech32 address returns error",
			proof:       types.NewProof(bech32PubKey, bech32SigData, hex.EncodeToString([]byte(plainText))),
			addressData: types.NewBech32Address("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn", "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "wrong Base58 address returns error",
			proof:       types.NewProof(base58PubKey, base58SigData, hex.EncodeToString([]byte(plainText))),
			addressData: types.NewBase58Address("HWQ14mk82aRMAad2TdxFHbeqLeUGo5SiBxTXyZyTesJT"),
			shouldErr:   true,
		},
		{
			name:        "wrong Hex address returns error",
			proof:       types.NewProof(hexPubKey, hexSigData, hex.EncodeToString([]byte(plainText))),
			addressData: types.NewHexAddress("0xcdAFfbFd8c131464fEE561e3d9b585141e403719", "0x"),
			shouldErr:   true,
		},
		{
			name:        "correct proof with Base58 address returns no error",
			proof:       types.NewProof(base58PubKey, base58SigData, hex.EncodeToString([]byte(plainText))),
			addressData: types.NewBase58Address(base58Addr),
			shouldErr:   false,
		},
		{
			name:        "correct proof with Bech32 address returns no error",
			proof:       types.NewProof(bech32PubKey, bech32SigData, hex.EncodeToString([]byte(plainText))),
			addressData: types.NewBech32Address(bech32Addr, "cosmos"),
			shouldErr:   false,
		},
		{
			name:        "correct proof with Hex address returns no error",
			proof:       types.NewProof(hexPubKey, hexSigData, hex.EncodeToString([]byte(plainText))),
			addressData: types.NewHexAddress(hexAddr, "0x"),
			shouldErr:   false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cdc, _ := app.MakeCodecs()
			err := tc.proof.Verify(cdc, tc.addressData)

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestBech32Address_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		address   *types.Bech32Address
		shouldErr bool
	}{
		{
			name:      "empty address returns error",
			address:   types.NewBech32Address("", ""),
			shouldErr: true,
		},
		{
			name:      "empty prefix returns error",
			address:   types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", ""),
			shouldErr: true,
		},
		{
			name:      "wrong prefix returns error",
			address:   types.NewBech32Address("desmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos"),
			shouldErr: true,
		},
		{
			name:      "invalid address returns error",
			address:   types.NewBech32Address("desmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0", "desmos"),
			shouldErr: true,
		},
		{
			name:      "valid address returns no error",
			address:   types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.address.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBech32Address_GetValue(t *testing.T) {
	data := types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos")
	require.Equal(t, "cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", data.GetValue())
}

// --------------------------------------------------------------------------------------------------------------------

func TestBase58Address_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		address   *types.Base58Address
		shouldErr bool
	}{
		{
			name:      "empty address returns error",
			address:   types.NewBase58Address(""),
			shouldErr: true,
		},
		{
			name:      "invalid address returns error",
			address:   types.NewBase58Address("0OiIjJ"),
			shouldErr: true,
		},
		{
			name:      "valid address returns no error",
			address:   types.NewBase58Address("5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.address.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBase58Address_GetValue(t *testing.T) {
	data := types.NewBase58Address("5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN")
	require.Equal(t, "5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN", data.GetValue())
}

// --------------------------------------------------------------------------------------------------------------------

func TestHexAddress_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		address   *types.HexAddress
		shouldErr bool
	}{
		{
			name:      "empty and blank address returns error",
			address:   types.NewHexAddress("  ", ""),
			shouldErr: true,
		},
		{
			name:      "address value shorter than prefix returns error",
			address:   types.NewHexAddress("0", "0x"),
			shouldErr: true,
		},
		{
			name:      "not matching prefix returns error",
			address:   types.NewHexAddress("0184", "0x"),
			shouldErr: true,
		},
		{
			name:      "invalid address returns error",
			address:   types.NewHexAddress("0x0OiIjJ", "0x"),
			shouldErr: true,
		},
		{
			name:      "spaced address returns error",
			address:   types.NewHexAddress("0x 941991947B6eC9F5537bcaC30C1295E8154Df4cC", "0x"),
			shouldErr: true,
		},
		{
			name:      "valid address returns no error",
			address:   types.NewHexAddress("0x941991947B6eC9F5537bcaC30C1295E8154Df4cC", "0x"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.address.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestHexAddress_GetValue(t *testing.T) {
	data := types.NewHexAddress("0x941991947B6eC9F5537bcaC30C1295E8154Df4cC", "0x")
	require.Equal(t, "0x941991947B6eC9F5537bcaC30C1295E8154Df4cC", data.GetValue())
}

// --------------------------------------------------------------------------------------------------------------------

func TestUnpackAddressData(t *testing.T) {
	testCases := []struct {
		name        string
		address     *codectypes.Any
		shouldError bool
	}{
		{
			name:        "invalid address returns error",
			address:     testutil.NewAny(secp256k1.GenPrivKey()),
			shouldError: true,
		},
		{
			name:        "valid Bech32 data returns no error",
			address:     testutil.NewAny(types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos")),
			shouldError: false,
		},
		{
			name:        "valid Base58 data returns no error",
			address:     testutil.NewAny(types.NewBase58Address("5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN")),
			shouldError: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cdc, _ := app.MakeCodecs()
			_, err := types.UnpackAddressData(cdc, tc.address)

			if tc.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestChainLink_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		chainLink types.ChainLink
		shouldErr bool
	}{
		{
			name: "empty address returns error",
			chainLink: types.ChainLink{
				User:         "cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
				Proof:        types.NewProof(secp256k1.GenPrivKey().PubKey(), testutil.SingleSignatureProtoFromHex("74657874"), "74657874"),
				ChainConfig:  types.NewChainConfig("cosmos"),
				CreationTime: time.Now(),
			},
			shouldErr: true,
		},
		{
			name: "invalid user returns error",
			chainLink: types.NewChainLink(
				"",
				types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), testutil.SingleSignatureProtoFromHex("74657874"), "74657874"),
				types.NewChainConfig("cosmos"),
				time.Now(),
			),
			shouldErr: true,
		},
		{
			name: "invalid proof returns error",
			chainLink: types.NewChainLink(
				"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
				types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), nil, "74657874"),
				types.NewChainConfig("cosmos"),
				time.Now(),
			),
			shouldErr: true,
		},
		{
			name: "invalid chain config returns error",
			chainLink: types.NewChainLink(
				"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
				types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), testutil.SingleSignatureProtoFromHex("74657874"), "74657874"),
				types.NewChainConfig(""),
				time.Now(),
			),
			shouldErr: true,
		},
		{
			name: "invalid time returns error",
			chainLink: types.NewChainLink(
				"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
				types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), testutil.SingleSignatureProtoFromHex("74657874"), "74657874"),
				types.NewChainConfig("cosmos"),
				time.Time{},
			),
			shouldErr: true,
		},
		{
			name: "valid chain link returns no error",
			chainLink: types.NewChainLink(
				"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
				types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), testutil.SingleSignatureProtoFromHex("74657874"), "74657874"),
				types.NewChainConfig("cosmos"),
				time.Now(),
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.chainLink.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestChainLinkMarshaling(t *testing.T) {
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()

	addr, err := bech32.ConvertAndEncode("cosmos", pubKey.Address().Bytes())
	require.NoError(t, err)

	chainLink := types.NewChainLink(
		"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
		types.NewBech32Address(addr, "cosmos"),
		types.NewProof(pubKey, &types.SingleSignatureData{}, "plain-text"),
		types.NewChainConfig("cosmos"),
		time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
	)

	cdc, _ := app.MakeCodecs()
	marshalled := types.MustMarshalChainLink(cdc, chainLink)
	unmarshalled := types.MustUnmarshalChainLink(cdc, marshalled)
	require.Equal(t, chainLink, unmarshalled)
}
