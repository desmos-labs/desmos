package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	subspacestypes "github.com/desmos-labs/desmos/v2/x/subspaces/types"

	"github.com/desmos-labs/desmos/v2/x/relationships/types"
)

// GetQueryCmd returns the command allowing to perform queries
func GetQueryCmd() *cobra.Command {
	profileQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the relationships module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	profileQueryCmd.AddCommand(
		GetCmdQueryRelationships(),
		GetCmdQueryBlocks(),
	)
	return profileQueryCmd
}

// GetCmdQueryRelationships returns the command allowing to query the relationships with optional user and subspace
func GetCmdQueryRelationships() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relationships [subspace-id] [[creator]] [[counterparty]]",
		Short: "Retrieve all the relationships inside a given subspace, with optional creator and counterparty",
		Args:  cobra.RangeArgs(1, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			var creator string
			if len(args) > 1 {
				creator = args[1]
			}

			var counterparty string
			if len(args) > 2 {
				counterparty = args[2]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Relationships(
				context.Background(),
				types.NewQueryRelationshipsRequest(subspaceID, creator, counterparty, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "user relationships")

	return cmd
}

// GetCmdQueryBlocks returns the command allowing to query all the blocks with optional user and subspace
func GetCmdQueryBlocks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blocks [subspace-id] [[blocker]] [[blocked]]",
		Short: "Retrieve the list of all the user blocks present inside the given subspace with optional blocker and blocked addresses",
		Args:  cobra.RangeArgs(1, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			var blocker string
			if len(args) > 1 {
				blocker = args[1]
			}

			var blocked string
			if len(args) > 2 {
				blocked = args[2]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Blocks(
				context.Background(),
				types.NewQueryBlocksRequest(subspaceID, blocker, blocked, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "user blocks")

	return cmd
}
