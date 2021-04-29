package types

import "github.com/cosmos/cosmos-sdk/codec"

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

// RemoveUser removes the given user address from the provided users slice.
// If the user is found, returns the slice with it removed and true.
// Otherwise, returns the original slice and false
func RemoveUser(users []string, address string) ([]string, bool) {
	for index, user := range users {
		if user == address {
			return append(users[:index], users[index+1:]...), true
		}
	}
	return users, false
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
