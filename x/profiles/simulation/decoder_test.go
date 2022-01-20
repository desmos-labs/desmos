package simulation_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/app"

	"github.com/desmos-labs/desmos/v2/x/profiles/simulation"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
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

	relationship := types.NewRelationship(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	userBlock := types.NewUserBlock(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"reason",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	clientID := "client_id"

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
			Key: types.RelationshipsStoreKey(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			Value: cdc.MustMarshal(&relationship),
		},
		{
			Key: types.UserBlockStoreKey(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			Value: cdc.MustMarshal(&userBlock),
		},
		{
			Key: types.ApplicationLinkExpiringTimeKey(
				time.Date(2022, 1, 1, 0, 0, 00, 000, time.UTC),
				clientID,
			),
			Value: []byte(clientID),
		},
	}}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"DTags", fmt.Sprintf("DTagAddressA: %s\nDTagAddressB: %s\n", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")},
		{"DTag transfer request", fmt.Sprintf("RequestA: %s\nRequestB: %s\n", request, request)},
		{"Relationship", fmt.Sprintf("Relationships A: %s\nRelationships B: %s\n", relationship, relationship)},
		{"User block", fmt.Sprintf("User block A: %s\nUser block B: %s\n", userBlock, userBlock)},
		{"ClientID", fmt.Sprintf("Client ID A: %s\nClient ID B: %s\n", clientID, clientID)},
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
