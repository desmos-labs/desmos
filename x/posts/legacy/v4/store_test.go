package v4_test

import (
	"testing"
	"time"

	v4 "github.com/desmos-labs/desmos/v6/x/posts/legacy/v4"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/testutil/storetesting"
	"github.com/desmos-labs/desmos/v6/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(paramstypes.StoreKey, subspacestypes.StoreKey, types.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "user answers are deleted properly for tallied poll",
			store: func(ctx sdk.Context) {
				poll := types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					types.NewPollTallyResults([]types.PollTallyResults_AnswerResult{
						types.NewAnswerResult(0, 1),
					}),
				))

				userAnswer := types.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{1},
					"cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw",
				)

				store := ctx.KVStore(keys[types.StoreKey])
				store.Set(types.AttachmentStoreKey(1, 1, 1), cdc.MustMarshal(&poll))
				store.Set(types.PollAnswerStoreKey(1, 1, 1, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"), cdc.MustMarshal(&userAnswer))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var stored types.Attachment
				err := cdc.Unmarshal(store.Get(types.AttachmentStoreKey(1, 1, 1)), &stored)
				require.NoError(t, err)
				require.Equal(t, types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					types.NewPollTallyResults([]types.PollTallyResults_AnswerResult{
						types.NewAnswerResult(0, 1),
					}),
				)), stored)

				require.False(t, store.Has(types.PollAnswerStoreKey(1, 1, 1, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw")))
			},
		},
		{
			name: "user answers are not deleted for not tallied poll",
			store: func(ctx sdk.Context) {
				poll := types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					nil,
				))

				userAnswer := types.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{1},
					"cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw",
				)

				store := ctx.KVStore(keys[types.StoreKey])
				store.Set(types.AttachmentStoreKey(1, 1, 1), cdc.MustMarshal(&poll))
				store.Set(types.PollAnswerStoreKey(1, 1, 1, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"), cdc.MustMarshal(&userAnswer))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var stored types.Attachment
				err := cdc.Unmarshal(store.Get(types.AttachmentStoreKey(1, 1, 1)), &stored)
				require.NoError(t, err)
				require.Equal(t, types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					nil,
				)), stored)

				require.True(t, store.Has(types.PollAnswerStoreKey(1, 1, 1, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw")))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := storetesting.BuildContext(keys, tKeys, memKeys)
			if tc.store != nil {
				tc.store(ctx)
			}

			err := v4.MigrateStore(ctx, keys[types.StoreKey], cdc)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
