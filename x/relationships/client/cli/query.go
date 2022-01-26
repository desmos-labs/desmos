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
		Use:   "relationships [[address]] [[subspace-id]]",
		Short: "Retrieve all the relationships with optional address and subspace",
		Args:  cobra.RangeArgs(0, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var user string
			if len(args) >= 1 {
				user = args[0]
			}

			var subspace uint64
			if len(args) == 2 {
				subspace, err = subspacestypes.ParseSubspaceID(args[1])
				if err != nil {
					return err
				}
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Relationships(
				context.Background(),
				&types.QueryRelationshipsRequest{User: user, SubspaceId: subspace, Pagination: pageReq},
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
		Use:   "blocks [[address]] [[subspace-id]] ",
		Short: "Retrieve the list of all the blocked users with optional address and subspace",
		Args:  cobra.RangeArgs(0, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var user string
			if len(args) >= 1 {
				user = args[0]
			}

			var subspace uint64
			if len(args) == 2 {
				subspace, err = subspacestypes.ParseSubspaceID(args[1])
				if err != nil {
					return err
				}
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Blocks(
				context.Background(),
				&types.QueryBlocksRequest{User: user, SubspaceId: subspace, Pagination: pageReq})
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
