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

	cliutils "github.com/desmos-labs/desmos/v4/x/posts/client/utils"
	"github.com/desmos-labs/desmos/v4/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
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
		GetCmdChangePostOwner(),
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

// GetCmdChangePostOwner returns the command allowing to change the owner of a existing post
func GetCmdChangePostOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "change-post-owner [subspace-id] [post-id] [new-owner]",
		Args:    cobra.ExactArgs(3),
		Short:   "Change the owner of a existing post",
		Example: fmt.Sprintf(`%s tx posts change-post-owner 1 1 desmos1e209r8nc8qdkmqujahwrq4xrlxhk3fs9k7yzmw --from alice`, version.AppName),
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

			newOwner := args[2]
			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgChangePostOwner(subspaceID, postID, newOwner, signer)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
