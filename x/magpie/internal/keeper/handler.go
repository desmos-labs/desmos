package keeper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// NewHandler returns a handler for "magpie" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgCreateSession:
			return handleMsgCreateSession(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Magpie message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreateSession(ctx sdk.Context, keeper Keeper, msg types.MsgCreateSession) (*sdk.Result, error) {

	// query if a previous TX with the same namespace and external owner exists
	// if a query exists,
	// see if current time is between creation time and expiry time
	// if yes, then continue and emit event
	// else return error

	// Get the public key used to sign the message
	pkBytes, _ := base64.StdEncoding.DecodeString(msg.PubKey)
	var pkBytes33 = [33]byte{}
	copy(pkBytes33[:], pkBytes)
	pubkey := secp256k1.PubKeySecp256k1(pkBytes33)

	// Create the StdSignDoc by using the given message data, with an empty string
	signedMsg := msg
	signedMsg.Signature = ""

	stdSignDoc := auth.StdSignDoc{
		AccountNumber: 0,
		ChainID:       msg.Namespace,
		Fee:           json.RawMessage(auth.NewStdFee(200000, nil).Bytes()),
		Memo:          "",
		Msgs:          []json.RawMessage{json.RawMessage(signedMsg.GetSignBytes())},
		Sequence:      0,
	}

	// Create the signature bytes
	signedBytes := sdk.MustSortJSON(keeper.Cdc.MustMarshalJSON(stdSignDoc))
	sig, _ := base64.StdEncoding.DecodeString(msg.Signature)

	// Verify the signature
	if !pubkey.VerifyBytes(signedBytes, sig) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "the session signature is not valid")
	}

	// Create the session
	session := types.Session{
		SessionID:     keeper.GetLastSessionID(ctx).Next(),
		Created:       ctx.BlockHeight(),
		Expiry:        ctx.BlockHeight() + 240, // 24 hours, counting a 6 secs block interval
		Owner:         msg.Owner,
		Namespace:     msg.Namespace,
		ExternalOwner: msg.ExternalOwner,
		PubKey:        msg.PubKey,
		Signature:     msg.Signature,
	}

	// Check for any previously existing session
	if _, found := keeper.GetSession(ctx, session.SessionID); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("session with id %s already exists", session.SessionID))
	}

	// Save the session
	keeper.SaveSession(ctx, session)

	createSessionEvent := sdk.NewEvent(
		types.EventTypeCreateSession,
		sdk.NewAttribute(types.AttributeKeySessionID, session.SessionID.String()),
		sdk.NewAttribute(types.AttributeKeyNamespace, session.Namespace),
		sdk.NewAttribute(types.AttributeKeyExternalOwner, session.ExternalOwner),
		sdk.NewAttribute(types.AttributeKeyExpiry, strconv.FormatInt(session.Expiry, 10)),
	)
	ctx.EventManager().EmitEvent(createSessionEvent)

	result := sdk.Result{
		Data:   types.ModuleCdc.MustMarshalBinaryLengthPrefixed(session.SessionID),
		Events: sdk.Events{createSessionEvent},
	}
	return &result, nil
}
