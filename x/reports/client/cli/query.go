package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd adds the query commands
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	postQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the reports module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	postQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryPostReports(cdc),
	)...)
	return postQueryCmd
}

// GetCmdQueryPostReports queries a post's reports
func GetCmdQueryPostReports(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "all [post-id]",
		Short: "Returns all the reports of the posts with the given ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			postID := args[0]

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryReports, postID)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find post with id %s \n", postID)
				return nil
			}

			var out types.ReportsQueryResponse
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
