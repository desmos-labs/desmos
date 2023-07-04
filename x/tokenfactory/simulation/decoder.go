package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding tokenfactory type
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, []byte(types.DenomsPrefixKey)):
			var metadataA, metadataB types.DenomAuthorityMetadata
			cdc.MustUnmarshal(kvA.Value, &metadataA)
			cdc.MustUnmarshal(kvB.Value, &metadataB)
			return fmt.Sprintf("Authority MetadataA: %s\nAuthority MetadataB: %s\n", metadataA, metadataB)

		case bytes.HasPrefix(kvA.Key, []byte(types.CreatorPrefixKey)):
			return fmt.Sprintf("DenomA: %s\nDenomB: %s\n", kvA.Value, kvB.Value)

		case bytes.HasPrefix(kvA.Key, []byte(types.ParamsPrefixKey)):
			var paramA, paramB types.Params
			cdc.MustUnmarshal(kvA.Value, &paramA)
			cdc.MustUnmarshal(kvA.Value, &paramB)
			return fmt.Sprintf("ParamsA: %s\nParamsB: %s\n", paramA, paramB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
