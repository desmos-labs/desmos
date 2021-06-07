package cli_test

import (
	"fmt"
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

func (s *IntegrationTestSuite) TestCmdQueryProfile() {
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
		expectedOutput types.QueryProfileResponse
	}{
		{
			name: "non existing profile",
			args: []string{
				s.network.Validators[1].Address.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryProfileResponse{
				Profile: nil,
			},
		},
		{
			name: "existing profile is returned properly",
			args: []string{
				"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryProfileResponse{
				Profile: profileAny,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryProfile()
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

func (s *IntegrationTestSuite) TestCmdSaveProfile() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "invalid dtag returns error",
			args: []string{
				"",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "invalid creator returns error",
			args: []string{
				"dtag",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, ""),
			},
			expErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
				"dtag",
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
			cmd := cli.GetCmdSaveProfile()
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

func (s *IntegrationTestSuite) TestCmdDeleteProfile() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid user returns error",
			args:   []string{fmt.Sprintf("--%s=%s", flags.FlagFrom, "")},
			expErr: true,
		},
		{
			name: "correct data returns no error",
			args: []string{
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
			cmd := cli.GetCmdDeleteProfile()
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
