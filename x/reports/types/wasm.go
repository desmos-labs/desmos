package types

import "encoding/json"

type ReportsMsg struct {
	CreateReport          *json.RawMessage `json:"create_post"`
	DeleteReport          *json.RawMessage `json:"edit_post"`
	SupportStandardReason *json.RawMessage `json:"support_standard_reason"`
	AddReason             *json.RawMessage `json:"add_reason"`
	RemoveReason          *json.RawMessage `json:"remove_reason"`
}

type ReportsQuery struct {
	Reports *json.RawMessage `json:"reports"`
	Report  *json.RawMessage `json:"report"`
	Reasons *json.RawMessage `json:"reasons"`
	Reason  *json.RawMessage `json:"reason"`
}
