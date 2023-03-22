package keeper_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/go-bip39"

	"github.com/cosmos/cosmos-sdk/crypto/hd"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	"github.com/desmos-labs/desmos/v4/app"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v4/x/profiles/testutil"
	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	storeKey       sdk.StoreKey
	k              keeper.Keeper
	ak             authkeeper.AccountKeeper
	paramsKeeper   paramskeeper.Keeper

	rk            *testutil.MockRelationshipsKeeper
	channelKeeper *testutil.MockChannelKeeper
	portKeeper    *testutil.MockPortKeeper
	scopedKeeper  *testutil.MockScopedKeeper

	// Used for IBC testing
	// coordinator *ibctesting.Coordinator
	// chainA      *ibctesting.TestChain
	// chainB      *ibctesting.TestChain
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
	keys := sdk.NewKVStoreKeys(types.StoreKey, authtypes.StoreKey, paramstypes.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)

	suite.storeKey = keys[types.StoreKey]

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, sdk.StoreTypeIAVL, memDB)
	}
	for _, tKey := range tKeys {
		ms.MountStoreWithDB(tKey, sdk.StoreTypeTransient, memDB)
	}
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	// Dependencies initializations
	suite.paramsKeeper = paramskeeper.NewKeeper(
		suite.cdc, suite.legacyAminoCdc, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey],
	)
	suite.ak = authkeeper.NewAccountKeeper(
		suite.cdc,
		keys[authtypes.StoreKey],
		suite.paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
	)

	// Mocks initializations
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	suite.rk = testutil.NewMockRelationshipsKeeper(ctrl)
	suite.channelKeeper = testutil.NewMockChannelKeeper(ctrl)
	suite.portKeeper = testutil.NewMockPortKeeper(ctrl)
	suite.scopedKeeper = testutil.NewMockScopedKeeper(ctrl)

	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.legacyAminoCdc,
		suite.storeKey,
		suite.paramsKeeper.Subspace(types.DefaultParamsSpace),
		suite.ak,
		suite.rk,
		suite.channelKeeper,
		suite.portKeeper,
		suite.scopedKeeper,
	)

	// Set the IBC data
	//suite.SetupIBCTest()
}

// func (suite *KeeperTestSuite) SetupIBCTest() {
// 	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
// 	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(1))
// 	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(2))
// }

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
