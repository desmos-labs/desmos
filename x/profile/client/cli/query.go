package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd adds the query commands
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	postQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the posts module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	postQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryAccount(cdc),
		GetCmdQueryAccounts(cdc),
	)...)
	return postQueryCmd
}

// GetCmdQueryAccount queries an account
func GetCmdQueryAccount(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "profile [moniker]",
		Short: "Retrieve the profile having the moniker, if any.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryProfile, args[0])
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find a profile with moniker %s \n", args[0])
				return nil
			}

			var out types.Profile
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdQueryAccounts queries all the accounts
func GetCmdQueryAccounts(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "profiles",
		Short: "Retrieve all the profiles.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryProfiles)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find any profile")
				return nil
			}

			var out types.Profiles
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
