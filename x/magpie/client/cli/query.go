package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

// GetQueryCmd adds the query commands
func GetQueryCmd() *cobra.Command {
	magpieQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the magpie module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	magpieQueryCmd.AddCommand(GetCmdSession())
	return magpieQueryCmd
}

// GetCmdSession queries a session by PostID
func GetCmdSession() *cobra.Command {
	return &cobra.Command{
		Use:   "session [id]",
		Short: "Returns the session having the specified id, if any.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Session(
				context.Background(),
				&types.QuerySessionRequest{Id: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Session)
		},
	}
}
