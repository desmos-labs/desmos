package types

//DONTCOVER

const (
	ModuleName = "profiles"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionSaveProfile               = "save_profile"
	ActionDeleteProfile             = "delete_profile"
	ActionRequestDtag               = "request_dtag"
	ActionAcceptDtagTransfer        = "accept_dtag_request"
	ActionRefuseDTagTransferRequest = "refuse_dtag_request"
	ActionCancelDTagTransferRequest = "cancel_dtag_request"

	QuerierRoute              = ModuleName
	QueryProfile              = "profile"
	QueryIncomingDTagRequests = "incoming-dtag-requests"
	QueryParams               = "params"

	DoNotModify = "[do-not-modify]"
)

var (
	ProfileStorePrefix         = []byte("profile")
	DtagStorePrefix            = []byte("dtag")
	DTagTransferRequestsPrefix = []byte("transfer_requests")
)

// ProfileStoreKey turns an address to a key used to store a profile into the profiles store
func ProfileStoreKey(address string) []byte {
	return append(ProfileStorePrefix, address...)
}

// DtagStoreKey turns a dtag to a key used to store a dtag -> address couple
func DtagStoreKey(dtag string) []byte {
	return append(DtagStorePrefix, []byte(dtag)...)
}

// DtagTransferRequestStoreKey turns an address to a key used to store a transfer request into the profiles store
func DtagTransferRequestStoreKey(address string) []byte {
	return append(DTagTransferRequestsPrefix, address...)
}
