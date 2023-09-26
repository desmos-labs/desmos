package v5_test

import (
	"testing"
	"time"

	v4 "github.com/desmos-labs/desmos/v6/x/posts/legacy/v4"
	v5 "github.com/desmos-labs/desmos/v6/x/posts/legacy/v5"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/testutil/storetesting"
	"github.com/desmos-labs/desmos/v6/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	// Build all the necessary keys
	keys := storetypes.NewKVStoreKeys(paramstypes.StoreKey, subspacestypes.StoreKey, types.StoreKey)
	tKeys := storetypes.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "media attachments are still valid even without migration",
			store: func(ctx sdk.Context) {
				media := v4.NewAttachment(1, 1, 1, v4.NewMedia("https://example.com/image.png", "image/png"))

				store := ctx.KVStore(keys[types.StoreKey])
				store.Set(types.AttachmentStoreKey(1, 1, 1), cdc.MustMarshal(&media))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var stored types.Attachment
				err := cdc.Unmarshal(store.Get(types.AttachmentStoreKey(1, 1, 1)), &stored)
				require.NoError(t, err)
				require.Equal(t, types.NewAttachment(1, 1, 1, types.NewMedia(
					"https://example.com/image.png",
					"image/png",
				)), stored)
			},
		},
		{
			name: "poll provided answers attachments are migrated properly",
			store: func(ctx sdk.Context) {
				poll := v4.NewAttachment(1, 1, 1, v4.NewPoll(
					"What animal is best?",
					[]v4.Poll_ProvidedAnswer{
						v4.NewProvidedAnswer("Cat", []v4.Attachment{
							v4.NewAttachment(1, 1, 1, v4.NewMedia("ftp://user:password@example.com/cat.png", "image/png")),
						}),
						v4.NewProvidedAnswer("Dog", []v4.Attachment{
							v4.NewAttachment(1, 1, 1, v4.NewPoll(
								"Is this a nested poll?",
								[]v4.Poll_ProvidedAnswer{
									v4.NewProvidedAnswer("Yes", nil),
									v4.NewProvidedAnswer("No", nil),
								},
								time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
								false,
								false,
								nil,
							)),
						}),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					v4.NewPollTallyResults([]v4.PollTallyResults_AnswerResult{
						v4.NewAnswerResult(0, 1),
					}),
				))

				store := ctx.KVStore(keys[types.StoreKey])
				store.Set(types.AttachmentStoreKey(1, 1, 1), cdc.MustMarshal(&poll))
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
						types.NewProvidedAnswer("Cat", []types.AttachmentContent{
							types.NewMedia("ftp://user:password@example.com/cat.png", "image/png"),
						}),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					types.NewPollTallyResults([]types.PollTallyResults_AnswerResult{
						types.NewAnswerResult(0, 1),
					}),
				)), stored)
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

			err := v5.MigrateStore(ctx, keys[types.StoreKey], cdc)
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
