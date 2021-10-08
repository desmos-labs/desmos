# ADR 010: Enable mutual DTag exchange

## Changelog

- October 8th, 2021: First draft

## Status

DRAFT

## Abstract

We SHOULD edit the behavior of `MsgAcceptDTagTransferRequest` making the `newDTag` field an optional flag that
if specified, allows users to choose a new DTag otherwise it will simply swap the two users DTags. 

## Context

Currently, two users can't swap their DTag when transferring them.
For example, if Alice and Bob want to exchange their own DTag with each other they need to follow these steps:
* Alice transfers the DTag `@alice` to Bob;
* Alice select a random temporary DTag (e.g. `@charles`);
* Alice edits her profile to select the `@bob` DTag.

This flow could even be interrupted if, in the meantime, 
a third user create a profile with the now free `@bob` DTag before
Alice does it.

## Decision

To make possible the DTag mutual exchange, we need to make some changes on the logic that
handles `MsgAcceptDTagTransferRequest`.  
Firstly, we need to edit the `desmos tx profiles accept-dtag-transfer-request` so that the now
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

	cmd.Flags().String(FlagNickname, types.DoNotModify, "new DTag to be used")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
```

## Consequences

### Backwards Compatibility

### Positive

{positive consequences}

### Negative

{negative consequences}

### Neutral

{neutral consequences}

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section should contain a summary of issues to be solved in future iterations (usually referencing comments from a pull-request discussion).
Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

## Test Cases [optional]

Test cases for an implementation are mandatory for ADRs that are affecting consensus changes. Other ADRs can choose to include links to test cases if applicable.

## References

- Issue [#643](https://github.com/desmos-labs/desmos/issues/643)