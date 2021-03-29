package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
)

// NewLink returns a new Link object
func NewLink(source string, destination string) Link {
	return PollAnswer{
		Source:   source,
		Destination: destination,
	}
}
