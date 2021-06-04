package cli_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/testutil"
	"github.com/desmos-labs/desmos/x/staging/subspaces/client/cli"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"testing"
	"time"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func TestIntegrationTestSuite(t *testing.T) {
	// TODO re add this when out of staging
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := testutil.DefaultConfig()
	genesisState := cfg.GenesisState
	cfg.NumValidators = 2

	var subspacesData types.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &subspacesData))

	date, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	s.Require().NoError(err)

	subspacesData.Subspaces = []types.Subspace{
		{
			ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			Name:            "test",
			Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			CreationTime:    date,
			Type:            types.Open,
			Admins:          []string{},
			BannedUsers:     []string{},
			RegisteredUsers: []string{},
		},
	}

	subspacesDataBz, err := cfg.Codec.MarshalJSON(&subspacesData)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = subspacesDataBz
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

func (s *IntegrationTestSuite) TestCmdQuerySubspace() {
	val := s.network.Validators[0]
	date, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	s.Require().NoError(err)

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QuerySubspaceResponse
	}{
		{
			name:      "non existing subspace",
			args:      []string{"subspace_id"},
			expectErr: true,
		},
		{
			name: "existing subspace is returned correctly",
			args: []string{
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QuerySubspaceResponse{
				Subspace: types.Subspace{
					ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					Name:            "test",
					Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					CreationTime:    date,
					Type:            types.Open,
					Admins:          []string{},
					BannedUsers:     []string{},
					RegisteredUsers: []string{},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQuerySubspace()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QuerySubspaceResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Subspace, tc.expectedOutput.Subspace)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQuerySubspaces() {
	val := s.network.Validators[0]
	date, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	s.Require().NoError(err)

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QuerySubspacesResponse
	}{
		{
			name: "subspaces is returned correctly",
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QuerySubspacesResponse{
				Subspaces: []types.Subspace{
					{
						ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						Name:            "test",
						Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						CreationTime:    date,
						Type:            types.Open,
						Admins:          []string{},
						BannedUsers:     []string{},
						RegisteredUsers: []string{},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQuerySubspaces()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QuerySubspacesResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput.Subspaces, response.Subspaces)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdCreateSubspace() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid subspace ID returns error",
			args:   []string{"subspace"},
			expErr: true,
		},
		{
			name: "invalid name returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"",
			},
			expErr: true,
		},
		{
			name: "invalid subspace type returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"mooncake",
				fmt.Sprintf("--%s=%s", cli.FlagSubspaceType, "inv"),
			},
			expErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"mooncake",
				fmt.Sprintf("--%s=%s", cli.FlagSubspaceType, "close"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			expErr:   false,
			respType: &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error (with --subspace-type flag)",
			args: []string{
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"mooncake",
				fmt.Sprintf("--%s=%s", cli.FlagSubspaceType, "open"),
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
			cmd := cli.GetCmdCreateSubspace()
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

func (s *IntegrationTestSuite) TestCmdEditSubspace() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid subspace ID returns error",
			args:   []string{"subspace"},
			expErr: true,
		},
		{
			name: "invalid owner flag returns error",
			args: []string{
				"subspace",
				fmt.Sprintf("--%s=%s", cli.FlagOwner, "abd"),
			},
			expErr: true,
		},
		{
			name: "invalid subspace type returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"mooncake",
				fmt.Sprintf("--%s=%s", cli.FlagSubspaceType, "inv"),
			},
			expErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", cli.FlagName, "new"),
				fmt.Sprintf("--%s=%s", cli.FlagSubspaceType, "open"),
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
			cmd := cli.GetCmdEditSubspace()
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

func (s *IntegrationTestSuite) TestCmdAddAdmin() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid subspace ID returns error",
			args:   []string{"subspace"},
			expErr: true,
		},
		{
			name: "invalid address returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"",
			},
			expErr: true,
		},
		{
			name: "valid data returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
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
			cmd := cli.GetCmdAddAdmin()
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

func (s *IntegrationTestSuite) TestCmdRemoveAdmin() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid subspace ID returns error",
			args:   []string{"subspace"},
			expErr: true,
		},
		{
			name: "invalid address returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"",
			},
			expErr: true,
		},
		{
			name: "valid data returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
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
			cmd := cli.GetCmdRemoveAdmin()
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

func (s *IntegrationTestSuite) TestCmdRegisterUser() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid subspace ID returns error",
			args:   []string{"subspace"},
			expErr: true,
		},
		{
			name: "invalid address returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"",
			},
			expErr: true,
		},
		{
			name: "valid data returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
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
			cmd := cli.GetCmdRegisterUser()
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

func (s *IntegrationTestSuite) TestCmdUnregisterUser() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid subspace ID returns error",
			args:   []string{"subspace"},
			expErr: true,
		},
		{
			name: "invalid address returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"",
			},
			expErr: true,
		},
		{
			name: "valid data returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
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
			cmd := cli.GetCmdUnregisterUser()
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

func (s *IntegrationTestSuite) TestCmdBanUser() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid subspace ID returns error",
			args:   []string{"subspace"},
			expErr: true,
		},
		{
			name: "invalid address returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"",
			},
			expErr: true,
		},
		{
			name: "valid data returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
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
			cmd := cli.GetCmdBanUser()
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

func (s *IntegrationTestSuite) TestCmdUnbanUser() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid subspace ID returns error",
			args:   []string{"subspace"},
			expErr: true,
		},
		{
			name: "invalid address returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"",
			},
			expErr: true,
		},
		{
			name: "valid data returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
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
			cmd := cli.GetCmdUnbanUser()
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
