package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RelationshipID represents a unique relationship id
type RelationshipID string

// Valid tells if the id can be used safely
func (id RelationshipID) Valid() bool {
	return strings.TrimSpace(id.String()) != ""
}

// String implements fmt.Stringer
func (id RelationshipID) String() string {
	return string(id)
}

// Equals compares two RelationshipID instances
func (id RelationshipID) Equals(other RelationshipID) bool {
	return id == other
}

// Relationship represents a single relationship between two users. Creator is the one that first
// sent the relationship request, and Recipient is the one that received it and (optionally) accepted it.
type Relationship interface {
	RelationshipID() RelationshipID
	Creator() sdk.AccAddress
	Recipient() sdk.AccAddress
}

type Relationships []Relationship

// MonodirectionalRelationship implements Relationship and represents a monodirectional
// relationships that does not require the receiver to accept it before being effective.
type MonodirectionalRelationship struct {
	ID       RelationshipID `json:"id" yaml:"id"`
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
}

func NewMonodirectionalRelationship(sender, receiver sdk.AccAddress) MonodirectionalRelationship {
	rel := MonodirectionalRelationship{
		ID:       "",
		Sender:   sender,
		Receiver: receiver,
	}

	// RelationshipID calculation
	bz := []byte("monodirectional" + sender.String() + receiver.String())
	hash := sha256.Sum256(bz)

	rel.ID = RelationshipID(hex.EncodeToString(hash[:]))

	return rel
}

func (mr MonodirectionalRelationship) RelationshipID() RelationshipID {
	return mr.ID
}

// Creator implements Relationship.Creator
func (mr MonodirectionalRelationship) Creator() sdk.AccAddress {
	return mr.Sender
}

// Recipient implements Relationship.Recipient
func (mr MonodirectionalRelationship) Recipient() sdk.AccAddress {
	return mr.Receiver
}

// String implement fmt.Stringer
func (mr MonodirectionalRelationship) String() string {
	out := "Mono directional Relationship:\n"
	out += fmt.Sprintf("[RelationshipID] %s [Sender] %s -> [Receiver] %s", mr.ID, mr.Sender, mr.Receiver)
	return out
}

// Equals allows to check whether the contents of mr are the same of other
func (mr MonodirectionalRelationship) Equals(other MonodirectionalRelationship) bool {
	return mr.Sender.Equals(other.Sender) && mr.Receiver.Equals(other.Receiver)
}

// Validate checks the validity of the MonodirectionalRelationship
func (mr MonodirectionalRelationship) Validate() error {
	if mr.Sender.Empty() {
		return fmt.Errorf("relationship sender cannot be empty")
	}

	if mr.Receiver.Empty() {
		return fmt.Errorf("relationship receiver cannot be empty")
	}

	if mr.Sender.Equals(mr.Receiver) {
		return fmt.Errorf("you can't create a relationship with yourself")
	}

	return nil
}

type MonoDirectionalRelationships []MonodirectionalRelationship

// AppendIfMissing appends the given mr to the mrs slice if it does not exist inside it yet.
// It returns a new slice of MonoDirectionalRelationships containing such otherMr and a true value meaning that the otherMr has been appended.
func (mrs MonoDirectionalRelationships) AppendIfMissing(otherMr MonodirectionalRelationship) (MonoDirectionalRelationships, bool) {
	for _, mr := range mrs {
		if mr.Equals(otherMr) {
			return nil, false
		}
	}
	return append(mrs, otherMr), true
}

// BidirectionalRelationship implements Relationship and represents a bidirectional relationship
// that can have different statuses and requires the receiver to accept it before becoming effective.
type BidirectionalRelationship struct {
	ID       RelationshipID     `json:"id" yaml:"id"`
	Sender   sdk.AccAddress     `json:"sender" yaml:"sender"`
	Receiver sdk.AccAddress     `json:"receiver" yaml:"receiver"`
	Status   RelationshipStatus `json:"status" yaml:"status"`
}

// RelationshipStatus represents the status of a bidirectional relationship
type RelationshipStatus int

// String implements fmt.Stringer
func (rs RelationshipStatus) String() string {
	return strconv.Itoa(int(rs))
}

const (
	Sent     RelationshipStatus = 0 // Indicates that the relationship has been sent but not yet accepted or denied
	Accepted RelationshipStatus = 1 // Tells that the relationships has been accepted by the receiver
	Denied   RelationshipStatus = 2 // Tells that the relationship has been denied from the receiver
)

func NewBiDirectionalRelationship(sender, receiver sdk.AccAddress, status RelationshipStatus) BidirectionalRelationship {
	rel := BidirectionalRelationship{
		ID:       "",
		Sender:   sender,
		Receiver: receiver,
		Status:   status,
	}

	// RelationshipID calculation
	bz := []byte("bidirectional" + sender.String() + receiver.String())
	hash := sha256.Sum256(bz)

	rel.ID = RelationshipID(hex.EncodeToString(hash[:]))

	return rel
}

func (br BidirectionalRelationship) RelationshipID() RelationshipID {
	return br.ID
}

// Creator implements Relationship.Creator
func (br BidirectionalRelationship) Creator() sdk.AccAddress {
	return br.Sender
}

// Recipient implements Relationship.Recipient
func (br BidirectionalRelationship) Recipient() sdk.AccAddress {
	return br.Receiver
}

// String implement fmt.Stringer
func (br BidirectionalRelationship) String() string {
	out := "Bidirectional Relationship:\n"
	out += fmt.Sprintf("[RelationshipID] %s [Sender] %s <-> [Receiver] %s\n", br.ID, br.Sender, br.Receiver)

	switch br.Status {
	case 0:
		out += fmt.Sprintf("Status: %s", "Relationship not yet accepted or denied")
	case 1:
		out += fmt.Sprintf("Status: %s", "Relationship accepted")
	case 2:
		out += fmt.Sprintf("Status: %s", "Relationship denied")
	}

	return out
}

// Equals allows to check whether the contents of br are the same of other
func (br BidirectionalRelationship) Equals(other BidirectionalRelationship) bool {
	return br.Sender.Equals(other.Sender) && br.Receiver.Equals(other.Receiver) && br.Status == other.Status
}

// Validate checks the validity of the BidirectionalRelationship
func (br BidirectionalRelationship) Validate() error {
	if br.Sender.Empty() {
		return fmt.Errorf("relationship sender cannot be empty")
	}

	if br.Receiver.Empty() {
		return fmt.Errorf("relationship receiver cannot be empty")
	}

	if br.Sender.Equals(br.Receiver) {
		return fmt.Errorf("you can't create a relationship with yourself")
	}

	return nil
}

type BidirectionalRelationships []BidirectionalRelationship

// AppendIfMissing appends the given br to the brs slice if it does not exist inside it yet.
// It returns a new slice of BidirectionalRelationships containing such otherBr and a true value meaning that the otherBr has been appended.
func (brs BidirectionalRelationships) AppendIfMissing(otherBr BidirectionalRelationship) (BidirectionalRelationships, bool) {
	for _, br := range brs {
		if br.Equals(otherBr) {
			return nil, false
		}
	}
	return append(brs, otherBr), true
}
