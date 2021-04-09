package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/desmos-labs/desmos/x/links/types"
)

// LinksUnmarshaler defines the expected encoding store functions.
type LinksUnmarshaler interface {
	UnmarshalLink([]byte) (types.Link, error)
}

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding link type
func NewDecodeStore(cdc codec.BinaryMarshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.LinksStorePrefix):
			var link types.Link
			cdc.MustUnmarshalBinaryBare(kvA.Value, &link)
			return fmt.Sprintf("Link: %s\n", link)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
