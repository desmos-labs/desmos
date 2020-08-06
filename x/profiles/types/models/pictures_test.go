package models_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profiles/types/common"
	"github.com/desmos-labs/desmos/x/profiles/types/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPictures_Equals(t *testing.T) {
	tests := []struct {
		name      string
		pictures  *models.Pictures
		otherPics *models.Pictures
		expBool   bool
	}{
		{
			name:      "Different pictures returns false",
			pictures:  models.NewPictures(common.NewStrPtr("cover"), common.NewStrPtr("profile")),
			otherPics: models.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "First picture with nil value returns false (profile)",
			pictures:  models.NewPictures(nil, common.NewStrPtr("cover")),
			otherPics: models.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "First picture with nil value returns false (cover)",
			pictures:  models.NewPictures(common.NewStrPtr("profile"), nil),
			otherPics: models.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "Second picture with nil value returns false (profile)",
			pictures:  models.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			otherPics: models.NewPictures(nil, common.NewStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "Second picture with nil value returns false (cover)",
			pictures:  models.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			otherPics: models.NewPictures(common.NewStrPtr("profile"), nil),
			expBool:   false,
		},
		{
			name:      "Equals pictures returns true",
			pictures:  models.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			otherPics: models.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			expBool:   true,
		},
		{
			name:      "Same values but different pointers return true",
			pictures:  models.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			otherPics: models.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
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
		pictures *models.Pictures
		expErr   error
	}{
		{
			name:     "Valid Pictures",
			pictures: models.NewPictures(&profilePic, &profileCov),
			expErr:   nil,
		},
		{
			name:     "Invalid Pictures profile uri",
			pictures: models.NewPictures(&invalidURI, &profileCov),
			expErr:   fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name:     "Invalid Pictures cover uri",
			pictures: models.NewPictures(&profilePic, &invalidURI),
			expErr:   fmt.Errorf("invalid profile cover uri provided"),
		},
	}

	for _, test := range tests {
		actErr := test.pictures.Validate()
		require.Equal(t, test.expErr, actErr)
	}
}
