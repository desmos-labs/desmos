package keeper_test

import (
	"testing"
	"time"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	profileskeeper "github.com/desmos-labs/desmos/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/x/profiles/types"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	subspaceskeeper "github.com/desmos-labs/desmos/x/staging/subspaces/keeper"
	subspacetypes "github.com/desmos-labs/desmos/x/staging/subspaces/types"

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

	"github.com/desmos-labs/desmos/x/staging/posts/keeper"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
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
	rk             profileskeeper.Keeper
	sk             subspaceskeeper.Keeper

	stakingKeeper stakingkeeper.Keeper
	IBCKeeper     *ibckeeper.Keeper

	testData TestData
}

type TestData struct {
	postID                 string
	postOwner              string
	postCreationDate       time.Time
	postEndPollDate        time.Time
	postEndPollDateExpired time.Time
	answers                types.PollAnswers
	registeredReaction     types.RegisteredReaction
	post                   types.Post
	subspace               subspacetypes.Subspace
	otherSubspace          subspacetypes.Subspace
}

func (suite *KeeperTestSuite) SetupTest() {
	// Define the store keys
	keys := sdk.NewMemoryStoreKeys(types.StoreKey, paramstypes.StoreKey, profilestypes.StoreKey, subspacetypes.StoreKey,
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

	pk := paramskeeper.NewKeeper(
		suite.cdc,
		suite.legacyAminoCdc,
		keys[paramstypes.StoreKey],
		tKeys[paramstypes.TStoreKey],
	)

	ak := authkeeper.NewAccountKeeper(
		suite.cdc,
		keys[authtypes.StoreKey],
		pk.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
	)

	capabilityKeeper := capabilitykeeper.NewKeeper(suite.cdc, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	ScopedProfilesKeeper := capabilityKeeper.ScopeToModule(types.ModuleName)

	scopedIBCKeeper := capabilityKeeper.ScopeToModule(ibchost.ModuleName)
	IBCKeeper := ibckeeper.NewKeeper(
		suite.cdc,
		keys[ibchost.StoreKey],
		pk.Subspace(ibchost.ModuleName),
		suite.stakingKeeper,
		scopedIBCKeeper,
	)

	suite.rk = profileskeeper.NewKeeper(
		suite.cdc,
		suite.storeKey,
		pk.Subspace(profilestypes.DefaultParamsSpace),
		ak,
		IBCKeeper.ChannelKeeper,
		&IBCKeeper.PortKeeper,
		ScopedProfilesKeeper,
	)

	suite.sk = subspaceskeeper.NewKeeper(
		suite.storeKey,
		suite.cdc,
	)

	suite.k = keeper.NewKeeper(
		suite.cdc,
		keys[types.StoreKey],
		pk.Subspace(types.DefaultParamSpace),
		suite.rk,
		suite.sk,
	)

	suite.testData.postID = "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"
	suite.testData.postOwner = "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"

	// Setup data

	suite.testData.subspace = subspacetypes.NewSubspace(
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"test",
		suite.testData.postOwner,
		suite.testData.postOwner,
		subspacetypes.SubspaceTypeOpen,
		blockTime,
	)

	suite.testData.otherSubspace = subspacetypes.NewSubspace(
		"5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"test",
		suite.testData.postOwner,
		suite.testData.postOwner,
		subspacetypes.SubspaceTypeOpen,
		blockTime,
	)

	suite.testData.postCreationDate = blockTime
	suite.testData.postEndPollDate, _ = time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	suite.testData.postEndPollDateExpired, _ = time.Parse(time.RFC3339, "2019-01-01T01:15:00.000Z")
	suite.testData.answers = types.PollAnswers{
		types.NewPollAnswer("1", "Yes"),
		types.NewPollAnswer("2", "No"),
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
		PollData: &types.PollData{
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

}
