package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

var _ types.QueryServer = Keeper{}

// Session implements the Query/Session gRPC method
func (k Keeper) Session(ctx context.Context, request *types.QuerySessionRequest) (*types.QuerySessionResponse, error) {
	id, err := types.ParseSessionID(request.Id)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("invalid session id: %s", request.Id))
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	session, found := k.GetSession(sdkCtx, id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("session with id %s not found", request.Id))
	}

	return &types.QuerySessionResponse{Session: &session}, nil
}
