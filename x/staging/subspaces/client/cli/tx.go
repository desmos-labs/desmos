package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/spf13/cobra"
)

// NewTxCmd returns a new command to perform subspaces transactions
func NewTxCmd() *cobra.Command {
	subspacesTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Subspaces transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	subspacesTxCmd.AddCommand(
		GetCmdCreateSubspace(),
		GetCmdAddSubspaceAdmin(),
		GetCmdRemoveSubspaceAdmin(),
		GetCmdEnablePostsForUser(),
		GetCmdDisablePostsForUser(),
		GetCmdTransferOwnership(),
	)

	return subspacesTxCmd
}

// GetCmdCreateSubspace returns the command used to create a subspace
func GetCmdCreateSubspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [subspace-id] [name]",
		Args:  cobra.ExactArgs(2),
		Short: "Create a subspace with the given [subspace-id] and [name]",
		Long: fmt.Sprintf(`Create a new subspace with the given [subspace-id] and name. 
The given id must be a sha256 string identifying the subspace 
%s tx subspaces create 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID := args[0]
			subspaceName := args[1]

			msg := types.NewMsgCreateSubspace(subspaceID, subspaceName, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetCmdAddSubspaceAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-admin [address] [subspace-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Add a new admin to the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			newAdminAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			subspaceID := args[1]

			msg := types.NewMsgAddAdmin(subspaceID, newAdminAddress.String(), clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdRemoveSubspaceAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-admin [address] [subspace-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Remove an existent admin from the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			existentAdminAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			subspaceID := args[1]

			msg := types.NewMsgRemoveAdmin(subspaceID, existentAdminAddress.String(), clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdEnablePostsForUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable-posts [address] [subspace-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Enable the possibility to post inside the subspace with the given [subspace-id] for the user with the given [address]",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			user := args[0]
			subspaceID := args[1]
			admin := clientCtx.FromAddress.String()
			msg := types.NewMsgEnableUserPosts(user, subspaceID, admin)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdDisablePostsForUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable-posts [address] [subspace-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Disable the possibility to post inside the subspace with the given [subspace-id] for the user with the given [address]",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			user := args[0]
			subspaceID := args[1]
			admin := clientCtx.FromAddress.String()
			msg := types.NewMsgDisableUserPosts(user, subspaceID, admin)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdTransferOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-ownership [address] [subspace-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Transfer the ownership of the subspace with the given [subspace-id] to the user with the given [address]",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			newOwner := args[0]
			subspaceID := args[1]
			owner := clientCtx.FromAddress.String()
			msg := types.NewMsgTransferOwnership(newOwner, subspaceID, owner)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
