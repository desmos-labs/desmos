package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
	"github.com/tendermint/tendermint/libs/kv"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding profile type
func DecodeStore(cdc *codec.Codec, kvA, kvB kv.Pair) string {
	switch {
	case bytes.HasPrefix(kvA.Key, types.RelationshipsStorePrefix):
		var relationshipsA, relationshipsB []sdk.AccAddress
		cdc.MustUnmarshalBinaryBare(kvA.Value, &relationshipsA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &relationshipsB)
		return fmt.Sprintf("Relationships: %s\nRelationships: %s\n", relationshipsA, relationshipsB)
	default:
		panic(fmt.Sprintf("invalid relationships key %X", kvA.Key))
	}
}
