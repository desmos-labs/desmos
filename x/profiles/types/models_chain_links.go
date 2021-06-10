package types

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/mr-tron/base58"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
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
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewProof is a constructor function for Proof
// nolint:interfacer
func NewProof(pubKey cryptotypes.PubKey, signature string, plainText string) Proof {
	pubKeyAny, err := codectypes.NewAnyWithValue(pubKey)
	if err != nil {
		panic("failed to pack public key to any type")
	}
	return Proof{
		PubKey:    pubKeyAny,
		Signature: signature,
		PlainText: plainText,
	}
}

// Validate checks the validity of the Proof
func (p Proof) Validate() error {
	if p.PubKey == nil {
		return fmt.Errorf("public key field cannot be nil")
	}

	_, err := hex.DecodeString(p.Signature)
	if err != nil {
		return fmt.Errorf("invalid hex-encoded signature")
	}

	if strings.TrimSpace(p.PlainText) == "" {
		return fmt.Errorf("plain text cannot be empty or blank")
	}

	return nil
}

// Verify verifies the signature using the given plain text and public key.
// It returns and error if something is invalid.
func (p Proof) Verify(unpacker codectypes.AnyUnpacker) error {
	var pubkey cryptotypes.PubKey
	err := unpacker.UnpackAny(p.PubKey, &pubkey)
	if err != nil {
		return fmt.Errorf("failed to unpack the public key")
	}

	sig, _ := hex.DecodeString(p.Signature)
	if !pubkey.VerifySignature([]byte(p.PlainText), sig) {
		return fmt.Errorf("failed to verify the signature")
	}

	return nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (p *Proof) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pubKey cryptotypes.PubKey
	return unpacker.UnpackAny(p.PubKey, &pubKey)
}

// --------------------------------------------------------------------------------------------------------------------

// AddressData is an interface representing a generic external chain address
type AddressData interface {
	proto.Message

	// Validate checks the validity of the AddressData
	Validate() error

	// GetAddress returns the address value
	GetAddress() string
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

// GetAddress implements AddressData
func (b Bech32Address) GetAddress() string {
	return b.Value
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

// GetAddress implements AddressData
func (b Base58Address) GetAddress() string {
	return b.Value
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
func NewChainLink(address AddressData, proof Proof, chainConfig ChainConfig, creationTime time.Time) ChainLink {
	addressAny, err := codectypes.NewAnyWithValue(address)
	if err != nil {
		panic("failed to pack address data to any type")
	}
	return ChainLink{
		Address:      addressAny,
		Proof:        proof,
		ChainConfig:  chainConfig,
		CreationTime: creationTime,
	}
}

// Validate checks the validity of the ChainLink
func (link ChainLink) Validate() error {
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

// MustMarshalChainLink serializes the given chain link using the provided BinaryMarshaler
func MustMarshalChainLink(cdc codec.BinaryMarshaler, link ChainLink) []byte {
	return cdc.MustMarshalBinaryBare(&link)
}

// MustUnmarshalChainLink deserializes the given byte array as a chain link using
// the provided BinaryMarshaler
func MustUnmarshalChainLink(codec codec.BinaryMarshaler, bz []byte) ChainLink {
	var link ChainLink
	codec.MustUnmarshalBinaryBare(bz, &link)
	return link
}

// RemoveChainLinkIfPresent iterates over the given chain links slice and tries removing the link
// with the given chain name and address.
// If found, returns a new slice that does not contain that link and "true".
// Otherwise, it returns the original slice and "false".
func RemoveChainLinkIfPresent(links []ChainLink, chainName, address string) ([]ChainLink, bool) {
	for index, link := range links {
		addressData := link.GetAddress().GetCachedValue().(AddressData)
		if link.ChainConfig.Name == chainName && addressData.GetAddress() == address {
			return append(links[:index], links[index+1:]...), true
		}
	}
	return links, false
}
