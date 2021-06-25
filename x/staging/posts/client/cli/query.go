package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

//DONTCOVER
// TODO remove the above when x/posts is out of staging

// GetQueryCmd returns the command allowing to perform queries
func GetQueryCmd() *cobra.Command {
	postQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the posts module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	postQueryCmd.AddCommand(
		GetCmdQueryPost(),
		GetCmdQueryReports(),
		GetCmdQueryPosts(),
		GetCmdQueryUserAnswers(),
		GetCmdQueryRegisteredReactions(),
		GetCmdQueryParams(),
		GetCmdQueryPostReactions(),
	)
	return postQueryCmd
}

// GetCmdQueryPost returns the command allowing to query a post
func GetCmdQueryPost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post [id]",
		Short: "Retrieve the post having the given id, if any.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Post(
				context.Background(),
				&types.QueryPostRequest{PostId: args[0]},
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

// GetCmdQueryPosts returns the command allowing to query a list of posts
func GetCmdQueryPosts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "posts [subspace-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query posts with optional pagination",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for paginated posts inside the subspace:

Example:
$ %s query posts posts 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e
$ %s query posts posts --page=2 --limit=100
`,
				version.AppName, version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Posts(context.Background(), &types.QueryPostsRequest{SubspaceId: args[0], Pagination: pageReq})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(FlagSubspace, "", "(optional) filter the posts part of the subspace")

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, types.QueryPosts)

	return cmd
}

// GetCmdQueryUserAnswers returns the command allowing to query the answers of a poll
func GetCmdQueryUserAnswers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user-answers [id] [[user]]",
		Short: "Retrieve the user answers of the post with given id and the given user address",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			postID := args[0]
			var user string
			if len(args) == 2 {
				user = args[1]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.UserAnswers(
				context.Background(),
				&types.QueryUserAnswersRequest{PostId: postID, User: user, Pagination: pageReq},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, types.QueryUserAnswers)

	return cmd
}

// GetCmdQueryRegisteredReactions returns the command allowing to query the registered reactions
func GetCmdQueryRegisteredReactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registered-reactions [[subspace-id]]",
		Short: "Retrieve the registered reactions with optional subspace",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var subspaceID string
			if len(args) == 1 {
				subspaceID = args[0]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.RegisteredReactions(context.Background(), &types.QueryRegisteredReactionsRequest{SubspaceId: subspaceID, Pagination: pageReq})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, types.QueryRegisteredReactions)

	return cmd
}

// GetCmdQueryReports returns the command that allows to query the reports of a post
func GetCmdQueryReports() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reports [id]",
		Short: "Returns all the reports of the posts with the given id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Reports(
				context.Background(),
				&types.QueryReportsRequest{PostId: args[0]},
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

// GetCmdQueryParams returns the command allowing to query the module params
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parameters",
		Short: "Retrieve all the posts module parameters",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(
				context.Background(),
				&types.QueryParamsRequest{},
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

// GetCmdQueryPostReactions returns the command allowing to query the reactions of a post
func GetCmdQueryPostReactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-reactions [post-id]",
		Short: "Retrieve the reactions of the post having the given id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.PostReactions(
				context.Background(),
				&types.QueryPostReactionsRequest{PostId: args[0], Pagination: pageReq},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, types.QueryPostReactions)

	return cmd
}

// GetCmdQueryPostComments returns the command allowing to query the comments of a post
func GetCmdQueryPostComments() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post-comments [post-id]",
		Short: "Retrieve tha comments of the post with the given id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.PostComments(
				context.Background(),
				&types.QueryPostCommentsRequest{PostId: args[0], Pagination: pageReq},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, types.QueryPostComments)

	return cmd
}
