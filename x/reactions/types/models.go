package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rivo/uniseg"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

// ParseReactionID parses the given value as a reaction id, returning an error if it's invalid
func ParseReactionID(value string) (uint32, error) {
	if value == "" {
		return 0, nil
	}

	reactionID, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid reaction id: %s", err)
	}
	return uint32(reactionID), nil
}

// NewReaction returns a new Reaction instance
func NewReaction(subspaceID uint64, postID uint64, id uint32, value ReactionValue, author string) Reaction {
	valueAny, err := codectypes.NewAnyWithValue(value)
	if err != nil {
		panic("failed to pack value to any type")
	}

	return Reaction{
		SubspaceID: subspaceID,
		ID:         id,
		PostID:     postID,
		Value:      valueAny,
		Author:     author,
	}
}

// Validate implements fmt.Validator
func (r Reaction) Validate() error {
	if r.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", r.SubspaceID)
	}

	if r.ID == 0 {
		return fmt.Errorf("invalid id: %d", r.ID)
	}

	if r.PostID == 0 {
		return fmt.Errorf("invalid post id: %d", r.PostID)
	}

	err := r.Value.GetCachedValue().(ReactionValue).Validate()
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(r.Author)
	if err != nil {
		return fmt.Errorf("invalid author address: %s", err)
	}

	return nil
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (r *Reaction) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var target ReactionValue
	return unpacker.UnpackAny(r.Value, &target)
}

// --------------------------------------------------------------------------------------------------------------------

// ReactionValue represents a generic reaction value
type ReactionValue interface {
	proto.Message

	Equal(other interface{}) bool
	isReactionValue()
	Validate() error
}

// --------------------------------------------------------------------------------------------------------------------

var _ ReactionValue = &RegisteredReactionValue{}

// NewRegisteredReactionValue returns a new RegisteredReactionValue instance
func NewRegisteredReactionValue(registeredReactionID uint32) *RegisteredReactionValue {
	return &RegisteredReactionValue{
		RegisteredReactionID: registeredReactionID,
	}
}

// isReactionValue implements ReactionValue
func (v *RegisteredReactionValue) isReactionValue() {}

