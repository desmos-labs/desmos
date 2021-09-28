package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/testutil"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_IterateProfile() {
	date, err := time.Parse(time.RFC3339, "2010-10-02T12:10:00.000Z")
	suite.Require().NoError(err)

	addr1, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.Require().NoError(err)

	addr2, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.Require().NoError(err)

	addr3, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.Require().NoError(err)

	addr4, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	suite.Require().NoError(err)

	profiles := []*types.Profile{
		suite.CheckProfileNoError(
			types.NewProfileFromAccount("first", authtypes.NewBaseAccountWithAddress(addr1), date),
		),
		suite.CheckProfileNoError(
			types.NewProfileFromAccount("second", authtypes.NewBaseAccountWithAddress(addr2), date),
		),
		suite.CheckProfileNoError(
			types.NewProfileFromAccount("not", authtypes.NewBaseAccountWithAddress(addr3), date),
		),
		suite.CheckProfileNoError(
			types.NewProfileFromAccount("third", authtypes.NewBaseAccountWithAddress(addr4), date),
		),
	}

	expProfiles := []*types.Profile{
		profiles[0],
		profiles[1],
		profiles[3],
	}

	for _, profile := range profiles {
		err := suite.k.StoreProfile(suite.ctx, profile)
		suite.Require().NoError(err)
	}

	var validProfiles []*types.Profile
	suite.k.IterateProfiles(suite.ctx, func(_ int64, profile *types.Profile) (stop bool) {
		if profile.DTag == "not" {
			return false
		}
		validProfiles = append(validProfiles, profile)
		return false
	})

	suite.Len(expProfiles, len(validProfiles))
	for _, profile := range validProfiles {
		suite.Contains(expProfiles, profile)
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetProfiles() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		expProfiles []*types.Profile
	}{
		{
			name: "non empty profiles list is returned properly",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
			},
			expProfiles: []*types.Profile{
				testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			},
		},
		{
			name:        "empty profiles list is returned properly",
			expProfiles: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res := suite.k.GetProfiles(ctx)
			suite.Require().Equal(tc.expProfiles, res)
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_IterateUserIncomingDTagTransferRequests() {
	address := "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"
	requests := []types.DTagTransferRequest{
		types.NewDTagTransferRequest(
			"DTag1",
			"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			address,
		),
		types.NewDTagTransferRequest(
			"DTag2",
			"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			address,
		),
		types.NewDTagTransferRequest(
			"DTag3",
			"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			address,
		),
	}

	for _, request := range requests {
		profile := testutil.ProfileFromAddr(address)
		err := suite.k.StoreProfile(suite.ctx, profile)
		suite.Require().NoError(err)

		err = suite.k.SaveDTagTransferRequest(suite.ctx, request)
		suite.Require().NoError(err)
	}

	iterations := 0
	suite.k.IterateUserIncomingDTagTransferRequests(suite.ctx, address, func(index int64, request types.DTagTransferRequest) (stop bool) {
		iterations += 1
		return index == 1
	})
	suite.Require().Equal(iterations, 2)
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_IterateUserApplicationLinks() {
	address := "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"
	links := []types.ApplicationLink{
		types.NewApplicationLink(
			address,
			types.NewData("github", "github-user"),
			types.ApplicationLinkStateInitialized,
			types.NewOracleRequest(
				0,
				1,
				types.NewOracleRequestCallData("github", "call_data"),
				"client_id",
			),
			nil,
			time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		),
		types.NewApplicationLink(
			address,
			types.NewData("reddit", "reddit-user"),
			types.ApplicationLinkStateInitialized,
			types.NewOracleRequest(
				0,
				1,
				types.NewOracleRequestCallData("reddit", "call_data"),
				"client_id",
			),
			nil,
			time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		),
		types.NewApplicationLink(
			address,
			types.NewData("twitter", "twitter-user"),
			types.ApplicationLinkStateInitialized,
			types.NewOracleRequest(
				0,
				1,
				types.NewOracleRequestCallData("twitter", "call_data"),
				"client_id",
			),
			nil,
			time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		),
	}

	ctx, _ := suite.ctx.CacheContext()

	for _, link := range links {
		suite.ak.SetAccount(ctx, testutil.ProfileFromAddr(link.User))

		err := suite.k.SaveApplicationLink(ctx, link)
		suite.Require().NoError(err)
	}

	var iterated []types.ApplicationLink
	suite.k.IterateUserApplicationLinks(ctx, address, func(index int64, link types.ApplicationLink) (stop bool) {
		iterated = append(iterated, link)
		return index == 1
	})

	suite.Require().Equal([]types.ApplicationLink{links[0], links[1]}, iterated)
}

func (suite *KeeperTestSuite) TestKeeper_GetApplicationLinks() {
	address := "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"
	links := []types.ApplicationLink{
		types.NewApplicationLink(
			address,
			types.NewData("github", "github-user"),
			types.ApplicationLinkStateInitialized,
			types.NewOracleRequest(
				0,
				1,
				types.NewOracleRequestCallData("github", "call_data"),
				"client_id",
			),
			nil,
			time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		),
		types.NewApplicationLink(
			address,
			types.NewData("reddit", "reddit-user"),
			types.ApplicationLinkStateInitialized,
			types.NewOracleRequest(
				0,
				1,
				types.NewOracleRequestCallData("reddit", "call_data"),
				"client_id",
			),
			nil,
			time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		),
		types.NewApplicationLink(
			address,
			types.NewData("twitter", "twitter-user"),
			types.ApplicationLinkStateInitialized,
			types.NewOracleRequest(
				0,
				1,
				types.NewOracleRequestCallData("twitter", "call_data"),
				"client_id",
			),
			nil,
			time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		),
	}

	ctx, _ := suite.ctx.CacheContext()

	for _, link := range links {
		suite.ak.SetAccount(ctx, testutil.ProfileFromAddr(link.User))

		err := suite.k.SaveApplicationLink(ctx, link)
		suite.Require().NoError(err)
	}

	suite.Require().Equal(links, suite.k.GetApplicationLinks(ctx))
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_GetChainLinks() {
	pub1 := secp256k1.GenPrivKey().PubKey()
	pub2 := secp256k1.GenPrivKey().PubKey()

	testCases := []struct {
		name     string
		store    func(ctx sdk.Context)
		expLinks []types.ChainLink
	}{
		{
			name:     "non existent link returns empty array",
			expLinks: nil,
		},
		{
			name: "existent links returns all links",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(
					types.ChainLinksStoreKey("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos", "cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f"),
					types.MustMarshalChainLink(
						suite.cdc,
						types.NewChainLink(
							"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
							types.NewBech32Address("cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f", "cosmos"),
							types.NewProof(pub1, "signature", "plain_text"),
							types.NewChainConfig("cosmos"),
							time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
						),
					),
				)
				store.Set(
					types.ChainLinksStoreKey("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos", "cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs"),
					types.MustMarshalChainLink(
						suite.cdc,
						types.NewChainLink(
							"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
							types.NewBech32Address("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", "cosmos"),
							types.NewProof(pub2, "signature", "plain_text"),
							types.NewChainConfig("cosmos"),
							time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
						),
					),
				)
			},
			expLinks: []types.ChainLink{
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f", "cosmos"),
					types.NewProof(pub1, "signature", "plain_text"),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				),
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", "cosmos"),
					types.NewProof(pub2, "signature", "plain_text"),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				),
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			links := suite.k.GetChainLinks(ctx)
			suite.Require().Equal(tc.expLinks, links)
		})
	}
}
