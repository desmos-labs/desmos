package keeper_test

import (
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v2/testutil"
	"github.com/desmos-labs/desmos/v2/x/profiles/keeper"
	v200 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v200"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func (suite KeeperTestSuite) saveAppLinkV200(link v200.ApplicationLink) error {
	if !suite.k.HasProfile(suite.ctx, link.User) {
		return sdkerrors.Wrapf(types.ErrProfileNotFound, "a profile is required to link an application")
	}

	// Get the keys
	userApplicationLinkKey := types.UserApplicationLinkKey(link.User, link.Data.Application, link.Data.Username)
	applicationLinkClientIDKey := types.ApplicationLinkClientIDKey(link.OracleRequest.ClientID)

	// Store the data
	store := suite.ctx.KVStore(suite.k.StoreKey)
	store.Set(userApplicationLinkKey, v200.MustMarshalApplicationLink(suite.k.Cdc, link))
	store.Set(applicationLinkClientIDKey, userApplicationLinkKey)

	return nil
}

func (suite KeeperTestSuite) Test_Migrate4to5() {
	suite.ctx = suite.ctx.WithBlockTime(time.Date(2022, 5, 0, 00, 00, 00, 000, time.UTC))
	ctx, _ := suite.ctx.CacheContext()

	suite.k.SetParams(ctx, types.NewParams(
		types.DefaultNicknameParams(),
		types.DefaultDTagParams(),
		types.DefaultBioParams(),
		types.DefaultOracleParams(),
		types.NewAppLinksParams(time.Date(1, 1, 0, 00, 00, 00, 000, time.UTC)),
	))

	profile := testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
	suite.Require().NoError(suite.k.StoreProfile(ctx, profile))

	v200Links := []v200.ApplicationLink{
		v200.NewApplicationLink(
			"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			v200.NewData("twitter", "twitteruser"),
			v200.ApplicationLinkStateInitialized,
			v200.NewOracleRequest(
				0,
				1,
				v200.NewOracleRequestCallData("twitter", "calldata"),
				"client_id",
			),
			nil,
			time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		),
		v200.NewApplicationLink(
			"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			v200.NewData("reddit", "reddituser"),
			v200.ApplicationLinkStateInitialized,
			v200.NewOracleRequest(
				0,
				1,
				v200.NewOracleRequestCallData("reddit", "calldata"),
				"client_id2",
			),
			nil,
			time.Date(2022, 5, 1, 00, 00, 00, 000, time.UTC),
		),
	}

	expectedAppLink := types.NewApplicationLink(
		"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
		types.NewData("reddit", "reddituser"),
		types.ApplicationLinkStateInitialized,
		types.NewOracleRequest(
			0,
			1,
			types.NewOracleRequestCallData("reddit", "calldata"),
			"client_id2",
		),
		nil,
		time.Date(2022, 5, 1, 00, 00, 00, 000, time.UTC),
		time.Date(2023, 6, 1, 00, 00, 00, 000, time.UTC),
	)

	suite.ctx = ctx
	for _, link := range v200Links {
		err := suite.saveAppLinkV200(link)
		suite.Require().NoError(err)
	}

	migrator := keeper.NewMigrator(suite.k, suite.legacyAminoCdc, nil)

	err := migrator.Migrate4to5(suite.ctx)
	suite.Require().NoError(err)

	migratedAppLinks := suite.k.GetApplicationLinks(suite.ctx)

	suite.Require().Equal([]types.ApplicationLink{expectedAppLink}, migratedAppLinks)
}
