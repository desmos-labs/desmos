package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the command allowing to perform queries
func GetQueryCmd() *cobra.Command {
	subspaceQuerycmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the subspaces module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	subspaceQuerycmd.AddCommand(
		GetCmdQuerySubspace(),
		GetCmdQuerySubspaces(),
	)
	return subspaceQuerycmd
}

// GetCmdQuerySubspace returns the command to query the subspace with the given id
func GetCmdQuerySubspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subspace [id]",
		Short: "Get the subspace with the given [id]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Subspace(
				context.Background(),
				&types.QuerySubspaceRequest{SubspaceID: args[0]},
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

func GetCmdQuerySubspaces() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subspaces",
		Short: "Get all the subspaces",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Subspaces(
				context.Background(),
				&types.QuerySubspacesRequest{},
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
