package models

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "profiles"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionSaveProfile        = "save_profile"
	ActionDeleteProfile      = "delete_profile"
	ActionRequestDtag        = "request_dtag"
	ActionAcceptDtagTransfer = "accept_dtag_request"

	//Queries
	QuerierRoute      = ModuleName
	QueryProfile      = "profile"
	QueryProfiles     = "all"
	QueryDTagRequests = "dtag-requests"
	QueryParams       = "params"
)

var (
	ProfileStorePrefix    = []byte("profile")
	DtagStorePrefix       = []byte("dtag")
	TransferRequestPrefix = []byte("transfer_requests")
)

// ProfileStoreKey turns an address to a key used to store a profile into the profiles store
func ProfileStoreKey(address sdk.AccAddress) []byte {
	return append(ProfileStorePrefix, address...)
}

// DtagStoreKey turns a dtag to a key used to store a dtag -> address couple
func DtagStoreKey(dtag string) []byte {
	return append(DtagStorePrefix, []byte(dtag)...)
}

// DtagTransferRequestStoreKey turns an address to a key used to store a transfer request into the profiles store
func DtagTransferRequestStoreKey(address sdk.AccAddress) []byte {
	return append(TransferRequestPrefix, address...)
}
