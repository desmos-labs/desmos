package testutil

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"

	profilescli "github.com/desmos-labs/desmos/x/profiles/client/cli"
)

var commonArgs = []string{
	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
}

// MsgSaveProfile creates a tx for creating a relationship
func MsgSaveProfile(clientCtx client.Context, dtag, from string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := append([]string{
		dtag,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}, commonArgs...)

	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, profilescli.GetCmdSaveProfile(), args)
}

// NewMsgRequestDTagTransfer creates a tx for blocking a user
func NewMsgRequestDTagTransfer(clientCtx client.Context, user, from string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := append([]string{
		user,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}, commonArgs...)

	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, profilescli.GetCmdRequestDTagTransfer(), args)
}
