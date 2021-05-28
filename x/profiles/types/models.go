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

	"github.com/desmos-labs/desmos/x/commons"
)

// NewPictures is a constructor function for Pictures
func NewPictures(profile, cover string) Pictures {
	return Pictures{
		Profile: profile,
		Cover:   cover,
	}
}

// Validate check the validity of the Pictures
func (pic Pictures) Validate() error {
	if pic.Profile != "" {
		valid := commons.IsURIValid(pic.Profile)
		if !valid {
			return fmt.Errorf("invalid profile picture uri provided")
		}
	}

	if pic.Cover != "" {
		valid := commons.IsURIValid(pic.Cover)
		if !valid {
			return fmt.Errorf("invalid profile cover uri provided")
		}
	}

	return nil
}

// ___________________________________________________________________________________________________________________

func NewDTagTransferRequest(dTagToTrade string, sender, receiver string) DTagTransferRequest {
	return DTagTransferRequest{
		DTagToTrade: dTagToTrade,
		Receiver:    receiver,
		Sender:      sender,
	}
}

// Validate checks the request validity
func (request DTagTransferRequest) Validate() error {
	_, err := sdk.AccAddressFromBech32(request.Sender)
	if err != nil {
		return fmt.Errorf("invalid DTag transfer request sender address: %s", request.Sender)
	}

	_, err = sdk.AccAddressFromBech32(request.Receiver)
	if err != nil {
		return fmt.Errorf("invalid receiver address: %s", request.Receiver)
	}

	if request.Receiver == request.Sender {
		return fmt.Errorf("the sender and receiver must be different")
	}

	if strings.TrimSpace(request.DTagToTrade) == "" {
		return fmt.Errorf("invalid DTag to trade: %s", request.DTagToTrade)
	}

	return nil
}

// NewDTagTransferRequests returns a DTagTransferRequests instance wrapping the given requests
func NewDTagTransferRequests(requests []DTagTransferRequest) DTagTransferRequests {
	return DTagTransferRequests{Requests: requests}
}

// ___________________________________________________________________________________________________________________

// NewRelationship returns a new relationships with the given recipient and subspace
func NewRelationship(creator string, recipient string, subspace string) Relationship {
	return Relationship{
		Creator:   creator,
		Recipient: recipient,
		Subspace:  subspace,
	}
}

