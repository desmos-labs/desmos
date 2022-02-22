package types

import "encoding/json"

type ProfilesMsgsRoutes struct {
	Profiles ProfilesMsgs `json:"profiles"`
}

type ProfilesMsgs struct {
	SaveProfile               *MsgSaveProfile               `json:"save_profile,omitempty"`
	DeleteProfile             *MsgDeleteProfile             `json:"delete_profile,omitempty"`
	RequestDtagTransfer       *MsgRequestDTagTransfer       `json:"request_dtag_transfer"`
	AcceptDtagTransferRequest *MsgAcceptDTagTransferRequest `json:"accept_dtag_transfer_request"`
	RefuseDtagTransferRequest *MsgRefuseDTagTransferRequest `json:"refuse_dtag_transfer_request"`
	CancelDtagTransferRequest *MsgCancelDTagTransferRequest `json:"cancel_dtag_transfer_request"`
}

type ProfilesQueryRoutes struct {
	Profiles ProfilesQueryRequests `json:"profiles"`
}

type ProfilesQueryRequests struct {
	Profile                      *QueryProfileRequest                      `json:"profile"`
	Relationships                *QueryRelationshipsRequest                `json:"relationships"`
	IncomingDtagTransferRequests *QueryIncomingDTagTransferRequestsRequest `json:"incoming_dtag_transfer_requests"`
	Blocks                       *QueryBlocksRequest                       `json:"blocks"`
	ChainLinks                   *QueryChainLinksRequest                   `json:"chain_links"`
	UserChainLink                *QueryUserChainLinkRequest                `json:"user_chain_link"`
	AppLinks                     *QueryApplicationLinksRequest             `json:"app_links"`
	UserAppLinks                 *QueryUserApplicationLinkRequest          `json:"user_app_links"`
	ApplicationLinkByClientID    *QueryApplicationLinkByClientIDRequest    `json:"application_link_by_client_id"`
}

// UpdateDtagAuctionStatus represent the sudo message that's triggered from the profile module to update the status of an auction
// for the given user inside a DTag Auctioneer contract
type UpdateDtagAuctionStatus struct {
	User           string `json:"user"`
	TransferStatus string `json:"transfer_status"`
}

type SudoMsg struct {
	UpdateDtagAuctionStatus UpdateDtagAuctionStatus `json:"update_dtag_auction_status"`
}

func NewUpdateDTagAuctionStatusMsg(user, transferStatus string) SudoMsg {
	return SudoMsg{
		UpdateDtagAuctionStatus: UpdateDtagAuctionStatus{
			TransferStatus: transferStatus,
			User:           user,
		},
	}
}

func (sudoMsg SudoMsg) Marshal() ([]byte, error) {
	bz, err := json.Marshal(&sudoMsg)
	if err != nil {
		return nil, err
	}
	return bz, nil
}
