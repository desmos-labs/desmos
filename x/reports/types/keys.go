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
	ReportIDPrefix = []byte{0x01}
	ReportPrefix   = []byte{0x02}

	ReasonIDPrefix = []byte{0x03}
	ReasonPrefix   = []byte{0x04}
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

// ReportIDStoreKey returns the key used to store the next report id for the given subspace
func ReportIDStoreKey(subspaceID uint64) []byte {
	return append(ReportIDPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// GetSubspaceIDReportIDBytes returns the byte representation of the subspaceID and the reportID concatenated
func GetSubspaceIDReportIDBytes(subspaceID uint64, reportID uint64) []byte {
	return append(subspacestypes.GetSubspaceIDBytes(subspaceID), GetReportIDBytes(reportID)...)
}

// ReportStoreKey returns the key used  to store the report with the given subspace id and report id
func ReportStoreKey(subspaceID uint64, reportID uint64) []byte {
	return append(ReportPrefix, GetSubspaceIDReportIDBytes(subspaceID, reportID)...)
}

// --------------------------------------------------------------------------------------------------------------------

// GetReasonIDBytes returns the byte representation of the reasonID
func GetReasonIDBytes(reasonID uint32) (reasonIDBz []byte) {
	reasonIDBz = make([]byte, 4)
	binary.BigEndian.Uint32(reasonIDBz)
	return reasonIDBz
}

// GetReasonIDFromBytes returns reasonID in uint32 format from a byte array
func GetReasonIDFromBytes(bz []byte) (reasonID uint32) {
	return binary.BigEndian.Uint32(bz)
}

// ReasonIDStoreKey returns the key used to store the next reason id for the given subspace
func ReasonIDStoreKey(subspaceID uint64) []byte {
	return append(ReasonIDPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// GetSubspaceIDReasonIDBytes returns the byte representation of the given subspaceID concatenated with reasonID
func GetSubspaceIDReasonIDBytes(subspaceID uint64, reasonID uint32) []byte {
	return append(subspacestypes.GetSubspaceIDBytes(subspaceID), GetReasonIDBytes(reasonID)...)
}

// ReasonStoreKey returns the key used to store the reason with the given subspace id and reason id
func ReasonStoreKey(subspaceID uint64, reasonID uint32) []byte {
	return append(ReasonPrefix, GetSubspaceIDReasonIDBytes(subspaceID, reasonID)...)
}
