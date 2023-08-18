package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

// DONTCOVER
// Tests will use single commands and not the global tx one

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
		GetCmdLinkChainAccount(),
		GetCmdUnlinkChainAccount(),
		GetCmdSetDefaultExternalAddress(),
		GetCmdLinkApplication(),
		GetCmdUnlinkApplication(),
	)

	return profileTxCmd
}
