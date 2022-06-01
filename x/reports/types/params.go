package types

import (
	"fmt"
	"strings"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultParamsSpace represents the default paramspace for the Params keeper
	DefaultParamsSpace = ModuleName
)

var (
	// DefaultReasons represents the default set of reasons that can be adopted by subspaces
	DefaultReasons []StandardReason
)

var (
	// ReasonsKey represents the params key used to store available default reasons
	ReasonsKey = []byte("StandardReasons")
)

// --------------------------------------------------------------------------------------------------------------------

// ParamKeyTable returns the key declaration for the parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().
		RegisterParamSet(&Params{})
}

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

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of reports module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(ReasonsKey, &params.StandardReasons, ValidateStandardReasonsParam),
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
