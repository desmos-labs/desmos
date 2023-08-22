package simulation_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/tokenfactory/simulation"
	"github.com/desmos-labs/desmos/v6/x/tokenfactory/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	decoder := simulation.NewDecodeStore(cdc)

	metadata := types.DenomAuthorityMetadata{
		Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	}

	denom := "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken"

	params := types.DefaultParams()

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   append(types.GetDenomPrefixStore("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"), []byte(types.DenomAuthorityMetadataKey)...),
			Value: cdc.MustMarshal(&metadata),
		},
		{
			Key:   append(types.GetCreatorPrefix("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"), []byte(denom)...),
			Value: []byte(denom),
		},
		{
			Key:   []byte(types.ParamsPrefixKey),
			Value: cdc.MustMarshal(&params),
		},
		{
			Key:   []byte("Unknown key"),
			Value: nil,
		},
	}}

	testCases := []struct {
		name        string
		expectedLog string
	}{
		{"Authority Metadata", fmt.Sprintf("Authority MetadataA: %s\nAuthority MetadataB: %s\n",
			metadata, metadata)},
		{"Denom", fmt.Sprintf("DenomA: %s\nDenomB: %s\n",
			denom, denom)},
		{"Params", fmt.Sprintf("ParamsA: %s\nParamsB: %s\n",
			params, params)},
		{"other", ""},
	}

	for i, tc := range testCases {
		i, tc := i, tc
		t.Run(tc.name, func(t *testing.T) {
			switch i {
			case len(testCases) - 1:
				require.Panics(t, func() { decoder(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tc.name)
			default:
				require.Equal(t, tc.expectedLog, decoder(kvPairs.Pairs[i], kvPairs.Pairs[i]), tc.name)
			}
		})
	}
}
