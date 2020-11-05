package keeper

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	legacyauth "github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the magpie MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateSession(goCtx context.Context, msg *types.MsgCreateSession) (*types.MsgCreateSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

	pubkey := secp256k1.PubKey{Key: pkBytes}

	// Create the signature bytes using the given message  with an empty signature
	clearMsg := msg
	clearMsg.Signature = ""

	//nolint:staticcheck
	signedBytes := legacyauth.StdSignBytes(
		msg.Namespace,
		0,
		0,
		0,
		legacyauth.NewStdFee(200000, nil),
		[]sdk.Msg{clearMsg},
		"",
	)

	// Decode the signature
	sig, err := base64.StdEncoding.DecodeString(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot decode base64 signature")
	}

	// Verify the signature
	if !pubkey.VerifySignature(signedBytes, sig) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "the session signature is not valid")
	}

	// Create the session
	session := types.NewSession(
		k.GetLastSessionID(ctx).Next(),
		msg.Owner,
		uint64(ctx.BlockHeight()),
		uint64(ctx.BlockHeight())+k.GetDefaultSessionLength(ctx),
		msg.Namespace,
		msg.ExternalOwner,
		msg.PubKey,
		msg.Signature,
	)

	// Check for any previously existing session
	if _, found := k.GetSession(ctx, session.SessionId); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"session with id %d already exists", session.SessionId)
	}

	// Save the session
	k.SaveSession(ctx, session)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCreateSession,
		sdk.NewAttribute(types.AttributeKeySessionID, session.SessionId.String()),
		sdk.NewAttribute(types.AttributeKeyNamespace, session.Namespace),
		sdk.NewAttribute(types.AttributeKeyExternalOwner, session.ExternalOwner),
		sdk.NewAttribute(types.AttributeKeyExpiry, fmt.Sprintf("%d", session.ExpirationTime)),
	))

	return &types.MsgCreateSessionResponse{}, nil
}
