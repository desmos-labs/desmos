package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/client/tx"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// NewTxCmd returns a new command allowing to perform posts transactions
func NewTxCmd() *cobra.Command {
	postsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Posts transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	postsTxCmd.AddCommand(
		GetCmdCreatePost(),
		GetCmdEditPost(),
		GetCmdAddPostReaction(),
		GetCmdRemovePostReaction(),
		GetCmdAnswerPoll(),
		GetCmdRegisterReaction(),
	)

	return postsTxCmd
}

// getAttachments parses the attachments of a post.
// If one or more attachments are found, it returns them. Otherwise returns `nil` instead.
func getAttachments(cmd *cobra.Command) (types.Attachments, error) {
	mediasStrings, err := cmd.Flags().GetStringArray(FlagAttachment)
	if err != nil {
		return nil, fmt.Errorf("invalid flag value: %s", FlagAttachment)
	}

	attachments := types.Attachments{}
	for _, mediaString := range mediasStrings {
		argz := strings.Split(mediaString, ",")
		var tags []string

		// If some tags are specified
		if len(argz) < 2 {
			return nil, fmt.Errorf("if attachments are specified, the arguments has to be at least 2 and in this order: \"URI,Mime-Type\", please use the --help flag to know more")
		} else if len(argz) > 2 {
			tags = append(tags, argz[2:]...)
		}

		attachment := types.NewAttachment(argz[0], argz[1], tags)
		attachments = attachments.AppendIfMissing(attachment)
	}

	return attachments, nil
}

// getPollData parses the pollData of a post. If no poll data is found returns `nil` instead.
func getPollData(cmd *cobra.Command) (*types.PollData, error) {
	pollDetailsMap, err := cmd.Flags().GetStringToString(FlagPollDetails)
	if err != nil {
		return nil, fmt.Errorf("invalid %s value", FlagPollDetails)
	}

	pollAnswersSlice := viper.GetStringSlice(FlagPollAnswer)
	if len(pollDetailsMap) == 0 && len(pollAnswersSlice) > 0 {
		return nil, fmt.Errorf("poll answers specified but no poll details found. Please use %s to specify the poll details", FlagPollDetails)
	}

	if len(pollDetailsMap) > 0 && len(pollAnswersSlice) == 0 {
		return nil, fmt.Errorf("poll details specified but answers are not. Please use the %s to specify one or more answer", FlagPollAnswer)
	}

	var pollData *types.PollData
	if len(pollDetailsMap) > 0 && len(pollAnswersSlice) > 0 {
		date, err := time.Parse(time.RFC3339, pollDetailsMap[keyEndDate])
		if err != nil {
			return nil, fmt.Errorf(
				"end date should be provided in RFC3339 format, e.g 2020-01-01T12:00:00Z, %s found",
				pollDetailsMap[keyEndDate],
			)
		}

		if date.Before(time.Now().UTC()) {
			return nil, fmt.Errorf("poll's end date can't be in the past")
		}

		if len(strings.TrimSpace(pollDetailsMap[keyQuestion])) == 0 {
			return nil, fmt.Errorf("question should be provided and not be empty")
		}

		question := pollDetailsMap[keyQuestion]

		allowMultipleAnswers, err := strconv.ParseBool(pollDetailsMap[keyMultipleAnswers])
		if err != nil {
			return nil, fmt.Errorf("multiple-answers can only be true or false")
		}

		allowsAnswerEdits, err := strconv.ParseBool(pollDetailsMap[keyAllowsAnswerEdits])
		if err != nil {
			return nil, fmt.Errorf("allows-answer-edits can only be only true or false")
		}

		answers := types.PollAnswers{}
		for index, answer := range pollAnswersSlice {
			if strings.TrimSpace(answer) == "" {
				return nil, fmt.Errorf("invalid answer text at index %d", index)
			}

			pollAnswer := types.NewPollAnswer(fmt.Sprint(index), answer)
			answers = answers.AppendIfMissing(pollAnswer)
		}

		pollData = types.NewPollData(question, date, answers, allowMultipleAnswers, allowsAnswerEdits)
	}

	return pollData, nil
}

