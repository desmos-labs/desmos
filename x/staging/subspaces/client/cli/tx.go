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
	"github.com/spf13/viper"
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
		GetCmdRegisterUser(),
		GetCmdUnregisterUser(),
		GetCmdBlockUser(),
		GetCmdUnblockUser(),
		GetCmdEditSubspace(),
	)

	return subspacesTxCmd
}

// GetCmdCreateSubspace returns the command used to create a subspace
func GetCmdCreateSubspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [subspace-id] [name]",
		Args:  cobra.ExactArgs(2),
		Short: "Create a new subspace",
		Long: fmt.Sprintf(`Create a new subspace.
The id must be a valid SHA-256 hash uniquely identifying the subspace.

The name shall be a human readable name, while the --open flag can be used to tell whether 
the subspace allow users to post messages freely. 
e.g 1) %s tx subspaces create 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e "mooncake"
	2) %s tx subspaces create 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e "mooncake" --open
`, version.AppName, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID := args[0]
			subspaceName := args[1]
			open := viper.GetBool(FlagOpen)

			msg := types.NewMsgCreateSubspace(subspaceID, subspaceName, clientCtx.FromAddress.String(), open)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Bool(FlagOpen, false, "Tells if the subspace let post messages freely or not")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdAddSubspaceAdmin returns the command to add an admin to a subspace
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

// GetCmdRemoveSubspaceAdmin returns the command to remove an admin from a subspace
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

// GetCmdRegisterUser returns the command to register a user inside a subspace
func GetCmdRegisterUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-user [address] [subspace-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Register a user inside the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			user := args[0]
			subspaceID := args[1]
			admin := clientCtx.FromAddress.String()
			msg := types.NewMsgRegisterUser(user, subspaceID, admin)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdUnregisterUser returns the command to unregister a user from a subspace
func GetCmdUnregisterUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unregister-user [address] [subspace-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Unregister a user from the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			user := args[0]
			subspaceID := args[1]
			admin := clientCtx.FromAddress.String()
			msg := types.NewMsgUnregisterUser(user, subspaceID, admin)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdBlockUser returns the command to block a user inside a subspace
func GetCmdBlockUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block-user [address] [subspace-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Block a user inside the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			user := args[0]
			subspaceID := args[1]
			admin := clientCtx.FromAddress.String()
			msg := types.NewMsgBlockUser(user, subspaceID, admin)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdUnblockUser returns the command to unblock a user from a subspace
func GetCmdUnblockUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unblock-user [address] [subspace-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Unblock a user inside the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			user := args[0]
			subspaceID := args[1]
			admin := clientCtx.FromAddress.String()
			msg := types.NewMsgBlockUser(user, subspaceID, admin)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdEditSubspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [subspace-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Edit an existent subspace with the given id",
		Long: fmt.Sprintf(`Edit a subspace with the given id.
E.g 
1) Edit the owner only
%s tx subspaces edit 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e --owner "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"
2) Edit the name only
%s tx subspaces edit 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e --name "star"
3) Edit both owner and name
%s tx subspaces edit 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e --name "star" --owner "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"
`, version.AppName, version.AppName, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID := args[0]

			newOwner := viper.GetString(FlagOwner)
			name := viper.GetString(FlagName)

			owner := clientCtx.FromAddress.String()
			msg := types.NewMsgEditSubspace(subspaceID, newOwner, name, owner)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	cmd.Flags().String(FlagName, "", "New human readable name of the subspace")
	cmd.Flags().String(FlagOwner, "", "New owner of the subspace")

	return cmd
}
