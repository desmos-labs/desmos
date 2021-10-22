# ADR 006: Support multisig chain link

## Changelog

- September 29th 2021: Initial draft
- October 21th, 2021: Proposed
- October 21th, 2021: First Review

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
// It returns and error if something is invalid.
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

In order to generate right chain link json for both single signature and multi signature account, we propose 
to separate `generateChainLinkJSON` into `generateChainLinkJSONForSinglesigAccount` and 
`generateChainLinkJSONForMultisigAccount`.

In `generateChainLinkJSONForSinglesigAccount`, we will change signature into `SingleSignatureData`
from simple single signature bytes. Subsequently, convert it into `SignatureDescriptor_Data`:
```go
// generateChainLinkJSONForSinglesigAccount returns build a new ChainLinkJSON instance using the provided a single mnemonic and chain configuration
func generateChainLinkJSONForSinglesigAccount(
	cdc codec.BinaryCodec,
	mnemonic string,
	chain chainlinktypes.Chain
) (profilescliutils.ChainLinkJSON, error) {
	// Create an in-memory keybase for signing
	keyBase := keyring.NewInMemory()
	keyName := "chainlink"
	_, err := keyBase.NewAccount(keyName, mnemonic, "", chain.DerivationPath, hd.Secp256k1)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	// Generate the proof signing it with the key
	key, _ := keyBase.Key(keyName)
	addr, _ := sdk.Bech32ifyAddressBytes(chain.Prefix, key.GetAddress())
	value := []byte(addr)
	sigBz, pubkey, err := keyBase.Sign(keyName, value)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}
	sigData := &signing.SingleSignatureData{
		SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
		Signature: sigBz,
	}
	sig := signing.SignatureDataToProto(sigData)

	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(pubkey, sig, hex.EncodeToString([]byte(addr))),
		profilestypes.NewChainConfig(chain.Name),
	), nil
}
```

In `generateChainLinkJSONForMultisigAccount`, it requires the mnemonics and threshold to generate the multisig address.
Then, signing the plain text in order to create the multi signatures. Subsequently, convert it into `SignatureDescriptor_Data`:
```go
// generateChainLinkJSONForMultisigAccount returns build a new ChainLinkJSON instance using the multisig reference and chain configuration
func generateChainLinkJSONForMultiAccount(
	cdc codec.BinaryCodec, mnemonics []string, 
	threshold int, 
	chain chainlinktypes.Chain
) (profilescliutils.ChainLinkJSON, error) {
	pubkeys := []cryptotypes.PubKey{}
	mSig := multisig.NewMultisig(len(mnemonics))

	// Create an in-memory keybase for signing and generating multisig
	keyBase := keyring.NewInMemory()
	for i, m := range mnemonics {
		keyName := "chainlink" + strconv.Itoa(i)
		key, err := keyBase.NewAccount(keyName, m, "", chain.DerivationPath, hd.Secp256k1)
		if err != nil {
			return profilescliutils.ChainLinkJSON{}, err
		}
		pubkeys = append(pubkeys, key.GetPubKey())
	}

	sort.Slice(pubkeys, func(i, j int) bool {
		return bytes.Compare(pubkeys[i].Address(), pubkeys[j].Address()) < 0
	})

	mPubkey := kmultisig.NewLegacyAminoPubKey(threshold, pubkeys)
	addr, _ := sdk.Bech32ifyAddressBytes(chain.Prefix, mPubkey.Address().Bytes())

	// Sign
	keys, _ := keyBase.List()
	for _, key := range keys {
		fmt.Println(key.GetAddress().String())
		sigBz, pubkey, err := keyBase.Sign(key.GetName(), []byte(addr))
		if err != nil {
			return profilescliutils.ChainLinkJSON{}, err
		}
		sigData := &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
			Signature: sigBz,
		}
		multisig.AddSignatureFromPubKey(mSig, sigData, pubkey, pubkeys)
	}

	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(mPubkey, signing.SignatureDataToProto(mSig), hex.EncodeToString([]byte(addr))),
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
