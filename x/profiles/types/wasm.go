package types

import "encoding/json"

type ProfilesMsg struct {
	SaveProfile               *MsgSaveProfile               `json:"save_profile,omitempty"`
	DeleteProfile             *MsgDeleteProfile             `json:"delete_profile,omitempty"`
	RequestDtagTransfer       *MsgRequestDTagTransfer       `json:"request_dtag_transfer"`
	AcceptDtagTransferRequest *MsgAcceptDTagTransferRequest `json:"accept_dtag_transfer_request"`
	RefuseDtagTransferRequest *MsgRefuseDTagTransferRequest `json:"refuse_dtag_transfer_request"`
	CancelDtagTransferRequest *MsgCancelDTagTransferRequest `json:"cancel_dtag_transfer_request"`
}

type ProfilesQueryRoutes struct {
	Profile *ProfileQuery `json:"profile"`
}

type ProfileQuery struct {
	Request *QueryProfileRequest `json:"request"`
}

// UpdateDtagAuctionStatus represent the sudo message that's triggered from the profile module to update the status of an auction
// for the given user inside a DTag Auctioneer contract
type UpdateDtagAuctionStatus struct {
	User           string `json:"user"`
	TransferStatus string `json:"transfer_status"`
}

func NewUpdateDTagAuctionStatus(user, transferStatus string) UpdateDtagAuctionStatus {
	return UpdateDtagAuctionStatus{
		User:           user,
		TransferStatus: transferStatus,
	}
}

func (updateAS UpdateDtagAuctionStatus) Marshal() ([]byte, error) {
	bz, err := json.Marshal(&updateAS)
	if err != nil {
		return nil, err
	}
	return bz, nil
}
