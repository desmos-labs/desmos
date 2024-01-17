//go:build norace
// +build norace

package cli_test

import (
	"fmt"

	"cosmossdk.io/math"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"

	"github.com/desmos-labs/desmos/v6/x/profiles/client/cli"
	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func (s *IntegrationTestSuite) TestCmdQueryDTagRequests() {
	val := s.network.Validators[0]

	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expRequests []types.DTagTransferRequest
	}{
		{
			name: "existing requests are returned properly",
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos122u6u9gpdr2rp552fkkvlgyecjlmtqhkascl5a",
					"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				),
			},
		},
		{
			name: "empty slice is returned properly",
			args: []string{
				"cosmos1nqwf7chwfywdw2379sxmwlcgcfvvy86t6mpunz",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr:   false,
			expRequests: []types.DTagTransferRequest{},
		},
		{
			name: "existing requests of the given user are returned properly",
			args: []string{
				"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos122u6u9gpdr2rp552fkkvlgyecjlmtqhkascl5a",
					"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryDTagRequests()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryIncomingDTagTransferRequestsResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expRequests, response.Requests)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdRequestDTagTransfer() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid user address returns error",
			args: []string{
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "same user returns error",
			args: []string{
				val.Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, math.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdRequestDTagTransfer()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdCancelDTagTransfer() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid recipient returns error",
			args: []string{
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "same user and recipient returns error",
			args: []string{
				val.Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, math.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdCancelDTagTransfer()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdAcceptDTagTransfer() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid new dtag returns error",
			args: []string{
				"",
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid request sender returns error",
			args: []string{
				"dtag",
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
				"dtag",
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, math.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdAcceptDTagTransfer()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdRefuseDTagTransfer() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid sender returns error",
			args: []string{
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "same user and sender returns error",
			args: []string{
				val.Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, math.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdRefuseDTagTransfer()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}
