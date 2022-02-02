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

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context, paramsSpace paramstypes.Subspace)
	}{
		{
			name: "valid data returns no error",
			store: func(ctx sdk.Context) {
				// setup legacy params
				params := v200.DefaultParams()
				paramsSpace := paramstypes.NewSubspace(cdc, legacyAminoCdc, profilesKey, transientKey, "profiles")
				paramsSpace = paramsSpace.WithKeyTable(v200.ParamKeyTable())
				paramsSpace.SetParamSet(ctx, &params)

				store := ctx.KVStore(profilesKey)

				applicationLinkKey := types.UserApplicationLinkKey(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"twitter",
					"user",
				)

				applicationLinkBz := v200.MustMarshalApplicationLink(cdc, v200.NewApplicationLink(
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
				))

				store.Set(applicationLinkKey, applicationLinkBz)
			},
			check: func(ctx sdk.Context, paramSpace paramstypes.Subspace) {
				store := ctx.KVStore(profilesKey)
				applicationLinkKey := types.UserApplicationLinkKey(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"twitter",
					"user",
				)

				var expectedAppLink = types.NewApplicationLink(
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
					time.Date(2022, 7, 2, 00, 00, 00, 000, time.UTC),
				)

				var storedAppLink types.ApplicationLink
				cdc.MustUnmarshal(store.Get(applicationLinkKey), &storedAppLink)
				require.Equal(t, expectedAppLink, storedAppLink)

				var storedParams types.Params
				paramSpace.GetParamSet(ctx, &storedParams)

				require.Equal(t, types.DefaultParams(), storedParams)
			},
		},
	}

	// Make sure the new values are set properly
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := testutil.DefaultContext(profilesKey, transientKey)
			if tc.store != nil {
				tc.store(ctx)
			}

			newParams := types.DefaultParams()
			newParamSpace := paramstypes.NewSubspace(cdc, legacyAminoCdc, profilesKey, transientKey, "profiles")
			newParamSpace = newParamSpace.WithKeyTable(types.ParamKeyTable())
			newParamSpace.SetParamSet(ctx, &newParams)

			err := v231.MigrateStore(ctx, profilesKey, newParamSpace, cdc, legacyAminoCdc)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tc.check != nil {
					tc.check(ctx, newParamSpace)
				}
			}
		})
	}
}
