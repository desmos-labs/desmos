package simulation_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/kv"

	sim "github.com/desmos-labs/desmos/x/profiles/simulation"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

var (
	privKey            = ed25519.GenPrivKey().PubKey()
	accountCreatorAddr = sdk.AccAddress(privKey.Address())
	bio                = "Hollywood Actor. Proud environmentalist"

	anotherKey      = ed25519.GenPrivKey().PubKey()
	anotherUserAddr = sdk.AccAddress(anotherKey.Address())

	profile = types.Profile{
		DTag:    "leoDiCap",
		Bio:     &bio,
		Creator: accountCreatorAddr,
	}

	relationship = types.NewMonodirectionalRelationship(accountCreatorAddr, anotherUserAddr)
)

func makeTestCodec() (cdc *codec.Codec) {
	cdc = codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	return
}

func TestDecodeStore(t *testing.T) {
	cdc := makeTestCodec()

	kvPairs := kv.Pairs{
		kv.Pair{Key: types.ProfileStoreKey(profile.Creator), Value: cdc.MustMarshalBinaryBare(&profile)},
		kv.Pair{Key: types.DtagStoreKey(profile.DTag), Value: cdc.MustMarshalBinaryBare(&profile.Creator)},
		kv.Pair{Key: types.RelationshipsStoreKey(relationship.ID), Value: cdc.MustMarshalBinaryBare(&relationship)},
		kv.Pair{Key: types.UserRelationshipsStoreKey(accountCreatorAddr), Value: cdc.MustMarshalBinaryBare(&relationship.ID)},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Profile", fmt.Sprintf("ProfileA: %s\nProfileB: %s\n", profile, profile)},
		{"Address", fmt.Sprintf("AddressA: %s\nAddressB: %s\n", profile.Creator, profile.Creator)},
		{"Relationship", fmt.Sprintf("RelationshipA: %s\nRelationshipB: %s\n", relationship, relationship)},
		{"RelationshipID", fmt.Sprintf("RelationshipIDA: %s\nRelationshipIDB: %s\n", relationship.ID, relationship.ID)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { sim.DecodeStore(cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, sim.DecodeStore(cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
