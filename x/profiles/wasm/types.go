package wasm

import "github.com/desmos-labs/desmos/v2/x/profiles/types"

type ProfilesMessage struct {
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
