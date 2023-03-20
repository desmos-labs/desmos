//go:build norace
// +build norace

package cli_test

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/testutil"
	"github.com/desmos-labs/desmos/v4/x/posts/client/cli"
	cliutils "github.com/desmos-labs/desmos/v4/x/posts/client/utils"
	"github.com/desmos-labs/desmos/v4/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
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
			subspacestypes.NewSubspaceData(1, 2, 1),
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
		[]subspacestypes.Section{
			subspacestypes.NewSection(1, 1, 0, "Test section", "Test section"),
		},
		nil,
		nil,
		nil,
		nil,
	)
	subspacesDataBz, err := cfg.Codec.MarshalJSON(subspacesGenesis)
	s.Require().NoError(err)
	genesisState[subspacestypes.ModuleName] = subspacesDataBz

	// Initialize the module genesis data
	postsGenesis := types.NewGenesisState(
		[]types.SubspaceDataEntry{
			types.NewSubspaceDataEntry(1, 3),
		},
		[]types.Post{
			types.NewPost(
				1,
				0,
				1,
				"External ID",
				"This is a text",
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				0,
				nil,
				nil,
				[]types.PostReference{},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			types.NewPost(
				1,
				1,
				2,
				"External ID",
				"This is a text",
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				0,
				nil,
				nil,
				[]types.PostReference{},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
		},
		[]types.PostDataEntry{
			types.NewPostDataEntry(1, 1, 2),
			types.NewPostDataEntry(1, 2, 1),
		},
		[]types.Attachment{
			types.NewAttachment(1, 1, 1, types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("Dog", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				true,
				false,
				nil,
			)),
		},
		nil,
		[]types.UserAnswer{
			types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
			types.NewUserAnswer(1, 1, 1, []uint32{0, 1}, "cosmos1u65w3xnhga8ngyg44eudh07zdxmkzny6uaudfc"),
		},
		types.DefaultParams(),
	)
	postsDataBz, err := cfg.Codec.MarshalJSON(postsGenesis)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = postsDataBz

	// Store the genesis data
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

func (s *IntegrationTestSuite) TestCmdQueryPost() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryPostResponse
	}{
		{
			name: "non existing post returns error",
			args: []string{
				"1", "10",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: true,
		},
		{
			name: "existing post is returned correctly",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryPostResponse{
				Post: types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					[]string{},
					[]types.PostReference{},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryPost()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryPostResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Post, response.Post)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryPosts() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse proto.Message
		expPosts    []types.Post
	}{
		{
			name: "posts are returned correctly with only subspace",
			args: []string{
				"1",
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr:   false,
			expResponse: &types.QuerySubspacePostsResponse{},
			expPosts: []types.Post{
				types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					[]string{},
					[]types.PostReference{},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				),
			},
		},
		{
			name: "posts are returned correctly with section",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr:   false,
			expResponse: &types.QuerySectionPostsResponse{},
			expPosts: []types.Post{
				types.NewPost(
					1,
					1,
					2,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					[]string{},
					[]types.PostReference{},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryPosts()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.expResponse), out.String())

				switch res := tc.expResponse.(type) {
				case *types.QuerySubspacePostsResponse:
					s.Require().Equal(tc.expPosts, res.Posts)
				case *types.QuerySectionPostsResponse:
					s.Require().Equal(tc.expPosts, res.Posts)
				}
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryPostAttachments() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryPostAttachmentsResponse
	}{
		{
			name: "attachments are returned correctly",
			args: []string{
				"1", "1",
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 1),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryPostAttachmentsResponse{
				Attachments: []types.Attachment{
					types.NewAttachment(1, 1, 1, types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						true,
						false,
						nil,
					)),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryPostAttachments()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryPostAttachmentsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Attachments, response.Attachments)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryPollAnswers() {
	val := s.network.Validators[0]
	testCases := []struct {
		name        string
		args        []string
		shouldErr   bool
		expResponse types.QueryPollAnswersResponse
	}{
		{
			name: "answers are returned correctly if no user is specified",
			args: []string{
				"1", "1", "1",
				fmt.Sprintf("--%s=%d", flags.FlagLimit, 2),
				fmt.Sprintf("--%s=%d", flags.FlagPage, 1),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryPollAnswersResponse{
				Answers: []types.UserAnswer{
					types.NewUserAnswer(1, 1, 1, []uint32{0, 1}, "cosmos1u65w3xnhga8ngyg44eudh07zdxmkzny6uaudfc"),
					types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
				},
			},
		},
		{
			name: "answer is returned correctly if a user is specified",
			args: []string{
				"1", "1", "1", "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryPollAnswersResponse{
				Answers: []types.UserAnswer{
					types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
				},
			},
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
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Answers, response.Answers)
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
		expResponse types.QueryParamsResponse
	}{
		{
			name: "params are returned correctly",
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expResponse: types.QueryParamsResponse{
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

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryParamsResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expResponse.Params, response.Params)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (s *IntegrationTestSuite) writeCreatePostJSONFile() (filePath string) {
	attachments, err := types.PackAttachments([]types.AttachmentContent{})
	s.Require().NoError(err)

	jsonData := cliutils.CreatePostJSON{
		ExternalID: "This is my external id",
		Text:       "This is my post text",
		Entities: types.NewEntities(nil, nil, []types.Url{
			types.NewURL(0, 3, "https://example.com", "This"),
		}),
		Attachments:    attachments,
		ConversationID: 1,
		ReplySettings:  types.REPLY_SETTING_EVERYONE,
		ReferencedPosts: []types.PostReference{
			types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
		},
	}

	cdc, _ := app.MakeCodecs()
	jsonBz := cdc.MustMarshalJSON(&jsonData)

	// Write the JSON to a temp file
	f, err := ioutil.TempFile(s.T().TempDir(), "create-post")
	s.Require().NoError(err)
	defer f.Close()

	err = ioutil.WriteFile(f.Name(), jsonBz, 0644)
	s.Require().NoError(err)
	return f.Name()
}

func (s *IntegrationTestSuite) TestCmdCreatePost() {
	filePath := s.writeCreatePostJSONFile()
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
				"0", "1", filePath,
			},
			shouldErr: true,
		},
		{
			name: "invalid section id returns error",
			args: []string{
				"1", "0", filePath,
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1", filePath,
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
			cmd := cli.GetCmdCreatePost()
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

func (s *IntegrationTestSuite) writeEditPostJSONFile() (filePath string) {
	jsonData := cliutils.EditPostJSON{
		Text: "This is my edited text",
		Entities: types.NewEntities(nil, nil, []types.Url{
			types.NewURL(0, 3, "https://example.com", "This"),
		}),
	}

	cdc, _ := app.MakeCodecs()
	jsonBz := cdc.MustMarshalJSON(&jsonData)

	// Write the JSON to a temp file
	f, err := ioutil.TempFile(s.T().TempDir(), "edit-post")
	s.Require().NoError(err)
	defer f.Close()

	err = ioutil.WriteFile(f.Name(), jsonBz, 0644)
	s.Require().NoError(err)
	return f.Name()
}

func (s *IntegrationTestSuite) TestCmdEditPost() {
	filePath := s.writeEditPostJSONFile()
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
				"0", "1", filePath,
			},
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			args: []string{
				"1", "0", filePath,
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1", filePath,
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
			cmd := cli.GetCmdEditPost()
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

func (s *IntegrationTestSuite) TestCmdDeletePost() {
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
			cmd := cli.GetCmdDeletePost()
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

func (s *IntegrationTestSuite) writeAttachmentJSONFile() (filePath string) {
	attachmentAny, err := codectypes.NewAnyWithValue(types.NewMedia(
		"ftp://user:password@example.com/image.png",
		"image/png",
	))
	s.Require().NoError(err)

	cdc, _ := app.MakeCodecs()
	jsonBz := cdc.MustMarshalJSON(attachmentAny)

	// Write the JSON to a temp file
	f, err := ioutil.TempFile(s.T().TempDir(), "attachment")
	s.Require().NoError(err)
	defer f.Close()

	err = ioutil.WriteFile(f.Name(), jsonBz, 0644)
	s.Require().NoError(err)
	return f.Name()
}

func (s *IntegrationTestSuite) TestCmdAddPostAttachment() {
	filePath := s.writeAttachmentJSONFile()
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
				"0", "1", filePath,
			},
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			args: []string{
				"1", "0", filePath,
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1", filePath,
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
			cmd := cli.GetCmdAddPostAttachment()
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

func (s *IntegrationTestSuite) TestCmdRemovePostAttachment() {
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
			name: "invalid attachment id returns error",
			args: []string{
				"1", "1", "0",
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
			cmd := cli.GetCmdRemovePostAttachment()
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

func (s *IntegrationTestSuite) TestCmdAnswerPoll() {
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
				"0", "1", "1", "0,1",
			},
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			args: []string{
				"1", "0", "1", "0,1",
			},
			shouldErr: true,
		},
		{
			name: "invalid poll id returns error",
			args: []string{
				"1", "1", "0", "0,1",
			},
			shouldErr: true,
		},
		{
			name: "invalid answer returns error",
			args: []string{
				"1", "1", "1", "",
			},
			shouldErr: true,
		},
		{
			name: "valid data returns no error",
			args: []string{
				"1", "1", "1", "0,1,2",
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
			cmd := cli.GetCmdAnswerPoll()
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
