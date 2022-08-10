# ADR 007: Address generation of chain link proof

## Changelog

- January 20th, 2022: Initial draft;
- July 28th, 2022: First review;

## Status

PROPOSED

## Abstract

Currently, Desmos allows linking other chains accounts which addresses are formatted using either the `Bech32`, `Base58` or `Hex` encoding and generated using a single algorithm specific to the encoding itself. Since Desmos idea is to support as many chains as possible, we SHOULD split the address generation algorithm from the encoding algorithm so that more chains can be linked properly.

## Context

The `x/profiles` module gives users the possibility to link their profile to different external accounts.  However, each one of the supported address formats currently supports only one generation algorithm.

`Bech32Address` relies on the address generation algorithm specific of Cosmos-SDK chains:
- for single signature accounts: 
  ```
  ripemd160(sha256(32 bytes public key))[:20]
  ``` 
- for multi-sig accounts: 
  ```
  sha256(aminoCdc.Marshal(multisig public key))[:20]
  ```

`Base58Address` relies on the Solana algorithm:
```
(32 bytes public key)[:32]
```

`HexAddress` relies on the Ethereum algorithm:
```
keccak64(64 bytes public key)[12:]
```

Due to the fact all the `AddressData` implementations are highly coupled to a single address generation function, it is currently impossible to link an address that needs to be encoded using a supported encoding method but is generated using a different algorithm. This is the case of Elrond addresses which are encoded using the `Bech32` encoding algorithm, but are generated using an algorithm that is different from the Cosmos one.

## Decision

In order to allow developers to integrate any kind of address, we will review how the current `Address` structure is made. Instead of using an interface with multiple implementations based on the encoding algorithm, we will use a single structure that allows to specify both the encoding algorithm and the hashing algorithm(s) that should be used to get the proper address: 

```protobuf
syntax = "proto3";

// Address contains the data of an external address
message Address {
  // Encoded value of the address
  string value = 1;
  
  // Algorithm that has been used in order to generate the address starting from the public key bytes
  GenerationAlgorithm generation_algorithm = 2;
  
  // Algorithm that needs to be used to properly encode the address 
  google.proto.Any encoding_algorithm = 3 [ (cosmos_proto.accepts_interface) = "AddressEncoding" ];
}

// GenerationAlgorithm represents various address generation algorithms
enum GenerationAlgorithm {
  // GENERATION_ALGORITHM_UNKNOWN represents an unknown algorithm and will be discarded 
  GENERATION_ALGORITHM_UNKNOWN = 0;
  
  // GENERATION_ALGORITHM_COSMOS represents the Cosmos generation algorithm
  GENERATION_ALGORITHM_COSMOS = 1;
  
  // GENERATION_ALGORITHM_EVM represents the EVM generation algorithm  
  GENERATION_ALGORITHM_EVM = 2;
  
  // GENERATION_ALGORITHM_DO_NOTHING should be used when the public key bytes do not need to be modified  
  GENERATION_ALGORITHM_DO_NOTHING = 3;
}

// Bech32Encoding represents the encoding algorithm based on the Bech32 format
message Bech32Encoding {
  option (cosmos_proto.implements_interface) = "AddressEncoding";
  
  // Prefix to be used
  string prefix = 1; 
}

// Base58Encoding represents the encoding algorithm based on the Base58 format
message Base58Encoding {
  option (cosmos_proto.implements_interface) = "AddressEncoding";

  // Prefix to be used
  string prefix = 1;
}

// HexEncoding represents the encoding algorithm based on the Hex format
message HexEncoding {
  option (cosmos_proto.implements_interface) = "AddressEncoding";
  
  // (optional) Prefix to be used
  string prefix = 1;
  
  // (optional) Whether the address should be upper case or not (default: false)
  bool uppercase = 2;
}
```

We will also define the following methods for the `Address` structure: 

```go
// Validate validates the given public key against this address, to make sure they match
func (a *Address) Validate(pubKey cryptotypes.PubKey) error {
	addressBytes, err := hex.DecodeString(a.Value)
	if err != nil {
		return err
	}
	
	// Generate the address bytes from the pub key
	generatedBytes, err := generateAddressBytes(pubKey, a.GenerationAlgorithm)
	if err != nil {
		return err
	}
	
	// Compare the bytes
	if !bytes.Equals(addressBytes, generatedBytes) {
		return fmt.Errrorf("address bytes do not match generated ones: expected %s but got %s", addressBytes, generatedBytes)	
	}
	
	return nil
}

// generateAddressBytes generates the address bytes starting from the given public key 
// and using the provided generation algorithm
func generateAddressBytes(pubKey cryptotypes.PubKey, generationAlgorithm GenerationAlgorithm) ([]byte, error) {
	// ...
}

// GetValue returns the string value of the address, encoded as it should be
func (a *Address) GetValue() string {
	return a.Encoding.GetCachedValue().(AddressEncoding).Encode(a.Value)
}
```


## Consequences

### Backwards Compatibility

Since this update will affect all the instances of `AddressData` by completely changing how it is defined, such changes are backwards incompatible. For this reason, we need to make sure we write a proper migration to update all the current `AddressData` instances to the new `Address` type. 

### Positive

- Easy to extend the address generation algorithm for `Address` instances

### Negative

### Neutral

## References

- [Cosmos address](https://docs.cosmos.network/master/architecture/adr-028-public-key-addresses.html#legacy-public-key-addresses-don-t-change)
- [Ethereum address](https://ethereum.org/en/developers/docs/accounts/#account-creation)
- [Solana address](https://docs.solana.com/terminology#account)
- [Elrond address](https://docs.elrond.com/technology/glossary/)