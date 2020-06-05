package keeper_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/stretchr/testify/require"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

func Test_queryProfile(t *testing.T) {

	tests := []struct {
		name          string
		path          []string
		storedAccount models.Profile
		expErr        error
	}{
		{
			name:          "Profile doesnt exist (address given)",
			path:          []string{models.QueryProfile, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			storedAccount: testProfile,
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("Profile with address %s doesn't exists", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			),
		},
		{
			name:          "Profile doesnt exist (blank path given)",
			path:          []string{models.QueryProfile, ""},
			storedAccount: testProfile,
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"Moniker or address cannot be empty or blank",
			),
		},
		{
			name:          "Profile doesnt exist (moniker given)",
			path:          []string{models.QueryProfile, "monk"},
			storedAccount: testProfile,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No address related to this moniker: monk"),
		},
		{
			name:          "Profile returned correctly (address given)",
			path:          []string{models.QueryProfile, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			storedAccount: testProfile,
			expErr:        nil,
		},
		{
			name:          "Profile returned correctly (moniker given)",
			path:          []string{models.QueryProfile, "moniker"},
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
		storedAccount *models.Profile
		expResult     models.Profiles
	}{
		{
			name:          "Empty Profiles",
			path:          []string{models.QueryProfiles},
			storedAccount: nil,
			expResult:     models.Profiles{},
		},
		{
			name:          "Profile returned correctly",
			path:          []string{models.QueryProfiles},
			storedAccount: &testProfile,
			expResult:     models.Profiles{testProfile},
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
