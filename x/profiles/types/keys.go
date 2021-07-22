package types

import (
	"encoding/binary"
	"strings"
)

// DONTCOVER

const (
	ModuleName = "profiles"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	DesmosChainName = "desmos"

	ActionSaveProfile               = "save_profile"
	ActionDeleteProfile             = "delete_profile"
	ActionRequestDTag               = "request_dtag_transfer"
	ActionAcceptDTagTransfer        = "accept_dtag_transfer_request"
	ActionRefuseDTagTransferRequest = "refuse_dtag_transfer_request"
	ActionCancelDTagTransferRequest = "cancel_dtag_transfer_request"
	ActionCreateRelationship        = "create_relationship"
	ActionDeleteRelationship        = "delete_relationship"
	ActionBlockUser                 = "block_user"
	ActionUnblockUser               = "unblock_user"
	ActionLinkChainAccount          = "link_chain_account"
	ActionUnlinkChainAccount        = "unlink_chain_account"
	ActionLinkApplication           = "link_application"
	ActionUnlinkApplication         = "unlink_application"

	QuerierRoute              = ModuleName
	QueryProfile              = "profile"
	QueryIncomingDTagRequests = "incoming-dtag-requests"
	QueryUserRelationships    = "user_relationships"
	QueryRelationships        = "relationships"
	QueryUserBlocks           = "user_blocks"
	QueryParams               = "params"

	DoNotModify = "[do-not-modify]"

	// IBCVersion represents the version of the IBC specification to be used.
	// Currently we use ICS20-1 because we need to communicate with Band which uses the same version.
	// TODO: Explore if it's possible to support multiple specifications at once, and see if we should create our own one.
	IBCVersion = "ics20-1"
	IBCPortID  = "ibc-profiles"
)

var (
	DTagPrefix                      = []byte("dtag")
	DTagTransferRequestPrefix       = []byte("transfer_request")
	RelationshipsStorePrefix        = []byte("relationships")
	UsersBlocksStorePrefix          = []byte("users_blocks")
	ChainLinksPrefix                = []byte("chain_links")
	UserApplicationLinkPrefix       = []byte("user_application_link")
	ApplicationLinkClientIDPrefix   = []byte("client_id")
	ApplicationLinkExpirationPrefix = []byte("expiration_application_link")

	// IBCPortKey defines the key to store the port ID in store
	IBCPortKey = []byte{0x01}
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

// UserRelationshipsPrefix returns the prefix used to store all relationships created
// by the user with the given address
func UserRelationshipsPrefix(user string) []byte {
	return append(RelationshipsStorePrefix, []byte(user)...)
}

// UserRelationshipsSubspacePrefix returns the prefix used to store all the relationships created by the user
// with the given address for the subspace having the given id
func UserRelationshipsSubspacePrefix(user, subspace string) []byte {
	return append(UserRelationshipsPrefix(user), []byte(subspace)...)
}

// RelationshipsStoreKey returns the store key used to store the relationships containing the given data
func RelationshipsStoreKey(user, subspace, recipient string) []byte {
	return append(UserRelationshipsSubspacePrefix(user, subspace), []byte(recipient)...)
}

// BlockerPrefix returns the store prefix used to store the blocks created by the given blocker
func BlockerPrefix(blocker string) []byte {
	return append(UsersBlocksStorePrefix, []byte(blocker)...)
}

// BlockerSubspacePrefix returns the store prefix used to store the blocks that the given blocker
// has created inside the specified subspace
func BlockerSubspacePrefix(blocker string, subspace string) []byte {
	return append(BlockerPrefix(blocker), []byte(subspace)...)
}

// UserBlockStoreKey returns the store key used to save the block made by the given blocker,
// inside the specified subspace and towards the given blocked user
func UserBlockStoreKey(blocker string, subspace string, blockedUser string) []byte {
	return append(BlockerSubspacePrefix(blocker, subspace), []byte(blockedUser)...)
}

// UserChainLinksPrefix returns the store prefix used to identify all the chain links for the given user
func UserChainLinksPrefix(user string) []byte {
	return append(ChainLinksPrefix, []byte(user)...)
}

// ChainLinksStoreKey returns the store key used to store the chain links containing the given data
func ChainLinksStoreKey(user, chainName, address string) []byte {
	return append(UserChainLinksPrefix(user), []byte(chainName+address)...)
}

// UserApplicationLinksPrefix returns the store prefix used to identify all the application links for the given user
func UserApplicationLinksPrefix(user string) []byte {
	return append(UserApplicationLinkPrefix, []byte(user)...)
}

// UserApplicationLinkKey returns the key used to store the data about the application link
// of the given user for the specified application and username
func UserApplicationLinkKey(user, application, username string) []byte {
	return append(UserApplicationLinksPrefix(user), []byte(strings.ToLower(application)+strings.ToLower(username))...)
}

// ApplicationLinkClientIDKey returns the key used to store the reference to the application link
// associated with the specified client id
func ApplicationLinkClientIDKey(clientID string) []byte {
	return append(ApplicationLinkClientIDPrefix, []byte(clientID)...)
}

// ExpiringApplicationLinkPrefix returns the prefix used to store client id keys of the expiring application links
// at the given block height
func ExpiringApplicationLinkPrefix(blockHeight uint64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	binary.BigEndian.PutUint64(buf, blockHeight)
	return append(ApplicationLinkExpirationPrefix, buf...)
}

// ExpiringApplicationLinkKey returns the key used to store the client id key of the application link
// which will be expired at the given block height
func ExpiringApplicationLinkKey(blockHeight uint64, clientID string) []byte {
	return append(ExpiringApplicationLinkPrefix(blockHeight), []byte(clientID)...)
}
