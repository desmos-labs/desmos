package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/desmos-labs/desmos/x/reports/types"
	"github.com/spf13/cobra"
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

// GetCmdQueryPostReports returns the command that allows to query a post's reports
func GetCmdQueryPostReports() *cobra.Command {
	return &cobra.Command{
		Use:   "post [id]",
		Short: "Returns all the reports of the posts with the given ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
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

			return clientCtx.PrintOutput(res)
		},
	}
}
