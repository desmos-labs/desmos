# ADR 011: Wrong json profile pics fields fix

## Changelog

- October 13th, 2021: Proposed
- October 14th, 2021: First review
## Status

PROPOSED

## Abstract

We MUST edit the `profile_picture` and `cover_picture` fields inside `msg_profile.proto` 
to make it possible to omit them.

## Context

Currently, when we encode a `MsgSaveProfile` instance as a JSON object using Protobuf, the result will
cause an error when we try to broadcast it to the chain. This is due to the fact that
both `profile_picture` and `cover_picture` fields have a wrong notation that make the Proto compiler
produce a file with wrong JSON options. In addition to this, these fields can't be omitted because they miss
the `omitempty` notation.

## Decision

In order to produce the correct proto file we only need to edit the notation of 
these fields inside the `msg_profile.proto` file by:
1. Renaming them from snake_case to camelCase;
2. Appending the `omitempty` string inside the _jsontag_ field.

```protobuf
message MsgSaveProfile {
  
  ...
  
  string profilePicture = 4 [
    (gogoproto.jsontag) = "profile_picture,omitempty",
    (gogoproto.moretags) = "yaml:\"profile_picture\""
  ];

  string coverPicture = 5 [
    (gogoproto.jsontag) = "cover_picture,omitempty",
    (gogoproto.moretags) = "yaml:\"cover_picture\""
  ];
  
  ...
}
```
## Consequences

### Backwards Compatibility

These changes will not produce any backwards compatibility problem.
### Positive

- Make clients able to build the correct JSON field to forward a `MsgSaveProfile`.

### Negative

- None knows

### Neutral

- None knows

## Further Discussions

## Test Cases [optional]

## References

- Issue [#644](https://github.com/desmos-labs/desmos/issues/644)