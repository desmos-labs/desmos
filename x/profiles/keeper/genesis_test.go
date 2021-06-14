package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_ExportGenesis() {
	usecases := []struct {
		name  string
		state struct {
			DTagRequests  []types.DTagTransferRequest
			Relationships []types.Relationship
			Blocks        []types.UserBlock
			Params        types.Params
			IBCPortID     string
		}
		expGenesis *types.GenesisState
	}{
		{
			name: "empty state",
			state: struct {
				DTagRequests  []types.DTagTransferRequest
				Relationships []types.Relationship
				Blocks        []types.UserBlock
				Params        types.Params
				IBCPortID     string
			}{
				DTagRequests:  nil,
				Params:        types.DefaultParams(),
				Relationships: nil,
				Blocks:        nil,
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, types.DefaultParams(), "", nil),
		},
		{
			name: "non-empty state",
			state: struct {
				DTagRequests  []types.DTagTransferRequest
				Relationships []types.Relationship
				Blocks        []types.UserBlock
				Params        types.Params
				IBCPortID     string
			}{
				DTagRequests: []types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				},
				Relationships: []types.Relationship{
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				Blocks: []types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				Params: types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				),
				IBCPortID: "port-id",
			},
			expGenesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				},
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				),
				"port-id",
				nil,
			),
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()

			for _, req := range uc.state.DTagRequests {
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(suite.ctx, req))
			}
			for _, rel := range uc.state.Relationships {
				suite.Require().NoError(suite.k.SaveRelationship(suite.ctx, rel))
			}
			for _, block := range uc.state.Blocks {
				suite.Require().NoError(suite.k.SaveUserBlock(suite.ctx, block))
			}
			suite.k.SetParams(suite.ctx, uc.state.Params)
			suite.k.SetPort(suite.ctx, uc.state.IBCPortID)

			exported := suite.k.ExportGenesis(suite.ctx)
			suite.Require().Equal(uc.expGenesis, exported)
		})
	}
}

func (suite *KeeperTestSuite) Test_InitGenesis() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.Require().NoError(err)

	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.Require().NoError(err)

	addr3, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.Require().NoError(err)

	profile1 := suite.CheckProfileNoError(types.NewProfile(
		"dtag-1",
		"nickname-1",
		"bio-1",
		types.NewPictures("profile-1", "cover-1"),
		time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		authtypes.NewBaseAccountWithAddress(addr1),
		nil,
	))

	profile2 := suite.CheckProfileNoError(types.NewProfile(
		"dtag-2",
		"nickname-2",
		"bio-2",
		types.NewPictures("profile-2", "cover-2"),
		time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
		authtypes.NewBaseAccountWithAddress(addr2),
		nil,
	))

	addr4, err := sdk.AccAddressFromBech32("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0")
	suite.Require().NoError(err)

	pubKey4, err := sdk.GetPubKeyFromBech32(
		sdk.Bech32PubKeyTypeAccPub,
		"cosmospub1addwnpepq0j8zw4t6tg3v8gh7d2d799gjhue7ewwmpg2hwr77f9kuuyzgqtrw5r6wec",
	)
	suite.Require().NoError(err)

	doubleLinksProfile := suite.CheckProfileNoError(types.NewProfile(
		"dtag-3",
		"nickname-3",
		"bio-3",
		types.NewPictures("profile-3", "cover-3"),
		time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
		authtypes.NewBaseAccountWithAddress(addr4),
		[]types.ChainLink{
			types.NewChainLink(
				types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
				types.NewProof(pubKey4, "sig_hex", "addr"),
				types.NewChainConfig("cosmos"),
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
			types.NewChainLink(
				types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
				types.NewProof(pubKey4, "sig_hex", "addr"),
				types.NewChainConfig("cosmos"),
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
		},
	))

	usecases := []struct {
		name         string
		authAccounts []authtypes.AccountI
		genesis      *types.GenesisState
		expErr       bool
		expState     struct {
			Profiles             []*types.Profile
			DTagTransferRequests []types.DTagTransferRequest
			Relationships        []types.Relationship
			Blocks               []types.UserBlock
			Params               types.Params
			IBCPortID            string
		}
	}{
		{
			name:    "empty genesis",
			genesis: types.NewGenesisState(nil, nil, nil, types.DefaultParams(), types.IBCPortID, nil),
			expState: struct {
				Profiles             []*types.Profile
				DTagTransferRequests []types.DTagTransferRequest
				Relationships        []types.Relationship
				Blocks               []types.UserBlock
				Params               types.Params
				IBCPortID            string
			}{
				Profiles:             nil,
				DTagTransferRequests: nil,
				Relationships:        nil,
				Blocks:               nil,
				Params:               types.DefaultParams(),
				IBCPortID:            "profiles-port-id",
			},
		},
		{
			name: "double relationships panics",
			genesis: types.NewGenesisState(
				nil,
				[]types.Relationship{
					types.NewRelationship("creator", "recipient", "subspace"),
					types.NewRelationship("creator", "recipient", "subspace"),
				},
				[]types.UserBlock{},
				types.DefaultParams(),
				"profiles-port-id",
				nil,
			),
			expErr: true,
		},
		{
			name: "double user block panics",
			genesis: types.NewGenesisState(
				nil,
				[]types.Relationship{},
				[]types.UserBlock{
					types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
					types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
				},
				types.DefaultParams(),
				"profiles-port-id",
				nil,
			),
			expErr: true,
		},
		{
			name: "double chain link panics",
			authAccounts: []authtypes.AccountI{
				doubleLinksProfile,
			},
			genesis: types.NewGenesisState(
				nil,
				[]types.Relationship{},
				[]types.UserBlock{},
				types.DefaultParams(),
				"profiles-port-id",
				nil,
			),
			expErr: true,
		},
		{
			name: "non-empty genesis",
			authAccounts: []authtypes.AccountI{
				profile1,
				profile2,
				authtypes.NewBaseAccountWithAddress(addr3),
			},
			genesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				},
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				),
				"profiles-port-id",
				nil,
			),
			expState: struct {
				Profiles             []*types.Profile
				DTagTransferRequests []types.DTagTransferRequest
				Relationships        []types.Relationship
				Blocks               []types.UserBlock
				Params               types.Params
				IBCPortID            string
			}{
				Profiles: []*types.Profile{
					profile1,
					profile2,
				},
				DTagTransferRequests: []types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				},
				Relationships: []types.Relationship{
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				Blocks: []types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				Params: types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				),
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()

			for _, acc := range uc.authAccounts {
				suite.ak.SetAccount(suite.ctx, acc)
			}

			if uc.expErr {
				suite.Require().Panics(func() { suite.k.InitGenesis(suite.ctx, *uc.genesis) })
			} else {
				suite.Require().NotPanics(func() { suite.k.InitGenesis(suite.ctx, *uc.genesis) })

				suite.Require().Equal(uc.expState.Profiles, suite.k.GetProfiles(suite.ctx))
				suite.Require().Equal(uc.expState.DTagTransferRequests, suite.k.GetDTagTransferRequests(suite.ctx))
				suite.Require().Equal(uc.expState.Params, suite.k.GetParams(suite.ctx))
			}
		})
	}
}
