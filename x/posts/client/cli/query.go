package cli

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/posts/types"
)

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
		GetCmdQueryPosts(),
		GetCmdQueryPollAnswers(),
		GetCmdQueryRegisteredReactions(),
		GetCmdQueryParams(),
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
		Use:   "posts",
		Short: "Query posts with optional filters",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for paginated posts that match optional filters:

Example:
$ %s query posts posts --creator desmos1qugw5ux0ea0v3cdxj7n9jnrz69f9wyc4668ek5
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

			page := viper.GetUint64(flagPage)
			limit := viper.GetUint64(flagNumLimit)

			// Default params
			params := DefaultQueryPostsRequest(page, limit)

			// SortBy
			if sortBy := viper.GetString(flagSortBy); len(sortBy) > 0 {
				params.SortBy = sortBy
			}

			// SortOrder
			if sortOrder := viper.GetString(flagSorOrder); len(sortOrder) > 0 {
				params.SortOrder = sortOrder
			}

			// ParentID
			if parentID := viper.GetString(FlagParentID); len(parentID) > 0 {
				idParent := parentID
				if !types.IsValidPostID(idParent) {
					return fmt.Errorf("invalid postID: %s", idParent)
				}
				params.ParentID = parentID
			}

			// CreationTime
			if creationTime := viper.GetString(FlagCreationTime); len(creationTime) > 0 {
				parsedTime, err := time.Parse(time.RFC3339, creationTime)
				if err != nil {
					return err
				}

				params.CreationTime = &parsedTime
			}

			// Subspace
			if subspace := viper.GetString(FlagSubspace); len(subspace) > 0 {
				params.Subspace = subspace
			}

			// Hashtags
			if hashtags := viper.GetStringSlice(FlagHashtag); len(hashtags) > 0 {
				params.Hashtags = hashtags
			}

			// Creator
			if bech32CreatorAddress := viper.GetString(FlagCreator); len(bech32CreatorAddress) != 0 {
				depositorAddr, err := sdk.AccAddressFromBech32(bech32CreatorAddress)
				if err != nil {
					return err
				}
				params.Creator = depositorAddr.String()
			}

			res, err := queryClient.Posts(context.Background(), &params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().Uint64(flagPage, 1, "pagination page of posts to to query for")
	cmd.Flags().Uint64(flagNumLimit, 100, "pagination limit of posts to query for")

	cmd.Flags().String(flagSortBy, "", "(optional) sort the posts based on this field")
	cmd.Flags().String(flagSorOrder, "", "(optional) sort the posts using this order (ascending/descending)")

	cmd.Flags().String(FlagParentID, "", "(optional) filter the posts with given parent id")
	cmd.Flags().String(FlagCreationTime, "", "(optional) filter the posts created at block height")
	cmd.Flags().String(FlagSubspace, "", "(optional) filter the posts part of the subspace")
	cmd.Flags().String(FlagCreator, "", "(optional) filter the posts created by creator")
	cmd.Flags().StringSlice(FlagHashtag, []string{}, "(optional) filter the posts that contain the specified hashtags")

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryPollAnswers returns the command allowing to query the answers of a poll
func GetCmdQueryPollAnswers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "poll-answers [id]",
		Short: "Retrieve tha poll answers of the post with given id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.PollAnswers(
				context.Background(),
				&types.QueryPollAnswersRequest{PostId: args[0]},
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

// GetCmdQueryRegisteredReactions returns the command allowing to query the registered reactions
func GetCmdQueryRegisteredReactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registered-reactions",
		Short: "Retrieve tha poll answers of the post with given id",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.RegisteredReactions(context.Background(), &types.QueryRegisteredReactionsRequest{})
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
