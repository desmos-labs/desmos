package types

import (
	"encoding/binary"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// DONTCOVER

const (
	ModuleName = "reports"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateReport   = "create_report"
	ActionDeleteReport   = "delete_report"
	ActionSupportReasons = "support_reasons"
	ActionAddReason      = "add_reason"
	ActionRemoveReason   = "remove_reason"
)

var (
	NextReportIDPrefix = []byte{0x01}
	ReportPrefix       = []byte{0x02}

	NextReasonIDPrefix = []byte{0x10}
	ReasonPrefix       = []byte{0x11}
)

// GetReportIDBytes returns the byte representation of the reportID
func GetReportIDBytes(reportID uint64) (reportIDBz []byte) {
	reportIDBz = make([]byte, 8)
	binary.BigEndian.PutUint64(reportIDBz, reportID)
	return reportIDBz
}

// GetReportIDFromBytes returns reportID in uint64 format from a byte array
func GetReportIDFromBytes(bz []byte) (reportID uint64) {
	return binary.BigEndian.Uint64(bz)
}

// NextReportIDStoreKey returns the key used to store the next report id for the given subspace
func NextReportIDStoreKey(subspaceID uint64) []byte {
	return append(NextReportIDPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// SubspaceReportsPrefix returns the store prefix used to store all the reports related to the given subspace
func SubspaceReportsPrefix(subspaceID uint64) []byte {
	return append(ReportPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// ReportStoreKey returns the key used  to store the report with the given subspace id and report id
func ReportStoreKey(subspaceID uint64, reportID uint64) []byte {
	return append(SubspaceReportsPrefix(subspaceID), GetReportIDBytes(reportID)...)
}

// --------------------------------------------------------------------------------------------------------------------

// GetReasonIDBytes returns the byte representation of the reasonID
func GetReasonIDBytes(reasonID uint32) (reasonIDBz []byte) {
	reasonIDBz = make([]byte, 4)
	binary.BigEndian.PutUint32(reasonIDBz, reasonID)
	return reasonIDBz
}

// GetReasonIDFromBytes returns reasonID in uint32 format from a byte array
func GetReasonIDFromBytes(bz []byte) (reasonID uint32) {
	return binary.BigEndian.Uint32(bz)
}

// NextReasonIDStoreKey returns the key used to store the next reason id for the given subspace
func NextReasonIDStoreKey(subspaceID uint64) []byte {
	return append(NextReasonIDPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// SubspaceReasonsPrefix returns the store prefix used to store all the reports for the given subspace
func SubspaceReasonsPrefix(subspaceID uint64) []byte {
	return append(ReasonPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// ReasonStoreKey returns the key used to store the reason with the given subspace id and reason id
func ReasonStoreKey(subspaceID uint64, reasonID uint32) []byte {
	return append(SubspaceReasonsPrefix(subspaceID), GetReasonIDBytes(reasonID)...)
}
