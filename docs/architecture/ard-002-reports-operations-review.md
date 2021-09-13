# ADR 002: Reports save and delete operations

## Changelog

- September 13th, 2021: Initial draft

## Status

DRAFT

## Abstract

Inside Desmos, most of the common types you deal with has two methods: `save` (to create & edit) and `delete`.
The `report` type can only be created, without the possibility to edit or delete it later.
These methods SHOULD be implemented to improve the handling of `reports` and standardize the way in which
structures can be manipulated in Desmos.

## Context

The idea behind Desmos `reports` is to give users a way to express dissent toward a post leaving 
developers the freedom to implement their own ToS around them.
Since Desmos is a protocol that will serve as a common field for many social-enabled applications, 
it would have been extremely complicated to build a one-for-all reporting system. 
Different social apps have different scopes, and want to deal with contents in different ways.   
In order to leave everyone free to manage the contents of their social, the creation of the best
ToS-agnostic system to reports contents is crucial.

## Decision

In order to build a complete report system, we SHOULD edit the actual one and expand it to match the common
operations we're already using in other types like `Profile`.
The two main operations that will comprise the new reporting system will be:
 * `Save`: to both create and edit a `report`;
 * `Delete`: to delete a previously made `report`.



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