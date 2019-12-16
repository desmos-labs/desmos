package cli

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/viper"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
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
	postQueryCmd.AddCommand(client.GetCommands(
		GetCmdQueryPost(cdc),
		GetCmdQueryPosts(cdc),
		GetCmdQueryLike(cdc),
	)...)
	return postQueryCmd
}

// GetCmdQueryPost queries a post
func GetCmdQueryPost(queryRoute string, cdc *codec.Codec) *cobra.Command {
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
				fmt.Printf("could not find post - %s \n", postID)
				return nil
			}

			var out types.Post
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdQueryLike queries a like
func GetCmdQueryLike(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "like [id]",
		Short: "like id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			likeID := args[0]

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryLike, likeID)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("could not find like - %s \n", likeID)
				return nil
			}

			var out types.Like
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
			bech32CreatorAddress := viper.GetString(flagCreator)
			parentID := viper.GetString(flagParentID)
			creationTime := viper.GetInt(flagCreationTime)

			var creatorAddr sdk.AccAddress

			params := types.NewQueryPostsParams(page, limit, nil, sdk.NewInt(-1), creatorAddr)

			if len(bech32CreatorAddress) != 0 {
				depositorAddr, err := sdk.AccAddressFromBech32(bech32CreatorAddress)
				if err != nil {
					return err
				}
				params.Creator = depositorAddr
			}

			if len(parentID) > 0 {
				parentID, err := types.ParsePostID(parentID)
				if err != nil {
					return err
				}

				params.ParentID = &parentID
			}

			if creationTime >= 0 {
				params.CreationTime = sdk.NewInt(int64(creationTime))
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
	cmd.Flags().String(flagCreator, "", "(optional) filter the posts created by creator")
	cmd.Flags().String(flagParentID, "", "(optional) filter the posts with given parent id")
	cmd.Flags().Int(flagCreationTime, -1, "(optional) filter the posts created at block height")

	return cmd
}
