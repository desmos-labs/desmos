//go:build norace
// +build norace

package cli_test

import (
	"fmt"
	"path"
	"time"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/protobuf/proto"

	"github.com/desmos-labs/desmos/v4/x/profiles/client/cli"
	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

func (s *IntegrationTestSuite) TestCmdQueryChainLinks() {
	val := s.network.Validators[0]
	useCases := []struct {
		name           string
		args           []string
		shouldErr      bool
		expectedOutput types.QueryChainLinksResponse
	}{
		{
			name: "existing chain links are returned properly",
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryChainLinksResponse{
				Links: []types.ChainLink{
					s.testChainLinkAccount.GetBech32ChainLink(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "empty array is returned properly",
			args: []string{
				val.Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryChainLinksResponse{
				Links: []types.ChainLink{},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "existing chain links of the given user are returned properly",
			args: []string{
				"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryChainLinksResponse{
				Links: []types.ChainLink{
					s.testChainLinkAccount.GetBech32ChainLink(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
	}

	for _, uc := range useCases {
		uc := uc

		s.Run(uc.name, func() {
			cmd := cli.GetCmdQueryChainLinks()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, uc.args)

			if uc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryChainLinksResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(uc.expectedOutput.Pagination, response.Pagination)
				for i := range uc.expectedOutput.Links {
					s.Require().True(uc.expectedOutput.Links[i].Equal(response.Links[i]))
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryChainLinkOwners() {
	val := s.network.Validators[0]
	target := s.testChainLinkAccount.GetBech32ChainLink(
		"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
		time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
	).GetAddressData().GetValue()

	useCases := []struct {
		name           string
		args           []string
		shouldErr      bool
		expectedOutput types.QueryChainLinkOwnersResponse
	}{
		{
			name: "existing chain link owners are returned properly",
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryChainLinkOwnersResponse{
				Owners: []types.QueryChainLinkOwnersResponse_ChainLinkOwnerDetails{
					{
						User:      "cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						ChainName: "cosmos",
						Target:    target,
					},
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "empty array is returned properly",
			args: []string{
				"desmos",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryChainLinkOwnersResponse{
				Owners: []types.QueryChainLinkOwnersResponse_ChainLinkOwnerDetails{},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "existing chain link owners of the given chain name are returned properly",
			args: []string{
				"cosmos",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryChainLinkOwnersResponse{
				Owners: []types.QueryChainLinkOwnersResponse_ChainLinkOwnerDetails{
					{
						User:      "cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						ChainName: "cosmos",
						Target:    target,
					},
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
	}

	for _, uc := range useCases {
		uc := uc

		s.Run(uc.name, func() {
			cmd := cli.GetCmdQueryChainLinkOwners()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, uc.args)

			if uc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryChainLinkOwnersResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(uc.expectedOutput.Pagination, response.Pagination)
				s.Require().Equal(uc.expectedOutput.Owners, response.Owners)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryDefaultExternalAddresses() {
	val := s.network.Validators[0]
	useCases := []struct {
		name           string
		args           []string
		shouldErr      bool
		expectedOutput types.QueryDefaultExternalAddressesResponse
	}{
		{
			name: "all default chain links are returned properly",
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryDefaultExternalAddressesResponse{
				Links: []types.ChainLink{
					s.testChainLinkAccount.GetBech32ChainLink(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "empty array is returned properly",
			args: []string{
				val.Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryDefaultExternalAddressesResponse{
				Links: []types.ChainLink{},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "existing chain links of the given owner are returned properly",
			args: []string{
				"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryDefaultExternalAddressesResponse{
				Links: []types.ChainLink{
					s.testChainLinkAccount.GetBech32ChainLink(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
	}

	for _, uc := range useCases {
		uc := uc

		s.Run(uc.name, func() {
			cmd := cli.GetCmdQueryDefaultExternalAddresses()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, uc.args)

			if uc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryChainLinksResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(uc.expectedOutput.Pagination, response.Pagination)
				for i := range uc.expectedOutput.Links {
					s.Require().True(uc.expectedOutput.Links[i].Equal(response.Links[i]))
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdLinkChainAccount() {
	cliCtx := s.network.Validators[0].ClientCtx
	cliCtx.Keyring = s.keyBase

	filePath := path.Join(s.T().TempDir(), "data.json")
	s.writeChainLinkJSONFile(filePath)

	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "could not get destination key returns error",
			args: []string{
				"src",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, ""),
			},
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "could not get source key returns error",
			args: []string{
				"wrong",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "dest"),
			},
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid request works properly",
			args: []string{
				filePath,
				fmt.Sprintf("--%s=%s", flags.FlagFrom, dstKeyName),
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
			cmd := cli.GetCmdLinkChainAccount()
			out, err := clitestutil.ExecTestCLICmd(cliCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(cliCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdUnlinkChainAccount() {
	cliCtx := s.network.Validators[0].ClientCtx
	cliCtx.Keyring = s.keyBase
	src, err := s.keyBase.Key("src")
	s.Require().NoError(err)
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "empty chain name returns error",
			args: []string{
				"",
				src.GetAddress().String(),
			},
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "empty address returns error",
			args: []string{
				"cosmos",
				"",
			},
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid request works properly",
			args: []string{
				"cosmos",
				src.GetAddress().String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, dstKeyName),
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
			cmd := cli.GetCmdUnlinkChainAccount()
			out, err := clitestutil.ExecTestCLICmd(cliCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(cliCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdSetDefaultExternalAddress() {
	cliCtx := s.network.Validators[0].ClientCtx
	cliCtx.Keyring = s.keyBase
	src, err := s.keyBase.Key("src")
	s.Require().NoError(err)
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name: "empty chain name returns error",
			args: []string{
				"",
				src.GetAddress().String(),
			},
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "empty address returns error",
			args: []string{
				"cosmos",
				"",
			},
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid request works properly",
			args: []string{
				"cosmos",
				src.GetAddress().String(),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, dstKeyName),
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
			cmd := cli.GetCmdSetDefaultExternalAddress()
			out, err := clitestutil.ExecTestCLICmd(cliCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(cliCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}
