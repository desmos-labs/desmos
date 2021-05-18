package cli

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"

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
		GetCmdRegisterUser(),
		GetCmdBlockUser(),
		GetCmdEditSubspace(),
	)

	return subspacesTxCmd
}

// GetCmdCreateSubspace returns the command used to create a subspace
func GetCmdCreateSubspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [subspace-id] [name] [[open]]",
		Args:  cobra.RangeArgs(2, 3),
		Short: "Create a subspace with the given [subspace-id], [name] and [open]",
		Long: fmt.Sprintf(`The [subspace-id] must be a sha256 string identifying the subspace, 
the [name] a human readable name and the [open] a boolean identifying whether the subspace allow users posts without registration.
Leaving [open] empty will set it to false by default. 
e.g 1) %s tx subspaces create 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e "mooncake"
	2) %s tx subspaces create 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e "mooncake" true
`, version.AppName, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID := args[0]
			subspaceName := args[1]
			open := false
			if len(args) > 2 {
				open, err = strconv.ParseBool(args[2])
				if err != nil {
					return fmt.Errorf("open field can only be true or false")
				}
			}

			msg := types.NewMsgCreateSubspace(subspaceID, subspaceName, clientCtx.FromAddress.String(), open)
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
		Short: "Add a new admin to the subspace with the given [subspace-id]",
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

func GetCmdRegisterUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-user [address] [subspace-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Register a user inside the subspace with the given [subspace-id] to let him post in it",
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

func GetCmdBlockUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block-user [address] [subspace-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Block a user to post inside the subspace with the given [subspace-id]",
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
		Use:   "edit [subspace-id] [[new_owner]]",
		Args:  cobra.RangeArgs(1, 2),
		Short: "Edit an existent subspace with the given [subspace-id]",
		Long: fmt.Sprintf(`Edit a subspace with the given [subspace-id].
E.g 
1) Edit the owner only
%s tx subspaces edit 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"
2) Edit the name only
%s tx subspaces edit 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e 
	--new-name "star"
3) Edit both owner and name
%s tx subspaces edit 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"
	--new-name "star"
`, version.AppName, version.AppName, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID := args[0]
			var newOwner string

			if len(args) > 1 {
				newOwner = args[1]
			}

			newName := viper.GetString(FlagNewName)

			owner := clientCtx.FromAddress.String()
			msg := types.NewMsgEditSubspace(subspaceID, newOwner, newName, owner)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	cmd.Flags().String(FlagNewName, "", "New human readable name of the subspace")

	return cmd
}
