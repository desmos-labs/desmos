package keeper_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

func Test_queryProfile(t *testing.T) {

	tests := []struct {
		name          string
		path          []string
		storedAccount types.Profile
		expErr        error
	}{
		{
			name:          "Profile doesnt exist (address given)",
			path:          []string{types.QueryProfile, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			storedAccount: testProfile,
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("Profile with address %s doesn't exists", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			),
		},
		{
			name:          "Profile doesnt exist (blank path given)",
			path:          []string{types.QueryProfile, ""},
			storedAccount: testProfile,
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"DTag or address cannot be empty or blank",
			),
		},
		{
			name:          "Profile doesnt exist (dtag given)",
			path:          []string{types.QueryProfile, "monk"},
			storedAccount: testProfile,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No address related to this dtag: monk"),
		},
		{
			name:          "Profile returned correctly (address given)",
			path:          []string{types.QueryProfile, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			storedAccount: testProfile,
			expErr:        nil,
		},
		{
			name:          "Profile returned correctly (dtag given)",
			path:          []string{types.QueryProfile, "dtag"},
			storedAccount: testProfile,
			expErr:        nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			err := k.SaveProfile(ctx, test.storedAccount)
			require.Nil(t, err)

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if result != nil {
				require.Nil(t, err)
				expectedIndented, err := codec.MarshalJSONIndent(k.Cdc, &test.storedAccount)
				require.NoError(t, err)
				require.Equal(t, string(expectedIndented), string(result))
			}

			if result == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
				require.Nil(t, result)
			}

		})
	}

}

func Test_queryProfiles(t *testing.T) {

	tests := []struct {
		name          string
		path          []string
		storedAccount *types.Profile
		expResult     types.Profiles
	}{
		{
			name:          "Empty Profiles",
			path:          []string{types.QueryProfiles},
			storedAccount: nil,
			expResult:     types.Profiles{},
		},
		{
			name:          "Profile returned correctly",
			path:          []string{types.QueryProfiles},
			storedAccount: &testProfile,
			expResult:     types.Profiles{testProfile},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.storedAccount != nil {
				err := k.SaveProfile(ctx, *test.storedAccount)
				require.Nil(t, err)
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if result != nil {
				require.Nil(t, err)
				expectedIndented, err := codec.MarshalJSONIndent(k.Cdc, &test.expResult)
				require.NoError(t, err)
				require.Equal(t, string(expectedIndented), string(result))
			}

		})
	}

}
