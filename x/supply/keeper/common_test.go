package keeper_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/go-bip39"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/desmos-labs/desmos/v2/app"
	"github.com/desmos-labs/desmos/v2/x/supply/keeper"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"
)

type KeeperTestSuite struct {
	suite.Suite

	storeKeys      map[string]*sdk.MemoryStoreKey
	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	k              keeper.Keeper
	ak             authkeeper.AccountKeeper
	bk             bankkeeper.Keeper
	dk             distributionkeeper.Keeper
	sk             stakingkeeper.Keeper
	paramsKeeper   paramskeeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define store keys
	suite.storeKeys = sdk.NewMemoryStoreKeys(authtypes.StoreKey, banktypes.StoreKey, distributiontypes.StoreKey, stakingtypes.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.StoreKey)

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range suite.storeKeys {
		ms.MountStoreWithDB(key, sdk.StoreTypeIAVL, memDB)
	}

	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	suite.paramsKeeper = paramskeeper.NewKeeper(
		suite.cdc, suite.legacyAminoCdc, suite.storeKeys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey],
	)

	maccPerms := app.GetMaccPerms()
	maccPerms["multiPerm"] = []string{authtypes.Burner, authtypes.Minter, authtypes.Staking}

	suite.ak = authkeeper.NewAccountKeeper(
		suite.cdc,
		suite.storeKeys[authtypes.StoreKey],
		suite.paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
	)

	suite.bk = bankkeeper.NewBaseKeeper(
		suite.cdc,
		suite.storeKeys[banktypes.StoreKey],
		suite.ak,
		suite.paramsKeeper.Subspace(banktypes.ModuleName),
		nil,
	)

	suite.sk = stakingkeeper.NewKeeper(
		suite.cdc,
		suite.storeKeys[stakingtypes.StoreKey],
		suite.ak,
		suite.bk,
		suite.paramsKeeper.Subspace(stakingtypes.ModuleName),
	)

	suite.dk = distributionkeeper.NewKeeper(
		suite.cdc,
		suite.storeKeys[distributiontypes.StoreKey],
		suite.paramsKeeper.Subspace(distributiontypes.ModuleName),
		suite.ak,
		suite.bk,
		suite.sk,
		"",
		nil,
	)

	// Define keeper
	suite.k = keeper.NewKeeper(suite.cdc, suite.ak, suite.bk, suite.dk)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) CreateBaseAccount() *authtypes.BaseAccount {
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
	return authtypes.NewBaseAccount(sdk.AccAddress(privKey.PubKey().Address()), privKey.PubKey(), 0, 0)
}

func (suite *KeeperTestSuite) SetSupply(coin sdk.Coin) {
	intBytes, err := coin.Amount.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal amount value %v", err))
	}

	str := suite.ctx.KVStore(suite.storeKeys[banktypes.StoreKey])
	supplyStore := prefix.NewStore(str, banktypes.SupplyKey)

	// Bank invariants and IBC requires to remove zero coins.
	if coin.IsZero() {
		supplyStore.Delete([]byte(coin.GetDenom()))
	} else {
		supplyStore.Set([]byte(coin.GetDenom()), intBytes)
	}
}
