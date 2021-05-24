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
func NewVerificationData(method, value string) *VerificationData {
	return &VerificationData{
		Method: method,
		Value:  value,
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
func NewErrorResult(error string) *Result {
	return &Result{
		Sum: &Result_Failed_{
			Failed: &Result_Failed{
				Error: error,
			},
		},
	}
}

// NewSuccessResult allows to build a new Result from the given value and signature
func NewSuccessResult(value, signature string) *Result {
	return &Result{
		Sum: &Result_Success_{
			Success: &Result_Success{
				Value:     value,
				Signature: signature,
			},
		},
	}
}

// -------------------------------------------------------------------------------------------------------------------

// NewConnection allows to build a new Connection instance
func NewConnection(
	user string, application *ApplicationData, verification *VerificationData,
	state ConnectionState, oracleRequest *OracleRequest, result *Result, creationTime time.Time,
) *Connection {
	return &Connection{
		User:          user,
		Application:   application,
		Verification:  verification,
		State:         state,
		OracleRequest: oracleRequest,
		Result:        result,
		CreationTime:  creationTime,
	}
}
