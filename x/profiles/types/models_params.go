package types

import (
	"fmt"
	"strings"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// DefaultParamsSpace represents the default paramspace for the Params keeper
	DefaultParamsSpace = ModuleName
)

// Default profile paramsModule
var (
	DefaultMinNicknameLength = sdk.NewInt(2)
	DefaultMaxNicknameLength = sdk.NewInt(1000) //longest name on earth count 954 chars
	DefaultRegEx             = `^[A-Za-z0-9_]+$`
	DefaultMinDTagLength     = sdk.NewInt(3)
	DefaultMaxDTagLength     = sdk.NewInt(30)
	DefaultMaxBioLength      = sdk.NewInt(1000)

	// OracleScriptID represents the oracle script to be called on Band Protocol
	DefaultOracleScriptID int64 = 32

	DefaultOracleAskCount uint64 = 10
	DefaultOracleMinCount uint64 = 6

	DefaultOraclePrepareGas uint64 = 50_000
	DefaultOracleExecuteGas uint64 = 200_000

	DefaultOracleFeePayer = "desmos-ibc-profiles"
	DefaultOracleFeeCoins = sdk.NewCoins(sdk.NewCoin("band", sdk.NewInt(10)))
)

// Parameters store keys
var (
	NicknameLenParamsKey = []byte("NicknameParams")
	DTagLenParamsKey     = []byte("DTagParams")
	MaxBioLenParamsKey   = []byte("MaxBioLen")
	OracleParamsKey      = []byte("OracleParams")
)

// ___________________________________________________________________________________________________________________

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().
		RegisterParamSet(&Params{})
}

// NewParams creates a new ProfileParams obj
func NewParams(nicknameParams NicknameParams, dTagParams DTagParams, maxBioLen sdk.Int, oracleParams OracleParams) Params {
	return Params{
		NicknameParams: nicknameParams,
		DTagParams:     dTagParams,
		MaxBioLength:   maxBioLen,
		Oracle:         oracleParams,
	}
}

// DefaultParams return default paramsModule
func DefaultParams() Params {
	return Params{
		NicknameParams: DefaultNicknameParams(),
		DTagParams:     DefaultDTagParams(),
		MaxBioLength:   DefaultMaxBioLength,
		Oracle:         DefaultOracleParams(),
	}
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of profile module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(NicknameLenParamsKey, &params.NicknameParams, ValidateNicknameParams),
		paramstypes.NewParamSetPair(DTagLenParamsKey, &params.DTagParams, ValidateDTagParams),
		paramstypes.NewParamSetPair(MaxBioLenParamsKey, &params.MaxBioLength, ValidateBioParams),
		paramstypes.NewParamSetPair(OracleParamsKey, &params.Oracle, ValidateOracleParams),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateNicknameParams(params.NicknameParams); err != nil {
		return err
	}

	if err := ValidateDTagParams(params.DTagParams); err != nil {
		return err
	}

	return ValidateBioParams(params.MaxBioLength)
}

// ___________________________________________________________________________________________________________________

// NewNicknameParams creates a new NicknameParams obj
func NewNicknameParams(minLen, maxLen sdk.Int) NicknameParams {
	return NicknameParams{
		MinNicknameLength: minLen,
		MaxNicknameLength: maxLen,
	}
}

// DefaultNicknameParams return default nickname params
func DefaultNicknameParams() NicknameParams {
	return NewNicknameParams(
		DefaultMinNicknameLength,
		DefaultMaxNicknameLength,
	)
}

func ValidateNicknameParams(i interface{}) error {
	params, areNicknameParams := i.(NicknameParams)
	if !areNicknameParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	minLength := params.MinNicknameLength
	if minLength.IsNil() || minLength.LT(DefaultMinNicknameLength) {
		return fmt.Errorf("invalid minimum nickname length param: %s", minLength)
	}

	// TODO make sense to cap this? I've done this thinking "what's the sense of having names higher that 1000 chars?"
	maxLength := params.MaxNicknameLength
	if maxLength.IsNil() || maxLength.IsNegative() || maxLength.GT(DefaultMaxNicknameLength) {
		return fmt.Errorf("invalid max nickname length param: %s", maxLength)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// NewDTagParams creates a new DTagParams obj
func NewDTagParams(regEx string, minLen, maxLen sdk.Int) DTagParams {
	return DTagParams{
		RegEx:         regEx,
		MinDTagLength: minLen,
		MaxDTagLength: maxLen,
	}
}

// DefaultDTagParams return default paramsModule
func DefaultDTagParams() DTagParams {
	return NewDTagParams(
		DefaultRegEx,
		DefaultMinDTagLength,
		DefaultMaxDTagLength,
	)
}

func ValidateDTagParams(i interface{}) error {
	params, isDtagParams := i.(DTagParams)
	if !isDtagParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if len(strings.TrimSpace(params.RegEx)) == 0 {
		return fmt.Errorf("empty dTag regEx param")
	}

	if params.MinDTagLength.IsNegative() || params.MinDTagLength.LT(DefaultMinDTagLength) {
		return fmt.Errorf("invalid minimum dTag length param: %s", params.MinDTagLength)
	}

	if params.MaxDTagLength.IsNegative() {
		return fmt.Errorf("invalid max dTag length param: %s", params.MaxDTagLength)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

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

// ___________________________________________________________________________________________________________________

// NewOracleParams creates a new Oracle Params obj
func NewOracleParams(
	scriptID int64,
	askCount,
	minCount,
	prepareGas,
	executeGas uint64,
	feePayer string,
	feeAmount ...sdk.Coin,
) OracleParams {
	return OracleParams{
		ScriptID:   scriptID,
		AskCount:   askCount,
		MinCount:   minCount,
		PrepareGas: prepareGas,
		ExecuteGas: executeGas,
		FeePayer:   feePayer,
		FeeAmount:  feeAmount,
	}
}

func DefaultOracleParams() OracleParams {
	return NewOracleParams(
		DefaultOracleScriptID,
		DefaultOracleAskCount,
		DefaultOracleMinCount,
		DefaultOraclePrepareGas,
		DefaultOracleExecuteGas,
		DefaultOracleFeePayer,
		DefaultOracleFeeCoins...,
	)
}

func ValidateOracleParams(i interface{}) error {
	params, isOracleParams := i.(OracleParams)
	if !isOracleParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.AskCount < params.MinCount {
		return fmt.Errorf("invalid ask count: %d, min count: %d", params.AskCount, params.MinCount)
	}

	if params.MinCount < 0 {
		return fmt.Errorf("invalid min count: %d", params.MinCount)
	}

	if params.PrepareGas <= 0 {
		return fmt.Errorf("invalid prepare gas: %d", params.PrepareGas)
	}

	if params.ExecuteGas <= 0 {
		return fmt.Errorf("invalid execute gas: %d", params.ExecuteGas)
	}

	if strings.TrimSpace(params.FeePayer) == "" {
		return fmt.Errorf("invalid fee payer: %s", params.FeePayer)
	}

	err := params.FeeAmount.Validate()
	if err != nil {
		return err
	}

	return nil
}
