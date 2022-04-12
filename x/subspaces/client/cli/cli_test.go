package cli_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/gogo/protobuf/proto"

	"github.com/desmos-labs/desmos/v3/x/subspaces/client/cli"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/testutil"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
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

	// Initialize the module genesis data
	genesis := types.NewGenesisState(
		3,
		[]types.GenesisSubspace{
			types.NewGenesisSubspace(
				types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
				1,
			),
			types.NewGenesisSubspace(
				types.NewSubspace(
					2,
					"Another test subspace",
					"This is another test subspace",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
				),
				3,
			),
			types.NewGenesisSubspace(
				types.NewSubspace(
					3,
					"Subspace to delete",
					"This is a test subspace that will be deleted",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
				),
				1,
			),
		},
		[]types.ACLEntry{
			types.NewACLEntry(1, "cosmos1xw69y2z3yf00rgfnly99628gn5c0x7fryyfv5e", types.PermissionWrite),
			types.NewACLEntry(2, "cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd", types.PermissionManageGroups),
		},
		[]types.UserGroup{
			types.NewUserGroup(1, 1, "Test group", "", types.PermissionWrite),
			types.NewUserGroup(2, 1, "Another test group", "", types.PermissionManageGroups),
			types.NewUserGroup(2, 2, "Third group", "", types.PermissionWrite),
		},
		[]types.UserGroupMembersEntry{
			types.NewUserGroupMembersEntry(1, 1, []string{
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			}),
			types.NewUserGroupMembersEntry(2, 1, []string{
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			}),
		},
	)

	// Store the genesis data
	subspacesDataBz, err := cfg.Codec.MarshalJSON(genesis)
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

// --------------------------------------------------------------------------------------------------------------------

