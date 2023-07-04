package cli

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/cobra"

	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

const (
	FlagOutputPath = "output-path"
)

// DONTCOVER

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Token factory transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdCreateDenom(),
		GetCmdMint(),
		GetCmdBurn(),
		GetCmdSetDenomMetadata(),
		GetCmdDraftDenomMetadata(),
	)

	return txCmd
}

// GetCmdCreateDenom returns the command used to create a denom
func GetCmdCreateDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-denom [subspace-id] [subdenom]",
		Short: fmt.Sprintf("create a new denom from an account. (Costs %s though!)", sdk.DefaultBondDenom),
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			sender := clientCtx.FromAddress
			msg := types.NewMsgCreateDenom(subspaceID, sender.String(), args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdMint returns the command used to mint a denom to an address
func GetCmdMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [subspace-id] [amount]",
		Short: "Mint a denom to an address. Must have permissions to do so.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			sender := clientCtx.FromAddress
			msg := types.NewMsgMint(subspaceID, sender.String(), amount)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdBurn returns the command used to burn a denom from the treasury account
func GetCmdBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [subspace-id] [amount]",
		Short: "Burn tokens from the treasury account. Must have permissions to do so.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			sender := clientCtx.FromAddress
			msg := types.NewMsgBurn(subspaceID, sender.String(), amount)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdSetDenomMetadata returns the command used to set the metadata of the denom
func GetCmdSetDenomMetadata() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-denom-metadata [subspace-id] [json-path]",
		Short: "Set a subspace token metadata. Must have permissions to do so.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			var metadata banktypes.Metadata
			bz, err := os.ReadFile(args[1])
			if err != nil {
				return err
			}

			err = clientCtx.Codec.UnmarshalJSON(bz, &metadata)
			if err != nil {
				return err
			}

			sender := clientCtx.FromAddress
			msg := types.NewMsgSetDenomMetadata(subspaceID, sender.String(), metadata)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDraftDenomMetadata returns the command used to draft a denom metadata
func GetCmdDraftDenomMetadata() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "draft-denom-metadata",
		Short: "Draft a subspace token metadata for setting denom metadata",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			output, err := cmd.Flags().GetString(FlagOutputPath)
			if err != nil {
				return err
			}

			var metadata banktypes.Metadata
			bz, err := clientCtx.Codec.MarshalJSON(&metadata)
			if err != nil {
				return err
			}

			return os.WriteFile(output, bz, 0600)
		},
	}

	cmd.Flags().String(FlagOutputPath, "metadata.json", "output file path of the draft metadata")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