// Validate implements ReactionValue
func (v *RegisteredReactionValue) Validate() error {
	if v.RegisteredReactionID == 0 {
		return fmt.Errorf("invalid reaction id: %d", v.RegisteredReactionID)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

var _ ReactionValue = &FreeTextValue{}

// NewFreeTextValue returns a new FreeTextValue instance
func NewFreeTextValue(text string) *FreeTextValue {
	return &FreeTextValue{
		Text: text,
	}
}

// isReactionValue implements ReactionValue
func (v *FreeTextValue) isReactionValue() {}

// GetLength returns the length of the reaction value
func (v *FreeTextValue) GetLength() int {
	// Counting graphemes instead of runes of bytes can provide a more accurate length of the text.
	// This will also ensure that emojis are counted as a single character, which will grant a more consistent
	// user experience with clients as well.
	// Example: 🏳️‍🌈 (rainbow flag emoji) is 1 grapheme, 4 runes, and 14 bytes.
	return uniseg.GraphemeClusterCount(v.Text)

}

// Validate implements ReactionValue
func (v *FreeTextValue) Validate() error {
	if strings.TrimSpace(v.Text) == "" {
		return fmt.Errorf("invalid text: %s", v.Text)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// ParseRegisteredReactionID parses the given value as a registered reaction id, returning an error if it's invalid
func ParseRegisteredReactionID(value string) (uint32, error) {
	if value == "" {
		return 0, nil
	}

	reactionID, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid registered reaction id: %s", err)
	}
	return uint32(reactionID), nil
}

// NewRegisteredReaction returns a new RegisteredReaction instance
func NewRegisteredReaction(subspaceID uint64, id uint32, shorthandCode string, displayValue string) RegisteredReaction {
	return RegisteredReaction{
		SubspaceID:    subspaceID,
		ID:            id,
		ShorthandCode: shorthandCode,
		DisplayValue:  displayValue,
	}
}

// Validate implements fmt.Validator
func (r RegisteredReaction) Validate() error {
	if r.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", r.SubspaceID)
	}

	if r.ID == 0 {
		return fmt.Errorf("invalid id: %d", r.ID)
	}

	if strings.TrimSpace(r.ShorthandCode) == "" {
		return fmt.Errorf("invalid shorthand code: %s", r.ShorthandCode)
	}

	if strings.TrimSpace(r.DisplayValue) == "" {
		return fmt.Errorf("invalid display value: %s", r.DisplayValue)
	}

	return nil
}

// Update updates the fields of a given registered reaction without validating it.
// Before storing the updated reaction, a validation with Validate() should
// be performed.
func (r RegisteredReaction) Update(update RegisteredReactionUpdate) RegisteredReaction {
	if update.ShorthandCode == DoNotModify {
		update.ShorthandCode = r.ShorthandCode
	}

	if update.DisplayValue == DoNotModify {
		update.DisplayValue = r.DisplayValue
	}

	return NewRegisteredReaction(
		r.SubspaceID,
		r.ID,
		update.ShorthandCode,
		update.DisplayValue,
	)
}

// RegisteredReactionUpdate contains all the data that can be updated about a registered reaction.
// When performing an update, if a field should not be edited then it must be set to types.DoNotModify
type RegisteredReactionUpdate struct {
	ShorthandCode string
	DisplayValue  string
}

// NewRegisteredReactionUpdate builds a new RegisteredReactionUpdate instance containing the given data
func NewRegisteredReactionUpdate(shorthandCode, displayValue string) RegisteredReactionUpdate {
	return RegisteredReactionUpdate{
		ShorthandCode: shorthandCode,
		DisplayValue:  displayValue,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// NewSubspaceReactionsParams returns a new SubspaceReactionsParams instance
func NewSubspaceReactionsParams(
	subspaceID uint64,
	registeredReactionParams RegisteredReactionValueParams,
	freeTextParams FreeTextValueParams,
) SubspaceReactionsParams {
	return SubspaceReactionsParams{
		SubspaceID:         subspaceID,
		RegisteredReaction: registeredReactionParams,
		FreeText:           freeTextParams,
	}
}

// DefaultReactionsParams returns the default params for the given subspace
func DefaultReactionsParams(subspaceID uint64) SubspaceReactionsParams {
	return NewSubspaceReactionsParams(
		subspaceID,
		NewRegisteredReactionValueParams(true),
		NewFreeTextValueParams(true, 5, ""),
	)
}

// Validate implements fmt.Validator
func (p SubspaceReactionsParams) Validate() error {
	if p.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", p.SubspaceID)
	}

	err := p.FreeText.Validate()
	if err != nil {
		return fmt.Errorf("invalid free text value params: %s", err)
	}

	return nil
}

// NewRegisteredReactionValueParams returns a new RegisteredReactionValueParams instance
func NewRegisteredReactionValueParams(enabled bool) RegisteredReactionValueParams {
	return RegisteredReactionValueParams{
		Enabled: enabled,
	}
}

// NewFreeTextValueParams returns a new FreeTextValueParams instance
func NewFreeTextValueParams(enabled bool, maxLength uint32, regEx string) FreeTextValueParams {
	return FreeTextValueParams{
		Enabled:   enabled,
		MaxLength: maxLength,
		RegEx:     regEx,
	}
}

// Validate implements fmt.Validator
func (p FreeTextValueParams) Validate() error {
	if p.MaxLength == 0 {
		return fmt.Errorf("invalid max length: %d", p.MaxLength)
	}

	if p.RegEx != "" {
		_, err := regexp.Compile(p.RegEx)
		if err != nil {
			return fmt.Errorf("invalid regex: %s", err)
		}
	}

	return nil
}
