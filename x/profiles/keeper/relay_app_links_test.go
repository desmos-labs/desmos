package keeper_test

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
	"time"

	"github.com/desmos-labs/desmos/v2/testutil"

	"github.com/desmos-labs/desmos/v2/pkg/obi"

	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/modules/core/exported"

	"github.com/desmos-labs/desmos/v2/testutil/ibctesting"

	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	oracletypes "github.com/desmos-labs/desmos/v2/x/oracle/types"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func createRequestPacketData(clientID string) oracletypes.OracleRequestPacketData {
	return oracletypes.NewOracleRequestPacketData(
		clientID,
		1,
		nil,
		1,
		1,
		sdk.NewCoins(),
		1,
		1,
	)
}

func createResponsePacketData(
	clientID string, requestID uint64, status oracletypes.ResolveStatus, result string,
) oracletypes.OracleResponsePacketData {
	var resultBz []byte

	if strings.TrimSpace(result) != "" {
		bz, err := base64.StdEncoding.DecodeString(result)
		if err != nil {
			panic(err)
		}
		resultBz = bz
	}

	return oracletypes.OracleResponsePacketData{
		ClientID:      clientID,
		RequestID:     oracletypes.RequestID(requestID),
		AnsCount:      1,
		RequestTime:   1,
		ResolveTime:   1,
		ResolveStatus: status,
		Result:        resultBz,
	}
}

