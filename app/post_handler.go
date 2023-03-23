package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/posthandler"
)

func NewPostHandler() (sdk.PostHandler, error) {
	return posthandler.NewPostHandler(
		posthandler.HandlerOptions{},
	)
}
