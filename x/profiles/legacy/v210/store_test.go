package v210_test

import (
	v200 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v200"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/app"
	v210 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v210"
	v230 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v230"
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
			oldValue: v200.MustMarshalApplicationLink(cdc, v200.NewApplicationLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				v200.NewData("twitter", "user"),
				v200.AppLinkStateVerificationStarted,
				v200.NewOracleRequest(
					1,
					1,
					v200.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				v200.NewSuccessResult("value", "signature"),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			)),
			newValue: v200.MustMarshalApplicationLink(cdc, v200.NewApplicationLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				v200.NewData("twitter", "user"),
				v200.AppLinkStateVerificationStarted,
				v200.NewOracleRequest(
					1,
					1,
					v200.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				v200.NewSuccessResult("76616c7565", "signature"), // The value should be HEX
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
			oldValue: v230.MustMarshalChainLink(cdc, v230.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				v230.NewBech32Address("desmos13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu", "desmos"),
				v230.NewProof(&secp256k1.PubKey{Key: []byte{1}}, "signature", "wrong"),
				v230.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			)),
			newValue: v230.MustMarshalChainLink(cdc, v230.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				v230.NewBech32Address("desmos13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu", "desmos"),
				v230.NewProof(&secp256k1.PubKey{Key: []byte{1}}, "signature", "77726f6e67"), // Plain text is now in HEX
				v230.NewChainConfig("cosmos"),
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
