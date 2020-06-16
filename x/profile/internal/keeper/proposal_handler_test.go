package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandleNameSurnameParamsEdit(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	storedMin := sdk.NewInt(3)
	storedMax := sdk.NewInt(50)

	completeNsParams := models.NewNameSurnameLenParams(&validMin, &validMax)
	onlyMinParam := models.NewNameSurnameLenParams(&validMin, nil)
	onlyMaxParam := models.NewNameSurnameLenParams(nil, &validMax)

	storedParams := models.NewNameSurnameLenParams(&storedMin, &storedMax)

	tests := []struct {
		name          string
		proposal      gov.Content
		storedParams  models.NameSurnameLenParams
		expParameters models.NameSurnameLenParams
	}{
		{
			name: "Proposal changes both parameters",
			proposal: models.NewNameSurnameParamsEditProposal(
				"Param proposal",
				"change params",
				completeNsParams,
			),
			storedParams:  storedParams,
			expParameters: completeNsParams,
		},
		{
			name: "Proposal changes only min parameter",
			proposal: models.NewNameSurnameParamsEditProposal(
				"Param proposal",
				"change params",
				onlyMinParam,
			),
			storedParams:  storedParams,
			expParameters: models.NewNameSurnameLenParams(&validMin, &storedMax),
		},
		{
			name: "Proposal changes only max parameter",
			proposal: models.NewNameSurnameParamsEditProposal(
				"Param proposal",
				"change params",
				onlyMaxParam,
			),
			storedParams:  storedParams,
			expParameters: models.NewNameSurnameLenParams(&storedMin, &validMax),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			k.SetNameSurnameLenParams(ctx, test.storedParams)

			handler := keeper.NewEditParamsProposalHandler(k)
			err := handler(ctx, test.proposal)
			require.NoError(t, err)
			require.Equal(t, test.expParameters, k.GetNameSurnameLenParams(ctx))
		})
	}
}

func TestHandleMonikerParamsEdit(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	storedMin := sdk.NewInt(3)
	storedMax := sdk.NewInt(50)

	completeMonikerParams := models.NewMonikerLenParams(&validMin, &validMax)
	onlyMinParam := models.NewMonikerLenParams(&validMin, nil)
	onlyMaxParam := models.NewMonikerLenParams(nil, &validMax)

	storedParams := models.NewMonikerLenParams(&storedMin, &storedMax)

	tests := []struct {
		name          string
		proposal      gov.Content
		storedParams  models.MonikerLenParams
		expParameters models.MonikerLenParams
	}{
		{
			name: "Proposal changes both parameters",
			proposal: models.NewMonikerParamsEditProposal(
				"Param proposal",
				"change params",
				completeMonikerParams,
			),
			storedParams:  storedParams,
			expParameters: completeMonikerParams,
		},
		{
			name: "Proposal changes only min parameter",
			proposal: models.NewMonikerParamsEditProposal(
				"Param proposal",
				"change params",
				onlyMinParam,
			),
			storedParams:  storedParams,
			expParameters: models.NewMonikerLenParams(&validMin, &storedMax),
		},
		{
			name: "Proposal changes only max parameter",
			proposal: models.NewMonikerParamsEditProposal(
				"Param proposal",
				"change params",
				onlyMaxParam,
			),
			storedParams:  storedParams,
			expParameters: models.NewMonikerLenParams(&storedMin, &validMax),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			k.SetMonikerLenParams(ctx, test.storedParams)

			handler := keeper.NewEditParamsProposalHandler(k)
			err := handler(ctx, test.proposal)
			require.NoError(t, err)
			require.Equal(t, test.expParameters, k.GetMonikerLenParams(ctx))
		})
	}
}

func TestHandleBioParamsEdit(t *testing.T) {
	tests := []struct {
		name          string
		proposal      gov.Content
		storedParams  models.BioLenParams
		expParameters models.BioLenParams
	}{
		{
			name: "Proposal changes parameter",
			proposal: models.NewBioParamsEditProposal(
				"Param proposal",
				"change params",
				models.NewBioLenParams(sdk.NewInt(30)),
			),
			storedParams:  models.NewBioLenParams(sdk.NewInt(50)),
			expParameters: models.NewBioLenParams(sdk.NewInt(30)),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			k.SetBioLenParams(ctx, test.storedParams)

			handler := keeper.NewEditParamsProposalHandler(k)
			err := handler(ctx, test.proposal)
			require.NoError(t, err)
			require.Equal(t, test.expParameters, k.GetBioLenParams(ctx))
		})
	}
}
