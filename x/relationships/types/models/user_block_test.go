package models_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	address1, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	address2, _ = sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")

	userBlock = types.UserBlock{
		Blocker: address1,
		Blocked: address2,
		Reason:  "idk",
	}
)

func TestNewUserBlock(t *testing.T) {
	actual := types.NewUserBlock(address1, address2, "idk")

	require.Equal(t, userBlock, actual)
}

func TestUserBlock_String(t *testing.T) {
	require.Equal(t, "User Block: [Blocker] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 [Blocked] cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 [Reason] idk", userBlock.String())
}

func TestUserBlock_Validate(t *testing.T) {
	tests := []struct {
		name      string
		userBlock types.UserBlock
		expError  error
	}{
		{
			name:      "empty blocker address returns error",
			userBlock: types.NewUserBlock(nil, address2, "reason"),
			expError:  fmt.Errorf("blocker address cannot be empty"),
		},
		{
			name:      "empty blocked address returns error",
			userBlock: types.NewUserBlock(address1, nil, "reason"),
			expError:  fmt.Errorf("the address of the blocked user cannot be empty"),
		},
		{
			name:      "equals blocker and blocked addresses returns error",
			userBlock: types.NewUserBlock(address1, address1, "reason"),
			expError:  fmt.Errorf("blocker and blocked addresses cannot be equals"),
		},
		{
			name:      "correct user block returns no error",
			userBlock: types.NewUserBlock(address1, address2, "reason"),
			expError:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expError, test.userBlock.Validate())
		})
	}
}

func TestUserBlock_Equals(t *testing.T) {
	tests := []struct {
		name           string
		userBlock      types.UserBlock
		otherUserBlock types.UserBlock
		expBool        bool
	}{
		{
			name:           "Equals user block returns true",
			userBlock:      types.NewUserBlock(address1, address2, "reason"),
			otherUserBlock: types.NewUserBlock(address1, address2, "reason"),
			expBool:        true,
		},
		{
			name:           "Non equals user block returns false",
			userBlock:      types.NewUserBlock(address2, address1, "reason"),
			otherUserBlock: types.NewUserBlock(address1, address2, "reason"),
			expBool:        false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.userBlock.Equals(test.otherUserBlock))
		})
	}
}
