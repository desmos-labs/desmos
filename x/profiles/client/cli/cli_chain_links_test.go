package cli_test

import (
	"fmt"
	"path"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/desmos-labs/desmos/x/profiles/client/cli"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (s *IntegrationTestSuite) TestCmdQueryProfileByChainLink() {
	val := s.network.Validators[0]

	addr, err := sdk.AccAddressFromBech32("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs")
	s.Require().NoError(err)

	profile, err := types.NewProfile(
		"dtag",
		"nickname",
		"bio",
		types.Pictures{},
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		authtypes.NewBaseAccountWithAddress(addr),
	)
	s.Require().NoError(err)

	srcKey, err := s.keyBase.Key("src")
	profile.ChainsLinks = []types.ChainLink{
		types.NewChainLink(
			types.NewBech32Address(srcKey.GetAddress().String(), "cosmos"),
			types.NewProof(srcKey.GetPubKey(), "signature", "plain_text"),
			types.NewChainConfig("cosmos"),
			time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		),
	}

	profileAny, err := codectypes.NewAnyWithValue(profile)
	s.Require().NoError(err)

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryProfileByChainLinkResponse
	}{
		{
			name: "no link associated to profile",
			args: []string{
				"cosmos",
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: true,
			expectedOutput: types.QueryProfileByChainLinkResponse{
				Profile: nil,
			},
		},
		{
			name: "existing profile is returned properly",
			args: []string{
				"cosmos",
				srcKey.GetAddress().String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryProfileByChainLinkResponse{
				Profile: profileAny,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryProfileByChainLink()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryProfileResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().True(tc.expectedOutput.Profile.Equal(response.Profile))
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
