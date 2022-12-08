package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewChainConfig allows to build a new ChainConfig instance
func NewChainConfig(name string) ChainConfig {
	return ChainConfig{
		Name: name,
	}
}

// Validate checks the validity of the ChainConfig
func (c ChainConfig) Validate() error {
	if strings.TrimSpace(c.Name) == "" {
		return fmt.Errorf("chain name cannot be empty or blank")
	}
	if c.Name != strings.ToLower(c.Name) {
		return fmt.Errorf("chain name must be lowercase")
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewProof is a constructor function for Proof
//nolint:interfacer
func NewProof(pubKey cryptotypes.PubKey, signature Signature, plainText string) Proof {
	pubKeyAny, err := codectypes.NewAnyWithValue(pubKey)
	if err != nil {
		panic("failed to pack public key to any type")
	}

	signatureAny, err := codectypes.NewAnyWithValue(signature)
	if err != nil {
		panic("failed to pack signature to any type")
	}

	return Proof{
		PubKey:    pubKeyAny,
		Signature: signatureAny,
		PlainText: plainText,
	}
}

// GetSignature returns the Signature associated to this proof
func (p Proof) GetSignature() (Signature, error) {
	sigValue := p.Signature.GetCachedValue()
	if sigValue == nil {
		return nil, fmt.Errorf("nil signature")
	}

	signature, ok := sigValue.(Signature)
	if !ok {
		return nil, fmt.Errorf("invalid signature type: %T", sigValue)
	}
	return signature, nil
}

// Validate checks the validity of the Proof
func (p Proof) Validate() error {
	if p.PubKey == nil {
		return fmt.Errorf("public key field cannot be nil")
	}

	if p.Signature == nil {
		return fmt.Errorf("signature field cannot be nil")
	}

	if strings.TrimSpace(p.PlainText) == "" {
		return fmt.Errorf("plain text cannot be empty or blank")
	}

	_, err := hex.DecodeString(p.PlainText)
	if err != nil {
		return fmt.Errorf("invalid hex-encoded plain text")
	}

	return nil
}

// Verify verifies the signature using the given plain text and public key.
// It returns and error if something is invalid.
func (p Proof) Verify(cdc codec.BinaryCodec, amino *codec.LegacyAmino, owner string, address Address) error {
	// Decode the value
	value, err := hex.DecodeString(p.PlainText)
	if err != nil {
		return fmt.Errorf("error while decoding proof text: %s", err)
	}

	// Get the signature
	signature, err := p.GetSignature()
	if err != nil {
		return err
	}

	// Validate the signature
	err = signature.Validate(cdc, amino, value, owner)
	if err != nil {
		return fmt.Errorf("invalid signature: %s", err)
	}

	// Verify the signature
	pubKey, err := signature.Verify(cdc, p.PubKey, value)
	if err != nil {
		return fmt.Errorf("error while verifying the signature: %s", err)
	}

	// Verify the public key
	err = address.VerifyPubKey(pubKey)
	if err != nil {
		return err
	}

	return nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (p *Proof) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pubKey cryptotypes.PubKey
	err := unpacker.UnpackAny(p.PubKey, &pubKey)
	if err != nil {
		return err
	}

	var signatureData Signature
	return unpacker.UnpackAny(p.Signature, &signatureData)
}

// --------------------------------------------------------------------------------------------------------------------

// Signature represents a generic signature data
type Signature interface {
	proto.Message

	// GetValueType returns the type of the signature
	GetValueType() (SignatureValueType, error)

	// Validate checks the validity of the Signature
	Validate(cdc codec.BinaryCodec, amino *codec.LegacyAmino, plainText []byte, owner string) error

	// Verify allows to verify this signature using the given public key against the given plain text.
	// If the signature is valid, it returns the public key instance used to verify it
	Verify(cdc codec.BinaryCodec, pubKey *codectypes.Any, plainText []byte) (cryptotypes.PubKey, error)
}

// --------------------------------------------------------------------------------------------------------------------

// ValidateRawValue tells whether the given value has been properly encoded as a raw value
func ValidateRawValue(value []byte, expectedValue string) error {
	if string(value) != expectedValue {
		return fmt.Errorf("invalid signed value: expected %s, got %s", expectedValue, value)
	}

	return nil
}

// ValidateDirectTxValue tells whether the given value has been properly encoded as a Protobuf transaction containing
// the expected value as the memo field value
func ValidateDirectTxValue(value []byte, expectedMemo string, cdc codec.BinaryCodec) error {
	// Unmarshal the SignDoc
	var signDoc tx.SignDoc
	err := cdc.Unmarshal(value, &signDoc)
	if err != nil {
		return err
	}

	// Check to make sure the value was a SignDoc. If that's not the case, the two arrays will not match
	if !bytes.Equal(value, cdc.MustMarshal(&signDoc)) {
		return fmt.Errorf("invalid signed doc")
	}

	// Get the TxBody
	var txBody tx.TxBody
	err = cdc.Unmarshal(signDoc.BodyBytes, &txBody)
	if err != nil {
		return err
	}

	// Check the memo field
	if txBody.Memo != expectedMemo {
		return fmt.Errorf("invalid signed memo: expected %s, got %s", expectedMemo, txBody.Memo)
	}

	return nil
}

// ValidateAminoTxValue tells whether the given value has been properly encoded as an Amino transaction containing
// the expected value as the memo field value
func ValidateAminoTxValue(value []byte, expectedMemo string, cdc *codec.LegacyAmino) error {
	// Unmarshal the StdSignDoc
	var signDoc legacytx.StdSignDoc
	err := cdc.UnmarshalJSON(value, &signDoc)
	if err != nil {
		return err
	}

	// Check the memo field
	if signDoc.Memo != expectedMemo {
		return fmt.Errorf("invalid signed memo: expected %s, got %s", expectedMemo, signDoc.Memo)
	}

	return nil
}

// ValidatePersonalSignValue tells whether the given value has been properly encoded using the EVM persona_sign specification
func ValidatePersonalSignValue(value []byte, expectedValue string) error {
	expectedSignedValue := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(expectedValue), expectedValue)
	if string(value) != expectedSignedValue {
		return fmt.Errorf("invalid signed value: expected %s but got %s", expectedSignedValue, value)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

var _ Signature = &SingleSignature{}

// NewSingleSignature returns a new CosmosSignature instance
func NewSingleSignature(valueType SignatureValueType, signature []byte) *SingleSignature {
	return &SingleSignature{
		ValueType: valueType,
		Signature: signature,
	}
}

// GetValueType implements CosmosSignature
func (s *SingleSignature) GetValueType() (SignatureValueType, error) {
	return s.ValueType, nil
}

// Validate implements Signature
func (s *SingleSignature) Validate(cdc codec.BinaryCodec, amino *codec.LegacyAmino, plainText []byte, owner string) error {
	// Validate the signature itself
	switch s.ValueType {
	case SIGNATURE_VALUE_TYPE_COSMOS_DIRECT:
		return ValidateDirectTxValue(plainText, owner, cdc)
	case SIGNATURE_VALUE_TYPE_COSMOS_AMINO:
		return ValidateAminoTxValue(plainText, owner, amino)
	case SIGNATURE_VALUE_TYPE_RAW:
		return ValidateRawValue(plainText, owner)
	case SIGNATURE_VALUE_TYPE_EVM_PERSONAL_SIGN:
		return ValidatePersonalSignValue(plainText, owner)
	default:
		return fmt.Errorf("invalid signature type: %s", s.ValueType)
	}
}

// Verify implements Signature
func (s *SingleSignature) Verify(cdc codec.BinaryCodec, pubKey *codectypes.Any, plainText []byte) (cryptotypes.PubKey, error) {
	// Get the pub key
	var pubkey cryptotypes.PubKey
	err := cdc.UnpackAny(pubKey, &pubkey)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack the public key")
	}

	// Verify the signature
	if !pubkey.VerifySignature(plainText, s.Signature) {
		return nil, fmt.Errorf("failed to verify the signature")
	}

	return pubkey, nil
}

// --------------------------------------------------------------------------------------------------------------------

var _ Signature = &CosmosMultiSignature{}

// NewCosmosMultiSignature returns a new CosmosMultiSignature instance
func NewCosmosMultiSignature(bitArray *cryptotypes.CompactBitArray, signatures []Signature) *CosmosMultiSignature {
	sigsAnys := make([]*codectypes.Any, len(signatures))
	for i, sig := range signatures {
		sigAny, err := codectypes.NewAnyWithValue(sig)
		if err != nil {
			panic("failed to pack signature to any type")
		}
		sigsAnys[i] = sigAny
	}

	return &CosmosMultiSignature{
		BitArray:   bitArray,
		Signatures: sigsAnys,
	}
}

// GetValueType implements CosmosSignature
func (s *CosmosMultiSignature) GetValueType() (SignatureValueType, error) {
	signMode := SIGNATURE_VALUE_TYPE_UNSPECIFIED
	for i, signature := range s.Signatures {
		// Unwrap the signature
		cosmosSig, ok := signature.GetCachedValue().(Signature)
		if !ok {
			return SIGNATURE_VALUE_TYPE_UNSPECIFIED, fmt.Errorf("invalid signature type at index %d: %T", i, cosmosSig)
		}

		// Get the signature sign mode
		signatureSignMode, err := cosmosSig.GetValueType()
		if err != nil {
			return SIGNATURE_VALUE_TYPE_UNSPECIFIED, err
		}

		if signatureSignMode == SIGNATURE_VALUE_TYPE_UNSPECIFIED {
			return SIGNATURE_VALUE_TYPE_UNSPECIFIED, fmt.Errorf("invalid signature signing mode: %s", signatureSignMode)
		}

		if signMode != SIGNATURE_VALUE_TYPE_UNSPECIFIED && signMode != signatureSignMode {
			return SIGNATURE_VALUE_TYPE_UNSPECIFIED, fmt.Errorf("signature at index %d has different signing mode than others", i)
		}
		signMode = signatureSignMode
	}
	return signMode, nil
}

// Validate implements Signature
func (s *CosmosMultiSignature) Validate(cdc codec.BinaryCodec, amino *codec.LegacyAmino, plainText []byte, owner string) error {
	signMode, err := s.GetValueType()
	if err != nil {
		return err
	}

	// Validate the signature itself
	switch signMode {
	case SIGNATURE_VALUE_TYPE_RAW:
		return ValidateRawValue(plainText, owner)
	case SIGNATURE_VALUE_TYPE_COSMOS_DIRECT:
		return ValidateDirectTxValue(plainText, owner, cdc)
	case SIGNATURE_VALUE_TYPE_COSMOS_AMINO:
		return ValidateAminoTxValue(plainText, owner, amino)
	default:
		return fmt.Errorf("invalid signing mode: %s", signMode)
	}
}

// Verify implements Signature
func (s *CosmosMultiSignature) Verify(cdc codec.BinaryCodec, pubKey *codectypes.Any, plainText []byte) (cryptotypes.PubKey, error) {
	// Convert the signature data to the Cosmos type
	cosmosSigData, err := SignatureToCosmosSignatureData(cdc, s)
	if err != nil {
		return nil, err
	}

	// Make sure the sig data is of the correct type
	sigData, ok := cosmosSigData.(*signing.MultiSignatureData)
	if !ok {
		return nil, fmt.Errorf("invalid signature data type: %T", sigData)
	}

	// Get the pub key
	var multiPubkey multisig.PubKey
	err = cdc.UnpackAny(pubKey, &multiPubkey)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack the public key")
	}

	// Verify the signature
	err = multiPubkey.VerifyMultisignature(func(mode signing.SignMode) ([]byte, error) {
		return plainText, nil
	}, sigData)
	if err != nil {
		return nil, err
	}

	return multiPubkey, nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (s *CosmosMultiSignature) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, signature := range s.Signatures {
		var signatureData Signature
		err := unpacker.UnpackAny(signature, &signatureData)
		if err != nil {
			return err
		}
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// SignatureToCosmosSignatureData allows to convert the given Signature to a Cosmos SignatureData instance
// unpacking the proto.Any instance using the given unpacker.
func SignatureToCosmosSignatureData(unpacker codectypes.AnyUnpacker, s Signature) (signing.SignatureData, error) {
	switch data := s.(type) {
	case *SingleSignature:
		return &signing.SingleSignatureData{
			Signature: data.Signature,
			SignMode:  signing.SignMode_SIGN_MODE_UNSPECIFIED, // This can be unknown since we don't use it anyway
		}, nil

	case *CosmosMultiSignature:
		sigs, err := unpackSignatures(unpacker, data.Signatures)
		if err != nil {
			return nil, err
		}

		return &signing.MultiSignatureData{
			BitArray:   data.BitArray,
			Signatures: sigs,
		}, nil

	default:
		return nil, fmt.Errorf("signature type not supported: %T", s)
	}
}

// CosmosSignatureDataToSignature allows to create a Signature instance from the given Cosmos signature data
func CosmosSignatureDataToSignature(data signing.SignatureData) (Signature, error) {
	switch data := data.(type) {
	case *signing.SingleSignatureData:
		var signatureType SignatureValueType
		switch data.SignMode {
		case signing.SignMode_SIGN_MODE_DIRECT:
			signatureType = SIGNATURE_VALUE_TYPE_COSMOS_DIRECT
		case signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON:
			signatureType = SIGNATURE_VALUE_TYPE_COSMOS_AMINO
		case signing.SignMode_SIGN_MODE_TEXTUAL:
			signatureType = SIGNATURE_VALUE_TYPE_RAW
		default:
			return nil, fmt.Errorf("unsupported signing mode: %s", data.SignMode)
		}

		return &SingleSignature{
			ValueType: signatureType,
			Signature: data.Signature,
		}, nil

	case *signing.MultiSignatureData:
		sigAnys := make([]*codectypes.Any, len(data.Signatures))
		for i, data := range data.Signatures {
			sigData, err := CosmosSignatureDataToSignature(data)
			if err != nil {
				return nil, err
			}
			sigAny, err := codectypes.NewAnyWithValue(sigData)
			if err != nil {
				return nil, err
			}
			sigAnys[i] = sigAny
		}
		return &CosmosMultiSignature{
			BitArray:   data.BitArray,
			Signatures: sigAnys,
		}, nil
	default:
		return nil, fmt.Errorf("unexpected case %+v", data)
	}
}

// unpackSignatures unpacks the given signatures using the provided unpacker
func unpackSignatures(unpacker codectypes.AnyUnpacker, sigs []*codectypes.Any) ([]signing.SignatureData, error) {
	var signatures = make([]signing.SignatureData, len(sigs))
	for i, sig := range sigs {
		var signature Signature
		if err := unpacker.UnpackAny(sig, &signature); err != nil {
			return nil, err
		}

		cosmosSigData, err := SignatureToCosmosSignatureData(unpacker, signature)
		if err != nil {
			return nil, err
		}
		signatures[i] = cosmosSigData
	}

	return signatures, nil
}

// --------------------------------------------------------------------------------------------------------------------

func NewAddress(value string, generationAlgorithm GenerationAlgorithm, encodingAlgorithm AddressEncoding) Address {
	encodingAlgorithmAny, err := codectypes.NewAnyWithValue(encodingAlgorithm)
	if err != nil {
		panic("failed to pack encoding algorithm to any type")
	}
	return Address{
		Value:               value,
		GenerationAlgorithm: generationAlgorithm,
		EncodingAlgorithm:   encodingAlgorithmAny,
	}
}

// Validate checks the validity of the Address
func (a Address) Validate() error {
	if strings.TrimSpace(a.Value) == "" {
		return fmt.Errorf("value cannot be empty or blank")
	}
	if a.GenerationAlgorithm == GENERATION_ALGORITHM_UNSPECIFIED {
		return fmt.Errorf("unknown address generation algorithm")
	}
	if a.EncodingAlgorithm == nil {
		return fmt.Errorf("address encoding algorithm field can not be nil")
	}
	return nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (a *Address) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var encoding AddressEncoding
	return unpacker.UnpackAny(a.EncodingAlgorithm, &encoding)
}

// VerifyPubKey verifies if the address is generated in the proper algorithms from the public key
func (a *Address) VerifyPubKey(pubKey cryptotypes.PubKey) error {
	generatedBz, err := generateAddressBytes(pubKey, a.GenerationAlgorithm)
	if err != nil {
		return err
	}
	encoded, err := a.EncodingAlgorithm.GetCachedValue().(AddressEncoding).Encode(generatedBz)
	if a.Value != encoded {
		return fmt.Errorf("address bytes do not match generated ones: expected %s but got %s", a.Value, encoded)
	}
	return nil
}

// generateAddressBytes generates the address bytes starting from the given public key
// and using the provided generation algorithm
func generateAddressBytes(key cryptotypes.PubKey, generationAlgorithm GenerationAlgorithm) ([]byte, error) {
	switch generationAlgorithm {
	case GENERATION_ALGORITHM_COSMOS:
		return key.Address().Bytes(), nil

	case GENERATION_ALGORITHM_DO_NOTHING:
		return key.Bytes(), nil

	case GENERATION_ALGORITHM_EVM:
		pubKey, err := btcec.ParsePubKey(key.Bytes(), btcec.S256())
		if err != nil {
			return nil, err
		}
		uncompressedPubKey := pubKey.SerializeUncompressed()
		return crypto.Keccak256(uncompressedPubKey[1:])[12:], nil

	default:
		return nil, fmt.Errorf("unsupported generation algorithm")
	}
}

// --------------------------------------------------------------------------------------------------------------------

type AddressEncoding interface {
	proto.Message

	// Validate checks the validity of the AddressEncoding
	Validate() error

	// Encode encodes the address bytes into a proper address string by the encoding algorithm
	Encode(value []byte) (string, error)
}

// --------------------------------------------------------------------------------------------------------------------

var _ AddressEncoding = &Bech32Encoding{}

// NeBech32Encoding returns a new Bech32Encoding instance
func NewBech32Encoding(prefix string) *Bech32Encoding {
	return &Bech32Encoding{Prefix: prefix}
}

// Validate implements AddressEncoding
func (b Bech32Encoding) Validate() error {
	if strings.TrimSpace(b.Prefix) == "" {
		return fmt.Errorf("prefix cannot be empty or blank")
	}
	return nil
}

// Encode implements AddressEncoding
func (b *Bech32Encoding) Encode(value []byte) (string, error) {
	return bech32.ConvertAndEncode(b.Prefix, value)
}

// --------------------------------------------------------------------------------------------------------------------

var _ AddressEncoding = &Base58Encoding{}

// NewBase58Encoding returns a new Base58Encoding instance
func NewBase58Encoding(prefix string) *Base58Encoding {
	return &Base58Encoding{Prefix: prefix}
}

// Validate implements AddressEncoding
func (b Base58Encoding) Validate() error {
	if len(b.Prefix) != 0 && strings.TrimSpace(b.Prefix) == "" {
		return fmt.Errorf("prefix cannot be blank")
	}
	_, err := hex.DecodeString(b.Prefix)
	if err != nil {
		return err
	}
	return nil
}

// Encode implements AddressEncoding
func (b *Base58Encoding) Encode(value []byte) (string, error) {
	prefixBz, err := hex.DecodeString(b.Prefix)
	if err != nil {
		return "", err
	}
	return base58.Encode(append(prefixBz, value...)), nil
}

// --------------------------------------------------------------------------------------------------------------------

var _ AddressEncoding = &HexEncoding{}

// NewHexEncoding returns a new HexEncoding instance
func NewHexEncoding(prefix string, isEIP55 bool) *HexEncoding {
	return &HexEncoding{Prefix: prefix, IsEIP55: isEIP55}
}

// Validate implements AddressEncoding
func (h HexEncoding) Validate() error {
	if len(h.Prefix) != 0 && strings.TrimSpace(h.Prefix) == "" {
		return fmt.Errorf("prefix cannot be blank")
	}
	return nil
}

// GetValue implements AddressEncoding
func (h HexEncoding) Encode(value []byte) (string, error) {
	hexAddr := hex.EncodeToString(value)
	if h.IsEIP55 {
		return h.Prefix + common.HexToAddress(hexAddr).Hex()[2:], nil
	}
	return h.Prefix + hexAddr, nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewChainLink returns a new ChainLink instance
//nolint:interfacer
func NewChainLink(user string, address Address, proof Proof, chainConfig ChainConfig, creationTime time.Time) ChainLink {
	return ChainLink{
		User:         user,
		Address:      address,
		Proof:        proof,
		ChainConfig:  chainConfig,
		CreationTime: creationTime,
	}
}

// Validate checks the validity of the ChainLink
func (link ChainLink) Validate() error {
	if _, err := sdk.AccAddressFromBech32(link.User); err != nil {
		return fmt.Errorf("invalid creator address: %s", link.User)
	}

	err := link.Address.Validate()
	if err != nil {
		return err
	}

	err = link.Proof.Validate()
	if err != nil {
		return err
	}

	err = link.ChainConfig.Validate()
	if err != nil {
		return err
	}

	if link.CreationTime.IsZero() {
		return fmt.Errorf("creation time cannot be zero")
	}

	return nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (link *ChainLink) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	err := link.Address.UnpackInterfaces(unpacker)
	if err != nil {
		return err
	}

	err = link.Proof.UnpackInterfaces(unpacker)
	if err != nil {
		return err
	}

	return nil
}

// MustMarshalChainLink serializes the given chain link using the provided BinaryCodec
func MustMarshalChainLink(cdc codec.BinaryCodec, link ChainLink) []byte {
	return cdc.MustMarshal(&link)
}

// MustUnmarshalChainLink deserializes the given byte array as a chain link using
// the provided BinaryCodec
func MustUnmarshalChainLink(codec codec.BinaryCodec, bz []byte) ChainLink {
	var link ChainLink
	codec.MustUnmarshal(bz, &link)
	return link
}
