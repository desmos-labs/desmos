package keeper_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibchost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	ibckeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/posts/types"
	profileskeeper "github.com/desmos-labs/desmos/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/x/profiles/types"
	subspaceskeeper "github.com/desmos-labs/desmos/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/x/subspaces/types"
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

type KeeperTestSuite struct {
	suite.Suite

	cdc            codec.BinaryMarshaler
	legacyAminoCdc *codec.LegacyAmino
	ctx            sdk.Context
	k              keeper.Keeper
	storeKey       sdk.StoreKey
	ak             authkeeper.AccountKeeper
	rk             profileskeeper.Keeper
	sk             subspaceskeeper.Keeper

	stakingKeeper stakingkeeper.Keeper
	IBCKeeper     *ibckeeper.Keeper

	testData TestData
}

// TestProfile represents a test profile
type TestProfile struct {
	*profilestypes.Profile

	privKey cryptotypes.PrivKey
}

type TestData struct {
	postID                 string
	postOwner              string
	postCreationDate       time.Time
	postEndPollDate        time.Time
	postEndPollDateExpired time.Time
	answers                types.ProvidedAnswers
	registeredReaction     types.RegisteredReaction
	post                   types.Post
	profile                TestProfile
	subspace               subspacestypes.Subspace
	otherSubspace          subspacestypes.Subspace
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define the store keys
	keys := sdk.NewKVStoreKeys(types.StoreKey, authtypes.StoreKey, paramstypes.StoreKey, profilestypes.StoreKey, subspacestypes.StoreKey,
		ibchost.StoreKey, capabilitytypes.StoreKey)
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

	blockTime, _ := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	suite.ctx = sdk.NewContext(
		ms,
		tmproto.Header{ChainID: "test-chain-id", Time: blockTime},
		false,
		log.NewNopLogger(),
	)
	suite.cdc, suite.legacyAminoCdc = app.MakeCodecs()

	paramsKeeper := paramskeeper.NewKeeper(
		suite.cdc,
		suite.legacyAminoCdc,
		keys[paramstypes.StoreKey],
		tKeys[paramstypes.TStoreKey],
	)

	suite.ak = authkeeper.NewAccountKeeper(
		suite.cdc,
		keys[authtypes.StoreKey],
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
	)

	capabilityKeeper := capabilitykeeper.NewKeeper(
		suite.cdc,
		keys[capabilitytypes.StoreKey],
		memKeys[capabilitytypes.MemStoreKey],
	)

	ScopedProfilesKeeper := capabilityKeeper.ScopeToModule(types.ModuleName)
	scopedIBCKeeper := capabilityKeeper.ScopeToModule(ibchost.ModuleName)

	IBCKeeper := ibckeeper.NewKeeper(
		suite.cdc,
		keys[ibchost.StoreKey],
		paramsKeeper.Subspace(ibchost.ModuleName),
		suite.stakingKeeper,
		scopedIBCKeeper,
	)

	suite.sk = subspaceskeeper.NewKeeper(
		keys[subspacestypes.StoreKey],
		suite.cdc,
	)

	suite.rk = profileskeeper.NewKeeper(
		suite.cdc,
		keys[profilestypes.StoreKey],
		paramsKeeper.Subspace(profilestypes.DefaultParamsSpace),
		suite.ak,
		IBCKeeper.ChannelKeeper,
		&IBCKeeper.PortKeeper,
		ScopedProfilesKeeper,
		suite.sk,
	)

	suite.k = keeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		paramsKeeper.Subspace(types.DefaultParamSpace),
		suite.rk,
		suite.sk,
	)

	suite.testData.postID = "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"
	suite.testData.postOwner = "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"

	// Setup data

	suite.testData.subspace = subspacestypes.NewSubspace(
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"test",
		"description",
		"https://logo-png.com",
		suite.testData.postOwner,
		suite.testData.postOwner,
		subspacestypes.SubspaceTypeOpen,
		blockTime,
	)

	suite.testData.otherSubspace = subspacestypes.NewSubspace(
		"5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"test",
		"description",
		"https://logo-png.com",
		suite.testData.postOwner,
		suite.testData.postOwner,
		subspacestypes.SubspaceTypeOpen,
		blockTime,
	)

	suite.testData.postCreationDate = blockTime
	suite.testData.postEndPollDate, _ = time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	suite.testData.postEndPollDateExpired, _ = time.Parse(time.RFC3339, "2019-01-01T01:15:00.000Z")
	suite.testData.answers = types.ProvidedAnswers{
		types.NewProvidedAnswer("1", "Yes"),
		types.NewProvidedAnswer("2", "No"),
	}
	suite.testData.post = types.Post{
		PostID:               suite.testData.postID,
		Message:              "Post message",
		Created:              suite.testData.postCreationDate,
		LastEdited:           suite.testData.postCreationDate.Add(1),
		CommentsState:        types.CommentsStateBlocked,
		Subspace:             suite.testData.subspace.ID,
		AdditionalAttributes: nil,
		Creator:              suite.testData.postOwner,
		Attachments: types.NewAttachments(
			types.NewAttachment("https://uri.com", "text/plain", []string{suite.testData.postOwner}),
		),
		Poll: &types.Poll{
			Question:              "poll?",
			ProvidedAnswers:       types.NewPollAnswers(suite.testData.answers[0], suite.testData.answers[1]),
			EndDate:               suite.testData.postEndPollDate,
			AllowsMultipleAnswers: true,
			AllowsAnswerEdits:     true,
		},
	}

	suite.testData.registeredReaction = types.NewRegisteredReaction(
		suite.testData.postOwner,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	suite.initProfile()
}

func (suite *KeeperTestSuite) initProfile() {
	mnemonic := "ugly like hockey joy digital glow learn remove pet promote screen twenty phone beach aspect mechanic gate piano antenna island loyal possible acoustic jewel"
	derivedPrivKey, err := hd.Secp256k1.Derive()(mnemonic, "", sdk.FullFundraiserPath)
	suite.Require().NoError(err)

	privKey := hd.Secp256k1.Generate()(derivedPrivKey)

	// Create the base account and set inside the auth keeper.
	// This is done in order to make sure that when we try to create a profile using the above address, the profile
	// can be created properly. Not storing the base account would end up in the following error since it's null:
	// "the given account cannot be serialized using Protobuf"
	baseAcc := authtypes.NewBaseAccount(sdk.AccAddress(privKey.PubKey().Address()), privKey.PubKey(), 0, 0)
	suite.ak.SetAccount(suite.ctx, baseAcc)

	profile, err := profilestypes.NewProfile(
		"dtag",
		"test-user",
		"biography",
		profilestypes.NewPictures(
			"https://shorturl.at/adnX3",
			"https://shorturl.at/cgpyF",
		),
		time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
		baseAcc,
	)
	suite.Require().NoError(err)

	suite.testData.profile = TestProfile{
		Profile: profile,
		privKey: privKey,
	}

	err = suite.rk.StoreProfile(suite.ctx, profile)
	suite.Require().NoError(err)
}
