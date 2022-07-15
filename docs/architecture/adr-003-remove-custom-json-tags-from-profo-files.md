# ADR 003: Remove custom JSON tags from Proto files

## Changelog

- October 13th, 2021: Proposed;
- October 14th, 2021: First review;
- October 18th, 2021: Second review;
- October 18th, 2021: Third review.

## Status

ACCEPTED Implemented

## Abstract

We SHOULD remove the custom `jsontag` from every proto file and let Protobuf generate the `json` tags
with its own conventions. 

## Context

Currently, when we encode a custom message instance as a JSON object using Protobuf, the result CAN
cause an error when we try to broadcast it to the chain. This is due to the fact that
custom `jsontag` option makes the proto compiler generate a wrong JSON notation. 
In addition to this, these fields can't be omitted because they miss the `omitempty` notation.

## Decision

In order to produce the correct proto file we only need to remove the custom `jsontag` 
from proto files that use it. So, taking `MsgSaveProfile` as an example, this is the file after
removing the `jsontag`s from it:

```protobuf
message MsgSaveProfile {
  
  ...
  
  string profilePicture = 4 [
    (gogoproto.moretags) = "yaml:\"profile_picture\""
  ];

  string coverPicture = 5 [
    (gogoproto.moretags) = "yaml:\"cover_picture\""
  ];
  
  ...
}
```
## Consequences

### Backwards Compatibility

These changes will not produce any backwards compatibility problem.

### Positive

- Make clients able to build the correct JSON field to forward correct messages.

### Negative

(none known)

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#644](https://github.com/desmos-labs/desmos/issues/644)