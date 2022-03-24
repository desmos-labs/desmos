package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/x/fees/types"
)

func TestMinFee_Validate(t *testing.T) {
	tests := []struct {
		name     string
		minFee   types.MinFee
		expError error
	}{
		{
			name:     "empty message type returns error",
			minFee:   types.NewMinFee("", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			expError: fmt.Errorf("invalid minimum fee message type"),
		},
		{
			name:     "invalid min fee amount returns error",
			minFee:   types.NewMinFee("create_post", sdk.Coins{sdk.Coin{Denom: "stakE", Amount: sdk.NewInt(0)}}),
			expError: fmt.Errorf("invalid minimum fee amount"),
		},
		{
			name:     "correct fee returns no errors",
			minFee:   types.NewMinFee("create_post", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			expError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.minFee.Validate()
			if err != nil {
				require.Equal(t, test.expError, err)
			} else {
				require.Equal(t, test.expError, err)
			}
		})
	}
}
