package keeper_test

import (
	"encoding/hex"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_ExportGenesis() {
	var pubKey, err = sdk.GetPubKeyFromBech32(
		sdk.Bech32PubKeyTypeAccPub,
		"cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d",
	)
	suite.Require().NoError(err)

	usecases := []struct {
		name  string
		state struct {
			DTagRequests  []types.DTagTransferRequest
			Relationships []types.Relationship
			Blocks        []types.UserBlock
			Params        types.Params
			IBCPortID     string
			ChainLinks    []types.ChainLink
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
				ChainLinks    []types.ChainLink
			}{
				DTagRequests:  nil,
				Params:        types.DefaultParams(),
				Relationships: nil,
				Blocks:        nil,
				ChainLinks:    nil,
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
				ChainLinks    []types.ChainLink
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
				ChainLinks: []types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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

			suite.Require().NoError(suite.k.StoreProfile(suite.ctx, suite.testData.profile))
			for _, l := range uc.state.ChainLinks {
				suite.Require().NoError(suite.k.StoreChainLink(suite.ctx, l))
			}

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
	))

	profile2 := suite.CheckProfileNoError(types.NewProfile(
		"dtag-2",
		"nickname-2",
		"bio-2",
		types.NewPictures("profile-2", "cover-2"),
		time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
		authtypes.NewBaseAccountWithAddress(addr2),
	))

	suite.Require().NoError(err)

	// Generate keys
	priv1 := secp256k1.GenPrivKey()
	pubKey1 := priv1.PubKey()

	priv2 := secp256k1.GenPrivKey()
	pubKey2 := priv2.PubKey()

	// Get Bech32 encoded addresses
	linkAddr1, err := bech32.ConvertAndEncode("cosmos", pubKey1.Address().Bytes())
	suite.Require().NoError(err)
	linkAddr2, err := bech32.ConvertAndEncode("cosmos", pubKey2.Address().Bytes())
	suite.Require().NoError(err)

	// Get signature by signing with keys
	sig1, err := priv1.Sign([]byte(linkAddr1))
	suite.Require().NoError(err)
	sigHex1 := hex.EncodeToString(sig1)

	// Get signature by signing with keys
	sig2, err := priv2.Sign([]byte(linkAddr2))
	suite.Require().NoError(err)
	sigHex2 := hex.EncodeToString(sig2)

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
			ChainLinks           []types.ChainLink
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
				ChainLinks           []types.ChainLink
			}{
				Profiles:             nil,
				DTagTransferRequests: nil,
				Relationships:        nil,
				Blocks:               nil,
				Params:               types.DefaultParams(),
				IBCPortID:            "profiles-port-id",
				ChainLinks:           nil,
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
				profile1,
			},
			genesis: types.NewGenesisState(
				nil,
				[]types.Relationship{},
				[]types.UserBlock{},
				types.DefaultParams(),
				"profiles-port-id",
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(linkAddr1, "cosmos"),
						types.NewProof(pubKey1, sigHex1, linkAddr1),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(linkAddr1, "cosmos"),
						types.NewProof(pubKey2, sigHex2, linkAddr2),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				},
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
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(linkAddr1, "cosmos"),
						types.NewProof(pubKey1, sigHex1, linkAddr1),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(linkAddr2, "cosmos"),
						types.NewProof(pubKey2, sigHex2, linkAddr2),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				},
			),
			expState: struct {
				Profiles             []*types.Profile
				DTagTransferRequests []types.DTagTransferRequest
				Relationships        []types.Relationship
				Blocks               []types.UserBlock
				Params               types.Params
				IBCPortID            string
				ChainLinks           []types.ChainLink
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
				ChainLinks: []types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(linkAddr1, "cosmos"),
						types.NewProof(pubKey1, sigHex1, linkAddr1),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(linkAddr2, "cosmos"),
						types.NewProof(pubKey2, sigHex2, linkAddr2),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				},
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
