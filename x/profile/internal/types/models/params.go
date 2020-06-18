package models

import (
	"encoding/json"
	"fmt"
	paramsModule "github.com/cosmos/cosmos-sdk/x/params/subspace"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// default paramspace for paramsModule keeper
	DefaultParamspace = ModuleName
)

// Default profile paramsModule
var (
	DefaultMinNameSurnameLength = sdk.NewInt(2)
	DefaultMaxNameSurnameLength = sdk.NewInt(1000) //longest name on earth count 954 chars
	DefaultMinMonikerLength     = sdk.NewInt(2)
	DefaultMaxMonikerLength     = sdk.NewInt(30)
	DefaultMaxBioLength         = sdk.NewInt(1000)
)

// Parameters store keys
var (
	NameSurnameLenParamsKey = []byte("nameSurnameLenParams")
	MonikerLenParamsKey     = []byte("monikerLenParams")
	BioLenParamsKey         = []byte("bioLenParams")

	ProfileParamsKey = []byte("profileParams")
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramsModule.KeyTable {
	return paramsModule.NewKeyTable().RegisterParamSet(&Params{})
}

type Params struct {
	NameSurnameLengths NameSurnameLengths `json:"name_surname_lengths" yaml:"name_surname_lengths"`
	MonikerLengths     MonikerLengths     `json:"moniker_lengths" yaml:"moniker_lengths"`
	BiographyLengths   BiographyLengths   `json:"biography_lengths" yaml:"biography_lengths"`
}

// NewParams creates a new ProfileParams obj
func NewParams(nsLen NameSurnameLengths, monikerLen MonikerLengths, bioLen BiographyLengths) Params {
	return Params{
		NameSurnameLengths: nsLen,
		MonikerLengths:     monikerLen,
		BiographyLengths:   bioLen,
	}
}

// DefaultParams return default paramsModule
func DefaultParams() Params {
	return Params{
		NameSurnameLengths: DefaultNameSurnameLenParams(),
		MonikerLengths:     DefaultMonikerLenParams(),
		BiographyLengths:   DefaultBioLenParams(),
	}
}

// WithNameSurnameParams replace the given non nil nsParams with the existent one
func (params Params) WithNameSurnameParams(nsParams NameSurnameLengths) Params {
	if nsParams.MinNameSurnameLen == nil {
		params.NameSurnameLengths.MaxNameSurnameLen = nsParams.MaxNameSurnameLen
	} else if nsParams.MaxNameSurnameLen == nil {
		params.NameSurnameLengths.MinNameSurnameLen = nsParams.MinNameSurnameLen
	} else {
		params.NameSurnameLengths = nsParams
	}
	return params
}

// WithMonikerParams replace the given non nil moniker with the existent one
func (params Params) WithMonikerParams(mParams MonikerLengths) Params {
	if mParams.MinMonikerLen == nil {
		params.MonikerLengths.MaxMonikerLen = mParams.MaxMonikerLen
	} else if mParams.MaxMonikerLen == nil {
		params.MonikerLengths.MinMonikerLen = mParams.MinMonikerLen
	} else {
		params.MonikerLengths = mParams
	}
	return params
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of profile module's parameters.
func (params *Params) ParamSetPairs() paramsModule.ParamSetPairs {
	return paramsModule.ParamSetPairs{
		paramsModule.NewParamSetPair(NameSurnameLenParamsKey, &params.NameSurnameLengths, ValidateNameSurnameLenParams),
		paramsModule.NewParamSetPair(MonikerLenParamsKey, &params.MonikerLengths, ValidateMonikerLenParams),
		paramsModule.NewParamSetPair(BioLenParamsKey, &params.BiographyLengths, ValidateBioLenParams),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateNameSurnameLenParams(params.NameSurnameLengths); err != nil {
		return err
	}

	if err := ValidateMonikerLenParams(params.MonikerLengths); err != nil {
		return err
	}

	return ValidateBioLenParams(params.BiographyLengths)
}

// NameSurnameLengths defines the paramsModule around names and surnames len
type NameSurnameLengths struct {
	MinNameSurnameLen *sdk.Int `json:"min_name_surname_len" yaml:"min_name_surname_len"`
	MaxNameSurnameLen *sdk.Int `json:"max_name_surname_len" yaml:"max_name_surname_len"`
}

// NewNameSurnameLenParams creates a new NameSurnameLengths obj
func NewNameSurnameLenParams(minLen, maxLen *sdk.Int) NameSurnameLengths {
	return NameSurnameLengths{
		MinNameSurnameLen: minLen,
		MaxNameSurnameLen: maxLen,
	}
}

// DefaultNameSurnameLenParams return default paramsModule
func DefaultNameSurnameLenParams() NameSurnameLengths {
	return NewNameSurnameLenParams(&DefaultMinNameSurnameLength,
		&DefaultMaxNameSurnameLength)
}

// String implements stringer interface
func (params NameSurnameLengths) String() string {
	out, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func ValidateNameSurnameLenParams(i interface{}) error {
	params, isNameSurnParams := i.(NameSurnameLengths)

	if !isNameSurnParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.MinNameSurnameLen != nil && params.MinNameSurnameLen.IsNegative() || params.MinNameSurnameLen.LT(DefaultMinNameSurnameLength) {
		return fmt.Errorf("invalid minimum name/surname length param: %s", params.MinNameSurnameLen)
	}

	// TODO make sense to cap this? I've done this thinking "what's the sense of having names higher that 1000 chars?"
	if params.MaxNameSurnameLen != nil && params.MaxNameSurnameLen.IsNegative() || params.MaxNameSurnameLen.GT(DefaultMaxNameSurnameLength) {
		return fmt.Errorf("invalid max name/surname length param: %s", params.MaxNameSurnameLen)
	}

	return nil
}

// MonikerLengths defines the paramsModule around profiles' monikers
type MonikerLengths struct {
	MinMonikerLen *sdk.Int `json:"min_moniker_len" yaml:"min_moniker_len"`
	MaxMonikerLen *sdk.Int `json:"max_moniker_len" yaml:"min_moniker_len"`
}

// NewMonikerLenParams creates a new MonikerLengths obj
func NewMonikerLenParams(minLen, maxLen *sdk.Int) MonikerLengths {
	return MonikerLengths{
		MinMonikerLen: minLen,
		MaxMonikerLen: maxLen,
	}
}

// DefaultMonikerLenParams return default paramsModule
func DefaultMonikerLenParams() MonikerLengths {
	return NewMonikerLenParams(&DefaultMinMonikerLength,
		&DefaultMaxMonikerLength)
}

// String implements stringer interface
func (params MonikerLengths) String() string {
	out, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func ValidateMonikerLenParams(i interface{}) error {
	params, isMonikerParams := i.(MonikerLengths)
	if !isMonikerParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.MinMonikerLen != nil && params.MinMonikerLen.IsNegative() || params.MinMonikerLen.LT(DefaultMinMonikerLength) {
		return fmt.Errorf("invalid minimum moniker length param: %s", params.MinMonikerLen)
	}

	if params.MaxMonikerLen != nil && params.MaxMonikerLen.IsNegative() {
		return fmt.Errorf("invalid max moniker length param: %s", params.MaxMonikerLen)
	}

	return nil
}

// BiographyLengths defines the paramsModule around profiles' biography
type BiographyLengths struct {
	MaxBioLen sdk.Int `json:"max_bio_len" yaml:"max_moniker_len"`
}

// NewBioLenParams creates a new BiographyLengths obj
func NewBioLenParams(maxLen sdk.Int) BiographyLengths {
	return BiographyLengths{
		MaxBioLen: maxLen,
	}
}

// DefaultBioLenParams returns default paramsModule
func DefaultBioLenParams() BiographyLengths {
	return NewBioLenParams(DefaultMaxBioLength)
}

// String implements stringer interface
func (params BiographyLengths) String() string {
	out, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func ValidateBioLenParams(i interface{}) error {
	params, isBioLenParams := i.(BiographyLengths)
	if !isBioLenParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.MaxBioLen.IsNegative() {
		return fmt.Errorf("invalid max bio length param: %s", params.MaxBioLen)
	}

	return nil
}
