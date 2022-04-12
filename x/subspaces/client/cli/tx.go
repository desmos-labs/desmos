package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// DONTCOVER

const (
	FlagName        = "name"
	FlagDescription = "description"
	FlagTreasury    = "treasury"
	FlagOwner       = "owner"
	FlagPermissions = "permissions"
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
		GetCmdDeleteSubspace(),

		NewGroupsTxCmd(),

		GetCmdSetUserPermissions(),
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

// -------------------------------------------------------------------------------------------------------------------

// NewGroupsTxCmd returns a new command to perform subspaces groups transactions
func NewGroupsTxCmd() *cobra.Command {
	groupsTxCmd := &cobra.Command{
		Use:                        "groups",
		Short:                      "Subspace groups transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	groupsTxCmd.AddCommand(
		GetCmdCreateUserGroup(),
		GetCmdEditUserGroup(),
		GetCmdSetUserGroupPermissions(),
		GetCmdDeleteUserGroup(),
		GetCmdAddUserToUserGroup(),
		GetCmdRemoveUserFromUserGroup(),
	)

	return groupsTxCmd
}

// GetCmdCreateUserGroup returns the command to create a user group
func GetCmdCreateUserGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [subspace-id] [group-name]",
		Args:  cobra.MinimumNArgs(2),
		Short: "Create a new user group within a subspace",
		Long: fmt.Sprintf(`Create a new user group within the subspace having the provided id.

An optional description of this group can be provided with the %[1]s flag.

The permissions of this group can be set using the %[2]s flag.
If no permissions are set, the default PermissionNothing will be used instead.
Multiple permissions must be specified separating them with a comma (,).`, FlagDescription, FlagPermissions),
		Example: fmt.Sprintf(`
%s tx subspaces groups create 1 "Admins" \
  --description "Group of the subspace admins" \
  --permissions "Write,ModerateContent,SetUserPermissions" \
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

			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			permissions, err := cmd.Flags().GetStringSlice(FlagPermissions)
			if err != nil {
				return err
			}

			permission := types.PermissionNothing
			for _, permArg := range permissions {
				perm, err := types.ParsePermission(permArg)
				if err != nil {
					return err
				}
				permission = types.CombinePermissions(permission, perm)
			}

			msg := types.NewMsgCreateUserGroup(subspaceID, name, description, permission, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagDescription, "", "Description of the group")
	cmd.Flags().StringSlice(FlagPermissions, []string{types.SerializePermission(types.PermissionNothing)}, "Permissions of the group")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditUserGroup returns the command to edit a user group
func GetCmdEditUserGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [subspace-id] [group-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Edit the group with the given id",
		Example: fmt.Sprintf(`
%s tx subspaces groups edit 1 1 \
  --name "Super admins"
  --description "This is the group of super users" \
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

			groupID, err := types.ParseGroupID(args[1])
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

			msg := types.NewMsgEditUserGroup(subspaceID, groupID, name, description, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagName, types.DoNotModify, "New human readable name of the group")
	cmd.Flags().String(FlagDescription, types.DoNotModify, "Description of the group")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdSetUserGroupPermissions returns the command to set the permissions for a user group
func GetCmdSetUserGroupPermissions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-permissions [subspace-id] [group-id] [permissions]",
		Args:  cobra.ExactArgs(3),
		Short: "Set the permissions for a specific group",
		Long: `Set the permissions for a specific user group a given subspace.
It is mandatory to specify at least one permission to be set.
When specifying multiple permissions, they must be separated by a comma (,).`,
		Example: fmt.Sprintf(`
%s tx subspaces groups set-permissions 1 1 "Write,ModerateContent,SetUserPermissions" \
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

			groupID, err := types.ParseGroupID(args[1])
			if err != nil {
				return err
			}

			permission := types.PermissionNothing
			for _, arg := range strings.Split(args[2], ",") {
				perm, err := types.ParsePermission(arg)
				if err != nil {
					return err
				}
				permission = types.CombinePermissions(permission, perm)
			}

			msg := types.NewMsgSetUserGroupPermissions(subspaceID, groupID, permission, clientCtx.FromAddress.String())
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
		Use:   "delete [subspace-id] [group-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Delete a user group from a subspace",
		Example: fmt.Sprintf(`
%s tx subspaces groups delete 1 1 --from alice
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

			groupID, err := types.ParseGroupID(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteUserGroup(subspaceID, groupID, clientCtx.FromAddress.String())
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
		Use:   "add-user [subspace-id] [group-id] [user]",
		Args:  cobra.ExactArgs(3),
		Short: "Add a user to a user group",
		Example: fmt.Sprintf(`
%s tx subspaces groups add-user 1 1 desmos1p8r4guvdze03md4g9zclhh6mr8ljvtd80pehr3 \
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

			groupID, err := types.ParseGroupID(args[1])
			if err != nil {
				return err
			}

			user := args[2]

			msg := types.NewMsgAddUserToUserGroup(subspaceID, groupID, user, clientCtx.FromAddress.String())
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
		Use:   "remove-user [subspace-id] [group-id] [user]",
		Args:  cobra.ExactArgs(3),
		Short: "Remove a user from a user group",
		Example: fmt.Sprintf(`
%s tx subspaces groups remove-user 1 1 desmos1p8r4guvdze03md4g9zclhh6mr8ljvtd80pehr3 \
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

			groupID, err := types.ParseGroupID(args[1])
			if err != nil {
				return err
			}

			user := args[2]

			msg := types.NewMsgRemoveUserFromUserGroup(subspaceID, groupID, user, clientCtx.FromAddress.String())
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

// GetCmdSetUserPermissions returns the command to set the permissions for a user
func GetCmdSetUserPermissions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-user-permissions [subspace-id] [user] [permissions]",
		Args:  cobra.ExactArgs(3),
		Short: "Set the permissions for a specific user",
		Long: `Set the permissions for a specific user inside a given subspace.
It is mandatory to specify at least one permission to be set.
When specifying multiple permissions, they must be separated by a comma (,).`,
		Example: fmt.Sprintf(`
%s tx subspaces set-user-permissions 1 desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7 "Write,ModerateContent,SetUserPermissions" \
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

			user := args[1]

			permission := types.PermissionNothing
			for _, arg := range strings.Split(args[2], ",") {
				perm, err := types.ParsePermission(arg)
				if err != nil {
					return err
				}
				permission = types.CombinePermissions(permission, perm)
			}

			msg := types.NewMsgSetUserPermissions(subspaceID, user, permission, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
