package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// NewTxCmd returns a new command allowing to perform profiles transactions
func NewTxCmd() *cobra.Command {
	profileTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Profiles transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	profileTxCmd.AddCommand(
		GetCmdSaveProfile(),
		GetCmdDeleteProfile(),
		GetCmdRequestDTagTransfer(),
		GetCmdAcceptDTagTransfer(),
		GetCmdRefuseDTagTransfer(),
		GetCmdCancelDTagTransfer(),
		GetCmdCreateRelationship(),
		GetCmdDeleteRelationship(),
		GetCmdBlockUser(),
		GetCmdUnblockUser(),
		GetCmdLinkChainAccount(),
		GetCmdUnlinkChainAccount(),
		GetCmdLinkApplication(),
		GetCmdUnlinkApplication(),
	)

	return profileTxCmd
}
