package simulation_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/app"

	"github.com/desmos-labs/desmos/v5/x/profiles/simulation"
	"github.com/desmos-labs/desmos/v5/x/profiles/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	addr, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	request := types.NewDTagTransferRequest(
		"dtag",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	chainLink := types.NewChainLink(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		types.NewBech32Address("cosmos", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
		types.Proof{},
		types.NewChainConfig("cosmos"),
		time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
	)
	applicationLink := types.NewApplicationLink(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		types.NewData("application", "username"),
		types.ApplicationLinkStateInitialized,
		types.OracleRequest{},
		nil,
		time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
	)

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.DTagStoreKey("AAkvohxhflhXsuyMg"),
			Value: addr,
		},
		{
			Key: types.DTagTransferRequestStoreKey(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			Value: cdc.MustMarshal(&request),
		},
		{
			Key: types.ChainLinksStoreKey(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			Value: cdc.MustMarshal(&chainLink),
		},
		{
			Key: types.UserApplicationLinkKey(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"application",
				"username",
			),
			Value: cdc.MustMarshal(&applicationLink),
		},
		{
			Key: types.ApplicationLinkExpiringTimeKey(
				time.Date(2022, 1, 1, 0, 0, 00, 000, time.UTC),
				"client_id",
			),
			Value: []byte("client_id"),
		},
		{
			Key: types.DefaultExternalAddressKey(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos",
			),
			Value: []byte("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
		},
		{
			Key:   []byte("invalid"),
			Value: []byte("value"),
		},
	}}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"DTags", fmt.Sprintf("DTagAddressA: %s\nDTagAddressB: %s\n", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")},
		{"DTag transfer request", fmt.Sprintf("RequestA: %s\nRequestB: %s\n", request, request)},
		{"Chain link", fmt.Sprintf("ChainLinkA: %s\nChainLinkB: %s\n", chainLink, chainLink)},
		{"Application link", fmt.Sprintf("ApplicationLinkA: %s\nApplicationLinkB: %s\n", &applicationLink, &applicationLink)},
		{"Expiring Application link", fmt.Sprintf("ExpiringClientIDA: %s\nExpiringClientIDB: %s\n", "client_id", "client_id")},
		{"External address", fmt.Sprintf("ExternalAddressA: %s\nExternalAddressB: %s\n", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
