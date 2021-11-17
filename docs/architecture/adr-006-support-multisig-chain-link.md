# ADR 006: Support multisig chain link

## Changelog

- September 29th 2021: Initial draft
- October 21th, 2021: Proposed
- October 21th, 2021: First Review
- October 22th, 2021: Second Review
- November 15th, 2021: Third Review

## Status

PROPOSED

## Abstract

Currently, it is not possible to create a chain link using a multisig account. Since many validators MIGHT use multisig accounts, we SHOULD change the `Proof` type and its `Verify` function in order to support this kind of account as well.

## Context

Currently, the `x/profiles` module gives users the possibility to link their profile to different external accounts. 
In particular, to link other blockchains accounts to a profile, we follow this process:
1. the user signs a message with their own private key;
2. the signature and the signed value are placed inside a `Proof` object;
3. Desmos verifies the `Proof` object to guarantee that the user really owns such account and thus it can be linked successfully to their Desmos profile.

Currently, this process works properly for single-signature accounts, but it does not support multi-signature accounts. This is due to the fact that the `Proof` type only supports single-account's signatures, and its `Verify` function is only able to verify such signature type.

## Decision

We propose to change the `Proof#Signature` field to be of type [`SignatureDescriptor_Data`](https://github.com/cosmos/cosmos-sdk/blob/master/proto/cosmos/tx/signing/v1beta1/signing.proto#L57).

### Proof implementation

`SignatureDescriptor_Data` supports storing both single and multi-sig signatures. Moreover, the Cosmos SDK provides a [function](https://github.com/cosmos/cosmos-sdk/blob/master/types/tx/signing/signature.go#L65) to convert such type into the corresponding interface ([`SignatureData`](https://github.com/cosmos/cosmos-sdk/blob/master/types/tx/signing/signature_data.go#L10)), which is required from the `Verify` function in order to validate such signature. 

The verification process can be implemented as follows:
1. If it's a `SingleSignatureData`, make sure the account public key is a cryptotypes. `PubKey` and 
then use the `VerifySignature` method to verify the signature.
2. If it's a `MultiSignatureData`, make sure the account public key is a `multisig.PubKey` and 
then use the `VerifyMultisignature` method to verify the signature.

### CLI implementation

We propose to use an interactive prompt to create a new chain-link JSON for both single signature and multi signature account. For the single signature account, the process will remain the same as before. For the multi signature account, we will add the new feature beyond the remained interactive prompt.

To simplify things for multisig accounts, we propose to create a chain link JSON starting from a transaction signed with a multisig account. Inside the Cosmos SDK, multisig account transactions signing depends on the `tx sign` and 
`tx multisign` commands. To send a transaction to the node, the following process takes place: 
1. the multisig account owners create a raw transaction file
2. the threshold number of signers sign it to using the `tx sign` command
3. all individual signatures are gathered and used as input to the `tx multisign` the signed transaction. 

The multisign file includes not only all public keys and threshold of the multisig account, but also the required number of signatures. Thanks to this, we can leverage all the information stored inside such file to create a proper chain link JSON starting from a multisigned transaction.

The whole process in code is presented below:
```go
// getChainLinkJSONFromMultiSign generates the chain-link JSON from the multisign file and its raw transaction file
func getChainLinkJSONFromMultiSign(
	cmd *cobra.Command,
	txFile string,
	chainID string,
	chain chainlinktypes.Chain,
) (profilescliutils.ChainLinkJSON, error) {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	parsedTx, err := authclient.ReadTxFromFile(clientCtx, txFile)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	txCfg := clientCtx.TxConfig
	txBuilder, err := txCfg.WrapTxBuilder(parsedTx)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	txFactory := tx.NewFactoryCLI(clientCtx, cmd.Flags())
	if txFactory.SignMode() == signingtypes.SignMode_SIGN_MODE_UNSPECIFIED {
		txFactory = txFactory.WithSignMode(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON)
	}

	sigs, err := txBuilder.GetTx().GetSignaturesV2()
	if len(sigs) != 1 {
		return profilescliutils.ChainLinkJSON{}, fmt.Errorf("invalid number of signatures")
	}
	multisigSig := sigs[0]

	signingData := authSigning.SignerData{
		ChainID:       chainID,
		AccountNumber: txFactory.AccountNumber(),
		Sequence:      txFactory.Sequence(),
	}
	// the bytes of plain text
	value, err := txCfg.SignModeHandler().GetSignBytes(txFactory.SignMode(), signingData, txBuilder.GetTx())
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	addr, _ := sdk.Bech32ifyAddressBytes(chain.Prefix, multisigSig.PubKey.Address().Bytes())
	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(multisigSig.PubKey, signingtypes.SignatureDataToProto(multisigSig.Data), hex.EncodeToString(value)),
		profilestypes.NewChainConfig(chain.Name),
	), nil
}
```

## Consequences

### Backwards Compatibility

Since we change the `Proof#Signature` field from a hex-encode string into a `SignatureDescriptor_Data`, this feature is not backwards compatible. Migrating the old chain link signatures to `SignatureDescriptor_Data`
is required.

### Positive

- Give the possibility to link a multisig account to a Desmos profile

### Negative

- Raise the complexity to generate and verify the signature

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

The following tests cases MUST to be present:
- Verify `Proof` including wrong format signature returns error
- Verify `Proof` including non-matched single signature and pubkey returns error
- Verify `Proof` including non-matched multi signatures and pubkeys returns error
- Verify `Proof` including proper single signature and pubkey returns no error
- Verify `Proof` including proper multi signatures and pubkeys returns no error


## References

- Issue [#633](https://github.com/desmos-labs/desmos/issues/633)
- [SignatureData](https://github.com/cosmos/cosmos-sdk/blob/master/types/tx/signing/signature_data.go)
- [Signature](https://github.com/cosmos/cosmos-sdk/blob/master/types/tx/signing/signature.go)
- [Multisig](https://github.com/cosmos/cosmos-sdk/blob/master/crypto/keys/multisig/multisig.go)
