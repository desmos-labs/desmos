package cli_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/desmos-labs/desmos/testutil"
	"github.com/desmos-labs/desmos/x/staging/posts/client/cli"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func TestIntegrationTestSuite(t *testing.T) {
	//TODO restore this when out of staging
	//suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := testutil.DefaultConfig()
	genesisState := cfg.GenesisState
	cfg.NumValidators = 2

	var postsData types.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &postsData))

	creationDate, err := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	s.Require().NoError(err)

	pollEndDate, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	s.Require().NoError(err)

	postsData.RegisteredReactions = []types.RegisteredReaction{
		types.NewRegisteredReaction(
			"cosmos1lhhkerae9cu3fa442vt50t32grlajun5lmrv3g",
			":reaction:",
			"https://example.com/reaction.jpg",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewRegisteredReaction(
			"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			":smile-jpg:",
			"https://smile.jpg",
			"5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}
	postsData.Posts = []types.Post{
		{
			PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			Message:              "Post message",
			Created:              creationDate,
			LastEdited:           creationDate.Add(1),
			Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			CommentsState:        types.CommentsStateAllowed,
			AdditionalAttributes: nil,
			Creator:              "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			Attachments: types.NewAttachments(
				types.NewAttachment(
					"https://uri.com",
					"text/plain",
					[]string{"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
				),
			),
			PollData: &types.PollData{
				Question: "poll?",
				ProvidedAnswers: types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
				),
				EndDate:               pollEndDate,
				AllowsMultipleAnswers: true,
				AllowsAnswerEdits:     true,
			},
		},
	}
	postsData.UsersPollAnswers = []types.UserAnswer{
		types.NewUserAnswer(
			"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			"cosmos1unacjuhyamzks5yu7qwlfuahdedd838e6fmdta",
			[]string{"1"},
		),
		types.NewUserAnswer(
			"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			[]string{"1"},
		),
	}
	postsData.PostsReactions = []types.PostReactionsEntry{
		types.NewPostReactionsEntry(
			"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			[]types.PostReaction{
				types.NewPostReaction(
					":broken_heart:",
					"💔",
					"cosmos12t08qkk4dm2pqgyy8hmq5hx92y2m29zedmdw7f",
				),
			}),
	}

	postsData.Reports = []types.Report{
		types.NewReport(
			"2b6284dd0361c20022ce366f4355c052165c0c23d7f588da5ac3572d68fda2f2",
			"scam",
			"Test report",
			"cosmos1azqm9kmyxunkx2yt332hmnr8sa3lclhjlg9w5k",
		),
	}

	postsData.Params = types.DefaultParams()

	postsDataBz, err := cfg.Codec.MarshalJSON(&postsData)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = postsDataBz
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

func (s *IntegrationTestSuite) TestCmdQueryPost() {
	val := s.network.Validators[0]

	creationDate, err := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	s.Require().NoError(err)

	pollEndDate, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	s.Require().NoError(err)

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryPostResponse
	}{
		{
			name:      "non existing post",
			args:      []string{"post_id"},
			expectErr: true,
		},
		{
			name: "existing post is returned properly",
			args: []string{
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryPostResponse{
				Post: types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post message",
					Created:              creationDate,
					LastEdited:           creationDate.Add(1),
					CommentsState:        types.CommentsStateAllowed,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: []types.Attribute{},
					Creator:              "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					Attachments: types.NewAttachments(
						types.NewAttachment(
							"https://uri.com",
							"text/plain",
							[]string{"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
						),
					),
					PollData: &types.PollData{
						Question: "poll?",
						ProvidedAnswers: types.NewPollAnswers(
							types.NewPollAnswer("1", "Yes"),
							types.NewPollAnswer("2", "No"),
						),
						EndDate:               pollEndDate,
						AllowsMultipleAnswers: true,
						AllowsAnswerEdits:     true,
					},
				},
				PollAnswers: []types.UserAnswer{
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1unacjuhyamzks5yu7qwlfuahdedd838e6fmdta",
						[]string{"1"},
					),
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						[]string{"1"},
					),
				},
				Reactions: []types.PostReaction{
					types.NewPostReaction(
						":broken_heart:",
						"💔",
						"cosmos12t08qkk4dm2pqgyy8hmq5hx92y2m29zedmdw7f",
					),
				},
				Children: []string{},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryPost()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryPostResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(response.Post, tc.expectedOutput.Post)
				s.Require().NotEmpty(response.Reactions)
				s.Require().NotEmpty(response.PollAnswers)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryPosts() {
	val := s.network.Validators[0]

	creationDate, err := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	s.Require().NoError(err)

	pollEndDate, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	s.Require().NoError(err)

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryPostsResponse
	}{
		{
			name: "existing posts are returned properly",
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryPostsResponse{
				Posts: []types.QueryPostResponse{
					{
						Post: types.Post{
							PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
							Message:              "Post message",
							Created:              creationDate,
							LastEdited:           creationDate.Add(1),
							Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
							AdditionalAttributes: []types.Attribute{},
							Creator:              "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
							Attachments: types.NewAttachments(
								types.NewAttachment(
									"https://uri.com",
									"text/plain",
									[]string{"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
								),
							),
							PollData: &types.PollData{
								Question: "poll?",
								ProvidedAnswers: types.NewPollAnswers(
									types.NewPollAnswer("1", "Yes"),
									types.NewPollAnswer("2", "No"),
								),
								EndDate:               pollEndDate,
								AllowsMultipleAnswers: true,
								AllowsAnswerEdits:     true,
							},
						},
						PollAnswers: []types.UserAnswer{
							types.NewUserAnswer(
								"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
								"cosmos1unacjuhyamzks5yu7qwlfuahdedd838e6fmdta",
								[]string{"1"},
							),
							types.NewUserAnswer(
								"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
								"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
								[]string{"1"},
							),
						},
						Reactions: []types.PostReaction{
							types.NewPostReaction(
								":broken_heart:",
								"💔",
								"cosmos12t08qkk4dm2pqgyy8hmq5hx92y2m29zedmdw7f",
							),
						},
						Children: []string{},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryPosts()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryPostsResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().NotEmpty(response.Posts)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryPollAnswers() {
	val := s.network.Validators[0]

	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		expLen    int
	}{
		{
			name:      "invalid post id",
			args:      []string{"post_id"},
			shouldErr: true,
		},
		{
			name: "valid data is returned properly",
			args: []string{
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expLen:    2,
		},
		{
			name: "valid data with pagination is returned properly",
			args: []string{
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expLen:    1,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryPollAnswers()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryPollAnswersResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expLen, len(response.Answers))
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryRegisteredReactions() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryRegisteredReactionsResponse
	}{
		{
			name:      "data without subspace and pagination is returned properly",
			args:      []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			expectErr: false,
			expectedOutput: types.QueryRegisteredReactionsResponse{
				RegisteredReactions: []types.RegisteredReaction{
					types.NewRegisteredReaction(
						"cosmos1lhhkerae9cu3fa442vt50t32grlajun5lmrv3g",
						":reaction:",
						"https://example.com/reaction.jpg",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRegisteredReaction(
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						":smile-jpg:",
						"https://smile.jpg",
						"5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name:      "data with subspace is returned properly",
			args:      []string{"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			expectErr: false,
			expectedOutput: types.QueryRegisteredReactionsResponse{
				RegisteredReactions: []types.RegisteredReaction{
					types.NewRegisteredReaction(
						"cosmos1lhhkerae9cu3fa442vt50t32grlajun5lmrv3g",
						":reaction:",
						"https://example.com/reaction.jpg",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				Pagination: &query.PageResponse{
					NextKey: nil,
					Total:   0,
				},
			},
		},
		{
			name: "data with pagination is returned properly",
			args: []string{
				fmt.Sprintf("--%s=%d", flags.FlagPage, 2),
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryRegisteredReactionsResponse{
				RegisteredReactions: []types.RegisteredReaction{
					types.NewRegisteredReaction(
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						":smile-jpg:",
						"https://smile.jpg",
						"5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
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
			cmd := cli.GetCmdQueryRegisteredReactions()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryRegisteredReactionsResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryReports() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryReportsResponse
	}{
		{
			name: "missing post id",
			args: []string{
				"not-found",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr:      false,
			expectedOutput: types.QueryReportsResponse{Reports: []types.Report{}},
		},
		{
			name: "valid post id",
			args: []string{
				"2b6284dd0361c20022ce366f4355c052165c0c23d7f588da5ac3572d68fda2f2",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			expectErr: false,
			expectedOutput: types.QueryReportsResponse{Reports: []types.Report{
				types.NewReport(
					"2b6284dd0361c20022ce366f4355c052165c0c23d7f588da5ac3572d68fda2f2",
					"scam",
					"Test report",
					"cosmos1azqm9kmyxunkx2yt332hmnr8sa3lclhjlg9w5k",
				),
			}},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryReports()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryReportsResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryParams() {
	val := s.network.Validators[0]

	testCases := []struct {
		name           string
		args           []string
		expectErr      bool
		expectedOutput types.QueryParamsResponse
	}{
		{
			name:      "data is returned properly",
			args:      []string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			expectErr: false,
			expectedOutput: types.QueryParamsResponse{
				Params: types.DefaultParams(),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryParams()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryParamsResponse
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput, response)
			}
		})
	}
}

// ___________________________________________________________________________________________________________________

func (s *IntegrationTestSuite) TestCmdCreatePost() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid subspace returns error",
			args:   []string{"subspace"},
			expErr: true,
		},
		{
			name: "empty message returns error without flags",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			},
			expErr: true,
		},
		{
			name: "invalid parent id returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				fmt.Sprintf("--%s=%s", cli.FlagParentID, "parent_id"),
			},
			expErr: true,
		},
		{
			name: "invalid comments state value returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"my post",
				fmt.Sprintf("--%s=%s", cli.FlagCommentsState, "comment_state"),
			},
			expErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"my post",
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
			cmd := cli.GetCmdCreatePost()
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

func (s *IntegrationTestSuite) TestCmdEditPost() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid post id returns error",
			args:   []string{"post_id"},
			expErr: true,
		},
		{
			name: "empty message returns error without flags",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			},
			expErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"My message",
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
			cmd := cli.GetCmdEditPost()
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

func (s *IntegrationTestSuite) TestCmdReportPost() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid post id",
			expErr: true,
			args:   []string{"1", "scam", "message"},
		},
		{
			name:   "invalid report type",
			expErr: true,
			args:   []string{"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd", "", "message"},
		},
		{
			name:   "valid report",
			expErr: false,
			args: []string{
				"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
				"scam",
				"message",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
			},
			respType: &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdReportPost()
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

func (s *IntegrationTestSuite) TestCmdAddPostReaction() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid post id returns error",
			args:   []string{"post_id"},
			expErr: true,
		},
		{
			name: "invalid value returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"value",
			},
			expErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"💔",
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
			cmd := cli.GetCmdAddPostReaction()
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

func (s *IntegrationTestSuite) TestCmdRemovePostReaction() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid post id returns error",
			args:   []string{"post_id"},
			expErr: true,
		},
		{
			name: "invalid value returns error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"value",
			},
			expErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"💔",
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
			cmd := cli.GetCmdRemovePostReaction()
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

func (s *IntegrationTestSuite) TestCmdAnswerPoll() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name:   "invalid post id returns error",
			args:   []string{"post_id"},
			expErr: true,
		},
		{
			name: "empty answers return error",
			args: []string{
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			},
			expErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"1", "2",
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
			cmd := cli.GetCmdAnswerPoll()
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

func (s *IntegrationTestSuite) TestCmdRegisterReaction() {
	val := s.network.Validators[0]

	testCases := []struct {
		name     string
		args     []string
		expErr   bool
		respType proto.Message
	}{
		{
			name: "emoji short code returns error",
			args: []string{
				":broken_heart:",
				"https://example.com/reaction.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "invalid short code returns error",
			args: []string{
				":^76554:",
				"https://example.com/reaction.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "invalid value returns error",
			args: []string{
				":reaction_2:",
				"value",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "invalid subspace returns error",
			args: []string{
				":reaction_2:",
				"https://example.com/reaction_2.jpg",
				"subspace",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			},
			expErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				":new_reaction:",
				"https://example.com/new_reaction.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
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
			cmd := cli.GetCmdRegisterReaction()
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
