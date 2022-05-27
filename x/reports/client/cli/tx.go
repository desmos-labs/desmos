package cli

// DONTCOVER

import (
	"fmt"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

const (
	flagMessage = "message"
)

// NewTxCmd returns a new command to perform reports transactions
func NewTxCmd() *cobra.Command {
	subspacesTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Reports transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	subspacesTxCmd.AddCommand(
		GetCmdReportUser(),
		GetCmdReportPost(),
		GetCmdDeleteReport(),
		GetCmdSupportStandardReason(),
		GetCmdAddReason(),
		GetCmdRemoveReason(),
	)

	return subspacesTxCmd
}

// GetCmdReportUser returns the command allowing to report a user
func GetCmdReportUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-user [subspace-id] [reason-id] [user-address]",
		Args:  cobra.ExactArgs(3),
		Short: "Report a user, optionally specifying a message",
		Long:  "Report the user inside the specific subspace for the reason having the given id",
		Example: fmt.Sprintf(`
%s tx reports report-user 1 1 desmos1cs0gu6006rz9wnmltjuhnuz8k3a2wg6jzmmgyu \
  --message "Please admins review this report!" \
  --from alice
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			reasonID, err := types.ParseReasonID(args[1])
			if err != nil {
				return err
			}

			userAddr, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}

			message, err := cmd.Flags().GetString(flagMessage)
			if err != nil {
				return err
			}

			reporter := clientCtx.FromAddress.String()

			msg := types.NewMsgCreateReport(
				subspaceID,
				reasonID,
				message,
				types.NewUserData(userAddr.String()),
				reporter,
			)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagMessage, "", "Optional message associated with the report")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdReportPost returns the command allowing to report a user
func GetCmdReportPost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-post [subspace-id] [reason-id] [post-id]",
		Args:  cobra.ExactArgs(3),
		Short: "Report a post, optionally specifying a message",
		Long:  "Report the post having the specified id inside the specific subspace for the reason having the given id",
		Example: fmt.Sprintf(`
%s tx reports report-post 1 1 1 \
  --message "Please admins review this report!" \
  --from alice
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			reasonID, err := types.ParseReasonID(args[1])
			if err != nil {
				return err
			}

			postID, err := poststypes.ParsePostID(args[2])
			if err != nil {
				return err
			}

			message, err := cmd.Flags().GetString(flagMessage)
			if err != nil {
				return err
			}

			reporter := clientCtx.FromAddress.String()

			msg := types.NewMsgCreateReport(
				subspaceID,
				reasonID,
				message,
				types.NewPostData(postID),
				reporter,
			)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagMessage, "", "Optional message associated with the report")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDeleteReport returns the command allowing to delete a report
func GetCmdDeleteReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete-report [subspace-id] [report-id]",
		Args:    cobra.ExactArgs(2),
		Short:   "Delete a report",
		Long:    "Delete the report having the given id from the specified subspace",
		Example: fmt.Sprintf(`%s tx reports delete-report 1 1 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			reportID, err := types.ParseReportID(args[1])
			if err != nil {
				return err
			}

			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgDeleteReport(subspaceID, reportID, signer)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdSupportStandardReason returns the command allowing to support a standard reason
func GetCmdSupportStandardReason() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "support-standard-reason [subspace-id] [report-id]",
		Args:    cobra.ExactArgs(2),
		Short:   "Support a standard reporting reason",
		Long:    "Add the support for the specific standard reporting reason inside the subspaces",
		Example: fmt.Sprintf(`%s tx reports support-standard-reason 1 1 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			reasonID, err := types.ParseReasonID(args[1])
			if err != nil {
				return err
			}

			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgSupportStandardReason(subspaceID, reasonID, signer)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAddReason returns the command allowing to add a new reporting reason
func GetCmdAddReason() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-reason [subspace-id] [title] [[description]]",
		Args:  cobra.RangeArgs(2, 3),
		Short: "Add a new reporting reason",
		Long:  "Add a new reporting reason with the given title and optional desciption to a subspace",
		Example: fmt.Sprintf(`
%s tx reports add-reason "Spam" "Spam content or spammer user" \
  --from alice
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			title := args[2]

			var description string
			if len(args) > 2 {
				description = args[2]
			}

			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgAddReason(subspaceID, title, description, signer)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRemoveReason returns the command allowing to remove a reporting reason
func GetCmdRemoveReason() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove-reason [subspace-id] [reason-id]",
		Args:    cobra.ExactArgs(2),
		Short:   "Remove a reporting reason",
		Long:    "Remove the reporting reason having the given id from the specified subspace",
		Example: fmt.Sprintf(`%s tx reports remove-reason 1 1 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			reasonID, err := types.ParseReasonID(args[1])
			if err != nil {
				return err
			}

			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgRemoveReason(subspaceID, reasonID, signer)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
