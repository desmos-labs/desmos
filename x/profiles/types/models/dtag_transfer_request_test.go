package models_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/stretchr/testify/require"
)

func TestNewDTagTransferRequest(t *testing.T) {
	var user, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	otherUser, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	require.Equal(t, types.DTagTransferRequest{CurrentOwner: user, ReceivingUser: otherUser},
		types.NewDTagTransferRequest(user, otherUser))
}

func TestDTagTransferRequest_Equals(t *testing.T) {
	var user, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	otherUser, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name     string
		request  types.DTagTransferRequest
		otherReq types.DTagTransferRequest
		expBool  bool
	}{
		{
			name:     "Equals requests return true",
			request:  types.NewDTagTransferRequest(user, otherUser),
			otherReq: types.NewDTagTransferRequest(user, otherUser),
			expBool:  true,
		},
		{
			name:     "Non equals requests return false",
			request:  types.NewDTagTransferRequest(user, otherUser),
			otherReq: types.NewDTagTransferRequest(user, user),
			expBool:  false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.request.Equals(test.otherReq))
		})
	}
}

func TestDTagTransferRequest_String(t *testing.T) {
	var user, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	otherUser, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	require.Equal(t,
		"DTag transfer request:\n[Current CurrentOwner] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 [Receiving User] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		types.NewDTagTransferRequest(user, otherUser).String(),
	)
}

func TestDTagTransferRequest_Validate(t *testing.T) {
	var user, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	otherUser, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name    string
		request types.DTagTransferRequest
		expErr  error
	}{
		{
			name:    "Empty current owner returns error",
			request: types.NewDTagTransferRequest(nil, otherUser),
			expErr:  fmt.Errorf("current owner address cannot be empty"),
		},
		{
			name:    "Empty receiving user returns error",
			request: types.NewDTagTransferRequest(user, nil),
			expErr:  fmt.Errorf("receiving user address cannot be empty"),
		},
		{
			name:    "Equals current owner and receiving user addresses return error",
			request: types.NewDTagTransferRequest(user, user),
			expErr:  fmt.Errorf("the receiving user and current owner must be different"),
		},
		{
			name:    "Valid request returns no error",
			request: types.NewDTagTransferRequest(user, otherUser),
			expErr:  nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.request.Validate())
		})
	}
}
