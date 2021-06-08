package types_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/profiles/types"

	"github.com/stretchr/testify/require"
)

func TestPictures_Validate(t *testing.T) {
	tests := []struct {
		name     string
		pictures types.Pictures
		expErr   error
	}{
		{
			name:     "Valid Pictures",
			pictures: types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
			expErr:   nil,
		},
		{
			name:     "Invalid Pictures profile uri",
			pictures: types.NewPictures("invalid", "https://shorturl.at/cgpyF"),
			expErr:   fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name:     "Invalid Pictures cover uri",
			pictures: types.NewPictures("https://shorturl.at/adnX3", "invalid"),
			expErr:   fmt.Errorf("invalid profile cover uri provided"),
		},
	}

	for _, test := range tests {
		actErr := test.pictures.Validate()
		require.Equal(t, test.expErr, actErr)
	}
}
