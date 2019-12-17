package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd adds the query commands
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	postQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the posts module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	postQueryCmd.AddCommand(client.GetCommands(
		GetCmdPost(storeKey, cdc),
		GetCmdLike(storeKey, cdc),
	)...)
	return postQueryCmd
}

// GetCmdPost queries a post
func GetCmdPost(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "post [id]",
		Short: "Retrieve the post having the given id, if any.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			postID := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/post/%s", queryRoute, postID), nil)
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

// GetCmdLike queries a like
func GetCmdLike(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "like [id]",
		Short: "like id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			likeID := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/like/%s", queryRoute, likeID), nil)
			if err != nil {
				fmt.Printf("could not find like - %s \n", likeID)
				return nil
			}

			var out types.Reaction
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
