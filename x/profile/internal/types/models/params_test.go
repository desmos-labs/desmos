package models_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDefaultNameSurnameLenParams(t *testing.T) {
	nameSurnameParams := models.NewNameSurnameLenParams(sdk.NewInt(2), sdk.NewInt(1000))
	defaultNSParams := models.DefaultNameSurnameLenParams()
	require.Equal(t, defaultNSParams, nameSurnameParams)
}

func TestDefaultMonikerLenParams(t *testing.T) {
	monikerParams := models.NewMonikerLenParams(sdk.NewInt(2), sdk.NewInt(30))
	defaultMonikerParams := models.DefaultMonikerLenParams()
	require.Equal(t, defaultMonikerParams, monikerParams)
}

func TestDefaultBioLenParams(t *testing.T) {
	bioParams := models.NewBioLenParams(sdk.NewInt(1000))
	defaultBioParams := models.DefaultBioLenParams()
	require.Equal(t, defaultBioParams, bioParams)
}

func TestNameSurnameLenParams_String(t *testing.T) {
	nameSurnameParams := models.NewNameSurnameLenParams(sdk.NewInt(2), sdk.NewInt(1000))
	actual := nameSurnameParams.String()
	require.Equal(t, "{\"min_name_surname_len\":\"2\",\"max_name_surname_len\":\"1000\"}", actual)
}

func TestMonikerLenParams_String(t *testing.T) {
	monikerParams := models.NewMonikerLenParams(sdk.NewInt(2), sdk.NewInt(30))
	actual := monikerParams.String()
	require.Equal(t, "{\"min_moniker_len\":\"2\",\"max_moniker_len\":\"30\"}", actual)
}

func TestBioLenParams_String(t *testing.T) {
	bioParams := models.NewBioLenParams(sdk.NewInt(1000))
	actual := bioParams.String()
	require.Equal(t, "{\"max_bio_len\":\"1000\"}", actual)
}

func TestValidateNameSurnameLenParams(t *testing.T) {
	tests := []struct {
		name   string
		params interface{}
		expErr error
	}{
		{
			name:   "Wrong params returns error",
			params: models.NewBioLenParams(sdk.NewInt(1000)),
			expErr: fmt.Errorf("invalid parameters type: {\"max_bio_len\":\"1000\"}"),
		},
		{
			name:   "Invalid min param returns error",
			params: models.NewNameSurnameLenParams(sdk.NewInt(1), sdk.NewInt(1000)),
			expErr: fmt.Errorf("invalid minimum name/surname length param: 1"),
		},
		{
			name:   "Invalid max param returns error",
			params: models.NewNameSurnameLenParams(sdk.NewInt(2), sdk.NewInt(-10)),
			expErr: fmt.Errorf("invalid max name/surname length param: -10"),
		},
		{
			name:   "Valid params returns no error",
			params: models.NewNameSurnameLenParams(sdk.NewInt(2), sdk.NewInt(1000)),
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
	tests := []struct {
		name   string
		params interface{}
		expErr error
	}{
		{
			name:   "Wrong params returns error",
			params: models.NewBioLenParams(sdk.NewInt(1000)),
			expErr: fmt.Errorf("invalid parameters type: {\"max_bio_len\":\"1000\"}"),
		},
		{
			name:   "Invalid min param return error",
			params: models.NewMonikerLenParams(sdk.NewInt(1), sdk.NewInt(30)),
			expErr: fmt.Errorf("invalid minimum moniker length param: 1"),
		},
		{
			name:   "Invalid max param return error",
			params: models.NewMonikerLenParams(sdk.NewInt(2), sdk.NewInt(-30)),
			expErr: fmt.Errorf("invalid max moniker length param: -30"),
		},
		{
			name:   "Valid params returns no error",
			params: models.NewMonikerLenParams(sdk.NewInt(3), sdk.NewInt(30)),
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
			name:   "Wrong params returns error",
			params: models.NewMonikerLenParams(sdk.NewInt(3), sdk.NewInt(30)),
			expErr: fmt.Errorf("invalid parameters type: {\"min_moniker_len\":\"3\",\"max_moniker_len\":\"30\"}"),
		},
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