func (s *IntegrationTestSuite) TestCmdQuerySubspace() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QuerySubspaceResponse
	}{
		{
			name: "non existing subspace returns error",
			args: []string{
				"10",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "existing subspace is returned correctly",
			args: []string{
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QuerySubspaceResponse{
				Subspace: types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQuerySubspace()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QuerySubspaceResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Subspace, tc.expResponse.Subspace)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQuerySubspaces() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QuerySubspacesResponse
	}{
		{
			name: "subspaces are returned correctly",
			args: []string{
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QuerySubspacesResponse{
				Subspaces: []types.Subspace{
					types.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
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

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QuerySubspacesResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Subspaces, response.Subspaces)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryUserGroups() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryUserGroupsResponse
	}{
		{
			name: "user groups are returned correctly",
			args: []string{
				"2",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryUserGroupsResponse{
				Groups: []types.UserGroup{
					types.DefaultUserGroup(2),
					types.NewUserGroup(2, 1, "Another test group", "", types.PermissionManageGroups),
					types.NewUserGroup(2, 2, "Third group", "", types.PermissionWrite),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryUserGroups()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryUserGroupsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Groups, response.Groups)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryUserGroup() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryUserGroupResponse
	}{
		{
			name: "user group is returned correctly",
			args: []string{
				"2",
				"0",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryUserGroupResponse{
				Group: types.DefaultUserGroup(2),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryUserGroup()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryUserGroupResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Group, response.Group)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryUserGroupMembers() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryUserGroupMembersResponse
	}{
		{
			name:      "subspace not found returns error",
			args:      []string{"10", "1"},
			shouldErr: true,
		},
		{
			name:      "group not found returns error",
			args:      []string{"1", "10"},
			shouldErr: true,
		},
		{
			name: "members are returned correctly",
			args: []string{
				"2", "1",
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryUserGroupMembersResponse{
				Members: []string{
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryUserGroupMembers()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryUserGroupMembersResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Members, response.Members)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryUserPermissions() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryUserPermissionsResponse
	}{
		{
			name: "subspace not found returns error",
			args: []string{
				"11", "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "user permissions are returned correctly",
			args: []string{
				"2", "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryUserPermissionsResponse{
				Permissions: types.PermissionManageGroups,
				Details: []types.PermissionDetail{
					types.NewPermissionDetailGroup(0, types.PermissionNothing),
					types.NewPermissionDetailGroup(1, types.PermissionManageGroups),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryUserPermissions()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryUserPermissionsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Permissions, response.Permissions)
				s.Require().Equal(tc.expResponse.Details, response.Details)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdCreateSubspace() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid name returns error",
			args: []string{
				"",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"Test subspace",
				fmt.Sprintf("--%s=%s", cli.FlagDescription, "This is a test subspace"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error with custom treasury and owner",
			args: []string{
				"Another test subspace",
				fmt.Sprintf("--%s=%s", cli.FlagDescription, "Another test subspace"),
				fmt.Sprintf("--%s=%s", cli.FlagTreasury, "cosmos1lqjd4p6uxvsvus6kf3gf00uz7luj4jcxpyahul"),
				fmt.Sprintf("--%s=%s", cli.FlagOwner, "cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u"),
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
			cmd := cli.GetCmdCreateSubspace()
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

func (s *IntegrationTestSuite) TestCmdEditSubspace() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"subspace"},
			shouldErr: true,
		},
		{
			name: "invalid name returns error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagName, ""),
			},
			shouldErr: true,
		},
		{
			name: "invalid owner flag returns error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagOwner, "abd"),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"2",
				fmt.Sprintf("--%s=%s", cli.FlagName, "Edited name"),
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
			cmd := cli.GetCmdEditSubspace()
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

func (s *IntegrationTestSuite) TestCmdDeleteSubspace() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"subspace"},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"3",
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
			cmd := cli.GetCmdDeleteSubspace()
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

func (s *IntegrationTestSuite) TestCmdCreateUserGroup() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"id", "testing-group"},
			shouldErr: true,
		},
		{
			name:      "invalid name returns error",
			args:      []string{"1", ""},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"2", "testing-group",
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
			cmd := cli.GetCmdCreateUserGroup()
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

func (s *IntegrationTestSuite) TestCmdEditUserGroup() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"0", "1"},
			shouldErr: true,
		},
		{
			name:      "invalid group id returns error",
			args:      []string{"1", "g"},
			shouldErr: true,
		},
		{
			name: "valid data returns no error - group = 0",
			args: []string{
				"1", "0",
				fmt.Sprintf("--%s=%s", flags.FlagName, "This is my new group name"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error - group > 0",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%s", flags.FlagName, "This is my new group name"),
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
			cmd := cli.GetCmdEditUserGroup()
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

func (s *IntegrationTestSuite) TestCmdSetUserGroupPermissions() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"0", "1"},
			shouldErr: true,
		},
		{
			name:      "invalid group id returns error",
			args:      []string{"1", "g"},
			shouldErr: true,
		},
		{
			name: "valid data returns no error - group id = 0",
			args: []string{
				"1", "0", types.SerializePermission(types.PermissionWrite),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error - group id > 0",
			args: []string{
				"1", "1", types.SerializePermission(types.PermissionWrite),
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
			cmd := cli.GetCmdSetUserGroupPermissions()
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

func (s *IntegrationTestSuite) TestCmdDeleteUserGroup() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"subspace-id", "testing-group"},
			shouldErr: true,
		},
		{
			name:      "invalid name returns error",
			args:      []string{"1", ""},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"2", "1",
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
			cmd := cli.GetCmdDeleteUserGroup()
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

func (s *IntegrationTestSuite) TestCmdAddUserToUserGroup() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"subspace-id", "testing-group"},
			shouldErr: true,
		},
		{
			name:      "invalid name returns error",
			args:      []string{"1", ""},
			shouldErr: true,
		},
		{
			name:      "invalid user returns error",
			args:      []string{"1", "1", ""},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"2", "1", "cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
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
			cmd := cli.GetCmdAddUserToUserGroup()
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

func (s *IntegrationTestSuite) TestCmdRemoveUserFromUserGroup() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"subspace-id", "testing-group"},
			shouldErr: true,
		},
		{
			name:      "invalid name returns error",
			args:      []string{"1", ""},
			shouldErr: true,
		},
		{
			name:      "invalid user returns error",
			args:      []string{"1", "group", ""},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"2", "1", "cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
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
			cmd := cli.GetCmdRemoveUserFromUserGroup()
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

func (s *IntegrationTestSuite) TestCmdSetPermissions() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"id", "group", "Write"},
			shouldErr: true,
		},
		{
			name:      "invalid target returns error",
			args:      []string{"1", "", "Write"},
			shouldErr: true,
		},
		{
			name:      "invalid permission returns error",
			args:      []string{"1", "group", "NonExistingPermission"},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "cosmos1xw69y2z3yf00rgfnly99628gn5c0x7fryyfv5e", "Write",
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
			cmd := cli.GetCmdSetUserPermissions()
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
