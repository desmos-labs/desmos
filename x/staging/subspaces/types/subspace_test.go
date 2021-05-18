package types_test

import (
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSubspace_WithName(t *testing.T) {
	sub := types.NewSubspace("123", "name", "", "", true, time.Unix(1, 2))

	sub.WithName("sub")

	assert.Equal(t, "sub", sub.Name)
}
