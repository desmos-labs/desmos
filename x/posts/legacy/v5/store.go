package v5

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MigrateStore performs the migration from version 4 to version 5 of the store.
// To do this, it iterates over all the post attachments, and converts them to
// the new storing format (AttachmentContent instead of Attachment).
// It also removes all the Polls that have been saved as a Poll_ProvidedAnswer's attachment.
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	// TODO: Implement this
	panic("TODO: Implement me")
}
