package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	"github.com/desmos-labs/desmos/v3/x/relationships/types"
)

// NewTxCmd returns a new command allowing to perform profiles transactions
func NewTxCmd() *cobra.Command {
	profileTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Relationships transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	profileTxCmd.AddCommand(
		GetCmdCreateRelationship(),
		GetCmdDeleteRelationship(),
		GetCmdBlockUser(),
		GetCmdUnblockUser(),
	)

	return profileTxCmd
}

// GetCmdCreateRelationship returns the command allowing to create a relationship
func GetCmdCreateRelationship() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-relationship [counterparty] [subspace-id]",
		Short:   "Create a relationship with the given receiver address",
		Example: fmt.Sprintf(`%s tx relationships create-relationship desmos1e209r8nc8qdkmqujahwrq4xrlxhk3fs9k7yzmw 1`, version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRelationship(clientCtx.FromAddress.String(), args[0], subspaceID)
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
		Use:     "delete-relationship [counterparty] [subspace-id]",
		Short:   "Delete the relationship with the given user",
		Example: fmt.Sprintf(`%s tx relationships delete-relationship desmos1e209r8nc8qdkmqujahwrq4xrlxhk3fs9k7yzmw 1`, version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRelationship(clientCtx.FromAddress.String(), args[0], subspaceID)
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
		Example: fmt.Sprintf(`%s tx relationships block desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud 1
%s tx relationships block desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud 1 "reason for block"`, version.AppName, version.AppName),
		Args: cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[1])
			if err != nil {
				return err
			}

			reason := ""
			if len(args) == 3 {
				reason = args[2]
			}

			msg := types.NewMsgBlockUser(clientCtx.FromAddress.String(), args[0], reason, subspaceID)
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
		Use:     "unblock [address] [subspace]",
		Short:   "Unblock the user with the given address",
		Example: fmt.Sprintf(`%s tx relationships unblock desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud 1`, version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := subspacestypes.ParseSubspaceID(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgUnblockUser(clientCtx.FromAddress.String(), args[0], subspaceID)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
