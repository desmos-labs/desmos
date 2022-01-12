package chainlink_test

import (
	"io/ioutil"

	cmd "github.com/desmos-labs/desmos/v2/app/desmos/cmd/chainlink"
	"github.com/desmos-labs/desmos/v2/app/desmos/cmd/chainlink/types"
	"github.com/desmos-labs/desmos/v2/testutil"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	profilescliutils "github.com/desmos-labs/desmos/v2/x/profiles/client/utils"
	profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func BuildMockChainLinkJSONBuilderProvider(getter MockGetter) types.ChainLinkJSONBuilderProvider {
	return func(isSingleAccount bool) types.ChainLinkJSONBuilder {
		if isSingleAccount {
			return types.NewSingleAccountChainLinkJSONBuilder(getter)
		}
		return types.NewMultisigAccountChainLinkJSONBuilder(getter)
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
	_, err = keyBase.NewAccount(types.KeyName, mnemonic, "", "m/44'/118'/0'/0/0", hd.Secp256k1)
	suite.Require().NoError(err)

	// Get the key from the keybase
	key, err := keyBase.Key(types.KeyName)
	suite.Require().NoError(err)

	expected := profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address("cosmos13j7p6faa9jr8ty6lvqv0prldprr6m5xenmafnt", "cosmos"),
		profilestypes.NewProof(
			key.GetPubKey(),
			testutil.SingleSignatureProtoFromHex("c3bd014b2178d63d94b9c28e628bfcf56736de28f352841b0bb27d6fff2968d62c13a10aeddd1ebfe3b13f3f8e61f79a2c63ae6ff5cb78cb0d64e6b0a70fae57"),
			"636f736d6f7331336a377036666161396a72387479366c7671763070726c64707272366d3578656e6d61666e74"),
		profilestypes.NewChainConfig("cosmos"),
	)

	suite.Require().Equal(expected, data)
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
			"memo": "",
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
			"CkAv4+a/BrQeFNM2ETyv8w5NTRigi4N6qF+Ry5Vx9/C4RWBd4EesFQhm/KBKuzFWq6QFNolXd/SH0ZjyQDd/darECkAcEtkxg/x/0ZqZdud7eI3yvTMn1TKSiu+KawEHBgUsSFvyh8ViIAmu1nLUVEXUOuD+PBmAI0BG0LL9Lnwfwjmg"
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
			testutil.MultiSignatureProtoFromAnyHex(suite.Codec, "0a2b2f6465736d6f732e70726f66696c65732e763162657461312e4d756c74695369676e61747572654461746112f3010a0508031201c012740a2c2f6465736d6f732e70726f66696c65732e763162657461312e53696e676c655369676e6174757265446174611244087f12402fe3e6bf06b41e14d336113caff30e4d4d18a08b837aa85f91cb9571f7f0b845605de047ac150866fca04abb3156aba40536895777f487d198f240377f75aac412740a2c2f6465736d6f732e70726f66696c65732e763162657461312e53696e676c655369676e6174757265446174611244087f12401c12d93183fc7fd19a9976e77b788df2bd3327d532928aef8a6b010706052c485bf287c5622009aed672d45445d43ae0fe3c1980234046d0b2fd2e7c1fc239a0"),
			"7b226163636f756e745f6e756d626572223a2230222c22636861696e5f6964223a22636f736d6f73222c22666565223a7b22616d6f756e74223a5b5d2c22676173223a22323030303030227d2c226d656d6f223a22222c226d736773223a5b7b2274797065223a22636f736d6f732d73646b2f4d7367566f7465222c2276616c7565223a7b226f7074696f6e223a312c2270726f706f73616c5f6964223a2231222c22766f746572223a22636f736d6f73316578646a6b6678756438797a7174767561336864643933787530676d656b356c343772387261227d7d5d2c2273657175656e6365223a2230227d"),
		profilestypes.NewChainConfig("cosmos"),
	)

	suite.Require().Equal(expected, data)
}
