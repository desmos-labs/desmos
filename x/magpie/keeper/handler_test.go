package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/magpie/keeper"
	"github.com/desmos-labs/desmos/x/magpie/types"
	"github.com/stretchr/testify/require"
)

func (suite *KeeperTestSuite) Test_handleMsgCreateSession() {
	owner, err := sdk.AccAddressFromBech32("cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0")
	suite.NoError(err)

	testData := []struct {
		name  string
		msg   types.MsgCreateSession
		error error
	}{
		{
			name: "Empty signature returns error",
			msg: types.MsgCreateSession{
				Owner:         suite.testData.owner,
				Namespace:     suite.testData.session.Namespace,
				ExternalOwner: suite.testData.session.ExternalOwner,
				PubKey:        suite.testData.session.PubKey,
				Signature:     "",
			},
			error: sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "the session signature is not valid"),
		},
		{
			name: "Invalid signature returns error",
			msg: types.MsgCreateSession{
				Owner:     suite.testData.owner,
				PubKey:    "ArDhBMh0X/3Akfc58oF1zFE00L/rLpgMMVvmcj0QlaN1",
				Signature: "3KXX5DmlsDAyO0pmgDT3pTyyuTfGr9ocJCOcaPwZDilAiwAp6U9egpHr1qOtx4dLLrtIVWE8npHK49BKKyyacg==",
			},
			error: sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "the session signature is not valid"),
		},
		{
			name: "Valid signature works properly",
			msg: types.MsgCreateSession{
				Owner:         owner,
				Namespace:     "cosmoshub-2",
				ExternalOwner: "cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0",
				PubKey:        "ArDhBMh0X/3Akfc58oF1zFE00L/rLpgMMVvmcj0QlaN1",
				Signature:     "3KXX5DmlsDAyO0pmgDT3pTyyuTfGr9ocJCOcaPwZDilAiwAp6U9egpHr1qOtx4dLLrtIVWE8npHK49BKKyyacg==",
			},
		},
	}

	for _, test := range testData {
		test := test
		suite.Run(test.name, func() {
			sessionLength := int64(240)
			err := suite.keeper.SetDefaultSessionLength(suite.ctx, sessionLength)
			suite.NoError(err)

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Valid response
			if res != nil {
				// Check the stored session
				expectedID := suite.keeper.GetLastSessionID(suite.ctx)
				session := types.Session{
					SessionID:     expectedID,
					Created:       suite.ctx.BlockHeight(),
					Expiry:        suite.ctx.BlockHeight() + sessionLength,
					Owner:         test.msg.Owner,
					Namespace:     test.msg.Namespace,
					ExternalOwner: test.msg.ExternalOwner,
					PubKey:        test.msg.PubKey,
					Signature:     test.msg.Signature,
				}

				var stored types.Session
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.SessionStoreKey(expectedID)), &stored)
				suite.Equal(session, stored)

				// Check the events
				creationEvent := sdk.NewEvent(
					types.EventTypeCreateSession,
					sdk.NewAttribute(types.AttributeKeySessionID, session.SessionID.String()),
					sdk.NewAttribute(types.AttributeKeyNamespace, session.Namespace),
					sdk.NewAttribute(types.AttributeKeyExternalOwner, session.ExternalOwner),
					sdk.NewAttribute(types.AttributeKeyExpiry, strconv.FormatInt(session.Expiry, 10)),
				)

				suite.NotNil(res)
				suite.Contains(res.Events, creationEvent)
			}

			// Invalid response
			if res == nil {
				suite.NotNil(err)
				suite.Equal(err.Error(), test.error.Error())
			}
		})
	}
}
