package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

	postsTxCmd.AddCommand(client.PostCommands(
		GetCmdCreatePost(cdc),
		GetCmdEditPost(cdc),
		GetCmdAddLike(cdc),
		GetCmdRemoveLike(cdc),
	)...)

	return postsTxCmd
}

// GetCmdCreatePost is the CLI command for creating a post
func GetCmdCreatePost(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "create [subspace] [message] [allows-comments] [[[uri],[mime-type]]...] [[EndDate] " +
			"[[answer-id],[text]...] [multiple-answers] [answers-edits]]",
		Short: "Create a new post",
		Long: fmt.Sprintf(`
				Create a new post, specifying the subspace, message and whether or not it will allow for comments.
				Optional media attachments and polls are also supported.
				If you want to add one or more media attachment, you have to specify a uri and a mime type for each.
				Each attachment can be added only once, otherwise and error will occur.
                You can do so by concatenating them together separated by a comma (,).
				Usage examples:

				- tx posts create "desmos" "Hello world!" true
				- tx posts create "desmos" "A post with media" true "https://example.com,text/plain"
				- tx posts create "desmos" "A post with multiple medias" false "https://example.com/media1,text/plain" "https://example.com/media2,application/json"

				If you want to add a poll to your post you have to specifiy:
					1. The end date of your poll after which no further answers will be accepted
					2. A slice of answers that will be provided to the users once they want to take part in poll votations.
                       Each answer should have:
						- an ID in form of uint64 identifying the answer 
					    - the text of the answer itself
					3. A boolean value that indicates the possibility of multiple answers from users
					4. A boolean value that indicates the possibility to edit the answers in future
				Usage examples:
				-  tx posts create "desmos" "A post with a poll" true "2020-01-01T12:00:00Z" "1,answer1" "2,answer2" "3,answer3" false true
		`),
		Args: cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

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

			mediasStrings := strings.Split(viper.GetString(flagMedia), " ")
			medias := types.PostMedias{}
			for _, mediaString := range mediasStrings {
				argz := strings.Split(mediaString, ",")
				if len(argz) != 2 {
					return fmt.Errorf("if medias are specified, they shouldn't have empty fields, please use the --help flag to know more")
				}

				media := types.NewPostMedia(argz[0], argz[1])
				medias = medias.AppendIfMissing(media)
			}

			pollDetailsMap := viper.GetStringMapString(flagPollDetails)
			pollAnswersSlice := viper.GetStringSlice(flagPollAnswer)

			date, err := time.Parse(time.RFC3339, pollDetailsMap["end-date"])
			if err != nil {
				return fmt.Errorf("end date shoulb be provided in RFC3339 format, e.g 2020-01-01T12:00:00Z")
			}

			allowMultipleAnswers, err := strconv.ParseBool(pollDetailsMap["multiple-answers"])
			if err != nil {
				return fmt.Errorf("multiple-answers could be only true or false")
			}

			allowsAnswerEdits, err := strconv.ParseBool(pollDetailsMap["allows-answer-edits"])
			if err != nil {
				return fmt.Errorf("allows-answer-edits could be only true or false")
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

			pollData := types.PollData{
				Open:                  true,
				EndDate:               date,
				ProvidedAnswers:       answers,
				AllowsMultipleAnswers: allowMultipleAnswers,
				AllowsAnswerEdits:     allowsAnswerEdits,
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
	cmd.Flags().StringSlice(flagMedia, nil, "Current post's media")
	cmd.Flags().StringToString(flagPollDetails, nil, "Current post's poll details")
	cmd.Flags().StringSlice(flagPollAnswer, nil, "Current post's poll answer")

	return cmd
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

			msg := types.NewMsgEditPost(postID, args[1], from, time.Now().UTC())
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

			msg := types.NewMsgAddPostReaction(postID, args[1], from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdRemoveLike is the CLI command for removing a like from a post
func GetCmdRemoveLike(cdc *codec.Codec) *cobra.Command {
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

			msg := types.NewMsgRemovePostReaction(postID, from, args[1])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
