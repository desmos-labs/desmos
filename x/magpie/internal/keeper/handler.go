package keeper

import (
	"fmt"
	"time"

	"encoding/base64"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kwunyeung/desmos/x/magpie/internal/types"
	"github.com/rs/xid"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// NewHandler returns a handler for "magpie" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgCreatePost:
			return handleMsgCreatePost(ctx, keeper, msg)
		case types.MsgEditPost:
			return handleMsgEditPost(ctx, keeper, msg)
		case types.MsgLike:
			return handleMsgLike(ctx, keeper, msg)
		case types.MsgCreateSession:
			return handleMsgCreateSession(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Magpie Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle creating a new post
func handleMsgCreatePost(ctx sdk.Context, keeper Keeper, msg types.MsgCreatePost) sdk.Result {

	post := types.Post{
		ID:            xid.New().String(),
		ParentID:      msg.ParentID,
		Message:       msg.Message,
		Created:       msg.Created,
		Likes:         0,
		Owner:         msg.Owner,
		Namespace:     msg.Namespace,
		ExternalOwner: msg.ExternalOwner,
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	)

	if err := keeper.CreatePost(ctx, post); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePost,
			sdk.NewAttribute(types.AttributeKeyPostID, post.ID),
			sdk.NewAttribute(types.AttributeKeyNamespace, post.Namespace),
			sdk.NewAttribute(types.AttributeKeyExternalOwner, post.ExternalOwner),
		),
	)

	return sdk.Result{
		Data:   keeper.cdc.MustMarshalBinaryLengthPrefixed(post.ID),
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgEditPost(ctx sdk.Context, keeper Keeper, msg types.MsgEditPost) sdk.Result {
	existing, found := keeper.GetPost(ctx, msg.ID)
	if found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s not found", msg.ID)).Result()
	}

	// checks if the the msg sender is the same as the current owner
	if !msg.Owner.Equals(existing.Owner) {
		return sdk.ErrUnauthorized("Incorrect owner").Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	)

	if err := keeper.EditPostMessage(ctx, existing, msg.Message); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeEditPost,
			sdk.NewAttribute(types.AttributeKeyPostID, msg.ID),
		),
	)

	return sdk.Result{
		Data:   keeper.cdc.MustMarshalBinaryLengthPrefixed(msg.ID),
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgLike(ctx sdk.Context, keeper Keeper, msg types.MsgLike) sdk.Result {
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return sdk.ErrUnknownRequest("Post doesn't exist").Result()
	}

	like := types.Like{
		ID:            xid.New().String(),
		Created:       msg.Created,
		PostID:        msg.PostID,
		Owner:         msg.Liker,
		Namespace:     msg.Namespace,
		ExternalOwner: msg.ExternalOwner,
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Liker.String()),
		),
	)

	if err := keeper.SavePostLike(ctx, post, like); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLikePost,
			sdk.NewAttribute(types.AttributeKeyLikeID, like.ID),
			sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID),
			sdk.NewAttribute(types.AttributeKeyNamespace, msg.Namespace),
			sdk.NewAttribute(types.AttributeKeyExternalOwner, msg.ExternalOwner),
		),
	)

	return sdk.Result{
		Data:   keeper.cdc.MustMarshalBinaryLengthPrefixed(like.ID),
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgCreateSession(ctx sdk.Context, keeper Keeper, msg types.MsgCreateSession) sdk.Result {

	// query if a previous TX with the same namespace and external owner exists
	// if a query exists,
	// see if current time is between creation time and expiry time
	// if yes, then continue and emit event
	// else return error

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	)

	// check if the signature is signed by the external address
	// addr, err := utils.GetAccAddressFromExternal(msg.ExternalOwner, msg.Namespace)

	// if err != nil {
	// 	return err.Result()
	// }

	// acc := auth.NewBaseAccountWithAddress(addr)

	// pubkey := acc.GetPubKey()

	// pubkey := sdk.MustGetAccPubKeyBech32(msg.Pubkey)

	pkBytes, _ := base64.StdEncoding.DecodeString(msg.Pubkey)

	var pkBytes33 = [33]byte{}

	copy(pkBytes33[:], pkBytes)

	pubkey := secp256k1.PubKeySecp256k1(pkBytes33)

	message := fmt.Sprintf(`{"account_number":"0","chain_id":"%s","fee":{"amount":[],"gas":"200000"},"memo":"","msgs":[{"type":"desmos/MsgCreateSession","value":{"created":"%s","external_owner":"%s","namespace":"%s","owner":"%s","pubkey":"%s","signature":null}}],"sequence":"0"}`,
		ctx.ChainID(), msg.Created.Format(time.RFC3339Nano), msg.ExternalOwner, msg.Namespace, msg.Owner.String(), msg.Pubkey)

	// message := `{"account_number":"0","chain_id":"tesmos-1","fee":{"amount":[],"gas":"200000"},"memo":"","msgs":[{"type":"desmos/MsgCreateSession","value":{"created":"2019-07-19T10:08:05.161Z","external_owner":"cosmos10505nl7yftsme9jk2glhjhta7w0475uv6pzj70","namespace":"cosmos","owner":"desmos186vmnukgywe9hwr233x8jcyvavm7zpven4jxlr","signature":null}}],"sequence":"0"}`
	sig, _ := base64.StdEncoding.DecodeString(msg.Signature)

	if !pubkey.VerifyBytes([]byte(message), sig) {
		return sdk.ErrUnauthorized("The session signature is not correct. " + message).Result()
		// panic("The session signature is not correct.")
	}

	session := types.Session{
		ID:            xid.New().String(),
		Created:       msg.Created,
		Expiry:        msg.Created.Add(time.Minute * 14400),
		Owner:         msg.Owner,
		Namespace:     msg.Namespace,
		ExternalOwner: msg.ExternalOwner,
		Pubkey:        msg.Pubkey,
		Signature:     msg.Signature,
	}

	if err := keeper.SaveSession(ctx, session); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateSession,
			sdk.NewAttribute(types.AttributeKeySessionID, session.ID),
			sdk.NewAttribute(types.AttributeKeyNamespace, msg.Namespace),
			sdk.NewAttribute(types.AttributeKeyExternalOwner, msg.ExternalOwner),
			sdk.NewAttribute(types.AttributeKeyExpiry, session.Expiry.Format(time.RFC3339Nano)),
		),
	)

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}
}
