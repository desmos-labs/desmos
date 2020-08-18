package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// GetQueryCmd adds the query commands
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	profileQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the profiles module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	profileQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryProfile(cdc),
		GetCmdQueryProfiles(cdc),
		GetCmdQueryProfileParams(cdc),
		GetCmdQueryUserRelationships(cdc),
	)...)
	return profileQueryCmd
}

// GetCmdQueryProfile queries a profile from the given address or dtag
func GetCmdQueryProfile(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "profile [address_or_dtag]",
		Short: "Retrieve the profile having the specified user address or profile dtag, if any.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryProfile, args[0])
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find a profile with dtag %s \n", args[0])
				return nil
			}

			var out types.Profile
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
		Args:  cobra.NoArgs,
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

// GetCmdQueryProfileParams queries all the profiles' module params
func GetCmdQueryProfileParams(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "parameters",
		Short: "Retrieve all the profile module parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParams)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find profile parameters")
				return nil
			}

			var out types.Params
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdQueryUserRelationships queries all the profiles' users' relationships
func GetCmdQueryUserRelationships(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "relationships [address]",
		Short: "Retrieve all the user's relationships",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryRelationships, args[0])
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find profile parameters")
				return nil
			}

			var out types.Relationships
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
