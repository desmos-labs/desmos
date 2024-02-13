package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_AfterProfileSaved() {
	suite.hooks.EXPECT().AfterProfileSaved(sdk.Context{}, nil)
	suite.k.SetHooks(suite.hooks)

	suite.k.AfterProfileSaved(sdk.Context{}, nil)
}

func (suite *KeeperTestSuite) TestKeeper_AfterProfileDeleted() {
	suite.hooks.EXPECT().AfterProfileDeleted(sdk.Context{}, nil)
	suite.k.SetHooks(suite.hooks)

	suite.k.AfterProfileDeleted(sdk.Context{}, nil)
}

func (suite *KeeperTestSuite) TestKeeper_AfterDTagTransferRequestCreated() {
	suite.hooks.EXPECT().AfterDTagTransferRequestCreated(sdk.Context{}, types.DTagTransferRequest{})
	suite.k.SetHooks(suite.hooks)

	suite.k.AfterDTagTransferRequestCreated(sdk.Context{}, types.DTagTransferRequest{})
}

func (suite *KeeperTestSuite) TestKeeper_AfterDTagTransferRequestAccepted() {
	suite.hooks.EXPECT().AfterDTagTransferRequestAccepted(sdk.Context{}, types.DTagTransferRequest{}, "")
	suite.k.SetHooks(suite.hooks)

	suite.k.AfterDTagTransferRequestAccepted(sdk.Context{}, types.DTagTransferRequest{}, "")
}

func (suite *KeeperTestSuite) TestKeeper_AfterDTagTransferRequestDeleted() {
	suite.hooks.EXPECT().AfterDTagTransferRequestDeleted(sdk.Context{}, "", "")
	suite.k.SetHooks(suite.hooks)

	suite.k.AfterDTagTransferRequestDeleted(sdk.Context{}, "", "")
}

func (suite *KeeperTestSuite) TestKeeper_AfterChainLinkSaved() {
	suite.hooks.EXPECT().AfterChainLinkSaved(sdk.Context{}, types.ChainLink{})
	suite.k.SetHooks(suite.hooks)

	suite.k.AfterChainLinkSaved(sdk.Context{}, types.ChainLink{})
}

func (suite *KeeperTestSuite) TestKeeper_AfterChainLinkDeleted() {
	suite.hooks.EXPECT().AfterChainLinkDeleted(sdk.Context{}, types.ChainLink{})
	suite.k.SetHooks(suite.hooks)

	suite.k.AfterChainLinkDeleted(sdk.Context{}, types.ChainLink{})
}

func (suite *KeeperTestSuite) TestKeeper_AfterApplicationLinkSaved() {
	suite.hooks.EXPECT().AfterApplicationLinkSaved(sdk.Context{}, types.ApplicationLink{})
	suite.k.SetHooks(suite.hooks)

	suite.k.AfterApplicationLinkSaved(sdk.Context{}, types.ApplicationLink{})
}

func (suite *KeeperTestSuite) TestKeeper_AfterApplicationLinkDeleted() {
	suite.hooks.EXPECT().AfterApplicationLinkDeleted(sdk.Context{}, types.ApplicationLink{})
	suite.k.SetHooks(suite.hooks)

	suite.k.AfterApplicationLinkDeleted(sdk.Context{}, types.ApplicationLink{})
}
