package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
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

// GetCmdQueryRelationships returns the command that allows to query for all the stored relationships
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

			return clientCtx.PrintOutput(res)
		},
	}
}

// GetCmdQueryUserRelationships returns the command allowing to query all the relationships of a specific user
func GetCmdQueryUserRelationships() *cobra.Command {
	return &cobra.Command{
		Use:   "user [address]",
		Short: "Retrieve all the user's relationships",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.UserRelationships(
				context.Background(),
				&types.QueryUserRelationshipsRequest{User: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
}

// GetCmdQueryUserBlocks returns the command allowing to query all the blocks of a single user
func GetCmdQueryUserBlocks() *cobra.Command {
	return &cobra.Command{
		Use:   "blacklist [address]",
		Short: "Retrieve the list of all the blocked users of the given address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.UserBlocks(
				context.Background(),
				&types.QueryUserBlocksRequest{User: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
}
