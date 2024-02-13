package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantcli "github.com/cosmos/cosmos-sdk/x/feegrant/client/cli"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/pflag"

	subspacesauthz "github.com/desmos-labs/desmos/v7/x/subspaces/authz"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	authzcli "github.com/cosmos/cosmos-sdk/x/authz/client/cli"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

// DONTCOVER

const (
	FlagName           = "name"
	FlagDescription    = "description"
	FlagParent         = "parent"
	FlagSection        = "section"
	FlagOwner          = "owner"
	FlagPermissions    = "permissions"
	FlagInitialMembers = "initial-members"
	FlagUserGrantee    = "user"
	FlagGroupGrantee   = "group"

	delegate   = "delegate"
	redelegate = "redelegate"
	unbond     = "unbond"
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
		GetTreasuryTxCmd(),
		GetAllowancesTxCmd(),
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

			owner, err := cmd.Flags().GetString(FlagOwner)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSubspace(name, description, owner, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagDescription, "", "Description of the subspace")
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

			owner, err := cmd.Flags().GetString(FlagOwner)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditSubspace(subspaceID, name, description, owner, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagName, types.DoNotModify, "New human readable name of the subspace")
	cmd.Flags().String(FlagDescription, types.DoNotModify, "Description of the subspace")
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
%s tx subspaces groups set-permissions 1 1 "WRITE_CONTENT,INTERACT_WITH_CONTENT,EDIT_OWN_CONTENT,MODERATE_CONTENT,
EDIT_SUBSPACE,DELETE_SUBSPACE,MANAGE_SECTIONS,MANAGE_GROUPS,SET_PERMISSIONS,EVERYTHING" \
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
  "WRITE_CONTENT,INTERACT_WITH_CONTENT,EDIT_OWN_CONTENT,MODERATE_CONTENT,
EDIT_SUBSPACE,DELETE_SUBSPACE,MANAGE_SECTIONS,MANAGE_GROUPS,SET_PERMISSIONS,EVERYTHING" \
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
			expiration := time.Unix(exp, 0)
			msg, err := authz.NewMsgGrant(clientCtx.GetFromAddress(), grantee, authorization, &expiration)
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

// --------------------------------------------------------------------------------------------------------------------

// GetTreasuryTxCmd returns a new command to perform subspaces treasury transactions
func GetTreasuryTxCmd() *cobra.Command {
	treasuryTxCmd := &cobra.Command{
		Use:                        "treasury",
		Short:                      "Tx commands for subspace treasury",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	treasuryTxCmd.AddCommand(
		GetCmdGrantTreasuryAuthorization(),
		GetCmdRevokeTreasuryAuthorization(),
	)

	return treasuryTxCmd
}

// GetCmdGrantTreasuryAuthorization returns the command used to grant a treasury authorization
func GetCmdGrantTreasuryAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant [subspace-id] [grantee] [authorization_type=\"send\"|\"generic\"|\"delegate\"|\"unbond\"|\"redelegate\"] --from [granter]",
		Short: "Grant a treasury authorization to a user",
		Long: strings.TrimSpace(
			fmt.Sprintf(`grant treasury authorization to an address to execute a transaction on your behalf:
Examples:
 $ %[1]s tx %[2]s grant desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7 send --spend-limit=1000stake --from=desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7
 $ %[1]s tx %[2]s grant desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7 generic --msg-type=/cosmos.gov.v1beta1.MsgVote --from=desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7
	`, version.AppName, types.ModuleName),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			exp, err := cmd.Flags().GetInt64(authzcli.FlagExpiration)
			if err != nil {
				return err
			}

			var authorization authz.Authorization
			switch args[2] {
			case "send":
				authorization, err = getSendAuthorization(cmd.Flags())
				if err != nil {
					return err
				}

			case "generic":
				authorization, err = getGenericAuthorization(cmd.Flags())
				if err != nil {
					return err
				}

			case delegate, unbond, redelegate:
				authorization, err = getStakeAuthorization(cmd.Flags(), args[2])
				if err != nil {
					return err
				}

			default:
				return fmt.Errorf("invalid authorization type, %s", args[2])
			}

			expiration := time.Unix(exp, 0)
			msg := types.NewMsgGrantTreasuryAuthorization(subspaceID, clientCtx.GetFromAddress().String(), args[1], authorization, &expiration)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(authzcli.FlagMsgType, "", "The Msg method name for which we are creating a GenericAuthorization")
	cmd.Flags().String(authzcli.FlagSpendLimit, "", "SpendLimit for Send Authorization, an array of Coins allowed spend")
	cmd.Flags().StringSlice(authzcli.FlagAllowedValidators, []string{}, "Allowed validators addresses separated by ,")
	cmd.Flags().StringSlice(authzcli.FlagDenyValidators, []string{}, "Deny validators addresses separated by ,")
	cmd.Flags().Int64(authzcli.FlagExpiration, time.Now().AddDate(1, 0, 0).Unix(), "The Unix timestamp. Default is one year.")
	cmd.Flags().StringSlice(authzcli.FlagAllowList, []string{}, "Allowed addresses grantee is allowed to send funds separated by ,")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// getStakeAuthorization returns a send authorization from the given command flags
func getSendAuthorization(flags *pflag.FlagSet) (*banktypes.SendAuthorization, error) {
	limit, err := flags.GetString(authzcli.FlagSpendLimit)
	if err != nil {
		return nil, err
	}

	spendLimit, err := sdk.ParseCoinsNormalized(limit)
	if err != nil {
		return nil, err
	}

	if !spendLimit.IsAllPositive() {
		return nil, fmt.Errorf("spend-limit should be greater than zero")
	}

	allowed, err := getAllowedListFromFlags(flags)
	if err != nil {
		return nil, err
	}

	return banktypes.NewSendAuthorization(spendLimit, allowed), nil
}

// getStakeAuthorization returns a generic authorization from the given command flags
func getGenericAuthorization(flags *pflag.FlagSet) (*authz.GenericAuthorization, error) {
	msgType, err := flags.GetString(authzcli.FlagMsgType)
	if err != nil {
		return nil, err
	}
	return authz.NewGenericAuthorization(msgType), nil
}

// getStakeAuthorization returns a stake authorization from the given command flags
func getStakeAuthorization(flags *pflag.FlagSet, stakingType string) (*stakingtypes.StakeAuthorization, error) {
	limit, err := flags.GetString(authzcli.FlagSpendLimit)
	if err != nil {
		return nil, err
	}

	var delegateLimit *sdk.Coin
	if limit != "" {
		spendLimit, err := sdk.ParseCoinNormalized(limit)
		if err != nil {
			return nil, err
		}

		if !spendLimit.IsPositive() {
			return nil, fmt.Errorf("spend-limit should be greater than zero")
		}
		delegateLimit = &spendLimit
	}

	allowed, err := getValidatorAddressesFromFlags(flags, authzcli.FlagAllowedValidators)
	if err != nil {
		return nil, err
	}

	denied, err := getValidatorAddressesFromFlags(flags, authzcli.FlagDenyValidators)
	if err != nil {
		return nil, err
	}

	var authorizationType stakingtypes.AuthorizationType
	switch stakingType {
	case delegate:
		authorizationType = stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE
	case unbond:
		authorizationType = stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE
	default:
		authorizationType = stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE
	}

	return stakingtypes.NewStakeAuthorization(allowed, denied, authorizationType, delegateLimit)
}

// getValidatorAddressesFromFlags returns validator addresses with type (allowed or deny) from flags
func getValidatorAddressesFromFlags(flags *pflag.FlagSet, typ string) ([]sdk.ValAddress, error) {
	validators, err := flags.GetStringSlice(typ)
	if err != nil {
		return nil, err
	}

	validatorAddrs := make([]sdk.ValAddress, len(validators))
	for i, validator := range validators {
		addr, err := sdk.ValAddressFromBech32(validator)
		if err != nil {
			return nil, err
		}
		validatorAddrs[i] = addr
	}
	return validatorAddrs, nil
}

// getAllowedListFromFlags converts the allowed list addresses into sdk.AccAddress instances
func getAllowedListFromFlags(flags *pflag.FlagSet) ([]sdk.AccAddress, error) {
	allowList, err := flags.GetStringSlice(authzcli.FlagAllowList)
	if err != nil {
		return nil, err
	}

	addrs := make([]sdk.AccAddress, len(allowList))
	for i, addr := range allowList {
		accAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		addrs[i] = accAddr
	}
	return addrs, nil
}

// GetCmdRevokeTreasuryAuthorization returns the command used to revoke a treasury authorization
func GetCmdRevokeTreasuryAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke [subspace-id] [grantee] [msg_type] --from=[granter]",
		Short: "revoke a treasury authorization",
		Long: strings.TrimSpace(
			fmt.Sprintf(`revoke treasury authorization from a granter to a grantee:
Example:
 $ %s tx %s revoke desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7 %s --from=desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7
			`, version.AppName, authz.ModuleName, banktypes.SendAuthorization{}.MsgTypeURL()),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeTreasuryAuthorization(subspaceID, clientCtx.GetFromAddress().String(), args[1], args[2])
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

// GetAllowancesTxCmd returns a new command to perform subspaces treasury transactions
func GetAllowancesTxCmd() *cobra.Command {
	treasuryTxCmd := &cobra.Command{
		Use:                        "allowances",
		Short:                      "Tx commands for subspace treasury",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	treasuryTxCmd.AddCommand(
		GetCmdGrantAllowance(),
		GetCmdRevokeAllowance(),
	)

	return treasuryTxCmd
}

// GetCmdGrantAllowance returns the command used to grant a fee allowance
func GetCmdGrantAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant [subspace-id]",
		Short: "Grant a fee allowance to an address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			grantee, err := getGranteeFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			allowance, err := getAllowanceFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgGrantAllowance(subspaceID, clientCtx.FromAddress.String(), grantee, allowance)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagUserGrantee, "", "Address of the user being the allowance grantee")
	cmd.Flags().Uint32(FlagGroupGrantee, 0, "Id of group being the allowance grantee")
	cmd.Flags().StringSlice(feegrantcli.FlagAllowedMsgs, []string{}, "Set of allowed messages for fee allowance")
	cmd.Flags().String(feegrantcli.FlagExpiration, "", "The RFC 3339 timestamp after which the grant expires for the user")
	cmd.Flags().String(feegrantcli.FlagSpendLimit, "", "Spend limit specifies the max limit can be used, if not mentioned there is no limit")
	cmd.Flags().Int64(feegrantcli.FlagPeriod, 0, "Period specifies the time duration in which period_spend_limit coins can be spent before that allowance is reset")
	cmd.Flags().String(feegrantcli.FlagPeriodLimit, "", "Period limit specifies the maximum number of coins that can be spent in the period")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRevokeAllowance returns the command used to revoke a fee allowance
func GetCmdRevokeAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke [subspace-id]",
		Short: "Revoke a fee allowance from an address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			grantee, err := getGranteeFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeAllowance(subspaceID, clientCtx.FromAddress.String(), grantee)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagUserGrantee, "", "Address of the user being the allowance grantee")
	cmd.Flags().Uint32(FlagGroupGrantee, 0, "Id of group being the allowance grantee")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// getGranteeFromFlags returns a grantee from flags
func getGranteeFromFlags(flags *pflag.FlagSet) (types.Grantee, error) {
	userGrantee, err := flags.GetString(FlagUserGrantee)
	if err != nil {
		return nil, err
	}

	groupGrantee, err := flags.GetUint32(FlagGroupGrantee)
	if err != nil {
		return nil, err
	}

	switch {
	case userGrantee != "" && groupGrantee != 0:
		return nil, fmt.Errorf("only one of --%s or --%s must be used", FlagUserGrantee, FlagGroupGrantee)

	case userGrantee != "":
		return types.NewUserGrantee(userGrantee), nil

	case groupGrantee != 0:
		return types.NewGroupGrantee(groupGrantee), nil
	}

	return nil, fmt.Errorf("one of --%s or --%s must be used", FlagUserGrantee, FlagGroupGrantee)
}

// getAllowanceFromFlags returns an allowance from flags
func getAllowanceFromFlags(flags *pflag.FlagSet) (feegrant.FeeAllowanceI, error) {
	spendLimit, err := flags.GetString(feegrantcli.FlagSpendLimit)
	if err != nil {
		return nil, err
	}

	// if `FlagSpendLimit` isn't set, limit will be nil
	limit, err := sdk.ParseCoinsNormalized(spendLimit)
	if err != nil {
		return nil, err
	}

	expired, err := flags.GetString(feegrantcli.FlagExpiration)
	if err != nil {
		return nil, err
	}

	periodClock, err := flags.GetInt64(feegrantcli.FlagPeriod)
	if err != nil {
		return nil, err
	}

	periodLimit, err := flags.GetString(feegrantcli.FlagPeriodLimit)
	if err != nil {
		return nil, err
	}

	allowedMsgs, err := flags.GetStringSlice(feegrantcli.FlagAllowedMsgs)
	if err != nil {
		return nil, err
	}

	// Build basic allowance
	var allowance feegrant.FeeAllowanceI
	basic := feegrant.BasicAllowance{
		SpendLimit: limit,
	}

	// Add expiration to allowance if expiration is set
	var expiresAtTime time.Time
	if expired != "" {
		expiresAtTime, err = time.Parse(time.RFC3339, expired)
		if err != nil {
			return nil, err
		}
		basic.Expiration = &expiresAtTime
	}
	allowance = &basic

	// Check any of period or periodLimit flags set, consider it as periodic fee allowance if set
	if periodClock > 0 || periodLimit != "" {
		periodLimit, err := sdk.ParseCoinsNormalized(periodLimit)
		if err != nil {
			return nil, err
		}

		if periodClock <= 0 {
			return nil, fmt.Errorf("period clock was not set")
		}

		if periodLimit == nil {
			return nil, fmt.Errorf("period limit was not set")
		}

		periodReset := getPeriodReset(periodClock)
		if basic.Expiration != nil && periodReset.Sub(expiresAtTime) > 0 {
			return nil, fmt.Errorf("period (%d) cannot reset after expiration (%v)", periodClock, expired)
		}

		periodAllowance := &feegrant.PeriodicAllowance{
			Basic:            basic,
			Period:           getPeriod(periodClock),
			PeriodReset:      periodReset,
			PeriodSpendLimit: periodLimit,
			PeriodCanSpend:   periodLimit,
		}

		allowance = periodAllowance
	}

	// Check if allowedMsgs flags set, consider it as allowed msg allowance if set
	if len(allowedMsgs) > 0 {
		filteredAllowance, err := feegrant.NewAllowedMsgAllowance(allowance, allowedMsgs)
		if err != nil {
			return nil, err
		}
		allowance = filteredAllowance
	}

	return allowance, nil
}

// getPeriodReset generates a next period reset time from a duration
func getPeriodReset(duration int64) time.Time {
	return time.Now().Add(getPeriod(duration))
}

// getPeriod turns duration type from int64 into time.Duration
func getPeriod(duration int64) time.Duration {
	return time.Duration(duration) * time.Second
}
