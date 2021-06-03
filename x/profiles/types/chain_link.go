package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ghodss/yaml"
	"github.com/gogo/protobuf/proto"
)

// NewChainConfig is a constructor function for ChainConfig
func NewChainConfig(name string) ChainConfig {
	return ChainConfig{
		Name: name,
	}
}

func (chainConfig ChainConfig) Validate() error {
	if strings.TrimSpace(chainConfig.Name) == "" {
		return fmt.Errorf("chain name cannot be empty or blank")
	}
	return nil
}

// ___________________________________________________________________________________________________________________

// NewProof is a constructor function for Proof
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

func (proof Proof) Validate() error {
	if proof.PubKey == nil {
		return fmt.Errorf("public key field cannot be empty")
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

// -------------------------------------------------------------------------------------------------------------------

type prettyProof struct {
	PubKey    string `json:"public_key" yaml:"public_key"`
	Signature string `json:"signature" yaml:"signature"`
	PlainText string `json:"plain_text" yaml:"plain_text"`
}

// String implements Proof implements stringer
func (proof *Proof) String() string {
	out, _ := proof.MarshalYAML()
	return out.(string)
}

// MarshalYAML returns the YAML representation of a Proof
func (proof *Proof) MarshalYAML() (interface{}, error) {
	bs, err := yaml.Marshal(prettyProof{
		PubKey:    proof.PubKey.String(),
		Signature: proof.Signature,
		PlainText: proof.PlainText,
	})

	if err != nil {
		return nil, err
	}

	return string(bs), nil
}

// MarshalJSON returns the JSON representation of a Proof
func (proof Proof) MarshalJSON() ([]byte, error) {
	return json.Marshal(prettyProof{
		PubKey:    proof.PubKey.String(),
		Signature: proof.Signature,
		PlainText: proof.PlainText,
	})
}

// ___________________________________________________________________________________________________________________

type AddressData interface {
	proto.Message
	Validate() error
	GetAddressString() string
}

func NewBech32Address(value, prefix string) *Bech32Address {
	return &Bech32Address{Value: value, Prefix: prefix}
}

func (address Bech32Address) Validate() error {
	if strings.TrimSpace(address.Value) == "" {
		return fmt.Errorf("address cannot be empty or blank")
	}
	if strings.TrimSpace(address.Prefix) == "" {
		return fmt.Errorf("prefix cannot be empty or blank")
	}
	return nil
}

func (address Bech32Address) GetAddressString() string {
	return address.Value
}

func NewBase58Address(value, prefix string) *Base58Address {
	return &Base58Address{Value: value}
}

func (address Base58Address) Validate() error {
	if strings.TrimSpace(address.Value) == "" {
		return fmt.Errorf("address cannot be empty or blank")
	}
	return nil
}

func (address Base58Address) GetAddressString() string {
	return address.Value
}

func UnpackAddress(unpacker codectypes.AnyUnpacker, addressAny *codectypes.Any) (AddressData, error) {
	var address AddressData
	if err := unpacker.UnpackAny(addressAny, &address); err != nil {
		return nil, err
	}
	return address, nil
}

// ___________________________________________________________________________________________________________________

// NewChainLink is a constructor function for ChainLink
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
		var pubKey cryptotypes.PubKey
		if err := unpacker.UnpackAny(link.Proof.PubKey, &pubKey); err != nil {
			return err
		}
		var address AddressData
		if err := unpacker.UnpackAny(link.Address, &address); err != nil {
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
