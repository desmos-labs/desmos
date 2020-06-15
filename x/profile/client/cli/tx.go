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

func getFlagValueOrNilOnDefault(flag string) *string {
	flagValue := viper.GetString(flag)
	if flagValue == "" {
		return nil
	}
	return &flagValue
}

// GetCmdSaveProfile is the CLI command for saving an profile
func GetCmdSaveProfile(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save a profile",
		Long: fmt.Sprintf(`
Save a new profile or edit an existing one specifying the dtag, name, surname, bio, a profile picture and cover.
Every data is optional except for the dtag.
If you are editing an existing profile you should fill all the existent fields otherwise they will be set as nil.

%s tx profiles save \
    --dtag "DiCapLeo" \
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

			dtag := viper.GetString(flagDtag)
			picture := getFlagValueOrNilOnDefault(flagProfilePic)
			cover := getFlagValueOrNilOnDefault(flagProfileCover)
			bio := getFlagValueOrNilOnDefault(flagBio)

			msg := types.NewMsgSaveProfile(dtag, bio, picture, cover, cliCtx.FromAddress)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagDtag, "", "DTag of the profile")
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
