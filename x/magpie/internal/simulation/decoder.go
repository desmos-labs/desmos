package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"github.com/tendermint/tendermint/libs/kv"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding magpie type
func DecodeStore(cdc *codec.Codec, kvA, kvB kv.Pair) string {
	switch {
	case bytes.Equal(kvA.Key, types.SessionLengthKey):
		var idA, idB int64
		cdc.MustUnmarshalBinaryBare(kvA.Value, &idA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &idB)
		return fmt.Sprintf("DefaultSessionLengthA: %d\nDefaultSessionLengthB: %d\n", idA, idB)
	case bytes.Equal(kvA.Key, types.LastSessionIDStoreKey):
		var postA, postB types.SessionID
		cdc.MustUnmarshalBinaryBare(kvA.Value, &postA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &postB)
		return fmt.Sprintf("LastSessionIDA: %s\nLastSessionIDB: %s\n", postA, postB)
	case bytes.HasPrefix(kvA.Key, types.SessionStorePrefix):
		var commentsA, commentsB types.Session
		cdc.MustUnmarshalBinaryBare(kvA.Value, &commentsA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &commentsB)
		return fmt.Sprintf("SessionA: %s\nSessionB: %s\n", commentsA, commentsB)
	default:
		panic(fmt.Sprintf("invalid magpie key %X", kvA.Key))
	}
}
