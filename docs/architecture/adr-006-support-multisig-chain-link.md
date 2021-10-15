# ADR 006: Support multisig chain link

## Changelog

- September 29th 2021: Initial draft

## Status

DRAFT

## Abstract

> "If you can't explain it simply, you don't understand it well enough." Provide a simplified and layman-accessible explanation of the ADR.
> A short (~200 word) description of the issue being addressed.

## Context

> This section describes the forces at play, including technological, political, social, and project local. These forces are probably in tension, and should be called out as such. The language in this section is value-neutral. It is simply describing facts. It should clearly explain the problem and motivation that the proposal aims to resolve.
> {context body}

## Decision

> This section describes our response to these forces. It is stated in full sentences, with active voice. "We will ..."
> {decision body}


```go
type Proof struct {
    PlainText []byte
    Signature string // hex-encoded string of signing.SignatureDescriptor_Data
    PublicKey PubKey
}

// Verify verifies the signature using the given plain text and public key.
// It returns and error if something is invalid.
func (p Proof) Verify(cdc codec.BinaryCodec, address AddressData) error {
	sigBz, _ := hex.DecodeString(p.Signature)
	var sigProto signing.SignatureDescriptor_Data
	if err := cdc.Unmarshal(sigBz, &sigProto); err != nil {
		return err
	}
	switch sigData := (signing.SignatureDataFromProto(&sigProto)).(type) {
	case *signing.SingleSignatureData:
		value, _ := hex.DecodeString(p.PlainText)
		var pubkey cryptotypes.PubKey
		err := cdc.UnpackAny(p.PubKey, &pubkey)
		if err != nil {
			return fmt.Errorf("failed to unpack the public key")
		}
		if !pubkey.VerifySignature(value, sigData.Signature) {
			return fmt.Errorf("failed to verify the signature")
		}
		valid, err := address.VerifyPubKey(pubkey)
		if err != nil {
			return err
		}
		if !valid {
			return fmt.Errorf("invalid address and public key combination provided")
		}

	case *signing.MultiSignatureData:
		value, _ := hex.DecodeString(p.PlainText)
		var pubkey multisig.PubKey
		err := cdc.UnpackAny(p.PubKey, &pubkey)
		if err != nil {
			return fmt.Errorf("failed to unpack the public key")
		}
		if err := pubkey.VerifyMultisignature(
			func(mode signing.SignMode) ([]byte, error) {
				return value, nil
			},
			sigData,
		); err != nil {
			return err
		}

		valid, err := address.VerifyPubKey(pubkey)
		if err != nil {
			return err
		}
		if !valid {
			return fmt.Errorf("invalid address and public key combination provided")
		}
	}
	return nil
}
```


## Consequences

> This section describes the resulting context, after applying the decision. All consequences should be listed here, not just the "positive" ones. A particular decision may have positive, negative, and neutral consequences, but all of them affect the team and project in the future.

### Backwards Compatibility

> All ADRs that introduce backwards incompatibilities must include a section describing these incompatibilities and their severity. The ADR must explain how the author proposes to deal with these incompatibilities. ADR submissions without a sufficient backwards compatibility treatise may be rejected outright.

### Positive

{positive consequences}

### Negative

{negative consequences}

### Neutral

{neutral consequences}

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section should contain a summary of issues to be solved in future iterations (usually referencing comments from a pull-request discussion).
Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

## Test Cases [optional]

Test cases for an implementation are mandatory for ADRs that are affecting consensus changes. Other ADRs can choose to include links to test cases if applicable.

## References

- {reference link}