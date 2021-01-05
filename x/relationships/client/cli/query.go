package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client/flags"

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
		GetCmdQueryUserBlocks(),
	)
	return cmd
}

// GetCmdQueryUserRelationships returns the command allowing to query all the relationships of a specific user
func GetCmdQueryUserRelationships() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user [address]",
		Short: "Retrieve all the user's relationships",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
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

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryUserBlocks returns the command allowing to query all the blocks of a single user
func GetCmdQueryUserBlocks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blocklist [address]",
		Short: "Retrieve the list of all the blocked users of the given address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
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

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
