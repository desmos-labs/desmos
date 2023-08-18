package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"

	"github.com/stretchr/testify/require"
)

func TestDTagTransferRequest_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		request   types.DTagTransferRequest
		shouldErr bool
	}{
		{
			name: "empty DTag to trade returns error",
			request: types.NewDTagTransferRequest(
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "empty request sender returns error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "empty request receiver returns error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
			),
			shouldErr: true,
		},
		{
			name: "equals request receiver and request sender addresses return error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.request.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
