package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
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
	case bytes.HasPrefix(kvA.Key, types.DtagStorePrefix):
		var addressA, addressB sdk.AccAddress
		cdc.MustUnmarshalBinaryBare(kvA.Value, &addressA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &addressB)
		return fmt.Sprintf("AddressA: %s\nAddressB: %s\n", addressA, addressB)
	case bytes.HasPrefix(kvA.Key, types.RelationshipsStorePrefix):
		var relationshipA, relationshipB types.Relationship
		cdc.MustUnmarshalBinaryBare(kvA.Value, &relationshipA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &relationshipB)
		return fmt.Sprintf("RelationshipA: %s\nRelationshipB: %s\n", relationshipA, relationshipB)
	case bytes.HasPrefix(kvA.Key, types.UserRelationshipsStorePrefix):
		var relationshipIDA, relationshipIDB types.RelationshipID
		cdc.MustUnmarshalBinaryBare(kvA.Value, &relationshipIDA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &relationshipIDB)
		return fmt.Sprintf("RelationshipIDA: %s\nRelationshipIDB: %s\n", relationshipIDA, relationshipIDB)
	default:
		panic(fmt.Sprintf("invalid profile key %X", kvA.Key))
	}
}
