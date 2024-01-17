package v3_test

import (
	"testing"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	"github.com/stretchr/testify/require"

	v2 "github.com/desmos-labs/desmos/v6/x/posts/legacy/v2"
	v3 "github.com/desmos-labs/desmos/v6/x/posts/legacy/v3"
	v4 "github.com/desmos-labs/desmos/v6/x/posts/legacy/v4"

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

	lastEditDate := time.Date(2022, 11, 30, 8, 0, 0, 0, time.UTC)
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "posts are migrated properly",
			store: func(ctx sdk.Context) {
				post := v2.NewPost(
					1,
					0,
					1,
					"External id",
					"This is a post text that does not contain any useful information",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
					v2.NewEntities(
						[]v2.Tag{
							v2.NewTag(1, 3, "tag"),
						},
						[]v2.Tag{
							v2.NewTag(4, 6, "tag"),
						},
						[]v2.Url{
							v2.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					[]v2.PostReference{
						v2.NewPostReference(v2.POST_REFERENCE_TYPE_QUOTE, 1, 0),
					},
					v2.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					&lastEditDate,
				)

				store := ctx.KVStore(keys[types.StoreKey])
				store.Set(types.PostStoreKey(1, 1), cdc.MustMarshal(&post))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var stored v4.Post
				err := cdc.Unmarshal(store.Get(types.PostStoreKey(1, 1)), &stored)
				require.NoError(t, err)
				require.Equal(t, v4.NewPost(
					1,
					0,
					1,
					"External id",
					"This is a post text that does not contain any useful information",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
					v4.NewEntities(
						[]v4.TextTag{
							v4.NewTextTag(1, 3, "tag"),
						},
						[]v4.TextTag{
							v4.NewTextTag(4, 6, "tag"),
						},
						[]v4.Url{
							v4.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					nil,
					[]v4.PostReference{
						v4.NewPostReference(v4.POST_REFERENCE_TYPE_QUOTE, 1, 0),
					},
					v4.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					&lastEditDate,
				), stored)
			},
		},
		{
			name: "attachments are migrated properly",
			store: func(ctx sdk.Context) {
				contentAny, err := codectypes.NewAnyWithValue(&v2.Media{
					Uri:      "https://example.com?image=hello.png",
					MimeType: "image/png",
				})
				require.NoError(t, err)

				attachment := v2.Attachment{
					SubspaceID: 1,
					SectionID:  0,
					PostID:     1,
					ID:         1,
					Content:    contentAny,
				}
				store := ctx.KVStore(keys[types.StoreKey])
				store.Set(types.AttachmentStoreKey(1, 1, 1), cdc.MustMarshal(&attachment))
			},
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var stored v4.Attachment
				err := cdc.Unmarshal(store.Get(types.AttachmentStoreKey(1, 1, 1)), &stored)
				require.NoError(t, err)
				require.Equal(t, v4.NewAttachment(
					1,
					1,
					1,
					v4.NewMedia(
						"https://example.com?image=hello.png",
						"image/png",
					),
				), stored)
			},
		},
		{
			name: "user answers are migrated properly",
			store: func(ctx sdk.Context) {
				answer := v2.UserAnswer{
					SubspaceID:     1,
					SectionID:      0,
					PostID:         1,
					PollID:         1,
					AnswersIndexes: []uint32{0, 1},
					User:           "cosmos1xkahvzfyt4hsu265733ndwydv9pazqgfu586nw",
				}
				store := ctx.KVStore(keys[types.StoreKey])
				store.Set(types.PollAnswerStoreKey(1, 1, 1, "cosmos1xkahvzfyt4hsu265733ndwydv9pazqgfu586nw"), cdc.MustMarshal(&answer))
			},
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var stored v4.UserAnswer
				err := cdc.Unmarshal(store.Get(types.PollAnswerStoreKey(1, 1, 1, "cosmos1xkahvzfyt4hsu265733ndwydv9pazqgfu586nw")), &stored)
				require.NoError(t, err)
				require.Equal(t, v4.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{0, 1},
					"cosmos1xkahvzfyt4hsu265733ndwydv9pazqgfu586nw",
				), stored)
			},
		},
		{
			name: "attachments are migrated properly",
			store: func(ctx sdk.Context) {
				contentAny, err := codectypes.NewAnyWithValue(&v2.Media{
					Uri:      "https://example.com?image=hello.png",
					MimeType: "image/png",
				})
				require.NoError(t, err)

				attachment := v2.Attachment{
					SubspaceID: 1,
					SectionID:  0,
					PostID:     1,
					ID:         1,
					Content:    contentAny,
				}
				store := ctx.KVStore(keys[types.StoreKey])
				store.Set(types.AttachmentStoreKey(1, 1, 1), cdc.MustMarshal(&attachment))
			},
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var stored v4.Attachment
				err := cdc.Unmarshal(store.Get(types.AttachmentStoreKey(1, 1, 1)), &stored)
				require.NoError(t, err)
				require.Equal(t, v4.NewAttachment(
					1,
					1,
					1,
					v4.NewMedia(
						"https://example.com?image=hello.png",
						"image/png",
					),
				), stored)
			},
		},
		{
			name: "user answers are migrated properly",
			store: func(ctx sdk.Context) {
				answer := v2.UserAnswer{
					SubspaceID:     1,
					SectionID:      0,
					PostID:         1,
					PollID:         1,
					AnswersIndexes: []uint32{0, 1},
					User:           "cosmos1xkahvzfyt4hsu265733ndwydv9pazqgfu586nw",
				}
				store := ctx.KVStore(keys[types.StoreKey])
				store.Set(types.PollAnswerStoreKey(1, 1, 1, "cosmos1xkahvzfyt4hsu265733ndwydv9pazqgfu586nw"), cdc.MustMarshal(&answer))
			},
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var stored v4.UserAnswer
				err := cdc.Unmarshal(store.Get(types.PollAnswerStoreKey(1, 1, 1, "cosmos1xkahvzfyt4hsu265733ndwydv9pazqgfu586nw")), &stored)
				require.NoError(t, err)
				require.Equal(t, v4.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{0, 1},
					"cosmos1xkahvzfyt4hsu265733ndwydv9pazqgfu586nw",
				), stored)
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

			err := v3.MigrateStore(ctx, keys[types.StoreKey], cdc)
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
