package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

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
		GetCmdCreateRelationship(),
		GetCmdDeleteRelationship(),
		GetCmdBlockUser(),
		GetCmdUnblockUser(),
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
Save a new profile or edit the existing one specifying a DTag, a nickname, biography, profile picture and cover picture.
Every data given through the flags is optional.
If you are editing an existing profile you should fill only the fields that you want to edit. 
The empty ones will be filled with a special [do-not-modify] flag that tells the system to not edit them.

%s tx profiles save LeoDiCap \
	%s "Leonardo Di Caprio" \
	%s "Hollywood actor. Proud environmentalist" \
	%s "https://profilePic.jpg"
	%s "https://profileCover.jpg"
`, version.AppName, FlagNickname, FlagBio, FlagProfilePic, FlagCoverPic),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			dTag := args[0]
			nickname, _ := cmd.Flags().GetString(FlagNickname)
			bio, _ := cmd.Flags().GetString(FlagBio)
			profilePic, _ := cmd.Flags().GetString(FlagProfilePic)
			coverPic, _ := cmd.Flags().GetString(FlagCoverPic)

			msg := types.NewMsgSaveProfile(dTag, nickname, bio, profilePic, coverPic, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagNickname, types.DoNotModify, "Nickname to be used")
	cmd.Flags().String(FlagBio, types.DoNotModify, "Biography to be used")
	cmd.Flags().String(FlagProfilePic, types.DoNotModify, "Profile picture")
	cmd.Flags().String(FlagCoverPic, types.DoNotModify, "Cover picture")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDeleteProfile returns the command used to delete an existing profile
func GetCmdDeleteProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an existent profile related to the user's address",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteProfile(clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRequestDTagTransfer returns the command to create a DTag transfer request
func GetCmdRequestDTagTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-dtag [address]",
		Short: "Make a request to get the DTag of the user having the given address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			requestRecipient, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestDTagTransfer(clientCtx.FromAddress.String(), requestRecipient.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdCancelDTagTransfer returns the command to cancel an outgoing DTag transfer request
func GetCmdCancelDTagTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-dtag-transfer [recipient]",
		Short: "Cancel a DTag transfer made to the given recipient address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelDTagTransferRequest(clientCtx.FromAddress.String(), owner.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAcceptDTagTransfer returns the command to accept a DTag transfer request
func GetCmdAcceptDTagTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept-dtag-transfer [newDTag] [address]",
		Short: "Accept a DTag transfer request made by the user with the given address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			receivingUser, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAcceptDTagTransfer(args[0], receivingUser.String(), clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRefuseDTagTransfer returns the command to refuse an incoming DTag transfer request
func GetCmdRefuseDTagTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refuse-dtag-transfer [sender]",
		Short: "Refuse a DTag transfer made by the given sender address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRefuseDTagTransferRequest(args[0], clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdCreateRelationship returns the command allowing to create a relationship
func GetCmdCreateRelationship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-relationship [receiver] [subspace]",
		Short: "Create a relationship with the given receiver address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRelationship(clientCtx.FromAddress.String(), args[0], args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDeleteRelationship returns the command allowing to delete a relationships
func GetCmdDeleteRelationship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-relationship [receiver] [subspace]",
		Short: "Delete the relationship with the given user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRelationship(clientCtx.FromAddress.String(), args[0], args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdBlockUser returns the command allowing to block a user
func GetCmdBlockUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block [address] [subspace] [[reason]]",
		Short: "Block the user with the given address, optionally specifying the reason for the block",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			reason := ""
			if len(args) == 3 {
				reason = args[2]
			}

			msg := types.NewMsgBlockUser(clientCtx.FromAddress.String(), args[0], reason, args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdUnblockUser returns the command allowing to unblock a user
func GetCmdUnblockUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unblock [address] [subspace]",
		Short: "Unblock the user with the given address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUnblockUser(clientCtx.FromAddress.String(), args[0], args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