// Validate implement Validator
func (r Relationship) Validate() error {
	_, err := sdk.AccAddressFromBech32(r.Creator)
	if err != nil {
		return fmt.Errorf("invalid creator address: %s", r.Creator)
	}

	_, err = sdk.AccAddressFromBech32(r.Recipient)
	if err != nil {
		return fmt.Errorf("invalid recipient address: %s", r.Recipient)
	}

	if r.Creator == r.Recipient {
		return fmt.Errorf("creator and recipient cannot be the same user")
	}

	if !commons.IsValidSubspace(r.Subspace) {
		return fmt.Errorf("subspace must be a valid sha-256")
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// RemoveRelationship removes the given relationships from the provided relationships slice.
// If the relationship was found, returns the slice with it removed and true.
// Otherwise, returns the original slice and false
func RemoveRelationship(relationships []Relationship, relationship Relationship) ([]Relationship, bool) {
	for index, rel := range relationships {
		if rel.Equal(relationship) {
			return append(relationships[:index], relationships[index+1:]...), true
		}
	}
	return relationships, false
}

// MustMarshalRelationships serializes the given relationships using the provided BinaryMarshaler
func MustMarshalRelationships(cdc codec.BinaryMarshaler, relationships []Relationship) []byte {
	wrapped := Relationships{Relationships: relationships}
	return cdc.MustMarshalBinaryBare(&wrapped)
}

// MustUnmarshalRelationships deserializes the given byte array as an array of relationships using
// the provided BinaryMarshaler
func MustUnmarshalRelationships(codec codec.BinaryMarshaler, bz []byte) []Relationship {
	var wrapped Relationships
	codec.MustUnmarshalBinaryBare(bz, &wrapped)
	return wrapped.Relationships
}

// ___________________________________________________________________________________________________________________

// NewUserBlock returns a new object representing the fact that one user has blocked another one
// for a specific reason on the given subspace.
func NewUserBlock(blocker, blocked string, reason, subspace string) UserBlock {
	return UserBlock{
		Blocker:  blocker,
		Blocked:  blocked,
		Reason:   reason,
		Subspace: subspace,
	}
}

// Validate implements validator
func (ub UserBlock) Validate() error {
	if len(ub.Blocker) == 0 {
		return fmt.Errorf("blocker address cannot be empty")
	}

	if len(ub.Blocked) == 0 {
		return fmt.Errorf("the address of the blocked user cannot be empty")
	}

	if ub.Blocker == ub.Blocked {
		return fmt.Errorf("blocker and blocked addresses cannot be equals")
	}

	if !commons.IsValidSubspace(ub.Subspace) {
		return fmt.Errorf("subspace must be a valid sha-256 hash")
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// RemoveUserBlock removes the block made from the blocker towards the blocked inside the subspace,
// from the provided slice of blocks.
// If the block is found, returns the new slice with it removed and true.
// If the block is not found, returns the original fl
func RemoveUserBlock(blocks []UserBlock, blocker, blocked, subspace string) ([]UserBlock, bool) {
	for index, ub := range blocks {
		if ub.Blocker == blocker && ub.Blocked == blocked && ub.Subspace == subspace {
			return append(blocks[:index], blocks[index+1:]...), true

		}
	}
	return blocks, false
}

// MustMarshalUserBlocks serializes the given blocks using the provided BinaryMarshaler
func MustMarshalUserBlocks(cdc codec.BinaryMarshaler, block []UserBlock) []byte {
	wrapped := UserBlocks{Blocks: block}
	return cdc.MustMarshalBinaryBare(&wrapped)
}

// MustUnmarshalUserBlocks deserializes the given byte array as an array of blocks using
// the provided BinaryMarshaler
func MustUnmarshalUserBlocks(cdc codec.BinaryMarshaler, bz []byte) []UserBlock {
	var wrapped UserBlocks
	cdc.MustUnmarshalBinaryBare(bz, &wrapped)
	return wrapped.Blocks
}

// ___________________________________________________________________________________________________________________

// NewChainConfig is a constructor function for ChainConfig
func NewChainConfig(name string, prefix string) ChainConfig {
	return ChainConfig{
		Name:             name,
		Bech32AddrPrefix: prefix,
	}
}

func (chainConfig ChainConfig) Validate() error {
	if chainConfig.Name == "" {
		return fmt.Errorf("chain name cannot be empty")
	}
	return nil
}

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
	_, err := hex.DecodeString(proof.Signature)
	if err != nil {
		return fmt.Errorf("failed to decode hex string of signature")
	}

	if strings.TrimSpace(proof.PlainText) == "" {
		return fmt.Errorf("plain text can not be empty or blank")
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

// NewChainLink is a constructor function for ChainLink
func NewChainLink(address string, proof Proof, chainConfig ChainConfig, creationTime time.Time) ChainLink {
	return ChainLink{
		Address:      address,
		Proof:        proof,
		ChainConfig:  chainConfig,
		CreationTime: creationTime,
	}
}

func (chainLink ChainLink) Validate() error {

	if chainLink.Address == "" {
		return fmt.Errorf("source address cannot be empty")
	}

	if err := chainLink.ChainConfig.Validate(); err != nil {
		return err
	}

	if err := chainLink.Proof.Validate(); err != nil {
		return err
	}

	return nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (chainLink *ChainLink) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pubKey cryptotypes.PubKey
	return unpacker.UnpackAny(chainLink.Proof.PubKey, &pubKey)
}

// MustMarshalUserBlocks serializes the given chain link using the provided BinaryMarshaler
func MustMarshalChainLink(cdc codec.BinaryMarshaler, chainLink ChainLink) []byte {
	return cdc.MustMarshalBinaryBare(&chainLink)
}

// MustUnmarshalChainLink deserializes the given byte array as a chain link using
// the provided BinaryMarshaler
func MustUnmarshalChainLink(codec codec.BinaryMarshaler, bz []byte) ChainLink {
	var chainlink ChainLink
	codec.MustUnmarshalBinaryBare(bz, &chainlink)
	return chainlink
}
