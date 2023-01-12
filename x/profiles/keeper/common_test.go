package keeper_test

import (
	"path/filepath"
	"testing"
	"time"

	relationshipskeeper "github.com/desmos-labs/desmos/v4/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v4/x/relationships/types"

	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"

	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/go-bip39"

	"github.com/cosmos/cosmos-sdk/crypto/hd"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/testutil/ibctesting"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibchost "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"

	"github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc              codec.Codec
	legacyAminoCdc   *codec.LegacyAmino
	ctx              sdk.Context
	storeKey         sdk.StoreKey
	k                keeper.Keeper
	ak               authkeeper.AccountKeeper
	rk               relationshipskeeper.Keeper
	sk               subspaceskeeper.Keeper
	paramsKeeper     paramskeeper.Keeper
	bankKeeper       bankkeeper.Keeper
	stakingKeeper    stakingkeeper.Keeper
	upgradeKeeper    upgradekeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper
	capabilityKeeper *capabilitykeeper.Keeper

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
	keys := sdk.NewKVStoreKeys(
		types.StoreKey, relationshipstypes.StoreKey, subspacestypes.StoreKey,
		authtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, capabilitytypes.StoreKey,
	)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

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
	for _, memKey := range memKeys {
		ms.MountStoreWithDB(memKey, sdk.StoreTypeMemory, nil)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

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

	suite.capabilityKeeper = capabilitykeeper.NewKeeper(suite.cdc, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	homeDir := filepath.Join(suite.T().TempDir(), "x_upgrade_keeper_test")
	suite.upgradeKeeper = upgradekeeper.NewKeeper(
		nil,
		keys[upgradetypes.StoreKey],
		suite.cdc,
		homeDir,
		nil,
	)

	suite.bankKeeper = bankkeeper.NewBaseKeeper(
		suite.cdc,
		keys[banktypes.StoreKey],
		suite.ak,
		suite.paramsKeeper.Subspace(banktypes.ModuleName),
		nil,
	)

	suite.stakingKeeper = stakingkeeper.NewKeeper(
		suite.cdc,
		keys[stakingtypes.StoreKey],
		suite.ak,
		suite.bankKeeper,
		suite.paramsKeeper.Subspace(stakingtypes.ModuleName),
	)

	scopedIBCKeeper := suite.capabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedProfilesKeeper := suite.capabilityKeeper.ScopeToModule(types.ModuleName)
	suite.IBCKeeper = ibckeeper.NewKeeper(
		suite.cdc,
		keys[ibchost.StoreKey],
		suite.paramsKeeper.Subspace(ibchost.ModuleName),
		suite.stakingKeeper,
		suite.upgradeKeeper,
		scopedIBCKeeper,
	)

	suite.sk = subspaceskeeper.NewKeeper(suite.cdc, keys[subspacestypes.StoreKey], nil, nil)
	suite.rk = relationshipskeeper.NewKeeper(suite.cdc, keys[relationshipstypes.StoreKey], suite.sk)
	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.legacyAminoCdc,
		suite.storeKey,
		suite.paramsKeeper.Subspace(types.DefaultParamsSpace),
		suite.ak,
		suite.rk,
		suite.IBCKeeper.ChannelKeeper,
		&suite.IBCKeeper.PortKeeper,
		scopedProfilesKeeper,
	)

	// Set the IBC data
	suite.SetupIBCTest()
}

func (suite *KeeperTestSuite) SetupIBCTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
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
