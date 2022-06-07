package keeper

import (
	"bytes"
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

var _ types.QueryServer = &Keeper{}

// getReportsStorePrefix returns the store prefix to be used to iterate reports based on the given request data
func getReportsStorePrefix(request *types.QueryReportsRequest) ([]byte, error) {
	if request.Target == nil {
		return types.SubspaceReportsPrefix(request.SubspaceId), nil
	}

	switch target := request.Target.GetCachedValue().(type) {
	case *types.UserTarget:
		if request.Reporter != "" {
			return types.UserReportStoreKey(request.SubspaceId, target.User, request.Reporter), nil
		}
		return types.UserReportsPrefix(request.SubspaceId, target.User), nil

	case *types.PostTarget:
		if request.Reporter != "" {
			return types.PostReportStoreKey(request.SubspaceId, target.PostID, request.Reporter), nil
		}
		return types.PostReportsPrefix(request.SubspaceId, target.PostID), nil

	default:
		return nil, fmt.Errorf("invalid target type")
	}
}

// Reports implements the QueryReports gRPC method
func (k Keeper) Reports(ctx context.Context, request *types.QueryReportsRequest) (*types.QueryReportsResponse, error) {
	if request.SubspaceId == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	// Get the proper store prefix
	storePrefix, err := getReportsStorePrefix(request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	reportsStore := prefix.NewStore(store, storePrefix)

	var reports []types.Report
	pageRes, err := query.Paginate(reportsStore, request.Pagination, func(key []byte, value []byte) error {
		var report types.Report

		switch {
		case bytes.HasPrefix(storePrefix, types.ReportPrefix):
			err = k.cdc.Unmarshal(value, &report)

		case bytes.HasPrefix(storePrefix, types.UsersReportsPrefix),
			bytes.HasPrefix(storePrefix, types.PostsReportsPrefix):
			reportID := types.GetReportIDFromBytes(value)
			storedReport, found := k.GetReport(sdkCtx, request.SubspaceId, reportID)
			if !found {
				err = fmt.Errorf("report reference not found")
			}
			report = storedReport

		}

		if err != nil {
			return err
		}

		reports = append(reports, report)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryReportsResponse{
		Reports:    reports,
		Pagination: pageRes,
	}, nil
}

// Reasons implements the QueryReasons gRPC method
func (k Keeper) Reasons(ctx context.Context, request *types.QueryReasonsRequest) (*types.QueryReasonsResponse, error) {
	if request.SubspaceId == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)
	reasonsStore := prefix.NewStore(store, types.SubspaceReasonsPrefix(request.SubspaceId))

	var reasons []types.Reason
	pageRes, err := query.Paginate(reasonsStore, request.Pagination, func(key []byte, value []byte) error {
		var reason types.Reason
		if err := k.cdc.Unmarshal(value, &reason); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		reasons = append(reasons, reason)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryReasonsResponse{
		Reasons:    reasons,
		Pagination: pageRes,
	}, nil
}

// Params implements the QueryParams gRPC method
func (k Keeper) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}
