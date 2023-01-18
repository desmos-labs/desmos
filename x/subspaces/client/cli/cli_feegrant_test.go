//go:build norace
// +build norace

package cli_test

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v4/x/subspaces/client/cli"
	"github.com/gogo/protobuf/proto"
)

func (s *IntegrationTestSuite) TestCmdGrantAllowance() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"id"},
			shouldErr: true,
		},
		{
			name:      "empty grantee returns error",
			args:      []string{"1"},
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				fmt.Sprintf("--%s=%s", cli.FlagGroupGrantee, "1"),
			},
			shouldErr: true,
		},
		{
			name: "invalid msg returns error",
			args: []string{
				"0",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			},
			shouldErr: true,
		},
		{
			name: "invalid expiration returns error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				fmt.Sprintf("--%s=%s", cli.FlagExpiration, "invalid"),
			},
			shouldErr: true,
		},
		{
			name: "invalid periodic allowance returns error - period is greater than expiration",
			args: append(
				[]string{
					"1",
					fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					fmt.Sprintf("--%s=%s", cli.FlagSpendLimit, "100stake"),
					fmt.Sprintf("--%s=%s", cli.FlagExpiration, getFormattedExpiration(3600)),
					fmt.Sprintf("--%s=%d", cli.FlagPeriod, 36000),
					fmt.Sprintf("--%s=%s", cli.FlagPeriodLimit, "10stake"),
				},
			),
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "invalid periodic allowance returns error - invalid number of args",
			args: append(
				[]string{
					"1",
					fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					fmt.Sprintf("--%s=%s", cli.FlagSpendLimit, "100stake"),
					fmt.Sprintf("--%s=%s", cli.FlagPeriodLimit, "10stake"),
				},
			),
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data with default allowance returns no error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data with basic allowance returns no error",
			args: append(
				[]string{
					"1",
					fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					fmt.Sprintf("--%s=%s", cli.FlagSpendLimit, "100stake"),
					fmt.Sprintf("--%s=%s", cli.FlagExpiration, getFormattedExpiration(3600)),
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
					fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
				},
			),
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data with periodic allowance returns no error",
			args: append(
				[]string{
					"1",
					fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					fmt.Sprintf("--%s=%s", cli.FlagSpendLimit, "100stake"),
					fmt.Sprintf("--%s=%d", cli.FlagPeriod, 3600),
					fmt.Sprintf("--%s=%s", cli.FlagPeriodLimit, "10stake"),
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
					fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
				},
			),
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data with filtered allowance returns no error",
			args: append(
				[]string{
					"1",
					fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					fmt.Sprintf("--%s=%s", cli.FlagAllowedMsgs, "/desmos.posts.v2.MsgCreatPost"),
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
					fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
				},
			),
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdGrantAllowance()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdRevokeAllowance() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"id"},
			shouldErr: true,
		},
		{
			name:      "empty grantee returns error",
			args:      []string{"1"},
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				fmt.Sprintf("--%s=%s", cli.FlagGroupGrantee, "1"),
			},
			shouldErr: true,
		},
		{
			name: "invalid msg returns error",
			args: []string{
				"0",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdRevokeAllowance()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func getFormattedExpiration(duration int64) string {
	return time.Now().Add(time.Duration(duration) * time.Second).Format(time.RFC3339)
}
