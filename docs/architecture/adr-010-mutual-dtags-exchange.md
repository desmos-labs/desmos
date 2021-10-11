# ADR 010: Enable mutual DTag exchange

## Changelog

- October 8th, 2021: First draft;
- October 11th, 2021: Moved from draft to proposed.

## Status

PROPOSED

## Abstract

We SHOULD edit the behavior of `MsgAcceptDTagTransferRequest` making the `newDTag` field an optional flag. 
When specified, it will allow the DTag sender to choose a new DTag before trading his one. 

## Context

Currently, the DTag transfer process doesn't allow two users to swap their DTag when transferring them.
For example, if Alice and Bob want to exchange their own DTag with each other, they need to follow these steps:
* Alice transfers the DTag `@alice` to Bob;
* Alice select a random temporary DTag (e.g. `@charles`);
* Alice edits her profile to select the `@bob` DTag.

This flow works. However, it may be interrupted if in the meantime 
a third user creates a profile with the now free `@bob` DTag before Alice does. 
If this happens, Alice will be forced to choose a new DTag or send a new transfer request 
to the third user.

## Decision

To make possible the mutual DTag transfer, we need to make some changes on the logic that
handles `MsgAcceptDTagTransferRequest`.  
First, we need to edit the `desmos tx profiles accept-dtag-transfer-request` so that the now
required `newDTag` field becomes an optional flag:
```go
func GetCmdAcceptDTagTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept-dtag-transfer-request [address]",
		Short: "Accept a DTag transfer request made by the user with the given address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			receivingUser, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			dtag, _ := cmd.Flags().GetString(FlagDTag)
			
			msg := types.NewMsgAcceptDTagTransferRequest(dtag, receivingUser.String(), clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagDTag, "", "new DTag to be used")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
```

Second, we need to edit the logic that handles the `MsgAcceptDTagTransferRequest` in order to let it perform a 
DTag swap when no new DTag has been specified:
```go
func (k msgServer) AcceptDTagTransferRequest(goCtx context.Context, msg *types.MsgAcceptDTagTransferRequest) (*types.MsgAcceptDTagTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	request, found, err := k.GetDTagTransferRequest(ctx, msg.Sender, msg.Receiver)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no request made from %s", msg.Sender)
	}

	// Get the current owner profile
	currentOwnerProfile, exist, err := k.GetProfile(ctx, msg.Receiver)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "profile of %s doesn't exist", msg.Receiver)
	}

	// Get the DTag to trade and make sure its correct
	dTagWanted := request.DTagToTrade
	dTagToTrade := currentOwnerProfile.DTag
	if dTagWanted != dTagToTrade {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the owner's DTag is different from the one to be exchanged")
	}

	// Check for an existent profile of the receiving user
	receiverProfile, exist, err := k.GetProfile(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}

	// Check if the DTag is not specified then perform DTag swap if the receiver has a profile
	if  strings.TrimSpace(msg.NewDTag) == "" {
		if !exist {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "receiver profile doesn't exist and thus can't swap their DTag")
		}
		if err = k.SwapDTag(ctx, currentOwnerProfile, receiverProfile); err != nil {
			return nil, err
		}
	} else {
		// Change the DTag and validate the profile
		currentOwnerProfile.DTag = msg.NewDTag
		err = k.ValidateProfile(ctx, currentOwnerProfile)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}

		// Store the profile
		err = k.StoreProfile(ctx, currentOwnerProfile)
		if err != nil {
			return nil, err
		}

		if !exist {
			add, err := sdk.AccAddressFromBech32(msg.Sender)
			if err != nil {
				return nil, err
			}

			senderAcc := k.ak.GetAccount(ctx, add)
			if senderAcc == nil {
				senderAcc = authtypes.NewBaseAccountWithAddress(add)
			}

			receiverProfile, err = types.NewProfileFromAccount(dTagToTrade, senderAcc, ctx.BlockTime())
			if err != nil {
				return nil, err
			}
		} else {
			receiverProfile.DTag = dTagToTrade
		}

		// Validate the receiver profile
		err = k.ValidateProfile(ctx, receiverProfile)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}

		// Save the receiver profile
		err = k.StoreProfile(ctx, receiverProfile)
		if err != nil {
			return nil, err
		}
	}

	k.DeleteAllUserIncomingDTagTransferRequests(ctx, msg.Receiver)
	k.DeleteAllUserIncomingDTagTransferRequests(ctx, msg.Sender)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDTagTransferAccept,
		sdk.NewAttribute(types.AttributeDTagToTrade, dTagToTrade),
		sdk.NewAttribute(types.AttributeNewDTag, msg.NewDTag),
		sdk.NewAttribute(types.AttributeRequestSender, msg.Sender),
		sdk.NewAttribute(types.AttributeRequestReceiver, msg.Receiver),
	))

	return &types.MsgAcceptDTagTransferRequestResponse{}, nil
}
```

Here follows the specification of the new `SwapDTag` method introduce above at line 113.  
It comprehends also a private utility method called `replaceProfileDTag` that handles the
operations related to the DTag replacement.

```go
// replaceProfileDTag replace the given profile dTag with the given dTag.
// It returns error if the edited profile is invalid.
func (k Keeper) replaceProfileDTag(ctx sdk.Context, store sdk.KVStore, profile *types.Profile, dTag string) error {
profile.DTag = dTag
if err := k.ValidateProfile(ctx, profile); err != nil {
return err
}
store.Set(types.DTagStoreKey(profile.DTag), profile.GetAddress())
k.ak.SetAccount(ctx, profile)
return nil
}

// SwapDTag swap the profileA DTag with the profileB DTag
// It returns an error if one of the two profiles is invalid after the swap.
func (k Keeper) SwapDTag(ctx sdk.Context, profileA, profileB *types.Profile) error {
store := ctx.KVStore(k.storeKey)

dTagA := profileA.DTag
dTagB := profileB.DTag

// save profileA with profileB dTag
if err := k.replaceProfileDTag(ctx, store, profileA, dTagB); err != nil {
return err
}

if err := k.replaceProfileDTag(ctx, store, profileB, dTagA); err != nil {
return err
}

return nil
}
```

## Consequences

### Backwards Compatibility

There are no backwards compatibility issues related to these changes.

### Positive

- Improve the UX of DTag transfer requests
- Give the possibility to swap DTags between users

### Negative

- None knows

### Neutral

- None knows

## Further Discussions

## Test Cases [optional]
The following tests cases needs to be added:
1) Accept a transfer without specifying a new DTag;
2) Swap two profiles DTags;
3) Replace DTag with valid DTag;
4) Replace DTag with invalid DTag.

## References

- Issue [#643](https://github.com/desmos-labs/desmos/issues/643)