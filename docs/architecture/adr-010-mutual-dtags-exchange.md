# ADR 010: Enable mutual DTag exchange

## Changelog

- October 8th, 2021: First draft;
- October 11th, 2021: Moved from draft to proposed;
- October 18th, 2021: First review;
- October 18th, 2021: Second review;
- October 18th, 2021: Third review;

## Status

PROPOSED

## Abstract

We SHOULD edit the inner logic of the DTag transfer acceptance in order to make
the mutual exchange of DTags possible between users. 

## Context

Currently, the DTag transfer process doesn't allow two users to swap their DTag when accepting a transfer request.
For example, if Alice and Bob want to exchange their own DTag with each other, they need to follow these steps:
1. Alice transfers the DTag `@alice` to Bob;
2. Alice selects a random temporary DTag (e.g. `@charles`);
3. Alice edits her profile to select the `@bob` DTag.

Although this flow works, in between steps 2 and 3, a third user MIGHT create a profile with the now free `@bob` DTag before Alice does. 
If this happens, Alice will be forced to choose a new DTag or send a new transfer request to the third user in order to 
obtain Bob's original DTag. For this reason, we should make it possible for the DTag transfer recipient to claim the 
sender's original DTag without performing additional steps later.

## Decision

In order to properly support DTag swaps, we will edit how `MsgAcceptDTagTransferRequest` are handled in order to allow 
the request receiver to specify the request sender's DTag as their new DTag.

To make this more clear to the users, we can add a description with `Short` and `Long` to specify this new behavior:
```go
func GetCmdAcceptDTagTransfer() *cobra.Command {
cmd := &cobra.Command{
Use:   "accept-dtag-transfer-request [DTag] [address]",
...
Short:  `Accept a DTag transfer request made by the user with the given address.
When accepting the request, you can specify the request recipient DTag as your new DTag. 
If this happens, your DTag and the other user's one will be effectively swapped.`
Long:  `Accept a DTag transfer request made by the user with the given address.
When accepting the request, you can specify the request recipient DTag as your new DTag. 
If this happens, your DTag and the other user's one will be effectively swapped.`
...
}
```

The major change will however be inside the `AcceptDTagTransferRequest` method of the `msgServer` implementation for 
the `x/profiles` module. Here we need to make sure that if the accepting user specifies a DTag that is equal to the one
of the receiving user, the method correctly handles the request by swapping users DTags:

```go
func (k msgServer) AcceptDTagTransferRequest(goCtx context.Context, msg *types.MsgAcceptDTagTransferRequest) (*types.MsgAcceptDTagTransferRequestResponse, error) {
	...
	// Check if the msg NewDTag is equal to the receiver one, if so, perform the DTags swap
	if  exist && msg.NewDTag == receiverProfile.DTag {
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
	...
}
```

Here follows the specification of the new `SwapDTag` method introduce above at line 111.  
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

- Give the possibility to swap DTags between users

### Negative

- None knows

### Neutral

- None knows

## Further Discussions

## Test Cases [optional]
The following tests cases needs to be added:
1) Accept a transfer specifying the receiver DTag;
2) Swap two profiles DTags;
3) Replace DTag with valid DTag;
4) Replace DTag with invalid DTag.

## References

- Issue [#643](https://github.com/desmos-labs/desmos/issues/643)