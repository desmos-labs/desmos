package wasm

import (
	"encoding/json"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

type ProfilesMsg struct {
	SaveProfile               *types.MsgSaveProfile               `json:"save_profile,omitempty"`
	DeleteProfile             *types.MsgDeleteProfile             `json:"delete_profile,omitempty"`
	RequestDtagTransfer       *types.MsgRequestDTagTransfer       `json:"request_dtag_transfer"`
	AcceptDtagTransferRequest *types.MsgAcceptDTagTransferRequest `json:"accept_dtag_transfer_request"`
	RefuseDtagTransferRequest *types.MsgRefuseDTagTransferRequest `json:"refuse_dtag_transfer_request"`
	CancelDtagTransferRequest *types.MsgCancelDTagTransferRequest `json:"cancel_dtag_transfer_request"`
}

type ProfilesQueryRoutes struct {
	Profile *ProfileQuery `json:"profile"`
}

type ProfileQuery struct {
	Request *types.QueryProfileRequest `json:"request"`
}

// UpdateDTagAuctionStatus represent the sudo message that's triggered from the profile module to update the status of an auction
// for the given user inside a DTag Auctioneer contract
type UpdateDTagAuctionStatus struct {
	User           string `json:"user"`
	TransferStatus string `json:"transfer_status"`
}

func NewUpdateDTagAuctionStatus(user, transferStatus string) UpdateDTagAuctionStatus {
	return UpdateDTagAuctionStatus{
		User:           user,
		TransferStatus: transferStatus,
	}
}

func (updateAS UpdateDTagAuctionStatus) Marshal() ([]byte, error) {
	bz, err := json.Marshal(&updateAS)
	if err != nil {
		return nil, err
	}
	return bz, nil
}
