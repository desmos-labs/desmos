package cli

// DONTCOVER

import (
	"context"
	"fmt"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v3/x/reactions/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// GetQueryCmd returns the command allowing to perform queries
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the reactions module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	queryCmd.AddCommand(
		GetCmdQueryReactions(),
		GetCmdQueryRegisteredReactions(),
		GetCmdQueryParams(),
	)
	return queryCmd
}

// GetCmdQueryReactions returns the command to query the reactions inside a subspace
func GetCmdQueryReactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reactions [subspace-id]",
		Short: "Query the reactions inside the specified subspace with an optional post id",
		Example: fmt.Sprintf(`
%s query reactions reactions 1
%s query reactions reactions 1 1
`, version.AppName, version.AppName),
		Args: cobra.RangeArgs(1, 2),
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

			var postID uint64
			if len(args) > 1 {
				postID, err = poststypes.ParsePostID(args[1])
				if err != nil {
					return err
				}
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Reactions(
				context.Background(),
				types.NewQueryReactionsRequest(subspaceID, postID, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "reactions")

	return cmd
}

// GetCmdQueryRegisteredReactions returns the command to query the registered reactions of a subspace
func GetCmdQueryRegisteredReactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "registered-reactions [subspace-id]",
		Short:   "Query the registered reactions of a subspace",
		Example: fmt.Sprintf(`%s query reactions registered-reactions 1`, version.AppName),
		Args:    cobra.ExactArgs(2),
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

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.RegisteredReactions(
				context.Background(),
				types.NewQueryRegisteredReactionsRequest(subspaceID, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "registered reactions")

	return cmd
}

// GetCmdQueryParams returns the command to query the reaction params of a subspace
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "params [subspace-id]",
		Short:   "Query the reaction params of a subspace",
		Example: fmt.Sprintf(`%s query reactions params 1`, version.AppName),
		Args:    cobra.ExactArgs(1),
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

			res, err := queryClient.ReactionsParams(
				context.Background(),
				types.NewQueryReactionsParamsRequest(subspaceID),
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
