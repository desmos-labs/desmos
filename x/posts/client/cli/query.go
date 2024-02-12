package cli

// DONTCOVER

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v7/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

// GetQueryCmd returns the command allowing to perform queries
func GetQueryCmd() *cobra.Command {
	subspaceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the posts module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	subspaceQueryCmd.AddCommand(
		GetCmdQueryPost(),
		GetCmdQueryPosts(),
		GetCmdQueryPostAttachments(),
		GetCmdQueryPollAnswers(),
		GetCmdQueryParams(),
		GetCmdQueryPostOwnerTransferRequests(),
	)
	return subspaceQueryCmd
}

// GetCmdQueryPost returns the command to query the post having the given id
func GetCmdQueryPost() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "post [subspace-id] [post-id]",
		Short:   "Query the post with the given id",
		Example: fmt.Sprintf(`%s query posts post 1 1`, version.AppName),
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

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			res, err := queryClient.Post(context.Background(), types.NewQueryPostRequest(subspaceID, postID))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryPosts returns the command to query all the posts inside a subspace
func GetCmdQueryPosts() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "posts [subspace-id] [[section-id]]",
		Short:   "Query the posts inside a specific subspace with optional section",
		Example: fmt.Sprintf(`%s query posts posts 1 --page=2 --limit=100`, version.AppName),
		Args:    cobra.RangeArgs(1, 2),
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

			var res proto.Message
			if len(args) == 1 {
				res, err = queryClient.SubspacePosts(context.Background(), types.NewQuerySubspacePostsRequest(subspaceID, pageReq))
				if err != nil {
					return err
				}
			} else {
				sectionID, err := subspacestypes.ParseSectionID(args[1])
				if err != nil {
					return err
				}

				res, err = queryClient.SectionPosts(context.Background(), types.NewQuerySectionPostsRequest(subspaceID, sectionID, pageReq))
				if err != nil {
					return err
				}
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "posts")

	return cmd
}

// GetCmdQueryPostAttachments returns the command to query all the attachments for the post having the given id
func GetCmdQueryPostAttachments() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "attachments [subspace-id] [post-id]",
		Short:   "Query the attachments for the post having the given id",
		Example: fmt.Sprintf(`%s query posts attachments 1 1 --page=2 --limit=100`, version.AppName),
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

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.PostAttachments(context.Background(), types.NewQueryPostAttachmentsRequest(subspaceID, postID, pageReq))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "post attachments")

	return cmd
}

// GetCmdQueryPollAnswers returns the command to query all the answers to a given poll
func GetCmdQueryPollAnswers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "answers [subspace-id] [post-id] [poll-id] [[user]]",
		Short: "Query the answers related to a specific poll with an optional user",
		Long: `Query the answers related to the poll having the given id, and associated to the specified post inside the specified subspace.
If a user address is provided, only the answer of that user will be returned (if any).
`,
		Example: fmt.Sprintf(`
%s query posts answers 1 1 1 --page=2 --limit=100
%s query posts answers 1 1 1 desmos1mc0mrx23aawryc6gztvdyrupph00yz8lk42v40 --page=2 --limit=100
`, version.AppName, version.AppName),
		Args: cobra.RangeArgs(3, 4),
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

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			pollID, err := types.ParseAttachmentID(args[2])
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var user string
			if len(args) > 3 {
				user = args[3]
			}

			res, err := queryClient.PollAnswers(context.Background(), types.NewQueryPollAnswersRequest(subspaceID, postID, pollID, user, pageReq))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "poll answers")

	return cmd
}

// GetCmdQueryParams returns the command to query the module params
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "params",
		Short:   "Query the module parameters",
		Example: fmt.Sprintf(`%s query posts params`, version.AppName),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), types.NewQueryParamsRequest())
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryPostOwnerTransferRequests returns the command allowing to query all the post owner transfer requests made towards a user
func GetCmdQueryPostOwnerTransferRequests() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "incoming-post-owner-transfer-requests [subspace-id] [[receiver]]",
		Short: "Retrieve the post owner transfer requests with subspace id, optional address and pagination",
		Example: fmt.Sprintf(`%s query posts incoming-post-owner-transfer-requests
%s query posts incoming-post-owner-transfer-requests 1 --page=2 --limit=100
%s query posts incoming-post-owner-transfer-requests 1 desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
`, version.AppName, version.AppName, version.AppName),
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

			var receiver string
			if len(args) == 2 {
				receiver = args[1]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.IncomingPostOwnerTransferRequests(
				context.Background(),
				types.NewQueryIncomingPostOwnerTransferRequestsRequest(subspaceID, receiver, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "post owner transfer requests")

	return cmd
}
