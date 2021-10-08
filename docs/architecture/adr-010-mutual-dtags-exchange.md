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