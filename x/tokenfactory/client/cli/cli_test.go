//go:build norace
// +build norace

package cli_test

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v5/testutil"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
	"github.com/desmos-labs/desmos/v5/x/tokenfactory/client/cli"
	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
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
				"cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
		},
		nil, nil, nil, nil, nil,
	)

	subspacesDataBz, err := cfg.Codec.MarshalJSON(subspacesGenesis)
	s.Require().NoError(err)
	genesisState[subspacestypes.ModuleName] = subspacesDataBz

	tokenfactoryGenesis := &types.GenesisState{
		Params: types.DefaultParams(),
		FactoryDenoms: []types.GenesisDenom{
			{
				Denom: "factory/cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9/uminttoken",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				},
			},
			{
				Denom: "factory/cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9/utesttoken",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9",
				},
			},
		},
	}

	tokenfactoryDataBz, err := cfg.Codec.MarshalJSON(tokenfactoryGenesis)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = tokenfactoryDataBz

	cfg.GenesisState = genesisState

	s.cfg = cfg
	s.network, err = network.New(s.T(), s.T().TempDir(), cfg)
	s.Require().NoError(err)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

// --------------------------------------------------------------------------------------------------------------------

func (s *IntegrationTestSuite) TestCmdQueryParams() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryParamsResponse
	}{
		{
			name: "params is returned properly",
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryParamsResponse{
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

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryParamsResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Params, response.Params)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQuerySubspaceDenoms() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QuerySubspaceDenomsResponse
	}{
		{
			name: "invalid subspace id returns error",
			args: []string{
				"X",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "denoms are returned properly",
			args: []string{
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QuerySubspaceDenomsResponse{
				Denoms: []string{
					"factory/cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9/uminttoken",
					"factory/cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9/utesttoken",
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQuerySubspaceDenoms()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QuerySubspaceDenomsResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Denoms, response.Denoms)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (s *IntegrationTestSuite) TestCmdCreateDenom() {
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
				"X", "utrytoken",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid subdenom returns error",
			args: []string{
				"1", "",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			args: []string{
				"1", "utrytoken",
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
			cmd := cli.GetCmdCreateDenom()
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

func (s *IntegrationTestSuite) TestCmdMint() {
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
				"X",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid amount returns error",
			args: []string{
				"1", "X",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid sender returns no error",
			args: []string{
				"1", "10factory/cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9/uminttoken",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			args: []string{
				"1", "10factory/cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9/uminttoken",
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
			cmd := cli.GetCmdMint()
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

func (s *IntegrationTestSuite) TestCmdBurn() {
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
				"X",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid amount returns no error",
			args: []string{
				"1", "X",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid sender returns no error",
			args: []string{
				"1", "10factory/cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9/uminttoken",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			args: []string{
				"1", "10factory/cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9/uminttoken",
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
			cmd := cli.GetCmdBurn()
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

func (s *IntegrationTestSuite) TestCmdSetDenomMetadata() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		setupFile func() string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid subspace id returns error",
			args: []string{
				"X",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid path returns error",
			args: []string{
				"1", "X",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid metadata schema return error",
			setupFile: func() string {
				os.CreateTemp(os.TempDir(), "metadata.json")
				return path.Join(os.TempDir(), "metadata.json")
			},
			args: []string{
				"1", path.Join(os.TempDir(), "metadata.json"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			setupFile: func() string {
				file, _ := os.CreateTemp(os.TempDir(), "metadata.json")

				bz := s.cfg.Codec.MustMarshalJSON(&banktypes.Metadata{
					Name:        "Mint Token",
					Symbol:      "MTK",
					Description: "The custom token of the test subspace.",
					DenomUnits: []*banktypes.DenomUnit{
						{Denom: "factory/cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9/uminttoken", Exponent: uint32(0), Aliases: nil},
						{Denom: "minttoken", Exponent: uint32(6), Aliases: []string{"minttoken"}},
					},
					Base:    "factory/cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9/uminttoken",
					Display: "minttoken",
				})
				file.Write(bz)
				return path.Join(os.TempDir(), "metadata.json")
			},
			args: []string{
				"1", path.Join(os.TempDir(), "metadata.json"),
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
			cmd := cli.GetCmdCreateDenom()
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

func (s *IntegrationTestSuite) TestCmdDraftDenomMetadata() {
	val := s.network.Validators[0]

	s.Run("draft denom metadata properly", func() {
		outputPath := path.Join(os.TempDir(), "metadata.json")
		cmd := cli.GetCmdDraftDenomMetadata()
		clientCtx := val.ClientCtx

		_, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{fmt.Sprintf("--%s=%s", cli.FlagOutputPath, outputPath)})
		s.Require().NoError(err)

		out, err := os.ReadFile(outputPath)
		s.Require().NoError(err)

		var metadata banktypes.Metadata
		err = clientCtx.Codec.UnmarshalJSON(out, &metadata)
		s.Require().NoError(err)
		s.Require().Equal(banktypes.Metadata{DenomUnits: []*banktypes.DenomUnit{}}, metadata)
	})
}
