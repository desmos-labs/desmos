package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

func TestKeeper_SetHooks(t *testing.T) {
	testCases := []struct {
		name      string
		setup     func(k keeper.Keeper) keeper.Keeper
		shouldErr bool
	}{
		{
			name: "setting already set hooks returns error",
			setup: func(k keeper.Keeper) keeper.Keeper {
				return k.SetHooks(types.MultiSubspacesHooks{})
			},
			shouldErr: true,
		},
		{
			name:      "setting hooks not set works properly",
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			k := keeper.NewKeeper(nil, nil)
			if tc.setup != nil {
				k = tc.setup(k)
			}

			if tc.shouldErr {
				require.Panics(t, func() { k.SetHooks(types.MultiSubspacesHooks{}) })
			} else {
				require.NotPanics(t, func() { k.SetHooks(types.MultiSubspacesHooks{}) })
			}
		})
	}
}
