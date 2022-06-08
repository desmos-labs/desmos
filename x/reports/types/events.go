package types

// DONTCOVER

const (
	EventTypeCreateReport          = "create_report"
	EventTypeReportPost            = "report_post"
	EventTypeReportUser            = "report_user"
	EventTypeDeleteReport          = "delete_report"
	EventTypeSupportStandardReason = "support_standard_reason"
	EventTypeAddReason             = "add_reason"
	EventTypeRemoveReason          = "remove_reason"

	AttributeValueCategory       = ModuleName
	AttributeKeySubspaceID       = "subspace_id"
	AttributeKeyReportID         = "report_id"
	AttributeKeyReporter         = "reporter"
	AttributeKeyCreationTime     = "creation_time"
	AttributeKeyStandardReasonID = "standard_reason_id"
	AttributeKeyReasonID         = "reason_id"
	AttributeKeyPostID           = "post_id"
	AttributeKeyUser             = "user"
)
