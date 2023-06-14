package cli

// DONTCOVER

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	cliutils "github.com/desmos-labs/desmos/v5/x/posts/client/utils"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

// NewTxCmd returns a new command to perform subspaces transactions
func NewTxCmd() *cobra.Command {
	subspacesTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Posts transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	subspacesTxCmd.AddCommand(
		GetCmdCreatePost(),
		GetCmdEditPost(),
		GetCmdDeletePost(),
		GetCmdAddPostAttachment(),
		GetCmdRemovePostAttachment(),
		GetCmdAnswerPoll(),
		GetCmdMovePost(),
	)

	return subspacesTxCmd
}

// GetCmdCreatePost returns the command allowing to create a new post
func GetCmdCreatePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [subspace-id] [section-id] [json-file-path]",
		Args:    cobra.ExactArgs(3),
		Short:   "Create a new post",
		Long:    `Create a new post containing the data specified inside the JSON file located at the provided path.`,
		Example: fmt.Sprintf(`%s tx posts create 1 1 /path/to/my/file.json --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			sectionID, err := subspacestypes.ParseSectionID(args[1])
			if err != nil {
				return err
			}

			data, err := cliutils.ParseCreatePostJSON(clientCtx.Codec, args[2])
			if err != nil {
				return err
			}

			author := clientCtx.FromAddress.String()

			attachments, err := types.UnpackAttachments(clientCtx.Codec, data.Attachments)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePost(
				subspaceID,
				sectionID,
				data.ExternalID,
				data.Text,
				data.ConversationID,
				data.ReplySettings,
				data.Entities,
				data.Tags,
				attachments,
				data.ReferencedPosts,
				author,
			)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditPost returns the command allowing to edit an existing post
func GetCmdEditPost() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "edit [subspace-id] [post-id] [json-file-path]",
		Args:    cobra.ExactArgs(3),
		Short:   "Edit an existing post",
		Long:    `Edit an existing post by using the data specified inside the JSON file located at the provided path.`,
		Example: fmt.Sprintf(`%s tx posts edit 1 1 /path/to/my/file.json --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			data, err := cliutils.ParseEditPostJSON(clientCtx.Codec, args[2])
			if err != nil {
				return err
			}

			editor := clientCtx.FromAddress.String()

			msg := types.NewMsgEditPost(
				subspaceID,
				postID,
				data.Text,
				data.Entities,
				data.Tags,
				editor,
			)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDeletePost returns the command allowing to delete an existing post
func GetCmdDeletePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [subspace-id] [post-id]",
		Args:    cobra.ExactArgs(2),
		Short:   "Delete an existing post",
		Example: fmt.Sprintf(`%s tx posts delete 1 1 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgDeletePost(subspaceID, postID, signer)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAddPostAttachment returns the command allowing to add an attachment to an existing post
func GetCmdAddPostAttachment() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-attachment [subspace-id] [post-id] [json-file-path]",
		Args:    cobra.ExactArgs(3),
		Short:   "Add an attachment to an existing post",
		Long:    `Add an attachment to an existing post by using the data specified inside the JSON file located at the provided path.`,
		Example: fmt.Sprintf(`%s tx posts add-attachment 1 1 /path/to/my/file.json --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			content, err := cliutils.ParseAttachmentContent(clientCtx.Codec, args[2])
			if err != nil {
				return err
			}

			editor := clientCtx.FromAddress.String()

			msg := types.NewMsgAddPostAttachment(subspaceID, postID, content, editor)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRemovePostAttachment returns the command allowing to delete an existing post attachment
func GetCmdRemovePostAttachment() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove-attachment [subspace-id] [post-id] [attachment-id]",
		Args:    cobra.ExactArgs(3),
		Short:   "Remove an existing post attachment",
		Example: fmt.Sprintf(`%s tx posts remove-attachment 1 1 1 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			attachmentID, err := types.ParseAttachmentID(args[2])
			if err != nil {
				return err
			}

			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgRemovePostAttachment(subspaceID, postID, attachmentID, signer)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAnswerPoll returns the command allowing to answer an existing poll
func GetCmdAnswerPoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "answer-poll [subspace-id] [post-id] [poll-id] [answers-indexes]",
		Args:  cobra.ExactArgs(4),
		Short: "Answer an existing poll",
		Long: `Answer an existing poll with the provided answers index.
If you want to specify multiple answers, separate them using a comma.

Answer indexes must be the indexes that each of your answer has inside the poll's provided answer array. 
E.g. Suppose you have a poll with the following provided answer array: 
- Question: What animal is the best?
  Provided answers:
  - Cat
  - Dog

Then, the "Cat" answer has index 0 and and the "Dog" answer has index 1.
`,
		Example: fmt.Sprintf(`%s tx posts answer-poll 1 1 1 0,1,2 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

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

			indexes := strings.Split(args[3], ",")
			answers := make([]uint32, len(indexes))
			for i, answer := range indexes {
				answerIndex, err := strconv.ParseUint(answer, 10, 32)
				if err != nil {
					return err
				}
				answers[i] = uint32(answerIndex)
			}

			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgAnswerPoll(subspaceID, postID, pollID, answers, signer)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdMovePost returns the command allowing to move an existing post to another subspace
func GetCmdMovePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "move-post [subspace-id] [post-id] [target-subspace-id] [target-section-id]",
		Args:    cobra.ExactArgs(4),
		Short:   "Move an existing post to another subspace",
		Example: fmt.Sprintf(`%s tx posts move-post 1 1 2 1 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			targetSubspaceID, err := subspacestypes.ParseSubspaceID(args[2])
			if err != nil {
				return err
			}

			targetSectionID, err := subspacestypes.ParseSectionID(args[3])
			if err != nil {
				return err
			}

			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgMovePost(subspaceID, postID, targetSubspaceID, targetSectionID, signer)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRequestPostOwnerTransfer returns the command to create a post owner transfer request
func GetCmdRequestPostOwnerTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "request-post-owner-transfer [subspace-id] [post-id] [receiver]",
		Short:   "Make a request to transfer the owner of the post to the receiver",
		Example: fmt.Sprintf(`%s tx posts request-post-owner-transfer 1 1 desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud`, version.AppName),
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestPostOwnerTransfer(subspaceID, postID, args[2], clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdCancelPostOwnerTransfer returns the command to cancel an outgoing post owner transfer request
func GetCmdCancelPostOwnerTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cancel-post-owner-transfer [subspace-id] [post-id]",
		Short:   "Cancel a post owner transfer request with the given post id",
		Example: fmt.Sprintf(`%s tx posts cancel-dtag-transfer-request 1, 1`, version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelPostOwnerTransferRequest(subspaceID, postID, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAcceptPostOwnerTransfer returns the command to accept a post owner transfer request
func GetCmdAcceptPostOwnerTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "accept-post-owner-transfer [subspace-id] [post-id]",
		Short:   `Accept a post owner transfer request with the given post id`,
		Example: fmt.Sprintf(`%s tx posts accept-dtag-transfer-request 1 1`, version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAcceptPostOwnerTransferRequest(subspaceID, postID, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRefusePostOwnerTransfer returns the command to refuse an incoming post owner transfer request
func GetCmdRefusePostOwnerTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "refuse-post-owner-transfer-request [subspace-id] [post-id]",
		Short:   "Refuse a post owner transfer request with the given post id",
		Example: fmt.Sprintf(`%s tx posts refuse-dtag-transfer-request 1 1`, version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgRefusePostOwnerTransferRequest(subspaceID, postID, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
