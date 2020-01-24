package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd adds the query commands
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	magpieQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the magpie module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	magpieQueryCmd.AddCommand(
		GetCmdSession(storeKey, cdc),
	)
	return magpieQueryCmd
}

// GetCmdSession queries a session by PostID
func GetCmdSession(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "session [id]",
		Short: "Returns the session having the specified id, if any.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			sessionsID := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/session/%s", queryRoute, sessionsID), nil)

			if err != nil {
				fmt.Printf("Could not find session with id %s \n", sessionsID)
				return nil
			}

			var out types.Session
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
