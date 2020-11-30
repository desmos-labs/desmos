package testutil

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"

	postscli "github.com/desmos-labs/desmos/x/posts/client/cli"
)

var commonArgs = []string{
	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
}

// MsgCreatePost creates a tx for creating a new post
func MsgCreatePost(clientCtx client.Context, from, subspace string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := append([]string{
		subspace,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from),
	}, commonArgs...)

	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, postscli.GetCmdCreatePost(), args)
}

// QueryPostsExec queries the posts present on chain
func QueryPostsExec(clientCtx client.Context, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{fmt.Sprintf("--%s=json", cli.OutputFlag)}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, postscli.GetCmdQueryPosts(), args)
}
