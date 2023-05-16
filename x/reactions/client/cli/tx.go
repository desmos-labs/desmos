package cli

// DONTCOVER

import (
	"fmt"

	cliutils "github.com/desmos-labs/desmos/v5/x/reactions/client/utils"

	poststypes "github.com/desmos-labs/desmos/v5/x/posts/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v5/x/reactions/types"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

const (
	FlagRegisteredReaction = "registered-reaction"
	FlagFreeTextReaction   = "free-text"
	FlagShorthandCode      = "shorthand-code"
	FlagDisplayValue       = "display-value"
)

// NewTxCmd returns a new command to perform reactions transactions
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Reactions transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdAddReaction(),
		GetCmdRemoveReaction(),
		GetRegisteredReactionsTxCmd(),
		GetCmdSetParams(),
	)

	return txCmd
}

// GetCmdAddReaction returns the command allowing to add a reaction to a post
func GetCmdAddReaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [subspace-id] [post-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Add a reaction to a post",
		Long: fmt.Sprintf(`
Add a reaction to the post with the given id inside the specified subspace.
In order to specify the reaction value, either --%s or --%s must be used`, FlagRegisteredReaction, FlagFreeTextReaction),
		Example: fmt.Sprintf(`
%[1]s tx reactions add 1 2 --%[2]s 1 --from alice
%[1]s tx reactions add 1 2 --%[3]s "🚀" --from alice
`, version.AppName, FlagRegisteredReaction, FlagFreeTextReaction),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := poststypes.ParsePostID(args[1])
			if err != nil {
				return err
			}

			registeredReactionValue, err := cmd.Flags().GetUint32(FlagRegisteredReaction)
			if err != nil {
				return err
			}

			freeTextValue, err := cmd.Flags().GetString(FlagFreeTextReaction)
			if err != nil {
				return err
			}

			var value types.ReactionValue
			switch {
			case registeredReactionValue != 0 && freeTextValue != "":
				return fmt.Errorf("please use only one of either --%s or --%s", FlagRegisteredReaction, FlagFreeTextReaction)

			case registeredReactionValue != 0:
				value = types.NewRegisteredReactionValue(registeredReactionValue)

			case freeTextValue != "":
				value = types.NewFreeTextValue(freeTextValue)

			default:
				return fmt.Errorf("at least one of either --%s or --%s must be specfied", FlagRegisteredReaction, FlagFreeTextReaction)
			}

			user := clientCtx.FromAddress.String()

			msg := types.NewMsgAddReaction(subspaceID, postID, value, user)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint32(FlagRegisteredReaction, 0, "Registered reaction id with which to react")
	cmd.Flags().String(FlagFreeTextReaction, "", "Free text value with which to react")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRemoveReaction returns the command allowing to remove a reaction from a post
func GetCmdRemoveReaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove [subspace-id] [post-id] [reaction-id]",
		Args:    cobra.ExactArgs(3),
		Short:   "Remove the reaction with the given id from the provided post",
		Example: fmt.Sprintf(`%s tx reactions remove 1 1 1 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			postID, err := poststypes.ParsePostID(args[1])
			if err != nil {
				return err
			}

			reactionID, err := types.ParseReactionID(args[2])
			if err != nil {
				return err
			}

			user := clientCtx.FromAddress.String()

			msg := types.NewMsgRemoveReaction(subspaceID, postID, reactionID, user)
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

// GetRegisteredReactionsTxCmd returns a new command to perform registered reactions transactions
func GetRegisteredReactionsTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "registered",
		Short:                      "Registered reactions transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdAddRegisteredReaction(),
		GetCmdEditRegisteredReaction(),
		GetCmdRemoveRegisteredReaction(),
	)

	return txCmd
}

// GetCmdAddRegisteredReaction returns the command allowing to add a registered reaction
func GetCmdAddRegisteredReaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add [subspace-id] [shorthand-code] [display-value]",
		Args:    cobra.ExactArgs(3),
		Short:   "Register a new reaction",
		Long:    "Register a new reaction with the specified shorthand code and display value inside the given subspace",
		Example: fmt.Sprintf(`%s tx reactions registered add 1 ":hello:" "https://example.com?image=hello.png" --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			shorthandCode := args[1]
			displayValue := args[2]
			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgAddRegisteredReaction(subspaceID, shorthandCode, displayValue, signer)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditRegisteredReaction returns the command allowing to edit a registered reaction
func GetCmdEditRegisteredReaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [subspace-id] [reaction-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Edit an existing registered reaction",
		Example: fmt.Sprintf(`
%s tx reactions registered edit 1 1 \
  --shorthand-code ":wave:" \
  --display-value "https://example.com?image=wave.png" \
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

			reactionID, err := types.ParseReactionID(args[1])
			if err != nil {
				return err
			}

			shorthandCode, err := cmd.Flags().GetString(FlagShorthandCode)
			if err != nil {
				return err
			}

			displayValue, err := cmd.Flags().GetString(FlagDisplayValue)
			if err != nil {
				return err
			}

			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgEditRegisteredReaction(subspaceID, reactionID, shorthandCode, displayValue, signer)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagShorthandCode, types.DoNotModify, "New shorthand code to be set")
	cmd.Flags().String(FlagDisplayValue, types.DoNotModify, "New display value to be set")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRemoveRegisteredReaction returns the command allowing to remove a registered reaction
func GetCmdRemoveRegisteredReaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove [subspace-id] [reaction-id]",
		Args:    cobra.ExactArgs(2),
		Short:   "Remove a registered reaction from a subspace",
		Example: fmt.Sprintf(`%s tx reactions registered remove 1 1 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			reactionID, err := types.ParseRegisteredReactionID(args[1])
			if err != nil {
				return err
			}

			signer := clientCtx.FromAddress.String()

			msg := types.NewMsgRemoveRegisteredReaction(subspaceID, reactionID, signer)
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

// GetCmdSetParams returns the command allowing to set reactions parameters
func GetCmdSetParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-params [subspace-id] [json-file-path]",
		Args:  cobra.ExactArgs(2),
		Short: "Set the reactions params of a subspace",
		Long:  "Set the reactions params for the given subspace, reading it from the JSON file located at the given path",
		Example: fmt.Sprintf(`
%s tx reactions set-params 1 /path/to/file.json \
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

			data, err := cliutils.ParseSetReactionsParamsJSON(clientCtx.Codec, args[1])
			if err != nil {
				return err
			}

			user := clientCtx.FromAddress.String()

			msg := types.NewMsgSetReactionsParams(subspaceID, data.RegisteredReactionParams, data.FreeTextParams, user)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
