package models_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDefaultParams(t *testing.T) {
	nsMin := sdk.NewInt(2)
	nsMax := sdk.NewInt(1000)
	nameSurnameParams := models.NewNameSurnameLenParams(&nsMin, &nsMax)
	mMin := sdk.NewInt(2)
	mMax := sdk.NewInt(30)
	monikerParams := models.NewMonikerLenParams(&mMin, &mMax)
	bioParams := models.NewBioLenParams(sdk.NewInt(1000))

	params := models.NewParams(nameSurnameParams, monikerParams, bioParams)

	require.Equal(t, params, models.DefaultParams())
}

func TestParams_WithNameSurnameParams(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	storedMin := sdk.NewInt(3)
	storedMax := sdk.NewInt(50)

	completeNsParams := models.NewNameSurnameLenParams(&validMin, &validMax)
	onlyMinParam := models.NewNameSurnameLenParams(&validMin, nil)
	onlyMaxParam := models.NewNameSurnameLenParams(nil, &validMax)

	storedNSParams := models.NewNameSurnameLenParams(&storedMin, &storedMax)

	params := models.NewParams(storedNSParams, models.DefaultMonikerLenParams(), models.DefaultBioLenParams())

	tests := []struct {
		name          string
		params        models.NameSurnameLengths
		storedParams  models.Params
		expParameters models.NameSurnameLengths
	}{
		{
			name:          "Proposal changes both parameters",
			params:        completeNsParams,
			storedParams:  params,
			expParameters: completeNsParams,
		},
		{
			name:          "Proposal changes only min parameter",
			params:        onlyMinParam,
			storedParams:  params,
			expParameters: models.NewNameSurnameLenParams(&validMin, &storedMax),
		},
		{
			name:          "Proposal changes only max parameter",
			params:        onlyMaxParam,
			storedParams:  params,
			expParameters: models.NewNameSurnameLenParams(&storedMin, &validMax),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actualParams := test.storedParams.WithNameSurnameParams(test.params)
			require.Equal(t, test.expParameters, actualParams.NameSurnameLengths)
		})
	}
}

func TestParams_WithMonikerParams(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	storedMin := sdk.NewInt(3)
	storedMax := sdk.NewInt(50)

	completeMonikerParams := models.NewMonikerLenParams(&validMin, &validMax)
	onlyMinParam := models.NewMonikerLenParams(&validMin, nil)
	onlyMaxParam := models.NewMonikerLenParams(nil, &validMax)

	storedParams := models.NewMonikerLenParams(&storedMin, &storedMax)

	params := models.NewParams(models.DefaultNameSurnameLenParams(), storedParams, models.DefaultBioLenParams())

	tests := []struct {
		name          string
		params        models.MonikerLengths
		storedParams  models.Params
		expParameters models.MonikerLengths
	}{
		{
			name:          "Proposal changes both parameters",
			params:        completeMonikerParams,
			storedParams:  params,
			expParameters: completeMonikerParams,
		},
		{
			name:          "Proposal changes only min parameter",
			params:        onlyMinParam,
			storedParams:  params,
			expParameters: models.NewMonikerLenParams(&validMin, &storedMax),
		},
		{
			name:          "Proposal changes only max parameter",
			params:        onlyMaxParam,
			storedParams:  params,
			expParameters: models.NewMonikerLenParams(&storedMin, &validMax),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actualParams := test.storedParams.WithMonikerParams(test.params)
			require.Equal(t, test.expParameters, actualParams.MonikerLengths)
		})
	}
}

func TestValidateParams(t *testing.T) {
	invalidNameMin := sdk.NewInt(1)
	validNameMax := sdk.NewInt(1000)
	validMonikerMin := sdk.NewInt(2)
	invalidMonikerMax := sdk.NewInt(-30)

	tests := []struct {
		name   string
		params models.Params
		expErr error
	}{
		{
			name:   "Invalid min name/surname param returns error",
			params: models.NewParams(models.NewNameSurnameLenParams(&invalidNameMin, &validNameMax), models.DefaultMonikerLenParams(), models.DefaultBioLenParams()),
			expErr: fmt.Errorf("invalid minimum name/surname length param: 1"),
		},
		{
			name:   "Invalid max param return error",
			params: models.NewParams(models.DefaultNameSurnameLenParams(), models.NewMonikerLenParams(&validMonikerMin, &invalidMonikerMax), models.DefaultBioLenParams()),
			expErr: fmt.Errorf("invalid max moniker length param: -30"),
		},
		{
			name:   "Invalid max param returns error",
			params: models.NewParams(models.DefaultNameSurnameLenParams(), models.DefaultMonikerLenParams(), models.NewBioLenParams(sdk.NewInt(-1000))),
			expErr: fmt.Errorf("invalid max bio length param: -1000"),
		},
		{
			name:   "Valid params return no error",
			params: models.NewParams(models.DefaultNameSurnameLenParams(), models.DefaultMonikerLenParams(), models.DefaultBioLenParams()),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.params.Validate())
		})
	}
}

func TestDefaultNameSurnameLenParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	nameSurnameParams := models.NewNameSurnameLenParams(&min, &max)
	defaultNSParams := models.DefaultNameSurnameLenParams()
	require.Equal(t, defaultNSParams, nameSurnameParams)
}

func TestDefaultMonikerLenParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(30)
	monikerParams := models.NewMonikerLenParams(&min, &max)
	defaultMonikerParams := models.DefaultMonikerLenParams()
	require.Equal(t, defaultMonikerParams, monikerParams)
}

func TestDefaultBioLenParams(t *testing.T) {
	bioParams := models.NewBioLenParams(sdk.NewInt(1000))
	defaultBioParams := models.DefaultBioLenParams()
	require.Equal(t, defaultBioParams, bioParams)
}

func TestNameSurnameLenParams_String(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	nameSurnameParams := models.NewNameSurnameLenParams(&min, &max)
	actual := nameSurnameParams.String()
	require.Equal(t, "{\"min_name_surname_len\":\"2\",\"max_name_surname_len\":\"1000\"}", actual)
}

func TestMonikerLenParams_String(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(30)
	monikerParams := models.NewMonikerLenParams(&min, &max)
	actual := monikerParams.String()
	require.Equal(t, "{\"min_moniker_len\":\"2\",\"max_moniker_len\":\"30\"}", actual)
}

func TestBioLenParams_String(t *testing.T) {
	bioParams := models.NewBioLenParams(sdk.NewInt(1000))
	actual := bioParams.String()
	require.Equal(t, "{\"max_bio_len\":\"1000\"}", actual)
}

func TestValidateNameSurnameLenParams(t *testing.T) {
	invalidNameMin := sdk.NewInt(1)
	invalidNameMax := sdk.NewInt(-10)
	validNameMin := sdk.NewInt(2)
	validNameMax := sdk.NewInt(1000)

	tests := []struct {
		name   string
		params interface{}
		expErr error
	}{
		{
			name:   "Invalid min param returns error",
			params: models.NewNameSurnameLenParams(&invalidNameMin, &validNameMax),
			expErr: fmt.Errorf("invalid minimum name/surname length param: 1"),
		},
		{
			name:   "Invalid max param returns error",
			params: models.NewNameSurnameLenParams(&validNameMin, &invalidNameMax),
			expErr: fmt.Errorf("invalid max name/surname length param: -10"),
		},
		{
			name:   "Valid params returns no error",
			params: models.NewNameSurnameLenParams(&validNameMin, &validNameMax),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, models.ValidateNameSurnameLenParams(test.params))
		})
	}
}

func TestValidateMonikerLenParams(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	invalidMin := sdk.NewInt(1)
	invalidMax := sdk.NewInt(-30)

	tests := []struct {
		name   string
		params interface{}
		expErr error
	}{
		{
			name:   "Invalid min param return error",
			params: models.NewMonikerLenParams(&invalidMin, &validMax),
			expErr: fmt.Errorf("invalid minimum moniker length param: 1"),
		},
		{
			name:   "Invalid max param return error",
			params: models.NewMonikerLenParams(&validMin, &invalidMax),
			expErr: fmt.Errorf("invalid max moniker length param: -30"),
		},
		{
			name:   "Valid params returns no error",
			params: models.NewMonikerLenParams(&validMin, &validMax),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, models.ValidateMonikerLenParams(test.params))
		})
	}
}

func TestValidateBioLenParams(t *testing.T) {
	tests := []struct {
		name   string
		params interface{}
		expErr error
	}{
		{
			name:   "Invalid max param returns error",
			params: models.NewBioLenParams(sdk.NewInt(-1000)),
			expErr: fmt.Errorf("invalid max bio length param: -1000"),
		},
		{
			name:   "Valid params returns no error",
			params: models.NewBioLenParams(sdk.NewInt(1000)),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, models.ValidateBioLenParams(test.params))
		})
	}
}
