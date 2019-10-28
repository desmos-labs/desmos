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
	"github.com/kwunyeung/desmos/x/magpie/internal/types"
)

// GetTxCmd set the tx commands
func GetTxCmd(_ string, cdc *codec.Codec) *cobra.Command {
	magpieTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "magpie transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	magpieTxCmd.AddCommand(client.PostCommands(
		GetCmdCreateSession(cdc),
		client.LineBreak,
		GetCmdCreatePost(cdc),
		GetCmdEditPost(cdc),
		GetCmdAddLike(cdc),
	)...)

	return magpieTxCmd
}

// GetCmdCreatePost is the CLI command for creating a post
func GetCmdCreatePost(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create [message] [parent-post-id] [namespace] [external address]",
		Short: "create a new post",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)
			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			parentId, err := types.ParsePostId(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePost(args[0], parentId, time.Now(), from, args[2], args[3])
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
		Short: "edit an owned post",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)
			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			postId, err := types.ParsePostId(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgEditPost(postId, args[1], time.Now(), from)
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
		Use:   "like [post-id] [namespace] [external address]",
		Short: "like a post",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)
			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			postId, err := types.ParsePostId(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgLike(postId, time.Now(), from, args[1], args[2])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdCreateSession is the CLI command for creating a session for create post
func GetCmdCreateSession(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-session [namespace] [external address] [pubkey] [external signer signature]",
		Short: "record a session for external service to post a magpie",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			from := cliCtx.GetFromAddress()
			accGetter := authtypes.NewAccountRetriever(cliCtx)
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			msg := types.NewMsgCreateSession(time.Now(), from, args[0], args[1], args[2], args[3])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
