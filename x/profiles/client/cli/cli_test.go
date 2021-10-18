package cli_test

import (
	"encoding/hex"
	"io/ioutil"
	"testing"
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/desmos-labs/desmos/v2/app"
	"github.com/desmos-labs/desmos/v2/x/profiles/client/utils"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/go-bip39"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v2/testutil"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

const (
	srcKeyName = "src"
	dstKeyName = "dst"
)

type Keys struct {
	Src cryptotypes.PrivKey
	Dst cryptotypes.PrivKey
}

type IntegrationTestSuite struct {
	suite.Suite

	cfg         network.Config
	network     *network.Network
	keyBase     keyring.Keyring
	keys        Keys
	testProfile *types.Profile
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := testutil.DefaultConfig()
	genesisState := cfg.GenesisState
	cfg.NumValidators = 2

	var authData authtypes.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[authtypes.ModuleName], &authData))

	// Generate test keys
	s.setupKeyBase()

	// Get keys
	srcKey, err := s.keyBase.Key(srcKeyName)
	s.Require().NoError(err)

	destKey, err := s.keyBase.Key(dstKeyName)
	s.Require().NoError(err)

	// Store a profile account inside the auth genesis data
	addr, err := sdk.AccAddressFromBech32("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs")
	s.Require().NoError(err)

	account, err := types.NewProfile(
		"dtag",
		"nickname",
		"bio",
		types.Pictures{},
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		authtypes.NewBaseAccountWithAddress(addr),
	)
	s.Require().NoError(err)

	accountAny, err := codectypes.NewAnyWithValue(account)
	s.Require().NoError(err)

	authData.Accounts = append(authData.Accounts, accountAny)

	// Set accounts of keys inside the auth genesis data
	srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcKey.GetAddress())
	srcAccountAny, err := codectypes.NewAnyWithValue(srcBaseAcc)
	s.Require().NoError(err)
	authData.Accounts = append(authData.Accounts, srcAccountAny)
	destBaseAcc := authtypes.NewBaseAccountWithAddress(destKey.GetAddress())
	s.Require().NoError(err)
	destAccountAny, err := codectypes.NewAnyWithValue(destBaseAcc)
	s.Require().NoError(err)
	authData.Accounts = append(authData.Accounts, destAccountAny)

	authDataBz, err := cfg.Codec.MarshalJSON(&authData)
	s.Require().NoError(err)

	genesisState[authtypes.ModuleName] = authDataBz

	// Store the profiles genesis state
	var profilesData types.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &profilesData))

	profilesData.DTagTransferRequests = []types.DTagTransferRequest{
		types.NewDTagTransferRequest(
			"dtag",
			"cosmos122u6u9gpdr2rp552fkkvlgyecjlmtqhkascl5a",
			"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
		),
	}
	profilesData.Blocks = []types.UserBlock{
		types.NewUserBlock(
			addr.String(),
			"cosmos1zs70glquczqgt83g03jnvcqppu4jjj8yjxwlvh",
			"Test block",
			"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		),
	}
	profilesData.Relationships = []types.Relationship{
		types.NewRelationship(
			addr.String(),
			"cosmos1zs70glquczqgt83g03jnvcqppu4jjj8yjxwlvh",
			"60303ae22b998861bce3b28f33eec1be758a213c86c93c076dbe9f558c11c752",
		),
	}
	profilesData.ApplicationLinks = []types.ApplicationLink{
		types.NewApplicationLink(
			"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
			types.NewData("reddit", "reddit-user"),
			types.ApplicationLinkStateInitialized,
			types.NewOracleRequest(
				0,
				1,
				types.NewOracleRequestCallData("twitter", "call_data"),
				"client_id",
			),
			nil,
			time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
		),
	}

	profilesData.Params = types.DefaultParams()

	pubKey := testutil.PubKeyFromBech32(
		"cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d",
	)
	s.Require().NoError(err)

	stringAddr, err := sdk.Bech32ifyAddressBytes("cosmos", pubKey.Address())
	s.Require().NoError(err)

	profilesData.ChainLinks = []types.ChainLink{
		types.NewChainLink(
			"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
			types.NewBech32Address(stringAddr, "cosmos"),
			types.NewProof(
				pubKey,
				"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
				"74657874",
			),
			types.NewChainConfig("cosmos"),
			time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
		),
	}

	profilesDataBz, err := cfg.Codec.MarshalJSON(&profilesData)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = profilesDataBz
	cfg.GenesisState = genesisState

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) setupKeyBase() {
	keyBase := keyring.NewInMemory()
	algo := hd.Secp256k1
	hdPath := sdk.GetConfig().GetFullFundraiserPath()

	srcEntropySeed, _ := bip39.NewEntropy(256)
	destEntropySeed, _ := bip39.NewEntropy(256)

	srcMnemonic, _ := bip39.NewMnemonic(srcEntropySeed)
	destMnemonic, _ := bip39.NewMnemonic(destEntropySeed)

	_, err := keyBase.NewAccount(srcKeyName, srcMnemonic, "", hdPath, algo)
	s.Require().NoError(err)

	_, err = keyBase.NewAccount(dstKeyName, destMnemonic, "", hdPath, algo)
	s.Require().NoError(err)

	s.keyBase = keyBase

	derivedSrcPriv, err := algo.Derive()(srcMnemonic, "", hdPath)
	s.Require().NoError(err)

	derivedDstPriv, err := algo.Derive()(destMnemonic, "", hdPath)
	s.Require().NoError(err)

	s.keys = Keys{
		Src: algo.Generate()(derivedSrcPriv),
		Dst: algo.Generate()(derivedDstPriv),
	}
}

func (s *IntegrationTestSuite) writeChainLinkJSONFile(filePath string) {
	srcKey := s.keys.Src

	addStr, err := sdk.Bech32ifyAddressBytes("cosmos", srcKey.PubKey().Address())
	s.Require().NoError(err)

	plainText := addStr
	sigBz, err := srcKey.Sign([]byte(plainText))
	s.Require().NoError(err)

	jsonData := utils.NewChainLinkJSON(
		types.NewBech32Address(addStr, "cosmos"),
		types.NewProof(srcKey.PubKey(), hex.EncodeToString(sigBz), hex.EncodeToString([]byte(plainText))),
		types.NewChainConfig("cosmos"),
	)

	params := app.MakeTestEncodingConfig()
	jsonBz := params.Marshaler.MustMarshalJSON(&jsonData)

	// Write the JSON to a temp file
	s.Require().NoError(ioutil.WriteFile(filePath, jsonBz, 0666))
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}
