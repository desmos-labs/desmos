package types_test

import (
	"encoding/hex"

	"github.com/desmos-labs/desmos/v4/types/crypto/ethsecp256k1"

	"github.com/desmos-labs/desmos/v4/testutil/profilestesting"

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

	"github.com/desmos-labs/desmos/v4/app"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
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

func TestProof_GetSignature(t *testing.T) {
	pubKeyAny, err := codectypes.NewAnyWithValue(secp256k1.GenPrivKey().PubKey())
	require.NoError(t, err)

	testCases := []struct {
		name         string
		proof        types.Proof
		shouldErr    bool
		expSignature types.Signature
	}{
		{
			name: "invalid signature returns error",
			proof: types.Proof{
				PubKey:    pubKeyAny,
				Signature: pubKeyAny,
				PlainText: "74657874",
			},
			shouldErr: true,
		},
		{
			name: "valid signature is returned properly",
			proof: types.NewProof(
				secp256k1.GenPrivKey().PubKey(),
				profilestesting.SingleCosmosSignatureFromHex("74657874"),
				"74657874",
			),
			expSignature: profilestesting.SingleCosmosSignatureFromHex("74657874"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			signature, err := tc.proof.GetSignature()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expSignature, signature)
			}
		})
	}
}

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
			proof:     types.NewProof(secp256k1.GenPrivKey().PubKey(), profilestesting.SingleCosmosSignatureFromHex("74657874"), ""),
			shouldErr: true,
		},
		{
			name:      "invalid plain text format returns error",
			proof:     types.NewProof(secp256k1.GenPrivKey().PubKey(), profilestesting.SingleCosmosSignatureFromHex("74657874"), "="),
			shouldErr: true,
		},
		{
			name: "valid proof returns no error",
			proof: types.NewProof(
				secp256k1.GenPrivKey().PubKey(),
				profilestesting.SingleCosmosSignatureFromHex("74657874"),
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

// generateCosmosMultiSigSignature uses the given private keys to sign the given message using the multi sig algorithm,
// and returns the obtained MultiSignatureData instance.
func generateCosmosMultiSigSignature(t *testing.T, privKeys []cryptotypes.PrivKey, msg []byte) *types.CosmosMultiSignature {
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
		sigData := &signing.SingleSignatureData{SignMode: signing.SignMode_SIGN_MODE_TEXTUAL, Signature: sig}
		err = multisig.AddSignatureFromPubKey(cosmosMultisig, sigData, privKeys[i].PubKey(), pubKeys)
		require.NoError(t, err)
	}

	// Generate the signature data object
	signature, err := types.SignatureDataFromCosmosSignatureData(cosmosMultisig)
	require.NoError(t, err)
	return signature.(*types.CosmosMultiSignature)
}

func TestProof_Verify(t *testing.T) {
	privKeyBz, err := hex.DecodeString("bb98111da675930d32f79451fa8d05f188289699558c17148a5d9c82cdb31d1fe04fb0a0d9e689b436b59eff9676d7f2d788244cc4ccfc5768fe117efbd0f9d3")
	require.NoError(t, err)
	privKey := ed25519.PrivKey{Key: privKeyBz}

	pubKey := privKey.PubKey()
	pubKeyAny, err := codectypes.NewAnyWithValue(pubKey)
	require.NoError(t, err)

	anotherPubKey := secp256k1.GenPrivKey().PubKey()
	anotherPubKeyAny, err := codectypes.NewAnyWithValue(anotherPubKey)
	require.NoError(t, err)

	sigBz, err := privKey.Sign([]byte("cosmos1u55ywhk6thmhnxs7yn8vh8v7eznckcqjevnadx"))
	require.NoError(t, err)
	signature := types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, sigBz)
	sigAny, err := codectypes.NewAnyWithValue(signature)
	require.NoError(t, err)

	invalidAny, err := codectypes.NewAnyWithValue(&types.CosmosSingleSignature{})
	require.NoError(t, err)

	testCases := []struct {
		name        string
		proof       types.Proof
		owner       string
		addressData types.AddressData
		shouldErr   bool
	}{
		{
			name:        "invalid value returns error",
			proof:       types.Proof{PubKey: pubKeyAny, Signature: sigAny, PlainText: "value"},
			owner:       "cosmos1u55ywhk6thmhnxs7yn8vh8v7eznckcqjevnadx",
			addressData: types.NewBase58Address(base58.Encode(pubKey.Bytes())),
			shouldErr:   true,
		},
		{
			name:        "incorrect signature type returns error",
			proof:       types.Proof{PubKey: pubKeyAny, Signature: invalidAny, PlainText: "76616C7565"},
			owner:       "cosmos1u55ywhk6thmhnxs7yn8vh8v7eznckcqjevnadx",
			addressData: types.NewBase58Address(base58.Encode(pubKey.Bytes())),
			shouldErr:   true,
		},
		{
			name:        "invalid signature returns error",
			proof:       types.Proof{PubKey: invalidAny, Signature: sigAny, PlainText: "76616C7565"},
			owner:       "cosmos1u55ywhk6thmhnxs7yn8vh8v7eznckcqjevnadx",
			addressData: types.NewBase58Address(base58.Encode(pubKey.Bytes())),
			shouldErr:   true,
		},
		{
			name: "wrong signature returns error",
			proof: types.Proof{
				PubKey:    anotherPubKeyAny,
				Signature: sigAny,
				PlainText: "636F736D6F73317535357977686B3674686D686E787337796E387668387637657A6E636B63716A65766E616478",
			},
			owner:       "cosmos1u55ywhk6thmhnxs7yn8vh8v7eznckcqjevnadx",
			addressData: types.NewBase58Address(base58.Encode(pubKey.Bytes())),
			shouldErr:   true,
		},
		{
			name: "wrong signature returns error",
			proof: types.Proof{
				PubKey:    pubKeyAny,
				Signature: sigAny,
				PlainText: "636F736D6F73317535357977686B3674686D686E787337796E387668387637657A6E636B63716A65766E616478",
			},
			owner:       "cosmos1u55ywhk6thmhnxs7yn8vh8v7eznckcqjevnadx",
			addressData: types.NewBase58Address(base58.Encode(anotherPubKey.Bytes())),
			shouldErr:   true,
		},
		{
			name: "valid data returns no error",
			proof: types.Proof{
				PubKey:    pubKeyAny,
				Signature: sigAny,
				PlainText: "636F736D6F73317535357977686B3674686D686E787337796E387668387637657A6E636B63716A65766E616478",
			},
			owner:       "cosmos1u55ywhk6thmhnxs7yn8vh8v7eznckcqjevnadx",
			addressData: types.NewBase58Address(base58.Encode(pubKey.Bytes())),
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

// --------------------------------------------------------------------------------------------------------------------

func TestCosmosSingleSignature_Validate(t *testing.T) {
	cdc, amino := app.MakeCodecs()
	testCases := []struct {
		name      string
		signature *types.CosmosSingleSignature
		plainText []byte
		owner     string
		shouldErr bool
	}{
		{
			name:      "invalid direct value returns error",
			signature: types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_DIRECT, nil),
			plainText: cdc.MustMarshal(&types.Bech32Address{Prefix: "cosmos"}),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: true,
		},
		{
			name:      "valid direct value returns no error",
			signature: types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_DIRECT, nil),
			plainText: cdc.MustMarshal(&tx.SignDoc{BodyBytes: cdc.MustMarshal(&tx.TxBody{Memo: "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae"})}),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: false,
		},
		{
			name:      "invalid amino value returns error",
			signature: types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_AMINO, nil),
			plainText: amino.MustMarshalJSON(&legacytx.StdSignDoc{Memo: ""}),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: true,
		},
		{
			name:      "valid amino value returns error",
			signature: types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_AMINO, nil),
			plainText: amino.MustMarshalJSON(&legacytx.StdSignDoc{Memo: "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae"}),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: false,
		},
		{
			name:      "invalid raw value returns error",
			signature: types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, nil),
			plainText: []byte(""),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: true,
		},
		{
			name:      "valid raw value returns no error",
			signature: types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, nil),
			plainText: []byte("cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae"),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.signature.Validate(cdc, amino, tc.plainText, tc.owner)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCosmosSingleSignature_Verify(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	pubKeyAny, err := codectypes.NewAnyWithValue(pubKey)
	require.NoError(t, err)

	sigBz, err := privKey.Sign([]byte("cosmos10m20h8fy0qp2a8f46zzjpvg8pfl8flajgxsvmk"))
	require.NoError(t, err)

	invalidAny, err := codectypes.NewAnyWithValue(privKey)
	require.NoError(t, err)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		signature *types.CosmosSingleSignature
		pubKey    *codectypes.Any
		plainText []byte
		shouldErr bool
		expPubKey cryptotypes.PubKey
	}{
		{
			name:      "invalid pub key returns error",
			signature: types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, sigBz),
			pubKey:    invalidAny,
			plainText: []byte("cosmos10m20h8fy0qp2a8f46zzjpvg8pfl8flajgxsvmk"),
			shouldErr: true,
		},
		{
			name:      "invalid signature returns error",
			signature: types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, sigBz),
			pubKey:    pubKeyAny,
			plainText: []byte("value"),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			signature: types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, sigBz),
			pubKey:    pubKeyAny,
			plainText: []byte("cosmos10m20h8fy0qp2a8f46zzjpvg8pfl8flajgxsvmk"),
			shouldErr: false,
			expPubKey: pubKey,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			pubKey, err := tc.signature.Verify(cdc, tc.pubKey, tc.plainText)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.True(t, tc.expPubKey.Equals(pubKey))
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestCosmosMultiSignature_GetSignMode(t *testing.T) {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		signature   *types.CosmosMultiSignature
		shouldErr   bool
		expSignMode types.CosmosSignMode
	}{
		{
			name: "different sign modes return error",
			signature: types.NewCosmosMultiSignature(nil, []types.CosmosSignature{
				types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_DIRECT, nil),
				types.NewCosmosMultiSignature(nil, []types.CosmosSignature{
					types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, nil),
					types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_DIRECT, nil),
				}),
			}),
			shouldErr: true,
		},
		{
			name: "correct sign modes return proper value",
			signature: types.NewCosmosMultiSignature(nil, []types.CosmosSignature{
				types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, nil),
				types.NewCosmosMultiSignature(nil, []types.CosmosSignature{
					types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, nil),
					types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, nil),
				}),
			}),
			shouldErr:   false,
			expSignMode: types.COSMOS_SIGN_MODE_RAW,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			signMode, err := tc.signature.GetSignMode()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expSignMode, signMode)
			}
		})
	}
}

