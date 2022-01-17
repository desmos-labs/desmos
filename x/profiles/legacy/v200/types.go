package v200

import (
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
	"time"
)

// NewApplicationLink allows to build a new ApplicationLink instance
func NewApplicationLink(
	user string, data types.Data, state types.ApplicationLinkState, oracleRequest types.OracleRequest, result *types.Result,
	creationTime time.Time,
) ApplicationLink {
	return ApplicationLink{
		User:          user,
		Data:          data,
		State:         state,
		OracleRequest: oracleRequest,
		Result:        result,
		CreationTime:  creationTime,
	}
}
