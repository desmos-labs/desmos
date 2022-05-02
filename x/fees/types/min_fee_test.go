package types_test

import (
	"testing"

	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/fees/types"
)

func TestMinFee_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		fees      types.MinFee
		shouldErr bool
	}{
		{
			name:      "empty message type returns error",
			fees:      types.NewMinFee("", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			shouldErr: true,
		},
		{
			name: "invalid min fee amount returns error",
			fees: types.NewMinFee(
				sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
				sdk.Coins{sdk.Coin{Denom: "stakE", Amount: sdk.NewInt(0)}},
			),
			shouldErr: true,
		},
		{
			name: "correct fee returns no errors",
			fees: types.NewMinFee(
				sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
				sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000))),
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.fees.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
