# ADR 001: Reports type field review

## Changelog

- September 9th, 2021: Initial draft

## Status

DRAFT

## Abstract

Report's `type` field SHOULD be expanded and enhanced to make possible adding 
more reasons and elaborate more why a post has been reported. 

## Context

Centralized social networks give users the possibility to select multiple reasons when they're about to reporting a post.
By doing this, the report become itself more explicit and complete and will later help the moderator/community to judge and choose the best
way to handle it. 
Actually inside Desmos we don't handle the reporting system like described. Instead, we rely on our `Report` object 
which contains the `type` field. The `type` field takes care to contain the reason behind the report, but at the moment,
it can only store 1 reason at time. Despite this can be handy and quick during as a first implementation attempt, it
may be not enough for future decentralized social networks and social enabled apps.  

Users need an exhaustive way to report contents on decentralized platforms, so having a repeated `type` field that allows
to do this should be the better option to improve the whole report system. This ADR proposes such changes inside the current
`Report` type.

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