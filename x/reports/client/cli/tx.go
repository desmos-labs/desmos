package cli

import (
	"bufio"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/spf13/cobra"
)

// GetTxCmd set the tx commands
func GetTxCmd(_ string, cdc *codec.Codec) *cobra.Command {
	postsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Reports transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	postsTxCmd.AddCommand(flags.PostCommands(
		GetCmdReportPost(cdc),
	)...)

	return postsTxCmd
}

// GetCmdReportPost is the CLI command for reporting a post
func GetCmdReportPost(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "report [post-id] [reports-type] [reports-message]",
		Short: "reports a post",
		Long: fmt.Sprintf(`
Report an existent post specifying its ID, the reports's type and message.

E.g.
%s tx reports report a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc scam "this post is a scam" 
`, version.ClientName),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			postID := posts.PostID(args[0])
			if !postID.Valid() {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid postID: %s", postID))
			}

			msg := types.NewMsgReportPost(postID, args[1], args[2], cliCtx.GetFromAddress())

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
