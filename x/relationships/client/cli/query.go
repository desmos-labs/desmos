package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd adds the query commands
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	profileQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the relationships module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	profileQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryUserRelationships(cdc),
		GetCmdQueryRelationships(cdc),
	)...)
	return profileQueryCmd
}

func GetCmdQueryRelationships(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "relationships",
		Short: "Retrieve all the relationships",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryRelationships)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("No relationships found")
				return nil
			}

			var out map[string][]sdk.AccAddress
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdQueryUserRelationships queries all the profiles' users' relationships
func GetCmdQueryUserRelationships(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "user_relationships [address]",
		Short: "Retrieve all the user's relationships",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryUserRelationships, args[0])
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find any relationship associated with the given address %s", args[0])
				return nil
			}

			var out types.RelationshipsResponse
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
