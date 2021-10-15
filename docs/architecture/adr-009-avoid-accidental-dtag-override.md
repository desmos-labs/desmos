# ADR 009: Avoid accidental DTag override

## Changelog

- October 7th, 2021: Proposed
- October 8th, 2021: First review
- October 11th, 2021: Finalized ADR

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
	--%s "LeoDiCaprio" \
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

The following tests cases MUST to be present:
1. Creating a profile with an empty DTag returns an error;   
2. Creating a profile with DTag `[do-not-modify]` returns an error;   
3. Updating a profile with a different DTag changes its value and returns no error;   
4. Updating a profile with DTag `[do-not-modify]` does not update its value and returns no error.

## References

- Issue [#622](https://github.com/desmos-labs/desmos/issues/622)