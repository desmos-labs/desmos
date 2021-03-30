package keeper

import (
	"context"

	"github.com/desmos-labs/desmos/x/links/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateLink(ctx context.Context, msg *types.MsgCreateLink) (*types.MsgCreateLinkResponse, error) {
	return &types.MsgCreateLinkResponse{}, nil
}
