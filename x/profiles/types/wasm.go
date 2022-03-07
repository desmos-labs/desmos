package types

import "encoding/json"

type DesmosMsgRoute struct {
	Profiles ProfilesMsg `json:"profiles"`
}

type ProfilesMsg struct {
	SaveProfile               *MsgSaveProfile               `json:"save_profile"`
	DeleteProfile             *MsgDeleteProfile             `json:"delete_profile"`
	RequestDtagTransfer       *MsgRequestDTagTransfer       `json:"request_dtag_transfer"`
	AcceptDtagTransferRequest *MsgAcceptDTagTransferRequest `json:"accept_dtag_transfer_request"`
	RefuseDtagTransferRequest *MsgRefuseDTagTransferRequest `json:"refuse_dtag_transfer_request"`
	CancelDtagTransferRequest *MsgCancelDTagTransferRequest `json:"cancel_dtag_transfer_request"`
	LinkChainAccount          *MsgLinkChainAccount          `json:"link_chain_account"`
	LinkApplication           *MsgLinkApplication           `json:"link_application"`
	CreateRelationship        *MsgCreateRelationship        `json:"create_relationship"`
	DeleteRelationship        *MsgDeleteRelationship        `json:"delete_relationship"`
	BlockUser                 *MsgBlockUser                 `json:"block_user"`
	UnblockUser               *MsgUnblockUser               `json:"unblock_user"`
}

type DesmosQueryRoute struct {
	Profiles ProfilesQuery `json:"profiles"`
}

type ProfilesQuery struct {
	Profile                      *QueryProfileRequest                      `json:"profile"`
	Relationships                json.RawMessage                           `json:"relationships"`
	IncomingDtagTransferRequests *QueryIncomingDTagTransferRequestsRequest `json:"incoming_dtag_transfer_requests"`
	Blocks                       *QueryBlocksRequest                       `json:"blocks"`
	ChainLinks                   *QueryChainLinksRequest                   `json:"chain_links"`
	UserChainLink                *QueryUserChainLinkRequest                `json:"user_chain_link"`
	AppLinks                     *QueryApplicationLinksRequest             `json:"app_links"`
	UserAppLinks                 *QueryUserApplicationLinkRequest          `json:"user_app_links"`
	ApplicationLinkByClientID    *QueryApplicationLinkByClientIDRequest    `json:"application_link_by_client_id"`
}
