package cli_test

import (
	"encoding/hex"
	"io/ioutil"
	"testing"
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/desmos-labs/desmos/v3/app"
	"github.com/desmos-labs/desmos/v3/x/profiles/client/utils"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/go-bip39"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v3/testutil"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"
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

	cfg                  network.Config
	network              *network.Network
	keyBase              keyring.Keyring
	keys                 Keys
	testChainLinkAccount testutil.ChainLinkAccount
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

	profilesData.Params = types.DefaultParams()
	profilesData.DTagTransferRequests = []types.DTagTransferRequest{
		types.NewDTagTransferRequest(
			"dtag",
			"cosmos122u6u9gpdr2rp552fkkvlgyecjlmtqhkascl5a",
			"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
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

	s.testChainLinkAccount = testutil.GetChainLinkAccount("cosmos", "cosmos")
	profilesData.ChainLinks = []types.ChainLink{
		s.testChainLinkAccount.GetBech32ChainLink(
			"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
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
		types.NewProof(srcKey.PubKey(), testutil.SingleSignatureProtoFromHex(hex.EncodeToString(sigBz)), hex.EncodeToString([]byte(plainText))),
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
