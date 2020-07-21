package keeper

import (
	"encoding/base64"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

// NewHandler returns a handler for "magpie" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgCreateSession:
			return handleMsgCreateSession(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
				fmt.Sprintf("unrecognized magpie message type: %v", msg.Type()))
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
	pkBytes, err := base64.StdEncoding.DecodeString(msg.PubKey)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot decode base64 public key")
	}

	var pkBytes33 = [33]byte{}
	copy(pkBytes33[:], pkBytes)
	pubkey := secp256k1.PubKeySecp256k1(pkBytes33)

	// Create the signature bytes using the given message  with an empty signature
	clearMsg := msg
	clearMsg.Signature = ""
	signedBytes := auth.StdSignBytes(msg.Namespace, 0, 0, auth.NewStdFee(200000, nil), []sdk.Msg{clearMsg}, "")

	// Decode the signature
	sig, err := base64.StdEncoding.DecodeString(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot decode base64 signature")
	}

	// Verify the signature
	if !pubkey.VerifyBytes(signedBytes, sig) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "the session signature is not valid")
	}

	// Create the session
	session := types.Session{
		SessionID:     keeper.GetLastSessionID(ctx).Next(),
		Created:       ctx.BlockHeight(),
		Expiry:        ctx.BlockHeight() + keeper.GetDefaultSessionLength(ctx),
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
		Events: ctx.EventManager().Events(),
	}
	return &result, nil
}
