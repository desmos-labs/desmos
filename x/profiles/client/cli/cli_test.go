// +build norace

package cli_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/desmos-labs/desmos/testutil"
	"github.com/desmos-labs/desmos/x/profiles/client/cli"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := testutil.DefaultConfig()
	genesisState := cfg.GenesisState
	cfg.NumValidators = 2

	var profilesData types.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &profilesData))

	profilesData.Profiles = []types.Profile{
		types.NewProfile(
			"dtag",
			"moniker",
			"bio",
			types.Pictures{},
			time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
		),
	}
	profilesData.DtagTransferRequests = []types.DTagTransferRequest{
		types.NewDTagTransferRequest(
			"dtag",
			"cosmos122u6u9gpdr2rp552fkkvlgyecjlmtqhkascl5a",
			"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
		),
	}
	profilesData.Params = types.DefaultParams()

	profilesDataBz, err := cfg.Codec.MarshalJSON(&profilesData)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = profilesDataBz
	cfg.GenesisState = genesisState

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

// ___________________________________________________________________________________________________________________

func (s *IntegrationTestSuite) TestCmdQueryProfile() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryProfileResponse
	}{
		{
			name: "non existing profile",
			args: []string{
				s.network.Validators[1].Address.String(),
			},
			expectErr: true,
		},
		{
			name: "existing profile is returned properly",
			args: []string{
				"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryProfileResponse{
				Profile: types.NewProfile(
					"dtag",
					"moniker",
					"bio",
					types.Pictures{},
					time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryProfile()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryProfileResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryDTagRequests() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryDTagTransfersResponse
	}{
		{
			name: "empty slice is returned properly",
			args: []string{
				"cosmos1nqwf7chwfywdw2379sxmwlcgcfvvy86t6mpunz",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryDTagTransfersResponse{
				Requests: []types.DTagTransferRequest{},
			},
		},
		{
			name: "existing requests are returned properly",
			args: []string{
				"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryDTagTransfersResponse{
				Requests: []types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"cosmos122u6u9gpdr2rp552fkkvlgyecjlmtqhkascl5a",
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryDTagRequests()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryDTagTransfersResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryParams() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryParamsResponse
	}{
		{
			name:      "existing params are returned properly",
			args:      []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			expectErr: false,
			expectedOutput: types.QueryParamsResponse{
				Params: types.DefaultParams(),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryParams()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryParamsResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}

// ___________________________________________________________________________________________________________________

func (s *IntegrationTestSuite) TestCmdSaveProfile() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "invalid dtag returns error",
			args: []string{
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "invalid creator returns error",
			args: []string{
				"dtag",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, ""),
			},
			expErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
				"dtag",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			expErr:   false,
			respType: &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdSaveProfile()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdDeleteProfile() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid user returns error",
			args:   []string{fmt.Sprintf("--%s=%s", flags.FlagFrom, "")},
			expErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			expErr:   false,
			respType: &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdDeleteProfile()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdRequestDTagTransfer() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "invalid user address returns error",
			args: []string{
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "same user returns error",
			args: []string{
				val.Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			expErr:   false,
			respType: &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdRequestDTagTransfer()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdCancelDTagTransfer() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "invalid recipient returns error",
			args: []string{
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "same user and recipient returns error",
			args: []string{
				val.Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			expErr:   false,
			respType: &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdCancelDTagTransfer()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdAcceptDTagTransfer() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "invalid new dtag returns error",
			args: []string{
				"",
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "invalid request sender returns error",
			args: []string{
				"dtag",
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
				"dtag",
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			expErr:   false,
			respType: &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdAcceptDTagTransfer()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdRefuseDTagTransfer() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "invalid sender returns error",
			args: []string{
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "same user and sender returns error",
			args: []string{
				val.Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			expErr:   false,
			respType: &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdRefuseDTagTransfer()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}
