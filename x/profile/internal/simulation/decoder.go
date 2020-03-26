package simulation

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/tendermint/tendermint/libs/kv"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding profile type
func DecodeStore(cdc *codec.Codec, kvA, kvB kv.Pair) string {
	switch {
	case bytes.Equal(kvA.Key, types.AccountStorePrefix):
		var accA, accB types.Account
		cdc.MustUnmarshalBinaryBare(kvA.Value, &accA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &accB)
		return fmt.Sprintf("AccountA: %s\nAccountB: %s\n", accA, accB)
	default:
		panic(fmt.Sprintf("invalid account key %X", kvA.Key))
	}
}
