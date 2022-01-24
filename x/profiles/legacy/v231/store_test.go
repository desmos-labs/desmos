package v231_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/desmos-labs/desmos/v2/app"
	v200 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v200"
	v231 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v231"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
	"github.com/stretchr/testify/require"
)

func TestStoreMigration(t *testing.T) {
	cdc, legacyAminoCdc := app.MakeCodecs()
	profilesKey := sdk.NewKVStoreKey("profiles")
	transientKey := sdk.NewTransientStoreKey("profiles_transient")
	ctx := testutil.DefaultContext(profilesKey, transientKey)

	/// setup legacy params
	params := v200.DefaultParams()
	paramsSpace := paramstypes.NewSubspace(cdc, legacyAminoCdc, profilesKey, transientKey, "profiles")
	paramsSpace = paramsSpace.WithKeyTable(v200.ParamKeyTable())
	paramsSpace.SetParamSet(ctx, &params)

	store := ctx.KVStore(profilesKey)

	testCases := []struct {
		name           string
		key            []byte
		oldValue       []byte
		newValue       []byte
		expectedParams types.Params
	}{
		{
			name: "Store migration works correctly",
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
				v200.NewSuccessResult("76616c7565", "signature"), // The value should be HEX
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
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
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 6, 18, 00, 00, 00, 000, time.UTC),
			)),
			expectedParams: types.DefaultParams(),
		},
	}

	// Set all the old values to the store
	for _, tc := range testCases {
		store.Set(tc.key, tc.oldValue)
	}

	// set the new paramSpace for the migration
	newParams := types.DefaultParams()
	newParamSpace := paramstypes.NewSubspace(cdc, legacyAminoCdc, profilesKey, transientKey, "profiles")
	newParamSpace = newParamSpace.WithKeyTable(types.ParamKeyTable())
	newParamSpace.SetParamSet(ctx, &newParams)

	newParamSpace.GetParamSet(ctx, &newParams)

	// Run migrations
	err := v231.MigrateStore(ctx, profilesKey, newParamSpace, cdc, legacyAminoCdc)
	require.NoError(t, err)

	// Make sure the new values are set properly
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.newValue, store.Get(tc.key))
			require.Equal(t, tc.expectedParams, newParams)
		})
	}
}
