package cli

import (
	"bufio"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		GetCmdCreateAccount(cdc),
		GetCmdEditAccount(cdc),
		GetCmdDeleteAccount(cdc),
	)...)

	return postsTxCmd
}

// GetCmdCreateAccount is the CLI command for creating an account
func GetCmdCreateAccount(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [moniker]",
		Short: "Create a new account",
		Long: fmt.Sprintf(`
Create a new account specifying the moniker, name, surname, bio, a profile picture and cover.
Every data except moniker are optional, let be free to specify only what you want to.

E.g (only with moniker)
%s tx profile create leoDiCap 

E.g (with all the other optional fields)
%s tx profile create leoDiCap \
	--name "Leonardo" \
	--surname "Di Caprio" \
	--bio "Hollywood actor. Proud environmentalist" \
	--picture "https://profilePic.jpg"
	--cover "https://profileCover.jpg"
`, version.ClientName, version.ClientName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			picture := viper.GetString(flagProfilePic)
			cover := viper.GetString(flagProfileCover)
			pictures := types.NewPictures(picture, cover)

			name := viper.GetString(flagName)
			surname := viper.GetString(flagSurname)
			bio := viper.GetString(flagBio)

			msg := types.NewMsgCreateAccount(name, surname, args[0], bio, &pictures, cliCtx.FromAddress)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagName, "", "Name of the account")
	cmd.Flags().String(flagSurname, "", "Surname of the account")
	cmd.Flags().String(flagBio, "", "Biography of the account")
	cmd.Flags().String(flagProfilePic, "", "Account related profile picture")
	cmd.Flags().String(flagProfileCover, "", "Account related profile cover")

	return cmd
}

// GetCmdEditAccount is the CLI command for editing an account
func GetCmdEditAccount(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [moniker]",
		Short: "Edit an existent account",
		Long: fmt.Sprintf(`
Edit an existing account specifying the moniker, name, surname, bio, a profile picture and cover.
Every data except moniker are optional.

E.g (with all the other optional fields)
%s tx profile edit leoDiCap \
	--name "Leo" \
	--surname "Di Cap" \
	--bio "Hollywood actor. Proud environmentalist" \
	--picture "https://profilePic.jpg"
`, version.ClientName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			picture := viper.GetString(flagProfilePic)
			cover := viper.GetString(flagProfileCover)
			pictures := types.NewPictures(picture, cover)

			name := viper.GetString(flagName)
			surname := viper.GetString(flagSurname)
			bio := viper.GetString(flagBio)

			msg := types.NewMsgEditAccount(name, surname, args[0], bio, &pictures, cliCtx.FromAddress)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagName, "", "Name of the account")
	cmd.Flags().String(flagSurname, "", "Surname of the account")
	cmd.Flags().String(flagBio, "", "Biography of the account")
	cmd.Flags().String(flagProfilePic, "", "Account related profile picture")
	cmd.Flags().String(flagProfileCover, "", "Account related profile cover")

	return cmd
}

// GetCmdDeleteAccount is the CLI command for deleting an account
func GetCmdDeleteAccount(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [moniker]",
		Short: "Delete an existent account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			msg := types.NewMsgDeleteAccount(args[0], cliCtx.FromAddress)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
