package types

// DONTCOVER

const (
	EventTypeCreatedReport           = "created_report"
	EventTypeReportedPost            = "reported_post"
	EventTypeReportedUser            = "reported_user"
	EventTypeDeletedReport           = "deleted_report"
	EventTypeSupportedStandardReason = "supported_standard_reason"
	EventTypeAddedReportingReason    = "added_reason"
	EventTypeRemovedReportingReason  = "removed_reason"

	AttributeKeySubspaceID       = "subspace_id"
	AttributeKeyReportID         = "report_id"
	AttributeKeyReporter         = "reporter"
	AttributeKeyCreationTime     = "creation_time"
	AttributeKeyStandardReasonID = "standard_reason_id"
	AttributeKeyReasonID         = "reason_id"
	AttributeKeyPostID           = "post_id"
	AttributeKeyUser             = "user"
)
