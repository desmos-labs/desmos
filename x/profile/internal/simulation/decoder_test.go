package simulation

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/kv"
)

var (
	privKey            = ed25519.GenPrivKey().PubKey()
	accountCreatorAddr = sdk.AccAddress(privKey.Address())
	name               = "leo"
	surname            = "Di Caprio"
	bio                = "Hollywood Actor. Proud environmentalist"

	profile = models.Profile{
		Name:    &name,
		Surname: &surname,
		Moniker: "leoDiCap",
		Bio:     &bio,
		Creator: accountCreatorAddr,
	}
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
		kv.Pair{Key: models.ProfileStoreKey(profile.Creator), Value: cdc.MustMarshalBinaryBare(&profile)},
		kv.Pair{Key: models.MonikerStoreKey(profile.Moniker), Value: cdc.MustMarshalBinaryBare(&profile.Creator)},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Profile", fmt.Sprintf("ProfileA: %s\nProfileB: %s\n", profile, profile)},
		{"Moniker", fmt.Sprintf("AddressA: %s\nAddressB: %s\n", profile.Creator, profile.Creator)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { DecodeStore(cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, DecodeStore(cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
