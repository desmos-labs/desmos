# Chain link

## Upgradation guide of test cases

### Multisignature chain link account

After the proto version of `x/profiles` upgrades, the test case of multi signature chain link account test case must be updated to the specified version as well since the type uri of it will be `/desmos.profiles.{new-proto-version}.CosmosMultiSignature`. The process of upgrading is as follows:

#### 1. Create a tx.json with the content as follows:

```json
{
    "body": {
        "messages": [
            {
                "@type": "/cosmos.gov.v1beta1.MsgVote",
                "proposal_id": "1",
                "voter": "cosmos1exdjkfxud8yzqtvua3hdd93xu0gmek5l47r8ra",
                "option": "VOTE_OPTION_YES"
            }
        ],
        "memo": "desmos1n8345tvzkg3jumkm859r2qz0v6xsc3henzddcj",
        "timeout_height": "0",
        "extension_options": [],
        "non_critical_extension_options": []
    },
    "auth_info": {
        "signer_infos": [
            {
                "public_key": {
                    "@type": "/cosmos.crypto.multisig.LegacyAminoPubKey",
                    "threshold": 2,
                    "public_keys": [
                        {
                            "@type": "/cosmos.crypto.secp256k1.PubKey",
                            "key": "A4k1o4weHTkMVXqzT0zKRkRmWTwQEh3JGiPkJvCQ4VO7"
                        },
                        {
                            "@type": "/cosmos.crypto.secp256k1.PubKey",
                            "key": "ApfZ2jzyWcRxzgCnKEKr+oIMyrGIJMp+1FjouYPovluE"
                        },
                        {
                            "@type": "/cosmos.crypto.secp256k1.PubKey",
                            "key": "AziL8Ly6QrMOr+V7Vf6XCBjDjLPTq0Dtxv7PPzDRnFQe"
                        }
                    ]
                },
                "mode_info": {
                    "multi": {
                        "bitarray": {
                            "extra_bits_stored": 3,
                            "elems": "wA=="
                        },
                        "mode_infos": [
                            {
                                "single": {
                                    "mode": "SIGN_MODE_LEGACY_AMINO_JSON"
                                }
                            },
                            {
                                "single": {
                                    "mode": "SIGN_MODE_LEGACY_AMINO_JSON"
                                }
                            }
                        ]
                    }
                },
                "sequence": "0"
            }
        ],
        "fee": {
            "amount": [],
            "gas_limit": "200000",
            "payer": "",
            "granter": ""
        }
    },
    "signatures": [
        "CkAn/EVngYopgD7BP0KUBMcTHIGKzBlU9RLz1xozeefsdB0l3osUL2EVFlKwbveKrv/VhwcCPm6N++mMmQGFAWR2CkCTlMhmMOevuWGJmt2PwaIR0UuMw4cCxTyqcBhRVX81gywR4RUQ2k1nZXihmzQoZTF1R1SbK0vXjN+Ana+lUEH3"
    ]
}
```

#### 2. Execute the latest desmos cli with the command `desmos create-link-json` with `tx.json` to create chain link proof data `data.json`.

```bash
desmos create-chain-link-json
```

#### 3. Produce the `Any` value of `Signature` in hex encoding with the `data.json` produce at the 2nd step by the script:

```go
chainLinkJSON, err := os.ReadFile("path/data.json")
if err != nil {
    panic(err)
}
cdc, legacyCdc := app.MakeCodecs()
var link profilescliutils.ChainLinkJSON
if err := cdc.UnmarshalJSON(chainLinkJSON, &link); err != nil {
    panic(err)
}
var addrData profilestypes.AddressData
if err := cdc.UnpackAny(link.Address, &addrData); err != nil {
    panic(err)
}
if err := link.Proof.Verify(cdc, legacyCdc, "desmos1n8345tvzkg3jumkm859r2qz0v6xsc3henzddcj", addrData); err != nil {
    panic(err)
}
bz, err := link.Proof.Signature.Marshal()
if err != nil {
    panic(err)
}
fmt.Println(hex.EncodeToString(bz)) 
```

output:
```bash
0a282f6465736d6f732e70726f66696c65732e76332e436f736d6f734d756c74695369676e617475726512ed010a0508031201c012710a292f6465736d6f732e70726f66696c65732e76332e436f736d6f7353696e676c655369676e61747572651244087f124027fc4567818a29803ec13f429404c7131c818acc1954f512f3d71a3379e7ec741d25de8b142f61151652b06ef78aaeffd58707023e6e8dfbe98c99018501647612710a292f6465736d6f732e70726f66696c65732e76332e436f736d6f7353696e676c655369676e61747572651244087f12409394c86630e7afb961899add8fc1a211d14b8cc38702c53caa701851557f35832c11e11510da4d676578a19b342865317547549b2b4bd78cdf809dafa55041f7
```

#### 4. Replace the signature field of the multisignature account test case with the 3rd output value

```go
expected := profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address("cosmos1exdjkfxud8yzqtvua3hdd93xu0gmek5l47r8ra", "cosmos"),
		profilestypes.NewProof(
			suite.GetPubKeyFromTxFile(txFile),
			profilestesting.MultiCosmosSignatureFromHex(
				suite.Codec,
				"output value of 3rd",
			),
			"7b226163636f756e745f6e756d626572223a2230222c22636861696e5f6964223a22636f736d6f73222c22666565223a7b22616d6f756e74223a5b5d2c22676173223a22323030303030227d2c226d656d6f223a226465736d6f73316e3833343574767a6b67336a756d6b6d3835397232717a3076367873633368656e7a6464636a222c226d736773223a5b7b2274797065223a22636f736d6f732d73646b2f4d7367566f7465222c2276616c7565223a7b226f7074696f6e223a312c2270726f706f73616c5f6964223a2231222c22766f746572223a22636f736d6f73316578646a6b6678756438797a7174767561336864643933787530676d656b356c343772387261227d7d5d2c2273657175656e6365223a2230227d",
		),
		profilestypes.NewChainConfig("cosmos"),
	)
```