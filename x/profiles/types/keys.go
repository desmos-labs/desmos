package types

//DONTCOVER

const (
	ModuleName = "profiles"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

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

	QuerierRoute              = ModuleName
	QueryProfile              = "profile"
	QueryIncomingDTagRequests = "incoming-dtag-requests"
	QueryUserRelationships    = "user_relationships"
	QueryRelationships        = "relationships"
	QueryUserBlocks           = "user_blocks"
	QueryParams               = "params"

	DoNotModify = "[do-not-modify]"

	// IBC keys
	Version = "profiles-1"
	PortID  = "profiles"
)

var (
	DTagPrefix                 = []byte("dtag")
	DTagTransferRequestsPrefix = []byte("transfer_requests")
	RelationshipsStorePrefix   = []byte("relationships")
	UsersBlocksStorePrefix     = []byte("users_blocks")

	// PortKey defines the key to store the port ID in store
	PortKey = []byte("port")
)

// DTagStoreKey turns a DTag into the key used to store the address associated with it into the store
func DTagStoreKey(dTag string) []byte {
	return append(DTagPrefix, []byte(dTag)...)
}

// DTagTransferRequestStoreKey turns an address to a key used to store a transfer request into the profiles store
func DTagTransferRequestStoreKey(address string) []byte {
	return append(DTagTransferRequestsPrefix, address...)
}

// RelationshipsStoreKey turns a user address to a key used to store a Address -> []Address couple
func RelationshipsStoreKey(user string) []byte {
	return append(RelationshipsStorePrefix, []byte(user)...)
}

// UsersBlocksStoreKey turns a user address to a key used to store a Address -> []Address couple
func UsersBlocksStoreKey(user string) []byte {
	return append(UsersBlocksStorePrefix, []byte(user)...)
}
