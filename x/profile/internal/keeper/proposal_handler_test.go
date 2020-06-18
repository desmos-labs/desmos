package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandleNameSurnameParamsEdit(t *testing.T) {
	validMin := sdk.NewInt(5)
	validMax := sdk.NewInt(800)

	newNSParams := models.NewNameSurnameLenParams(&validMin, &validMax)
	storedParams := models.NewParams(
		models.DefaultNameSurnameLenParams(),
		models.DefaultMonikerLenParams(),
		models.DefaultBioLenParams(),
	)

	expectedParams := models.NewParams(
		newNSParams,
		models.DefaultMonikerLenParams(),
		models.DefaultBioLenParams(),
	)

	ctx, k := SetupTestInput()
	k.SetParams(ctx, storedParams)

	proposal := models.NewNameSurnameParamsEditProposal(
		"Param proposal",
		"change params",
		newNSParams,
	)

	handler := keeper.NewEditParamsProposalHandler(k)
	err := handler(ctx, proposal)
	require.NoError(t, err)
	require.Equal(t, expectedParams, k.GetParams(ctx))
}

func TestHandleMonikerParamsEdit(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	newMonikerParams := models.NewMonikerLenParams(&validMin, &validMax)
	storedParams := models.NewParams(
		models.DefaultNameSurnameLenParams(),
		models.DefaultMonikerLenParams(),
		models.DefaultBioLenParams(),
	)

	expectedParams := models.NewParams(
		models.DefaultNameSurnameLenParams(),
		newMonikerParams,
		models.DefaultBioLenParams(),
	)

	ctx, k := SetupTestInput()
	k.SetParams(ctx, storedParams)

	proposal := models.NewMonikerParamsEditProposal(
		"Param proposal",
		"change params",
		newMonikerParams,
	)

	handler := keeper.NewEditParamsProposalHandler(k)
	err := handler(ctx, proposal)
	require.NoError(t, err)
	require.Equal(t, expectedParams, k.GetParams(ctx))
}

func TestHandleBioParamsEdit(t *testing.T) {
	newBioParams := models.NewBioLenParams(sdk.NewInt(30))

	proposal := models.NewBioParamsEditProposal(
		"Param proposal",
		"change params",
		newBioParams,
	)

	storedParams := models.NewParams(
		models.DefaultNameSurnameLenParams(),
		models.DefaultMonikerLenParams(),
		models.NewBioLenParams(sdk.NewInt(50)),
	)

	expParams := models.NewParams(
		models.DefaultNameSurnameLenParams(),
		models.DefaultMonikerLenParams(),
		newBioParams,
	)

	ctx, k := SetupTestInput()
	k.SetParams(ctx, storedParams)
	handler := keeper.NewEditParamsProposalHandler(k)
	err := handler(ctx, proposal)
	require.NoError(t, err)
	require.Equal(t, expParams, k.GetParams(ctx))
}
