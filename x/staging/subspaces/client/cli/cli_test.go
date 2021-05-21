package cli_test

import (
	"fmt"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/desmos-labs/desmos/testutil"
	"github.com/desmos-labs/desmos/x/staging/subspaces/client/cli"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
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
			Open:            true,
			Admins:          []string{},
			BlockedUsers:    []string{},
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
					Open:            true,
					Admins:          []string{},
					BlockedUsers:    []string{},
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
			name: "existing subspace is returned correctly",
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
						Open:            true,
						Admins:          []string{},
						BlockedUsers:    []string{},
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
				s.Require().Equal(response.Subspaces, tc.expectedOutput.Subspaces)
			}
		})
	}
}
