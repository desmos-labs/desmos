package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// DONTCOVER

const (
	FlagName        = "name"
	FlagDescription = "description"
	FlagTreasury    = "treasury"
	FlagOwner       = "owner"
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
		GetCmdEditSubspace(),
		GetCmdCreateUserGroup(),
		GetCmdDeleteUserGroup(),
		GetCmdAddUserToUserGroup(),
		GetCmdRemoveUserFromUserGroup(),
		GetCmdSetPermissions(),
	)

	return subspacesTxCmd
}

// GetCmdCreateSubspace returns the command used to create a subspace
func GetCmdCreateSubspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [name]",
		Args:  cobra.ExactArgs(1),
		Short: "Create a new subspace",
		Long: `Create a new subspace.
The name must be a human readable name.`,
		Example: fmt.Sprintf(`
%s tx subspaces create "Desmos" \
  --description "The official subspace of Desmos" \
  --treasury desmos1jqk5p244yl4ktukq5xhavvlfzl8z4we4qfmuyh \
  --owner desmos1p8r4guvdze03md4g9zclhh6mr8ljvtd80pehr3 \
  --from alice
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]

			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			treasury, err := cmd.Flags().GetString(FlagTreasury)
			if err != nil {
				return err
			}

			owner, err := cmd.Flags().GetString(FlagOwner)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSubspace(name, description, treasury, owner, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagDescription, "", "Description of the subspace")
	cmd.Flags().String(FlagTreasury, "", "Treasury of the subspace")
	cmd.Flags().String(FlagOwner, "", "Owner of the subspace")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditSubspace returns the command to edit a subspace
func GetCmdEditSubspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [subspace-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Edit the subspace with the given id",
		Example: fmt.Sprintf(`
%s tx subspaces edit 1 \
  --name "Desmos - Democratizing social networks"
  --description "The official subspace of Desmos" \
  --treasury desmos1jqk5p244yl4ktukq5xhavvlfzl8z4we4qfmuyh \
  --owner desmos1p8r4guvdze03md4g9zclhh6mr8ljvtd80pehr3 \
  --from alice
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			name, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			treasury, err := cmd.Flags().GetString(FlagTreasury)
			if err != nil {
				return err
			}

			owner, err := cmd.Flags().GetString(FlagOwner)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditSubspace(subspaceID, name, description, treasury, owner, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagName, types.DoNotModify, "New human readable name of the subspace")
	cmd.Flags().String(FlagDescription, types.DoNotModify, "Description of the subspace")
	cmd.Flags().String(FlagTreasury, types.DoNotModify, "Treasury of the subspace")
	cmd.Flags().String(FlagOwner, types.DoNotModify, "Owner of the subspace")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDeleteSubspace returns the command to delete a subspace
func GetCmdDeleteSubspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [subspace-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Deletes the subspace with the given id",
		Example: fmt.Sprintf(`%s tx subspaces delete 1 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteSubspace(subspaceID, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdCreateUserGroup returns the command to create a user group
func GetCmdCreateUserGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user-group [subspace-id] [group-name] [[permissions]]",
		Args:  cobra.MinimumNArgs(2),
		Short: "Create a new user group within a subspace",
		Long: `Create a new user group within the subspace having the provided id.
The permissions of this group can be set using the third (optional) parameter.
If no permissions are set, the default PermissionNothing will be used instead.
Multiple permissions must be specified separating them with a comma (,).`,
		Example: fmt.Sprintf(`
%s tx subspaces create-user-group 1 "Admins" "Write,ModerateContent,SetPermissions" \
  --from alice
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			name := args[1]

			permission := types.PermissionNothing
			if len(args) > 2 {
				for _, permArg := range strings.Split(args[2], ",") {
					perm, err := types.ParsePermission(permArg)
					if err != nil {
						return err
					}
					permission = types.CombinePermissions(permission, perm)
				}
			}

			msg := types.NewMsgCreateUserGroup(subspaceID, name, permission, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDeleteUserGroup returns the command to delete a user group
func GetCmdDeleteUserGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-user-group [subspace-id] [group-name]",
		Args:  cobra.ExactArgs(2),
		Short: "Delete a user group from a subspace",
		Example: fmt.Sprintf(`
%s tx subspaces delete-user-group 1 "Admins" --from alice
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			name := args[1]

			msg := types.NewMsgDeleteUserGroup(subspaceID, name, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAddUserToUserGroup returns the command to add a user to a user group
func GetCmdAddUserToUserGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-user-to-group [subspace-id] [group-name] [user]",
		Args:  cobra.ExactArgs(3),
		Short: "Add a user to a user group",
		Example: fmt.Sprintf(`
%s tx subspaces add-user-to-user-group 1 "Admins" desmos1p8r4guvdze03md4g9zclhh6mr8ljvtd80pehr3 \
  --from alice
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			name := args[1]
			user := args[2]

			msg := types.NewMsgAddUserToUserGroup(subspaceID, name, user, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRemoveUserFromUserGroup returns the command to remove a user from a user group
func GetCmdRemoveUserFromUserGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-user-from-group [subspace-id] [group-name] [user]",
		Args:  cobra.ExactArgs(3),
		Short: "Remove a user from a user group",
		Example: fmt.Sprintf(`
%s tx subspaces remove-user-from-user-group 1 "Admins" desmos1p8r4guvdze03md4g9zclhh6mr8ljvtd80pehr3 \
  --from alice
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			name := args[1]
			user := args[2]

			msg := types.NewMsgRemoveUserFromUserGroup(subspaceID, name, user, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdSetPermissions returns the command to set the permissions for a target
func GetCmdSetPermissions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-permissions [subspace-id] [target] [permissions]",
		Args:  cobra.MinimumNArgs(3),
		Short: "Set the permissions for a specific target",
		Long: `Set the permissions for a specific target inside a given subspace.
The target can be either a group (in this case the name should be used), 
or a user (in this case the address should be used).

In both cases, it is mandatory to specify at least one permission to be set.
When specifying multiple permissions, they must be separated by a comma (,).`,
		Example: fmt.Sprintf(`
%s tx subspaces set-permissions 1 "Admins" "Write,ModerateContent,SetPermissions" \
  --from alice
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			target := args[1]

			permission := types.PermissionNothing
			for _, arg := range strings.Split(args[2], ",") {
				perm, err := types.ParsePermission(arg)
				if err != nil {
					return err
				}
				permission = types.CombinePermissions(permission, perm)
			}

			msg := types.NewMsgSetPermissions(subspaceID, target, permission, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
