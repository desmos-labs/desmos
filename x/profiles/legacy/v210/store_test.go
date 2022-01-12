package v210_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/app"
	v210 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v210"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func TestStoreMigration(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	profilesKey := sdk.NewKVStoreKey("profiles")
	ctx := testutil.DefaultContext(profilesKey, sdk.NewTransientStoreKey("transient_test"))
	store := ctx.KVStore(profilesKey)

	testCases := []struct {
		name     string
		key      []byte
		oldValue []byte
		newValue []byte
	}{
		{
			name: "application link",
			key: types.UserApplicationLinkKey(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				"twitter",
				"user",
			),
			oldValue: types.MustMarshalApplicationLink(cdc, types.NewApplicationLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewData("twitter", "user"),
				types.AppLinkStateVerificationStarted,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewSuccessResult("value", "signature"),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			)),
			newValue: types.MustMarshalApplicationLink(cdc, types.NewApplicationLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewData("twitter", "user"),
				types.AppLinkStateVerificationStarted,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewSuccessResult("76616c7565", "signature"), // The value should be HEX
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			)),
		},
		{
			name: "chain link",
			key: types.ChainLinksStoreKey(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				"desmos",
				"desmos13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu",
			),
			oldValue: types.MustMarshalChainLink(cdc, types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address("desmos13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu", "desmos"),
				types.NewProof(&secp256k1.PubKey{Key: []byte{1}}, "signature", "wrong"),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			)),
			newValue: types.MustMarshalChainLink(cdc, types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address("desmos13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu", "desmos"),
				types.NewProof(&secp256k1.PubKey{Key: []byte{1}}, "signature", "77726f6e67"), // Plain text is now in HEX
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			)),
		},
	}

	// Set all the old values to the store
	for _, tc := range testCases {
		store.Set(tc.key, tc.oldValue)
	}

	// Run migrations
	err := v210.MigrateStore(ctx, profilesKey, cdc)
	require.NoError(t, err)

	// Make sure the new values are set properly
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.newValue, store.Get(tc.key))
		})
	}
}
