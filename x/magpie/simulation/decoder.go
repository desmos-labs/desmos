package simulation

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

// DecodeStore unmarshalls the KVPair's Value to the corresponding magpie type
func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key, types.SessionLengthKey):
			idA := binary.LittleEndian.Uint64(kvA.Value)
			idB := binary.LittleEndian.Uint64(kvB.Value)
			return fmt.Sprintf("DefaultSessionLengthA: %d\nDefaultSessionLengthB: %d\n", idA, idB)

		case bytes.Equal(kvA.Key, types.LastSessionIDStoreKey):
			var idA, idB types.SessionID
			cdc.MustUnmarshalBinaryBare(kvA.Value, &idA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &idB)
			return fmt.Sprintf("LastSessionIDA: %d\nLastSessionIDB: %d\n", idA.Value, idB.Value)

		case bytes.HasPrefix(kvA.Key, types.SessionStorePrefix):
			var sessionA, sessionB types.Session
			cdc.MustUnmarshalBinaryBare(kvA.Value, &sessionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &sessionB)
			return fmt.Sprintf("SessionA: %s\nSessionB: %s\n", sessionA.String(), sessionB.String())

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
