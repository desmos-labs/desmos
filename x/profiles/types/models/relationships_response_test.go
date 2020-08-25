package models_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/stretchr/testify/require"
)

func TestRelationshipsResponse_String(t *testing.T) {
	address1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	address2, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	relationshipsResponse := types.NewRelationshipResponse([]sdk.AccAddress{address1, address2})

	require.Equal(t, "Relationships: [cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4]", relationshipsResponse.String())
}
