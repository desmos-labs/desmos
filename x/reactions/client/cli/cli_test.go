//go:build norace
// +build norace

package cli_test

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/v3/app"
	cliutils "github.com/desmos-labs/desmos/v3/x/reactions/client/utils"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/desmos-labs/desmos/v3/x/reactions/client/cli"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v3/testutil"
	"github.com/desmos-labs/desmos/v3/x/reactions/types"
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
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
		},
		nil, nil, nil, nil,
	)
	subspacesDataBz, err := cfg.Codec.MarshalJSON(subspacesGenesis)
	s.Require().NoError(err)
	genesisState[subspacestypes.ModuleName] = subspacesDataBz

	// Initialize the posts module genesis state
	postsGenesis := poststypes.NewGenesisState(
		[]poststypes.SubspaceDataEntry{
			poststypes.NewSubspaceDataEntry(1, 3),
		},
		[]poststypes.Post{
			poststypes.NewPost(
				1,
				0,
				1,
				"External ID",
				"This is a text",
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				0,
				nil,
				[]poststypes.PostReference{},
				poststypes.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			poststypes.NewPost(
				1,
				0,
				2,
				"External ID",
				"This is a text",
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				0,
				nil,
				[]poststypes.PostReference{},
				poststypes.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
		},
		[]poststypes.PostDataEntry{
			poststypes.NewPostDataEntry(1, 1, 1),
			poststypes.NewPostDataEntry(1, 1, 1),
		},
		nil,
		nil,
		nil,
		poststypes.DefaultParams(),
	)
	postsDataBz, err := cfg.Codec.MarshalJSON(postsGenesis)
	s.Require().NoError(err)
	genesisState[poststypes.ModuleName] = postsDataBz

	// Initialize the module genesis data
	genesis := types.NewGenesisState(
		[]types.SubspaceDataEntry{
			types.NewSubspaceDataEntry(1, 3),
		},
		[]types.RegisteredReaction{
			types.NewRegisteredReaction(
				1,
				1,
				":hello:",
				"https://example.com?image=hello.png",
			),
			types.NewRegisteredReaction(
				1,
				2,
				":wave:",
				"https://example.com?image=wave.png",
			),
		},
		[]types.PostDataEntry{
			types.NewPostDataEntry(1, 1, 4),
			types.NewPostDataEntry(1, 2, 2),
		},
		[]types.Reaction{
			types.NewReaction(
				1,
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1uh2ulr5unm800ttf05r6f7x82wg8ah4z8h8cr8",
			),
			types.NewReaction(
				1,
				1,
				2,
				types.NewFreeTextValue("🚀"),
				"cosmos1uh2ulr5unm800ttf05r6f7x82wg8ah4z8h8cr8",
			),
			types.NewReaction(
				1,
				1,
				3,
				types.NewRegisteredReactionValue(1),
				"cosmos14z8mn9ywhqu84alr5grxuljwj87jyz0zpxnlxy",
			),
			types.NewReaction(
				1,
				2,
				1,
				types.NewFreeTextValue("🚀"),
				"cosmos1uh2ulr5unm800ttf05r6f7x82wg8ah4z8h8cr8",
			),
		},
		[]types.SubspaceReactionsParams{
			types.NewSubspaceReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 10, ""),
			),
		},
	)

	// Store the genesis data
	reportsGenesisBz, err := cfg.Codec.MarshalJSON(genesis)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = reportsGenesisBz
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

