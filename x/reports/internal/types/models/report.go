package models

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Report is the struct of a post's reports
type Report struct {
	Type    string         `json:"type" yaml:"type"`       // Identifies the type of the reports
	Message string         `json:"message" yaml:"message"` // Contains the user message
	User    sdk.AccAddress `json:"user" yaml:"user"`       // Identifies the reporting user
}

// NewReport returns a Report
func NewReport(t string, message string, user sdk.AccAddress) Report {
	return Report{
		Type:    t,
		Message: message,
		User:    user,
	}
}

// Validate implements validator
func (r Report) Validate() error {
	if len(strings.TrimSpace(r.Type)) == 0 {
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
