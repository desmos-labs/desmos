package types

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

// NewChainConfig is a constructor function for ChainConfig
func NewChainConfig(ID string, prefix string) ChainConfig {
	return ChainConfig{
		Id:               ID,
		Bech32AddrPrefix: prefix,
	}
}

func (chainConfig ChainConfig) Validate() error {
	if chainConfig.Id == "" {
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
func NewLink(srcaddr string, destAddr string, proof Proof, chainConfig ChainConfig, creationTime time.Time) Link {
	return Link{
		SourceAddress:      srcaddr,
		DestinationAddress: destAddr,
		Proof:              proof,
		ChainConfig:        chainConfig,
		CreationTime:       creationTime,
	}
}

func (link Link) Validate() error {

	if link.SourceAddress == "" {
		return fmt.Errorf("source address cannot be empty")
	}

	if link.DestinationAddress == "" {
		return fmt.Errorf("destination address cannot be empty")
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

	proofContent := []byte(link.SourceAddress + "-" + link.DestinationAddress)

	if !pubKey.VerifySignature(proofContent, sigBz) {
		return fmt.Errorf("failed to verify signature")
	}

	return nil
}