func TestCosmosMultiSignature_Validate(t *testing.T) {
	cdc, amino := app.MakeCodecs()
	testCases := []struct {
		name      string
		signature *types.CosmosMultiSignature
		plainText []byte
		owner     string
		shouldErr bool
	}{
		{
			name: "invalid direct value returns error",
			signature: types.NewCosmosMultiSignature(nil, []types.CosmosSignature{
				types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_DIRECT, nil),
			}),
			plainText: cdc.MustMarshal(&types.Bech32Address{Prefix: "cosmos"}),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: true,
		},
		{
			name: "valid direct value returns no error",
			signature: types.NewCosmosMultiSignature(nil, []types.CosmosSignature{
				types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_DIRECT, nil),
			}),
			plainText: cdc.MustMarshal(&tx.SignDoc{BodyBytes: cdc.MustMarshal(&tx.TxBody{Memo: "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae"})}),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: false,
		},
		{
			name: "invalid amino value returns error",
			signature: types.NewCosmosMultiSignature(nil, []types.CosmosSignature{
				types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_AMINO, nil),
			}),
			plainText: amino.MustMarshalJSON(&legacytx.StdSignDoc{Memo: ""}),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: true,
		},
		{
			name: "valid amino value returns error",
			signature: types.NewCosmosMultiSignature(nil, []types.CosmosSignature{
				types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_AMINO, nil),
			}),
			plainText: amino.MustMarshalJSON(&legacytx.StdSignDoc{Memo: "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae"}),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: false,
		},
		{
			name: "invalid raw value returns error",
			signature: types.NewCosmosMultiSignature(nil, []types.CosmosSignature{
				types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, nil),
			}),
			plainText: []byte(""),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: true,
		},
		{
			name: "valid raw value returns no error",
			signature: types.NewCosmosMultiSignature(nil, []types.CosmosSignature{
				types.NewCosmosSingleSignature(types.COSMOS_SIGN_MODE_RAW, nil),
			}),
			plainText: []byte("cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae"),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.signature.Validate(cdc, amino, tc.plainText, tc.owner)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCosmosMultiSignature_Verify(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	privKeys, pubKey := generateMultiSigKeys(3)
	pubKeyAny, err := codectypes.NewAnyWithValue(pubKey)
	require.NoError(t, err)

	invalidAny, err := codectypes.NewAnyWithValue(&types.CosmosMultiSignature{})
	require.NoError(t, err)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		signature *types.CosmosMultiSignature
		pubKey    *codectypes.Any
		plainText []byte
		shouldErr bool
		expPubKey cryptotypes.PubKey
	}{
		{
			name:      "invalid pub key returns error",
			signature: generateCosmosMultiSigSignature(t, privKeys, []byte("cosmos1vfvst0mr79nzzxsk65uuhfklmwrnsfadhtn977")),
			pubKey:    invalidAny,
			plainText: []byte("cosmos1vfvst0mr79nzzxsk65uuhfklmwrnsfadhtn977"),
			shouldErr: true,
		},
		{
			name:      "invalid signature returns error",
			signature: generateCosmosMultiSigSignature(t, privKeys, []byte("cosmos1vfvst0mr79nzzxsk65uuhfklmwrnsfadhtn977")),
			pubKey:    pubKeyAny,
			plainText: []byte("value"),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			signature: generateCosmosMultiSigSignature(t, privKeys, []byte("cosmos1vfvst0mr79nzzxsk65uuhfklmwrnsfadhtn977")),
			pubKey:    pubKeyAny,
			plainText: []byte("cosmos1vfvst0mr79nzzxsk65uuhfklmwrnsfadhtn977"),
			shouldErr: false,
			expPubKey: pubKey,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			pubKey, err := tc.signature.Verify(cdc, tc.pubKey, tc.plainText)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.True(t, tc.expPubKey.Equals(pubKey))
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestValidateDirectTxSig(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	testCases := []struct {
		name         string
		value        []byte
		expectedMemo string
		shouldErr    bool
	}{
		{
			name:         "invalid message returns error",
			value:        cdc.MustMarshal(&types.Bech32Address{Prefix: "cosmos"}),
			expectedMemo: "memo",
			shouldErr:    true,
		},
		{
			name:         "wrong memo returns error",
			value:        cdc.MustMarshal(&tx.SignDoc{BodyBytes: cdc.MustMarshal(&tx.TxBody{Memo: "memo"})}),
			expectedMemo: "other memo",
			shouldErr:    true,
		},
		{
			name:         "valid data returns no error",
			value:        cdc.MustMarshal(&tx.SignDoc{BodyBytes: cdc.MustMarshal(&tx.TxBody{Memo: "memo"})}),
			expectedMemo: "memo",
			shouldErr:    false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateDirectTxSig(tc.value, tc.expectedMemo, cdc)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateAminoTxSig(t *testing.T) {
	_, legacyAmino := app.MakeCodecs()
	testCases := []struct {
		name         string
		value        []byte
		expectedMemo string
		shouldErr    bool
	}{
		{
			name:         "invalid message returns error",
			value:        legacyAmino.MustMarshalJSON(&types.Bech32Address{}),
			expectedMemo: "memo",
			shouldErr:    true,
		},
		{
			name:         "wrong memo returns error",
			value:        legacyAmino.MustMarshalJSON(&legacytx.StdSignDoc{}),
			expectedMemo: "memo",
			shouldErr:    true,
		},
		{
			name:         "valid data returns no error",
			value:        legacyAmino.MustMarshalJSON(&legacytx.StdSignDoc{Memo: "memo"}),
			expectedMemo: "memo",
			shouldErr:    false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateAminoTxSig(tc.value, tc.expectedMemo, legacyAmino)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateRawSig(t *testing.T) {
	testCases := []struct {
		name          string
		value         []byte
		expectedValue string
		shouldErr     bool
	}{
		{
			name:          "wrong value returns error",
			value:         []byte(""),
			expectedValue: "value",
			shouldErr:     true,
		},
		{
			name:          "correct value returns no error",
			value:         []byte("value"),
			expectedValue: "value",
			shouldErr:     false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateRawSig(tc.value, tc.expectedValue)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestEVMSignature_Validate(t *testing.T) {
	cdc, amino := app.MakeCodecs()
	testCases := []struct {
		name      string
		signature *types.EVMSignature
		plainText []byte
		owner     string
		shouldErr bool
	}{
		{
			name:      "invalid signing method returns error",
			signature: types.NewEVMSignature(types.EVM_SIGNATURE_METHOD_UNSPECIFIED, nil),
			plainText: []byte("\x19Ethereum Signed Message:\n45cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae"),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: true,
		},
		{
			name:      "invalid plain text returns error",
			signature: types.NewEVMSignature(types.EVM_SIGNATURE_METHOD_PERSONAL_SIGN, nil),
			plainText: []byte("\x19Ethereum Signed Message:\n5value"),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			signature: types.NewEVMSignature(types.EVM_SIGNATURE_METHOD_PERSONAL_SIGN, nil),
			plainText: []byte("\x19Ethereum Signed Message:\n45cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae"),
			owner:     "cosmos1s3p4hlhfnlsynauak7ggqv2y4hafwc0y6u0hae",
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.signature.Validate(cdc, amino, tc.plainText, tc.owner)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestEVMSignature_Verify(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	pubKeyHex := "0473b502fc87442aa6e3026ebd78ce4f7d631896f3f3abbe99c63c36fb3bedffce84618b8d83cea35eeaf9d89620c160df4effd72dfd41c7d6b3bb334e11cc990e"
	pubKeyBz, err := hex.DecodeString(pubKeyHex)
	require.NoError(t, err)

	pubKey := &ethsecp256k1.PubKey{Key: pubKeyBz}
	pubKeyAny, err := codectypes.NewAnyWithValue(pubKey)
	require.NoError(t, err)

	sigBz, err := hex.DecodeString("fc036f35f3d1aee08fbb49048f91eb8fc7fc0d83553067a3a70ce25d039e74aa310fd488dd0c493e75ab3923df4cc6b077ee9309df1501583ca330d1c7c1ff6b1b")
	require.NoError(t, err)

	invalidAny, err := codectypes.NewAnyWithValue(&types.CosmosSingleSignature{})
	require.NoError(t, err)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		signature *types.EVMSignature
		pubKey    *codectypes.Any
		plainText []byte
		shouldErr bool
		expPubKey cryptotypes.PubKey
	}{
		{
			name:      "invalid pub key returns error",
			signature: types.NewEVMSignature(types.EVM_SIGNATURE_METHOD_PERSONAL_SIGN, sigBz),
			pubKey:    invalidAny,
			plainText: []byte("\x19Ethereum Signed Message:\n45desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd"),
			shouldErr: true,
		},
		{
			name:      "invalid signature returns error",
			signature: types.NewEVMSignature(types.EVM_SIGNATURE_METHOD_PERSONAL_SIGN, sigBz),
			pubKey:    pubKeyAny,
			plainText: []byte("value"),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			signature: types.NewEVMSignature(types.EVM_SIGNATURE_METHOD_PERSONAL_SIGN, sigBz),
			pubKey:    pubKeyAny,
			plainText: []byte("\x19Ethereum Signed Message:\n45desmos16c60y8t8vra27zjg2arlcd58dck9cwn7p6fwtd"),
			shouldErr: false,
			expPubKey: pubKey,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			pubKey, err := tc.signature.Verify(cdc, tc.pubKey, tc.plainText)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.True(t, tc.expPubKey.Equals(pubKey))
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
		name      string
		address   *codectypes.Any
		shouldErr bool
	}{
		{
			name:      "invalid address returns error",
			address:   profilestesting.NewAny(secp256k1.GenPrivKey()),
			shouldErr: true,
		},
		{
			name:      "valid Bech32 data returns no error",
			address:   profilestesting.NewAny(types.NewBech32Address("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k", "cosmos")),
			shouldErr: false,
		},
		{
			name:      "valid Base58 data returns no error",
			address:   profilestesting.NewAny(types.NewBase58Address("5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN")),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cdc, _ := app.MakeCodecs()
			_, err := types.UnpackAddressData(cdc, tc.address)

			if tc.shouldErr {
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
				Proof:        types.NewProof(secp256k1.GenPrivKey().PubKey(), profilestesting.SingleCosmosSignatureFromHex("74657874"), "74657874"),
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
				types.NewProof(secp256k1.GenPrivKey().PubKey(), profilestesting.SingleCosmosSignatureFromHex("74657874"), "74657874"),
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
				types.NewProof(secp256k1.GenPrivKey().PubKey(), &types.CosmosSingleSignature{}, "="),
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
				types.NewProof(secp256k1.GenPrivKey().PubKey(), profilestesting.SingleCosmosSignatureFromHex("74657874"), "74657874"),
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
				types.NewProof(secp256k1.GenPrivKey().PubKey(), profilestesting.SingleCosmosSignatureFromHex("74657874"), "74657874"),
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
				types.NewProof(secp256k1.GenPrivKey().PubKey(), profilestesting.SingleCosmosSignatureFromHex("74657874"), "74657874"),
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
		types.NewProof(pubKey, profilestesting.SingleCosmosSignatureFromHex("74657874"), "plain-text"),
		types.NewChainConfig("cosmos"),
		time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
	)

	cdc, _ := app.MakeCodecs()
	marshalled := types.MustMarshalChainLink(cdc, chainLink)
	unmarshalled := types.MustUnmarshalChainLink(cdc, marshalled)
	require.Equal(t, chainLink, unmarshalled)
}
