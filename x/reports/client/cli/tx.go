package cli

// DONTCOVER

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// NewTxCmd returns a new command to perform reports transactions
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Reports transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdCreateReport(),
		GetCmdDeleteReport(),
		NewReasonsTxCmd(),
	)

	return cmd
}

// GetCmdCreateReport returns the command allowing to create a report
func GetCmdCreateReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [subspace-id] [reasons-ids]",
		Args:  cobra.ExactArgs(2),
		Short: "Create a report for user or a post, optionally specifying a message",
		Long: fmt.Sprintf(`
Report the specified user or post inside the specific subspace for the reasons having the given ids.
Multiple reasons can be specified. If so, each reason id must be separated using a comma.

To report a user, --%s must be used. 
To report a post, --%s must be used instead.`, FlagUser, FlagPostID),
		Example: fmt.Sprintf(`
%[1]s tx reports report 1 1,2,3 \
  --%s desmos1cs0gu6006rz9wnmltjuhnuz8k3a2wg6jzmmgyu \
  --message "This user is spammer" \
  --from alice

%[1]s tx reports report 1 1,2,3 \
  --%s 1 \
  --message "This port is spam" \
  --from alice
`, version.AppName, FlagUser, FlagPostID),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			reasons, err := types.ParseReasonsIDs(args[1])
			if err != nil {
				return err
			}

			target, err := ReadReportTarget(cmd.Flags())
			if err != nil {
				return err
			}

			if target == nil {
				return fmt.Errorf("at least one of --%s or --%s must be specfieid", FlagUser, FlagPostID)
			}

			message, err := cmd.Flags().GetString(FlagMessage)
			if err != nil {
				return err
			}

			reporter := clientCtx.FromAddress.String()

			msg := types.NewMsgCreateReport(subspaceID, reasons, message, target, reporter)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagUser, "", "Address of the user to be reported")
	cmd.Flags().Uint64(FlagPostID, 0, "Id the post to be reported")
	cmd.Flags().String(FlagMessage, "", "Optional message associated with the report")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDeleteReport returns the command allowing to delete a report
func GetCmdDeleteReport() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [subspace-id] [report-id]",
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

// --------------------------------------------------------------------------------------------------------------------

// NewReasonsTxCmd returns a new command to perform reasons transactions
func NewReasonsTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "reactions",
		Short:                      "Reports reactions transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdSupportStandardReason(),
		GetCmdAddReason(),
		GetCmdRemoveReason(),
	)

	return cmd
}

// GetCmdSupportStandardReason returns the command allowing to support a standard reason
func GetCmdSupportStandardReason() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "support-standard [subspace-id] [reason-id]",
		Args:    cobra.ExactArgs(2),
		Short:   "Support a standard reporting reason",
		Long:    "Add the support for the specific standard reporting reason inside the subspace",
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
		Use:   "add [subspace-id] [title] [[description]]",
		Args:  cobra.RangeArgs(2, 3),
		Short: "Add a new reporting reason",
		Long:  "Add a new reporting reason with the given title and optional description to a subspace",
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

			title := args[1]

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
		Use:     "remove [subspace-id] [reason-id]",
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
