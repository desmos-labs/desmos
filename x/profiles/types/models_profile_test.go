package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"

	"github.com/stretchr/testify/require"
)

func TestPictures_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		pictures  types.Pictures
		shouldErr bool
	}{
		{
			name:      "invalid profile uri returns error",
			pictures:  types.NewPictures("invalid", "https://shorturl.at/cgpyF"),
			shouldErr: true,
		},
		{
			name:      "invalid cover uri returns error",
			pictures:  types.NewPictures("https://shorturl.at/adnX3", "invalid"),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			pictures:  types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.pictures.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
