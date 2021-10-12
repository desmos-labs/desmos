# ADR 008: Change the plain text encoding of links to hex

## Changelog

- October 5th 2021: Initial draft;
- October 7th 2021: Moved from DRAFT to PROPOSED;

## Status

PROPOSED

## Abstract

Currently, when verifying an application link or a chain link we assume that their plain text values have been encoded using UTF-8.
However, there is a major problem with the UTF-8 encoding: it does not support all bytes properly. For this reason, we SHOULD change the encoding of the plain text from UTF-8 to HEX in order to avoid any possible encoding problem that would result in a signature that it's impossible to verify correctly. 

## Context

Desmos `profiles` module give the possibility to link the desmos profile to external account. There are two objects to prove 
the connection between them, which are application link for centralized network and chain link for blockchain network.
Both application link and chain link contains: 
 * A signature of the plain text signed by a private key;
 * The public key associated with the private key generating the signature. 
 In addition, the plain text used by the signature is assumed to be UTF-8 encoded 
but the problem occurs if the encoding of the plain text is another one such as UTF-16, Unicode and etc.

## Decision

We propose to change the encoding of the plain text of both application link and chain link to hex.

### The implementation of chain link

When saving a `ChainLink`, we use the `Proof` object in order to verify the signature. To make sure it supports the HEX encoding instead of the UTF-8 one, we need to change how the `Validate` method checks for the validity of such proof:
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
Second, we need to change how the `Proof#Verify` method verifies the signature provided in order to make sure that it deserializes the plain text as an HEX value instead of an UTF-8 one:
``go
// Verify verifies the signature using the given plain text and public key.
// It returns and error if something is invalid.
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
``

Finally, we need to modify the `generateChainLinkJSON` function to return a HEX encoded plain text:
```go
// generateChainLinkJSON returns build a new ChainLinkJSON intance using the provided mnemonic and chain configuration
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

### Chain link implementation

While dealing with application links, we use the `Result_Success_` type to identify a successfully verified link. In order to force the plain text to be HEX encoded, we need to modify the `Validate` function to perform such check:

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

Besides, we also need to change the function that is currently used by users to generate the signature using the Desmos CLI to make sure it returns the plain text using the HEX encoding:
```go
// GetSignCmd returns the command allowing to sign an arbitrary for later verification
func GetSignCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign [value]",
		Short: "Allows to sign the given value using the private key associated to the address or key specified using the --from flag",
		RunE: func(cmd *cobra.Command, args []string) error {
			
            ...

			value := args[0]
			bz, pubKey, err := txFactory.Keybase().Sign(key.GetName(), []byte(value))
			if err != nil {
				return err
			}

			// Build the signature data output
			signatureData := SignatureData{
				Address:   strings.ToLower(pubKey.Address().String()),
				Signature: strings.ToLower(hex.EncodeToString(bz)),
				PubKey:    strings.ToLower(hex.EncodeToString(pubKey.Bytes())),
				Value:     hex.EncodeToString([]byte(value)),
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

Currently, all the plain text in the old links are UTF-8 encoded, there is no problem with them since the signature was 
verified during the creation process and this ADR only targets to the new links. It can be kept consistent on-chain by the migration script
to transform all currently stored links into hex-encoded.
As a result, this feature is backwards compatible.

### Positive

* Ensure signing and verification process works properly with arbitrary strings

### Negative

(none known)

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#636](https://github.com/desmos-labs/desmos/issues/636)