package types

import (
	"fmt"
	"strings"
)

var (
	// DefaultReasons represents the default set of reasons that can be adopted by subspaces
	DefaultReasons []StandardReason
)

// --------------------------------------------------------------------------------------------------------------------

// NewParams returns a new Params instance
func NewParams(reasons []StandardReason) Params {
	return Params{
		StandardReasons: reasons,
	}
}

// DefaultParams returns the default params
func DefaultParams() Params {
	return Params{
		StandardReasons: DefaultReasons,
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	return ValidateStandardReasonsParam(params.StandardReasons)
}

// ValidateStandardReasonsParam validates the reasons params value
func ValidateStandardReasonsParam(i interface{}) error {
	reasons, ok := i.([]StandardReason)
	if !ok {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	err := NewStandardReasons(reasons...).Validate()
	if err != nil {
		return err
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

type StandardReasons []StandardReason

// NewStandardReasons returns a new instance of StandardReasons contains the given reasons
func NewStandardReasons(reasons ...StandardReason) StandardReasons {
	return reasons
}

// Validate implements fmt.Validator
func (s StandardReasons) Validate() error {
	ids := map[uint32]bool{}
	for _, reason := range s {
		if _, duplicated := ids[reason.ID]; duplicated {
			return fmt.Errorf("duplicated reason with id %d", reason.ID)
		}
		ids[reason.ID] = true

		err := reason.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// NewStandardReason returns a new StandardReason instance
func NewStandardReason(id uint32, title string, description string) StandardReason {
	return StandardReason{
		ID:          id,
		Title:       title,
		Description: description,
	}
}

// Validate implements fmt.Validator
func (r StandardReason) Validate() error {
	if r.ID == 0 {
		return fmt.Errorf("invalid id: %d", r.ID)
	}

	if strings.TrimSpace(r.Title) == "" {
		return fmt.Errorf("invalid title: %s", r.Title)
	}

	return nil
}
