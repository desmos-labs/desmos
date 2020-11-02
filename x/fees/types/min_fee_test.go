package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/fees/types"
	"github.com/stretchr/testify/require"
)

func TestMinFee_Validate(t *testing.T) {
	tests := []struct {
		name     string
		minFee   types.MinFee
		expError error
	}{
		{
			name:     "empty message type returns error",
			minFee:   types.NewMinFee("", sdk.NewDecWithPrec(1, 2)),
			expError: fmt.Errorf("invalid minimum fee message type"),
		},
		{
			name:     "negative minimum fee amount returns error",
			minFee:   types.NewMinFee("desmos/createPost", sdk.NewDec(-1)),
			expError: fmt.Errorf("minimum fee amout cannot be negative"),
		},
		{
			name:     "no errors",
			minFee:   types.NewMinFee("desmos/createPost", sdk.NewDec(1)),
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
