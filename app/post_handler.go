package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/posthandler"
)

// NewPostHandler returns an empty PostHandler chain.
// PostHandler is like AnteHandler but it executes after RunMsgs.
func NewPostHandler() (sdk.PostHandler, error) {
	return posthandler.NewPostHandler(
		posthandler.HandlerOptions{},
	)
}
