# ADR 008: Fee modules

## Changelog

- March 2, 2022: Initial draft.

## Status
DRAFTED

## Abstract
This ADR defines the `x/fees` module which enables the possibility to set custom fees for each desmos message
through governance proposals.

## Context
In order to better prevent any kind of spam, from zero-gas attack, to smart contracts implementation, to subspaces creation and, 
in the near future, through posts, it's useful to have a system with which the community can tweak the costs of desmos messages.

## Decision
We will create a module named `fees` which give desmos community the possibility to add minimum fees to any Desmos custom
message through governance.

### Types
Minimum (or custom) fees will contain the following elements:  

* a string identifying the message type delivered with the transaction;
* the amount of fees associated with the give message.

#### Min Fees
```go
type MinFee struct {
	// The message type identifying the Desmos message
	messageType string
	
	// The amount of fees to be paid for the message
	amount sdk.Coins
}
```

#### Params
We will save the `MinFee`s for messages in the module's `Params` in order to be able to change them with governance 
`ParameterChangeProposal`s. To pursue efficiency, the ideal implementation for them is the following:

```go
type Params struct {
	MinFees map[string]MinFee
}
```

This map will use the proto message type as key, allowing us to keep a constant complexity lookup for its elements.

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

This decorator will implement the following interface:
```go
// AnteDecorator wraps the next AnteHandler to perform custom pre- and post-processing.
type AnteDecorator interface {
	AnteHandle(ctx Context, tx Tx, simulate bool, next AnteHandler) (newCtx Context, err error)
}
```

### `Query` Service
The module will expose the following query:
```protobuf
// Query defines the gRPC querier service.
service Query {
  // Params queries the fees module params
  rpc Params(QueryParamsRequest) returns (QueryParamsResponse);
}
```

## Consequences

### Negative
Applying custom fees to messages requires to add an extra decorator to the `AnteHandler`, 
which will need to perform stateful checks that can eventually slow down the node a bit. 

## Test Cases

We will need to add the following test cases :
* The custom fees are applied correctly;
* Benchmark on `AnteHandler`.

## References

- [Ante Handlers](https://docs.cosmos.network/v0.44/modules/auth/03_antehandlers.html#antehandlers);
- [First issue about min fees](https://github.com/desmos-labs/desmos/issues/230).