package types

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
)

// ____________________________________________________________________________________________________________________

// IsPresent checks if the given address is a present inside the users slice
func (users Users) IsPresent(address string) bool {
	for _, user := range users.Users {
		if user == address {
			return true
		}
	}
	return false
}

// AppendUser append the given address to the users slice
func (users Users) AppendUser(address string) Users {
	users.Users = append(users.Users, address)
	return users
}

// RemoveUser remove the given address from the users slice
func (users Users) RemoveUser(address string) Users {
	for index, user := range users.Users {
		if user == address {
			users.Users = append(users.Users[:index], users.Users[index+1:]...)
		}
	}
	return users
}

// ValidateUsers checks the validity of the given wrapped users slice that contains users of the given userType.
// It returns error if one of them is invalid.
func (users Users) ValidateUsers(userType string) error {
	for _, user := range users.Users {
		if user == "" {
			return fmt.Errorf("empty %s address", userType)
		}
	}
	return nil
}

// MustMarshalUsers marshals the given users into an array of bytes.
// Panics on error.
func MustMarshalUsers(cdc codec.BinaryMarshaler, users Users) []byte {
	bz, err := cdc.MarshalBinaryBare(&users)
	if err != nil {
		panic(err)
	}
	return bz
}

// MustUnmarshalUsers tries unmarshalling the given bz to a slice of users.
// Panics on error.
func MustUnmarshalUsers(cdc codec.BinaryMarshaler, bz []byte) Users {
	var wrapped Users
	err := cdc.UnmarshalBinaryBare(bz, &wrapped)
	if err != nil {
		panic(err)
	}
	return wrapped
}
