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
		Short:                      "Profiles transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	profileTxCmd.AddCommand(flags.PostCommands(
		GetCmdSaveProfile(cdc),
		GetCmdDeleteProfile(cdc),
	)...)

	return profileTxCmd
}

// GetCmdSaveProfile is the CLI command for saving an profile
func GetCmdSaveProfile(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save a profile",
		Long: fmt.Sprintf(`
Save a new profile or edit an existing one specifying the moniker, name, surname, bio, a profile picture and cover.
Every data is optional except for the moniker.
If you are editing an existing profile you should fill all the existent fields otherwise they will be set as nil.

%s tx profiles save \
    --moniker "DiCapLeo" \
	--name "Leonardo" \
	--surname "Di Caprio" \
	--bio "Hollywood actor. Proud environmentalist" \
	--picture "https://profilePic.jpg"
	--cover "https://profileCover.jpg"
`, version.ClientName),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			picture := viper.GetString(flagProfilePic)
			cover := viper.GetString(flagProfileCover)
			moniker := viper.GetString(flagMoniker)
			name := viper.GetString(flagName)
			surname := viper.GetString(flagSurname)
			bio := viper.GetString(flagBio)

			msg := types.NewMsgSaveProfile(moniker, &name, &surname, &bio, &picture, &cover, cliCtx.FromAddress)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagMoniker, "", "Moniker of the profile")
	cmd.Flags().String(flagName, "", "Name of the profile")
	cmd.Flags().String(flagSurname, "", "Surname of the profile")
	cmd.Flags().String(flagBio, "", "Biography of the profile")
	cmd.Flags().String(flagProfilePic, "", "Profile related profile picture")
	cmd.Flags().String(flagProfileCover, "", "Profile related profile cover")

	return cmd
}

// GetCmdDeleteProfile is the CLI command for deleting an profile
func GetCmdDeleteProfile(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an existent profile related to the user's address",
		Args:  cobra.NoArgs,
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
