package types_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPictures_Equals(t *testing.T) {
	tests := []struct {
		name      string
		pictures  *types.Pictures
		otherPics *types.Pictures
		expBool   bool
	}{
		{
			name:      "Different pictures returns false",
			pictures:  types.NewPictures(newStrPtr("cover"), newStrPtr("profile")),
			otherPics: types.NewPictures(newStrPtr("profile"), newStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "First picture with nil value returns false (profile)",
			pictures:  types.NewPictures(nil, newStrPtr("cover")),
			otherPics: types.NewPictures(newStrPtr("profile"), newStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "First picture with nil value returns false (cover)",
			pictures:  types.NewPictures(newStrPtr("profile"), nil),
			otherPics: types.NewPictures(newStrPtr("profile"), newStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "Second picture with nil value returns false (profile)",
			pictures:  types.NewPictures(newStrPtr("profile"), newStrPtr("cover")),
			otherPics: types.NewPictures(nil, newStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "Second picture with nil value returns false (cover)",
			pictures:  types.NewPictures(newStrPtr("profile"), newStrPtr("cover")),
			otherPics: types.NewPictures(newStrPtr("profile"), nil),
			expBool:   false,
		},
		{
			name:      "Equals pictures returns true",
			pictures:  types.NewPictures(newStrPtr("profile"), newStrPtr("cover")),
			otherPics: types.NewPictures(newStrPtr("profile"), newStrPtr("cover")),
			expBool:   true,
		},
		{
			name:      "Same values but different pointers return true",
			pictures:  types.NewPictures(newStrPtr("profile"), newStrPtr("cover")),
			otherPics: types.NewPictures(newStrPtr("profile"), newStrPtr("cover")),
			expBool:   true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual := test.pictures.Equals(test.otherPics)
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
