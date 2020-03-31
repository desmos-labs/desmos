package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestPictures_Equals(t *testing.T) {
	tests := []struct {
		name     string
		pictures types.Pictures
		otherPic types.Pictures
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
			actual := test.pictures.Equals(&test.otherPic)
			require.Equal(t, test.expBool, actual)
		})
	}
}
