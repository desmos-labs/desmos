// +build norace

package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/testutil"
	"github.com/desmos-labs/desmos/x/relationships/client/cli"
	"github.com/desmos-labs/desmos/x/relationships/types"
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
	cfg.NumValidators = 3

	var relationshipsData types.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &relationshipsData))

	relationshipsData.Blocks = []types.UserBlock{
		types.NewUserBlock(
			"cosmos1azqm9kmyxunkx2yt332hmnr8sa3lclhjlg9w5k",
			"cosmos1zs70glquczqgt83g03jnvcqppu4jjj8yjxwlvh",
			"Test block",
			"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		),
	}
	relationshipsData.Relationships = []types.Relationship{
		types.NewRelationship(
			"cosmos122u6u9gpdr2rp552fkkvlgyecjlmtqhkascl5a",
			"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
			"60303ae22b998861bce3b28f33eec1be758a213c86c93c076dbe9f558c11c752",
		),
	}

	relationshipsDataBz, err := cfg.Codec.MarshalJSON(&relationshipsData)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = relationshipsDataBz
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

func (s *IntegrationTestSuite) TestCmdQueryUserRelationships() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryUserRelationshipsResponse
	}{
		{
			name: "empty array is returned properly",
			args: []string{
				s.network.Validators[2].Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryUserRelationshipsResponse{
				User:          s.network.Validators[2].Address.String(),
				Relationships: []types.Relationship{},
			},
		},
		{
			name: "existing relationship is returned properly",
			args: []string{
				"cosmos122u6u9gpdr2rp552fkkvlgyecjlmtqhkascl5a",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryUserRelationshipsResponse{
				User: "cosmos122u6u9gpdr2rp552fkkvlgyecjlmtqhkascl5a",
				Relationships: []types.Relationship{
					types.NewRelationship(
						"cosmos122u6u9gpdr2rp552fkkvlgyecjlmtqhkascl5a",
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						"60303ae22b998861bce3b28f33eec1be758a213c86c93c076dbe9f558c11c752",
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryUserRelationships()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryUserRelationshipsResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryUserBlocks() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryUserBlocksResponse
	}{
		{
			name: "empty slice is returned properly",
			args: []string{
				s.network.Validators[2].Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryUserBlocksResponse{
				Blocks: []types.UserBlock{},
			},
		},
		{
			name: "existing user blocks are returned properly",
			args: []string{
				"cosmos1azqm9kmyxunkx2yt332hmnr8sa3lclhjlg9w5k",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryUserBlocksResponse{
				Blocks: []types.UserBlock{
					types.NewUserBlock(
						"cosmos1azqm9kmyxunkx2yt332hmnr8sa3lclhjlg9w5k",
						"cosmos1zs70glquczqgt83g03jnvcqppu4jjj8yjxwlvh",
						"Test block",
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryUserBlocks()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryUserBlocksResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}

// ___________________________________________________________________________________________________________________

func (s *IntegrationTestSuite) TestCmdCreateRelationship() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "invalid subspace returns error",
			args: []string{
				val.Address.String(),
				"subspace",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "invalid blocked user returns error",
			args: []string{
				"address",
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "same user and counterparty returns error",
			args: []string{
				val.Address.String(),
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "valid parameters works properly",
			args: []string{
				s.network.Validators[1].Address.String(),
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
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
			cmd := cli.GetCmdCreateRelationship()
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

func (s *IntegrationTestSuite) TestCmdDeleteRelationship() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "invalid receiver returns error",
			args: []string{
				"receiver",
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "invalid subspace returns error",
			args: []string{
				s.network.Validators[1].Address.String(),
				"subspace",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "same user and counterparty returns error",
			args: []string{
				val.Address.String(),
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "valid request executes properly",
			args: []string{
				s.network.Validators[1].Address.String(),
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
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
			cmd := cli.GetCmdDeleteRelationship()
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

func (s *IntegrationTestSuite) TestCmdBlockUser() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "invalid blocked address returns error",
			args: []string{
				"address",
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "invalid subspace returns error",
			args: []string{
				s.network.Validators[1].Address.String(),
				"subspace",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "same blocker and blocked returns error",
			args: []string{
				val.Address.String(),
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "valid request works properly without reason",
			args: []string{
				s.network.Validators[1].Address.String(),
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			expErr:   false,
			respType: &sdk.TxResponse{},
		},
		{
			name: "valid request works properly with reason",
			args: []string{
				s.network.Validators[1].Address.String(),
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"Blocking reason",
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
			cmd := cli.GetCmdBlockUser()
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

func (s *IntegrationTestSuite) TestCmdUnblockUser() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "invalid blocked address returns error",
			args: []string{
				"address",
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "invalid subspace returns error",
			args: []string{
				s.network.Validators[1].Address.String(),
				"subspace",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "same blocker and blocked returns error",
			args: []string{
				val.Address.String(),
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "valid request works properly without reason",
			args: []string{
				s.network.Validators[1].Address.String(),
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			expErr:   false,
			respType: &sdk.TxResponse{},
		},
		{
			name: "valid request works properly with reason",
			args: []string{
				s.network.Validators[1].Address.String(),
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
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
			cmd := cli.GetCmdUnblockUser()
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
