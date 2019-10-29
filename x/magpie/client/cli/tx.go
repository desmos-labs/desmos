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
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
)

// GetTxCmd set the tx commands
func GetTxCmd(_ string, cdc *codec.Codec) *cobra.Command {
	magpieTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Magpie transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	magpieTxCmd.AddCommand(client.PostCommands(
		GetCmdCreateSession(cdc),
	)...)

	return magpieTxCmd
}

// GetCmdCreateSession is the CLI command for creating a session for create post
func GetCmdCreateSession(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-session [namespace] [external address] [pubkey] [external signer signature]",
		Short: "Creates a session for an external service to post",
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