func (s *IntegrationTestSuite) TestCmdQueryReaction() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryReactionResponse
	}{
		{
			name: "invalid subspace id returns error",
			args: []string{
				"0", "1", "2",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			args: []string{
				"1", "0", "2",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "invalid reaction id returns error",
			args: []string{
				"1", "1", "0",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "reaction is returned properly",
			args: []string{
				"1", "1", "2",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryReactionResponse{
				Reaction: types.NewReaction(
					1,
					1,
					2,
					types.NewFreeTextValue("🚀"),
					"cosmos1uh2ulr5unm800ttf05r6f7x82wg8ah4z8h8cr8",
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryReaction()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryReactionResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Reaction, tc.expResponse.Reaction)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryReactions() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryReactionsResponse
	}{
		{
			name: "posts reactions are returned properly",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 1),
				fmt.Sprintf("--%s=%d", flags.FlagOffset, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryReactionsResponse{
				Reactions: []types.Reaction{
					types.NewReaction(
						1,
						1,
						2,
						types.NewFreeTextValue("🚀"),
						"cosmos1uh2ulr5unm800ttf05r6f7x82wg8ah4z8h8cr8",
					),
				},
			},
		},
		{
			name: "user reactions are returned correctly",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%s", cli.FlagUser, "cosmos14z8mn9ywhqu84alr5grxuljwj87jyz0zpxnlxy"),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryReactionsResponse{
				Reactions: []types.Reaction{
					types.NewReaction(
						1,
						1,
						3,
						types.NewRegisteredReactionValue(1),
						"cosmos14z8mn9ywhqu84alr5grxuljwj87jyz0zpxnlxy",
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryReactions()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryReactionsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Reactions, tc.expResponse.Reactions)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryRegisteredReaction() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryRegisteredReactionResponse
	}{
		{
			name: "invalid subspace id returns error",
			args: []string{
				"0", "1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "invalid reaction id returns error",
			args: []string{
				"1", "0",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "reaction is returned properly",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryRegisteredReactionResponse{
				RegisteredReaction: types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryRegisteredReaction()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryRegisteredReactionResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.RegisteredReaction, tc.expResponse.RegisteredReaction)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryRegisteredReactions() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryRegisteredReactionsResponse
	}{
		{
			name: "registered reactions are returned correctly",
			args: []string{
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryRegisteredReactionsResponse{
				RegisteredReactions: []types.RegisteredReaction{
					types.NewRegisteredReaction(
						1,
						1,
						":hello:",
						"https://example.com?image=hello.png",
					),
					types.NewRegisteredReaction(
						1,
						2,
						":wave:",
						"https://example.com?image=wave.png",
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryRegisteredReactions()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryRegisteredReactionsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.RegisteredReactions, tc.expResponse.RegisteredReactions)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryParams() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryReactionsParamsResponse
	}{
		{
			name: "params are returned correctly",
			args: []string{
				"1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryReactionsParamsResponse{
				Params: types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 10, ""),
				),
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

				var response types.QueryReactionsParamsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Params, tc.expResponse.Params)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdAddReaction() {
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
			name: "invalid post id returns error",
			args: []string{
				"1", "0",
			},
			shouldErr: true,
		},
		{
			name: "missing flags return error",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: true,
		},
		{
			name: "both flags set return error",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%d", cli.FlagRegisteredReaction, 1),
				fmt.Sprintf("--%s=%s", cli.FlagFreeTextReaction, "🚀"),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: true,
		},
		{
			name: "valid registered reaction data returns no error",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%d", cli.FlagRegisteredReaction, 1),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid free text reaction data returns no error",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%s", cli.FlagFreeTextReaction, "🚀"),
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
			cmd := cli.GetCmdAddReaction()
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

func (s *IntegrationTestSuite) TestGetCmdRemoveReaction() {
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
				"0", "1", "1",
			},
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			args: []string{
				"1", "0", "1",
			},
			shouldErr: true,
		},
		{
			name: "invalid reaction id returns error",
			args: []string{
				"1", "1", "0",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1", "1",
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
			cmd := cli.GetCmdRemoveReaction()
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

func (s *IntegrationTestSuite) TestGetCmdAddRegisteredReaction() {
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
				"0", ":rocket:", "https://example.com?image=rocket.png",
			},
			shouldErr: true,
		},
		{
			name: "invalid shorthand code returns error",
			args: []string{
				"1", "", "https://example.com?image=rocket.png",
			},
			shouldErr: true,
		},
		{
			name: "invalid display value returns error",
			args: []string{
				"1", ":rocket:", "",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", ":rocket:", "https://example.com?image=rocket.png",
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
			cmd := cli.GetCmdAddRegisteredReaction()
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

func (s *IntegrationTestSuite) TestGetCmdEditRegisteredReaction() {
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
			name: "invalid reason id returns error",
			args: []string{
				"1", "0",
			},
			shouldErr: true,
		},
		{
			name: "invalid shorthand code value returns error",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%s", cli.FlagShorthandCode, ""),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: true,
		},
		{
			name: "invalid display value returns error",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%s", cli.FlagDisplayValue, ""),
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%s", cli.FlagShorthandCode, ":moon:"),
				fmt.Sprintf("--%s=%s", cli.FlagDisplayValue, "https://example.com?image=moon.png"),
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
			cmd := cli.GetCmdEditRegisteredReaction()
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

func (s *IntegrationTestSuite) TestGetCmdRemoveRegisteredReaction() {
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
			name: "invalid reason id returns error",
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
			cmd := cli.GetCmdRemoveRegisteredReaction()
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

func (s *IntegrationTestSuite) writeSetReactionsParamsJSONFile() (filePath string) {
	jsonData := cliutils.SetReactionsParamsJSON{
		RegisteredReactionParams: types.NewRegisteredReactionValueParams(true),
		FreeTextParams:           types.NewFreeTextValueParams(true, 10, "[a-z]"),
	}

	cdc, _ := app.MakeCodecs()
	jsonBz := cdc.MustMarshalJSON(&jsonData)

	// Write the JSON to a temp file
	f, err := ioutil.TempFile(s.T().TempDir(), "set-reactions-params")
	s.Require().NoError(err)
	defer f.Close()

	err = ioutil.WriteFile(f.Name(), jsonBz, 0644)
	s.Require().NoError(err)
	return f.Name()
}

func (s *IntegrationTestSuite) TestCmdSetParams() {
	filePath := s.writeSetReactionsParamsJSONFile()
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
				"0", filePath,
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", filePath,
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
			cmd := cli.GetCmdSetParams()
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
