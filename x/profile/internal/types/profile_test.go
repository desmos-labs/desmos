package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestProfile_String(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	var testAccount = types.Profile{
		Name:     "name",
		Surname:  "surname",
		Moniker:  "moniker",
		Bio:      "biography",
		Pictures: &testPictures,
		Creator:  owner,
	}

	require.Equal(t,
		`{"name":"name","surname":"surname","moniker":"moniker","bio":"biography","pictures":{"profile":"profile","cover":"cover"},"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}`,
		testAccount.String(),
	)
}

func TestProfile_Equals(t *testing.T) {
	var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	var testPictures = types.NewPictures("profile", "cover")

	var testAccount = types.Profile{
		Name:     "name",
		Surname:  "surname",
		Moniker:  "moniker",
		Bio:      "biography",
		Pictures: &testPictures,
		Creator:  testPostOwner,
	}

	var testAccount2 = types.Profile{
		Name:     "name",
		Surname:  "surname",
		Moniker:  "oniker",
		Bio:      "biography",
		Pictures: &testPictures,
		Creator:  testPostOwner,
	}

	tests := []struct {
		name     string
		account  types.Profile
		otherAcc types.Profile
		expBool  bool
	}{
		{
			name:     "Equals accounts returns true",
			account:  testAccount,
			otherAcc: testAccount,
			expBool:  true,
		},
		{
			name:     "Non equals account returns false",
			account:  testAccount,
			otherAcc: testAccount2,
			expBool:  false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual := test.account.Equals(test.otherAcc)
			require.Equal(t, test.expBool, actual)
		})
	}

}

//TODO add tests for chainLink and verifiedServices when implemented
func TestProfile_Validate(t *testing.T) {
	var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	tests := []struct {
		name    string
		account types.Profile
		expErr  error
	}{
		{
			name: "Empty profile creator returns error",
			account: types.Profile{
				Name:     "name",
				Surname:  "surname",
				Moniker:  "moniker",
				Bio:      "biography",
				Pictures: &testPictures,
				Creator:  nil,
			},
			expErr: fmt.Errorf("profile creator cannot be empty or blank"),
		},
		{
			name: "Empty profile moniker returns error",
			account: types.Profile{
				Name:     "name",
				Surname:  "surname",
				Moniker:  "",
				Bio:      "biography",
				Pictures: &testPictures,
				Creator:  testPostOwner,
			},
			expErr: fmt.Errorf("profile moniker cannot be empty or blank"),
		},
		{
			name: "Valid account returns no error",
			account: types.Profile{
				Name:     "name",
				Surname:  "surname",
				Moniker:  "moniker",
				Bio:      "biography",
				Pictures: &testPictures,
				Creator:  testPostOwner,
			},
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual := test.account.Validate()
			require.Equal(t, test.expErr, actual)
		})
	}
}
