# ADR 008: Fees modules

## Changelog

- March 2, 2022: Initial draft;
- March 8, 2022: First review;
- March 9, 2022: Second review;
- March 10, 2022: Third review.

## Status
ACCEPTED Implemented

## Abstract
This ADR defines the `x/fees` module which allows setting custom min fees for each message type.

## Context
In order to better prevent any kind of spam (e.g. zero-gas attacks, garbage smart contracts implementations, etc.) it's useful 
to have a system that allows to set a minimum amount of fees that needs to be paid when sending specific messages that can be vectors of such spam attacks. 
This system should allow changing dynamically such min fees amounts, so that the community can properly tweak them if necessary.

## Decision
We will create a module named `fees` that allows setting such min fee amounts of any kind of messages that our chain supports. 
This will then make sure that when such messages are broadcast inside a transaction, the transaction signer 
is paying at least the minimum amount of fees for each message.

### Types

#### Min Fees
```go
type MinFee struct {
    // The message type for which this min fee amount is valid
    messageType string

    // The min amount of fees to be paid for each instance of this type of message
    amount sdk.Coins
}
```

#### Params
We will save each `MinFee` instance as the module on-chain params in order to be able to later change them with governance proposals as follows:

```go
type Params struct {
	MinFees []MinFee
}
```

#### Ante Handler and Fee Decorator
We need to set up a custom [AnteHandler](https://github.com/cosmos/cosmos-sdk/blob/da36c46f3a3a8dec7eaa85b69e7fa154e9d64dce/types/handler.go#L8) 
in order to be able to manage custom fees.  
This one will look exactly like the default one except for the fact that it will have an extra decorator to handle
custom fees:
```go
type MinFeeDecorator struct {
	feesKeeper feeskeeper.Keeper
}
```

The custom `AnteHandler` will iterate over all the transaction's messages and perform the following operations:
- Fetch all the min fees to be paid for such message from the `x/fees` params (if no min fees are set, then `minFee = 0`);
- Sum all the min fees together;

After having calculated the total required min fees, it will check that fees are greater or equals to the min fees, 
and, if not, return an error.

## Consequences

### Positive
- Spam prevention of transaction containing specific messages

### Negative
- Applying custom fees to messages requires to add an extra decorator to the `AnteHandler`, 
which will need to perform stateful checks that can eventually slow down the node a bit. 

## Test Cases

We will need to add the following test cases:
- a transaction not having enough fees is rejected;
- a transaction having enough fees is accepted;
- `AnteHandler` benchmark tests to make sure it does not impact the transaction handling process too much.

## References

- [Ante Handlers](https://docs.cosmos.network/v0.44/modules/auth/03_antehandlers.html#antehandlers);
- [First issue about min fees](https://github.com/desmos-labs/desmos/issues/230).