package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/reports/types"
)

// GetQueryCmd adds the query commands
func GetQueryCmd() *cobra.Command {
	postQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the reports module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	postQueryCmd.AddCommand(
		GetCmdQueryPostReports(),
	)
	return postQueryCmd
}

// GetCmdQueryPostReports returns the command that allows to query the reports of a post
func GetCmdQueryPostReports() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post [id]",
		Short: "Returns all the reports of the posts with the given id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.PostReports(
				context.Background(),
				&types.QueryPostReportsRequest{PostId: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
