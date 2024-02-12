//go:build norace
// +build norace

package cli_test

import (
	"fmt"
	"testing"
	"time"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	"github.com/desmos-labs/desmos/v7/testutil"
	profilestypes "github.com/desmos-labs/desmos/v7/x/profiles/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v7/x/relationships/client/cli"
	"github.com/desmos-labs/desmos/v7/x/relationships/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
	keyBase keyring.Keyring
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := testutil.DefaultConfig()
	genesisState := cfg.GenesisState
	cfg.NumValidators = 2

	// Store a profile account inside the auth genesis data
	var authData authtypes.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[authtypes.ModuleName], &authData))

	addr, err := sdk.AccAddressFromBech32("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs")
	s.Require().NoError(err)

	profile, err := profilestypes.NewProfile(
		"TestDTag",
		"Test",
		"",
		profilestypes.Pictures{},
		time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
		authtypes.NewBaseAccountWithAddress(addr),
	)
	s.Require().NoError(err)
	profileAny, err := codectypes.NewAnyWithValue(profile)
	s.Require().NoError(err)
	authData.Accounts = append(authData.Accounts, profileAny)

	authDataBz, err := cfg.Codec.MarshalJSON(&authData)
	s.Require().NoError(err)

	genesisState[authtypes.ModuleName] = authDataBz

	// Store the profiles genesis state
	var relationshipsGenesis types.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &relationshipsGenesis))

	relationshipsGenesis.Blocks = []types.UserBlock{
		types.NewUserBlock(
			addr.String(),
			"cosmos1zs70glquczqgt83g03jnvcqppu4jjj8yjxwlvh",
			"Test block",
			2,
		),
	}
	relationshipsGenesis.Relationships = []types.Relationship{
		types.NewRelationship(
			addr.String(),
			"cosmos1zs70glquczqgt83g03jnvcqppu4jjj8yjxwlvh",
			2,
		),
	}

	relationshipsGenesisBz, err := cfg.Codec.MarshalJSON(&relationshipsGenesis)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = relationshipsGenesisBz
	cfg.GenesisState = genesisState

	s.cfg = cfg
	s.network, err = network.New(s.T(), s.T().TempDir(), cfg)
	s.Require().NoError(err)
	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestCmdQueryRelationships() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		shouldErr      bool
		expectedOutput types.QueryRelationshipsResponse
	}{
		{
			name: "empty array is returned properly",
			args: []string{
				"10",
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryRelationshipsResponse{
				Relationships: []types.Relationship{},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "existing relationships of the given user are returned properly",
			args: []string{
				"2",
				"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryRelationshipsResponse{
				Relationships: []types.Relationship{
					types.NewRelationship(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						"cosmos1zs70glquczqgt83g03jnvcqppu4jjj8yjxwlvh",
						2,
					),
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "existing relationships are returned properly",
			args: []string{
				"2",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryRelationshipsResponse{
				Relationships: []types.Relationship{
					types.NewRelationship(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						"cosmos1zs70glquczqgt83g03jnvcqppu4jjj8yjxwlvh",
						2,
					),
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryRelationships()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryRelationshipsResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryBlocks() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		shouldErr      bool
		expectedOutput types.QueryBlocksResponse
	}{
		{
			name: "empty slice is returned properly",
			args: []string{
				"10",
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryBlocksResponse{
				Blocks: []types.UserBlock{},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "existing user blocks for a given user are returned properly",
			args: []string{
				"2",
				"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryBlocksResponse{
				Blocks: []types.UserBlock{
					types.NewUserBlock(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						"cosmos1zs70glquczqgt83g03jnvcqppu4jjj8yjxwlvh",
						"Test block",
						2,
					),
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "existing user blocks are returned properly",
			args: []string{
				"2",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryBlocksResponse{
				Blocks: []types.UserBlock{
					types.NewUserBlock(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						"cosmos1zs70glquczqgt83g03jnvcqppu4jjj8yjxwlvh",
						"Test block",
						2,
					),
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryBlocks()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryBlocksResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdCreateRelationship() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid subspace returns error",
			args: []string{
				val.Address.String(),
				"0",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid blocked user returns error",
			args: []string{
				"address",
				"1",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "same user and counterparty returns error",
			args: []string{
				val.Address.String(),
				"1",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "valid parameters works properly",
			args: []string{
				s.network.Validators[1].Address.String(),
				"1",
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
			cmd := cli.GetCmdCreateRelationship()
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

func (s *IntegrationTestSuite) TestCmdDeleteRelationship() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid receiver returns error",
			args: []string{
				"receiver",
				"1",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid subspace returns error",
			args: []string{
				s.network.Validators[1].Address.String(),
				"0",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "same user and counterparty returns error",
			args: []string{
				val.Address.String(),
				"1",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "valid request executes properly",
			args: []string{
				s.network.Validators[1].Address.String(),
				"1",
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
			cmd := cli.GetCmdDeleteRelationship()
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

func (s *IntegrationTestSuite) TestCmdBlockUser() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid blocked address returns error",
			args: []string{
				"address",
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid subspace returns error",
			args: []string{
				s.network.Validators[1].Address.String(),
				"subspace",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "same blocker and blocked returns error",
			args: []string{
				val.Address.String(),
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "valid request works properly without reason",
			args: []string{
				s.network.Validators[1].Address.String(),
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid request works properly with reason",
			args: []string{
				s.network.Validators[1].Address.String(),
				"",
				"Blocking reason",
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
			cmd := cli.GetCmdBlockUser()
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

func (s *IntegrationTestSuite) TestCmdUnblockUser() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid blocked address returns error",
			args: []string{
				"address",
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid subspace returns error",
			args: []string{
				s.network.Validators[1].Address.String(),
				"subspace",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "same blocker and blocked returns error",
			args: []string{
				val.Address.String(),
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			shouldErr: true,
		},
		{
			name: "valid request works properly without reason",
			args: []string{
				s.network.Validators[1].Address.String(),
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid request works properly with reason",
			args: []string{
				s.network.Validators[1].Address.String(),
				"",
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
			cmd := cli.GetCmdUnblockUser()
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
