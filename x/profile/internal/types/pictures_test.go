package types_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestPictures_Equals(t *testing.T) {
	tests := []struct {
		name     string
		pictures *types.Pictures
		otherPic *types.Pictures
		expBool  bool
	}{
		{
			name:     "Equals pictures returns true",
			pictures: types.NewPictures("profile", "cover"),
			otherPic: types.NewPictures("profile", "cover"),
			expBool:  true,
		},
		{
			name:     "Different pictures returns false",
			pictures: types.NewPictures("profile", "cover"),
			otherPic: types.NewPictures("prof", "cover"),
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
	tests := []struct {
		name     string
		pictures *types.Pictures
		expErr   error
	}{
		{
			name:     "Valid Pictures",
			pictures: types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
			expErr:   nil,
		},
		{
			name:     "Invalid Pictures profile uri",
			pictures: types.NewPictures("adnX3", "https://shorturl.at/cgpyF"),
			expErr:   fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name:     "Invalid Pictures cover uri",
			pictures: types.NewPictures("https://shorturl.at/adnX3", "cgpyF"),
			expErr:   fmt.Errorf("invalid profile cover uri provided"),
		},
	}

	for _, test := range tests {
		actErr := test.pictures.Validate()
		require.Equal(t, test.expErr, actErr)
	}
}
