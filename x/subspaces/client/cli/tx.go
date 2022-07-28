package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/x/authz"

	subspacesauthz "github.com/desmos-labs/desmos/v4/x/subspaces/authz"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	authzcli "github.com/cosmos/cosmos-sdk/x/authz/client/cli"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// DONTCOVER

const (
	FlagName           = "name"
	FlagDescription    = "description"
	FlagParent         = "parent"
	FlagSection        = "section"
	FlagTreasury       = "treasury"
	FlagOwner          = "owner"
	FlagPermissions    = "permissions"
	FlagInitialMembers = "initial-members"
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
		NewSectionsTxCmd(),
		NewGroupsTxCmd(),
		GetCmdSetUserPermissions(),
		GetCmdGrantAuthorization(),
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

// NewSectionsTxCmd returns a new command to perform subspaces sections transactions
func NewSectionsTxCmd() *cobra.Command {
	groupsTxCmd := &cobra.Command{
		Use:                        "sections",
		Short:                      "Subspace sections transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	groupsTxCmd.AddCommand(
		GetCmdCreateSection(),
		GetCmdEditSection(),
		GetCmdMoveSection(),
		GetCmdDeleteSection(),
	)

	return groupsTxCmd
}

// GetCmdCreateSection returns the command used to create a subspace section
func GetCmdCreateSection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [subspace-id] [name]",
		Args:  cobra.ExactArgs(2),
		Short: "Create a new section within a subspace",
		Long: `Create a new section within the subspace with the provided id.
The name must be a human readable name.`,
		Example: fmt.Sprintf(`
%s tx subspaces sections create 1 "Custom section" \
  --description "This is my custom section" \
  --parent 1 \
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

			parentID, err := cmd.Flags().GetUint32(FlagParent)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSection(subspaceID, name, description, parentID, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagDescription, "", "Description of the section")
	cmd.Flags().Uint32(FlagParent, 0, "Id of the parent section")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditSection returns the command to edit a subspace section
func GetCmdEditSection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [subspace-id] [section-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Edit the subspace section with the given id",
		Example: fmt.Sprintf(`
%s tx subspaces edit 1 1 \
  --name "Desmos - Democratizing social networks"
  --description "The official subspace of Desmos" \
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

			sectionID, err := types.ParseSectionID(args[1])
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

			msg := types.NewMsgEditSection(subspaceID, sectionID, name, description, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagName, types.DoNotModify, "New human readable name of the section")
	cmd.Flags().String(FlagDescription, types.DoNotModify, "New description of the section")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdMoveSection returns the command to delete a subspace section
func GetCmdMoveSection() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "move [subspace-id] [section-id] [new-parent-id]",
		Args:    cobra.ExactArgs(3),
		Short:   "Move the subspace section with the given id to the new parent",
		Example: fmt.Sprintf(`%s tx subspaces delete 1 1 2 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			sectionID, err := types.ParseSectionID(args[1])
			if err != nil {
				return err
			}

			newParentID, err := types.ParseSectionID(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgMoveSection(subspaceID, sectionID, newParentID, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDeleteSection returns the command to delete a subspace section
func GetCmdDeleteSection() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [subspace-id] [section-id]",
		Args:    cobra.ExactArgs(2),
		Short:   "Deletes the subspace section with the given id",
		Example: fmt.Sprintf(`%s tx subspaces delete 1 1 --from alice`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			sectionID, err := types.ParseSectionID(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteSection(subspaceID, sectionID, clientCtx.FromAddress.String())
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
		GetCmdMoveUserGroup(),
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
		Args:  cobra.ExactArgs(2),
		Short: "Create a new user group within a subspace",
		Long: fmt.Sprintf(`Create a new user group within the subspace having the provided id.

An optional description of this group can be provided with the %[1]s flag.

The permissions of this group can be set using the %[2]s flag.
If no permissions are set, the default PermissionNothing will be used instead.
Multiple permissions must be specified separating them with a comma (,).`, FlagDescription, FlagPermissions),
		Example: fmt.Sprintf(`
%s tx subspaces groups create 1 "Admins" \
  --description "Group of the subspace admins" \
  --permissions "WRITE,MODERATE_CONTENT,SET_USER_PERMISSIONS" \
  --initial-members "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53,cosmos1g4yzh3q3grf804t4y4fuynrvrxtshgxy7j783f" \
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

			sectionID, err := cmd.Flags().GetUint32(FlagSection)
			if err != nil {
				return err
			}

			flagPermissions, err := cmd.Flags().GetStringSlice(FlagPermissions)
			if err != nil {
				return err
			}

			var permissions types.Permissions
			for _, arg := range flagPermissions {
				permissions = types.CombinePermissions(append(permissions, arg)...)
			}

			initialMembers, err := cmd.Flags().GetStringSlice(FlagInitialMembers)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUserGroup(subspaceID, sectionID, name, description, permissions, initialMembers, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint32(FlagSection, 0, "Id of the section inside which to create the group")
	cmd.Flags().String(FlagDescription, "", "Description of the group")
	cmd.Flags().StringSlice(FlagPermissions, nil, "Permissions of the group")
	cmd.Flags().StringSlice(FlagInitialMembers, nil, "Initial members of the group")

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

// GetCmdMoveUserGroup returns the command to move a user group to another section
func GetCmdMoveUserGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "move [subspace-id] [group-id] [new-section-id]",
		Args:    cobra.ExactArgs(3),
		Short:   "Move a user group to a new section",
		Example: fmt.Sprintf(`%s tx subspaces groups move 1 1 2 --from alice`, version.AppName),
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

			newSectionID, err := types.ParseSectionID(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgMoveUserGroup(subspaceID, groupID, newSectionID, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

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
%s tx subspaces groups set-permissions 1 1 "WRITE,MODERATE_CONTENT,SET_USER_PERMISSIONS" \
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

			var permissions types.Permissions
			for _, arg := range strings.Split(args[2], ",") {
				permissions = types.CombinePermissions(append(permissions, arg)...)
			}

			msg := types.NewMsgSetUserGroupPermissions(subspaceID, groupID, permissions, clientCtx.FromAddress.String())
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
%s tx subspaces set-user-permissions 1 \
  desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7 \
  "WRITE,MODERATE_CONTENT,SET_USER_PERMISSIONS" \
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

			sectionID, err := cmd.Flags().GetUint32(FlagSection)
			if err != nil {
				return err
			}

			user := args[1]

			var permissions types.Permissions
			for _, arg := range strings.Split(args[2], ",") {
				permissions = types.CombinePermissions(append(permissions, arg)...)
			}

			msg := types.NewMsgSetUserPermissions(subspaceID, sectionID, user, permissions, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint32(FlagSection, 0, "Id of the section inside which to set the permissions")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// --------------------------------------------------------------------------------------------------------------------

// GetCmdGrantAuthorization returns the command to grant a subspace authorization
func GetCmdGrantAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant [subspaces-ids] [grantee]",
		Short: "Grant an authorization to an address inside one or more subspaces",
		Long: `Grant an authorization to an address inside one or more subspaces.
If you want to grant the same authorization inside multiple subspaces, simply specify the subspaces ids separating them with a comma (,).`,
		Example: fmt.Sprintf(`
%s tx subspaces grant 1,2,3 desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7 --msg-type=%s --from alice
`, version.AppName, sdk.MsgTypeURL(&types.MsgSetUserPermissions{})),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspacesIDs, err := types.ParseSubspacesIDs(args[0])
			if err != nil {
				return err
			}

			grantee, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msgType, err := cmd.Flags().GetString(authzcli.FlagMsgType)
			if err != nil {
				return err
			}

			exp, err := cmd.Flags().GetInt64(authzcli.FlagExpiration)
			if err != nil {
				return err
			}

			authorization := subspacesauthz.NewGenericSubspaceAuthorization(subspacesIDs, msgType)
			msg, err := authz.NewMsgGrant(clientCtx.GetFromAddress(), grantee, authorization, time.Unix(exp, 0))
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(authzcli.FlagMsgType, "", "The msg method name for which we are creating the authorization")
	cmd.Flags().Int64(authzcli.FlagExpiration, time.Now().AddDate(1, 0, 0).Unix(), "The Unix timestamp. Default is one year.")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
