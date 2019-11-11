package cli

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// GetTxCmd set the tx commands
func GetTxCmd(_ string, cdc *codec.Codec) *cobra.Command {
	postsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Posts transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	postsTxCmd.AddCommand(client.PostCommands(
		GetCmdCreatePost(cdc),
		GetCmdEditPost(cdc),
		GetCmdAddLike(cdc),
	)...)

	return postsTxCmd
}

// GetCmdCreatePost is the CLI command for creating a post
func GetCmdCreatePost(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create [message] [parent-post-id] [namespace] [external address]",
		Short: "Create a new post",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)
			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			parentID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePost(args[0], parentID, from, args[2], args[3])
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdEditPost is the CLI command for editing a post
func GetCmdEditPost(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "edit [post-id] [message]",
		Short: "Edit a post you have previously created",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)
			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgEditPost(postID, args[1], time.Now(), from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdAddLike is the CLI command for adding a like to a post
func GetCmdAddLike(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "like [post-id]",
		Short: "Like a post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)
			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgLikePost(postID, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
