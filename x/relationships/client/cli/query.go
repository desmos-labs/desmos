package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

// GetQueryCmd adds the query commands
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the relationships module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(
		GetCmdQueryUserRelationships(),
		GetCmdQueryRelationships(),
		GetCmdQueryUserBlocks(),
	)
	return cmd
}

func GetCmdQueryRelationships() *cobra.Command {
	return &cobra.Command{
		Use:   "all",
		Short: "Retrieve all the relationships",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Relationships(context.Background(), &types.QueryRelationshipsRequest{})
			if err != nil {
				return fmt.Errorf("no relationships found")
			}

			return clientCtx.PrintOutput(res.Relationships)
		},
	}
}

// GetCmdQueryUserRelationships queries all the profiles' users' relationships
func GetCmdQueryUserRelationships(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "user [address]",
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

			var out types.Relationships
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdQueryUserBlocks(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "blacklist [address]",
		Short: "Retrieve the list of all the blocked users of the given address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryUserBlocks, args[0])
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find any user block associated with the given address %s", args[0])
				return nil
			}

			var out []types.UserBlock
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
