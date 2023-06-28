package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate implements fmt.Validator
func (metadata DenomAuthorityMetadata) Validate() error {
	if metadata.Admin != "" {
		_, err := sdk.AccAddressFromBech32(metadata.Admin)
		if err != nil {
			return err
		}
	}
	return nil
}
