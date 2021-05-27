package keeper_test

import (
	"encoding/base64"
	"time"

	oracletypes "github.com/bandprotocol/chain/x/oracle/types"

	"github.com/desmos-labs/desmos/x/ibc/applications/profiles/types"
)

func createResponsePacketData(clientID string, requestID int64, status oracletypes.ResolveStatus, result []byte) oracletypes.OracleResponsePacketData {
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

func (suite *KeeperTestSuite) Test_OnRecvPacket() {
	resultBz, err := base64.StdEncoding.DecodeString("AAAAAIBhMDBhN2Q1YmQ0NWU0MjYxNTY0NWZjYWViNGQ4MDBhZjIyNzA0ZTU0OTM3YWIyMzVlNWU1MGJlYmQzOGU4OGI3NjVmZGI2OTZjMjI3MTJjMGNhYjExNzY3NTZiNjM0NmNiYzExNDgxYzU0NGQxZjc4MjhjYjIzMzYyMGMwNjE3MwAAAAxyaWNtb250YWduaW4=")
	suite.Require().NoError(err)

	tests := []struct {
		name               string
		existingConnection *types.Connection
		data               oracletypes.OracleResponsePacketData
		shouldErr          bool
		storedConnection   *types.Connection
	}{
		{
			name:               "non existing connection returns error",
			existingConnection: nil,
			data: createResponsePacketData(
				"client_id",
				-1,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				[]byte{},
			),
			shouldErr: true,
		},
		{

			name: "resolve status expired updates connection properly",
			existingConnection: types.NewConnection(
				"user",
				types.NewApplicationData("twitter", "user"),
				types.NewVerificationData("tweet", "123456789"),
				types.CONNECTION_STATE_STARTED,
				types.NewOracleRequest(1, 1, "client_id"),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_EXPIRED,
				nil,
			),
			shouldErr: false,
			storedConnection: types.NewConnection(
				"user",
				types.NewApplicationData("twitter", "user"),
				types.NewVerificationData("tweet", "123456789"),
				types.CONNECTION_STATE_ERROR,
				types.NewOracleRequest(1, 1, "client_id"),
				types.NewErrorResult(types.ErrRequestExpired),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
		{

			name: "resolve status failure updates connection properly",
			existingConnection: types.NewConnection(
				"user",
				types.NewApplicationData("twitter", "user"),
				types.NewVerificationData("tweet", "123456789"),
				types.CONNECTION_STATE_STARTED,
				types.NewOracleRequest(1, 1, "client_id"),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_FAILURE,
				nil,
			),
			shouldErr: false,
			storedConnection: types.NewConnection(
				"user",
				types.NewApplicationData("twitter", "user"),
				types.NewVerificationData("tweet", "123456789"),
				types.CONNECTION_STATE_ERROR,
				types.NewOracleRequest(1, 1, "client_id"),
				types.NewErrorResult(types.ErrRequestFailed),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
		{

			name: "resolve status success updates connection properly",
			existingConnection: types.NewConnection(
				"user",
				types.NewApplicationData("twitter", "user"),
				types.NewVerificationData("tweet", "123456789"),
				types.CONNECTION_STATE_STARTED,
				types.NewOracleRequest(1, 1, "client_id"),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			data: createResponsePacketData(
				"client_id",
				1,
				oracletypes.RESOLVE_STATUS_SUCCESS,
				resultBz,
			),
			shouldErr: false,
			storedConnection: types.NewConnection(
				"user",
				types.NewApplicationData("twitter", "user"),
				types.NewVerificationData("tweet", "123456789"),
				types.CONNECTION_STATE_SUCCESS,
				types.NewOracleRequest(1, 1, "client_id"),
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

			if uc.existingConnection != nil {
				err := suite.k.SaveApplicationLink(suite.ctx, uc.existingConnection)
				suite.Require().NoError(err)
			}

			err := suite.k.OnRecvPacket(suite.ctx, uc.data)

			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				stored, err := suite.k.GetApplicationLinkByClientID(suite.ctx, uc.existingConnection.OracleRequest.ClientId)
				suite.Require().NoError(err)
				suite.Require().True(stored.Equal(uc.storedConnection))
			}
		})
	}
}
