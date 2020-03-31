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
	case bytes.HasPrefix(kvA.Key, types.ProfileStorePrefix):
		var profileA, profileB types.Profile
		cdc.MustUnmarshalBinaryBare(kvA.Value, &profileA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &profileB)
		return fmt.Sprintf("ProfileA: %s\nProfileB: %s\n", profileA, profileB)
	default:
		panic(fmt.Sprintf("invalid profile key %X", kvA.Key))
	}
}
