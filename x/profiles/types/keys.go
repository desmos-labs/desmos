package types

import "strings"

// DONTCOVER

const (
	ModuleName = "profiles"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	DesmosChainName = "desmos"

	ActionSaveProfile               = "save_profile"
	ActionDeleteProfile             = "delete_profile"
	ActionRequestDTag               = "request_dtag"
	ActionAcceptDTagTransfer        = "accept_dtag_request"
	ActionRefuseDTagTransferRequest = "refuse_dtag_request"
	ActionCancelDTagTransferRequest = "cancel_dtag_request"
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

	// IBC keys
	IBCVersion = "ibc-profiles-1"
	IBCPortID  = "ibc-profiles"
)

var (
	DTagPrefix                    = []byte("dtag")
	DTagTransferRequestsPrefix    = []byte("transfer_requests")
	RelationshipsStorePrefix      = []byte("relationships")
	UsersBlocksStorePrefix        = []byte("users_blocks")
	ChainLinksPrefix              = []byte("chain_links")
	UserApplicationLinkPrefix     = []byte("user_application_link")
	ApplicationLinkPrefix         = []byte("application_link")
	ApplicationLinkClientIDPrefix = []byte("client_id")

	// IBCPortKey defines the key to store the port ID in store
	IBCPortKey = []byte{0x01}
)

// DTagStoreKey turns a DTag into the key used to store the address associated with it into the store
func DTagStoreKey(dTag string) []byte {
	return append(DTagPrefix, []byte(strings.ToLower(dTag))...)
}

// DTagTransferRequestStoreKey turns an address to a key used to store a transfer request into the profiles store
func DTagTransferRequestStoreKey(address string) []byte {
	return append(DTagTransferRequestsPrefix, address...)
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

// BLockPrefix returns the store prefix used to store the blocks created by the given blocker
func BlockerPrefix(blocker string) []byte {
	return append(UsersBlocksStorePrefix, []byte(blocker)...)
}

func BlockerSubspacePrefix(blocker string, subspace string) []byte {
	return append(BlockerPrefix(blocker), []byte(subspace)...)
}

// UsersBlocksStoreKey turns a user address to a key used to store a Address -> []Address couple
func UsersBlocksStoreKey(blocker string, subspace string, blockedUser string) []byte {
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
