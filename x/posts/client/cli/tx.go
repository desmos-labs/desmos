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

	postsTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreatePost(cdc),
		GetCmdEditPost(cdc),
		GetCmdAddPostReaction(cdc),
		GetCmdRemovePostReaction(cdc),
		GetCmdAnswerPoll(cdc),
	)...)

	return postsTxCmd
}

// GetCmdCreatePost is the CLI command for creating a post
func GetCmdCreatePost(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [subspace] [message] [allows-comments]",
		Short: "Create a new post",
		Long: fmt.Sprintf(`
				Create a new post, specifying the subspace, message and whether or not it will allow for comments.
				Optional media attachments and polls are also supported.
				If you want to add one or more medias attachments, you have to use the --media flag.
				You need to specify both media's URI and mime-type in this order separeted by a comma
				Usage examples:

				- tx posts create "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e" "Hello world!" true
				- tx posts create "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e" "A post with multiple medias" false \
                  --media "https://example.com/media1,text/plain"
				- tx posts create "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e" "A post with multiple medias" false \
                  --media "https://example.com/media1,text/plain https://example.com/media2,application/json"
				
				If you want to add a poll to your post you need to specify it through two flags:
					1. --poll-details 
                       where you specify a map of the following fields:
						* The question of the poll (key: question)
						* The end date of your poll after which no further answers will be accepted (key: date)
						* A boolean value that indicates the possibility of multiple answers from users (key: multiple-answers)
						* A boolean value that indicates the possibility to edit the answers in future (key: allows-answer-edits)
					2. --poll-answer 
                       where you specify a slice of answers that will be provided to the users once they want to take part in poll votations.
						Each answer should is identified by the text of the answer itself
				Usage examples:
				- tx posts create "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e" "Post" true \
  					--poll-details question="Which dog do you prefer?",multiple-answers=false,allow-answer-edits=true,date=2020-01-01T15:00:00.000Z \
  					--poll-answer "Beagle" \
  					--poll-answer "Carlino" \
					--poll-answer "German Sheperd,Labrador"
		`),
		Args: cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)
			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			allowsComments, err := strconv.ParseBool(args[2])
			if err != nil {
				return err
			}

			parentID, err := types.ParsePostID(viper.GetString(flagParentID))
			if err != nil {
				return err
			}

			// medias' checks

			mediasStrings := viper.GetStringSlice(flagMedia)
			medias := types.PostMedias{}
			for _, mediaString := range mediasStrings {
				argz := strings.Split(mediaString, ",")
				if len(argz) != 2 {
					return fmt.Errorf("if medias are specified, the arguments has to be exactly 2 and in this order: \"URI,Mime-Type\", please use the --help flag to know more")
				}

				media := types.NewPostMedia(argz[0], argz[1])
				medias = medias.AppendIfMissing(media)
			}

			// polls' checks

			pollDetailsMap := viper.GetStringMapString(flagPollDetails)
			pollAnswersSlice := viper.GetStringSlice(flagPollAnswer)

			if len(pollDetailsMap) == 0 && len(pollAnswersSlice) > 0 {
				return fmt.Errorf("poll answers specified but no poll details found. Please use %s to specify the poll details", flagPollDetails)
			}

			if len(pollDetailsMap) > 0 && len(pollAnswersSlice) == 0 {
				return fmt.Errorf("poll details specified but answers are not. Please use the %s to specify one or more answer", flagPollAnswer)
			}

			pollData := types.PollData{}
			if len(pollDetailsMap) > 0 && len(pollAnswersSlice) > 0 {
				date, err := time.Parse(time.RFC3339, pollDetailsMap[keyEndDate])
				if err != nil {
					return fmt.Errorf("end date should be provided in RFC3339 format, e.g 2020-01-01T12:00:00Z")
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
						ID:   uint64(index),
						Text: answer,
					}

					answers = answers.AppendIfMissing(pollAnswer)
				}

				pollData = types.PollData{
					Question:              question,
					Open:                  true,
					EndDate:               date,
					ProvidedAnswers:       answers,
					AllowsMultipleAnswers: allowMultipleAnswers,
					AllowsAnswerEdits:     allowsAnswerEdits,
				}
			}

			msg := types.NewMsgCreatePost(
				args[1],
				parentID,
				allowsComments,
				args[0],
				map[string]string{},
				from,
				time.Now().UTC(),
				medias,
				&pollData,
			)

			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagParentID, "0", "Id of the post to which this one should be an answer to")
	cmd.Flags().StringArray(flagMedia, []string{}, "Current post's media")
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
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)
			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[0])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
			}

			msg := types.NewMsgEditPost(postID, args[1], from, time.Now().UTC())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

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
The value can be anything as long as it is ASCII supported.

E.g. 
%s tx posts add-reaction 12 like --from jack
%s tx posts add-reaction 12 👍 --from jack
`, version.ClientName, version.ClientName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)
			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[0])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
			}

			msg := types.NewMsgAddPostReaction(postID, args[1], from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

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
The value can be anything as long as it is ASCII supported.

E.g. 
%s tx posts remove-reaction 12 like --from jack
%s tx posts remove-reaction 12 👍 --from jack
`, version.ClientName, version.ClientName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)
			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[0])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, err.Error())
			}

			msg := types.NewMsgRemovePostReaction(postID, from, args[1])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

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

			accGetter := authtypes.NewAccountRetriever(cliCtx)
			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			postID, err := types.ParsePostID(args[0])
			if err != nil {
				return err
			}

			var answers []uint64
			for i := 1; i < len(args); i++ {
				answer, err := strconv.ParseUint(args[i], 10, 64)
				if err != nil {
					return err
				}

				answers = append(answers, answer)
			}

			msg := types.NewMsgAnswerPollPost(postID, answers, from)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
