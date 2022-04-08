package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/desmos-labs/desmos/v3/testutil"
	"github.com/desmos-labs/desmos/v3/x/supply/client/cli"
	"github.com/desmos-labs/desmos/v3/x/supply/types"
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

	var authData authtypes.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[authtypes.ModuleName], &authData))

	var bankData banktypes.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[banktypes.ModuleName], &bankData))

	var stakingData stakingtypes.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[stakingtypes.ModuleName], &stakingData))

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestCmdQueryTotalSupply() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		shouldErr      bool
		expectedOutput types.QueryTotalSupplyResponse
	}{
		{
			name: "invalid denom returns error",
			args: []string{
				"",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "total supply returned correctly without divider_exponent conversion applied",
			args: []string{
				"stake",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr:      false,
			expectedOutput: types.QueryTotalSupplyResponse{TotalSupply: sdk.NewInt(1000000020)},
		},
		{
			name: "total supply returned correctly with divider_exponent conversion applied",
			args: []string{
				"stake",
				"2",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr:      false,
			expectedOutput: types.QueryTotalSupplyResponse{TotalSupply: sdk.NewInt(10000000)},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryTotalSupply()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryTotalSupplyResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput.TotalSupply, response.TotalSupply)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryCirculatingSupply() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		shouldErr      bool
		expectedOutput types.QueryCirculatingSupplyResponse
	}{
		{
			name: "invalid denom returns error",
			args: []string{
				"",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "circulating supply returned correctly without divider_exponent conversion applied",
			args: []string{
				"stake",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr:      false,
			expectedOutput: types.QueryCirculatingSupplyResponse{CirculatingSupply: sdk.NewInt(1000000020)},
		},
		{
			name: "circulating supply returned correctly with divider_exponent conversion applied",
			args: []string{
				"stake",
				"2",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr:      false,
			expectedOutput: types.QueryCirculatingSupplyResponse{CirculatingSupply: sdk.NewInt(10000000)},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryCirculatingSupply()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryCirculatingSupplyResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput.CirculatingSupply, response.CirculatingSupply)
			}
		})
	}
}
