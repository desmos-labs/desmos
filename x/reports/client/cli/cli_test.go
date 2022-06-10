//go:build norace
// +build norace

package cli_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"
	"github.com/desmos-labs/desmos/v3/x/reports/client/cli"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v3/testutil"
	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := testutil.DefaultConfig()
	genesisState := cfg.GenesisState
	cfg.NumValidators = 2

	// Initialize the subspaces module genesis state
	subspacesGenesis := subspacestypes.NewGenesisState(
		2,
		[]subspacestypes.SubspaceData{
			subspacestypes.NewSubspaceData(1, 1, 1),
		},
		[]subspacestypes.Subspace{
			subspacestypes.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
		},
		nil, nil, nil, nil,
	)
	subspacesDataBz, err := cfg.Codec.MarshalJSON(subspacesGenesis)
	s.Require().NoError(err)
	genesisState[subspacestypes.ModuleName] = subspacesDataBz

	// Initialize the posts module genesis state
	postsGenesis := poststypes.NewGenesisState(
		[]poststypes.SubspaceDataEntry{
			poststypes.NewSubspaceDataEntry(1, 2),
		},
		[]poststypes.Post{
			poststypes.NewPost(
				1,
				0,
				1,
				"External ID",
				"This is a text",
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				0,
				nil,
				[]poststypes.PostReference{},
				poststypes.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
		},
		[]poststypes.PostDataEntry{
			poststypes.NewPostDataEntry(1, 1, 2),
		},
		nil,
		nil,
		nil,
		poststypes.DefaultParams(),
	)
	postsDataBz, err := cfg.Codec.MarshalJSON(postsGenesis)
	s.Require().NoError(err)
	genesisState[poststypes.ModuleName] = postsDataBz

	// Initialize the module genesis data
	genesis := types.NewGenesisState(
		[]types.SubspaceDataEntry{
			types.NewSubspacesDataEntry(1, 3, 3),
		},
		[]types.Reason{
			types.NewReason(
				1,
				1,
				"Spam",
				"This content is spam",
			),
			types.NewReason(
				1,
				2,
				"Harmful content",
				"This content contains self-harm/suicide",
			),
		},
		[]types.Report{
			types.NewReport(
				1,
				1,
				[]uint32{1},
				"This user is a spammer",
				types.NewUserTarget("cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm"),
				"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			types.NewReport(
				1,
				2,
				[]uint32{1},
				"This content is spam",
				types.NewPostTarget(1),
				"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
		},
		types.NewParams([]types.StandardReason{
			types.NewStandardReason(1, "Spam", "This content is spam/This user is a spammer"),
		}),
	)

	// Store the genesis data
	reportsGenesisBz, err := cfg.Codec.MarshalJSON(genesis)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = reportsGenesisBz
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

// --------------------------------------------------------------------------------------------------------------------

func (s *IntegrationTestSuite) TestCmdQueryUserReports() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryReportsResponse
	}{
		{
			name: "reports are returned correctly",
			args: []string{
				"1", "cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm",
				fmt.Sprintf("--%s=%s", cli.FlagReporter, "cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z"),
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryReportsResponse{
				Reports: []types.Report{
					types.NewReport(
						1,
						1,
						[]uint32{1},
						"This user is a spammer",
						types.NewUserTarget("cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm"),
						"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryUserReports()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryReportsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Reports, tc.expResponse.Reports)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryPostReports() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryReportsResponse
	}{
		{
			name: "reports are returned correctly",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%s", cli.FlagReporter, "cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z"),
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryReportsResponse{
				Reports: []types.Report{
					types.NewReport(
						1,
						2,
						[]uint32{1},
						"This content is spam",
						types.NewPostTarget(1),
						"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryPostReports()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryReportsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Reports, tc.expResponse.Reports)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryReports() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryReportsResponse
	}{
		{
			name: "reports are returned correctly",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryReportsResponse{
				Reports: []types.Report{
					types.NewReport(
						1,
						1,
						[]uint32{1},
						"This user is a spammer",
						types.NewUserTarget("cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm"),
						"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryReports()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryReportsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Reports, tc.expResponse.Reports)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryReasons() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryReasonsResponse
	}{
		{
			name: "reasons are returned correctly",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryReasonsResponse{
				Reasons: []types.Reason{
					types.NewReason(
						1,
						1,
						"Spam",
						"This content is spam",
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryReasons()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryReasonsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Reasons, tc.expResponse.Reasons)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryParams() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryParamsResponse
	}{
		{
			name: "params are returned correctly",
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryParamsResponse{
				Params: types.NewParams([]types.StandardReason{
					types.NewStandardReason(1, "Spam", "This content is spam/This user is a spammer"),
				}),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryParams()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryParamsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Params, tc.expResponse.Params)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdReportUser() {
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
				"", "cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w", "1",
			},
			shouldErr: true,
		},
		{
			name: "invalid user address returns error",
			args: []string{
				"1", "invalid-address", "1",
			},
			shouldErr: true,
		},
		{
			name: "invalid reason id returns error",
			args: []string{
				"1", "cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w", "0",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w", "1,2,3",
				fmt.Sprintf("--%s=%s", cli.FlagMessage, "This is a new report"),
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
			cmd := cli.GetCmdReportUser()
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

func (s *IntegrationTestSuite) TestCmdReportPost() {
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
				"", "1", "cosmos1wprgptc8ktt0eemrn2znpxv8crdxm8tdpkdr7w",
			},
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			args: []string{
				"1", "0", "1",
			},
			shouldErr: true,
		},
		{
			name: "invalid reason id returns error",
			args: []string{
				"1", "1", "0",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1", "1,2,3",
				fmt.Sprintf("--%s=%s", cli.FlagMessage, "This is a new report"),
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
			cmd := cli.GetCmdReportPost()
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

func (s *IntegrationTestSuite) TestCmdDeleteReport() {
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
				"0", "1",
			},
			shouldErr: true,
		},
		{
			name: "invalid report id returns error",
			args: []string{
				"1", "0",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1",
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
			cmd := cli.GetCmdDeleteReport()
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

func (s *IntegrationTestSuite) TestCmdSupportStandardReason() {
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
				"0", "1",
			},
			shouldErr: true,
		},
		{
			name: "invalid reason id returns error",
			args: []string{
				"1", "0",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1",
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
			cmd := cli.GetCmdSupportStandardReason()
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

func (s *IntegrationTestSuite) TGetCmdAddReason() {
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
				"0", "Spam",
			},
			shouldErr: true,
		},
		{
			name: "invalid title returns error",
			args: []string{
				"1", "",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "Spam", "This content is spam",
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
			cmd := cli.GetCmdAddReason()
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
