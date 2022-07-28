# ADR 001: Change chain link plain text encoding to hex

## Changelog

- October 5th 2021: Initial draft;
- October 7th 2021: Moved from DRAFT to PROPOSED;
- October 10th 2021: First review;

## Status

ACCEPTED Implemented

## Abstract

Currently, when verifying an application link or a chain link we assume that their plain text values have been encoded using UTF-8.
However, there is a major problem with the UTF-8 encoding: it does not support all bytes properly. For this reason, we SHOULD change 
the encoding of the plain text from UTF-8 to HEX in order to avoid any possible encoding problem that would result in a signature 
that it's impossible to verify correctly. 

## Context

The `x/profiles` module gives users the possibility to link their profile to some external account(s), either centralized applications 
(eg. GitHub, Reddit, Twitter) or other blockchains (eg. Cosmos, Solana, Polkadot). When linking a Desmos profile to any of these accounts, 
we use a signature-based authentication process in order to make sure that the user controls such accounts. In both the applications 
and chains links, we currently expect the user to use the UTF-8 encoding when sending over the plain text used to create the signature. 
However, since the UTF-8 encoding is not able to correctly represent all bytes, there might be cases in which we end up with a signature 
that it's impossible to verify. For example, this is what happens if the original plain text was encoded before being signed with another 
encoding such as UTF-16, Unicode, etc.

## Decision

We propose to change the encoding of the plain text of both application link and chain link to hex.

### Chain link implementation

When saving a `ChainLink`, we use the `Proof` object in order to verify the signature. To make sure it supports the HEX encoding 
instead of the UTF-8 one, we need to change how the `Validate` method checks for the validity of such proof:
```go
// Validate checks the validity of the Proof
func (p Proof) Validate() error {
    
    ...

    _, err = hex.DecodeString(p.PlainText)
    if err != nil {
        return fmt.Errorf("invalid hex-encoded plain text")
    }

	return nil
}
```
Second, we need to change how the `Proof#Verify` method verifies the signature provided in order to make sure that it deserializes
the plain text as an HEX value instead of an UTF-8 one:
```go
// Verify verifies the signature using the given plain text and public key.
// It returns an error if something is invalid.
func (p Proof) Verify(unpacker codectypes.AnyUnpacker, address AddressData) error {
	var pubkey cryptotypes.PubKey
	err := unpacker.UnpackAny(p.PubKey, &pubkey)
	if err != nil {
		return fmt.Errorf("failed to unpack the public key")
	}

	value, err := hex.DecodeString(p.PlainText)
	if err != nil {
		return fmt.Errorf("invalid hex-encoded plain text")
	}
	
	sig, err := hex.DecodeString(p.Signature)
	if err != nil {
		return fmt.Errorf("invalid hex-encoded signature")
	}

	if !pubkey.VerifySignature(value, sig) {
		return fmt.Errorf("failed to verify the signature")
	}

	valid, err := address.VerifyPubKey(pubkey)
	if err != nil {
		return err
	}

	if !valid {
		return fmt.Errorf("invalid address and public key combination provided")
	}

	return nil
}
```

Finally, we need to modify the `generateChainLinkJSON` function to return a HEX encoded plain text:
```go
// generateChainLinkJSON returns build a new ChainLinkJSON instance using the provided mnemonic and chain configuration
func generateChainLinkJSON(mnemonic string, chain chainlinktypes.Chain) (profilescliutils.ChainLinkJSON, error) {
	
    ...

	sig, pubkey, err := keyBase.Sign(keyName, []byte(value)) 
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(pubkey, hex.EncodeToString(sig), hex.EncodeToString([]byte(value))),
		profilestypes.NewChainConfig(chain.Name),
	), nil
}
```

### Application link implementation

While dealing with application links, we use the `Result_Success_` type to identify a successfully verified link. In order to force 
the plain text to be HEX encoded, we need to modify the `Validate` function to perform such check:

```go
// Validate returns an error if the instance does not contain valid data
func (r Result_Success_) Validate() error {
	
    ...

    if _, err := hex.DecodeString(r.Value); err != nil {
        return fmt.Errorf("invalid hex-encoded plain text")
    }

	return nil
}
```

Besides, we also need to change the function that is currently used by users to generate the signature using the Desmos CLI to 
make sure it returns the plain text using the HEX encoding:
```go
// GetSignCmd returns the command allowing to sign an arbitrary value for later verification
func GetSignCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign [value]",
		Short: "Allows to sign the given value using the private key associated to the address or key specified using the --from flag",
		RunE: func(cmd *cobra.Command, args []string) error {
			
            ...

			value := []byte(args[0])
			bz, pubKey, err := txFactory.Keybase().Sign(key.GetName(), value)
			if err != nil {
				return err
			}

			// Build the signature data output
			signatureData := SignatureData{
				Address:   strings.ToLower(pubKey.Address().String()),
				Signature: strings.ToLower(hex.EncodeToString(bz)),
				PubKey:    strings.ToLower(hex.EncodeToString(pubKey.Bytes())),
				Value:     hex.EncodeToString(value),
			}

			// Serialize the output as JSON and print it
			bz, err = json.Marshal(&signatureData)
			if err != nil {
				return err
			}

			return clientCtx.PrintBytes(bz)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
```
## Consequences

### Backwards Compatibility

With this approach there SHOULD not be any problem with old chain and application links since since the signature was 
verified during the creation process and this ADR only targets the new links that will be created. However, in order to 
make sure that clients can verify all the links at the same way, we SHOULD keep the on-chain data consistent using a migration script 
that transforms all currently stored plain texts from being UTF-8 encoded strings into HEX encoded strings.
As a result, this feature is backwards compatible.

### Positive

- Ensure the signing and verification process for chain and application links work properly with arbitrary bytes

### Negative

(none known)

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#636](https://github.com/desmos-labs/desmos/issues/636)
- [The basics of UTF-8](https://www.codeguru.com/cplusplus/the-basics-of-utf-8/)