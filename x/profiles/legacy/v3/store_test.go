package v300_test

import (
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/app"
	v2 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v2"
	v3 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v3"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, legacyAminoCdc := app.MakeCodecs()
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
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
				params := v2.DefaultParams()
				paramsSpace := paramstypes.NewSubspace(cdc, legacyAminoCdc, storeKey, transientKey, "profiles")
				paramsSpace = paramsSpace.WithKeyTable(v2.ParamKeyTable())
				paramsSpace.SetParamSet(ctx, &params)

				store := ctx.KVStore(storeKey)

				applicationLinkKey := types.UserApplicationLinkKey(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"twitter",
					"user",
				)

				applicationLinkBz := v2.MustMarshalApplicationLink(cdc, v2.NewApplicationLink(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					v2.NewData("twitter", "user"),
					v2.AppLinkStateVerificationStarted,
					v2.NewOracleRequest(
						1,
						1,
						v2.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					v2.NewSuccessResult("76616c7565", "signature"), // The value should be HEX
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				))

				store.Set(applicationLinkKey, applicationLinkBz)

				block := v2.UserBlock{
					Blocker:  "blocker",
					Blocked:  "blocked",
					Reason:   "reason",
					Subspace: "",
				}
				blockBz := cdc.MustMarshal(&block)
				store.Set(v2.UserBlockStoreKey(block.Blocker, block.Subspace, block.Blocked), blockBz)

				relationship := v2.Relationship{
					Creator:   "user",
					Recipient: "recipient",
					Subspace:  "2",
				}
				relBz := cdc.MustMarshal(&relationship)
				store.Set(append(types.RelationshipsStorePrefix, 0x01), relBz)
			},
			shouldErr: false,
			check: func(ctx sdk.Context, paramSpace paramstypes.Subspace) {
				store := ctx.KVStore(storeKey)

				expectedApplicationLinkKey := types.UserApplicationLinkKey(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"twitter",
					"user",
				)

				expectedAppLink := types.NewApplicationLink(
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
				cdc.MustUnmarshal(store.Get(expectedApplicationLinkKey), &storedAppLink)
				require.Equal(t, expectedAppLink, storedAppLink)

				oldBlockKey := v2.UserBlockStoreKey("blocker", "", "blocked")
				require.False(t, store.Has(oldBlockKey))

				expectedBlock := types.NewUserBlock("blocker", "blocked", "reason", 0)
				expectedBlockKey := types.UserBlockStoreKey(expectedBlock.Blocker, expectedBlock.SubspaceID, expectedBlock.Blocked)
				require.True(t, store.Has(expectedBlockKey))

				var storedBlock types.UserBlock
				cdc.MustUnmarshal(store.Get(expectedBlockKey), &storedBlock)
				require.Equal(t, expectedBlock, storedBlock)

				oldRelationshipKey := v2.RelationshipsStoreKey("user", "2", "recipient")
				require.False(t, store.Has(oldRelationshipKey))

				expectedRelationship := types.NewRelationship("user", "recipient", 2)
				expectedRelationshipKey := types.RelationshipsStoreKey(expectedRelationship.Creator, expectedRelationship.SubspaceID, expectedRelationship.Recipient)
				require.True(t, store.Has(expectedRelationshipKey))

				var storedRelationship types.Relationship
				cdc.MustUnmarshal(store.Get(expectedRelationshipKey), &storedRelationship)
				require.Equal(t, expectedRelationship, storedRelationship)

				var storedParams types.Params
				paramSpace.GetParamSet(ctx, &storedParams)
				require.Equal(t, types.DefaultParams(), storedParams)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := testutil.DefaultContext(storeKey, transientKey)
			if tc.store != nil {
				tc.store(ctx)
			}

			newParams := types.DefaultParams()
			newParamSpace := paramstypes.NewSubspace(cdc, legacyAminoCdc, storeKey, transientKey, "profiles")
			newParamSpace = newParamSpace.WithKeyTable(types.ParamKeyTable())
			newParamSpace.SetParamSet(ctx, &newParams)

			err := v3.MigrateStore(ctx, storeKey, newParamSpace, cdc, legacyAminoCdc)
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
