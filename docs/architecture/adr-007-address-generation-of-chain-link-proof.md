# ADR 006: Address generation of chain link proof

## Changelog

- January 20th, 2022: Initial draft;

## Status

PROPOSED

## Abstract

Currently, Desmos chain link only supports `Cosmos`, `Solana` and `Ethereum` address format. Since Desmos will support more links from any other chains in the future, we SHOULD split the address algorithm from `AddressData` instance and add more address generation methods.

## Context

The `x/profiles` module gives users the possibility to link their profile to different external accounts.
However, each of supported address format are sticked on only one generation algorithm now, says:
1. `Bech32Address` utilizes basic cosmos generation functions which are `ripemd160(sha256(32 bytes public key))[:20]` for simple account and `sha256(aminoCdc.Marshal(multisig public key))[:20]` for multisig account
2. `Base58Address` generates address in the solana way which is from 32 bytes public key directly
3. `HexAddress` uses the ethereum generation function keccak64(64 bytes public key)[12:]

Due to the fact the address data highly couples to the generation function, it is impossible to link the supported chain address format which uses the different address generation function.  

## Decision

We proposed to create a new interface to handle address generation from public key, which includes `GenerateAddress` function.
```go
type AddressGenerationStrategy interface{
    GenerateAddress(pk PubKey) []byte
}
```

Subsequently, the `AddressData` instance like `Bech32Address` can apply any value of the strategy instance which implements `AddressGenerationStrategy` as a property.
```proto
message Bech32Address{
    ...
    google.protobuf.Any strategy = 3 [(cosmos_proto.accepts_interface) = "AddressGenerationStrategy"];
}
```

Finally, the `VerifyPubKey` process can be implemented as follows:
1. Unpack strategy any value into strategy interface
1. Use the strategy to generate the address raw bytes from public key
2. Compare the value of address data instance if it is equal to the address raw bytes from the public key

## Consequences

### Backwards Compatibility

Since we changed all the instances of `AddressData`, like `Bech32Address`, to have a strategy object in the structure, this feature is not backwards compatible.
For this reason, it will be needed to migrate to new `AddressData` instances.

### Positive

- Easy to extend the address generation algorithm for `AddressData` instances

### Negative


### Neutral


## Further Discussions


## Test Cases [optional]


## References
