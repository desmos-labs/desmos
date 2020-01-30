package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/magpie/internal/keeper"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"github.com/stretchr/testify/assert"
)

func Test_handleMsgCreateSession(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0")

	testData := []struct {
		name  string
		msg   types.MsgCreateSession
		error error
	}{
		{
			name: "Empty signature returns error",
			msg: types.MsgCreateSession{
				Owner:         testSession.Owner,
				Namespace:     testSession.Namespace,
				ExternalOwner: testSession.ExternalOwner,
				PubKey:        testSession.PubKey,
				Signature:     "",
			},
			error: sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "the session signature is not valid"),
		},
		{
			name: "Invalid signature returns error",
			msg: types.MsgCreateSession{
				Owner:     testSession.Owner,
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
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			sessionLength := int64(240)
			k.SetDefaultSessionLength(ctx, sessionLength)

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				// Check the stored session
				expectedID := k.GetLastSessionID(ctx)
				session := types.Session{
					SessionID:     expectedID,
					Created:       ctx.BlockHeight(),
					Expiry:        ctx.BlockHeight() + sessionLength,
					Owner:         test.msg.Owner,
					Namespace:     test.msg.Namespace,
					ExternalOwner: test.msg.ExternalOwner,
					PubKey:        test.msg.PubKey,
					Signature:     test.msg.Signature,
				}

				var stored types.Session
				store := ctx.KVStore(k.StoreKey)
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.SessionStorePrefix+expectedID.String())), &stored)
				assert.Equal(t, session, stored)

				// Check the events
				creationEvent := sdk.NewEvent(
					types.EventTypeCreateSession,
					sdk.NewAttribute(types.AttributeKeySessionID, session.SessionID.String()),
					sdk.NewAttribute(types.AttributeKeyNamespace, session.Namespace),
					sdk.NewAttribute(types.AttributeKeyExternalOwner, session.ExternalOwner),
					sdk.NewAttribute(types.AttributeKeyExpiry, strconv.FormatInt(session.Expiry, 10)),
				)

				assert.NotNil(t, res)
				assert.Contains(t, res.Events, creationEvent)
			}

			// Invalid response
			if res == nil {
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), test.error.Error())
			}
		})
	}
}