// GetCmdCreatePost returns the CLI command to create a post
func GetCmdCreatePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [subspace] [[message]]",
		Short: "Create a new post",
		Long: fmt.Sprintf(`
Create a new post specifying the subspace and the message (optional if any kind of attachment is provided).
Optional attachments and polls are also supported. See the below sections to know how to include them.

E.g.
%s tx posts create "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e" "Hello world!"

Comments to the post could be locked by including the --allows-comments flag.
By default this field is set to true.

%s tx posts create "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e" "Hello world!" \
   --allows-comments false

=== Attachments ===
If you want to add one or more attachment(s), you have to use the --attachment flag.
You need to firstly specify the attachment URI and then its mime-type separeted by a comma.
You can also specify the desmos addresses tagged in the attachment you're sharing by adding as 
many address you want after the mime-type separated by a comma.

%s tx posts create "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e" "A post with a single attachment" \
  --attachment "https://example.com/attachment1,text/plain,desmos1ulmv2dyc8zjmhk9zlsq4ajpudwc8zjfm82aysr" \
  --allows-comments false
%s tx posts create "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e" "A post with multiple attachments" \
  --attachment "https://example.com/attachment1,text/plain,desmos1ulmv2dyc8zjmhk9zlsq4ajpudwc8zjfm82aysr" \
  --attachment "https://example.com/attachment2,application/json"

If attachments are provided, the post could be created even without any message as following:

%s tx posts create "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e" \
  --attachment "https://example.com/attachment1,text/plain,desmos1ulmv2dyc8zjmhk9zlsq4ajpudwc8zjfm82aysr" \
  --attachment "https://example.com/attachment2,application/json" \
  --allows-comments false

=== Polls ===
If you want to add a poll to your post you need to specify it through two flags:
  1. --poll-details, which accepts a map with the following keys:
     * question: the question of the poll
     * date: the end date of your poll after which no further answers will be accepted
     * multiple-answers: a boolean indicating the possibility of multiple answers from users
     * allows-answers-edits: a boolean value that indicates the possibility to edit the answers in the future
  2. --poll-answer, which accepts a slice of answers that will be provided to the users once they want to take part in the poll votations.	
     Each answer should be identified by the text of the answer itself.

If a poll is provided, the post can be created even without specifying any message as follows:

E.g.
%s tx posts create "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e" "Post with poll" \
	--poll-details "question=Which dog do you prefer?,multiple-answers=false,allows-answer-edits=true,end-date=2020-01-01T15:00:00.000Z" \
	--poll-answer "Beagle" \
	--poll-answer "Pug" \
	--poll-answer "German Sheperd"
`, version.AppName, version.AppName, version.AppName, version.AppName, version.AppName, version.AppName),
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Check parent id
			parentID := viper.GetString(FlagParentID)
			if parentID != "" && !types.IsValidPostID(parentID) {
				return sdkerrors.Wrap(types.ErrInvalidPostID, parentID)
			}

			// Check for attachments
			attachments, err := getAttachments(cmd)
			if err != nil {
				return err
			}

			// Check for poll
			pollData, err := getPollData(cmd)
			if err != nil {
				return err
			}

			text := ""
			if len(args) > 1 {
				text = args[1]
			}

			msg := types.NewMsgCreatePost(
				text,
				parentID,
				viper.GetBool(FlagAllowsComments),
				args[0],
				nil,
				clientCtx.FromAddress.String(),
				attachments,
				pollData,
			)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Bool(FlagAllowsComments, true, "Possibility to comment the post or not")
	cmd.Flags().String(FlagParentID, "", "Id of the post to which this one should be an answer to")
	cmd.Flags().StringArray(FlagAttachment, []string{}, "Current post's attachment")
	cmd.Flags().StringToString(FlagPollDetails, map[string]string{}, "Current post's poll details")
	cmd.Flags().StringSlice(FlagPollAnswer, []string{}, "Current post's poll answer")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditPost returns the CLI command to edit a post
