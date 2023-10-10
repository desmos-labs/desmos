package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

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

func TestMultiProfilesHooks_AfterProfileSaved(t *testing.T) {
	hook := MockHooks{CalledMap: make(map[string]bool)}

	hooks := types.NewMultiProfilesHooks(hook)
	hooks.AfterProfileSaved(sdk.Context{}, nil)

	require.True(t, hook.CalledMap["AfterProfileSaved"])
}

func TestMultiProfilesHooks_AfterProfileDeleted(t *testing.T) {
	hook := MockHooks{CalledMap: make(map[string]bool)}

	hooks := types.NewMultiProfilesHooks(hook)
	hooks.AfterProfileDeleted(sdk.Context{}, nil)

	require.True(t, hook.CalledMap["AfterProfileDeleted"])
}

func TestMultiProfilesHooks_AfterDTagTransferRequestCreated(t *testing.T) {
	hook := MockHooks{CalledMap: make(map[string]bool)}

	hooks := types.NewMultiProfilesHooks(hook)
	hooks.AfterDTagTransferRequestCreated(sdk.Context{}, types.DTagTransferRequest{})

	require.True(t, hook.CalledMap["AfterDTagTransferRequestCreated"])
}

func TestMultiProfilesHooks_AfterDTagTransferRequestAccepted(t *testing.T) {
	hook := MockHooks{CalledMap: make(map[string]bool)}

	hooks := types.NewMultiProfilesHooks(hook)
	hooks.AfterDTagTransferRequestAccepted(sdk.Context{}, types.DTagTransferRequest{}, "")

	require.True(t, hook.CalledMap["AfterDTagTransferRequestAccepted"])
}

func TestMultiProfilesHooks_AfterDTagTransferRequestDeleted(t *testing.T) {
	hook := MockHooks{CalledMap: make(map[string]bool)}

	hooks := types.NewMultiProfilesHooks(hook)
	hooks.AfterDTagTransferRequestDeleted(sdk.Context{}, "", "")

	require.True(t, hook.CalledMap["AfterDTagTransferRequestDeleted"])
}

func TestMultiProfilesHooks_AfterChainLinkSaved(t *testing.T) {
	hook := MockHooks{CalledMap: make(map[string]bool)}

	hooks := types.NewMultiProfilesHooks(hook)
	hooks.AfterChainLinkSaved(sdk.Context{}, types.ChainLink{})

	require.True(t, hook.CalledMap["AfterChainLinkSaved"])
}

func TestMultiProfilesHooks_AfterChainLinkDeleted(t *testing.T) {
	hook := MockHooks{CalledMap: make(map[string]bool)}

	hooks := types.NewMultiProfilesHooks(hook)
	hooks.AfterChainLinkDeleted(sdk.Context{}, types.ChainLink{})

	require.True(t, hook.CalledMap["AfterChainLinkDeleted"])
}

func TestMultiProfilesHooks_AfterApplicationLinkSaved(t *testing.T) {
	hook := MockHooks{CalledMap: make(map[string]bool)}

	hooks := types.NewMultiProfilesHooks(hook)
	hooks.AfterApplicationLinkSaved(sdk.Context{}, types.ApplicationLink{})

	require.True(t, hook.CalledMap["AfterApplicationLinkSaved"])
}

func TestMultiProfilesHooks_AfterApplicationLinkDeleted(t *testing.T) {
	hook := MockHooks{CalledMap: make(map[string]bool)}

	hooks := types.NewMultiProfilesHooks(hook)
	hooks.AfterApplicationLinkDeleted(sdk.Context{}, types.ApplicationLink{})

	require.True(t, hook.CalledMap["AfterApplicationLinkDeleted"])
}
