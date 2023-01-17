package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func TestKeeper_SetHooks(t *testing.T) {
	testCases := []struct {
		name      string
		hooks     types.SubspacesHooks
		shouldErr bool
	}{
		{
			name:      "setting already set hooks returns error",
			hooks:     types.MultiSubspacesHooks{},
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
			subspaceKeeper := keeper.NewKeeper(nil, nil, nil, nil)
			k := subspaceKeeper.SetHooks(tc.hooks)

			if tc.shouldErr {
				require.Panics(t, func() { k.SetHooks(types.MultiSubspacesHooks{}) })
			} else {
				require.NotPanics(t, func() { k.SetHooks(types.MultiSubspacesHooks{}) })
			}
		})
	}
}
