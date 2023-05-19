package authz_test

import (
	"testing"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmosauthz "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"

	relationshipstypes "github.com/desmos-labs/desmos/v5/x/relationships/types"

	"github.com/desmos-labs/desmos/v5/testutil/storetesting"
	"github.com/desmos-labs/desmos/v5/x/subspaces/authz"
)

func TestGenericSubspaceAuthorization_Accept(t *testing.T) {
	testCases := []struct {
		name          string
		authorization *authz.GenericSubspaceAuthorization
		msg           sdk.Msg
		shouldErr     bool
		expResponse   cosmosauthz.AcceptResponse
	}{
		{
			name:          "invalid message type returns error",
			authorization: authz.NewGenericSubspaceAuthorization([]uint64{1}, sdk.MsgTypeURL(&relationshipstypes.MsgCreateRelationship{})),
			msg:           &banktypes.MsgSend{},
			shouldErr:     true,
		},
		{
			name:          "wrong subspace id is rejected",
			authorization: authz.NewGenericSubspaceAuthorization([]uint64{1}, sdk.MsgTypeURL(&relationshipstypes.MsgCreateRelationship{})),
			msg: relationshipstypes.NewMsgCreateRelationship(
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				2,
			),
			shouldErr:   false,
			expResponse: cosmosauthz.AcceptResponse{Accept: false},
		},
		{
			name:          "valid subspace id is accepted",
			authorization: authz.NewGenericSubspaceAuthorization([]uint64{1}, sdk.MsgTypeURL(&relationshipstypes.MsgCreateRelationship{})),
			msg: relationshipstypes.NewMsgCreateRelationship(
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				1,
			),
			shouldErr:   false,
			expResponse: cosmosauthz.AcceptResponse{Accept: true},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := storetesting.BuildContext(nil, nil, nil)
			ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())

			response, err := tc.authorization.Accept(ctx, tc.msg)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResponse, response)
			}
		})
	}
}

func TestGenericSubspaceAuthorization_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name          string
		authorization *authz.GenericSubspaceAuthorization
		shouldErr     bool
	}{
		{
			name:          "empty subspaces ids return error",
			authorization: authz.NewGenericSubspaceAuthorization(nil, sdk.MsgTypeURL(&relationshipstypes.MsgCreateRelationship{})),
			shouldErr:     true,
		},
		{
			name:          "invalid subspace id returns error",
			authorization: authz.NewGenericSubspaceAuthorization([]uint64{0}, sdk.MsgTypeURL(&relationshipstypes.MsgCreateRelationship{})),
			shouldErr:     true,
		},
		{
			name:          "valid data returns no error",
			authorization: authz.NewGenericSubspaceAuthorization([]uint64{1, 2}, sdk.MsgTypeURL(&relationshipstypes.MsgCreateRelationship{})),
			shouldErr:     false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.authorization.ValidateBasic()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
