package types

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewApplicationLink allows to build a new ApplicationLink instance
func NewApplicationLink(
	user string, data Data, state ApplicationLinkState, oracleRequest OracleRequest, result *Result, creationTime time.Time,
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

// Validate returns an error if the instance does not contain valid data
func (l ApplicationLink) Validate() error {
	_, err := sdk.AccAddressFromBech32(l.User)
	if err != nil {
		return fmt.Errorf("invalid user address: %s", err)
	}

	err = l.Data.Validate()
	if err != nil {
		return err
	}

	err = l.OracleRequest.Validate()
	if err != nil {
		return err
	}

	if l.Result != nil {
		err = l.Result.Validate()
		if err != nil {
			return err
		}
	}

	if l.CreationTime.IsZero() {
		return fmt.Errorf("invalid creation time: %s", l.CreationTime)
	}

	return nil
}

// IsVerificationOngoing tells whether the verification for the link is still ongoing
func (l *ApplicationLink) IsVerificationOngoing() bool {
	return l.State == ApplicationLinkStateInitialized || l.State == AppLinkStateVerificationStarted
}

// IsVerificationCompleted tells whether the verification for the link has completed or not
func (l *ApplicationLink) IsVerificationCompleted() bool {
	return l.State == AppLinkStateVerificationSuccess ||
		l.State == AppLinkStateVerificationError ||
		l.State == AppLinkStateVerificationTimedOut
}

// MustMarshalApplicationLink serializes the given application link using the provided BinaryCodec
func MustMarshalApplicationLink(cdc codec.BinaryCodec, link ApplicationLink) []byte {
	return cdc.MustMarshal(&link)
}

// MustUnmarshalApplicationLink deserializes the given byte array as an application link using
// the provided BinaryCodec
func MustUnmarshalApplicationLink(cdc codec.BinaryCodec, bz []byte) ApplicationLink {
	var link ApplicationLink
	cdc.MustUnmarshal(bz, &link)
	return link
}

// --------------------------------------------------------------------------------------------------------------------

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

// --------------------------------------------------------------------------------------------------------------------

// NewOracleRequest allows to build a new OracleRequest instance
func NewOracleRequest(id uint64, scriptID uint64, callData OracleRequest_CallData, clientID string) OracleRequest {
	return OracleRequest{
		ID:             id,
		OracleScriptID: scriptID,
		CallData:       callData,
		ClientID:       clientID,
	}
}

// Validate returns an error if the instance does not contain valid data
func (o OracleRequest) Validate() error {
	if o.OracleScriptID <= 0 {
		return fmt.Errorf("invalid oracle script id: %d", o.OracleScriptID)
	}

	err := o.CallData.Validate()
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(o.ClientID)) == 0 {
		return fmt.Errorf("client id cannot be empty or blank")
	}

	return nil
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

// --------------------------------------------------------------------------------------------------------------------

// Validate returns an error if the instance does not contain valid data
func (r *Result) Validate() error {
	var err error
	switch result := (r.Sum).(type) {
	case *Result_Success_:
		err = result.Validate()
	case *Result_Failed_:
		err = result.Validate()
	}
	return err
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

// Validate returns an error if the instance does not contain valid data
func (r Result_Failed_) Validate() error {
	if len(strings.TrimSpace(r.Failed.Error)) == 0 {
		return fmt.Errorf("error message cannot be empty or blank")
	}

	return nil
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

// Validate returns an error if the instance does not contain valid data
func (r Result_Success_) Validate() error {
	if len(strings.TrimSpace(r.Success.Value)) == 0 {
		return fmt.Errorf("value cannot be empty or blank")
	}

	if len(strings.TrimSpace(r.Success.Signature)) == 0 {
		return fmt.Errorf("signature cannot be empty or blank")
	}

	if _, err := hex.DecodeString(r.Success.Signature); err != nil {
		return fmt.Errorf("invalid signature encoding; must be hex")
	}

	return nil
}
