package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// NewTxCmd returns a new command allowing to perform profiles transactions
func NewTxCmd() *cobra.Command {
	profileTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Profiles transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	profileTxCmd.AddCommand(
		GetCmdSaveProfile(),
		GetCmdDeleteProfile(),
		GetCmdRequestDTagTransfer(),
		GetCmdAcceptDTagTransfer(),
		GetCmdRefuseDTagTransfer(),
		GetCmdCancelDTagTransfer(),
	)

	return profileTxCmd
}

// GetCmdSaveProfile returns the command used to save a profile
func GetCmdSaveProfile() *cobra.Command {
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
`, version.AppName, flagMoniker, flagBio, flagProfilePic, flagCoverPic),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			dTag := args[0]
			moniker := viper.GetString(flagMoniker)
			bio := viper.GetString(flagBio)
			profilePic := viper.GetString(flagProfilePic)
			coverPic := viper.GetString(flagCoverPic)

			msg := types.NewMsgSaveProfile(dTag, moniker, bio, profilePic, coverPic, clientCtx.FromAddress.String())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagMoniker, types.DoNotModify, "Moniker to be used")
	cmd.Flags().String(flagBio, types.DoNotModify, "Biography to be used")
	cmd.Flags().String(flagProfilePic, types.DoNotModify, "Profile picture")
	cmd.Flags().String(flagCoverPic, types.DoNotModify, "Cover picture")

	return cmd
}

// GetCmdDeleteProfile returns the command used to delete an existing profile
func GetCmdDeleteProfile() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete an existent profile related to the user's address",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteProfile(clientCtx.FromAddress.String())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// GetCmdRequestDTagTransfer returns the command to create a DTag transfer request
func GetCmdRequestDTagTransfer() *cobra.Command {
	return &cobra.Command{
		Use:   "transfer-dtag [address]",
		Short: "Make a request to get the DTag of the user having the given address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			requestRecipient, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestDTagTransfer(requestRecipient.String(), clientCtx.FromAddress.String())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// GetCmdAcceptDTagTransfer returns the command to accept a DTag transfer request
func GetCmdAcceptDTagTransfer() *cobra.Command {
	return &cobra.Command{
		Use:   "accept-dtag-transfer [newDTag] [address]",
		Short: "Accept a DTag transfer request made by the user with the given address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			receivingUser, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAcceptDTagTransfer(args[0], clientCtx.FromAddress.String(), receivingUser.String())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// GetCmdRefuseDTagTransfer returns the command to refuse an incoming DTag transfer request
func GetCmdRefuseDTagTransfer() *cobra.Command {
	return &cobra.Command{
		Use:   "refuse-dtag-transfer [sender]",
		Short: "Refuse a DTag transfer made by the given sender address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			sender, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRefuseDTagTransferRequest(sender.String(), clientCtx.FromAddress.String())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// GetCmdCancelDTagTransfer returns the command to cancel an outgoing DTag transfer request
func GetCmdCancelDTagTransfer() *cobra.Command {
	return &cobra.Command{
		Use:   "cancel-dtag-transfer [owner]",
		Short: "Cancel a DTag transfer made to the given owner address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			owner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelDTagTransferRequest(clientCtx.FromAddress.String(), owner.String())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}
