package cli

import (
	"bufio"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetTxCmd set the tx commands
func GetTxCmd(_ string, cdc *codec.Codec) *cobra.Command {
	profileTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Profile transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	profileTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreateProfile(cdc),
		GetCmdEditProfile(cdc),
		GetCmdDeleteProfile(cdc),
	)...)

	return profileTxCmd
}

// GetCmdCreateProfile is the CLI command for creating a profile
func GetCmdCreateProfile(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [moniker]",
		Short: "Create a new profile",
		Long: fmt.Sprintf(`
Create a new profile specifying the moniker, name, surname, bio, a profile picture and cover.
Every data except the moniker is optional, feel free to specify only what you want other people to know publicly about you.

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

			msg := types.NewMsgCreateProfile(name, surname, args[0], bio, &pictures, cliCtx.FromAddress)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagName, "", "Name of the profile")
	cmd.Flags().String(flagSurname, "", "Surname of the profile")
	cmd.Flags().String(flagBio, "", "Biography of the profile")
	cmd.Flags().String(flagProfilePic, "", "Profile related picture")
	cmd.Flags().String(flagProfileCover, "", "Profile related cover picture")

	return cmd
}

// GetCmdEditProfile is the CLI command for editing an profile
func GetCmdEditProfile(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [previous_moniker]",
		Short: "Edit an existent profile",
		Long: fmt.Sprintf(`
Edit an existing profile specifying the previous moniker, new moniker, name, surname, bio, a profile picture and cover.
Every data except moniker is optional.

E.g (with all the other optional fields)
%s tx profile edit leoDiCap \
    --moniker "DiCapLeo" \
	--name "Leo" \
	--surname "Di Cap" \
	--bio "Hollywood actor. Proud environmentalist" \
	--picture "https://profilePic.jpg"
	--cover "https://profileCover.jpg"
`, version.ClientName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			picture := viper.GetString(flagProfilePic)
			cover := viper.GetString(flagProfileCover)

			newMoniker := viper.GetString(flagNewMoniker)
			name := viper.GetString(flagName)
			surname := viper.GetString(flagSurname)
			bio := viper.GetString(flagBio)

			prevMoniker := args[0]
			if newMoniker == "default" {
				newMoniker = prevMoniker
			}

			msg := types.NewMsgEditProfile(prevMoniker, newMoniker, name, surname, bio, picture, cover, cliCtx.FromAddress)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagNewMoniker, "default", "New moniker of the profile")
	cmd.Flags().String(flagName, "default", "Name of the profile")
	cmd.Flags().String(flagSurname, "default", "Surname of the profile")
	cmd.Flags().String(flagBio, "default", "Biography of the profile")
	cmd.Flags().String(flagProfilePic, "default", "Profile related profile picture")
	cmd.Flags().String(flagProfileCover, "default", "Profile related profile cover")

	return cmd
}

// GetCmdDeleteProfile is the CLI command for deleting an profile
func GetCmdDeleteProfile(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an existent profile related to the user's address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			msg := types.NewMsgDeleteProfile(cliCtx.FromAddress)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
