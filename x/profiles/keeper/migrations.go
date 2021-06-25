package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	v0163 "github.com/desmos-labs/desmos/x/profiles/legacy/v0163"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// DONTCOVER

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
	amino  *codec.LegacyAmino
}

// NewMigrator returns a new Migrators
func NewMigrator(amino *codec.LegacyAmino, keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
		amino:  amino,
	}
}

// Migrate0163to0170 migrates the current context from being compatible with Desmos v0.16.3
// to being compatible with Desmos v0.17.0
func (m Migrator) Migrate0163to0170(ctx sdk.Context) error {
	store := ctx.KVStore(m.keeper.storeKey)

	// Re-set all the DTag store entries to make them case-insensitive
	// Related fix: https://github.com/desmos-labs/desmos/issues/492
	m.keeper.IterateProfiles(ctx, func(index int64, profile *types.Profile) (stop bool) {
		v0163DTagKey := v0163.DTagStoreKey(profile.DTag)
		if store.Has(v0163DTagKey) {
			store.Delete(v0163DTagKey)
			store.Set(types.DTagStoreKey(profile.DTag), profile.GetAddress())
		}
		return false
	})

	m.keeper.ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {

		// Re-set all the relationships using the new way
		// Related fix: https://github.com/desmos-labs/desmos/issues/467
		v0163RelationshipsKey := v0163.RelationshipsStoreKey(account.GetAddress().String())
		if store.Has(v0163RelationshipsKey) {
			relationships := v0163.MustUnmarshalRelationships(m.keeper.cdc, store.Get(v0163RelationshipsKey))
			for _, relationship := range relationships {
				err := m.keeper.SaveRelationship(ctx, types.NewRelationship(
					relationship.Creator,
					relationship.Recipient,
					relationship.Subspace,
				))
				if err != nil {
					panic(err)
				}
			}
			store.Delete(v0163RelationshipsKey)
		}

		return false
	})

	return nil
}
