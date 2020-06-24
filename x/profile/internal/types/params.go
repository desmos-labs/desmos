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
	DefaultMinDtagLength    = sdk.NewInt(2)
	DefaultMaxDtagLength    = sdk.NewInt(30)
	DefaultMaxBioLength     = sdk.NewInt(1000)
)

// Parameters store keys
var (
	MonikerLenParamsKey = []byte("monikerLenParams")
	DtagLenParamsKey    = []byte("dtagLenParams")
	MaxBioLenParamsKey  = []byte("maxBioLenParams")
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramsModule.KeyTable {
	return paramsModule.NewKeyTable().RegisterParamSet(&Params{})
}

type Params struct {
	MonikerLengths MonikerLengths `json:"moniker_lengths" yaml:"moniker_lengths"`
	DtagLengths    DtagLengths    `json:"dtag_lengths" yaml:"dtag_lengths"`
	MaxBioLen      sdk.Int        `json:"max_bio_len" yaml:"max_bio_len"`
}

// NewParams creates a new ProfileParams obj
func NewParams(monikerLen MonikerLengths, dtagLen DtagLengths, maxBioLen sdk.Int) Params {
	return Params{
		MonikerLengths: monikerLen,
		DtagLengths:    dtagLen,
		MaxBioLen:      maxBioLen,
	}
}

// DefaultParams return default paramsModule
func DefaultParams() Params {
	return Params{
		MonikerLengths: DefaultMonikerLenParams(),
		DtagLengths:    DefaultDtagLenParams(),
		MaxBioLen:      DefaultMaxBioLength,
	}
}

func (params Params) String() string {
	out := "Profiles parameters:\n"
	out += fmt.Sprintf("%s\n%s\nBiography params lengths:\nMax accepted length: %s\n",
		params.MonikerLengths.String(),
		params.DtagLengths.String(),
		params.MaxBioLen,
	)

	return strings.TrimSpace(out)
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of profile module's parameters.
func (params *Params) ParamSetPairs() paramsModule.ParamSetPairs {
	return paramsModule.ParamSetPairs{
		paramsModule.NewParamSetPair(MonikerLenParamsKey, &params.MonikerLengths, ValidateMonikerLenParams),
		paramsModule.NewParamSetPair(DtagLenParamsKey, &params.DtagLengths, ValidateDtagLenParams),
		paramsModule.NewParamSetPair(MaxBioLenParamsKey, &params.MaxBioLen, ValidateBioLenParams),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateMonikerLenParams(params.MonikerLengths); err != nil {
		return err
	}

	if err := ValidateDtagLenParams(params.DtagLengths); err != nil {
		return err
	}

	return ValidateBioLenParams(params.MaxBioLen)
}

// MonikerLengths defines the paramsModule around moniker len
type MonikerLengths struct {
	MinMonikerLen sdk.Int `json:"min_moniker_len" yaml:"min_moniker_len"`
	MaxMonikerLen sdk.Int `json:"max_moniker_len" yaml:"max_moniker_len"`
}

// NewMonikerLenParams creates a new MonikerLengths obj
func NewMonikerLenParams(minLen, maxLen sdk.Int) MonikerLengths {
	return MonikerLengths{
		MinMonikerLen: minLen,
		MaxMonikerLen: maxLen,
	}
}

// DefaultMonikerLenParams return default moniker params
func DefaultMonikerLenParams() MonikerLengths {
	return NewMonikerLenParams(DefaultMinMonikerLength,
		DefaultMaxMonikerLength)
}

// String implements stringer interface
func (params MonikerLengths) String() string {
	out := "Moniker params lengths:\n"
	out += fmt.Sprintf("Min accepted length: %s\nMax accepted length: %s",
		params.MinMonikerLen,
		params.MaxMonikerLen,
	)

	return strings.TrimSpace(out)
}

func ValidateMonikerLenParams(i interface{}) error {
	params, isNameSurnParams := i.(MonikerLengths)

	if !isNameSurnParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.MinMonikerLen.IsNegative() || params.MinMonikerLen.LT(DefaultMinMonikerLength) {
		return fmt.Errorf("invalid minimum name/surname length param: %s", params.MinMonikerLen)
	}

	// TODO make sense to cap this? I've done this thinking "what's the sense of having names higher that 1000 chars?"
	if params.MaxMonikerLen.IsNegative() || params.MaxMonikerLen.GT(DefaultMaxMonikerLength) {
		return fmt.Errorf("invalid max name/surname length param: %s", params.MaxMonikerLen)
	}

	return nil
}

// DtagLengths defines the paramsModule around profiles' dtag
type DtagLengths struct {
	MinDtagLen sdk.Int `json:"min_dtag_len" yaml:"min_dtag_len"`
	MaxDtagLen sdk.Int `json:"max_dtag_len" yaml:"max_dtag_len"`
}

// NewDtagLenParams creates a new DtagLengths obj
func NewDtagLenParams(minLen, maxLen sdk.Int) DtagLengths {
	return DtagLengths{
		MinDtagLen: minLen,
		MaxDtagLen: maxLen,
	}
}

// DefaultDtagLenParams return default paramsModule
func DefaultDtagLenParams() DtagLengths {
	return NewDtagLenParams(DefaultMinDtagLength,
		DefaultMaxDtagLength)
}

// String implements stringer interface
func (params DtagLengths) String() string {
	out := "Dtag params lengths:\n"
	out += fmt.Sprintf("Min accepted length: %s\nMax accepted length: %s",
		params.MinDtagLen,
		params.MaxDtagLen,
	)

	return strings.TrimSpace(out)
}

func ValidateDtagLenParams(i interface{}) error {
	params, isMonikerParams := i.(DtagLengths)
	if !isMonikerParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.MinDtagLen.IsNegative() || params.MinDtagLen.LT(DefaultMinDtagLength) {
		return fmt.Errorf("invalid minimum moniker length param: %s", params.MinDtagLen)
	}

	if params.MaxDtagLen.IsNegative() {
		return fmt.Errorf("invalid max moniker length param: %s", params.MaxDtagLen)
	}

	return nil
}

func ValidateBioLenParams(i interface{}) error {
	bioLen, isBioLenParams := i.(sdk.Int)
	if !isBioLenParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if bioLen.IsNegative() {
		return fmt.Errorf("invalid max bio length param: %s", bioLen)
	}

	return nil
}
