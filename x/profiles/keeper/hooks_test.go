package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

var _ types.ProfilesHooks = &MockHooks{}

type MockHooks struct {
	CalledMap map[string]bool
}

func (h MockHooks) AfterProfileSaved(_ sdk.Context, _ *types.Profile) {
	h.CalledMap["AfterProfileSaved"] = true
}

func (h MockHooks) AfterProfileDeleted(_ sdk.Context, _ *types.Profile) {
	h.CalledMap["AfterProfileDeleted"] = true
}

func (h MockHooks) AfterDTagTransferRequestCreated(_ sdk.Context, _ types.DTagTransferRequest) {
	h.CalledMap["AfterDTagTransferRequestCreated"] = true
}

func (h MockHooks) AfterDTagTransferRequestAccepted(_ sdk.Context, _ types.DTagTransferRequest, _ string) {
	h.CalledMap["AfterDTagTransferRequestAccepted"] = true
}

func (h MockHooks) AfterDTagTransferRequestDeleted(_ sdk.Context, _, _ string) {
	h.CalledMap["AfterDTagTransferRequestDeleted"] = true
}

func (h MockHooks) AfterChainLinkSaved(_ sdk.Context, _ types.ChainLink) {
	h.CalledMap["AfterChainLinkSaved"] = true
}

func (h MockHooks) AfterChainLinkDeleted(_ sdk.Context, _ types.ChainLink) {
	h.CalledMap["AfterChainLinkDeleted"] = true
}

func (h MockHooks) AfterApplicationLinkSaved(_ sdk.Context, _ types.ApplicationLink) {
	h.CalledMap["AfterApplicationLinkSaved"] = true
}

func (h MockHooks) AfterApplicationLinkDeleted(_ sdk.Context, _ types.ApplicationLink) {
	h.CalledMap["AfterApplicationLinkDeleted"] = true
}

func (suite *KeeperTestSuite) TestKeeper_AfterProfileSaved() {
	hook := MockHooks{CalledMap: make(map[string]bool)}
	suite.k.SetHooks(hook).AfterProfileSaved(sdk.Context{}, nil)

	suite.Require().True(hook.CalledMap["AfterProfileSaved"])
}

func (suite *KeeperTestSuite) TestKeeper_AfterProfileDeleted() {
	hook := MockHooks{CalledMap: make(map[string]bool)}
	suite.k.SetHooks(hook).AfterProfileDeleted(sdk.Context{}, nil)

	suite.Require().True(hook.CalledMap["AfterProfileDeleted"])
}

func (suite *KeeperTestSuite) TestKeeper_AfterDTagTransferRequestCreated() {
	hook := MockHooks{CalledMap: make(map[string]bool)}
	suite.k.SetHooks(hook).AfterDTagTransferRequestCreated(sdk.Context{}, types.DTagTransferRequest{})

	suite.Require().True(hook.CalledMap["AfterDTagTransferRequestCreated"])
}

func (suite *KeeperTestSuite) TestKeeper_AfterDTagTransferRequestAccepted() {
	hook := MockHooks{CalledMap: make(map[string]bool)}
	suite.k.SetHooks(hook).AfterDTagTransferRequestAccepted(sdk.Context{}, types.DTagTransferRequest{}, "")

	suite.Require().True(hook.CalledMap["AfterDTagTransferRequestAccepted"])
}

func (suite *KeeperTestSuite) TestKeeper_AfterDTagTransferRequestDeleted() {
	hook := MockHooks{CalledMap: make(map[string]bool)}
	suite.k.SetHooks(hook).AfterDTagTransferRequestDeleted(sdk.Context{}, "", "")

	suite.Require().True(hook.CalledMap["AfterDTagTransferRequestDeleted"])
}

func (suite *KeeperTestSuite) TestKeeper_AfterChainLinkSaved() {
	hook := MockHooks{CalledMap: make(map[string]bool)}
	suite.k.SetHooks(hook).AfterChainLinkSaved(sdk.Context{}, types.ChainLink{})
	suite.Require().True(hook.CalledMap["AfterChainLinkSaved"])
}

func (suite *KeeperTestSuite) TestKeeper_AfterChainLinkDeleted() {
	hook := MockHooks{CalledMap: make(map[string]bool)}
	suite.k.SetHooks(hook).AfterChainLinkDeleted(sdk.Context{}, types.ChainLink{})

	suite.Require().True(hook.CalledMap["AfterChainLinkDeleted"])
}

func (suite *KeeperTestSuite) TestKeeper_AfterApplicationLinkSaved() {
	hook := MockHooks{CalledMap: make(map[string]bool)}
	suite.k.SetHooks(hook).AfterApplicationLinkSaved(sdk.Context{}, types.ApplicationLink{})

	suite.Require().True(hook.CalledMap["AfterApplicationLinkSaved"])
}

func (suite *KeeperTestSuite) TestKeeper_AfterApplicationLinkDeleted() {
	hook := MockHooks{CalledMap: make(map[string]bool)}
	suite.k.SetHooks(hook).AfterApplicationLinkDeleted(sdk.Context{}, types.ApplicationLink{})

	suite.Require().True(hook.CalledMap["AfterApplicationLinkDeleted"])
}
