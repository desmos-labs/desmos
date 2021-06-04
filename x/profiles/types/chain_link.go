package types

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/enigmampc/btcutil/base58"
	"github.com/gogo/protobuf/proto"
)

// NewChainConfig is a constructor function for ChainConfig
func NewChainConfig(name string) ChainConfig {
	return ChainConfig{
		Name: name,
	}
}

// Validate checks the validity of the ChainConfig
func (chainConfig ChainConfig) Validate() error {
	if strings.TrimSpace(chainConfig.Name) == "" {
		return fmt.Errorf("chain name cannot be empty or blank")
	}
	return nil
}

// ___________________________________________________________________________________________________________________

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
func (proof Proof) Validate() error {
	if proof.PubKey == nil {
		return fmt.Errorf("public key field cannot be nil")
	}
	_, err := hex.DecodeString(proof.Signature)
	if err != nil {
		return fmt.Errorf("failed to decode hex string of signature")
	}
	if strings.TrimSpace(proof.PlainText) == "" {
		return fmt.Errorf("plain text cannot be empty or blank")
	}
	return nil
}

// Verify verifies the signature is signed from the plain text by the public key and returns error if something invalid
// the Proof object has to unpack public key with given unpacker first
func (proof Proof) Verify(unpacker codectypes.AnyUnpacker) error {
	var pubkey cryptotypes.PubKey
	err := unpacker.UnpackAny(proof.PubKey, &pubkey)
	if err != nil {
		return fmt.Errorf("failed to unpack the pubkey")
	}
	sig, _ := hex.DecodeString(proof.Signature)
	if !pubkey.VerifySignature([]byte(proof.PlainText), sig) {
		return fmt.Errorf("failed to verify the signature")
	}
	return nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (proof *Proof) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pubKey cryptotypes.PubKey
	return unpacker.UnpackAny(proof.PubKey, &pubKey)
}

// ___________________________________________________________________________________________________________________

// AddressData is an interface which allows chain link to store multiple encoded types (Bech32, Base58 and etc.) of the address
type AddressData interface {
	proto.Message

	// Validate checks the validity of the AddressData
	Validate() error

	// GetAddress returns the address value of AddressData
	GetAddress() string
}

// NewBech32Address is a constructor function for Bech32Address
func NewBech32Address(value, prefix string) *Bech32Address {
	return &Bech32Address{Value: value, Prefix: prefix}
}

// Validate checks the validity of the Bech32Address
func (address Bech32Address) Validate() error {
	if strings.TrimSpace(address.Value) == "" {
		return fmt.Errorf("address cannot be empty or blank")
	}
	if strings.TrimSpace(address.Prefix) == "" {
		return fmt.Errorf("prefix cannot be empty or blank")
	}
	if _, err := sdk.GetFromBech32(address.Value, address.Prefix); err != nil {
		return nil
	}
	return nil
}

// GetAddress returns the address value
func (address Bech32Address) GetAddress() string {
	return address.Value
}

// NewBase58Address is a constructor function for Base58Address
func NewBase58Address(value string) *Base58Address {
	return &Base58Address{Value: value}
}

// Validate checks the validity of the Base58Address
func (address Base58Address) Validate() error {
	if strings.TrimSpace(address.Value) == "" {
		return fmt.Errorf("address cannot be empty or blank")
	}
	if _, _, err := base58.CheckDecode(address.Value); err != nil {
		return err
	}
	return nil
}

// GetAddress returns the address value
func (address Base58Address) GetAddress() string {
	return address.Value
}

// UnpackAddressData deserializes the given any type value as an address data using the provided unpacker
func UnpackAddressData(unpacker codectypes.AnyUnpacker, addressAny *codectypes.Any) (AddressData, error) {
	var address AddressData
	if err := unpacker.UnpackAny(addressAny, &address); err != nil {
		return nil, err
	}
	return address, nil
}

// ___________________________________________________________________________________________________________________

// NewChainLink is a constructor function for ChainLink
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

	if err := link.ChainConfig.Validate(); err != nil {
		return err
	}
	if err := link.Proof.Validate(); err != nil {
		return err
	}
	if link.CreationTime.IsZero() {
		return fmt.Errorf("createion time cannot be zero")
	}
	return nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (link *ChainLink) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if link != nil {
		var address AddressData
		if err := unpacker.UnpackAny(link.Address, &address); err != nil {
			return err
		}
		if err := link.Proof.UnpackInterfaces(unpacker); err != nil {
			return err
		}
	}
	return nil
}

// MustMarshalChainLinks serializes the given chain link using the provided BinaryMarshaler
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
