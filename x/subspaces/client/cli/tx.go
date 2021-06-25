package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/subspaces/types"
)

const (
	FlagSubspaceType = "type"
	FlagName         = "name"
	FlagDescription  = "description"
	FlagLogo         = "logo"
	FlagOwner        = "owner"

	DoNotEdit = "do-not-edit"
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
		GetCmdAddAdmin(),
		GetCmdRemoveAdmin(),
		GetCmdRegisterUser(),
		GetCmdUnregisterUser(),
		GetCmdBanUser(),
		GetCmdUnbanUser(),
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

The name shall be a human readable name, while the --type flag can be used to tell whether 
the subspace allow users to post messages freely. 
e.g 1) %s tx subspaces create 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e "mooncake" --type open
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id := args[0]
			name := args[1]
			description, _ := cmd.Flags().GetString(FlagDescription)
			logo, _ := cmd.Flags().GetString(FlagLogo)

			subType, _ := cmd.Flags().GetString(FlagSubspaceType)
			subspaceType, err := types.SubspaceTypeFromString(types.NormalizeSubspaceType(subType))
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSubspace(
				id,
				name,
				description,
				logo,
				clientCtx.FromAddress.String(),
				subspaceType,
			)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagSubspaceType, "close", "Tells if the subspace let post messages freely or not")
	cmd.Flags().String(FlagDescription, "", "The description of the subspace")
	cmd.Flags().String(FlagLogo, "", "The logo of the subspace")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditSubspace returns the command to edit a subspace
func GetCmdEditSubspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [subspace-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Edit an existent subspace with the given id",
		Long: fmt.Sprintf(`Create a new subspace.
The id must be a valid SHA-256 hash uniquely identifying the subspace.

The name shall be a human readable name, while the --type flag can be used to tell whether 
the subspace allow users to post messages freely. 
e.g 1) %s tx subspaces edit 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e --name "new" --type "open"
`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID := args[0]

			owner, _ := cmd.Flags().GetString(FlagOwner)
			name, _ := cmd.Flags().GetString(FlagName)
			description, _ := cmd.Flags().GetString(FlagDescription)
			logo, _ := cmd.Flags().GetString(FlagLogo)

			subType, _ := cmd.Flags().GetString(FlagSubspaceType)
			subspaceType, err := types.SubspaceTypeFromString(types.NormalizeSubspaceType(subType))
			if err != nil && subType != types.DoNotModify {
				return err
			}

			editor := clientCtx.FromAddress.String()
			msg := types.NewMsgEditSubspace(subspaceID, owner, name, description, logo, editor, subspaceType)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagName, types.DoNotModify, "New human readable name of the subspace")
	cmd.Flags().String(FlagOwner, "", "New owner of the subspace")
	cmd.Flags().String(FlagSubspaceType, types.DoNotModify, "Tells if the subspace let post messages freely or not")
	cmd.Flags().String(FlagDescription, types.DoNotModify, "The description of the subspace")
	cmd.Flags().String(FlagLogo, types.DoNotModify, "The logo of the subspace")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAddAdmin returns the command to add an admin to a subspace
func GetCmdAddAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-admin [subspace-id] [address]",
		Args:  cobra.ExactArgs(2),
		Short: "Add a new admin to the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			newAdminAddress, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			subspaceID := args[0]

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

// GetCmdRemoveAdmin returns the command to remove an admin from a subspace
func GetCmdRemoveAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-admin [subspace-id] [address]",
		Args:  cobra.ExactArgs(2),
		Short: "Remove an existent admin from the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			existentAdminAddress, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			subspaceID := args[0]

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
		Use:   "register-user [subspace-id] [address]",
		Args:  cobra.ExactArgs(2),
		Short: "Register a user inside the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID := args[0]
			user := args[1]
			admin := clientCtx.FromAddress.String()

			msg := types.NewMsgRegisterUser(subspaceID, user, admin)
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
		Use:   "unregister-user [subspace-id] [address]",
		Args:  cobra.ExactArgs(2),
		Short: "Unregister a user from the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID := args[0]
			user := args[1]
			admin := clientCtx.FromAddress.String()
			msg := types.NewMsgUnregisterUser(subspaceID, user, admin)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdBanUser returns the command to ban a user inside a subspace
func GetCmdBanUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ban-user [subspace-id] [address]",
		Args:  cobra.ExactArgs(2),
		Short: "Ban a user inside the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID := args[0]
			user := args[1]
			admin := clientCtx.FromAddress.String()
			msg := types.NewMsgBanUser(subspaceID, user, admin)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdUnbanUser returns the command to unban a user from a subspace
func GetCmdUnbanUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unban-user [subspace-id] [address]",
		Args:  cobra.ExactArgs(2),
		Short: "Unban a user inside the subspace with the given id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID := args[0]
			user := args[1]
			admin := clientCtx.FromAddress.String()
			msg := types.NewMsgBanUser(subspaceID, user, admin)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
