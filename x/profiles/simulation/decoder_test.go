package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/profiles/keeper"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/desmos-labs/desmos/x/profiles/simulation"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	profile := types.NewProfile(
		"leoDiCap",
		"",
		"Hollywood Actor. Proud environmentalist",
		types.NewPictures("", ""),
		time.Time{},
		ed25519.GenPrivKey().PubKey().Address().String(),
	)

	requests := keeper.NewWrappedDTagTransferRequests([]types.DTagTransferRequest{
		types.NewDTagTransferRequest("dtag", profile.Creator, profile.Creator),
	})

	owner := keeper.NewWrappedDTagOwner(profile.Creator)

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.ProfileStoreKey(profile.Creator),
			Value: cdc.MustMarshalBinaryBare(&profile),
		},
		{
			Key:   types.DtagStoreKey(profile.Dtag),
			Value: cdc.MustMarshalBinaryBare(&owner),
		},
		{
			Key:   types.DtagTransferRequestStoreKey(profile.Creator),
			Value: cdc.MustMarshalBinaryBare(&requests),
		},
	}}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Profile", fmt.Sprintf("ProfileA: %s\nProfileB: %s\n", profile, profile)},
		{"Address", fmt.Sprintf("AddressA: %s\nAddressB: %s\n", profile.Creator, profile.Creator)},
		{"Requests", fmt.Sprintf("RequestsA: %s\nRequestsB: %s\n", requests.Requests, requests.Requests)},
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
