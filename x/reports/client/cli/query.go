package cli

// DONTCOVER

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
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
		GetCmdQueryReports(),
		GetCmdQueryReasons(),
		GetCmdQueryParams(),
	)
	return subspaceQueryCmd
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

			target, err := ReadReportTarget(cmd.Flags())
			if err != nil {
				return err
			}

			reporter, err := cmd.Flags().GetString(FlagReporter)
			if err != nil {
				return err
			}

			res, err := queryClient.Reports(
				context.Background(),
				types.NewQueryReportsRequest(subspaceID, target, reporter, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(FlagUser, "", "Optional address of the reported user to query the reports for")
	cmd.Flags().Uint64(FlagPostID, 0, "Optional id of the post to query the reports for")
	cmd.Flags().String(FlagReporter, "", fmt.Sprintf("Optional address of the reporter, used only if either --%s or --%s is specified", FlagUser, FlagPostID))

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
