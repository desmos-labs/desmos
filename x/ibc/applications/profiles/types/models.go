package types

import (
	"time"
)

// NewApplicationData allows to build a new ApplicationData instance
func NewApplicationData(application, username string) *ApplicationData {
	return &ApplicationData{
		Name:     application,
		Username: username,
	}
}

// NewVerificationData allows to build a new VerificationData instance
func NewVerificationData(application, callData string) *VerificationData {
	return &VerificationData{
		Application: application,
		CallData:    callData,
	}
}

// NewOracleRequest Allows to build a new OracleRequest instance
func NewOracleRequest(id int64, oracleScriptID int64, clientID string) *OracleRequest {
	return &OracleRequest{
		ID:             id,
		OracleScriptID: oracleScriptID,
		ClientId:       clientID,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// NewErrorResult allows to build a new Result from the given error string
func NewErrorResult(error string) *ApplicationLinkResult {
	return &ApplicationLinkResult{
		Sum: &ApplicationLinkResult_Failed_{
			Failed: &ApplicationLinkResult_Failed{
				Error: error,
			},
		},
	}
}

// NewSuccessResult allows to build a new Result from the given value and signature
func NewSuccessResult(value, signature string) *ApplicationLinkResult {
	return &ApplicationLinkResult{
		Sum: &ApplicationLinkResult_Success_{
			Success: &ApplicationLinkResult_Success{
				Value:     value,
				Signature: signature,
			},
		},
	}
}

// -------------------------------------------------------------------------------------------------------------------

// NewApplicationLink allows to build a new Connection instance
func NewApplicationLink(
	user string, application *ApplicationData, verification *VerificationData,
	state ApplicationLinkState, oracleRequest *OracleRequest, result *ApplicationLinkResult, creationTime time.Time,
) *ApplicationLink {
	return &ApplicationLink{
		User:          user,
		Application:   application,
		Verification:  verification,
		State:         state,
		OracleRequest: oracleRequest,
		Result:        result,
		CreationTime:  creationTime,
	}
}
