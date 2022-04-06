package chainlink_test

import (
	"encoding/hex"
	"io/ioutil"

	"github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink/builder"

	cmd "github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink"
	"github.com/desmos-labs/desmos/v3/testutil"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	multibuilder "github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink/builder/multi"
	singlebuilder "github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink/builder/single"
	profilescliutils "github.com/desmos-labs/desmos/v3/x/profiles/client/utils"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func BuildMockChainLinkJSONBuilderProvider(getter MockGetter) builder.ChainLinkJSONBuilderProvider {
	return func(owner string, isSingleAccount bool) builder.ChainLinkJSONBuilder {
		if isSingleAccount {
			return singlebuilder.NewAccountChainLinkJSONBuilder(owner, getter)
		}
		return multibuilder.NewAccountChainLinkJSONBuilder(getter)
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *CreateJSONChainLinkTestSuite) TestSingleSignatureAccount() {
	fileName := suite.TempFile()
	getter := NewMockGetter(fileName, true, "")
	_, err := clitestutil.ExecTestCLICmd(
		suite.ClientCtx,
		cmd.GetCreateChainLinkJSON(getter, BuildMockChainLinkJSONBuilderProvider(getter)),
		[]string{},
	)
	suite.Require().NoError(err)

	out, err := ioutil.ReadFile(fileName)
	suite.Require().NoError(err)

	var data profilescliutils.ChainLinkJSON
	err = suite.Codec.UnmarshalJSON(out, &data)
	suite.Require().NoError(err)

	// Create an account inside the inmemory keybase
	keyBase := keyring.NewInMemory()
	mnemonic := "clip toilet stairs jaguar baby over mosquito capital speed mule adjust eye print voyage verify smart open crack imitate auto gauge museum planet rebel"
	_, err = keyBase.NewAccount(singlebuilder.KeyName, mnemonic, "", "m/44'/118'/0'/0/0", hd.Secp256k1)
	suite.Require().NoError(err)

	// Get the key from the keybase
	key, err := keyBase.Key(singlebuilder.KeyName)
	suite.Require().NoError(err)

	sig, _, err := keyBase.Sign(singlebuilder.KeyName, []byte(suite.Owner))
	suite.Require().NoError(err)

	expected := profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address("cosmos13j7p6faa9jr8ty6lvqv0prldprr6m5xenmafnt", "cosmos"),
		profilestypes.NewProof(
			key.GetPubKey(),
			testutil.SingleSignatureProtoFromHex(hex.EncodeToString(sig)),
			hex.EncodeToString([]byte(suite.Owner))),
		profilestypes.NewChainConfig("cosmos"),
	)

	suite.Require().Equal(expected, data)
	suite.Require().NoError(
		data.Proof.Verify(
			suite.Codec,
			suite.LegacyAmino,
			suite.Owner,
			profilestypes.NewBech32Address("cosmos13j7p6faa9jr8ty6lvqv0prldprr6m5xenmafnt", "cosmos"),
		),
	)
}

func (suite *CreateJSONChainLinkTestSuite) TestMultiSignatureAccount() {
	fileName := suite.TempFile()
	txFileData := `{
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
	}`

	// Define where the data will be saved
	txFile := suite.TempFile()
	err := ioutil.WriteFile(txFile, []byte(txFileData), 0600)
	suite.Require().NoError(err)

	getter := NewMockGetter(fileName, false, txFile)
	_, err = clitestutil.ExecTestCLICmd(
		suite.ClientCtx,
		cmd.GetCreateChainLinkJSON(getter, BuildMockChainLinkJSONBuilderProvider(getter)),
		[]string{},
	)
	suite.Require().NoError(err)

	out, err := ioutil.ReadFile(fileName)
	suite.Require().NoError(err)

	var data profilescliutils.ChainLinkJSON
	err = suite.Codec.UnmarshalJSON(out, &data)
	suite.Require().NoError(err)

	expected := profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address("cosmos1exdjkfxud8yzqtvua3hdd93xu0gmek5l47r8ra", "cosmos"),
		profilestypes.NewProof(
			suite.GetPubKeyFromTxFile(txFile),
			testutil.MultiSignatureProtoFromAnyHex(
				suite.Codec,
				"0a262f6465736d6f732e70726f66696c65732e76322e4d756c74695369676e61747572654461746112e9010a0508031201c0126f0a272f6465736d6f732e70726f66696c65732e76322e53696e676c655369676e6174757265446174611244087f124027fc4567818a29803ec13f429404c7131c818acc1954f512f3d71a3379e7ec741d25de8b142f61151652b06ef78aaeffd58707023e6e8dfbe98c990185016476126f0a272f6465736d6f732e70726f66696c65732e76322e53696e676c655369676e6174757265446174611244087f12409394c86630e7afb961899add8fc1a211d14b8cc38702c53caa701851557f35832c11e11510da4d676578a19b342865317547549b2b4bd78cdf809dafa55041f7",
			),
			"7b226163636f756e745f6e756d626572223a2230222c22636861696e5f6964223a22636f736d6f73222c22666565223a7b22616d6f756e74223a5b5d2c22676173223a22323030303030227d2c226d656d6f223a226465736d6f73316e3833343574767a6b67336a756d6b6d3835397232717a3076367873633368656e7a6464636a222c226d736773223a5b7b2274797065223a22636f736d6f732d73646b2f4d7367566f7465222c2276616c7565223a7b226f7074696f6e223a312c2270726f706f73616c5f6964223a2231222c22766f746572223a22636f736d6f73316578646a6b6678756438797a7174767561336864643933787530676d656b356c343772387261227d7d5d2c2273657175656e6365223a2230227d",
		),
		profilestypes.NewChainConfig("cosmos"),
	)
	suite.Require().Equal(expected, data)
	suite.Require().NoError(
		data.Proof.Verify(
			suite.Codec,
			suite.LegacyAmino,
			suite.Owner,
			profilestypes.NewBech32Address("cosmos1exdjkfxud8yzqtvua3hdd93xu0gmek5l47r8ra", "cosmos"),
		),
	)
}
