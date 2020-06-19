package models_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewNameSurnameParamsEditProposal(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)
	expectedProposal := models.EditNameSurnameParamsProposal{
		Title:             "proposal",
		Description:       "description",
		NameSurnameParams: models.NewNameSurnameLenParams(&validMin, &validMax),
	}

	actualProposal := models.NewNameSurnameParamsEditProposal(
		"proposal",
		"description",
		models.NewNameSurnameLenParams(&validMin, &validMax),
	)

	require.Equal(t, expectedProposal, actualProposal)
}

func TestNameSurnameParamsEditProposal_GetTitle(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	proposal := models.NewNameSurnameParamsEditProposal(
		"proposal",
		"description",
		models.NewNameSurnameLenParams(&validMin, &validMax),
	)

	require.Equal(t, "proposal", proposal.GetTitle())
}

func TestNameSurnameParamsEditProposal_GetDescription(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	proposal := models.NewNameSurnameParamsEditProposal(
		"proposal",
		"description",
		models.NewNameSurnameLenParams(&validMin, &validMax),
	)

	require.Equal(t, "description", proposal.GetDescription())
}

func TestNameSurnameParamsEditProposal_ProposalRoute(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	proposal := models.NewNameSurnameParamsEditProposal(
		"proposal",
		"description",
		models.NewNameSurnameLenParams(&validMin, &validMax),
	)

	require.Equal(t, "profiles", proposal.ProposalRoute())
}

func TestNameSurnameParamsEditProposal_ProposalType(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	proposal := models.NewNameSurnameParamsEditProposal(
		"proposal",
		"description",
		models.NewNameSurnameLenParams(&validMin, &validMax),
	)

	require.Equal(t, "EditNameSurnameParams", proposal.ProposalType())
}

func TestNameSurnameParamsEditProposal_String(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	proposal := models.NewNameSurnameParamsEditProposal(
		"proposal",
		"description",
		models.NewNameSurnameLenParams(&validMin, &validMax),
	)

	require.Equal(t, "Name/Surname Profiles' params edit proposal:\n  Title:       proposal\n  Description: description\n  Proposed name/Surname params lengths:\n  Min: 2\n  Max: 30\n", proposal.String())
}

func TestNameSurnameParamsEditProposal_ValidateBasic(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(500)

	invalidMin := sdk.NewInt(0)

	proposal := models.NewNameSurnameParamsEditProposal(
		"proposal",
		"description",
		models.NewNameSurnameLenParams(&validMin, &validMax),
	)

	invalidProposal := models.NewNameSurnameParamsEditProposal(
		"proposal",
		"description",
		models.NewNameSurnameLenParams(&invalidMin, &validMax),
	)

	tests := []struct {
		name     string
		proposal gov.Content
		expErr   error
	}{
		{
			name:     "Invalid proposal returns error",
			proposal: invalidProposal,
			expErr:   fmt.Errorf("invalid minimum name/surname length param: %s", invalidMin),
		},
		{
			name:     "Valid proposal returns no error",
			proposal: proposal,
			expErr:   nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.proposal.ValidateBasic()
			require.Equal(t, test.expErr, err)
		})
	}
}

func TestNewMonikerParamsEditProposal(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)
	params := models.NewMonikerLenParams(&validMin, &validMax)

	monikerProposal := models.EditMonikerParamsProposal{
		Title:         "title",
		Description:   "description",
		MonikerParams: params,
	}

	actualProposal := models.NewMonikerParamsEditProposal("title", "description", params)

	require.Equal(t, monikerProposal, actualProposal)
}

func TestMonikerParamsEditProposal_GetTitle(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)
	params := models.NewMonikerLenParams(&validMin, &validMax)

	proposal := models.NewMonikerParamsEditProposal("title", "description", params)

	require.Equal(t, "title", proposal.GetTitle())
}

func TestMonikerParamsEditProposal_GetDescription(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)
	params := models.NewMonikerLenParams(&validMin, &validMax)

	proposal := models.NewMonikerParamsEditProposal("title", "description", params)

	require.Equal(t, "description", proposal.GetDescription())
}

func TestMonikerParamsEditProposal_ProposalRoute(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)
	params := models.NewMonikerLenParams(&validMin, &validMax)

	proposal := models.NewMonikerParamsEditProposal("title", "description", params)

	require.Equal(t, "profiles", proposal.ProposalRoute())
}

