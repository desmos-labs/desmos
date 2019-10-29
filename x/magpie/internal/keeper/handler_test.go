package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/magpie/internal/keeper"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"github.com/stretchr/testify/assert"
)

func Test_handleMsgCreateSession_EmitsMessageEvent(t *testing.T) {
	ctx, k := SetupTestInput()

	msgShareDocument := types.MsgCreateSession{
		Owner:         session.Owner,
		Created:       session.Created,
		Namespace:     session.Namespace,
		ExternalOwner: session.ExternalOwner,
		PubKey:        session.PubKey,
		Signature:     session.Signature,
	}

	handler := keeper.NewHandler(k)
	_ = handler(ctx, msgShareDocument)

	expected := sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeyAction, types.ActionCreationSession),
		sdk.NewAttribute(sdk.AttributeKeySender, msgShareDocument.Owner.String()),
	)
	assert.Contains(t, ctx.EventManager().Events(), expected)
}
