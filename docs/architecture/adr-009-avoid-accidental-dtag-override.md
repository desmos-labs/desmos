# ADR 009: Avoid accidental DTag override

## Changelog

- October 7th, 2021: Proposed
- October 8th, 2021: First review

## Status

PROPOSED

## Abstract

We SHOULD edit the behavior of the current `MsgSaveProfile` making the `DTag` an optional flag
in order to prevent users from accidentally overriding their own DTags. In this way, users will edit
their DTag only when they specify them with the flag.


## Context

Currently, the `desmos profile tx save` CLI command always requires to specify a DTag when using it. This means that
users that do not want to edit their DTag need to specify it anyway. This could lead to the situation where a user 
accidentally makes a typo while inserting the DTag triggering its edit. If this happens, and the user doesn't notice it 
immediately, another user can register a profile with the now free DTag, stealing it from the original user.

## Decision

To avoid the situation described above, we need to perform some changes on the logic that handles a `MsgSaveProfile`.
First, we need to edit the `desmos tx profiles save` CLI command so that the now required `dtag` field becomes 
an optional flag:
```go
func GetCmdSaveProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Args:  cobra.NoArgs,
		Short: "Save your profile",
		Long: fmt.Sprintf(`
Save a new profile or edit the existing one specifying a DTag, a nickname, biography, profile picture and cover picture.
Every data given through the flags is optional.
If you are editing an existing profile you should fill only the fields that you want to edit.
The empty ones will be filled with a special [do-not-modify] flag that tells the system to not edit them.

%s tx profiles save 
    --%s "LeoDiCap" \
	--%s "Leonardo Di Caprio" \
	--%s "Hollywood actor. Proud environmentalist" \
	--%s "https://profilePic.jpg" \
	--%s "https://profileCover.jpg"
`, version.AppName, FlagDTag, FlagNickname, FlagBio, FlagProfilePic, FlagCoverPic),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			dTag, _ := cmd.Flags().GetString(FlagDTag)
			nickname, _ := cmd.Flags().GetString(FlagNickname)
			bio, _ := cmd.Flags().GetString(FlagBio)
			profilePic, _ := cmd.Flags().GetString(FlagProfilePic)
			coverPic, _ := cmd.Flags().GetString(FlagCoverPic)

			msg := types.NewMsgSaveProfile(dTag, nickname, bio, profilePic, coverPic, clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagDTag, types.DoNotModify, "DTag to be used")
	cmd.Flags().String(FlagNickname, types.DoNotModify, "Nickname to be used")
	cmd.Flags().String(FlagBio, types.DoNotModify, "Biography to be used")
	cmd.Flags().String(FlagProfilePic, types.DoNotModify, "Profile picture")
	cmd.Flags().String(FlagCoverPic, types.DoNotModify, "Cover picture")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
```
Second, we need to remove the check on DTag inside `ValidateBasic()` that currently makes it impossible to specify an 
empty DTag inside a `MsgSaveProfile`:
```go
func (msg MsgSaveProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator: %s", msg.Creator))
	}
	return nil
}
```

Third, we need to edit the behaviour of the `msgServer#SaveProfile` method:
```go
func (k msgServer) SaveProfile(goCtx context.Context, msg *types.MsgSaveProfile) (*types.MsgSaveProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	profile, found, err := k.GetProfile(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	// Create a new profile if not found
	if !found {
		if msg.DTag == types.DoNotModify {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "DTag need to be specified if user doesn't have a profile")
		}

		addr, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			return nil, err
		}

		profile, err = types.NewProfileFromAccount(msg.DTag, k.ak.GetAccount(ctx, addr), ctx.BlockTime())
		if err != nil {
			return nil, err
		}
	}
	
	// Update the existing profile with the values provided from the user
	updated, err = profile.Update(types.NewProfileUpdate(
		msg.DTag,
		msg.Nickname,
		msg.Bio,
		types.NewPictures(msg.ProfilePicture, msg.CoverPicture), 
	))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Validate the profile
	err = k.ValidateProfile(ctx, updated)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the profile
	err = k.StoreProfile(ctx, updated)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeProfileSaved,
		sdk.NewAttribute(types.AttributeProfileDTag, updated.DTag),
		sdk.NewAttribute(types.AttributeProfileCreator, updated.GetAddress().String()),
		sdk.NewAttribute(types.AttributeProfileCreationTime, updated.CreationDate.Format(time.RFC3339Nano)),
	))

	return &types.MsgSaveProfileResponse{}, nil
}
```
Finally, we need to edit the behavior of `Update` method. There are two possible ways to handle this:  

1) The method handles the empty DTag situation as it does when it finds the `DoNotModify` identifier. This
is a simple way to go, but it can cause some confusion in the user that tries to set an empty DTag and got no error back.
```go
func (p *Profile) Update(update *ProfileUpdate) (*Profile, error) {
	if update.DTag == DoNotModify {
		update.DTag = p.DTag
	}

	if update.Nickname == DoNotModify {
		update.Nickname = p.Nickname
	}

	if update.Bio == DoNotModify {
		update.Bio = p.Bio
	}

	if update.Pictures.Profile == DoNotModify {
		update.Pictures.Profile = p.Pictures.Profile
	}

	if update.Pictures.Cover == DoNotModify {
		update.Pictures.Cover = p.Pictures.Cover
	}

	newProfile, err := NewProfile(
		update.DTag,
		update.Nickname,
		update.Bio,
		update.Pictures,
		p.CreationDate,
		p.GetAccount(),
	)
	if err != nil {
		return nil, err
	}

	return newProfile, nil
}
```

## Consequences

### Backwards Compatibility

There are no backwards compatibility issues related to these changes.

### Positive

* Protect users from accidentally editing their DTag.

### Negative

- None known

### Neutral

- None known 

## Further Discussions

## Test Cases [optional]

## References

- Issue [#622](https://github.com/desmos-labs/desmos/issues/622)