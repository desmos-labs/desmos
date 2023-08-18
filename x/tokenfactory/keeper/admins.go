package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/tokenfactory/types"
)

// GetAuthorityMetadata returns the authority metadata for a specific denom
func (k Keeper) GetAuthorityMetadata(ctx sdk.Context, denom string) types.DenomAuthorityMetadata {
	bz := k.GetDenomPrefixStore(ctx, denom).Get([]byte(types.DenomAuthorityMetadataKey))

	var metadata types.DenomAuthorityMetadata
	k.cdc.MustUnmarshal(bz, &metadata)
	return metadata
}

// SetAuthorityMetadata stores authority metadata for a specific denom
func (k Keeper) SetAuthorityMetadata(ctx sdk.Context, denom string, metadata types.DenomAuthorityMetadata) error {
	err := metadata.Validate()
	if err != nil {
		return err
	}

	store := k.GetDenomPrefixStore(ctx, denom)

	store.Set([]byte(types.DenomAuthorityMetadataKey), k.cdc.MustMarshal(&metadata))
	return nil
}
