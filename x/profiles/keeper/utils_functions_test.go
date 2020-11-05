package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_IterateProfile() {
	date, err := time.Parse(time.RFC3339, "2010-10-02T12:10:00.000Z")
	suite.Require().NoError(err)

	profiles := []types.Profile{
		types.NewProfile(
			"first",
			"",
			"",
			types.NewPictures("", ""),
			date,
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		),
		types.NewProfile(
			"second",
			"",
			"",
			types.NewPictures("", ""),
			date,
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		),
		types.NewProfile(
			"not",
			"",
			"",
			types.NewPictures("", ""),
			date,
			"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
		),
		types.NewProfile(
			"third",
			"",
			"",
			types.NewPictures("", ""),
			date,
			"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
		),
	}

	expProfiles := []types.Profile{
		profiles[0],
		profiles[1],
		profiles[3],
	}

	for _, profile := range profiles {
		err := suite.k.StoreProfile(suite.ctx, profile)
		suite.Require().NoError(err)
	}

	var validProfiles []types.Profile
	suite.k.IterateProfiles(suite.ctx, func(_ int64, profile types.Profile) (stop bool) {
		if profile.Dtag == "not" {
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
