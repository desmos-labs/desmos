package v2

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
)

// NewApplicationLink allows to build a new ApplicationLink instance
func NewApplicationLink(
	user string, data Data, state ApplicationLinkState, oracleRequest OracleRequest, result *Result,
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

// NewOracleRequest allows to build a new OracleRequest instance
func NewOracleRequest(id uint64, scriptID uint64, callData OracleRequest_CallData, clientID string) OracleRequest {
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

// NewData allows to build a new Data instance
func NewData(application, username string) Data {
	return Data{
		Application: application,
		Username:    username,
	}
}

// NewSuccessResult allows to build a new Result instance representing a success
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

// NewErrorResult allows to build a new Result instance representing an error
func NewErrorResult(error string) *Result {
	return &Result{
		Sum: &Result_Failed_{
			Failed: &Result_Failed{
				Error: error,
			},
		},
	}
}

// MustMarshalApplicationLink serializes the given application link using the provided BinaryCodec
func MustMarshalApplicationLink(cdc codec.BinaryCodec, link ApplicationLink) []byte {
	return cdc.MustMarshal(&link)
}
