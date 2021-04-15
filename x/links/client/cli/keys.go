package cli

import "time"

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	FlagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)
