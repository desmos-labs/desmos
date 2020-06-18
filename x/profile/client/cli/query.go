package cli

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

// GetQueryCmd adds the query commands
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	profileQueryCmd := &cobra.Command{
		Use:                        models.ModuleName,
		Short:                      "Querying commands for the profiles module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	profileQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryProfile(cdc),
		GetCmdQueryProfiles(cdc),
		GetCmdQueryProfileParams(cdc),
	)...)
	return profileQueryCmd
}

// GetCmdQueryProfile queries a profile from the given address or moniker
func GetCmdQueryProfile(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "profile [address_or_moniker]",
		Short: "Retrieve the profile having the specified user address or profile moniker, if any.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s/%s", models.QuerierRoute, models.QueryProfile, args[0])
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find a profile with moniker %s \n", args[0])
				return nil
			}

			var out models.Profile
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdQueryProfiles queries all the profiles
func GetCmdQueryProfiles(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "all",
		Short: "Retrieve all the registered profiles.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", models.QuerierRoute, models.QueryProfiles)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find any profile")
				return nil
			}

			var out models.Profiles
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdQueryProfileParams queries all the profiles' module params
func GetCmdQueryProfileParams(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Retrieve all the profile's params",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", models.QuerierRoute, models.QueryParams)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find profile params")
				return nil
			}

			var out models.Params
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
