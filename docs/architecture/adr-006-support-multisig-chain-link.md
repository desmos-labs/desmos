# ADR 006: Support multisig chain link

## Changelog

- September 29th 2021: Initial draft
- October 21th, 2021: Proposed
- October 21th, 2021: First Review
- October 22th, 2021: Second Review

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

Currently, this process works properly for single-signature accounts, but it does not support multi-signature accounts. This is due to the fact that the `Proof` type only supports signatures made by a single-signature account, and its `Verify` function is only able to verify such signature type.

## Decision

We propose to change the `Signature` of `Proof` into be 
[SignatureDescriptor_Data](https://github.com/cosmos/cosmos-sdk/blob/master/proto/cosmos/tx/signing/v1beta1/signing.proto#L57) instance.

### Proof implementation

`SignatureDescriptor_Data` supports to store both single and multi signatures. Moreover, it has a 
[function](https://github.com/cosmos/cosmos-sdk/blob/master/types/tx/signing/signature.go#L65) 
to convert to [SignatureData](https://github.com/cosmos/cosmos-sdk/blob/master/types/tx/signing/signature_data.go#L10), 
which is helpful to the signature verification in `Verify` function. The verification process will be like:
1. If it's a `SingleSignatureData`, make sure the account public key is a cryptotypes. `PubKey` and 
then use the `VerifySignature` method to verify the signature.
2. If it's a `MultiSignatureData`, make sure the account public key is a `multisig.PubKey` and 
then use the `VerifyMultisignature` method to verify the signature.

The whole process in code will be like:
```go
type Proof struct {
    PlainText []byte
    Signature *SignatureDescriptor_Data 
    PublicKey PubKey
}

// Verify verifies the signature using the given plain text and public key.
// It returns an error if something is invalid.
func (p Proof) Verify(cdc codec.BinaryCodec, address AddressData) error {
	value, _ := hex.DecodeString(p.PlainText)
	var pubkey cryptotypes.PubKey

	switch sigData := (signing.SignatureDataFromProto(p.Signature)).(type) {
	case *signing.SingleSignatureData:
		err := cdc.UnpackAny(p.PubKey, &pubkey)
		if err != nil {
			return fmt.Errorf("failed to unpack the public key")
		}
		if !pubkey.VerifySignature(value, sigData.Signature) {
			return fmt.Errorf("failed to verify the signature")
		}

	case *signing.MultiSignatureData:
		var multiPubkey multisig.PubKey
		err := cdc.UnpackAny(p.PubKey, &multiPubkey)
		if err != nil {
			return fmt.Errorf("failed to unpack the public key")
		}
		if err := multiPubkey.VerifyMultisignature(
			func(mode signing.SignMode) ([]byte, error) {
				return value, nil
			},
			sigData,
		); err != nil {
			return err
		}
		pubkey = multiPubkey
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

### CLI implementation

We propose to use an interactive prompt to create a new chain-link JSON for both single signature and multi signature account. For the single signature account, the process
will remain the same as before. For the multi signature account, we will add the new feature beyond the remained interactive prompt.

```go
// GetCreateChainLinkJSON returns the command allowing to generate the chain-link JSON
// file that is required by the link-chain command
func GetCreateChainLinkJSON(getter chainlinktypes.ChainLinkReferenceGetter) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-chain-link-json",
		Short: "Start an interactive prompt to create a new chain-link JSON object",
		Long: `Start an interactive prompt to create a new chain-link JSON object that can be used to later link your Desmos profile to another chain.
Once you have built the JSON object using this command, you can then run the following command to complete the linkage:

desmos tx profiles link-chain [/path/to/json/file.json]

Note that this command will ask you the mnemonic that should be used to generate the private key of the address you want to link.
The mnemonic is only used temporarily and never stored anywhere.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get the data
			accountType, err := getter.GetAccountType()
			if err != nil {
				return err
			}

			var chainLinkJson profilescliutils.ChainLinkJSON
			if accountType == "single" {
				mnemonic, err := getter.GetMnemonic()
				if err != nil {
					return err
				}

				chainLinkJSON, err := createSingleSigChainLinkJSON(getter)
			}

			if accountType == "multisig" {
				chainLinkJSON, err := createMultiSigChainLinkJSON(cmd, getter)
				if err != nil {
					return err
				}
			}

			filename, err := getter.GetFilename()
			if err != nil {
				return err
			}

			cdc, _ := app.MakeCodecs()
			bz, err := cdc.MarshalJSON(&chainLinkJSON)
			if err != nil {
				return err
			}

			if filename != "" {
				err = ioutil.WriteFile(filename, bz, 0600)
				if err != nil {
					return err
				}
			}
			cmd.Println(fmt.Sprintf("Chain link JSON file stored at %s", filename))
			return nil
		},
	}
	cmd.Flags().String(flags.FlagChainID, "", "network chain ID")
	return cmd
} 

// createSingleSigAccountChainLinkJSON creates the chain-link JSON for single signature account via getter
func createSingleSigChainLinkJSON(
	getter chainlinktypes.ChainLinkReferenceGetter
) (profilescliutils.ChainLinkJSON, error) {
	mnemonic, err := getter.GetMnemonic()
	if err != nil {
		return err
	}

	chain, err := getter.GetChain()
	if err != nil {
		return err
	}

	return generateChainLinkJSON(mnemonic, chain)
}

// createMultiSigChainLinkJSON creates the chain-link JSON for multi signature account via getter
func createMultiSigChainLinkJSON(
	cmd *cobra.Command, getter chainlinktypes.ChainLinkReferenceGetter
)(profilescliutils.ChainLinkJSON, error) {

	txFile, err := getter.GetTxFile()
	if err != nil {
		return err
	}

	multisignFile, err := getter.GetMultiSignFile()
	if err != nil {
		return err
	}

	chainID, err := getter.GetChainID()
	if err != nil {
		return err
	}

	return getChainLinkJSONFromMultiSign(
		cmd,
		txFile,
		multiSign, 
		chain,
	)
}
```

We propose to create chain-link JSON from the multisign file. In cosmos-sdk, multisig account transaction signing depends on `tx sign` and 
`tx multisign` commands. To send a transaction to the node, the multisig account owners should create a raw transaction file, then the threshold 
number of them sign it to generate the signed files by using the `tx sign` command with their keys. Subsequently, one of them gathers all the signed 
files and uses the `tx multisign` command with the raw transaction file to get the multisign file. The multisign file includes not only all public keys and 
threshold of account but also the required number of signatures so that creating the chain link json from the multisign file is possible.

The whole process in code is presented below:
```go
// getChainLinkJSONFromMultiSign generates the chain-link JSON from the multisign file and its raw transaction file
func getChainLinkJSONFromMultiSign(
	cmd *cobra.Command,
	txFile string,
	multiSignFile string,
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

	parsedMultiTx, err := authclient.ReadTxFromFile(clientCtx, multiSignFile)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	txCfg := clientCtx.TxConfig
	txBuilder, err := txCfg.WrapTxBuilder(parsedTx)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	multiTxBuilder, err := txCfg.WrapTxBuilder(parsedMultiTx)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	txFactory := tx.NewFactoryCLI(clientCtx, cmd.Flags())
	if txFactory.SignMode() == signingtypes.SignMode_SIGN_MODE_UNSPECIFIED {
		txFactory = txFactory.WithSignMode(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON)
	}

	sigs, err := multiTxBuilder.GetTx().GetSignaturesV2()
	if len(sigs) != 1 {
		return profilescliutils.ChainLinkJSON{}, fmt.Errorf("invalid number of signature")
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

	sigData := sigs[0].Data

	addr, _ := sdk.Bech32ifyAddressBytes(chain.Prefix, multisigSig.PubKey.Address().Bytes())
	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(multisigSig.PubKey, signingtypes.SignatureDataToProto(sigData), hex.EncodeToString(value)),
		profilestypes.NewChainConfig(chain.Name),
	), nil
}
```

## Consequences

### Backwards Compatibility

We change `Signature` of `Proof` into `SignatureDescriptor_Data` from hex-encode string of signature bytes.
This feature is not backwards compatible. Migrating the old chain link signature to `SignatureDescriptor_Data`
is required.

### Positive

* Give the possibility to link multisig account to desmos profile

### Negative

* Raise the complexity to generate and verify the signature

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

The following tests cases MUST to be present:
* Verify `Proof` including wrong format signature returns error
* Verify `Proof` including non-matched single signature and pubkey returns error
* Verify `Proof` including non-matched multi signatures and pubkeys returns error
* Verify `Proof` including proper single signature and pubkey returns no error
* Verify `Proof` including proper multi signatures and pubkeys returns no error


## References

- Issue [#633](https://github.com/desmos-labs/desmos/issues/633)
- [SignatureData](https://github.com/cosmos/cosmos-sdk/blob/master/types/tx/signing/signature_data.go)
- [Signature](https://github.com/cosmos/cosmos-sdk/blob/master/types/tx/signing/signature.go)
- [Multisig](https://github.com/cosmos/cosmos-sdk/blob/master/crypto/keys/multisig/multisig.go)
