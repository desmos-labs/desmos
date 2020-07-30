package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/desmos-labs/desmos/x/posts/types"
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

	postsTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreatePost(cdc),
		GetCmdEditPost(cdc),
		GetCmdAddPostReaction(cdc),
		GetCmdRemovePostReaction(cdc),
		GetCmdAnswerPoll(cdc),
		GetCmdRegisterReaction(cdc),
	)...)

	return postsTxCmd
}

// GetCmdCreatePost is the CLI command for creating a post
func GetCmdCreatePost(cdc *codec.Codec) *cobra.Command {
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
     * date: the end date of your poll after which no further answers will be accepted
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
`, version.ClientName, version.ClientName, version.ClientName, version.ClientName, version.ClientName, version.ClientName),
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			allowsComments := viper.GetBool(flagAllowsComments)
			parentID := types.PostID(viper.GetString(flagParentID))

			// attachments' checks
			mediasStrings, err := cmd.Flags().GetStringArray(flagAttachment)
			if err != nil {
				return fmt.Errorf("invalid flag value: %s", flagAttachment)
			}

			attachments := types.Attachments{}
			for _, mediaString := range mediasStrings {
				argz := strings.Split(mediaString, ",")
				var tags []sdk.AccAddress
				// if some tags are specified
				if len(argz) < 2 {
					return fmt.Errorf("if attachments are specified, the arguments has to be at least 2 and in this order: \"URI,Mime-Type\", please use the --help flag to know more")
				} else if len(argz) > 2 {
					for _, addr := range argz[2:] {
						tag, err := sdk.AccAddressFromBech32(addr)
						if err != nil {
							return err
						}
						tags = append(tags, tag)
					}
				}
				attachment := types.NewAttachment(argz[0], argz[1], tags)
				attachments = attachments.AppendIfMissing(attachment)
			}

			// polls' checks
			pollDetailsMap, err := cmd.Flags().GetStringToString(flagPollDetails)
			if err != nil {
				return fmt.Errorf("invalid %s value", flagPollDetails)
			}

			pollAnswersSlice := viper.GetStringSlice(flagPollAnswer)
			if len(pollDetailsMap) == 0 && len(pollAnswersSlice) > 0 {
				return fmt.Errorf("poll answers specified but no poll details found. Please use %s to specify the poll details", flagPollDetails)
			}

			if len(pollDetailsMap) > 0 && len(pollAnswersSlice) == 0 {
				return fmt.Errorf("poll details specified but answers are not. Please use the %s to specify one or more answer", flagPollAnswer)
			}

			var pollData *types.PollData
			if len(pollDetailsMap) > 0 && len(pollAnswersSlice) > 0 {
				date, err := time.Parse(time.RFC3339, pollDetailsMap[keyEndDate])
				if err != nil {
					return fmt.Errorf(
						"end date should be provided in RFC3339 format, e.g 2020-01-01T12:00:00Z, %s found",
						pollDetailsMap[keyEndDate],
					)
				}

				if date.Before(time.Now().UTC()) {
					return fmt.Errorf("poll's end date can't be in the past")
				}

				if len(strings.TrimSpace(pollDetailsMap[keyQuestion])) == 0 {
					return fmt.Errorf("question should be provided and not be empty")
				}

				question := pollDetailsMap[keyQuestion]

				allowMultipleAnswers, err := strconv.ParseBool(pollDetailsMap[keyMultipleAnswers])
				if err != nil {
					return fmt.Errorf("multiple-answers can only be true or false")
				}

				allowsAnswerEdits, err := strconv.ParseBool(pollDetailsMap[keyAllowsAnswerEdits])
				if err != nil {
					return fmt.Errorf("allows-answer-edits can only be only true or false")
				}

				answers := types.PollAnswers{}
				for index, answer := range pollAnswersSlice {
					if strings.TrimSpace(answer) == "" {
						return fmt.Errorf("invalid answer text at index %s", string(index))
					}

					pollAnswer := types.PollAnswer{
						ID:   types.AnswerID(index),
						Text: answer,
					}

					answers = answers.AppendIfMissing(pollAnswer)
				}

				pollData = &types.PollData{
					Question:              question,
					EndDate:               date,
					ProvidedAnswers:       answers,
					AllowsMultipleAnswers: allowMultipleAnswers,
					AllowsAnswerEdits:     allowsAnswerEdits,
				}
			}

			text := ""
			if len(args) > 1 {
				text = args[1]
			}

			msg := types.NewMsgCreatePost(
				text,
				parentID,
				allowsComments,
				args[0],
				map[string]string{},
				cliCtx.GetFromAddress(),
				attachments,
				pollData,
			)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().Bool(flagAllowsComments, true, "Possibility to comment the post or not")
	cmd.Flags().String(flagParentID, "", "Id of the post to which this one should be an answer to")
	cmd.Flags().StringArray(flagAttachment, []string{}, "Current post's attachment")
	cmd.Flags().StringToString(flagPollDetails, map[string]string{}, "Current post's poll details")
	cmd.Flags().StringSlice(flagPollAnswer, []string{}, "Current post's poll answer")

	return cmd
}

// GetCmdEditPost is the CLI command for editing a post
func GetCmdEditPost(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "edit [post-id] [message]",
		Short: "Edit a post you have previously created",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			postID := types.PostID(args[0])
			if !postID.Valid() {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid postID: %s", postID))
			}

			msg := types.NewMsgEditPost(postID, args[1], cliCtx.GetFromAddress())
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdAddPostReaction is the CLI command for adding a like to a post
func GetCmdAddPostReaction(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-reaction [post-id] [value]",
		Short: "Adds a reaction to a post",
		Long: fmt.Sprintf(`
Add a reaction to the post having the given id with the specified value. 
The value has to be a reaction short code.

E.g. 
%s tx posts add-reaction a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc :thumbsup: --from jack
`, version.ClientName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			postID := types.PostID(args[0])
			if !postID.Valid() {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid postID: %s", postID))
			}

			msg := types.NewMsgAddPostReaction(postID, args[1], cliCtx.GetFromAddress())
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdRemovePostReaction is the CLI command for removing a like from a post
func GetCmdRemovePostReaction(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-reaction [post-id] [value]",
		Short: "Removes an existing reaction from a post",
		Long: fmt.Sprintf(`
Removes the reaction having the given value from the post having the given id. 
The value has to be a reaction short code.

E.g. 
%s tx posts remove-reaction a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc :thumbsup: --from jack
`, version.ClientName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			postID := types.PostID(args[0])
			if !postID.Valid() {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid postID: %s", postID))
			}

			msg := types.NewMsgRemovePostReaction(postID, cliCtx.GetFromAddress(), args[1])
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdAnswerPoll is the CLI command for answering a post's poll
func GetCmdAnswerPoll(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "answer-poll [post-id] [answer...]",
		Short: "Answer a post's poll'",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			postID := types.PostID(args[0])
			if !postID.Valid() {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid postID: %s", postID))
			}

			var answers []types.AnswerID
			for i := 1; i < len(args); i++ {
				answer, err := strconv.ParseUint(args[i], 10, 32)
				if err != nil {
					return err
				}

				answers = append(answers, types.AnswerID(answer))
			}

			msg := types.NewMsgAnswerPoll(postID, answers, cliCtx.FromAddress)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdRegisterReaction is the CLI command for registering a reaction
func GetCmdRegisterReaction(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "register-reaction [short-code] [value] [subspace]",
		Short: "Register a new reaction",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, found := types.GetEmojiByShortCodeOrValue(args[1]); found {
				return fmt.Errorf("%s represents an emoji shortcode and thus cannot be used to register another reaction", args[1])
			}

			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgRegisterReaction(cliCtx.FromAddress, args[0], args[1], args[2])
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
