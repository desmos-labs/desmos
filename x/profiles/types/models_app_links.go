package types

import (
	"time"
)

// NewApplicationLink allows to build a new ApplicationLink instance
func NewApplicationLink(
	data Data, state ApplicationLinkState, oracleRequest OracleRequest, result *Result, creationTime time.Time,
) ApplicationLink {
	return ApplicationLink{
		Data:          data,
		State:         state,
		OracleRequest: oracleRequest,
		Result:        result,
		CreationTime:  creationTime,
	}
}

// NewData allows to build a new Data instance
func NewData(application, username string) Data {
	return Data{
		Application: application,
		Username:    username,
	}
}

// NewOracleRequest allows to build a new OracleRequest instance
func NewOracleRequest(id int64, scriptID int64, callData OracleRequest_CallData, clientID string) OracleRequest {
	return OracleRequest{
		ID:             id,
		OracleScriptID: scriptID,
		CallData:       callData,
		ClientID:       clientID,
	}
}

// NewOracleRequestCallData allows to build a new OracleRequest_CallData instance
func NewOracleRequestCallData(application, callData string) OracleRequest_CallData {
	return OracleRequest_CallData{
		Application: application,
		CallData:    callData,
	}
}

func NewErrorResult(error string) *Result {
	return &Result{
		Sum: &Result_Failed_{
			Failed: &Result_Failed{
				Error: error,
			},
		},
	}
}

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

func NewClientRequest(user, application, username string) ClientRequest {
	return ClientRequest{
		User:        user,
		Application: application,
		Username:    username,
	}
}
