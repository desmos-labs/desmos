# ADR 004: Expiration of application links

## Changelog

- September 20th, 2021: Initial draft;
- September 21th, 2021: Moved from DRAFT to PROPOSED;
- December  22th, 2021: First review;
- January   04th, 2022: Second review;
- January   10th, 2022: Third review.

## Status

PROPOSED

## Abstract

Currently when a user links a centralized application with their Desmos profile, the created link contains a timestamp of when it has been created.   
Since centralized applications' username can be switched and sold, we SHOULD implement an "expiration date" system on links. 
This means that after a certain amount of time passes, the link will be automatically marked deleted and the user has to connect it again in order to keep it valid.

## Context

Desmos `profiles` module gives users possibility to link their desmos profile with both centralized application and 
other blockchains accounts. By doing this, a user can be verified as the owner of those accounts and prove to the system
that they're not impersonating anyone else. This verification however remains valid only if the user
never trades or sells their centralized-app username to someone else. If they do, the link to such username MUST be invalidated. 
Unfortunately for us, it's not possible to understand when this happens since it's off-chain. 
To prevent this situation, an "expiration time" SHOULD be added to the `ApplicationLink` object.

## Decision

To implement the link expiration we will act as follow:
1) Extend the `ApplicationLink` structure by adding an `expiration_time` field that represent the time when the link will expire and will be
   deleted from the store.
2) Save a reference of the expiring link inside the store under the prefix `ExpiringApplicationLinkPrefix`. It will be formed by a prefix and a 
   `timestamp` value which will make it easy to iterate over the expired links
3) Add a new keeper method that will allow to iterate over these expired links
4) Use the created iterator to implement another keeper method `DeleteExpiredApplicationLinks` that will take care of deleting all the expired
   links
5) Add this method to the `BeginBlock` function in order to trigger the deletion of expired link at each block

We will also need to introduce a new `x/profiles` on chain parameter, `AppLinkParams`, which will register the default expiration time
of all the app links and could be edited by the community with a governance proposal. 
The parameter will be later used inside `StartProfileConnection` to calculated the estimated expire time of each `AppLink`.

We also need to add a new `AppLinksParams` parameter to the `x/profiles` params, which will contain
an `expiration_time` timestamp that will be added to all the `AppLinks` before they're stored in the chain. 
The implementation for the new parameter is the following:

## Consequences

### Backwards Compatibility

This update will affect the `ApplicationLink` object by adding the new `ExpirationBlockHeight` 
field on it breaking the compatibility with the previous versions of the software. To allow
a smooth update and overcome these compatibility issues, we need to set up a proper migration
from the previous versions to the one that will include the additions contained in this ADR.

### Positive

* Considerably reduce the possibility of impersonation of entities and users of centralized apps;

### Negative

By adding the extra `expiration_time` field we are going to raise the overall `AppLinks` handling complexity.
Since we're performing an iteration over all the expired link references inside at the start of each block this can require
an amount of time that SHOULD be studied with benchmark tests during the implementation.

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#516](https://github.com/desmos-labs/desmos/issues/516)
- PR [#562](https://github.com/desmos-labs/desmos/pull/562)