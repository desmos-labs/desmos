package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
	"github.com/spf13/cobra"
)

// GetCmdQueryUserAllowances returns the command to query the fee allowances of a specific user
func GetCmdQueryUserAllowances() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "user-allowances [[subspace-id]] [[grantee]] [[granter]]",
		Short:   "Query allowances for the given user",
		Example: fmt.Sprintf(`%s query subspaces user-allowances 1 desmos1evj20rymftvecmgn8t0xv700wkjlgucwfy4f0c desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud`, version.AppName),
		Args:    cobra.RangeArgs(0, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var subspaceID uint64
			if len(args) > 0 {
				subspaceID, err = types.ParseSubspaceID(args[0])
				if err != nil {
					return err
				}
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var grantee string
			if len(args) > 1 {
				grantee = args[1]
				_, err := sdk.AccAddressFromBech32(grantee)
				if err != nil {
					return err
				}
			}

			var granter string
			if len(args) > 2 {
				granter = args[2]
				_, err := sdk.AccAddressFromBech32(granter)
				if err != nil {
					return err
				}
			}

			res, err := queryClient.UserAllowances(
				context.Background(),
				types.NewQueryUserAllowancesRequest(subspaceID, granter, grantee, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "user fee allowances")
	return cmd
}

// GetCmdQueryGroupAllowances returns the command to query the fee allowances of a specific group
func GetCmdQueryGroupAllowances() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "group-allowances [[subspace-id]] [[group-id]] [[granter]]",
		Short:   "Query allowances for the given group",
		Example: fmt.Sprintf(`%s query subspaces group-allowances 1 1 desmos1evj20rymftvecmgn8t0xv700wkjlgucwfy4f0c`, version.AppName),
		Args:    cobra.RangeArgs(0, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var subspaceID uint64
			if len(args) > 0 {
				subspaceID, err = types.ParseSubspaceID(args[0])
				if err != nil {
					return err
				}
			}

			var groupID uint32
			if len(args) > 1 && args[1] != "" {
				groupID, err = types.ParseGroupID(args[1])
				if err != nil {
					return err
				}
			}

			var granter string
			if len(args) > 2 {
				granter = args[2]
				_, err := sdk.AccAddressFromBech32(granter)
				if err != nil {
					return err
				}
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.GroupAllowances(
				context.Background(),
				types.NewQueryGroupAllowancesRequest(subspaceID, granter, groupID, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "user fee allowances")
	return cmd
}
