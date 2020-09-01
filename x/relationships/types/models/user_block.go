package models

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// UserBlock represents the fact that the Blocker has blocked the given Blocked user.
// The Reason field represents the reason the user has been blocked for, and is optional.
type UserBlock struct {
	Blocker sdk.AccAddress `json:"blocker" yaml:"blocker"`
	Blocked sdk.AccAddress `json:"blocked" yaml:"blocked"`
	Reason  string         `json:"reason,omitempty" yaml:"reason,omitempty"`
}

func NewUserBlock(blocker, blocked sdk.AccAddress, reason string) UserBlock {
	return UserBlock{
		Blocker: blocker,
		Blocked: blocked,
		Reason:  reason,
	}
}

// String implements fmt.Stringer
func (ub UserBlock) String() string {
	out := "User Block: "
	out += fmt.Sprintf("[Blocker] %s [Blocked] %s [Reason] %s", ub.Blocker, ub.Blocked, ub.Reason)
	return out
}

// Validate implements validator
func (ub UserBlock) Validate() error {
	if ub.Blocker.Empty() {
		return fmt.Errorf("blocker address cannot be empty")
	}

	if ub.Blocked.Empty() {
		return fmt.Errorf("the address of the blocked user cannot be empty")
	}

	if ub.Blocker.Equals(ub.Blocked) {
		return fmt.Errorf("blocker and blocked addresses cannot be equals")
	}

	return nil
}

// Equals checks if the two user blocks have the same content
func (ub UserBlock) Equals(other UserBlock) bool {
	return ub.Blocker.Equals(other.Blocker) &&
		ub.Blocked.Equals(other.Blocked) &&
		ub.Reason == other.Reason
}
