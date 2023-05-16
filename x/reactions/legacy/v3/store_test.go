package v3_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/app"
	"github.com/desmos-labs/desmos/v5/testutil/storetesting"
	poststypes "github.com/desmos-labs/desmos/v5/x/posts/types"
	v3 "github.com/desmos-labs/desmos/v5/x/reactions/legacy/v3"
	"github.com/desmos-labs/desmos/v5/x/reactions/testutil"
	"github.com/desmos-labs/desmos/v5/x/reactions/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(types.StoreKey)

	// Mocks initializations
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pk := testutil.NewMockPostsKeeper(ctrl)

	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "duplicated reactions does not exist works properly",
			setup: func() {
				pk.EXPECT().
					IteratePosts(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				reaction := types.NewReaction(1, 1, 1, types.NewFreeTextValue("test"), "author")
				store.Set(types.ReactionStoreKey(1, 1, 1), cdc.MustMarshal(&reaction))

				duplicatedReaction := types.NewReaction(1, 1, 2, types.NewRegisteredReactionValue(1), "author")
				store.Set(types.ReactionStoreKey(1, 1, 2), cdc.MustMarshal(&duplicatedReaction))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				require.True(t, store.Has(types.ReactionStoreKey(1, 1, 1)))
				require.True(t, store.Has(types.ReactionStoreKey(1, 1, 2)))
			},
		},
		{
			name: "delete duplicated reactions properly -- free text",
			setup: func() {
				pk.EXPECT().
					IteratePosts(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				reaction := types.NewReaction(1, 1, 1, types.NewFreeTextValue("test"), "author")
				store.Set(types.ReactionStoreKey(1, 1, 1), cdc.MustMarshal(&reaction))

				duplicatedReaction := types.NewReaction(1, 1, 2, types.NewFreeTextValue("test"), "author")
				store.Set(types.ReactionStoreKey(1, 1, 2), cdc.MustMarshal(&duplicatedReaction))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				require.True(t, store.Has(types.ReactionStoreKey(1, 1, 1)))
				require.False(t, store.Has(types.ReactionStoreKey(1, 1, 2)))
			},
		},
		{
			name: "delete duplicated reactions properly -- registered reaction",
			setup: func() {
				pk.EXPECT().
					IteratePosts(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				reaction := types.NewReaction(1, 1, 1, types.NewRegisteredReactionValue(1), "author")
				store.Set(types.ReactionStoreKey(1, 1, 1), cdc.MustMarshal(&reaction))

				duplicatedReaction := types.NewReaction(1, 1, 2, types.NewRegisteredReactionValue(1), "author")
				store.Set(types.ReactionStoreKey(1, 1, 2), cdc.MustMarshal(&duplicatedReaction))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				require.True(t, store.Has(types.ReactionStoreKey(1, 1, 1)))
				require.False(t, store.Has(types.ReactionStoreKey(1, 1, 2)))
			},
		},
		{
			name: "fix missing next id properly - without reactions",
			setup: func() {
				posts := []poststypes.Post{
					poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
				}

				pk.EXPECT().
					IteratePosts(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(post poststypes.Post) (stop bool)) {
						for _, post := range posts {
							fn(post)
						}
					})
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				require.True(t, store.Has(types.NextReactionIDStoreKey(1, 1)))
				require.Equal(t, store.Get(types.NextReactionIDStoreKey(1, 1)), types.GetReactionIDBytes(1))
			},
		},
		{
			name: "fix missing next id properly - with reactions",
			setup: func() {
				posts := []poststypes.Post{
					poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
				}

				pk.EXPECT().
					IteratePosts(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(post poststypes.Post) (stop bool)) {
						for _, post := range posts {
							fn(post)
						}
					})
			},
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				reaction := types.NewReaction(1, 1, 1, types.NewRegisteredReactionValue(1), "author")
				store.Set(types.ReactionStoreKey(1, 1, 1), cdc.MustMarshal(&reaction))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				require.True(t, store.Has(types.NextReactionIDStoreKey(1, 1)))
				require.Equal(t, store.Get(types.NextReactionIDStoreKey(1, 1)), types.GetReactionIDBytes(2))
			},
		},
		{
			name: "no missing next id works properly",
			setup: func() {
				posts := []poststypes.Post{
					poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
				}

				pk.EXPECT().
					IteratePosts(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(post poststypes.Post) (stop bool)) {
						for _, post := range posts {
							fn(post)
						}
					})
			},
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				reaction := types.NewReaction(1, 1, 1, types.NewRegisteredReactionValue(1), "author")
				store.Set(types.ReactionStoreKey(1, 1, 1), cdc.MustMarshal(&reaction))

				store.Set(types.NextReactionIDStoreKey(1, 1), types.GetReactionIDBytes(2))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				require.True(t, store.Has(types.NextReactionIDStoreKey(1, 1)))
				require.Equal(t, store.Get(types.NextReactionIDStoreKey(1, 1)), types.GetReactionIDBytes(2))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup()
			}

			ctx := storetesting.BuildContext(keys, nil, nil)
			if tc.store != nil {
				tc.store(ctx)
			}

			err := v3.MigrateStore(ctx, keys[types.StoreKey], pk, cdc)
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
