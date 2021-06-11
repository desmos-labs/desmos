package keeper_test

import (
	"encoding/base64"
	"time"

	oracletypes "github.com/bandprotocol/chain/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func createResponsePacketData(
	clientID string, requestID int64, status oracletypes.ResolveStatus, result []byte,
) oracletypes.OracleResponsePacketData {
	return oracletypes.OracleResponsePacketData{
		ClientID:      clientID,
		RequestID:     oracletypes.RequestID(requestID),
		AnsCount:      1,
		RequestTime:   1,
		ResolveTime:   1,
		ResolveStatus: status,
		Result:        result,
	}
}

func (suite *KeeperTestSuite) Test_OnRecvApplicationLinkPacketData() {
	resultBz, err := base64.StdEncoding.DecodeString("AAAAAIBhMDBhN2Q1YmQ0NWU0MjYxNTY0NWZjYWViNGQ4MDBhZjIyNzA0ZTU0OTM3YWIyMzVlNWU1MGJlYmQzOGU4OGI3NjVmZGI2OTZjMjI3MTJjMGNhYjExNzY3NTZiNjM0NmNiYzExNDgxYzU0NGQxZjc4MjhjYjIzMzYyMGMwNjE3MwAAAAxyaWNtb250YWduaW4=")
	suite.Require().NoError(err)

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
				[]byte{},
			),
			shouldErr: true,
		},
		{
			name: "Resolve status expired updates connection properly",
			store: func(context sdk.Context) {
				link := types.NewApplicationLink(
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

				err := suite.k.SaveApplicationLink(suite.ctx, "user", link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_EXPIRED,
				nil,
			),
			shouldErr: false,
			expLink: types.NewApplicationLink(
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
			store: func(context sdk.Context) {
				link := types.NewApplicationLink(
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

				err := suite.k.SaveApplicationLink(suite.ctx, "user", link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_FAILURE,
				nil,
			),
			shouldErr: false,
			expLink: types.NewApplicationLink(
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

			name: "Resolve status success updates connection properly",
			store: func(context sdk.Context) {
				link := types.NewApplicationLink(
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

				err := suite.k.SaveApplicationLink(suite.ctx, "user", link)
				suite.Require().NoError(err)
			},
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				resultBz,
			),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				types.NewData("twitter", "user"),
				types.AppLinkStateVerificationSuccess,
				types.NewOracleRequest(
					1,
					1,
					types.NewOracleRequestCallData("twitter", "tweet-123456789"),
					"client_id",
				),
				types.NewSuccessResult(
					"ricmontagnin",
					"a00a7d5bd45e42615645fcaeb4d800af22704e54937ab235e5e50bebd38e88b765fdb696c22712c0cab1176756b6346cbc11481c544d1f7828cb233620c06173",
				),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
	}

	for _, uc := range tests {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()
			if uc.store != nil {
				uc.store(suite.ctx)
			}

			err := suite.k.OnRecvApplicationLinkPacketData(suite.ctx, uc.data)

			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				_, stored, err := suite.k.GetApplicationLinkByClientID(suite.ctx, uc.expLink.OracleRequest.ClientID)
				suite.Require().NoError(err)
				suite.Require().True(stored.Equal(uc.expLink))
			}
		})
	}
}
