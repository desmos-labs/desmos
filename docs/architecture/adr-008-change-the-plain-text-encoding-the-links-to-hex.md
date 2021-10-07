# ADR 008: Change the plain text encoding of links to hex

## Changelog

- October 5th 2021: Initial draft

## Status

DRAFT

## Abstract

Currently, when verifying an application link or a chain link, we assume that the plain text value of them has been encoded using UTF-8.
Unfortunately, Ledger or common wallets does not support to sign something due to the fact that they do not allow to sign arbitrary strings.
We SHOULD change encoding of the plain text to hex in order to avoid this problem. 

## Context

Desmos `profiles` module give the possibility to link the desmos profile to external account. The object to prove 
the connection of themare application link for centralized applications accounts and chain link for other blockchains accounts.
Both application link and chain link contains a signature signed with the plain text by a private key and a public key 
from the private key generating the signature. In addition, the plain text used by the signature is assumed as UTF-8 encoded 
but it occurs the problem if the encoding of plain text is others like UTF-16, Unicode and etc.

## Decision

We propose to change the encoding of the plain text of both application link and chain link to hex.

### The implementation of chain link

In chain link, `Proof` is a object to contains the data related verification. We will check the plain text if it is hex-encoded.
The `Validate` function will be like: 
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
In addition, the function to generate chain link json in CLI tool is `generateChainLinkJSON`. 
We will modify it to use hex-encoded plain text, it will be like:
```go
// generateChainLinkJSON returns build a new ChainLinkJSON intance using the provided mnemonic and chain configuration
func generateChainLinkJSON(mnemonic string, chain chainlinktypes.Chain) (profilescliutils.ChainLinkJSON, error) {
	
    ...

	plainText := hex.EncodeToString([]byte(addr))
	sig, pubkey, err := keyBase.Sign(keyName, []byte(plainText))
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(pubkey, hex.EncodeToString(sig), plainText),
		profilestypes.NewChainConfig(chain.Name),
	), nil
}
```

### The implementation of application link

In application link, the object to show the proof is `Result_Success_` which is a sub object inside `Result`.
We will modify the `Validate` function to ensure the plain text is hex-encoded.

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

Besides, the function to generate signature in the CLI tool is `GetSignCmd`.
We will use hex-encoded plain text inside it, it will be like:
```go
// GetSignCmd returns the command allowing to sign an arbitrary for later verification
func GetSignCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign [value]",
		Short: "Allows to sign the given value using the private key associated to the address or key specified using the --from flag",
		RunE: func(cmd *cobra.Command, args []string) error {
			
            ...

			// Sign the data with the private key
			value := args[0]
            plainText := hex.EncodeToString([]byte(value))
			bz, pubKey, err := txFactory.Keybase().Sign(key.GetName(), []byte(plainText))
			if err != nil {
				return err
			}

			// Build the signature data output
			signatureData := SignatureData{
				Address:   strings.ToLower(pubKey.Address().String()),
				Signature: strings.ToLower(hex.EncodeToString(bz)),
				PubKey:    strings.ToLower(hex.EncodeToString(pubKey.Bytes())),
				Value:     plainText,
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

Currently, all the plain text in the old links are UTF-8 encoded. The verification of old links will work properly 
since this adr only changes the plain text encoding, not the verification process.
As a result, it is backwards compatibility.

### Positive

* Ensure signing and verification process works properly

### Negative

(none known)

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#636](https://github.com/desmos-labs/desmos/issues/636)