package cli

// DONTCOVER

import (
	"context"
	"fmt"

	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v4/x/reactions/types"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

const (
	FlagUser = "user"
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
		GetCmdQueryReaction(),
		GetCmdQueryReactions(),
		GetCmdQueryRegisteredReaction(),
		GetCmdQueryRegisteredReactions(),
		GetCmdQueryParams(),
	)
	return queryCmd
}

// GetCmdQueryReaction returns the command to query a reaction by its id
func GetCmdQueryReaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reaction [subspace-id] [post-id] [reaction-id]",
		Short:   "Query the reaction for the given post having the provided id",
		Example: fmt.Sprintf(`%s query reactions reaction 1 1 1`, version.AppName),
		Args:    cobra.ExactArgs(3),
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

			postID, err := poststypes.ParsePostID(args[1])
			if err != nil {
				return err
			}

			reactionID, err := types.ParseReactionID(args[2])
			if err != nil {
				return err
			}

			res, err := queryClient.Reaction(
				context.Background(),
				types.NewQueryReactionRequest(subspaceID, postID, reactionID),
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

// GetCmdQueryReactions returns the command to query the reactions inside a subspace
func GetCmdQueryReactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reactions [subspace-id] [post-id]",
		Short: "Query the reactions inside the specified subspace with an optional post id",
		Example: fmt.Sprintf(`
%s query reactions reactions 1 1 --%s=cosmos14z8mn9ywhqu84alr5grxuljwj87jyz0zpxnlxy
`, version.AppName, FlagUser),
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

			postID, err := poststypes.ParsePostID(args[1])
			if err != nil {
				return err
			}

			user, err := cmd.Flags().GetString(FlagUser)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Reactions(
				context.Background(),
				types.NewQueryReactionsRequest(subspaceID, postID, user, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(FlagUser, "", "Optional address of the user to query the reactions for")

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "reactions")

	return cmd
}

// GetCmdQueryRegisteredReaction returns the command to query a registered reaction by its id
func GetCmdQueryRegisteredReaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "registered-reaction [subspace-id] [reaction-id]",
		Short:   "Query the registered reaction having the provided id",
		Example: fmt.Sprintf(`%s query reactions registered-reaction 1 1`, version.AppName),
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

			reactionID, err := types.ParseRegisteredReactionID(args[1])
			if err != nil {
				return err
			}

			res, err := queryClient.RegisteredReaction(
				context.Background(),
				types.NewQueryRegisteredReactionRequest(subspaceID, reactionID),
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

// GetCmdQueryRegisteredReactions returns the command to query the registered reactions of a subspace
func GetCmdQueryRegisteredReactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "registered-reactions [subspace-id]",
		Short:   "Query the registered reactions of a subspace",
		Example: fmt.Sprintf(`%s query reactions registered-reactions 1`, version.AppName),
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
