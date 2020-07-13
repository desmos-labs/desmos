package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/posts/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd adds the query commands
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	postQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the posts module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	postQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryPost(cdc),
		GetCmdQueryPosts(cdc),
		GetCmdQueryPollAnswer(cdc),
		GetCmdQueryRegisteredReactions(cdc),
		GetCmdQueryPostsParams(cdc),
	)...)
	return postQueryCmd
}

// GetCmdQueryPost queries a post
func GetCmdQueryPost(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "post [id]",
		Short: "Retrieve the post having the given id, if any.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			postID := args[0]

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryPost, postID)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find post with id %s \n", postID)
				return nil
			}

			var out types.PostQueryResponse
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdQueryPosts(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "posts",
		Short: "Query posts with optional filters",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for paginated posts that match optional filters:

Example:
$ %s query posts posts --creator desmos1qugw5ux0ea0v3cdxj7n9jnrz69f9wyc4668ek5
$ %s query posts posts --page=2 --limit=100
`,
				version.ClientName, version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			page := viper.GetInt(flagPage)
			limit := viper.GetInt(flagNumLimit)

			// Default params
			params := types.DefaultQueryPostsParams(page, limit)

			// SortBy
			if sortBy := viper.GetString(flagSortBy); len(sortBy) > 0 {
				params.SortBy = sortBy
			}

			// SortOrder
			if sortOrder := viper.GetString(flagSorOrder); len(sortOrder) > 0 {
				params.SortOrder = sortOrder
			}

			// ParentID
			if parentID := viper.GetString(flagParentID); len(parentID) > 0 {
				idParent := types.PostID(parentID)
				if !idParent.Valid() {
					return fmt.Errorf("invalid postID: %s", idParent)
				}
				params.ParentID = &idParent
			}

			// CreationTime
			if creationTime := viper.GetString(flagCreationTime); len(creationTime) > 0 {
				parsedTime, err := time.Parse(time.RFC3339, creationTime)
				if err != nil {
					return err
				}

				params.CreationTime = &parsedTime
			}

			// AllowsComments
			if allowsComments := viper.GetString(flagAllowsComments); len(allowsComments) > 0 {
				allowsCommentsBool, err := strconv.ParseBool(allowsComments)
				if err != nil {
					return err
				}
				params.AllowsComments = &allowsCommentsBool
			}

			// Subspace
			if subspace := viper.GetString(flagSubspace); len(subspace) > 0 {
				params.Subspace = subspace
			}

			// Hashtags
			if hashtags := viper.GetStringSlice(flagHashtag); len(hashtags) > 0 {
				params.Hashtags = hashtags
			}

			// Creator
			if bech32CreatorAddress := viper.GetString(flagCreator); len(bech32CreatorAddress) != 0 {
				depositorAddr, err := sdk.AccAddressFromBech32(bech32CreatorAddress)
				if err != nil {
					return err
				}
				params.Creator = depositorAddr
			}

			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryPosts)
			res, height, err := cliCtx.QueryWithData(route, bz)
			if err != nil {
				return err
			}

			var matchingPosts types.Posts
			err = cdc.UnmarshalJSON(res, &matchingPosts)
			if err != nil {
				return err
			}

			if matchingPosts == nil {
				matchingPosts = types.Posts{}
			}

			cliCtx = cliCtx.WithHeight(height)
			return cliCtx.PrintOutput(matchingPosts) // nolint:errcheck
		},
	}

	cmd.Flags().Int(flagPage, 1, "pagination page of posts to to query for")
	cmd.Flags().Int(flagNumLimit, 100, "pagination limit of posts to query for")

	cmd.Flags().String(flagSortBy, "", "(optional) sort the posts based on this field")
	cmd.Flags().String(flagSorOrder, "", "(optional) sort the posts using this order (ascending/descending)")

	cmd.Flags().String(flagParentID, "", "(optional) filter the posts with given parent id")
	cmd.Flags().String(flagCreationTime, "", "(optional) filter the posts created at block height")
	cmd.Flags().String(flagAllowsComments, "", "(optional) filter the posts allowing comments")
	cmd.Flags().String(flagSubspace, "", "(optional) filter the posts part of the subspace")
	cmd.Flags().String(flagCreator, "", "(optional) filter the posts created by creator")
	cmd.Flags().StringSlice(flagHashtag, []string{}, "(optional) filter the posts that contain the specified hashtags")

	return cmd
}

func GetCmdQueryPollAnswer(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "poll-answers [id]",
		Short: "Retrieve tha poll answers of the post with given id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			postID := args[0]

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryPollAnswers, postID)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find post with id %s \n", postID)
				return nil
			}

			var out types.PollAnswersQueryResponse
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdQueryRegisteredReactions(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "registered-reactions",
		Short: "Retrieve tha poll answers of the post with given id",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryRegisteredReactions)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find any registered reaction \n")
				return nil
			}

			var out types.Reactions
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdQueryPostsParams queries all the posts' module params
func GetCmdQueryPostsParams(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Retrieve all the posts module params",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParams)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find profile params")
				return nil
			}

			var out types.Params
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
