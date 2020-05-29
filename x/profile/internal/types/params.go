package types

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramSubspace "github.com/cosmos/cosmos-sdk/x/params/subspace"
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
	ParamStoreKeyNameSurnameLenRange = []byte("name_surname_len_range")
	ParamStoreKeyMonikerLenRange     = []byte("moniker_len_range")
	ParamStoreKeyMaxBioLen           = []byte("max_bio_len")
)

// ParamKeyTable - Key declaration for params
func ParamKeyTable() paramSubspace.KeyTable {
	return paramSubspace.NewKeyTable(
		paramSubspace.NewParamSetPair(ParamStoreKeyNameSurnameLenRange, NameSurnameLenParams{}, validateNameSurnameLenParams),
		paramSubspace.NewParamSetPair(ParamStoreKeyMonikerLenRange, MonikerLenParams{}, validateMonikerLenParams),
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
	out, _ := json.Marshal(params)
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

// BioLenParams defines the params around profiles' biography
type BioLenParams struct {
	MaxBioLen sdk.Int `json:"max_bio_len" yaml:"max_moniker_len"`
}