func TestMonikerParamsEditProposal_ProposalType(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)
	params := models.NewMonikerLenParams(&validMin, &validMax)

	proposal := models.NewMonikerParamsEditProposal("title", "description", params)

	require.Equal(t, "EditMonikerParams", proposal.ProposalType())
}

func TestMonikerParamsEditProposal_String(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)
	params := models.NewMonikerLenParams(&validMin, &validMax)

	proposal := models.NewMonikerParamsEditProposal("title", "description", params)

	require.Equal(t, "Moniker Profiles' params edit proposal:\n  Title:       title\n  Description: description\n  Proposed moniker params lengths:\n  Min: 2\n  Max: 30\n", proposal.String())
}

func TestMonikerParamsEditProposal_ValidateBasic(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(500)

	invalidMin := sdk.NewInt(0)

	proposal := models.NewMonikerParamsEditProposal(
		"proposal",
		"description",
		models.NewMonikerLenParams(&validMin, &validMax),
	)

	invalidProposal := models.NewMonikerParamsEditProposal(
		"proposal",
		"description",
		models.NewMonikerLenParams(&invalidMin, &validMax),
	)

	tests := []struct {
		name     string
		proposal gov.Content
		expErr   error
	}{
		{
			name:     "Invalid proposal returns error",
			proposal: invalidProposal,
			expErr:   fmt.Errorf("invalid minimum moniker length param: %s", invalidMin),
		},
		{
			name:     "Valid proposal returns no error",
			proposal: proposal,
			expErr:   nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.proposal.ValidateBasic()
			require.Equal(t, test.expErr, err)
		})
	}
}

func TestNewBioParamsEditProposal(t *testing.T) {
	params := models.NewBioLenParams(sdk.NewInt(2000))

	proposal := models.EditBioParamsProposal{
		Title:       "title",
		Description: "description",
		BioParams:   params,
	}

	actualProp := models.NewBioParamsEditProposal("title", "description", proposal.BioParams)
	require.Equal(t, proposal, actualProp)
}

func TestBioParamsEditProposal_GetTitle(t *testing.T) {
	actualProp := models.NewBioParamsEditProposal(
		"title",
		"description",
		models.NewBioLenParams(sdk.NewInt(2000)),
	)

	require.Equal(t, "title", actualProp.GetTitle())
}

func TestBioParamsEditProposal_GetDescription(t *testing.T) {
	actualProp := models.NewBioParamsEditProposal(
		"title",
		"description",
		models.NewBioLenParams(sdk.NewInt(2000)),
	)

	require.Equal(t, "description", actualProp.GetDescription())
}

func TestBioParamsEditProposal_ProposalRoute(t *testing.T) {
	actualProp := models.NewBioParamsEditProposal(
		"title",
		"description",
		models.NewBioLenParams(sdk.NewInt(2000)),
	)

	require.Equal(t, "profiles", actualProp.ProposalRoute())
}

func TestBioParamsEditProposal_ProposalType(t *testing.T) {
	actualProp := models.NewBioParamsEditProposal(
		"title",
		"description",
		models.NewBioLenParams(sdk.NewInt(2000)),
	)

	require.Equal(t, "EditBioParams", actualProp.ProposalType())
}

func TestBioParamsEditProposal_String(t *testing.T) {
	actualProp := models.NewBioParamsEditProposal(
		"title",
		"description",
		models.NewBioLenParams(sdk.NewInt(2000)),
	)

	require.Equal(t, "Biography Profiles' params edit proposal:\n  Title:       title\n  Description: description\n  Proposed biography params lengths:\n  Max: 2000\n", actualProp.String())

}

func TestBioParamsEditProposal_ValidateBasic(t *testing.T) {
	invalidMax := sdk.NewInt(-1)
	proposal := models.NewBioParamsEditProposal(
		"proposal",
		"description",
		models.NewBioLenParams(sdk.NewInt(500)),
	)

	invalidProposal := models.NewBioParamsEditProposal(
		"proposal",
		"description",
		models.NewBioLenParams(invalidMax),
	)

	tests := []struct {
		name     string
		proposal gov.Content
		expErr   error
	}{
		{
			name:     "Invalid proposal returns error",
			proposal: invalidProposal,
			expErr:   fmt.Errorf("invalid max bio length param: -1"),
		},
		{
			name:     "Valid proposal returns no error",
			proposal: proposal,
			expErr:   nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.proposal.ValidateBasic()
			require.Equal(t, test.expErr, err)
		})
	}
}
