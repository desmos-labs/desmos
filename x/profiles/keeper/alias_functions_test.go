package keeper_test

import (
	"time"

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
	tests := []struct {
		name     string
		accounts []*types.Profile
	}{
		{
			name:     "Non empty Profiles list returned",
			accounts: []*types.Profile{suite.testData.profile},
		},
		{
			name:     "Profile not found",
			accounts: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, profile := range test.accounts {
				err := suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)
			}

			res := suite.k.GetProfiles(suite.ctx)
			suite.Require().Equal(test.accounts, res)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_IterateUserApplicationLinks() {
	address := "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"
	links := []types.ApplicationLink{
		types.NewApplicationLink(
			address,
			types.NewData("github", "github-user"),
			types.ApplicationLinkStateInitialized,
			types.NewOracleRequest(
				-1,
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
				-1,
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
				-1,
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
		suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(link.User))

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

func (suite *KeeperTestSuite) TestKeeper_GetApplicationLinksEntries() {
	address := "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"
	links := []types.ApplicationLink{
		types.NewApplicationLink(
			address,
			types.NewData("github", "github-user"),
			types.ApplicationLinkStateInitialized,
			types.NewOracleRequest(
				-1,
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
				-1,
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
				-1,
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
		suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(link.User))

		err := suite.k.SaveApplicationLink(ctx, link)
		suite.Require().NoError(err)
	}

	suite.Require().Equal(links, suite.k.GetApplicationLinks(ctx))
}
