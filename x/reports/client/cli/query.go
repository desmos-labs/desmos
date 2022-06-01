package cli

// DONTCOVER

import (
	"context"
	"fmt"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// GetQueryCmd returns the command allowing to perform queries
func GetQueryCmd() *cobra.Command {
	subspaceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the reports module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	subspaceQueryCmd.AddCommand(
		GetCmdQueryUserReports(),
		GetCmdQueryPostReports(),
		GetCmdQueryReports(),
		GetCmdQueryReasons(),
		GetCmdQueryParams(),
	)
	return subspaceQueryCmd
}

// GetCmdQueryUserReports returns the command to query the reports associated to a user
func GetCmdQueryUserReports() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "user-reports [subspace-id] [user-address]",
		Short:   "Query the reports made towards the specified user",
		Example: fmt.Sprintf(`%s query reports user-reports 1 desmos1cs0gu6006rz9wnmltjuhnuz8k3a2wg6jzmmgyu`, version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			userAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Reports(context.Background(), types.NewQueryReportsRequest(
				subspaceID,
				types.NewUserTarget(userAddr.String()),
				pageReq,
			))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "user reports")

	return cmd
}

// GetCmdQueryPostReports returns the command to query the reports associated to a post
func GetCmdQueryPostReports() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "post-reports [subspace-id] [post-id]",
		Short:   "Query the reports made towards the specified post",
		Example: fmt.Sprintf(`%s query reports post-reports 1 1`, version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := poststypes.ParsePostID(args[1])
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Reports(context.Background(), types.NewQueryReportsRequest(
				subspaceID,
				types.NewPostTarget(postID),
				pageReq,
			))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "post reports")

	return cmd
}

// GetCmdQueryReports returns the command to query the reports of a subspace
func GetCmdQueryReports() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reports [subspace-id]",
		Short:   "Query the reports from within the specified subspace",
		Example: fmt.Sprintf(`%s query reports reports 1`, version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Reports(context.Background(), types.NewQueryReportsRequest(subspaceID, nil, pageReq))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "reports")

	return cmd
}

// GetCmdQueryReasons returns the command to query the reasons of a subspace
func GetCmdQueryReasons() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reasons [subspace-id]",
		Short:   "Query the reasons from within the specified subspace",
		Example: fmt.Sprintf(`%s query reports reasons 1`, version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Reasons(context.Background(), types.NewQueryReasonsRequest(subspaceID, pageReq))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "reasons")

	return cmd
}

// GetCmdQueryParams returns the command to query the params of the module
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "params",
		Short:   "Query the module parameters",
		Example: fmt.Sprintf(`%s query reports params 1`, version.AppName),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), types.NewQueryParamsRequest())
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
