package keeper_test

import (
	"testing"

	"cosmossdk.io/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	db "github.com/cosmos/cosmos-db"

	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/go-bip39"
	"github.com/stretchr/testify/suite"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/supply/keeper"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	cdc         codec.Codec
	legacyAmino *codec.LegacyAmino

	k  keeper.Keeper
	ak authkeeper.AccountKeeper
	bk bankkeeper.Keeper
	dk distributionkeeper.Keeper
	sk *stakingkeeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	maccPerms := app.GetMaccPerms()

	// Define store keys
	keys := storetypes.NewKVStoreKeys(authtypes.StoreKey, banktypes.StoreKey, distributiontypes.StoreKey, stakingtypes.StoreKey)

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAmino = app.MakeCodecs()

	maccPerms[authtypes.Burner] = []string{authtypes.Burner}
	maccPerms[authtypes.Minter] = []string{authtypes.Minter}
	maccPerms[banktypes.ModuleName] = []string{authtypes.Burner, authtypes.Minter, authtypes.Staking}

	// Dependencies initialization
	suite.ak = authkeeper.NewAccountKeeper(
		suite.cdc,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		address.NewBech32Codec("cosmos"),
		"cosmos",
		authtypes.NewModuleAddress("gov").String(),
	)

	moduleAcc := authtypes.NewEmptyModuleAccount(banktypes.ModuleName, authtypes.Burner,
		authtypes.Minter, authtypes.Staking)

	suite.ak.SetModuleAccount(suite.ctx, moduleAcc)

	suite.bk = bankkeeper.NewBaseKeeper(
		suite.cdc,
		runtime.NewKVStoreService(keys[banktypes.StoreKey]),
		suite.ak,
		nil,
		authtypes.NewModuleAddress("gov").String(),
		log.NewNopLogger(),
	)

	suite.sk = stakingkeeper.NewKeeper(
		suite.cdc,
		runtime.NewKVStoreService(keys[stakingtypes.StoreKey]),
		suite.ak,
		suite.bk,
		authtypes.NewModuleAddress("gov").String(),
		address.NewBech32Codec("cosmosvaloper"),
		address.NewBech32Codec("cosmosvalcons"),
	)

	suite.dk = distributionkeeper.NewKeeper(
		suite.cdc,
		runtime.NewKVStoreService(keys[distributiontypes.StoreKey]),
		suite.ak,
		suite.bk,
		suite.sk,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress("gov").String(),
	)
	suite.dk.FeePool.Set(suite.ctx, distributiontypes.InitialFeePool())

	// Define keeper
	suite.k = keeper.NewKeeper(suite.cdc, suite.ak, suite.bk, suite.dk)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

// setupSupply setups the total token supply with the given totalSupply. Further, it sends vestedSupply funds to a vested
// account and communityPoolSupply to the community pool.
// If totalSupply < vestedSupply + communityPoolSupply the function returns error.
func (suite *KeeperTestSuite) setupSupply(ctx sdk.Context, totalSupply sdk.Coins, vestedSupply sdk.Coins, communityPoolSupply sdk.Coins) {
	moduleAcc := suite.ak.GetModuleAccount(ctx, banktypes.ModuleName)

	// Mint supply coins
	suite.Require().NoError(suite.bk.MintCoins(ctx, moduleAcc.GetName(), totalSupply))

	// Create a vesting account
	vestingAccount, err := vestingtypes.NewContinuousVestingAccount(
		suite.createBaseAccount(),
		vestedSupply,
		0,
		12324125423,
	)
	suite.Require().NoError(err)
	suite.ak.SetAccount(ctx, vestingAccount)

	// Send supply coins to the vesting account
	suite.Require().NoError(suite.bk.SendCoinsFromModuleToAccount(
		ctx,
		banktypes.ModuleName,
		vestingAccount.GetAddress(),
		vestedSupply,
	))

	// Fund community pool
	suite.Require().NoError(suite.dk.FundCommunityPool(
		ctx,
		communityPoolSupply,
		moduleAcc.GetAddress(),
	))
}

// createBaseAccount returns a mew random BaseAccount
func (suite *KeeperTestSuite) createBaseAccount() *authtypes.BaseAccount {
	// Read entropy seed straight from tmcrypto.Rand and convert to mnemonic
	entropySeed, err := bip39.NewEntropy(256)
	suite.Require().NoError(err)

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	suite.Require().NoError(err)

	// Generate a private key
	derivedPrivKey, err := hd.Secp256k1.Derive()(mnemonic, "", sdk.FullFundraiserPath)
	suite.Require().NoError(err)

	privKey := hd.Secp256k1.Generate()(derivedPrivKey)

	return authtypes.NewBaseAccount(sdk.AccAddress(privKey.PubKey().Address()), privKey.PubKey(), 0, 0)
}
