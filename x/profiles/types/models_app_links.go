package types

import (
	"encoding/hex"
	"fmt"
	"strings"
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

// Validate returns an error if the instance does not contain valid data
func (d Data) Validate() error {
	if len(strings.TrimSpace(d.Application)) == 0 {
		return fmt.Errorf("application name cannot be empty or blank")
	}

	if len(strings.TrimSpace(d.Username)) == 0 {
		return fmt.Errorf("application username cannot be empty or blank")
	}

	return nil
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

// Validate returns an error if the instance does not contain valid data
func (c OracleRequest_CallData) Validate() error {
	if len(strings.TrimSpace(c.Application)) == 0 {
		return fmt.Errorf("application cannot be empty or blank")
	}

	if len(strings.TrimSpace(c.CallData)) == 0 {
		return fmt.Errorf("call data cannot be empty or blank")
	}

	if _, err := hex.DecodeString(c.CallData); err != nil {
		return fmt.Errorf("invalid call data encoding: must be hex")
	}

	return nil
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

// NewClientRequest allows to build a new ClientRequest instance
func NewClientRequest(user, application, username string) ClientRequest {
	return ClientRequest{
		User:        user,
		Application: application,
		Username:    username,
	}
}
