package cli_test

import (
	"fmt"
	"path"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/protobuf/proto"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/desmos-labs/desmos/x/profiles/client/cli"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (s *IntegrationTestSuite) TestCmdQueryUserChainLinks() {
	val := s.network.Validators[0]

	pubKey, err := sdk.GetPubKeyFromBech32(
		sdk.Bech32PubKeyTypeAccPub,
		"cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d",
	)
	s.Require().NoError(err)

	useCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryUserChainLinksResponse
	}{
		{
			name: "empty array is returned properly",
			args: []string{
				val.Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryUserChainLinksResponse{
				Links: []types.ChainLink{},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "existing chain links is returned properly",
			args: []string{
				"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryUserChainLinksResponse{
				Links: []types.ChainLink{
					types.NewChainLink(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
						types.NewProof(
							pubKey,
							"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
							"text",
						),
						types.NewChainConfig("cosmos"),
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
					types.NewChainLink(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
						types.NewProof(
							pubKey,
							"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
							"text",
						),
						types.NewChainConfig("cosmos"),
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
			cmd := cli.GetCmdQueryUserChainLinks()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, uc.args)

			if uc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryUserChainLinksResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())

				s.Require().Equal(uc.expectedOutput.Pagination, response.Pagination)
				for i, link := range response.Links {
					s.Require().True(link.Equal(response.Links[i]))
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
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "could not get destination key returns error",
			args: []string{
				"src",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, ""),
			},
			expErr:   true,
			respType: &sdk.TxResponse{},
		},
		{
			name: "could not get source key returns error",
			args: []string{
				"wrong",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, "dest"),
			},
			expErr:   true,
			respType: &sdk.TxResponse{},
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
			expErr:   false,
			respType: &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdLinkChainAccount()
			out, err := clitestutil.ExecTestCLICmd(cliCtx, cmd, tc.args)

			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(cliCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
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
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "empty chain name returns error",
			args: []string{
				"",
				src.GetAddress().String(),
			},
			expErr:   true,
			respType: &sdk.TxResponse{},
		},
		{
			name: "empty address returns error",
			args: []string{
				"cosmos",
				"",
			},
			expErr:   true,
			respType: &sdk.TxResponse{},
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
			expErr:   false,
			respType: &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdUnlinkChainAccount()
			out, err := clitestutil.ExecTestCLICmd(cliCtx, cmd, tc.args)

			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(cliCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}
