package types

// DONTCOVER

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/proto"
	"github.com/mr-tron/base58"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
	pubKey, err := btcec.ParsePubKey(key.Bytes())
	if err != nil {
		return false, err
	}
	uncompressedPubKey := pubKey.SerializeUncompressed()
	return bytes.Equal(crypto.Keccak256(uncompressedPubKey[1:])[12:], bz), err
}
