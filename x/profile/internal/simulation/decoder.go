package simulation

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

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
	case bytes.HasPrefix(kvA.Key, types.MonikerStorePrefix):
		var addressA, addressB sdk.AccAddress
		if len(kvB.Value) == 141 {
			println(sdk.AccAddress(kvB.Value).String())
		}
		cdc.MustUnmarshalBinaryBare(kvA.Value, &addressA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &addressB)
		println(addressA.String())
		println(addressB.String())
		return fmt.Sprintf("AddressA: %s\nAddressB: %s\n", addressA, addressB)
	default:
		panic(fmt.Sprintf("invalid profile key %X", kvA.Key))
	}
}
