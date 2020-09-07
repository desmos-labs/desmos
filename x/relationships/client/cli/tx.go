package cli

import (
	"bufio"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/relationships/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	posts "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/spf13/cobra"
)

// GetTxCmd set the tx commands
func GetTxCmd(_ string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Profiles transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(flags.PostCommands(
		GetCmdCreateRelationship(cdc),
		GetCmdDeleteRelationship(cdc),
	)...)

	return cmd
}

// GetCmdCreateRelationship is the CLI command for creating a relationship
func GetCmdCreateRelationship(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [receiver] [subspace]",
		Short: "Create a relationship with the given receiver address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			receiver, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			// TODO edit this import to use commons when user blocks is merged
			if !posts.IsValidSubspace(args[1]) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
			}

			msg := types.NewMsgCreateRelationship(cliCtx.FromAddress, receiver, args[1])

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdDeleteRelationship is the CLI command for deleting a relationship
func GetCmdDeleteRelationship(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [receiver] [subspace]",
		Short: "Delete the relationship with the given user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			receiver, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid receiver address: %s", receiver))
			}

			// TODO edit this import to use commons when user blocks is merged
			if !posts.IsValidSubspace(args[1]) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
			}

			msg := types.NewMsgDeleteRelationship(cliCtx.FromAddress, receiver, args[1])
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
