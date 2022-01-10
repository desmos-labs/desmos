package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	signing "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/proto"
	"github.com/mr-tron/base58"

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
// nolint:interfacer
func NewProof(pubKey cryptotypes.PubKey, signature SignatureData, plainText string) Proof {
	pubKeyAny, err := codectypes.NewAnyWithValue(pubKey)
	if err != nil {
		panic("failed to pack public key to any type")
	}

	signatureAny, err := codectypes.NewAnyWithValue(signature)
	if err != nil {
		panic("failed to pack signature data to any type")
	}

	return Proof{
		PubKey:    pubKeyAny,
		Signature: signatureAny,
		PlainText: plainText,
	}
}

// Validate checks the validity of the Proof
func (p Proof) Validate() error {
	if p.PubKey == nil {
		return fmt.Errorf("public key field cannot be nil")
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
func (p Proof) Verify(unpacker codectypes.AnyUnpacker, address AddressData) error {
	value, err := hex.DecodeString(p.PlainText)
	if err != nil {
		return fmt.Errorf("error while decoding proof text: %s", err)
	}

	var sigData SignatureData
	err = unpacker.UnpackAny(p.Signature, &sigData)
	if err != nil {
		return fmt.Errorf("failed to unpack the signature")
	}
	// Convert the signature data to the Cosmos type
	cosmosSigData, err := SignatureDataToCosmosSignatureData(unpacker, sigData)
	if err != nil {
		return err
	}

	// Verify the signature
	var pubkey cryptotypes.PubKey
	switch sigData := cosmosSigData.(type) {
	case *signing.SingleSignatureData:
		err = unpacker.UnpackAny(p.PubKey, &pubkey)
		if err != nil {
			return fmt.Errorf("failed to unpack the public key")
		}
		if !pubkey.VerifySignature(value, sigData.Signature) {
			return fmt.Errorf("failed to verify the signature")
		}

	case *signing.MultiSignatureData:
		var multiPubkey multisig.PubKey
		err = unpacker.UnpackAny(p.PubKey, &multiPubkey)
		if err != nil {
			return fmt.Errorf("failed to unpack the public key")
		}
		err = multiPubkey.VerifyMultisignature(
			func(mode signing.SignMode) ([]byte, error) {
				return value, nil
			},
			sigData,
		)
		if err != nil {
			return err
		}
		pubkey = multiPubkey
	}

	// Verify the public key
	valid, err := address.VerifyPubKey(pubkey)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("invalid address and public key combination provided")
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

	var signatureData SignatureData
	return unpacker.UnpackAny(p.Signature, &signatureData)
}

// --------------------------------------------------------------------------------------------------------------------

// SignatureData represents a generic single- or multi- signature data
type SignatureData interface {
	proto.Message

	isSignatureData()
}

// SignatureDataToCosmosSignatureData allows to convert the given SignatureData to a Cosmos SignatureData instance
// unpacking the proto.Any instance using the given unpacker.
func SignatureDataToCosmosSignatureData(unpacker codectypes.AnyUnpacker, s SignatureData) (signing.SignatureData, error) {
	switch data := s.(type) {
	case *SingleSignatureData:
		return &signing.SingleSignatureData{
			Signature: data.Signature,
			SignMode:  data.Mode,
		}, nil

	case *MultiSignatureData:
		sigs, err := unpackSignatures(unpacker, data.Signatures)
		if err != nil {
			return nil, err
		}

		return &signing.MultiSignatureData{
			BitArray:   data.BitArray,
			Signatures: sigs,
		}, nil
	}

	return nil, fmt.Errorf("signature type not supported: %T", s)
}

// unpackSignatures unpacks the given signatures using the provided unpacker
func unpackSignatures(unpacker codectypes.AnyUnpacker, sigs []*codectypes.Any) ([]signing.SignatureData, error) {
	var signatures = make([]signing.SignatureData, len(sigs))
	for i, sig := range sigs {
		var signatureData SignatureData
		if err := unpacker.UnpackAny(sig, &signatureData); err != nil {
			return nil, err
		}

		cosmosSigData, err := SignatureDataToCosmosSignatureData(unpacker, signatureData)
		if err != nil {
			return nil, err
		}
		signatures[i] = cosmosSigData
	}

	return signatures, nil
}

// --------------------------------------------------------------------------------------------------------------------

var _ SignatureData = &SingleSignatureData{}

// isSignatureData implements SignatureData
func (s *SingleSignatureData) isSignatureData() {}

// --------------------------------------------------------------------------------------------------------------------

var _ SignatureData = &MultiSignatureData{}

// isSignatureData implements SignatureData
func (s *MultiSignatureData) isSignatureData() {}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (s *MultiSignatureData) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, sig := range s.Signatures {
		var data SignatureData
		err := unpacker.UnpackAny(sig, &data)
		if err != nil {
			return err
		}
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// AddressData is an interface representing a generic external chain address
type AddressData interface {
	proto.Message

	// Validate checks the validity of the AddressData
	Validate() error

	// GetValue returns the address value
	GetValue() string

	// VerifyPubKey verifies that the given public key is associated with this address data
	VerifyPubKey(key cryptotypes.PubKey) (bool, error)
}

// --------------------------------------------------------------------------------------------------------------------

var _ AddressData = &Bech32Address{}

// NewBech32Address returns a new Bech32Address instance
func NewBech32Address(value, prefix string) *Bech32Address {
	return &Bech32Address{Value: value, Prefix: prefix}
}

// Validate implements AddressData
func (b Bech32Address) Validate() error {
	if strings.TrimSpace(b.Value) == "" {
		return fmt.Errorf("value cannot be empty or blank")
	}

	if strings.TrimSpace(b.Prefix) == "" {
		return fmt.Errorf("prefix cannot be empty or blank")
	}

	_, err := sdk.GetFromBech32(b.Value, b.Prefix)
	if err != nil {
		return fmt.Errorf("invalid Bech32 value or wrong prefix")
	}

	return nil
}

// GetValue implements AddressData
func (b Bech32Address) GetValue() string {
	return b.Value
}

// VerifyPubKey implements AddressData
func (b Bech32Address) VerifyPubKey(key cryptotypes.PubKey) (bool, error) {
	_, bz, err := bech32.DecodeAndConvert(b.Value)
	if err != nil {
		return false, err
	}
	return bytes.Equal(bz, key.Address().Bytes()), nil
}

// --------------------------------------------------------------------------------------------------------------------

var _ AddressData = &Base58Address{}

// NewBase58Address returns a new Base58Address instance
func NewBase58Address(value string) *Base58Address {
	return &Base58Address{Value: value}
}

// Validate implements AddressData
func (b Base58Address) Validate() error {
	if strings.TrimSpace(b.Value) == "" {
		return fmt.Errorf("address cannot be empty or blank")
	}

	if _, err := base58.Decode(b.Value); err != nil {
		return fmt.Errorf("invalid Base58 address")
	}

	return nil
}

// GetValue implements AddressData
func (b Base58Address) GetValue() string {
	return b.Value
}

// VerifyPubKey implements AddressData
func (b Base58Address) VerifyPubKey(key cryptotypes.PubKey) (bool, error) {
	bz, err := base58.Decode(b.Value)
	return bytes.Equal(tmhash.SumTruncated(bz), key.Address().Bytes()), err
}

// --------------------------------------------------------------------------------------------------------------------

var _ AddressData = &HexAddress{}

// NewHexAddress returns a new HexAddress instance
func NewHexAddress(value, prefix string) *HexAddress {
	return &HexAddress{Value: value, Prefix: prefix}
}

// Validate implements AddressData
func (h HexAddress) Validate() error {
	if strings.TrimSpace(h.Value) == "" {
		return fmt.Errorf("value cannot be empty or blank")
	}

	if len(h.Value) <= len(h.Prefix) {
		return fmt.Errorf("address cannot be smaller than prefix")
	}

	prefix, addrWithoutPrefix := h.Value[:len(h.Prefix)], h.Value[len(h.Prefix):]
	if prefix != h.Prefix {
		return fmt.Errorf("prefix does not match")
	}

	if _, err := hex.DecodeString(addrWithoutPrefix); err != nil {
		return fmt.Errorf("invalid hex address")
	}
	return nil
}

// GetValue implements AddressData
func (h HexAddress) GetValue() string {
	return h.Value
}

// VerifyPubKey implements AddressData
func (h HexAddress) VerifyPubKey(key cryptotypes.PubKey) (bool, error) {
	addr := h.Value[len(h.Prefix):]
	bz, err := hex.DecodeString(addr)
	if err != nil {
		return false, err
	}
	pubKey, err := btcec.ParsePubKey(key.Bytes(), btcec.S256())
	if err != nil {
		return false, err
	}
	uncompressedPubKey := pubKey.SerializeUncompressed()
	return bytes.Equal(crypto.Keccak256(uncompressedPubKey[1:])[12:], bz), err
}

// --------------------------------------------------------------------------------------------------------------------

// UnpackAddressData deserializes the given any type value as an address data using the provided unpacker
func UnpackAddressData(unpacker codectypes.AnyUnpacker, addressAny *codectypes.Any) (AddressData, error) {
	var address AddressData
	if err := unpacker.UnpackAny(addressAny, &address); err != nil {
		return nil, err
	}
	return address, nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewChainLink returns a new ChainLink instance
// nolint:interfacer
func NewChainLink(user string, address AddressData, proof Proof, chainConfig ChainConfig, creationTime time.Time) ChainLink {
	addressAny, err := codectypes.NewAnyWithValue(address)
	if err != nil {
		panic("failed to pack address data to any type")
	}
	return ChainLink{
		User:         user,
		Address:      addressAny,
		Proof:        proof,
		ChainConfig:  chainConfig,
		CreationTime: creationTime,
	}
}

// GetAddressData returns the AddressData associated with this chain link
func (link ChainLink) GetAddressData() AddressData {
	return link.Address.GetCachedValue().(AddressData)
}

// Validate checks the validity of the ChainLink
func (link ChainLink) Validate() error {
	if _, err := sdk.AccAddressFromBech32(link.User); err != nil {
		return fmt.Errorf("invalid creator address: %s", link.User)
	}

	if link.Address == nil {
		return fmt.Errorf("address cannot be nil")
	}

	err := link.Proof.Validate()
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
	if link.Address != nil {
		var address AddressData
		err := unpacker.UnpackAny(link.Address, &address)
		if err != nil {
			return err
		}
	}

	err := link.Proof.UnpackInterfaces(unpacker)
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
