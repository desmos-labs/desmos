package types

// DONTCOVER

import (
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

var (
	// PermissionReportContent allows users to report contents
	PermissionReportContent = subspacestypes.RegisterPermission("report_content")

	// PermissionDeleteOwnReports allows users to delete existing reports made by their own
	PermissionDeleteOwnReports = subspacestypes.RegisterPermission("delete_own_reports")

	// PermissionManageReports allows users to manage other users reports
	PermissionManageReports = subspacestypes.RegisterPermission("manage_reports")

	// PermissionManageReasons allows users to manage a subspace reasons for reporting
	PermissionManageReasons = subspacestypes.RegisterPermission("manage_reasons")
)
