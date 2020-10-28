package keeper_test

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/magpie/keeper"
	"github.com/desmos-labs/desmos/x/magpie/types"
)

func (suite *KeeperTestSuite) Test_handleMsgCreateSession() {
	testData := []struct {
		name  string
		msg   *types.MsgCreateSession
		error error
	}{
		{
			name: "Empty signature returns error",
			msg: types.NewMsgCreateSession(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"",
				"",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "the session signature is not valid"),
		},
		{
			name: "Invalid signature returns error",
			msg: types.NewMsgCreateSession(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"ArDhBMh0X/3Akfc58oF1zFE00L/rLpgMMVvmcj0QlaN1",
				"3KXX5DmlsDAyO0pmgDT3pTyyuTfGr9ocJCOcaPwZDilAiwAp6U9egpHr1qOtx4dLLrtIVWE8npHK49BKKyyacg==",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "the session signature is not valid"),
		},
		{
			name: "Valid signature works properly",
			msg: types.NewMsgCreateSession(
				"cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0",
				"cosmoshub-2",
				"cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0",
				"ArDhBMh0X/3Akfc58oF1zFE00L/rLpgMMVvmcj0QlaN1",
				"3KXX5DmlsDAyO0pmgDT3pTyyuTfGr9ocJCOcaPwZDilAiwAp6U9egpHr1qOtx4dLLrtIVWE8npHK49BKKyyacg==",
			),
		},
	}

	for _, test := range testData {
		test := test
		suite.Run(test.name, func() {
			sessionLength := uint64(240)
			err := suite.keeper.SetDefaultSessionLength(suite.ctx, sessionLength)
			suite.Require().NoError(err)

			server := keeper.NewMsgServerImpl(suite.keeper)
			res, err := server.CreateSession(context.Background(), test.msg)

			// Valid response
			if res != nil {
				// Check the stored session
				expectedID := suite.keeper.GetLastSessionID(suite.ctx)
				session := types.Session{
					SessionId:      expectedID,
					CreationTime:   uint64(suite.ctx.BlockHeight()),
					ExpirationTime: uint64(suite.ctx.BlockHeight()) + sessionLength,
					Owner:          test.msg.Owner,
					Namespace:      test.msg.Namespace,
					ExternalOwner:  test.msg.ExternalOwner,
					PublicKey:      test.msg.PubKey,
					Signature:      test.msg.Signature,
				}

				var stored types.Session
				store := suite.ctx.KVStore(suite.keeper.storeKey)
				suite.keeper.cdc.MustUnmarshalBinaryBare(store.Get(types.SessionStoreKey(expectedID)), &stored)
				suite.Require().Equal(session, stored)

				// Check the events
				creationEvent := sdk.NewEvent(
					types.EventTypeCreateSession,
					sdk.NewAttribute(types.AttributeKeySessionID, session.SessionId.String()),
					sdk.NewAttribute(types.AttributeKeyNamespace, session.Namespace),
					sdk.NewAttribute(types.AttributeKeyExternalOwner, session.ExternalOwner),
					sdk.NewAttribute(types.AttributeKeyExpiry, fmt.Sprintf("%d", session.ExpirationTime)),
				)

				suite.Require().NotNil(res)
				suite.Require().Contains(suite.ctx.EventManager().Events(), creationEvent)
			}

			// Invalid response
			if res == nil {
				suite.Require().NotNil(err)
				suite.Require().Equal(err.Error(), test.error.Error())
			}
		})
	}
}