func GetCmdEditPost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [post-id] [[message]]",
		Short: "Edit a post you have previously created",
		Long: fmt.Sprintf(`Edit a post by specifying its ID.
E.g.
%s tx posts edit "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af" "Edit"

You can also edit post's attachments and pollData by providing the right flags like you do when creating a post.

=== Attachments ===
If you want to edit attachments, you have to use the --attachment flag.

%s tx posts edit "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af" "Edit a post with attachments" \
  --attachment "https://example.com/attachment1,text/plain,desmos1ulmv2dyc8zjmhk9zlsq4ajpudwc8zjfm82aysr" \

=== Polls ===
If you want to edit post's poll you need to specify it through two flags:
  1. --poll-details, which accepts a map with the following keys:
     * question: the question of the poll
     * date: the end date of your poll after which no further answers will be accepted
     * multiple-answers: a boolean indicating the possibility of multiple answers from users
     * allows-answers-edits: a boolean value that indicates the possibility to edit the answers in the future
  2. --poll-answer, which accepts a slice of answers that will be provided to the users once they want to take part in the poll votations.	
     Each answer should be identified by the text of the answer itself.

If a poll is provided, the post can be edited even without specifying any message:

E.g.
%s tx posts edit "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af" "Edit a post with attachments" \
	--poll-details "question=Which dog do you prefer?,multiple-answers=false,allows-answer-edits=true,end-date=2020-01-01T15:00:00.000Z" \
	--poll-answer "Beagle" \
	--poll-answer "Pug" \
	--poll-answer "German Sheperd"
`, version.AppName, version.AppName, version.AppName),
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			postID := args[0]
			if !types.IsValidPostID(postID) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, postID)
			}

			// Check for attachments
			attachments, err := getAttachments(cmd)
			if err != nil {
				return err
			}

			// Check for poll
			pollData, err := getPollData(cmd)
			if err != nil {
				return err
			}

			text := ""
			if len(args) > 1 {
				text = args[1]
			}

			msg := types.NewMsgEditPost(
				postID,
				text,
				attachments,
				pollData,
				clientCtx.FromAddress.String(),
			)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringArray(FlagAttachment, []string{}, "Current post's attachment")
	cmd.Flags().StringToString(FlagPollDetails, map[string]string{}, "Current post's poll details")
	cmd.Flags().StringSlice(FlagPollAnswer, []string{}, "Current post's poll answer")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAddPostReaction returns the CLI command to add a reaction to a post
func GetCmdAddPostReaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-reaction [post-id] [value]",
		Short: "Adds a reaction to a post",
		Long: fmt.Sprintf(`
Add a reaction to the post having the given id with the specified value. 
The value has to be a reaction short code.

E.g. 
%s tx posts add-reaction a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc :thumbsup: --from jack
`, version.AppName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			postID := args[0]
			if !types.IsValidPostID(postID) {
				return sdkerrors.Wrap(types.ErrInvalidPostID, postID)
			}

			msg := types.NewMsgAddPostReaction(postID, args[1], clientCtx.GetFromAddress().String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRemovePostReaction returns the CLI command to remove a reaction from a post
func GetCmdRemovePostReaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-reaction [post-id] [value]",
		Short: "Removes an existing reaction from a post",
		Long: fmt.Sprintf(`
Removes the reaction having the given value from the post having the given id. 
The value has to be a reaction short code.

E.g. 
%s tx posts remove-reaction a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc :thumbsup: --from jack
`, version.AppName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			postID := args[0]
			if !types.IsValidPostID(postID) {
				return sdkerrors.Wrap(types.ErrInvalidPostID, postID)
			}

			msg := types.NewMsgRemovePostReaction(postID, clientCtx.FromAddress.String(), args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAnswerPoll returns the CLI command to answer a poll
func GetCmdAnswerPoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "answer-poll [post-id] [answer...]",
		Short: "Answer a post's poll'",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			postID := args[0]
			if !types.IsValidPostID(postID) {
				return sdkerrors.Wrap(types.ErrInvalidPostID, postID)
			}

			var answers []string
			for i := 1; i < len(args); i++ {
				answers = append(answers, args[i])
			}

			msg := types.NewMsgAnswerPoll(postID, answers, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRegisterReaction returns the CLI command to register a new reaction
func GetCmdRegisterReaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-reaction [short-code] [value] [subspace]",
		Short: "Register a new reaction",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, found := types.GetEmojiByShortCodeOrValue(args[1]); found {
				return fmt.Errorf("%s represents an emoji shortcode and thus cannot be used to register another reaction", args[1])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterReaction(clientCtx.FromAddress.String(), args[0], args[1], args[2])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
