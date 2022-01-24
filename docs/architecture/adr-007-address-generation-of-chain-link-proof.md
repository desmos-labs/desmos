# ADR 006: Address generation of chain link proof

## Changelog

- January 20th, 2022: Initial draft;

## Status

PROPOSED

## Abstract

Currently, Desmos chain link supports `Bech32`, `Base58` and `Hex` address formats. Besides, each of them only couples one address generation algorithms. 
Since Desmos will support more addresses from other chains, which might apply the supported address format but with the unsupported algorithm, we SHOULD split the address algorithm from `AddressData` instances.

## Context

The `x/profiles` module gives users the possibility to link their profile to different external accounts.
However, each of supported address format are sticked on only one generation algorithm now, says:
1. `Bech32Address` utilizes basic cosmos generation functions which are `ripemd160(sha256(32 bytes public key))[:20]` for simple account and `sha256(aminoCdc.Marshal(multisig public key))[:20]` for multisig account
2. `Base58Address` generates address in the solana way which is from 32 bytes public key directly
3. `HexAddress` uses the ethereum generation function keccak64(64 bytes public key)[12:]

Due to the fact all the `AddressData` instances highly couples to the address generation function, it is impossible to link the supported address format with the different address generation algorithms.

## Decision

We propose to apply `Strategy` design pattern by a new interface to handle address generation from public key, which includes `GenerateAddress` function.
```go
type AddressGenerationStrategy interface{
    GenerateAddress(pk PubKey) []byte
}
```

Subsequently, the `AddressData` instance like `Bech32Address` can include the any value of the strategy instance which implements `AddressGenerationStrategy` as a property.
```proto
message Bech32Address{
    ...
    google.protobuf.Any strategy = 3 [(cosmos_proto.accepts_interface) = "AddressGenerationStrategy"];
}
```

Finally, the `VerifyPubKey` process can be implemented as follows:
1. Unpack the strategy from the any value
1. Generate the address raw bytes from public key by the strategy 
2. Compare the value of address data instance if it is equal to the address raw bytes from the public key

## Consequences

### Backwards Compatibility

Since this update will affect all the instances of `AddressData` by adding a `Strategy` field, this feature breaks the compatibility with the previous versions of the software.
For this reason, the migration to new `AddressData` instances from old ones will be needed to solve the compatibility issue.

### Positive

- Easy to extend the address generation algorithm for `AddressData` instances

### Negative


### Neutral


## Further Discussions


## Test Cases [optional]

- Verify public key including unpackable `AddressGenerationStrategy` returns error
- Verify public key including wrong `AddressGenerationStrategy`returns error
- Verify public key including proper `AddressGenerationStrategy` returns no error

## References

- [Strategy pattern](https://refactoring.guru/design-patterns/strategy)
- [Cosmos address](https://docs.cosmos.network/master/architecture/adr-028-public-key-addresses.html#legacy-public-key-addresses-don-t-change)
- [Ethereum address](https://ethereum.org/en/developers/docs/accounts/#account-creation)
- [Solana address](https://docs.solana.com/terminology#account)