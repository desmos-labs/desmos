package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/kwunyeung/desmos/x/magpie/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd adds the query commands
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	magpieQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the magpie module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	magpieQueryCmd.AddCommand(client.GetCommands(
		GetCmdPost(storeKey, cdc),
		GetCmdLike(storeKey, cdc),
		GetCmdSession(storeKey, cdc),
	)...)
	return magpieQueryCmd
}

// GetCmdPost queries a post
func GetCmdPost(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "post [id]",
		Short: "post id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			postID := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/post/%s", queryRoute, postID), nil)
			if err != nil {
				fmt.Printf("could not find post - %s \n", postID)
				return nil
			}

			var out types.QueryResPost
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

			var out types.Like
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdSession queries a session by ID
func GetCmdSession(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "session [id]",
		Short: "Get the session by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			sessionsID := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/session/%s", queryRoute, sessionsID), nil)

			if err != nil {
				fmt.Printf("could not find session - %s \n", sessionsID)
				return nil
			}

			var out types.Session
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// // GetCmdNames queries a list of all names
// func GetCmdNames(queryRoute string, cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "names",
// 		Short: "names",
// 		// Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)

// 			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/names", queryRoute), nil)
// 			if err != nil {
// 				fmt.Printf("could not get query names\n")
// 				return nil
// 			}

// 			var out types.QueryResNames
// 			cdc.MustUnmarshalJSON(res, &out)
// 			return cliCtx.PrintOutput(out)
// 		},
// 	}
// }
