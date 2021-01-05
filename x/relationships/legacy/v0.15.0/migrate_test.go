package v0150_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v0130relationships "github.com/desmos-labs/desmos/x/relationships/legacy/v0.13.0"
	v0150relationships "github.com/desmos-labs/desmos/x/relationships/legacy/v0.15.0"
)

func TestMigrate(t *testing.T) {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	v013GenState := v0130relationships.GenesisState{
		UsersRelationships: map[string][]v0130relationships.Relationship{
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47": {
				{
					Recipient: receiver,
					Subspace:  "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				},
			},
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns": {
				{
					Recipient: sender,
					Subspace:  "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				},
			},
		},
		UsersBlocks: []v0130relationships.UserBlock{
			{
				Blocker:  sender,
				Blocked:  receiver,
				Reason:   "reason",
				Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			},
		},
	}

	expGenState := v0150relationships.GenesisState{
		Relationships: []v0150relationships.Relationship{
			{
				Creator:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				Recipient: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Subspace:  "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			},
			{
				Creator:   "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Recipient: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				Subspace:  "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			},
		},
		Blocks: []v0150relationships.UserBlock{
			{
				Blocker:  sender.String(),
				Blocked:  receiver.String(),
				Reason:   "reason",
				Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			},
		},
	}

	migrated := v0150relationships.Migrate(v013GenState)

	require.Len(t, expGenState.Relationships, len(migrated.Relationships))
	for index, relationship := range migrated.Relationships {
		require.Equal(t, expGenState.Relationships[index], relationship)
	}

	require.Len(t, expGenState.Blocks, len(migrated.Blocks))
	for index, block := range migrated.Blocks {
		require.Equal(t, expGenState.Blocks[index], block)
	}
}
