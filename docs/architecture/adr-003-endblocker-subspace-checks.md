# ADR 003: Subspaces: End-blocker checks on unregistered users

## Changelog

- September 15th, 2021: Initial draft;

## Status

DRAFT

## Abstract

The `subspaces` module has the purpose to let dApp developers to create a "space" where their dApps can live
on with their own Term of Services. By doing this, all their users and posts will be associated with 
a particular subspace and should follow a set of rules and be subject to the changes that happens on the subspace itself.     
CONTINUA QUI

## Context
 
Currently, to obtain the right to post, create relationships and block users inside a subspace users firstly need to register
into it. After doing it, he will become a user of that subspace and be able to perform the set of ops we mentioned above here.  
To ensure a user can perform these operations, Desmos modules needs to check user's status on the subspace.   
Most of these checks have already been implemented inside the `posts` module and are crucial to make the whole `subspaces`
aspect to work correctly.

## Decision

> This section describes our response to these forces. It is stated in full sentences, with active voice. "We will ..."
> {decision body}

## Consequences

> This section describes the resulting context, after applying the decision. All consequences should be listed here, not just the "positive" ones. A particular decision may have positive, negative, and neutral consequences, but all of them affect the team and project in the future.

### Backwards Compatibility

> All ADRs that introduce backwards incompatibilities must include a section describing these incompatibilities and their severity. The ADR must explain how the author proposes to deal with these incompatibilities. ADR submissions without a sufficient backwards compatibility treatise may be rejected outright.

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

- {reference link}