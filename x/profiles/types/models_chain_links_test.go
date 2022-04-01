package types_test

import (
	"encoding/hex"

	"github.com/desmos-labs/desmos/v3/testutil"

	"github.com/mr-tron/base58"

	"testing"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/desmos-labs/desmos/v3/app"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func TestUnmarshalSignature(t *testing.T) {
	_, amino := app.MakeCodecs()

	expectedMemo := "desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd"

	aminoSigBytes := "7b226163636f756e745f6e756d626572223a2230222c22636861696e5f6964223a22616b6173686e65742d32222c22666565223a7b22616d6f756e74223a5b7b22616d6f756e74223a2230222c2264656e6f6d223a22616b74227d5d2c22676173223a2231227d2c226d656d6f223a226465736d6f7331366336307938743876726132377a6a673261726c6364353864636b3963776e37703666777464222c226d736773223a5b5d2c2273657175656e6365223a2230227d"
	aminoSigBz, err := hex.DecodeString(aminoSigBytes)
	require.NoError(t, err)

	var stdSigDoc legacytx.StdSignDoc
	err = amino.UnmarshalJSON(aminoSigBz, &stdSigDoc)
	require.NoError(t, err)
	require.Equal(t, expectedMemo, stdSigDoc.Memo)

	directSigBytes := "6465736d6f7331366336307938743876726132377a6a673261726c6364353864636b3963776e37703666777464"
	directSigBz, err := hex.DecodeString(directSigBytes)
	require.NoError(t, err)
	require.Equal(t, expectedMemo, string(directSigBz))
}

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
	pubKeyAny, err := codectypes.NewAnyWithValue(secp256k1.GenPrivKey().PubKey())
	require.NoError(t, err)

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
			proof:     types.Proof{PubKey: pubKeyAny, PlainText: "74657874"},
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

// generateMultiSigKeys returns the given amount of private keys, and a new multi sig public key made of such keys
func generateMultiSigKeys(n int) ([]cryptotypes.PrivKey, cryptotypes.PubKey) {
	privKeys := make([]cryptotypes.PrivKey, n)
	pubKeys := make([]cryptotypes.PubKey, n)
	for i := 0; i < n; i++ {
		// Generate the private key
		privKeys[i] = secp256k1.GenPrivKey()
		pubKeys[i] = privKeys[i].PubKey()
	}

	return privKeys, kmultisig.NewLegacyAminoPubKey(n, pubKeys)
}

// generateMultiSigSignatureData uses the given private keys to sign the given message using the multi sig algorithm,
// and returns the obtained MultiSignatureData instance.
func generateMultiSigSignatureData(t *testing.T, privKeys []cryptotypes.PrivKey, msg []byte) types.SignatureData {
	cosmosMultisig := multisig.NewMultisig(len(privKeys))
	pubKeys := make([]cryptotypes.PubKey, len(privKeys))
	for i, privKey := range privKeys {
		pubKeys[i] = privKey.PubKey()
	}

	for i, privKey := range privKeys {
		// Sign the message using the generated private key
		sig, err := privKey.Sign(msg)
		require.NoError(t, err)

		// Build the signature data for the single signature and add it to the multi signature data
		sigData := &signing.SingleSignatureData{Signature: sig}
		err = multisig.AddSignatureFromPubKey(cosmosMultisig, sigData, privKeys[i].PubKey(), pubKeys)
		require.NoError(t, err)
	}

	// Generate the signature data object
	sigData, err := types.SignatureDataFromCosmosSignatureData(cosmosMultisig)
	require.NoError(t, err)

	return sigData
}

