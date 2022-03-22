package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/simapp"
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
	"github.com/desmos-labs/desmos/v3/x/supply/keeper"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type KeeperTestSuite struct {
	suite.Suite

	app            *simapp.SimApp
	ctx            sdk.Context
	cdc            codec.Codec
	legacyAminoCdc *codec.LegacyAmino

	k     keeper.Keeper
	ak    authkeeper.AccountKeeper
	bk    bankkeeper.Keeper
	dk    distributionkeeper.Keeper
	sk    stakingkeeper.Keeper
	denom string
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = simapp.Setup(false)

	maccPerms := simapp.GetMaccPerms()
	encodingConfig := simapp.MakeTestEncodingConfig()

	suite.cdc = encodingConfig.Marshaler
	suite.legacyAminoCdc = encodingConfig.Amino
	suite.ctx = suite.app.NewContext(false, tmproto.Header{})

	suite.denom = "udsm"

	maccPerms[authtypes.Burner] = []string{authtypes.Burner}
	maccPerms[authtypes.Minter] = []string{authtypes.Minter}
	maccPerms[banktypes.ModuleName] = []string{authtypes.Burner, authtypes.Minter, authtypes.Staking}

	suite.app.AccountKeeper = authkeeper.NewAccountKeeper(
		suite.cdc,
		suite.app.GetKey(authtypes.StoreKey),
		suite.app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
	)
	suite.app.AccountKeeper.SetParams(suite.ctx, authtypes.DefaultParams())

	suite.app.BankKeeper = bankkeeper.NewBaseKeeper(
		suite.cdc,
		suite.app.GetKey(banktypes.StoreKey),
		suite.app.AccountKeeper,
		suite.app.GetSubspace(banktypes.ModuleName),
		nil,
	)
	suite.app.BankKeeper.SetParams(suite.ctx, banktypes.DefaultParams())

	moduleAcc := authtypes.NewEmptyModuleAccount(banktypes.ModuleName, authtypes.Burner,
		authtypes.Minter, authtypes.Staking)

	suite.app.AccountKeeper.SetModuleAccount(suite.ctx, moduleAcc)

	suite.app.StakingKeeper = stakingkeeper.NewKeeper(
		suite.cdc,
		suite.app.GetKey(stakingtypes.StoreKey),
		suite.app.AccountKeeper,
		suite.app.BankKeeper,
		suite.app.GetSubspace(stakingtypes.ModuleName),
	)

	suite.app.DistrKeeper = distributionkeeper.NewKeeper(
		suite.cdc,
		suite.app.GetKey(distributiontypes.StoreKey),
		suite.app.GetSubspace(distributiontypes.ModuleName),
		suite.app.AccountKeeper,
		suite.app.BankKeeper,
		suite.app.StakingKeeper,
		"",
		nil,
	)

	// Define keeper
	suite.k = keeper.NewKeeper(suite.cdc, suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.DistrKeeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

// SupplySetup set up the total token supply with the given totalSupply. Further, it sends vestedSupply funds to a vested
// account and communityPoolSupply to the community pool.
//If totalSupply < vestedSupply + communityPoolSupply the function returns error.
func (suite *KeeperTestSuite) SupplySetup(totalSupply int64, vestedSupply int64, communityPoolSupply int64) {
	moduleAcc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, banktypes.ModuleName)
	totSupply := sdk.NewCoins(sdk.NewCoin(suite.denom, sdk.NewInt(totalSupply)))

	// Mint supply coins
	suite.Require().NoError(suite.app.BankKeeper.MintCoins(suite.ctx, moduleAcc.GetName(), totSupply))

	// Create a vesting account
	vestingAccount := vestingtypes.NewContinuousVestingAccount(
		suite.createBaseAccount(),
		sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(vestedSupply))),
		0,
		12324125423,
	)
	suite.app.AccountKeeper.SetAccount(suite.ctx, vestingAccount)

	// Send supply coins to the vesting account
	suite.Require().NoError(suite.app.BankKeeper.SendCoinsFromModuleToAccount(
		suite.ctx,
		banktypes.ModuleName,
		vestingAccount.GetAddress(),
		sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(200_000))),
	))

	// Fund community pool
	suite.Require().NoError(suite.app.DistrKeeper.FundCommunityPool(
		suite.ctx,
		sdk.NewCoins(sdk.NewCoin("udsm", sdk.NewInt(communityPoolSupply))),
		moduleAcc.GetAddress(),
	))
}

// createBaseAccount initialize a random BaseAccount
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

	// Create the base profile and set inside the auth keeper.
	// This is done in order to make sure that when we try to create a profile using the above address, the profile
	// can be created properly. Not storing the base profile would end up in the following error since it's null:
	// "the given profile cannot be serialized using Protobuf"
	return authtypes.NewBaseAccount(sdk.AccAddress(privKey.PubKey().Address()), privKey.PubKey(), 0, 0)
}
