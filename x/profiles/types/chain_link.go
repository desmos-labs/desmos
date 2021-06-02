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

func NewAddress(address, prefix string) Address {
	var addr isAddress_Sum
	if prefix != "" {
		bech32 := Address_Bech32{
			Value:  address,
			Prefix: prefix,
		}
		addr = &Address_Bech_32{
			Bech_32: &bech32,
		}
	} else {
		base58 := Address_Base58{
			Value: address,
		}
		addr = &Address_Base_58{
			Base_58: &base58,
		}

	}
	return Address{Sum: addr}
}

func (address Address_Bech32) Validate() error {
	if strings.TrimSpace(address.Value) == "" {
		return fmt.Errorf("address cannot be empty or blank")
	}
	if strings.TrimSpace(address.Prefix) == "" {
		return fmt.Errorf("prefix cannot be empty or blank")
	}
	return nil
}

func (address Address_Base58) Validate() error {
	if strings.TrimSpace(address.Value) == "" {
		return fmt.Errorf("address cannot be empty or blank")
	}
	return nil
}

func (address Address) Validate() error {
	switch address.Sum.(type) {
	case *Address_Bech_32:
		return address.GetBech_32().Validate()
	case *Address_Base_58:
		return address.GetBase_58().Validate()
	default:
		return fmt.Errorf("unknown address type")
	}
}

func (address Address) GetValue() string {
	switch address.Sum.(type) {
	case *Address_Bech_32:
		return address.GetBech_32().GetValue()
	case *Address_Base_58:
		return address.GetBase_58().GetValue()
	default:
		return ""
	}
}

// ___________________________________________________________________________________________________________________

// NewChainLink is a constructor function for ChainLink
func NewChainLink(address Address, proof Proof, chainConfig ChainConfig, creationTime time.Time) ChainLink {
	return ChainLink{
		Address:      address,
		Proof:        proof,
		ChainConfig:  chainConfig,
		CreationTime: creationTime,
	}
}

func (link ChainLink) Validate() error {
	if err := link.Address.Validate(); err != nil {
		return err
	}
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
		return unpacker.UnpackAny(link.Proof.PubKey, &pubKey)
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
