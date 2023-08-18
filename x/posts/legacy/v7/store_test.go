package v7_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/testutil/storetesting"
	v7 "github.com/desmos-labs/desmos/v6/x/posts/legacy/v7"
	"github.com/desmos-labs/desmos/v6/x/posts/types"
)

func TestMigrate(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	keys := sdk.NewKVStoreKeys(types.StoreKey)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "post set new owner field to author correctly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				post := types.NewPost(
					1,
					0,
					2,
					"External id",
					"Text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					nil,
					[]string{"generic"},
					[]types.PostReference{
						types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"",
				)
				store.Set(types.PostStoreKey(post.SubspaceID, post.ID), cdc.MustMarshal(&post))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var stored types.Post
				cdc.MustUnmarshal(store.Get(types.PostStoreKey(1, 2)), &stored)

				require.Equal(t, types.NewPost(
					1,
					0,
					2,
					"External id",
					"Text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					nil,
					[]string{"generic"},
					[]types.PostReference{
						types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
				), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := storetesting.BuildContext(keys, nil, nil)

			if tc.store != nil {
				tc.store(ctx)
			}

			err := v7.MigrateStore(ctx, keys[types.StoreKey], cdc)
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