func TestProof_Verify(t *testing.T) {
	// Bech32
	bech32PrivKey := secp256k1.GenPrivKey()
	bech32PubKey := bech32PrivKey.PubKey()
	bech32Addr, err := sdk.Bech32ifyAddressBytes("cosmos", bech32PubKey.Address())
	require.NoError(t, err)

	bech32Owner := "cosmos10m20h8fy0qp2a8f46zzjpvg8pfl8flajgxsvmk"
	bech32Sig, err := bech32PrivKey.Sign([]byte(bech32Owner))
	require.NoError(t, err)
	bech32SigData := &types.SingleSignatureData{
		Mode:      signing.SignMode_SIGN_MODE_TEXTUAL,
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

	base58Owner := "cosmos1u55ywhk6thmhnxs7yn8vh8v7eznckcqjevnadx"
	base58Sig, err := base58PrivKey.Sign([]byte(base58Owner))
	require.NoError(t, err)
	base58SigData := &types.SingleSignatureData{
		Mode:      signing.SignMode_SIGN_MODE_TEXTUAL,
		Signature: base58Sig,
	}

	// Hex
	hexPrivKeyBz, err := hex.DecodeString("2842d8f3701d16711b9ee320f32efe38e6b0891e243eaf6515250e7b006de53e")
	require.NoError(t, err)
	hexPrivKey := secp256k1.PrivKey{Key: hexPrivKeyBz}
	hexPubKey := hexPrivKey.PubKey()
	hexAddr := "0x941991947B6eC9F5537bcaC30C1295E8154Df4cC"

	hexOwner := "cosmos1l0g43u695yvmwem09ncwgsxup6m8aklcyr38ph"
	hexSig, err := hexPrivKey.Sign([]byte(hexOwner))
	require.NoError(t, err)
	hexSigData := &types.SingleSignatureData{
		Mode:      signing.SignMode_SIGN_MODE_DIRECT,
		Signature: hexSig,
	}

	// Multisig
	privKeys, multiSigPubKey := generateMultiSigKeys(3)
	multisigAddr, err := sdk.Bech32ifyAddressBytes("cosmos", multiSigPubKey.Address())
	require.NoError(t, err)

	multiSigOwner := "cosmos1vfvst0mr79nzzxsk65uuhfklmwrnsfadhtn977"
	multiSigData := generateMultiSigSignatureData(t, privKeys, []byte(multiSigOwner))
	validMultisigDataAny, err := codectypes.NewAnyWithValue(multiSigData)
	require.NoError(t, err)

	validaPubKeyAny, err := codectypes.NewAnyWithValue(bech32PubKey)
	require.NoError(t, err)
	invalidAny, err := codectypes.NewAnyWithValue(bech32PrivKey)
	require.NoError(t, err)

	testCases := []struct {
		name        string
		proof       types.Proof
		owner       string
		addressData types.AddressData
		shouldErr   bool
	}{
		{
			name:        "invalid public key value returns error",
			proof:       types.Proof{PubKey: invalidAny, Signature: anySigData, PlainText: hex.EncodeToString([]byte(bech32Owner))},
			owner:       bech32Owner,
			addressData: types.NewBech32Address(bech32Addr, "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "invalid signature value returns error",
			proof:       types.Proof{PubKey: validaPubKeyAny, Signature: invalidAny, PlainText: hex.EncodeToString([]byte(bech32Owner))},
			owner:       bech32Owner,
			addressData: types.NewBech32Address(bech32Addr, "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "wrong plain text returns error",
			proof:       types.NewProof(bech32PubKey, bech32SigData, "wrong"),
			owner:       bech32Owner,
			addressData: types.NewBech32Address(bech32Addr, "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "wrong signature returns error",
			proof:       types.NewProof(bech32PubKey, testutil.SingleSignatureProtoFromHex("74657874"), hex.EncodeToString([]byte(bech32Owner))),
			owner:       bech32Owner,
			addressData: types.NewBech32Address(bech32Addr, "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "wrong Bech32 address returns error",
			proof:       types.NewProof(bech32PubKey, bech32SigData, hex.EncodeToString([]byte(bech32Owner))),
			owner:       bech32Owner,
			addressData: types.NewBech32Address("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn", "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "wrong Base58 address returns error",
			proof:       types.NewProof(base58PubKey, base58SigData, hex.EncodeToString([]byte(base58Owner))),
			owner:       base58Owner,
			addressData: types.NewBase58Address("HWQ14mk82aRMAad2TdxFHbeqLeUGo5SiBxTXyZyTesJT"),
			shouldErr:   true,
		},
		{
			name:        "wrong Hex address returns error",
			proof:       types.NewProof(hexPubKey, hexSigData, hex.EncodeToString([]byte(hexOwner))),
			owner:       hexOwner,
			addressData: types.NewHexAddress("0xcdAFfbFd8c131464fEE561e3d9b585141e403719", "0x"),
			shouldErr:   true,
		},
		{
			name:        "invalid multi sig pubkey returns error",
			proof:       types.Proof{PubKey: invalidAny, Signature: validMultisigDataAny, PlainText: hex.EncodeToString([]byte(multiSigOwner))},
			owner:       multiSigOwner,
			addressData: types.NewBech32Address("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn", "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "wrong multi sig address returns error",
			proof:       types.NewProof(multiSigPubKey, multiSigData, hex.EncodeToString([]byte(multiSigOwner))),
			owner:       multiSigOwner,
			addressData: types.NewBech32Address("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn", "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "wrong multi sig pubkey returns error",
			proof:       types.NewProof(bech32PubKey, multiSigData, hex.EncodeToString([]byte(multiSigOwner))),
			owner:       multiSigOwner,
			addressData: types.NewBech32Address("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn", "cosmos"),
			shouldErr:   true,
		},
		{
			name:        "correct proof with Base58 address returns no error",
			proof:       types.NewProof(base58PubKey, base58SigData, hex.EncodeToString([]byte(base58Owner))),
			owner:       base58Owner,
			addressData: types.NewBase58Address(base58Addr),
			shouldErr:   false,
		},
		{
			name:        "correct proof with Bech32 address returns no error",
			proof:       types.NewProof(bech32PubKey, bech32SigData, hex.EncodeToString([]byte(bech32Owner))),
			owner:       bech32Owner,
			addressData: types.NewBech32Address(bech32Addr, "cosmos"),
			shouldErr:   false,
		},
		{
			name:        "correct proof with Hex address returns no error",
			proof:       types.NewProof(hexPubKey, hexSigData, hex.EncodeToString([]byte(hexOwner))),
			owner:       hexOwner,
			addressData: types.NewHexAddress(hexAddr, "0x"),
			shouldErr:   false,
		},
		{
			name:        "correct proof with multisig address returns no error",
			proof:       types.NewProof(multiSigPubKey, multiSigData, hex.EncodeToString([]byte(multiSigOwner))),
			owner:       multiSigOwner,
			addressData: types.NewBech32Address(multisigAddr, "cosmos"),
			shouldErr:   false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cdc, legacyAmino := app.MakeCodecs()
			err := tc.proof.Verify(cdc, legacyAmino, tc.owner, tc.addressData)

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestIsValidTextSig(t *testing.T) {
	testCases := []struct {
		name          string
		value         []byte
		expectedValue string
		expValid      bool
	}{
		{
			name:          "wrong value returns false",
			value:         []byte(""),
			expectedValue: "value",
			expValid:      false,
		},
		{
			name:          "correct value returns true",
			value:         []byte("value"),
			expectedValue: "value",
			expValid:      true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			isValid := types.IsValidTextSig(tc.value, tc.expectedValue)
			require.Equal(t, tc.expValid, isValid)
		})
	}
}

func TestIsValidDirectTxSig(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	testCases := []struct {
		name         string
		value        []byte
		expectedMemo string
		expIsValid   bool
	}{
		{
			name:         "invalid message returns false",
			value:        cdc.MustMarshal(&types.Bech32Address{Prefix: "cosmos"}),
			expectedMemo: "memo",
			expIsValid:   false,
		},
		{
			name:         "wrong memo returns false",
			value:        cdc.MustMarshal(&tx.SignDoc{BodyBytes: cdc.MustMarshal(&tx.TxBody{Memo: "memo"})}),
			expectedMemo: "other memo",
			expIsValid:   false,
		},
		{
			name:         "valid data returns true",
			value:        cdc.MustMarshal(&tx.SignDoc{BodyBytes: cdc.MustMarshal(&tx.TxBody{Memo: "memo"})}),
			expectedMemo: "memo",
			expIsValid:   true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			isValid := types.IsValidDirectTxSig(tc.value, tc.expectedMemo, cdc)
			require.Equal(t, tc.expIsValid, isValid)
		})
	}
}

func TestIsValidAminoTxSig(t *testing.T) {
	_, legacyAmino := app.MakeCodecs()
	testCases := []struct {
		name         string
		value        []byte
		expectedMemo string
		expIsValid   bool
	}{
		{
			name:         "invalid message returns false",
			value:        legacyAmino.MustMarshalJSON(&types.Bech32Address{}),
			expectedMemo: "memo",
			expIsValid:   false,
		},
		{
			name:         "wrong memo returns false",
			value:        legacyAmino.MustMarshalJSON(&legacytx.StdSignDoc{}),
			expectedMemo: "memo",
			expIsValid:   false,
		},
		{
			name:         "valid data returns true",
			value:        legacyAmino.MustMarshalJSON(&legacytx.StdSignDoc{Memo: "memo"}),
			expectedMemo: "memo",
			expIsValid:   true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			isValid := types.IsValidAminoTxSig(tc.value, tc.expectedMemo, legacyAmino)
			require.Equal(t, tc.expIsValid, isValid)
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
				types.NewProof(secp256k1.GenPrivKey().PubKey(), &types.SingleSignatureData{}, "="),
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
		types.NewProof(pubKey, testutil.SingleSignatureProtoFromHex("74657874"), "plain-text"),
		types.NewChainConfig("cosmos"),
		time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
	)

	cdc, _ := app.MakeCodecs()
	marshalled := types.MustMarshalChainLink(cdc, chainLink)
	unmarshalled := types.MustUnmarshalChainLink(cdc, marshalled)
	require.Equal(t, chainLink, unmarshalled)
}
