package models_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/relationships/types/models"
)

var (
	address, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	subspace   = "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	relationship = models.Relationship{
		Recipient: address,
		Subspace:  subspace,
	}
)

func TestNewRelationship(t *testing.T) {
	actual := models.NewRelationship(address, subspace)
	require.Equal(t, relationship, actual)
}

func TestRelationship_String(t *testing.T) {
	require.Equal(t, "Relationship:[Recipient] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", relationship.String())
}

func TestRelationship_Validate(t *testing.T) {
	tests := []struct {
		name         string
		relationship models.Relationship
		expErr       error
	}{
		{
			name:         "Empty recipient returns error",
			relationship: models.NewRelationship(nil, ""),
			expErr:       fmt.Errorf("recipient can't be empty"),
		},
		{
			name:         "Invalid subspace returns error",
			relationship: models.NewRelationship(address, ""),
			expErr:       fmt.Errorf("subspace must be a valid sha-256"),
		},
		{
			name:         "Valid relationship returns no error",
			relationship: relationship,
			expErr:       nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.relationship.Validate())
		})
	}
}

func TestRelationship_Equals(t *testing.T) {
	tests := []struct {
		name         string
		relationship models.Relationship
		otherRel     models.Relationship
		expBool      bool
	}{
		{
			name:         "Equals relationships returns true",
			relationship: relationship,
			otherRel:     relationship,
			expBool:      true,
		},
		{
			name:         "Non equals relationships returns false",
			relationship: relationship,
			otherRel:     models.NewRelationship(address, "1234"),
			expBool:      false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.relationship.Equals(test.otherRel))
		})
	}
}
