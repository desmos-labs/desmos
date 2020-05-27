package models

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ReportType string

// Empty returns true if the report type value is empty
func (rt ReportType) Empty() bool {
	return len(strings.TrimSpace(string(rt))) == 0
}

type ReportTypes []ReportType

// Contains checks if the rts array contains the given reportType
func (rts ReportTypes) Contains(reportType ReportType) bool {
	//check for cli test purposes
	if rts == nil {
		return true
	}

	for _, rt := range rts {
		if rt == reportType {
			return true
		}
	}
	return false
}

// Report is the struct of a post's reports
type Report struct {
	Type    ReportType     `json:"type" yaml:"type"`       // Identifies the type of the reports
	Message string         `json:"message" yaml:"message"` // Contains the user message
	User    sdk.AccAddress `json:"user" yaml:"user"`       // Identifies the reporting user
}

// NewReport returns a Report
func NewReport(t ReportType, message string, user sdk.AccAddress) Report {
	return Report{
		Type:    t,
		Message: message,
		User:    user,
	}
}

// String implements fmt.Stringer
func (r Report) String() string {
	bytes, err := json.Marshal(&r)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Validate implements validator
func (r Report) Validate() error {
	if r.Type.Empty() {
		return fmt.Errorf("report type cannot be empty")
	}

	if len(strings.TrimSpace(r.Message)) == 0 {
		return fmt.Errorf("reports's message cannot be empty")
	}

	if r.User.Empty() {
		return fmt.Errorf("invalid user address %s", r.User)
	}

	return nil
}

// Equals checks if the two reports are the same or not
func (r Report) Equals(other Report) bool {
	return r.Type == other.Type &&
		r.Message == other.Message &&
		r.User.Equals(other.User)
}

type Reports []Report

// String implements stringer
func (reports Reports) String() string {
	out := "Type - Message - User\n"
	for _, rep := range reports {
		out += fmt.Sprintf("%s - %s - %s\n",
			rep.Type, rep.Message, rep.User)
	}
	return strings.TrimSpace(out)
}

// Validate implements validator
func (reports Reports) Validate() error {
	for _, rep := range reports {
		if err := rep.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Equals checks if the two reports slices are equal
func (reports Reports) Equals(other Reports) bool {
	if len(reports) != len(other) {
		return false
	}

	for index, rep := range reports {
		if !rep.Equals(other[index]) {
			return false
		}
	}

	return true
}
