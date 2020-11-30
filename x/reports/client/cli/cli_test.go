// +build norace

package cli_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/desmos-labs/desmos/testutil"
	"github.com/desmos-labs/desmos/x/reports/client/cli"
	"github.com/desmos-labs/desmos/x/reports/types"
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
	cfg.NumValidators = 1

	var reportsData types.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &reportsData))

	reportsData.Reports = []types.Report{
		types.NewReport(
			"2b6284dd0361c20022ce366f4355c052165c0c23d7f588da5ac3572d68fda2f2",
			"scam",
			"Test report",
			"cosmos1azqm9kmyxunkx2yt332hmnr8sa3lclhjlg9w5k",
		),
	}

	reportsDataBz, err := cfg.Codec.MarshalJSON(&reportsData)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = reportsDataBz
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

func (s *IntegrationTestSuite) TestCmdQueryPostReports() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryPostReportsResponse
	}{
		{
			name: "missing post id",
			args: []string{
				"not-found",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr:      false,
			expectedOutput: types.QueryPostReportsResponse{Reports: []types.Report{}},
		},
		{
			name: "valid post id",
			args: []string{
				"2b6284dd0361c20022ce366f4355c052165c0c23d7f588da5ac3572d68fda2f2",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryPostReportsResponse{Reports: []types.Report{
				types.NewReport(
					"2b6284dd0361c20022ce366f4355c052165c0c23d7f588da5ac3572d68fda2f2",
					"scam",
					"Test report",
					"cosmos1azqm9kmyxunkx2yt332hmnr8sa3lclhjlg9w5k",
				),
			}},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryPostReports()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryPostReportsResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}

// ___________________________________________________________________________________________________________________

func (s *IntegrationTestSuite) TestCmdReportPost() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid post id",
			expErr: true,
			args:   []string{"1", "scam", "message"},
		},
		{
			name:   "invalid report type",
			expErr: true,
			args:   []string{"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd", "", "message"},
		},
		{
			name:   "valid report",
			expErr: false,
			args: []string{
				"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
				"scam",
				"message",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			respType: &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdReportPost()
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
