package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramSubspace "github.com/cosmos/cosmos-sdk/x/params/subspace"
)

const (
	// default paramspace for params keeper
	DefaultParamspace = ModuleName
)

// Default profile params
var (
	DefaultMinNameSurnameLength = sdk.NewInt(2)
	DefaultMaxNameSurnameLength = sdk.NewInt(1000) //longest name on earth count 954 chars
	DefaultMinMonikerLength     = sdk.NewInt(2)
	DefaultMaxMonikerLength     = sdk.NewInt(30)
	DefaultMaxBioLength         = sdk.NewInt(1000)
)

// Parameters store keys
var (
	ParamStoreKeyNameSurnameLen = []byte("nameSurnameLenParams")
	ParamStoreKeyMonikerLen     = []byte("monikerLenParams")
	ParamStoreKeyMaxBioLen      = []byte("maxBioLenParams")
)

// ParamKeyTable - Key declaration for params
func ParamKeyTable() paramSubspace.KeyTable {
	return paramSubspace.NewKeyTable(
		paramSubspace.NewParamSetPair(ParamStoreKeyNameSurnameLen, NameSurnameLenParams{}, validateNameSurnameLenParams),
		paramSubspace.NewParamSetPair(ParamStoreKeyMonikerLen, MonikerLenParams{}, validateMonikerLenParams),
		paramSubspace.NewParamSetPair(ParamStoreKeyMaxBioLen, BioLenParams{}, validateBioLenParams),
	)
}

// NameSurnameLenParams defines the params around names and surnames len
type NameSurnameLenParams struct {
	MinNameSurnameLen sdk.Int `json:"min_name_surname_len" yaml:"min_name_surname_len"`
	MaxNameSurnameLen sdk.Int `json:"max_name_surname_len" yaml:"max_name_surname_len"`
}

// NewNameSurnameLenParams creates a new NameSurnameLenParams obj
func NewNameSurnameLenParams(minLen, maxLen sdk.Int) NameSurnameLenParams {
	return NameSurnameLenParams{
		MinNameSurnameLen: minLen,
		MaxNameSurnameLen: maxLen,
	}
}

// DefaultNameSurnameLenParams return default params
func DefaultNameSurnameLenParams() NameSurnameLenParams {
	return NewNameSurnameLenParams(
		DefaultMinNameSurnameLength,
		DefaultMaxNameSurnameLength,
	)
}

// String implements stringer interface
func (params NameSurnameLenParams) String() string {
	out, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func validateNameSurnameLenParams(i interface{}) error {
	params, isNameSurnParams := i.(NameSurnameLenParams)

	if !isNameSurnParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.MinNameSurnameLen.IsNegative() || params.MinNameSurnameLen.LT(DefaultMinNameSurnameLength) {
		return fmt.Errorf("invalid minimum name/surname length param: %s", params.MinNameSurnameLen)
	}

	// TODO make sense to cap this? I've done this thinking "what's the sense of having names higher that 1000 chars?"
	if params.MaxNameSurnameLen.IsNegative() || params.MaxNameSurnameLen.GT(DefaultMaxNameSurnameLength) {
		return fmt.Errorf("invalid max name/surname length param: %s", params.MaxNameSurnameLen)
	}

	return nil
}

// MonikerLenParams defines the params around profiles' monikers
type MonikerLenParams struct {
	MinMonikerLen sdk.Int `json:"min_moniker_len" yaml:"min_moniker_len"`
	MaxMonikerLen sdk.Int `json:"max_moniker_len" yaml:"min_moniker_len"`
}

// NewMonikerLenParams creates a new MonikerLenParams obj
func NewMonikerLenParams(minLen, maxLen sdk.Int) MonikerLenParams {
	return MonikerLenParams{
		MinMonikerLen: minLen,
		MaxMonikerLen: maxLen,
	}
}

// DefaultMonikerLenParams return default params
func DefaultMonikerLenParams() MonikerLenParams {
	return MonikerLenParams{
		MinMonikerLen: DefaultMinMonikerLength,
		MaxMonikerLen: DefaultMaxMonikerLength,
	}
}

// String implements stringer interface
func (params MonikerLenParams) String() string {
	out, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func validateMonikerLenParams(i interface{}) error {
	params, isMonikerParams := i.(MonikerLenParams)
	if !isMonikerParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.MinMonikerLen.IsNegative() || params.MinMonikerLen.LT(DefaultMinMonikerLength) {
		return fmt.Errorf("invalid minimum moniker length param: %s", params.MinMonikerLen)
	}

	if params.MaxMonikerLen.IsNegative() {
		return fmt.Errorf("invalid max moniker length param: %s", params.MaxMonikerLen)
	}

	return nil
}

// BioLenParams defines the params around profiles' biography
type BioLenParams struct {
	MaxBioLen sdk.Int `json:"max_bio_len" yaml:"max_moniker_len"`
}

// NewBioLenParams creates a new BioLenParams obj
func NewBioLenParams(maxLen sdk.Int) BioLenParams {
	return BioLenParams{
		MaxBioLen: maxLen,
	}
}

// DefaultBioLenParams returns default params
func DefaultBioLenParams() BioLenParams {
	return BioLenParams{
		MaxBioLen: DefaultMaxBioLength,
	}
}

// String implements stringer interface
func (params BioLenParams) String() string {
	out, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func validateBioLenParams(i interface{}) error {
	params, isBioLenParams := i.(BioLenParams)
	if !isBioLenParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.MaxBioLen.IsNegative() {
		return fmt.Errorf("invalid max bio length param: %s", params.MaxBioLen)
	}

	return nil
}
