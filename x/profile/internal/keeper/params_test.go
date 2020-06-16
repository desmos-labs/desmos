package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_SetNameSurnameLenParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	ctx, k := SetupTestInput()
	nsParams := models.NewNameSurnameLenParams(&min, &max)
	k.SetNameSurnameLenParams(ctx, nsParams)

	actualParams := k.GetNameSurnameLenParams(ctx)

	require.Equal(t, nsParams, actualParams)
}

func TestKeeper_GetNameSurnameLenParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	ctx, k := SetupTestInput()
	nsParams := models.NewNameSurnameLenParams(&min, &max)
	k.SetNameSurnameLenParams(ctx, nsParams)

	actualParams := k.GetNameSurnameLenParams(ctx)

	require.Equal(t, nsParams, actualParams)
}

func TestKeeper_SetMonikerLenParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	ctx, k := SetupTestInput()
	monikerParams := models.NewMonikerLenParams(&min, &max)
	k.SetMonikerLenParams(ctx, monikerParams)

	actualParams := k.GetMonikerLenParams(ctx)

	require.Equal(t, monikerParams, actualParams)
}

func TestKeeper_GetMonikerLenParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	ctx, k := SetupTestInput()
	monikerParams := models.NewMonikerLenParams(&min, &max)
	k.SetMonikerLenParams(ctx, monikerParams)

	actualParams := k.GetMonikerLenParams(ctx)

	require.Equal(t, monikerParams, actualParams)
}

func TestKeeper_SetBioLenParams(t *testing.T) {
	ctx, k := SetupTestInput()
	bioParams := models.NewBioLenParams(sdk.NewInt(100))
	k.SetBioLenParams(ctx, bioParams)

	actualParams := k.GetBioLenParams(ctx)

	require.Equal(t, bioParams, actualParams)
}

func TestKeeper_GetBioLenParams(t *testing.T) {
	ctx, k := SetupTestInput()
	bioParams := models.NewBioLenParams(sdk.NewInt(100))
	k.SetBioLenParams(ctx, bioParams)

	actualParams := k.GetBioLenParams(ctx)

	require.Equal(t, bioParams, actualParams)
}
