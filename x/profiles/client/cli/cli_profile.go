package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func GetCmdSaveProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Args:  cobra.NoArgs,
		Short: "Save your profile",
		Long: fmt.Sprintf(`
 Save a new profile or edit the existing one specifying a DTag, a nickname, biography, profile picture and cover picture.
 Every data given through the flags is optional.
 If you are editing an existing profile you should fill only the fields that you want to edit.
 The empty ones will be filled with a special [do-not-modify] flag that tells the system to not edit them.

 %s tx profiles save 
 	--%s "LeoDiCaprio" \
 	--%s "Leonardo Di Caprio" \
 	--%s "Hollywood actor. Proud environmentalist" \
 	--%s "https://profilePic.jpg" \
 	--%s "https://profileCover.jpg"
 `, version.AppName, FlagDTag, FlagNickname, FlagBio, FlagProfilePic, FlagCoverPic),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			dTag, _ := cmd.Flags().GetString(FlagDTag)
			nickname, _ := cmd.Flags().GetString(FlagNickname)
			bio, _ := cmd.Flags().GetString(FlagBio)
			profilePic, _ := cmd.Flags().GetString(FlagProfilePic)
			coverPic, _ := cmd.Flags().GetString(FlagCoverPic)

			msg := types.NewMsgSaveProfile(dTag, nickname, bio, profilePic, coverPic, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagDTag, types.DoNotModify, "DTag to be used")
	cmd.Flags().String(FlagNickname, types.DoNotModify, "Nickname to be used")
	cmd.Flags().String(FlagBio, types.DoNotModify, "Biography to be used")
	cmd.Flags().String(FlagProfilePic, types.DoNotModify, "Profile picture")
	cmd.Flags().String(FlagCoverPic, types.DoNotModify, "Cover picture")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDeleteProfile returns the command used to delete an existing profile
func GetCmdDeleteProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete an existent profile related to the user's address",
		Example: fmt.Sprintf(`%s tx profiles delete`, version.AppName),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteProfile(clientCtx.FromAddress.String())
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

// GetCmdQueryProfile returns the command that allows to query the profile of a specific user
func GetCmdQueryProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile [address_or_dtag]",
		Short: "Retrieve the profile having the specified user address or profile dtag, if any.",
		Example: fmt.Sprintf(`%s query profiles desmos12a2y7fflz6g4e5gn0mh0n9dkrzllj0q5vx7c6t
%s query profiles Alice`, version.AppName, version.AppName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Profile(
				context.Background(),
				&types.QueryProfileRequest{User: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
