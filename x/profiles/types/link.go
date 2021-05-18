package types

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

// NewChainConfig is a constructor function for ChainConfig
func NewChainConfig(id string, prefix string) ChainConfig {
	return ChainConfig{
		ID:               id,
		Bech32AddrPrefix: prefix,
	}
}

func (chainConfig ChainConfig) Validate() error {
	if chainConfig.ID == "" {
		return fmt.Errorf("chain config id cannot be empty")
	}

	if chainConfig.Bech32AddrPrefix == "" {
		return fmt.Errorf("bech32 addr prefix config id cannot be empty")
	}

	return nil
}

// NewProof is a constructor function for Proof
func NewProof(pubKey string, signature string) Proof {
	return Proof{
		PubKey:    pubKey,
		Signature: signature,
	}
}

func (proof Proof) Validate() error {
	if _, err := hex.DecodeString(proof.PubKey); err != nil {
		return fmt.Errorf("failed to decode hex string of pubkey")
	}

	if _, err := hex.DecodeString(proof.Signature); err != nil {
		return fmt.Errorf("failed to decode hex string of signature")
	}

	return nil
}

// NewLink is a constructor function for Link
func NewLink(address string, proof Proof, chainConfig ChainConfig, creationTime time.Time) Link {
	return Link{
		Address:      address,
		Proof:        proof,
		ChainConfig:  chainConfig,
		CreationTime: creationTime,
	}
}

func (link Link) Validate() error {

	if link.Address == "" {
		return fmt.Errorf("source address cannot be empty")
	}

	chainConfig := link.ChainConfig
	if err := chainConfig.Validate(); err != nil {
		return err
	}

	proof := link.Proof
	if err := proof.Validate(); err != nil {
		return err
	}

	pubKeyBz, _ := hex.DecodeString(proof.PubKey)
	sigBz, _ := hex.DecodeString(proof.Signature)
	pubKey := &secp256k1.PubKey{Key: pubKeyBz}

	proofContent := []byte(link.Address)

	if !pubKey.VerifySignature(proofContent, sigBz) {
		return fmt.Errorf("failed to verify signature")
	}

	return nil
}
