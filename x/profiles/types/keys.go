package types

import (
	"strings"
)

// DONTCOVER

const (
	ModuleName = "profiles"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionSaveProfile               = "save_profile"
	ActionDeleteProfile             = "delete_profile"
	ActionRequestDTag               = "request_dtag_transfer"
	ActionAcceptDTagTransfer        = "accept_dtag_transfer_request"
	ActionRefuseDTagTransferRequest = "refuse_dtag_transfer_request"
	ActionCancelDTagTransferRequest = "cancel_dtag_transfer_request"
	ActionLinkChainAccount          = "link_chain_account"
	ActionUnlinkChainAccount        = "unlink_chain_account"
	ActionLinkApplication           = "link_application"
	ActionUnlinkApplication         = "unlink_application"

	DoNotModify = "[do-not-modify]"

	// IBCPortID is the default port id that profiles module binds to.
	IBCPortID = "ibc-profiles"

	// Version defines the current version the IBC profiles module supports
	Version = "desmos-1"
)

var (
	Separator = []byte{0x00}

	// IBCPortKey defines the key to store the port ID in store
	IBCPortKey = []byte{0x01}

	DTagPrefix                    = []byte{0x10}
	DTagTransferRequestPrefix     = []byte{0x11}
	ChainLinksPrefix              = []byte{0x12}
	ApplicationLinkPrefix         = []byte{0x13}
	ApplicationLinkClientIDPrefix = []byte{0x14}
	ApplicationLinkAppPrefix      = []byte{0x15}
)

// DTagStoreKey turns a DTag into the key used to store the address associated with it into the store
func DTagStoreKey(dTag string) []byte {
	return append(DTagPrefix, []byte(strings.ToLower(dTag))...)
}

// IncomingDTagTransferRequestsPrefix returns the prefix used to store all the DTag transfer requests that
// have been made towards the given recipient
func IncomingDTagTransferRequestsPrefix(recipient string) []byte {
	return append(DTagTransferRequestPrefix, []byte(recipient)...)
}

// DTagTransferRequestStoreKey returns the store key used to save the DTag transfer request made
// from the sender towards the recipient
func DTagTransferRequestStoreKey(sender, recipient string) []byte {
	return append(IncomingDTagTransferRequestsPrefix(recipient), []byte(sender)...)
}

// UserChainLinksPrefix returns the store prefix used to identify all the chain links for the given user
func UserChainLinksPrefix(user string) []byte {
	return append(ChainLinksPrefix, []byte(user)...)
}

// UserChainLinksChainPrefix returns the store prefix used to identify all the chain links for the given user and chain
func UserChainLinksChainPrefix(user, chainName string) []byte {
	return append(UserChainLinksPrefix(user), []byte(chainName)...)
}

// ChainLinksStoreKey returns the store key used to store the chain links containing the given data
func ChainLinksStoreKey(user, chainName, address string) []byte {
	return append(UserChainLinksChainPrefix(user, chainName), []byte(address)...)
}

// UserApplicationLinksPrefix returns the store prefix used to identify all the application links for the given user
func UserApplicationLinksPrefix(user string) []byte {
	return append(ApplicationLinkPrefix, []byte(user)...)
}

// UserApplicationLinksApplicationPrefix returns the store prefix used to identify all the application
// links for the given user and application
func UserApplicationLinksApplicationPrefix(user, application string) []byte {
	return append(UserApplicationLinksPrefix(user), []byte(strings.ToLower(application))...)
}

// UserApplicationLinkKey returns the key used to store the data about the application link
// of the given user for the specified application and username
func UserApplicationLinkKey(user, application, username string) []byte {
	return append(UserApplicationLinksApplicationPrefix(user, application), []byte(strings.ToLower(username))...)
}

// ApplicationLinkClientIDKey returns the key used to store the reference to the application link
// associated with the specified client id
func ApplicationLinkClientIDKey(clientID string) []byte {
	return append(ApplicationLinkClientIDPrefix, []byte(clientID)...)
}

// ApplicationLinkAppKey returns the key used to store all the application
// links associated to the given application
func ApplicationLinkAppKey(application string) []byte {
	return append(ApplicationLinkAppPrefix, []byte(application)...)
}

// ApplicationLinkAppUsernameKey returns the key used to store all the application
// links for the given application and username
func ApplicationLinkAppUsernameKey(application, username string) []byte {
	return append(ApplicationLinkAppKey(application), append(Separator, []byte(username)...)...)
}

// ApplicationLinkOwnerKey returns the key used to store the given owner associating it to the application link
// having the provided application and username
func ApplicationLinkOwnerKey(application, username, owner string) []byte {
	return append(ApplicationLinkAppUsernameKey(application, username), []byte(owner)...)
}
