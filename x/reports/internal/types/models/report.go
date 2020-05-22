package models

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/reports/internal/types/models/common"
)

// Report is the struct of a post's reports
type Report struct {
	Type    string         `json:"type" yaml:"type"`       // Identifies the type of the reports
	Message string         `json:"message" yaml:"message"` // Contains the user message
	User    sdk.AccAddress `json:"user" yaml:"user"`       // Identifies the reporting user
}

// NewReport returns a Report
func NewReport(t, message string, user sdk.AccAddress) Report {
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

// ValidReportType checks if the reports's type is one of the preset list
func (r Report) ValidReportType() (valid bool) {
	for _, repType := range common.ReportsTypes {
		if r.Type == repType {
			return true
		}
	}
	return false
}

// Validate implements validator
func (r Report) Validate() error {
	if len(strings.TrimSpace(r.Type)) == 0 || !r.ValidReportType() {
		return fmt.Errorf("invalid reports type, please referes to our official reports's type list to check the valid ones")
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

// AppendIfMissing appends the given Report to the Reports slice if it does not exist inside it yet.
// It returns a new slice of Reports containing such reports and a boolean indicating whether or not the original
// slice has been modified.
func (reports Reports) AppendIfMissing(report Report) (Reports, bool) {
	for _, r := range reports {
		if r.Equals(report) {
			return reports, false
		}
	}
	return append(reports, report), true
}
