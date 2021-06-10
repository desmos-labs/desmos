package types

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
	DTagPrefix                 = []byte("dtag")
	DTagTransferRequestsPrefix = []byte("transfer_requests")
	RelationshipsStorePrefix   = []byte("relationships")
	UsersBlocksStorePrefix     = []byte("users_blocks")
	ChainsLinksPrefix          = []byte("chains_links")

	// IBCPortKey defines the key to store the port ID in store
	IBCPortKey = []byte("ibc-port")
)

// DTagStoreKey turns a DTag into the key used to store the address associated with it into the store
func DTagStoreKey(dTag string) []byte {
	return append(DTagPrefix, []byte(dTag)...)
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

// RelationshipsStoreKey turns an user address, a recipient address and a subspace to a key used to store a Address -> Address couple
func RelationshipsStoreKey(user, subspace, recipient string) []byte {
	return append(UserRelationshipsSubspacePrefix(user, subspace), []byte(recipient)...)
}

// UsersBlocksStoreKey turns a user address to a key used to store a Address -> []Address couple
func UsersBlocksStoreKey(user string) []byte {
	return append(UsersBlocksStorePrefix, []byte(user)...)
}

// ChainsLinksStoreKey turns an address and chain name to a key used to store a Link
func ChainsLinksStoreKey(chainName string, address string) []byte {
	return append(ChainsLinksPrefix, []byte(chainName+address)...)
}
