package keeper_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec/address"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/go-bip39"

	"github.com/cosmos/cosmos-sdk/crypto/hd"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	ibctesting "github.com/cosmos/ibc-go/v8/testing"
	"github.com/desmos-labs/desmos/v6/app"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	db "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	desmosibctesting "github.com/desmos-labs/desmos/v6/testutil/ibctesting"
	"github.com/desmos-labs/desmos/v6/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v6/x/profiles/testutil"
	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	storeKey       storetypes.StoreKey
	k              *keeper.Keeper
	ak             authkeeper.AccountKeeper

	ctrl          *gomock.Controller
	rk            *testutil.MockRelationshipsKeeper
	channelKeeper *testutil.MockChannelKeeper
	portKeeper    *testutil.MockPortKeeper
	scopedKeeper  *testutil.MockScopedKeeper

	// Used for IBC testing
	coordinator *ibctesting.Coordinator
	chainA      *ibctesting.TestChain
	chainB      *ibctesting.TestChain
}

// TestProfile represents a test profile
type TestProfile struct {
	*types.Profile

	privKey cryptotypes.PrivKey
}

// Sign allows to sign the given data using the profile private key
func (p TestProfile) Sign(data []byte) []byte {
	bz, err := p.privKey.Sign(data)
	if err != nil {
		panic(err)
	}
	return bz
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define the store keys
	keys := storetypes.NewKVStoreKeys(types.StoreKey, authtypes.StoreKey)

	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// Dependencies initializations
	suite.ak = authkeeper.NewAccountKeeper(
		suite.cdc,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
		address.NewBech32Codec("cosmos"),
		"cosmos",
		authtypes.NewModuleAddress("gov").String(),
	)

	// Mocks initializations
	suite.ctrl = gomock.NewController(suite.T())

	suite.rk = testutil.NewMockRelationshipsKeeper(suite.ctrl)
	suite.channelKeeper = testutil.NewMockChannelKeeper(suite.ctrl)
	suite.portKeeper = testutil.NewMockPortKeeper(suite.ctrl)
	suite.scopedKeeper = testutil.NewMockScopedKeeper(suite.ctrl)

	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.legacyAminoCdc,
		suite.storeKey,
		suite.ak,
		suite.rk,
		suite.channelKeeper,
		suite.portKeeper,
		suite.scopedKeeper,
		authtypes.NewModuleAddress("gov").String(),
	)

	// Set the IBC data
	suite.SetupIBCTest()
}

func (suite *KeeperTestSuite) SetupIBCTest() {
	suite.coordinator = desmosibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(2))
}

func (suite *KeeperTestSuite) GetRandomProfile() TestProfile {
	// Read entropy seed straight from tmcrypto.Rand and convert to mnemonic
	entropySeed, err := bip39.NewEntropy(256)
	suite.Require().NoError(err)

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	suite.Require().NoError(err)

	// Generate a private key
	derivedPrivKey, err := hd.Secp256k1.Derive()(mnemonic, "", sdk.FullFundraiserPath)
	suite.Require().NoError(err)

	privKey := hd.Secp256k1.Generate()(derivedPrivKey)

	// Create the base profile and set inside the auth keeper.
	// This is done in order to make sure that when we try to create a profile using the above address, the profile
	// can be created properly. Not storing the base profile would end up in the following error since it's null:
	// "the given profile cannot be serialized using Protobuf"
	baseAcc := authtypes.NewBaseAccount(sdk.AccAddress(privKey.PubKey().Address()), privKey.PubKey(), 0, 0)

	profile, err := types.NewProfile(
		baseAcc.Address,
		"Random user",
		"",
		types.NewPictures("", ""),
		time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
		baseAcc,
	)
	suite.Require().NoError(err)

	return TestProfile{
		Profile: profile,
		privKey: privKey,
	}
}

func (suite *KeeperTestSuite) CheckProfileNoError(profile *types.Profile, err error) *types.Profile {
	suite.Require().NoError(err)
	return profile
}

func (suite *KeeperTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func NewIBCProfilesPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = types.IBCPortID
	path.EndpointB.ChannelConfig.PortID = types.IBCPortID
	path.EndpointA.ChannelConfig.Version = "ics-20"
	path.EndpointB.ChannelConfig.Version = "ics-20"

	return path
}
