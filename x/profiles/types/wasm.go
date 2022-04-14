package types

import (
	"encoding/json"
)

type ProfilesMsg struct {
	SaveProfile               json.RawMessage `json:"save_profile"`
	DeleteProfile             json.RawMessage `json:"delete_profile"`
	RequestDtagTransfer       json.RawMessage `json:"request_dtag_transfer"`
	AcceptDtagTransferRequest json.RawMessage `json:"accept_dtag_transfer_request"`
	RefuseDtagTransferRequest json.RawMessage `json:"refuse_dtag_transfer_request"`
	CancelDtagTransferRequest json.RawMessage `json:"cancel_dtag_transfer_request"`
	LinkChainAccount          json.RawMessage `json:"link_chain_account"`
	LinkApplication           json.RawMessage `json:"link_application"`
}

type ProfilesQuery struct {
	Profile                      json.RawMessage `json:"profile"`
	IncomingDtagTransferRequests json.RawMessage `json:"incoming_dtag_transfer_requests"`
	ChainLinks                   json.RawMessage `json:"chain_links"`
	UserChainLink                json.RawMessage `json:"user_chain_link"`
	AppLinks                     json.RawMessage `json:"app_links"`
	UserAppLinks                 json.RawMessage `json:"user_app_links"`
	ApplicationLinkByClientID    json.RawMessage `json:"application_link_by_client_id"`
}
