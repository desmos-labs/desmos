package types

import "github.com/cosmos/cosmos-sdk/codec"

// ____________________________________________________________________________________________________________________

// IsAdmin checks if the given address is a subspace admin
func (admins Admins) IsAdmin(address string) bool {
	for _, admin := range admins.Admins {
		if admin == address {
			return true
		}
	}
	return false
}

// AddAdminIfMissing appends the given address to the admins slice if it does not exist inside it yet.
// It returns a new slice of  containing such address.
func AddAdminIfMissing(admins []string, address string) []string {
	for _, admin := range admins {
		if admin == address {
			return admins
		}
	}
	return append(admins, address)
}

// RemoveAdmin removes the given admin address from the provided admins slice.
// If the admin was found, returns the slice with it removed and true.
// Otherwise, returns the original slice and false
func RemoveAdmin(admins []string, address string) ([]string, bool) {
	for index, admin := range admins {
		if admin == address {
			return append(admins[:index], admins[index+1:]...), true
		}
	}
	return admins, false
}

// MustMarshalAdmins marshals the given admins into an array of bytes.
// Panics on error.
func MustMarshalAdmins(cdc codec.BinaryMarshaler, admins Admins) []byte {
	bz, err := cdc.MarshalBinaryBare(&admins)
	if err != nil {
		panic(err)
	}
	return bz
}

// MustUnmarshalAdmins tries unmarshalling the given bz to a slice of admins.
// Panics on error.
func MustUnmarshalAdmins(cdc codec.BinaryMarshaler, bz []byte) Admins {
	var wrapped Admins
	err := cdc.UnmarshalBinaryBare(bz, &wrapped)
	if err != nil {
		panic(err)
	}
	return wrapped
}
