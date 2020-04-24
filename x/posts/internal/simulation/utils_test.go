package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/simulation"
	"github.com/stretchr/testify/require"
)

func TestRandomEmoji(t *testing.T) {
	r := rand.New(rand.NewSource(5577006791947779410))
	emoji := simulation.RandomEmoji(r)
	require.NotEmpty(t, emoji.Shortcodes)
}
