package types_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestPictures_Equals(t *testing.T) {
	profilePic := "profile"
	profileCov := "cover"
	tests := []struct {
		name     string
		pictures *types.Pictures
		otherPic *types.Pictures
		expBool  bool
	}{
		{
			name:     "Equals pictures returns true",
			pictures: types.NewPictures(&profilePic, &profileCov),
			otherPic: types.NewPictures(&profilePic, &profileCov),
			expBool:  true,
		},
		{
			name:     "Different pictures returns false",
			pictures: types.NewPictures(&profileCov, &profilePic),
			otherPic: types.NewPictures(&profilePic, &profileCov),
			expBool:  false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual := test.pictures.Equals(test.otherPic)
			require.Equal(t, test.expBool, actual)
		})
	}
}

func TestPictures_Validate(t *testing.T) {
	profilePic := "https://shorturl.at/adnX3"
	profileCov := "https://shorturl.at/cgpyF"
	invalidURI := "invalid"
	tests := []struct {
		name     string
		pictures *types.Pictures
		expErr   error
	}{
		{
			name:     "Valid Pictures",
			pictures: types.NewPictures(&profilePic, &profileCov),
			expErr:   nil,
		},
		{
			name:     "Invalid Pictures profile uri",
			pictures: types.NewPictures(&invalidURI, &profileCov),
			expErr:   fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name:     "Invalid Pictures cover uri",
			pictures: types.NewPictures(&profilePic, &invalidURI),
			expErr:   fmt.Errorf("invalid profile cover uri provided"),
		},
	}

	for _, test := range tests {
		actErr := test.pictures.Validate()
		require.Equal(t, test.expErr, actErr)
	}
}
