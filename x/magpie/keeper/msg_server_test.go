package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/magpie/keeper"
	"github.com/desmos-labs/desmos/x/magpie/types"
)

func (suite *KeeperTestSuite) Test_handleMsgCreateSession() {
	testData := []struct {
		name              string
		sessionLength     uint64
		msg               *types.MsgCreateSession
		expError          error
		expEvents         sdk.Events
		expStoredSessions []types.Session
	}{
		{
			name:          "Empty signature returns error",
			sessionLength: 1,
			msg: types.NewMsgCreateSession(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"",
				"",
				"",
			),
			expEvents: sdk.EmptyEvents(),
			expError:  sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "the session signature is not valid"),
		},
		{
			name:          "Invalid signature returns error",
			sessionLength: 1,
			msg: types.NewMsgCreateSession(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"ArDhBMh0X/3Akfc58oF1zFE00L/rLpgMMVvmcj0QlaN1",
				"3KXX5DmlsDAyO0pmgDT3pTyyuTfGr9ocJCOcaPwZDilAiwAp6U9egpHr1qOtx4dLLrtIVWE8npHK49BKKyyacg==",
			),
			expEvents: sdk.EmptyEvents(),
			expError:  sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "the session signature is not valid"),
		},
		//{
		//	name:          "Valid signature works properly",
		//	sessionLength: 240,
		//	msg: types.NewMsgCreateSession(
		//		"cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0",
		//		"cosmoshub-2",
		//		"cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0",
		//		"ArDhBMh0X/3Akfc58oF1zFE00L/rLpgMMVvmcj0QlaN1",
		//		"3KXX5DmlsDAyO0pmgDT3pTyyuTfGr9ocJCOcaPwZDilAiwAp6U9egpHr1qOtx4dLLrtIVWE8npHK49BKKyyacg==",
		//	),
		//	expEvents: sdk.Events{
		//		sdk.NewEvent(
		//			types.EventTypeCreateSession,
		//			sdk.NewAttribute(types.AttributeKeySessionID, "1"),
		//			sdk.NewAttribute(types.AttributeKeyNamespace, "cosmoshub-2"),
		//			sdk.NewAttribute(types.AttributeKeyExternalOwner, "cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0"),
		//			sdk.NewAttribute(types.AttributeKeyExpiry, "240"),
		//		),
		//	},
		//	expStoredSessions: []types.Session{
		//		{
		//			SessionId:      types.NewSessionID(1),
		//			CreationTime:   0,
		//			ExpirationTime: 240,
		//			Owner:          "cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0",
		//			Namespace:      "cosmoshub-2",
		//			ExternalOwner:  "cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0",
		//			PublicKey:      "ArDhBMh0X/3Akfc58oF1zFE00L/rLpgMMVvmcj0QlaN1",
		//			Signature:      "3KXX5DmlsDAyO0pmgDT3pTyyuTfGr9ocJCOcaPwZDilAiwAp6U9egpHr1qOtx4dLLrtIVWE8npHK49BKKyyacg==",
		//		},
		//	},
		//},
	}

	for _, test := range testData {
		test := test
		suite.Run(test.name, func() {
			err := suite.keeper.SetDefaultSessionLength(suite.ctx, test.sessionLength)
			suite.Require().NoError(err)

			server := keeper.NewMsgServerImpl(suite.keeper)
			_, err = server.CreateSession(sdk.WrapSDKContext(suite.ctx), test.msg)

			suite.RequireErrorsEqual(test.expError, err)
			suite.Require().Equal(suite.ctx.EventManager().Events(), test.expEvents)

			stored := suite.keeper.GetSessions(suite.ctx)
			suite.Require().Equal(test.expStoredSessions, stored)
		})
	}
}
