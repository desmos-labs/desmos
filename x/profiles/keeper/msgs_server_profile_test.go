package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/x/profiles/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_handleMsgSaveProfile() {
	tests := []struct {
		name              string
		existentProfiles  []*types.Profile
		blockTime         time.Time
		msg               *types.MsgSaveProfile
		shouldErr         bool
		expEvents         sdk.Events
		expStoredProfiles []*types.Profile
	}{
		{
			name:      "Profile saved (with no previous profile created)",
			blockTime: suite.testData.profile.CreationDate,
			msg: types.NewMsgSaveProfile(
				"custom_dtag",
				"my-nickname",
				"my-bio",
				"https://test.com/profile-picture",
				"https://test.com/cover-pic",
				suite.testData.profile.GetAddress().String(),
			),
			expStoredProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"custom_dtag",
					"my-nickname",
					"my-bio",
					types.NewPictures(
						"https://test.com/profile-picture",
						"https://test.com/cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileSaved,
					sdk.NewAttribute(types.AttributeProfileDTag, "custom_dtag"),
					sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.GetAddress().String()),
					sdk.NewAttribute(types.AttributeProfileCreationTime, suite.testData.profile.CreationDate.Format(time.RFC3339)),
				),
			},
		},
		{
			name:      "Profile saved (with previous profile created)",
			blockTime: suite.testData.profile.CreationDate,
			existentProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"test_dtag",
					"old-nickname",
					"old-biography",
					types.NewPictures(
						"https://test.com/old-profile-pic",
						"https://test.com/old-cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
			msg: types.NewMsgSaveProfile(
				"other_dtag",
				"nickname",
				"biography",
				"https://test.com/profile-pic",
				"https://test.com/cover-pic",
				suite.testData.profile.GetAddress().String(),
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileSaved,
					sdk.NewAttribute(types.AttributeProfileDTag, "other_dtag"),
					sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.GetAddress().String()),
					sdk.NewAttribute(types.AttributeProfileCreationTime, suite.testData.profile.CreationDate.Format(time.RFC3339)),
				),
			},
			expStoredProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"other_dtag",
					"nickname",
					"biography",
					types.NewPictures(
						"https://test.com/profile-pic",
						"https://test.com/cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
		},
		{
			name:      "Profile saved with same DTag but capital first letter (with previous profile created)",
			blockTime: suite.testData.profile.CreationDate,
			existentProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"test",
					"old-nickname",
					"old-biography",
					types.NewPictures(
						"https://test.com/old-profile-pic",
						"https://test.com/old-cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
			msg: types.NewMsgSaveProfile(
				"Test",
				"nickname",
				"biography",
				"https://test.com/profile-pic",
				"https://test.com/cover-pic",
				suite.testData.profile.GetAddress().String(),
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileSaved,
					sdk.NewAttribute(types.AttributeProfileDTag, "Test"),
					sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.GetAddress().String()),
					sdk.NewAttribute(types.AttributeProfileCreationTime, suite.testData.profile.CreationDate.Format(time.RFC3339)),
				),
			},
			expStoredProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"Test",
					"nickname",
					"biography",
					types.NewPictures(
						"https://test.com/profile-pic",
						"https://test.com/cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
		},
		{
			name:      "Profile not saved because of the same DTag",
			blockTime: suite.testData.profile.CreationDate,
			existentProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"test",
					"nickname",
					"biography",
					types.NewPictures(
						"https://test.com/profile-pic",
						"https://test.com/cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
			msg: types.NewMsgSaveProfile(
				"Test",
				"another-one",
				"biography",
				"https://test.com/profile-pic",
				"https://test.com/cover-pic",
				suite.testData.otherUser,
			),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
			expStoredProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"test",
					"nickname",
					"biography",
					types.NewPictures(
						"https://test.com/profile-pic",
						"https://test.com/cover-pic",
					),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
		},
		{
			name:      "Profile not edited because of the invalid profile picture",
			blockTime: suite.testData.profile.CreationDate,
			existentProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"custom_dtag",
					"biography",
					"",
					types.NewPictures("", ""),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
			msg: types.NewMsgSaveProfile(
				"custom_dtag",
				"",
				"",
				"invalid-pic",
				"",
				suite.testData.profile.GetAddress().String(),
			),
			expEvents: sdk.EmptyEvents(),
			shouldErr: true,
			expStoredProfiles: []*types.Profile{
				suite.CheckProfileNoError(types.NewProfile(
					"custom_dtag",
					"biography",
					"",
					types.NewPictures("", ""),
					suite.testData.profile.CreationDate,
					suite.testData.profile.GetAccount(),
				)),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.k.SetParams(suite.ctx, types.DefaultParams())

			suite.ctx = suite.ctx.WithBlockTime(test.blockTime)
			for _, acc := range test.existentProfiles {
				err := suite.k.StoreProfile(suite.ctx, acc)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.SaveProfile(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
			}

			stored := suite.k.GetProfiles(suite.ctx)
			suite.Require().Len(stored, len(test.expStoredProfiles))
			for _, profile := range stored {
				suite.Require().Contains(test.expStoredProfiles, profile)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgDeleteProfile() {
	tests := []struct {
		name           string
		storedProfiles []*types.Profile
		msg            *types.MsgDeleteProfile
		expErr         error
		expEvents      sdk.Events
	}{
		{
			name:           "Profile doesn't exists",
			storedProfiles: nil,
			msg:            types.NewMsgDeleteProfile(suite.testData.profile.GetAddress().String()),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"no profile associated with the following address found: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "Profile deleted successfully",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			msg: types.NewMsgDeleteProfile(suite.testData.profile.GetAddress().String()),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileDeleted,
					sdk.NewAttribute(types.AttributeProfileCreator, suite.testData.profile.GetAddress().String()),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, profile := range test.storedProfiles {
				err := suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.DeleteProfile(sdk.WrapSDKContext(suite.ctx), test.msg)

			suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

			if test.expErr != nil {
				suite.Require().Error(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}
		})
	}
}
