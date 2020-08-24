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
	"github.com/desmos-labs/desmos/x/profiles/types"
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
		GetCmdCreateMonoDirectionalRelationship(cdc),
		GetCmdDeleteUserRelationship(cdc),
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
		Use:   "save [dtag]",
		Args:  cobra.ExactArgs(1),
		Short: "Save your profile associating to it the given DTag.",
		Long: fmt.Sprintf(`
Save a new profile or edit the existing one specifying a DTag, a moniker, biography, profile picture and cover picture.
Every data given through the flags is optional.
If you are editing an existing profile you should fill all the existent fields otherwise the existing values
will be removed.

%s tx profiles save LeoDiCap \
	%s "Leonardo Di Caprio" \
	%s "Hollywood actor. Proud environmentalist" \
	%s "https://profilePic.jpg"
	%s "https://profileCover.jpg"
`, version.ClientName, flagMoniker, flagBio, flagProfilePic, flagCoverPic),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			moniker := getFlagValueOrNilOnDefault(flagMoniker)
			picture := getFlagValueOrNilOnDefault(flagProfilePic)
			cover := getFlagValueOrNilOnDefault(flagCoverPic)
			bio := getFlagValueOrNilOnDefault(flagBio)

			msg := types.NewMsgSaveProfile(args[0], moniker, bio, picture, cover, cliCtx.FromAddress)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagMoniker, "", "Moniker to be used")
	cmd.Flags().String(flagBio, "", "Biography to be used")
	cmd.Flags().String(flagProfilePic, "", "Profile picture")
	cmd.Flags().String(flagCoverPic, "", "Cover picture")

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

// GetCmdCreateMonoDirectionalRelationship is the CLI command for creating a monoDirRelationship
func GetCmdCreateMonoDirectionalRelationship(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-relationship [receiver]",
		Short: "Create a mono directional relationship with the given receiver address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			receiver, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMonoDirectionalRelationship(cliCtx.FromAddress, receiver)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdDeleteUserRelationship is the CLI command for deleting a relationship
func GetCmdDeleteUserRelationship(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-relationship [receiver]",
		Short: "Delete the relationship with the given user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			relationshipID, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid receiver address: %s", relationshipID))
			}

			msg := types.NewMsgDeleteRelationship(relationshipID, cliCtx.FromAddress)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
