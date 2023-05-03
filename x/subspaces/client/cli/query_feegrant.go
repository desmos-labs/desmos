package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

// DONTCOVER

// GetAllowancesQueryCmd returns a new command to query subspace allowances
func GetAllowancesQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        "allowances",
		Short:                      "Query commands for subspace allowances",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		GetCmdQueryUserAllowances(),
		GetCmdQueryGroupAllowances(),
	)

	return queryCmd
}

// GetCmdQueryUserAllowances returns the command to query the fee allowances of a specific user
func GetCmdQueryUserAllowances() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "user-allowances [subspace-id] [[grantee]]",
		Short:   "Query user allowances within a subspace, for an optional given user",
		Example: fmt.Sprintf(`%s query subspaces user-allowances 1 desmos1evj20rymftvecmgn8t0xv700wkjlgucwfy4f0c`, version.AppName),
		Args:    cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var grantee string
			if len(args) > 1 {
				grantee = args[1]
			}

			res, err := queryClient.UserAllowances(
				context.Background(),
				types.NewQueryUserAllowancesRequest(subspaceID, grantee, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "user allowances")
	return cmd
}

// GetCmdQueryGroupAllowances returns the command to query the fee allowances of a specific group
func GetCmdQueryGroupAllowances() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "group-allowances [subspace-id] [[group-id]]",
		Short:   "Query the user group allowances within a subspace, for an optional given group",
		Example: fmt.Sprintf(`%s query subspaces group-allowances 1 1`, version.AppName),
		Args:    cobra.RangeArgs(1, 2),
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
			if len(args) > 1 {
				groupID, err = types.ParseGroupID(args[1])
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
				types.NewQueryGroupAllowancesRequest(subspaceID, groupID, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "group allowances")
	return cmd
}
