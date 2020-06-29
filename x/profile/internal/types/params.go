package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramsModule "github.com/cosmos/cosmos-sdk/x/params/subspace"
)

const (
	// default paramspace for paramsModule keeper
	DefaultParamspace = ModuleName
)

// Default profile paramsModule
var (
	DefaultMinMonikerLength = sdk.NewInt(2)
	DefaultMaxMonikerLength = sdk.NewInt(1000) //longest name on earth count 954 chars
	DefaultRegEx            = `^[A-Za-z0-9_]+$`
	DefaultMinDTagLength    = sdk.NewInt(3)
	DefaultMaxDTagLength    = sdk.NewInt(30)
	DefaultMaxBioLength     = sdk.NewInt(1000)
)

// Parameters store keys
var (
	MonikerLenParamsKey = []byte("MonikerParams")
	DtagLenParamsKey    = []byte("DtagParams")
	MaxBioLenParamsKey  = []byte("MaxBioLen")
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramsModule.KeyTable {
	return paramsModule.NewKeyTable().RegisterParamSet(&Params{})
}

type Params struct {
	MonikerParams MonikerParams `json:"moniker_params" yaml:"moniker_params"`
	DtagParams    DtagParams    `json:"dtag_params" yaml:"dtag_params"`
	MaxBioLen     sdk.Int       `json:"max_bio_length" yaml:"max_bio_length"`
}

// NewParams creates a new ProfileParams obj
func NewParams(monikerLen MonikerParams, dtagLen DtagParams, maxBioLen sdk.Int) Params {
	return Params{
		MonikerParams: monikerLen,
		DtagParams:    dtagLen,
		MaxBioLen:     maxBioLen,
	}
}

// DefaultParams return default paramsModule
func DefaultParams() Params {
	return Params{
		MonikerParams: DefaultMonikerParams(),
		DtagParams:    DefaultDtagParams(),
		MaxBioLen:     DefaultMaxBioLength,
	}
}

func (params Params) String() string {
	out := "Profiles parameters:\n"
	out += fmt.Sprintf("%s\n%s\nBiography params lengths:\nMax accepted length: %s\n",
		params.MonikerParams.String(),
		params.DtagParams.String(),
		params.MaxBioLen,
	)

	return strings.TrimSpace(out)
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of profile module's parameters.
func (params *Params) ParamSetPairs() paramsModule.ParamSetPairs {
	return paramsModule.ParamSetPairs{
		paramsModule.NewParamSetPair(MonikerLenParamsKey, &params.MonikerParams, ValidateMonikerParams),
		paramsModule.NewParamSetPair(DtagLenParamsKey, &params.DtagParams, ValidateDtagParams),
		paramsModule.NewParamSetPair(MaxBioLenParamsKey, &params.MaxBioLen, ValidateBioParams),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateMonikerParams(params.MonikerParams); err != nil {
		return err
	}

	if err := ValidateDtagParams(params.DtagParams); err != nil {
		return err
	}

	return ValidateBioParams(params.MaxBioLen)
}

// MonikerParams defines the paramsModule around moniker len
type MonikerParams struct {
	MinMonikerLen sdk.Int `json:"min_length" yaml:"min_length"`
	MaxMonikerLen sdk.Int `json:"max_length" yaml:"max_length"`
}

// NewMonikerParams creates a new MonikerParams obj
func NewMonikerParams(minLen, maxLen sdk.Int) MonikerParams {
	return MonikerParams{
		MinMonikerLen: minLen,
		MaxMonikerLen: maxLen,
	}
}

// DefaultMonikerParams return default moniker params
func DefaultMonikerParams() MonikerParams {
	return NewMonikerParams(DefaultMinMonikerLength,
		DefaultMaxMonikerLength)
}

// String implements stringer interface
func (params MonikerParams) String() string {
	out := "Moniker params lengths:\n"
	out += fmt.Sprintf("Min accepted length: %s\nMax accepted length: %s",
		params.MinMonikerLen,
		params.MaxMonikerLen,
	)

	return strings.TrimSpace(out)
}

func ValidateMonikerParams(i interface{}) error {
	params, isNameSurnParams := i.(MonikerParams)

	if !isNameSurnParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.MinMonikerLen.IsNegative() || params.MinMonikerLen.LT(DefaultMinMonikerLength) {
		return fmt.Errorf("invalid minimum moniker length param: %s", params.MinMonikerLen)
	}

	// TODO make sense to cap this? I've done this thinking "what's the sense of having names higher that 1000 chars?"
	if params.MaxMonikerLen.IsNegative() || params.MaxMonikerLen.GT(DefaultMaxMonikerLength) {
		return fmt.Errorf("invalid max moniker length param: %s", params.MaxMonikerLen)
	}

	return nil
}

// DtagParams defines the paramsModule around profiles' dtag
type DtagParams struct {
	RegEx      string  `json:"reg_ex" yaml:"reg_ex"`
	MinDtagLen sdk.Int `json:"min_length" yaml:"min_length"`
	MaxDtagLen sdk.Int `json:"max_length" yaml:"max_length"`
}

// NewDtagParams creates a new DtagParams obj
func NewDtagParams(regEx string, minLen, maxLen sdk.Int) DtagParams {
	return DtagParams{
		RegEx:      regEx,
		MinDtagLen: minLen,
		MaxDtagLen: maxLen,
	}
}

// DefaultDtagParams return default paramsModule
func DefaultDtagParams() DtagParams {
	return NewDtagParams(
		DefaultRegEx,
		DefaultMinDTagLength,
		DefaultMaxDTagLength,
	)
}

// String implements stringer interface
func (params DtagParams) String() string {
	out := "Dtag params:\n"
	out += fmt.Sprintf("RegEx: %s\nMin accepted length: %s\nMax accepted length: %s",
		params.RegEx,
		params.MinDtagLen,
		params.MaxDtagLen,
	)

	return strings.TrimSpace(out)
}

func ValidateDtagParams(i interface{}) error {
	params, isMonikerParams := i.(DtagParams)
	if !isMonikerParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if len(strings.TrimSpace(params.RegEx)) == 0 {
		return fmt.Errorf("empty dTag regEx param")
	}

	if params.MinDtagLen.IsNegative() || params.MinDtagLen.LT(DefaultMinDTagLength) {
		return fmt.Errorf("invalid minimum dTag length param: %s", params.MinDtagLen)
	}

	if params.MaxDtagLen.IsNegative() {
		return fmt.Errorf("invalid max dTag length param: %s", params.MaxDtagLen)
	}

	return nil
}

func ValidateBioParams(i interface{}) error {
	bioLen, isBioLenParams := i.(sdk.Int)
	if !isBioLenParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if bioLen.IsNegative() {
		return fmt.Errorf("invalid max bio length param: %s", bioLen)
	}

	return nil
}
