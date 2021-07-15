package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding relationships type
func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.DTagPrefix):
			addressA := sdk.AccAddress(kvA.Value).String()
			addressB := sdk.AccAddress(kvB.Value).String()
			return fmt.Sprintf("DTagAddressA: %s\nDTagAddressB: %s\n", addressA, addressB)

		case bytes.HasPrefix(kvA.Key, types.DTagTransferRequestPrefix):
			var requestA, requestB types.DTagTransferRequest
			cdc.MustUnmarshalBinaryBare(kvA.Value, &requestA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &requestB)
			return fmt.Sprintf("RequestA: %s\nRequestB: %s\n", requestA, requestB)

		case bytes.HasPrefix(kvA.Key, types.RelationshipsStorePrefix):
			var relationshipA, relationshipB types.Relationship
			cdc.MustUnmarshalBinaryBare(kvA.Value, &relationshipA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &relationshipB)
			return fmt.Sprintf("Relationships A: %s\nRelationships B: %s\n",
				relationshipA, relationshipB)

		case bytes.HasPrefix(kvA.Key, types.UsersBlocksStorePrefix):
			var userBlockA, userBlockB types.UserBlock
			cdc.MustUnmarshalBinaryBare(kvA.Value, &userBlockA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &userBlockB)
			return fmt.Sprintf("User block A: %s\nUser block B: %s\n",
				userBlockA, userBlockB)

		case bytes.HasPrefix(kvA.Key, types.ChainLinksPrefix):
			var chainLinkA, chainLinkB types.ChainLink
			cdc.MustUnmarshalBinaryBare(kvA.Value, &chainLinkA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &chainLinkB)
			return fmt.Sprintf("Chain link A: %s\nChain link B: %s\n",
				chainLinkA.String(), chainLinkB.String())

		case bytes.HasPrefix(kvA.Key, types.UserApplicationLinkPrefix):
			var applicationLinkA, applicationLinkB types.ApplicationLink
			cdc.MustUnmarshalBinaryBare(kvA.Value, &applicationLinkA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &applicationLinkB)
			return fmt.Sprintf("Application link A: %s\nApplication link B: %s\n",
				applicationLinkA.String(), applicationLinkB.String())

		case bytes.HasPrefix(kvA.Key, types.ApplicationLinkExpirationPrefix):
			var expiringClientKeyA, expiringClientKeyB string
			expiringClientKeyA = string(kvA.Value)
			expiringClientKeyB = string(kvB.Value)
			return fmt.Sprintf("Expiring client key A :%s\nExpiring client key B:%s \n",
				expiringClientKeyA, expiringClientKeyB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
