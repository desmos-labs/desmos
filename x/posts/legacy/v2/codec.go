package v2

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
)

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"desmos.posts.v1.AttachmentContent",
		(*AttachmentContent)(nil),
		&Poll{},
		&Media{},
	)
}
