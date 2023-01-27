package v4

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
)

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"desmos.posts.v2.AttachmentContent",
		(*AttachmentContent)(nil),
		&Poll{},
		&Media{},
	)
}
