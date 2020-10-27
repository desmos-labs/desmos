package cli

import (
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/commons"
	"github.com/desmos-labs/desmos/x/relationships/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetTxCmd set the tx commands
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Profiles transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdCreateRelationship(),
		GetCmdDeleteRelationship(),
		GetCmdBlockUser(),
		GetCmdUnblockUser(),
	)

	return cmd
}

// GetCmdCreateRelationship returns the command allowing to create a relationship
func GetCmdCreateRelationship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [receiver] [subspace]",
		Short: "Create a relationship with the given receiver address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			if !commons.IsValidSubspace(args[1]) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
			}

			msg := types.NewMsgCreateRelationship(clientCtx.FromAddress.String(), args[0], args[1])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

// GetCmdDeleteRelationship returns the command allowing to delete a relationships
func GetCmdDeleteRelationship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [receiver] [subspace]",
		Short: "Delete the relationship with the given user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			if !commons.IsValidSubspace(args[1]) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
			}

			msg := types.NewMsgDeleteRelationship(clientCtx.FromAddress.String(), args[0], args[1])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

// GetCmdBlockUser returns the command allowing to block a user
func GetCmdBlockUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block [address] [subspace] [[reason]]",
		Short: "Block the user with the given address, optionally specifying the reason for the block",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			var reason string
			if len(args) == 3 {
				reason = args[2]
			}

			if !commons.IsValidSubspace(args[1]) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
			}

			msg := types.NewMsgBlockUser(clientCtx.FromAddress.String(), args[0], reason, args[1])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

// GetCmdUnblockUser returns the command allowing to unblock a user
func GetCmdUnblockUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unblock [address] [subspace]",
		Short: "Unblock the user with the given address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			if !commons.IsValidSubspace(args[1]) {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
			}

			msg := types.NewMsgUnblockUser(clientCtx.FromAddress.String(), args[0], args[1])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