func (suite *KeeperTestSuite) TestKeeper_StartProfileConnection() {
	var (
		applicationData    types.Data
		callData           string
		channelA, channelB ibctesting.TestChannel
		err                error
	)

	testCases := []struct {
		name        string
		malleate    func()
		storeChainA func(ctx sdk.Context)
		expPass     bool
	}{
		{
			name: "source channel not found",
			malleate: func() {
				// channel references wrong ID
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, _ = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				channelA.ID = "IDisInvalid"
			},
			expPass: false,
		},
		{
			name: "next seq send not found",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA = suite.chainA.NextTestChannel(connA, types.IBCPortID)
				channelB = suite.chainB.NextTestChannel(connB, types.IBCPortID)

				// manually create channel so next seq send is never set
				suite.chainA.App.IBCKeeper.ChannelKeeper.SetChannel(
					suite.chainA.GetContext(),
					channelA.PortID, channelA.ID,
					channeltypes.NewChannel(
						channeltypes.OPEN,
						channeltypes.ORDERED,
						channeltypes.NewCounterparty(channelB.PortID, channelB.ID),
						[]string{connA.ID},
						"ics-20",
					),
				)
				suite.chainA.CreateChannelCapability(channelA.PortID, channelA.ID)
			},
			expPass: false,
		},
		{
			name: "channel capability not found",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				capability := suite.chainA.GetChannelCapability(channelA.PortID, channelA.ID)

				// Release channel capability
				err := suite.chainA.App.ScopedProfilesKeeper.ReleaseCapability(suite.chainA.GetContext(), capability)
				suite.Require().NoError(err)
			},
			expPass: false,
		},
		{
			name: "send without profile returns error",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				applicationData = types.NewData("twitter", "twitteruser")
				callData = "call_data"
			},
			expPass: false,
		},
		{
			name: "send with profile works properly",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				applicationData = types.NewData("twitter", "twitteruser")
				callData = "call_data"
			},
			storeChainA: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr(suite.chainA.Account.GetAddress().String())
				suite.chainA.App.AccountKeeper.SetAccount(ctx, profile)
			},
			expPass: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupIBCTest()

			tc.malleate()
			if tc.storeChainA != nil {
				tc.storeChainA(suite.chainA.GetContext())
			}

			err = suite.chainA.App.ProfileKeeper.StartProfileConnection(
				suite.chainA.GetContext(), applicationData, callData, suite.chainA.Account.GetAddress(),
				channelA.PortID, channelA.ID,
				clienttypes.NewHeight(0, 110), 0,
			)

			if tc.expPass {
				suite.Require().NoError(err)

				links := suite.chainA.App.ProfileKeeper.GetApplicationLinks(suite.chainA.GetContext())
				suite.Require().Len(links, 1)

				suite.Require().Equal(suite.chainA.Account.GetAddress().String(), links[0].User)
				suite.Require().Equal(types.ApplicationLinkStateInitialized, links[0].State)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_OnRecvApplicationLinkPacketData() {
	profile := suite.GetRandomProfile()
	username := "twitter-profile"
	hexValue := hex.EncodeToString([]byte(username))
	hexSig := hex.EncodeToString(profile.Sign([]byte("twitter-profile")))

	type resultData struct {
		Signature string `obi:"signature"`
		Value     string `obi:"value"`
		Username  string `obi:"username"`
	}
	result, err := obi.Encode(resultData{
		Signature: hexSig,
		Value:     hexValue,
		Username:  username,
	})
	suite.Require().NoError(err)
	resultBase64 := base64.StdEncoding.EncodeToString(result)

	testCases := []struct {
		name      string
		store     func(sdk.Context)
		data      oracletypes.OracleResponsePacketData
		shouldErr bool
		expLink   types.ApplicationLink
	}{
		{
			name: "non existing link returns no error",
			data: createResponsePacketData(
				"client_id",
				0,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				"",
			),
			shouldErr: false,
		},
		{
			name: "resolve status expired updates connection properly",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				suite.ak.SetAccount(ctx, profile)

				link := types.NewApplicationLink(
					profile.GetAddress().String(),
					types.NewData("twitter", "user"),
					types.AppLinkStateVerificationStarted,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_EXPIRED,
				"",
			),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewData("twitter", "user"),
				types.AppLinkStateVerificationError,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewErrorResult(types.ErrRequestExpired),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
		{
			name: "resolve status failure updates connection properly",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				suite.ak.SetAccount(ctx, profile)

				link := types.NewApplicationLink(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					types.NewData("twitter", "user"),
					types.AppLinkStateVerificationStarted,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_FAILURE,
				"",
			),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewData("twitter", "user"),
				types.AppLinkStateVerificationError,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewErrorResult(types.ErrRequestFailed),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
		{
			name: "wrongly encoded result returns error",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				suite.ak.SetAccount(ctx, profile)

				link := types.NewApplicationLink(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					types.NewData("twitter", "user"),
					types.AppLinkStateVerificationStarted,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				"dGVzdA==",
			),
			shouldErr: true,
		},
		{
			name: "different returned value (username) updates correctly",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				suite.ak.SetAccount(ctx, profile)

				link := types.NewApplicationLink(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					types.NewData("twitter", "user"),
					types.AppLinkStateVerificationStarted,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				"AAAAgDY1NTkwMDA2MWY5YTMwNmM2ODViYmJmNDQ2YTNjZDAyZjQ2OWY5OTVhMmVhZDVkZDY0YWUwYWMwZTkwMTYxYjQ1OGEzYTkxZGNlMzA4MGZiOTM1Yzk4NTg1Y2EyYzFlOTNiMTcyMmZmNTJjZGQ1YzU5ODQwZjQ1MTQzOGI4ZTJjAAAAGDcyNjk2MzZkNmY2ZTc0NjE2NzZlNjk2ZQAAAAxhbm90aGVyX3VzZXI=",
			),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewData("twitter", "user"),
				types.AppLinkStateVerificationError,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewErrorResult(types.ErrInvalidAppUsername),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
		{
			name: "wrongly encoded result signature error",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				suite.ak.SetAccount(ctx, profile)

				link := types.NewApplicationLink(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					types.NewData("twitter", "ricmontagnin"),
					types.AppLinkStateVerificationStarted,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				"AAAACXNpZ25hdHVyZQAAAAxyaWNtb250YWduaW4=",
			),
			shouldErr: true,
		},
		{
			name: "wrong signature updates connection properly",
			store: func(ctx sdk.Context) {
				suite.ak.SetAccount(ctx, profile.Profile)

				link := types.NewApplicationLink(
					profile.GetAddress().String(),
					types.NewData("twitter", "ricmontagnin"),
					types.AppLinkStateVerificationStarted,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				"AAAAgDY1NTkwMDA2MWY5YTMwNmM2ODViYmJmNDQ2YTNjZDAyZjQ2OWY5OTVhMmVhZDVkZDY0YWUwYWMwZTkwMTYxYjQ1OGEzYTkxZGNlMzA4MGZiOTM1Yzk4NTg1Y2EyYzFlOTNiMTcyMmZmNTJjZGQ1YzU5ODQwZjQ1MTQzOGI4ZTJjAAAAGDcyNjk2MzZkNmY2ZTc0NjE2NzZlNjk2ZQAAAAxyaWNtb250YWduaW4=",
			),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				profile.GetAddress().String(),
				types.NewData("twitter", "ricmontagnin"),
				types.AppLinkStateVerificationError,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewErrorResult(types.ErrInvalidSignature),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
		{
			name: "valid resolve status success updates connection properly",
			store: func(ctx sdk.Context) {
				suite.ak.SetAccount(ctx, profile.Profile)

				link := types.NewApplicationLink(
					profile.GetAddress().String(),
					types.NewData("twitter", username),
					types.AppLinkStateVerificationStarted,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data:      createResponsePacketData("client_id", 1, oracletypes.RESOLVE_STATUS_SUCCESS, resultBase64),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				profile.GetAddress().String(),
				types.NewData("twitter", username),
				types.AppLinkStateVerificationSuccess,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewSuccessResult(hexValue, hexSig),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
		{
			name: "timed out link does not get updated",
			store: func(ctx sdk.Context) {
				suite.ak.SetAccount(ctx, profile.Profile)

				link := types.NewApplicationLink(
					profile.GetAddress().String(),
					types.NewData("twitter", username),
					types.AppLinkStateVerificationTimedOut,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data:      createResponsePacketData("client_id", 1, oracletypes.RESOLVE_STATUS_SUCCESS, resultBase64),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				profile.GetAddress().String(),
				types.NewData("twitter", username),
				types.AppLinkStateVerificationTimedOut,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
		{
			name: "errored link does not get updated",
			store: func(ctx sdk.Context) {
				suite.ak.SetAccount(ctx, profile.Profile)

				link := types.NewApplicationLink(
					profile.GetAddress().String(),
					types.NewData("twitter", username),
					types.AppLinkStateVerificationError,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					types.NewErrorResult(types.ErrInvalidSignature),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data:      createResponsePacketData("client_id", 1, oracletypes.RESOLVE_STATUS_SUCCESS, resultBase64),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				profile.GetAddress().String(),
				types.NewData("twitter", username),
				types.AppLinkStateVerificationError,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewErrorResult(types.ErrInvalidSignature),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
		{
			name: "already verified link does not get updated",
			store: func(ctx sdk.Context) {
				suite.ak.SetAccount(ctx, profile.Profile)

				link := types.NewApplicationLink(
					profile.GetAddress().String(),
					types.NewData("twitter", username),
					types.AppLinkStateVerificationSuccess,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					types.NewSuccessResult("value", "signature"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data:      createResponsePacketData("client_id", 1, oracletypes.RESOLVE_STATUS_SUCCESS, resultBase64),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				profile.GetAddress().String(),
				types.NewData("twitter", username),
				types.AppLinkStateVerificationSuccess,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewSuccessResult("value", "signature"),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			err = suite.k.OnRecvApplicationLinkPacketData(ctx, tc.data)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				stored, _, err := suite.k.GetApplicationLinkByClientID(ctx, tc.expLink.OracleRequest.ClientID)
				suite.Require().NoError(err)
				suite.Require().Truef(tc.expLink.Equal(stored), "%s\n%s", tc.expLink, stored)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_OnOracleRequestAcknowledgementPacket() {
	result := oracletypes.OracleRequestPacketAcknowledgement{RequestID: 1000}

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		data      oracletypes.OracleRequestPacketData
		ack       channeltypes.Acknowledgement
		shouldErr bool
		expLink   types.ApplicationLink
	}{
		{
			name:      "non existing link returns no error",
			data:      createRequestPacketData("client_id"),
			ack:       channeltypes.NewErrorAcknowledgement("error"),
			shouldErr: false,
		},
		{
			name: "acknowledgment error updates link properly",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.ak.SetAccount(ctx, testutil.ProfileFromAddr(address))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data:      createRequestPacketData("client_id"),
			ack:       channeltypes.NewErrorAcknowledgement("error"),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.NewData("twitter", "twitteruser"),
				types.AppLinkStateVerificationError,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				types.NewErrorResult("error"),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
		{
			name: "invalid acknowledgment result returns error",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.ak.SetAccount(ctx, testutil.ProfileFromAddr(address))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data:      createRequestPacketData("client_id"),
			ack:       channeltypes.NewResultAcknowledgement([]byte("error")),
			shouldErr: true,
		},
		{
			name: "acknowledgment result updates link properly",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.ak.SetAccount(ctx, testutil.ProfileFromAddr(address))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data:      createRequestPacketData("client_id"),
			ack:       channeltypes.NewResultAcknowledgement(types.ModuleCdc.MustMarshalJSON(&result)),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.NewData("twitter", "twitteruser"),
				types.AppLinkStateVerificationStarted,
				types.NewOracleRequest(
					1000,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			err := suite.k.OnOracleRequestAcknowledgementPacket(ctx, tc.data, tc.ack)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				link, _, err := suite.k.GetApplicationLinkByClientID(ctx, tc.data.ClientID)
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expLink, link)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_OnOracleRequestTimeoutPacket() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		data      oracletypes.OracleRequestPacketData
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name:      "not found link returns no error",
			data:      createRequestPacketData("client_id"),
			shouldErr: false,
		},
		{
			name: "valid client id updates the link properly",
			store: func(ctx sdk.Context) {
				address := "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"
				link := types.NewApplicationLink(
					address,
					types.NewData("reddit", "reddit-user"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "call_data"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.ak.SetAccount(ctx, testutil.ProfileFromAddr(address))
				suite.Require().NoError(suite.k.SaveApplicationLink(ctx, link))
			},
			data:      createRequestPacketData("client_id"),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				link, _, err := suite.k.GetApplicationLinkByClientID(ctx, "client_id")
				suite.Require().NoError(err)

				expected := types.NewApplicationLink(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewData("reddit", "reddit-user"),
					types.AppLinkStateVerificationTimedOut,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "call_data"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				suite.Require().Equal(expected, link)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			err := suite.k.OnOracleRequestTimeoutPacket(ctx, tc.data)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
