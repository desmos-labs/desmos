package keeper_test

import (
	"encoding/base64"
	"strings"
	"time"

	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	oracletypes "github.com/desmos-labs/desmos/x/oracle/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func createRequestPacketData(clientID string) oracletypes.OracleRequestPacketData {
	return oracletypes.NewOracleRequestPacketData(
		clientID,
		1,
		nil,
		1,
		1,
		sdk.NewCoins(),
		"",
		1,
		1,
	)
}

func createResponsePacketData(
	clientID string, requestID int64, status oracletypes.ResolveStatus, result string,
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
		RequestID:     requestID,
		AnsCount:      1,
		RequestTime:   1,
		ResolveTime:   1,
		ResolveStatus: status,
		Result:        resultBz,
	}
}

func (suite *KeeperTestSuite) TestKeeper_OnRecvApplicationLinkPacketData() {
	tests := []struct {
		name      string
		store     func(sdk.Context)
		data      oracletypes.OracleResponsePacketData
		shouldErr bool
		expLink   types.ApplicationLink
	}{
		{
			name: "Non existing connection returns error",
			data: createResponsePacketData(
				"client_id",
				-1,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				"",
			),
			shouldErr: true,
		},
		{
			name: "Resolve status expired updates connection properly",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					suite.testData.profile.GetAddress().String(),
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

				suite.ak.SetAccount(ctx, suite.testData.profile)
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
				suite.testData.profile.GetAddress().String(),
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

			name: "Resolve status failure updates connection properly",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					suite.testData.profile.GetAddress().String(),
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

				suite.ak.SetAccount(ctx, suite.testData.profile)

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
				suite.testData.profile.GetAddress().String(),
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
			name: "Wrongly encoded result returns error",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					suite.testData.profile.GetAddress().String(),
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

				suite.ak.SetAccount(ctx, suite.testData.profile)
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
			name: "Different returned value (username) updates correctly",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					suite.testData.profile.GetAddress().String(),
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

				suite.ak.SetAccount(ctx, suite.testData.profile)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				"AAAAgDc0OWI2OWJiZjJlOTI2MDE1ZjVhZTVkOWRjODQxM2IyYjIxNDYzYzhmNjNhNDI4N2I2MjY0NTZhY2ViMzllNTEwOTA0ZTg2NDkyNTA1ZTgxYmM5ZDRjMzFmMzUwNDY4ZjM3MDY4OTFiNmI4M2UxYzVmMmY5N2JlMzU2MDJmODA0AAAADHJpY21vbnRhZ25pbg==",
			),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				suite.testData.profile.GetAddress().String(),
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
			name: "Wrongly encoded result signature error",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					suite.testData.profile.GetAddress().String(),
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

				suite.ak.SetAccount(ctx, suite.testData.profile)
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
			name: "Wrong signature updates connection properly",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					suite.testData.profile.GetAddress().String(),
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

				suite.ak.SetAccount(ctx, suite.testData.profile)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				"AAAAgDY1NTkwMDA2MWY5YTMwNmM2ODViYmJmNDQ2YTNjZDAyZjQ2OWY5OTVhMmVhZDVkZDY0YWUwYWMwZTkwMTYxYjQ1OGEzYTkxZGNlMzA4MGZiOTM1Yzk4NTg1Y2EyYzFlOTNiMTcyMmZmNTJjZGQ1YzU5ODQwZjQ1MTQzOGI4ZTJjAAAADHJpY21vbnRhZ25pbg==",
			),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				suite.testData.profile.GetAddress().String(),
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
			name: "Valid resolve status success updates connection properly",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					suite.testData.profile.GetAddress().String(),
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

				suite.ak.SetAccount(ctx, suite.testData.profile)
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				"AAAAgDc0OWI2OWJiZjJlOTI2MDE1ZjVhZTVkOWRjODQxM2IyYjIxNDYzYzhmNjNhNDI4N2I2MjY0NTZhY2ViMzllNTEwOTA0ZTg2NDkyNTA1ZTgxYmM5ZDRjMzFmMzUwNDY4ZjM3MDY4OTFiNmI4M2UxYzVmMmY5N2JlMzU2MDJmODA0AAAADHJpY21vbnRhZ25pbg==",
			),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				suite.testData.profile.GetAddress().String(),
				types.NewData("twitter", "ricmontagnin"),
				types.AppLinkStateVerificationSuccess,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewSuccessResult(
					"ricmontagnin",
					"749b69bbf2e926015f5ae5d9dc8413b2b21463c8f63a4287b626456aceb39e510904e86492505e81bc9d4c31f350468f3706891b6b83e1c5f2f97be35602f804",
				),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
	}

	for _, uc := range tests {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			err := suite.k.OnRecvApplicationLinkPacketData(ctx, uc.data)

			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				stored, err := suite.k.GetApplicationLinkByClientID(ctx, uc.expLink.OracleRequest.ClientID)
				suite.Require().NoError(err)
				suite.Require().Truef(uc.expLink.Equal(stored), "%s\n%s", uc.expLink, stored)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_OnOracleRequestAcknowledgementPacket() {
	usecases := []struct {
		name      string
		store     func(ctx sdk.Context)
		data      oracletypes.OracleRequestPacketData
		ack       channeltypes.Acknowledgement
		shouldErr bool
		verify    func(ctx sdk.Context)
	}{
		{
			name:      "invalid client id returns error",
			data:      createRequestPacketData("client_id"),
			ack:       channeltypes.NewErrorAcknowledgement("error"),
			shouldErr: true,
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			err := suite.k.OnOracleRequestAcknowledgementPacket(ctx, uc.data, uc.ack)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_OnOracleRequestTimeoutPacket() {
	usecases := []struct {
		name      string
		store     func(ctx sdk.Context)
		data      oracletypes.OracleRequestPacketData
		shouldErr bool
		verify    func(ctx sdk.Context)
	}{
		{
			name:      "invalid client id request returns error",
			data:      createRequestPacketData("client_id"),
			shouldErr: true,
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
						-1,
						1,
						types.NewOracleRequestCallData("twitter", "call_data"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))
				suite.Require().NoError(suite.k.SaveApplicationLink(ctx, link))
			},
			data:      createRequestPacketData("client_id"),
			shouldErr: false,
			verify: func(ctx sdk.Context) {
				link, err := suite.k.GetApplicationLinkByClientID(ctx, "client_id")
				suite.Require().NoError(err)

				expected := types.NewApplicationLink(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewData("reddit", "reddit-user"),
					types.AppLinkStateVerificationTimedOut,
					types.NewOracleRequest(
						-1,
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

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			err := suite.k.OnOracleRequestTimeoutPacket(ctx, uc.data)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			if uc.verify != nil {
				uc.verify(ctx)
			}
		})
	}
}
