//go:build norace
// +build norace

package cli_test

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authzcli "github.com/cosmos/cosmos-sdk/x/authz/client/cli"
	"github.com/gogo/protobuf/proto"

	"github.com/desmos-labs/desmos/v5/x/subspaces/client/cli"
)

func (s *IntegrationTestSuite) TestCmdRevokeGrantAuthorization() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid subspace id returns error",
			args: []string{
				"x",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
			},
			shouldErr: true,
		},
		{
			name: "invalid expiration returns error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "x"),
			},
			shouldErr: true,
		},
		{
			name: "invalid msg returns error",
			args: []string{
				"0",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "10000"),
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "10000"),
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdGrantTreasuryAuthorization()
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

func (s *IntegrationTestSuite) TestCmdRevokeGrantAuthorization__SendAuthorization() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid spend limit returns error - invalid coins",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "x"),
			},
			shouldErr: true,
		},
		{
			name: "invalid spend limit returns error - negative coins",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "-1stake"),
			},
			shouldErr: true,
		},
		{
			name: "invalid msg returns error",
			args: []string{
				"0",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "10000"),
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "10000"),
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error - without expiration",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdGrantTreasuryAuthorization()
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

func (s *IntegrationTestSuite) TestCmdRevokeGrantAuthorization__GenericAuthorization() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "valid data returns no error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"generic",
				fmt.Sprintf("--%s=%s", authzcli.FlagMsgType, "/cosmos.bank.v1beta1.MsgSend"),
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "10000"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdGrantTreasuryAuthorization()
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

func (s *IntegrationTestSuite) TestCmdRevokeGrantAuthorization__StakeAuthorization() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid spend limit returns error - invalid coins",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "x"),
			},
			shouldErr: true,
		},
		{
			name: "invalid spend limit returns error - negative coins",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "-1stake"),
			},
			shouldErr: true,
		},
		{
			name: "invalid allowed validators returns error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", authzcli.FlagAllowedValidators, "x"),
			},
			shouldErr: true,
		},
		{
			name: "invalid deny validators returns error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", authzcli.FlagDenyValidators, "x"),
			},
			shouldErr: true,
		},
		{
			name: "empty deny and allowed validators returns error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error - delegate authorization",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", authzcli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error - unbond authorization",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"unbond",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", authzcli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error - redelegate authorization",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"redelegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", authzcli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdGrantTreasuryAuthorization()
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

func (s *IntegrationTestSuite) TestCmdRevokeTreasuryAuthorization() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid subspace id returns error",
			args: []string{
				"x",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"/cosmos.bank.v1beta1.MsgSend",
			},
			shouldErr: true,
		},
		{
			name: "invalid msg returns error",
			args: []string{
				"0",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"/cosmos.bank.v1beta1.MsgSend",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"/cosmos.bank.v1beta1.MsgSend",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdRevokeTreasuryAuthorization()
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
