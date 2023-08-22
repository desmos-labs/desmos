//go:build norace
// +build norace

package cli_test

import (
	"fmt"
	"testing"
	"time"

	authzcli "github.com/cosmos/cosmos-sdk/x/authz/client/cli"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantcli "github.com/cosmos/cosmos-sdk/x/feegrant/client/cli"
	"github.com/stretchr/testify/suite"

	poststypes "github.com/desmos-labs/desmos/v6/x/posts/types"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/gogo/protobuf/proto"

	"github.com/desmos-labs/desmos/v6/x/subspaces/client/cli"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/testutil"
	"github.com/desmos-labs/desmos/v6/x/subspaces/types"
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
		4,
		[]types.SubspaceData{
			types.NewSubspaceData(1, 3, 2),
			types.NewSubspaceData(2, 1, 3),
			types.NewSubspaceData(3, 1, 1),
		},
		[]types.Subspace{
			types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			types.NewSubspace(
				2,
				"Another test subspace",
				"This is another test subspace",
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
				nil,
			),
			types.NewSubspace(
				3,
				"Subspace to delete",
				"This is a test subspace that will be deleted",
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
				nil,
			),
		},
		[]types.Section{
			types.NewSection(
				1,
				1,
				0,
				"Test section",
				"Test section",
			),
			types.NewSection(
				1,
				2,
				0,
				"Another test section",
				"Another test section",
			),
		},
		[]types.UserPermission{
			types.NewUserPermission(1, 0, "cosmos1xw69y2z3yf00rgfnly99628gn5c0x7fryyfv5e", types.NewPermissions(poststypes.PermissionWrite)),
			types.NewUserPermission(2, 0, "cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd", types.NewPermissions(types.PermissionManageGroups)),
		},
		[]types.UserGroup{
			types.NewUserGroup(1, 0, 1, "Test group", "", types.NewPermissions(poststypes.PermissionWrite)),
			types.NewUserGroup(2, 0, 1, "Another test group", "", types.NewPermissions(types.PermissionManageGroups)),
			types.NewUserGroup(2, 0, 2, "Third group", "", types.NewPermissions(poststypes.PermissionWrite)),
		},
		[]types.UserGroupMemberEntry{
			types.NewUserGroupMemberEntry(1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm"),
			types.NewUserGroupMemberEntry(2, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
			types.NewUserGroupMemberEntry(2, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
		},
		[]types.Grant{
			types.NewGrant(
				1,
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			types.NewGrant(
				1,
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				types.NewUserGrantee("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			types.NewGrant(
				1,
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			types.NewGrant(
				2,
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
		},
	)

	// Store the genesis data
	subspacesDataBz, err := cfg.Codec.MarshalJSON(genesis)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = subspacesDataBz
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
					sdk.Coins{},
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
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
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
						sdk.Coins{},
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
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Subspaces, response.Subspaces)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQuerySection() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QuerySectionResponse
	}{
		{
			name: "non existing section returns error",
			args: []string{
				"1", "10",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "existing section is returned correctly",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QuerySectionResponse{
				Section: types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQuerySection()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QuerySectionResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Section, tc.expResponse.Section)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQuerySections() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QuerySectionsResponse
	}{
		{
			name: "sections are returned correctly",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 2),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QuerySectionsResponse{
				Sections: []types.Section{
					types.NewSection(
						1,
						1,
						0,
						"Test section",
						"Test section",
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQuerySections()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QuerySectionsResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Sections, response.Sections)
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
					types.NewUserGroup(2, 0, 1, "Another test group", "", types.NewPermissions(types.PermissionManageGroups)),
					types.NewUserGroup(2, 0, 2, "Third group", "", types.NewPermissions(poststypes.PermissionWrite)),
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
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(len(tc.expResponse.Groups), len(response.Groups))
				for i, group := range tc.expResponse.Groups {
					s.Require().True(group.Equal(response.Groups[i]))
				}
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
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().True(tc.expResponse.Group.Equal(response.Group))
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
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
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
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
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
				"11", "0", "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "user permissions are returned correctly",
			args: []string{
				"2", "0", "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryUserPermissionsResponse{
				Permissions: types.NewPermissions(types.PermissionManageGroups),
				Details: []types.PermissionDetail{
					types.NewPermissionDetailGroup(2, 0, 0, nil),
					types.NewPermissionDetailGroup(2, 0, 1, types.NewPermissions(types.PermissionManageGroups)),
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
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Permissions, response.Permissions)
				for i, detail := range tc.expResponse.Details {
					s.Require().True(detail.Equal(response.Details[i]))
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryUserAllowances() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryUserAllowancesResponse
	}{
		{
			name: "invalid subspace id returns error",
			args: []string{
				"subspace",
				"grantee",
			},
			shouldErr: true,
		},
		{
			name: "valid query without grantee is returned correctly",
			args: []string{
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryUserAllowancesResponse{
				Grants: []types.Grant{
					types.NewGrant(
						1,
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
						&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
					),
				},
			},
		},
		{
			name: "valid query is returned correctly",
			args: []string{
				"1",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryUserAllowancesResponse{
				Grants: []types.Grant{
					types.NewGrant(
						1,
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
						&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryUserAllowances()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryUserAllowancesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				for i, grant := range tc.expResponse.Grants {
					s.Require().True(grant.Equal(response.Grants[i]))
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryGroupAllowances() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryGroupAllowancesResponse
	}{
		{
			name: "invalid subspace id returns error",
			args: []string{
				"subspace",
				"group",
			},
			shouldErr: true,
		},
		{
			name: "invalid group id returns error",
			args: []string{
				"1",
				"group",
			},
			shouldErr: true,
		},
		{
			name: "valid query without group id is returned correctly",
			args: []string{
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryGroupAllowancesResponse{
				Grants: []types.Grant{
					types.NewGrant(
						1,
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						types.NewGroupGrantee(1),
						&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
					),
				},
			},
		},
		{
			name: "valid query is returned correctly",
			args: []string{
				"1",
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryGroupAllowancesResponse{
				Grants: []types.Grant{
					types.NewGrant(
						1,
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						types.NewGroupGrantee(1),
						&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryGroupAllowances()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryGroupAllowancesResponse
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &response), out.String())
				for i, grant := range tc.expResponse.Grants {
					s.Require().True(grant.Equal(response.Grants[i]))
				}
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

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
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error with custom owner",
			args: []string{
				"Another test subspace",
				fmt.Sprintf("--%s=%s", cli.FlagDescription, "Another test subspace"),
				fmt.Sprintf("--%s=%s", cli.FlagOwner, "cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u"),
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
			cmd := cli.GetCmdCreateSubspace()
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
			cmd := cli.GetCmdEditSubspace()
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
			cmd := cli.GetCmdDeleteSubspace()
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

// --------------------------------------------------------------------------------------------------------------------

func (s *IntegrationTestSuite) TestCmdCreateSection() {
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
				"0", "This is a new section",
			},
			shouldErr: true,
		},
		{
			name: "invalid name returns error",
			args: []string{
				"1", "",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "Test section",
				fmt.Sprintf("--%s=%s", cli.FlagDescription, "This is a test section"),
				fmt.Sprintf("--%s=%s", cli.FlagParent, "1"),
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
			cmd := cli.GetCmdCreateSection()
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

func (s *IntegrationTestSuite) TestCmdEditSection() {
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
			name: "invalid name returns error",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%s", cli.FlagName, ""),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%s", cli.FlagName, "Edited name"),
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
			cmd := cli.GetCmdEditSection()
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

func (s *IntegrationTestSuite) TestCmdMoveSection() {
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
				"0", "1", "2",
			},
			shouldErr: true,
		},
		{
			name: "invalid section id returns error",
			args: []string{
				"1", "0", "2",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1", "2",
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
			cmd := cli.GetCmdMoveSection()
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

func (s *IntegrationTestSuite) TestCmdDeleteSection() {
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
			name: "invalid section id returns error",
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
			cmd := cli.GetCmdDeleteSection()
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

// --------------------------------------------------------------------------------------------------------------------

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
				fmt.Sprintf("--%s=%s", cli.FlagPermissions, types.PermissionSetPermissions),
				fmt.Sprintf("--%s=%s", cli.FlagInitialMembers, "cosmos1g4yzh3q3grf804t4y4fuynrvrxtshgxy7j783f"),
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
			cmd := cli.GetCmdCreateUserGroup()
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
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
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
			cmd := cli.GetCmdEditUserGroup()
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

func (s *IntegrationTestSuite) TestCmdMoveUserGroup() {
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
				"0", "1", "2",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1", "2",
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
			cmd := cli.GetCmdMoveUserGroup()
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
				"1", "0", types.SerializePermission(poststypes.PermissionWrite),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error - group id > 0",
			args: []string{
				"1", "1", types.SerializePermission(poststypes.PermissionWrite),
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
			cmd := cli.GetCmdSetUserGroupPermissions()
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
			cmd := cli.GetCmdDeleteUserGroup()
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
			cmd := cli.GetCmdAddUserToUserGroup()
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
			cmd := cli.GetCmdRemoveUserFromUserGroup()
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

// --------------------------------------------------------------------------------------------------------------------

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
			cmd := cli.GetCmdSetUserPermissions()
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

// --------------------------------------------------------------------------------------------------------------------

func (s *IntegrationTestSuite) TestCmdGrantTreasuryAuthorization() {
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
				"x",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
			},
			shouldErr: true,
		},
		{
			name: "invalid expiration returns error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "x"),
			},
			shouldErr: true,
		},
		{
			name: "invalid msg returns error",
			args: []string{
				"0",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "10000"),
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "10000"),
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
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
			cmd := cli.GetCmdGrantTreasuryAuthorization()
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

func (s *IntegrationTestSuite) TestCmdGrantTreasuryAuthorization_SendAuthorization() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid spend limit returns error - invalid coins",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "x"),
			},
			shouldErr: true,
		},
		{
			name: "invalid spend limit returns error - negative coins",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "-1stake"),
			},
			shouldErr: true,
		},
		{
			name: "invalid msg returns error",
			args: []string{
				"0",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "10000"),
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "10000"),
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error - without expiration",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"send",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
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
			cmd := cli.GetCmdGrantTreasuryAuthorization()
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

func (s *IntegrationTestSuite) TestCmdGrantTreasuryAuthorization_GenericAuthorization() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "valid data returns no error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"generic",
				fmt.Sprintf("--%s=%s", authzcli.FlagMsgType, "/cosmos.bank.v1beta1.MsgSend"),
				fmt.Sprintf("--%s=%s", authzcli.FlagExpiration, "10000"),
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
			cmd := cli.GetCmdGrantTreasuryAuthorization()
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

func (s *IntegrationTestSuite) TestCmdGrantTreasuryAuthorization_StakeAuthorization() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "invalid spend limit returns error - invalid coins",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "x"),
			},
			shouldErr: true,
		},
		{
			name: "invalid spend limit returns error - negative coins",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "-1stake"),
			},
			shouldErr: true,
		},
		{
			name: "invalid allowed validators returns error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", authzcli.FlagAllowedValidators, "x"),
			},
			shouldErr: true,
		},
		{
			name: "invalid deny validators returns error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", authzcli.FlagDenyValidators, "x"),
			},
			shouldErr: true,
		},
		{
			name: "empty deny and allowed validators returns error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error - delegate authorization",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"delegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", authzcli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error - unbond authorization",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"unbond",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", authzcli.FlagAllowedValidators, val.ValAddress.String()),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data returns no error - redelegate authorization",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"redelegate",
				fmt.Sprintf("--%s=%s", authzcli.FlagSpendLimit, "100stake"),
				fmt.Sprintf("--%s=%s", authzcli.FlagAllowedValidators, val.ValAddress.String()),
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
			cmd := cli.GetCmdGrantTreasuryAuthorization()
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

func (s *IntegrationTestSuite) TestCmdRevokeTreasuryAuthorization() {
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
				"x",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"/cosmos.bank.v1beta1.MsgSend",
			},
			shouldErr: true,
		},
		{
			name: "invalid msg returns error",
			args: []string{
				"0",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"/cosmos.bank.v1beta1.MsgSend",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1",
				"cosmos1et50whs236j9dacacz7feh05jjum9jk04cdt9u",
				"/cosmos.bank.v1beta1.MsgSend",
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
			cmd := cli.GetCmdRevokeTreasuryAuthorization()
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

// --------------------------------------------------------------------------------------------------------------------

func (s *IntegrationTestSuite) TestCmdGrantAllowance() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"id"},
			shouldErr: true,
		},
		{
			name:      "empty grantee returns error",
			args:      []string{"1"},
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				fmt.Sprintf("--%s=%s", cli.FlagGroupGrantee, "1"),
			},
			shouldErr: true,
		},
		{
			name: "invalid msg returns error",
			args: []string{
				"0",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			},
			shouldErr: true,
		},
		{
			name: "invalid expiration returns error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				fmt.Sprintf("--%s=%s", feegrantcli.FlagExpiration, "invalid"),
			},
			shouldErr: true,
		},
		{
			name: "invalid periodic allowance returns error - period is greater than expiration",
			args: append(
				[]string{
					"1",
					fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					fmt.Sprintf("--%s=%s", feegrantcli.FlagSpendLimit, "100stake"),
					fmt.Sprintf("--%s=%s", feegrantcli.FlagExpiration, getFormattedExpiration(3600)),
					fmt.Sprintf("--%s=%d", feegrantcli.FlagPeriod, 36000),
					fmt.Sprintf("--%s=%s", feegrantcli.FlagPeriodLimit, "10stake"),
				},
			),
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "invalid periodic allowance returns error - invalid number of args",
			args: append(
				[]string{
					"1",
					fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					fmt.Sprintf("--%s=%s", feegrantcli.FlagSpendLimit, "100stake"),
					fmt.Sprintf("--%s=%s", feegrantcli.FlagPeriodLimit, "10stake"),
				},
			),
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data with default allowance returns no error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data with basic allowance returns no error",
			args: append(
				[]string{
					"1",
					fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					fmt.Sprintf("--%s=%s", feegrantcli.FlagSpendLimit, "100stake"),
					fmt.Sprintf("--%s=%s", feegrantcli.FlagExpiration, getFormattedExpiration(3600)),
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
					fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
				},
			),
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data with periodic allowance returns no error",
			args: append(
				[]string{
					"1",
					fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					fmt.Sprintf("--%s=%s", feegrantcli.FlagSpendLimit, "100stake"),
					fmt.Sprintf("--%s=%d", feegrantcli.FlagPeriod, 3600),
					fmt.Sprintf("--%s=%s", feegrantcli.FlagPeriodLimit, "10stake"),
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
					fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
				},
			),
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid data with filtered allowance returns no error",
			args: append(
				[]string{
					"1",
					fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					fmt.Sprintf("--%s=%s", feegrantcli.FlagAllowedMsgs, "/desmos.posts.v3.MsgCreatPost"),
					fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
					fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
					fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
					fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
				},
			),
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdGrantAllowance()
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

func (s *IntegrationTestSuite) TestCmdRevokeAllowance() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "invalid subspace id returns error",
			args:      []string{"id"},
			shouldErr: true,
		},
		{
			name:      "empty grantee returns error",
			args:      []string{"1"},
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				fmt.Sprintf("--%s=%s", cli.FlagGroupGrantee, "1"),
			},
			shouldErr: true,
		},
		{
			name: "invalid msg returns error",
			args: []string{
				"0",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%s", cli.FlagUserGrantee, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
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
			cmd := cli.GetCmdRevokeAllowance()
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

func getFormattedExpiration(duration int64) string {
	return time.Now().Add(time.Duration(duration) * time.Second).Format(time.RFC3339)
}
