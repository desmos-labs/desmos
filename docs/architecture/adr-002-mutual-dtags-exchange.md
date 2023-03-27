# ADR 002: Mutual DTags exchange

## Changelog

- October 8th, 2021: First draft;
- October 11th, 2021: Moved from draft to proposed;
- October 18th, 2021: First review;
- October 18th, 2021: Second review;
- October 18th, 2021: Third review;
- October 18th, 2021: Fourth review;
- October 19th, 2021: Fifth review.

## Status

ACCEPTED Implemented

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
...
}
```

The major change will however be inside the `AcceptDTagTransferRequest` method of the `msgServer` implementation for 
the `x/profiles` module. Here we need to make sure that if the accepting user specifies a DTag that is equal to the one
of the receiving user, the method correctly handles the request by swapping users DTags:

```go
func (k msgServer) AcceptDTagTransferRequest(goCtx context.Context, msg *types.MsgAcceptDTagTransferRequest) (*types.MsgAcceptDTagTransferRequestResponse, error) {
    ...
    // Check for an existent profile of the receiving user
    receiverProfile, exist, err := k.GetProfile(ctx, msg.Sender)
    if err != nil {
        return nil, err
    }

    if exist && msg.NewDTag == receiverProfile.DTag {
        err = k.storeProfileWithoutDTagCheck(ctx, currentOwnerProfile)
        if err != nil {
            return nil, err
        }
    } else {
        err = k.StoreProfile(ctx, currentOwnerProfile)
        if err != nil {
            return nil, err
        }
    }
    ...
}
```

Here follows the specification of the new `StoreProfileWithoutDTagCheck` method introduced above at line 69:
```go
// StoreProfileWithoutDTagCheck stores the given profile inside the current context
// without checking if another profile with the same DTag exists.
// It assumes that the given profile has already been validated.
func (k Keeper) storeProfileWithoutDTagCheck(ctx sdk.Context, profile *types.Profile) error {
	store := ctx.KVStore(k.storeKey)

	oldProfile, found, err := k.GetProfile(ctx, profile.GetAddress().String())
	if err != nil {
		return err
	}
	if found && oldProfile.DTag != profile.DTag {
		// Remove the previous DTag association (if the DTag has changed)
		store.Delete(types.DTagStoreKey(oldProfile.DTag))

		// Remove all incoming DTag transfer requests if the DTag has changed since these will be invalid now
		k.DeleteAllUserIncomingDTagTransferRequests(ctx, profile.GetAddress().String())
	}

	// Store the DTag -> Address association
	store.Set(types.DTagStoreKey(profile.DTag), profile.GetAddress())

	// Store the account inside the auth keeper
	k.ak.SetAccount(ctx, profile)

	k.Logger(ctx).Info("stored profile", "DTag", profile.DTag, "from", profile.GetAddress())
	return nil
}
```

By introducing this method, we SHOULD also edit the `StoreProfile` method to use the new function in order
to pursue the DRY principle:
```go
// StoreProfile stores the given profile inside the current context.
// It assumes that the given profile has already been validated.
// It returns an error if a profile with the same DTag from a different creator already exists
func (k Keeper) StoreProfile(ctx sdk.Context, profile *types.Profile) error {
	addr := k.GetAddressFromDTag(ctx, profile.DTag)
	if addr != "" && addr != profile.GetAddress().String() {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest,
			"a profile with DTag %s has already been created", profile.DTag)
	}

	return k.storeProfileWithoutDTagCheck(ctx, profile)
}
```

## Consequences

### Backwards Compatibility

There are no backwards compatibility issues related to these changes.

### Positive

- Give the possibility to swap DTags between users

### Negative

(none known)

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]
The following tests cases needs to be added:
1. Accepting a DTag transfer request specifying the receiver DTag works properly;
2. Swapping two profiles DTags works properly.

## References

- Issue [#643](https://github.com/desmos-labs/desmos/issues/643)