package types

import (
	"fmt"
	"strings"
	"time"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// DefaultParamsSpace represents the default paramspace for the Params keeper
	DefaultParamsSpace           = ModuleName
	FourteenDaysCorrectionFactor = time.Hour * 24 * 14 // This value is the equivalent of 14 days in minutes
)

// Default profile paramsModule
var (
	DefaultMinNicknameLength        = math.NewInt(2)
	DefaultMaxNicknameLength        = math.NewInt(1000) // Longest name on earth count 954 chars
	DefaultRegEx                    = `^[A-Za-z0-9_]+$`
	DefaultMinDTagLength            = math.NewInt(3)
	DefaultMaxDTagLength            = math.NewInt(30)
	DefaultMaxBioLength             = math.NewInt(1000)
	DefaultAppLinksValidityDuration = time.Hour * 24 * 365 // 1 year
)

// ___________________________________________________________________________________________________________________

// NewParams creates a new ProfileParams obj
func NewParams(nickname NicknameParams, dTag DTagParams, bio BioParams, oracle OracleParams, appLinks AppLinksParams) Params {
	return Params{
		Nickname: nickname,
		DTag:     dTag,
		Bio:      bio,
		Oracle:   oracle,
		AppLinks: appLinks,
	}
}

// DefaultParams return default paramsModule
func DefaultParams() Params {
	return Params{
		Nickname: DefaultNicknameParams(),
		DTag:     DefaultDTagParams(),
		Bio:      DefaultBioParams(),
		Oracle:   DefaultOracleParams(),
		AppLinks: DefaultAppLinksParams(),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateNicknameParams(params.Nickname); err != nil {
		return err
	}

	if err := ValidateDTagParams(params.DTag); err != nil {
		return err
	}

	if err := ValidateBioParams(params.Bio); err != nil {
		return err
	}

	if err := ValidateOracleParams(params.Oracle); err != nil {
		return err
	}

	return ValidateAppLinksParams(params.AppLinks)
}

// ___________________________________________________________________________________________________________________

// NewNicknameParams creates a new NicknameParams obj
func NewNicknameParams(minLen, maxLen math.Int) NicknameParams {
	return NicknameParams{
		MinLength: minLen,
		MaxLength: maxLen,
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

	minLength := params.MinLength
	if minLength.IsNil() || minLength.LT(DefaultMinNicknameLength) {
		return fmt.Errorf("invalid minimum nickname length param: %s", minLength)
	}

	// TODO make sense to cap this? I've done this thinking "what's the sense of having names higher that 1000 chars?"
	maxLength := params.MaxLength
	if maxLength.IsNil() || maxLength.IsNegative() || maxLength.GT(DefaultMaxNicknameLength) {
		return fmt.Errorf("invalid max nickname length param: %s", maxLength)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// NewDTagParams creates a new DTagParams obj
func NewDTagParams(regEx string, minLen, maxLen math.Int) DTagParams {
	return DTagParams{
		RegEx:     regEx,
		MinLength: minLen,
		MaxLength: maxLen,
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

	if params.MinLength.IsNegative() || params.MinLength.LT(DefaultMinDTagLength) {
		return fmt.Errorf("invalid minimum dTag length param: %s", params.MinLength)
	}

	if params.MaxLength.IsNegative() {
		return fmt.Errorf("invalid max dTag length param: %s", params.MaxLength)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// NewBioParams creates a new BioParams obj
func NewBioParams(maxLength math.Int) BioParams {
	return BioParams{
		MaxLength: maxLength,
	}
}

// DefaultBioParams returns default params module
func DefaultBioParams() BioParams {
	return NewBioParams(DefaultMaxBioLength)
}

func ValidateBioParams(i interface{}) error {
	bioParams, isBioParams := i.(BioParams)
	if !isBioParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if bioParams.MaxLength.IsNegative() {
		return fmt.Errorf("invalid max bio length param: %s", bioParams.MaxLength)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// NewOracleParams creates a new Oracle Params instance
func NewOracleParams(
	scriptID uint64,
	askCount,
	minCount,
	prepareGas,
	executeGas uint64,
	feeAmount ...sdk.Coin,
) OracleParams {
	return OracleParams{
		ScriptID:   scriptID,
		AskCount:   askCount,
		MinCount:   minCount,
		PrepareGas: prepareGas,
		ExecuteGas: executeGas,
		FeeAmount:  feeAmount,
	}
}

// DefaultOracleParams returns the default instance of OracleParams
func DefaultOracleParams() OracleParams {
	return NewOracleParams(
		0,
		1,
		1,
		50_000,
		200_000,
		sdk.NewCoin("band", math.NewInt(10)),
	)
}

// ValidateOracleParams returns an error if interface does not represent a valid OracleParams instance
func ValidateOracleParams(i interface{}) error {
	params, isOracleParams := i.(OracleParams)
	if !isOracleParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.AskCount < params.MinCount {
		return fmt.Errorf("invalid ask count: %d, min count: %d", params.AskCount, params.MinCount)
	}

	if params.MinCount <= 0 {
		return fmt.Errorf("invalid min count: %d", params.MinCount)
	}

	if params.PrepareGas <= 0 {
		return fmt.Errorf("invalid prepare gas: %d", params.PrepareGas)
	}

	if params.ExecuteGas <= 0 {
		return fmt.Errorf("invalid execute gas: %d", params.ExecuteGas)
	}

	err := params.FeeAmount.Validate()
	if err != nil {
		return err
	}

	return nil
}

// ___________________________________________________________________________________________________________________

func NewAppLinksParams(validityDuration time.Duration) AppLinksParams {
	return AppLinksParams{
		ValidityDuration: validityDuration,
	}
}

func DefaultAppLinksParams() AppLinksParams {
	return NewAppLinksParams(DefaultAppLinksValidityDuration)
}

func ValidateAppLinksParams(i interface{}) error {
	params, isAppLinksParams := i.(AppLinksParams)
	if !isAppLinksParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.ValidityDuration < FourteenDaysCorrectionFactor {
		return fmt.Errorf("validity duration must be not less than 14 days: %s", params.ValidityDuration)
	}

	return nil
}
