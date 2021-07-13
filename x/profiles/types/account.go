package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ghodss/yaml"
	"github.com/gogo/protobuf/proto"
)

var (
	_ authtypes.AccountI = (*Profile)(nil)
)

// NewProfile builds a new profile having the given DTag, creator and creation date
func NewProfile(
	dTag string, nickname, bio string, pictures Pictures, creationDate time.Time,
	account authtypes.AccountI,
) (*Profile, error) {
	// Make sure myAccount is a proto.Message, e.g. a BaseAccount etc.
	protoAccount, ok := account.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("the given account cannot be serialized using Protobuf")
	}

	myAccountAny, err := codectypes.NewAnyWithValue(protoAccount)
	if err != nil {
		return nil, err
	}

	return &Profile{
		DTag:         dTag,
		Nickname:     nickname,
		Bio:          bio,
		Pictures:     pictures,
		CreationDate: creationDate,
		Account:      myAccountAny,
	}, nil
}

// NewProfileFromAccount allows to build a new Profile instance from a provided DTag, and account and a creation time
func NewProfileFromAccount(dTag string, account authtypes.AccountI, creationTime time.Time) (*Profile, error) {
	return NewProfile(
		dTag,
		"",
		"",
		NewPictures("", ""),
		creationTime,
		account,
	)
}

// GetAccount returns the underlying account as an authtypes.AccountI instance
func (p *Profile) GetAccount() authtypes.AccountI {
	return p.Account.GetCachedValue().(authtypes.AccountI)
}

// setAccount sets the given account as the underlying account instance.
// This should be called after updating anything about the account (eg. after calling SetSequence).
func (p *Profile) setAccount(account authtypes.AccountI) error {
	accAny, err := codectypes.NewAnyWithValue(account)
	if err != nil {
		return err
	}

	p.Account = accAny
	return nil
}

// GetAddress implements authtypes.AccountI
func (p *Profile) GetAddress() sdk.AccAddress {
	return p.GetAccount().GetAddress()
}

// SetAddress implements authtypes.AccountI
func (p *Profile) SetAddress(addr sdk.AccAddress) error {
	acc := p.GetAccount()
	err := acc.SetAddress(addr)
	if err != nil {
		return err
	}

	return p.setAccount(acc)
}

// GetPubKey implements authtypes.AccountI
func (p *Profile) GetPubKey() cryptotypes.PubKey {
	return p.GetAccount().GetPubKey()
}

// SetPubKey implements authtypes.AccountI
func (p *Profile) SetPubKey(pubKey cryptotypes.PubKey) error {
	acc := p.GetAccount()
	err := acc.SetPubKey(pubKey)
	if err != nil {
		return err
	}

	return p.setAccount(acc)
}

// GetAccountNumber implements authtypes.AccountI
func (p *Profile) GetAccountNumber() uint64 {
	return p.GetAccount().GetAccountNumber()
}

// SetAccountNumber implements authtypes.AccountI
func (p *Profile) SetAccountNumber(accountNumber uint64) error {
	acc := p.GetAccount()
	err := acc.SetAccountNumber(accountNumber)
	if err != nil {
		return err
	}

	return p.setAccount(acc)
}

// GetSequence implements authtypes.AccountI
func (p *Profile) GetSequence() uint64 {
	return p.GetAccount().GetSequence()
}

// SetSequence implements authtypes.AccountI
func (p *Profile) SetSequence(sequence uint64) error {
	acc := p.GetAccount()
	err := acc.SetSequence(sequence)
	if err != nil {
		return err
	}

	return p.setAccount(acc)
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (p *Profile) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if p.Account != nil {
		var account authtypes.AccountI
		err := unpacker.UnpackAny(p.Account, &account)
		if err != nil {
			return err
		}
	}

	return nil
}

// Validate check the validity of the Profile
func (p *Profile) Validate() error {
	if strings.TrimSpace(p.DTag) == "" || p.DTag == DoNotModify {
		return fmt.Errorf("invalid profile DTag: %s", p.DTag)
	}

	if p.Nickname == DoNotModify {
		return fmt.Errorf("invalid profile nickname: %s", p.Nickname)
	}

	if p.Bio == DoNotModify {
		return fmt.Errorf("invalid profile bio: %s", p.Bio)
	}

	if p.Pictures.Profile == DoNotModify {
		return fmt.Errorf("invalid profile picture: %s", p.Pictures.Profile)
	}

	if p.Pictures.Cover == DoNotModify {
		return fmt.Errorf("invalid profile cover: %s", p.Pictures.Cover)
	}

	if len(p.GetAddress()) == 0 {
		return fmt.Errorf("invalid address: %s", p.GetAddress().String())
	}

	return p.Pictures.Validate()
}

// -------------------------------------------------------------------------------------------------------------------

type profilePretty struct {
	Address       sdk.AccAddress `json:"address" yaml:"address"`
	PubKey        string         `json:"public_key" yaml:"public_key"`
	AccountNumber uint64         `json:"account_number" yaml:"account_number"`
	Sequence      uint64         `json:"sequence" yaml:"sequence"`
	DTag          string         `json:"dtag" yaml:"dtag"`
	Nickname      string         `json:"nickname" yaml:"nickname"`
	Bio           string         `json:"bio" yaml:"bio"`
	Pictures      Pictures       `json:"pictures" yaml:"pictures"`
	CreationDate  time.Time      `json:"creation_date" yaml:"creation_date"`
}

// String implements authtypes.AccountI implements stringer
func (p *Profile) String() string {
	out, _ := p.MarshalYAML()
	return out.(string)
}

// MarshalYAML returns the YAML representation of a Profile.
func (p *Profile) MarshalYAML() (interface{}, error) {
	bs, err := yaml.Marshal(profilePretty{
		Address:       p.GetAddress(),
		PubKey:        p.GetPubKey().String(),
		AccountNumber: p.GetAccountNumber(),
		Sequence:      p.GetSequence(),
		DTag:          p.DTag,
		Nickname:      p.Nickname,
		Bio:           p.Bio,
		Pictures:      p.Pictures,
		CreationDate:  p.CreationDate,
	})

	if err != nil {
		return nil, err
	}

	return string(bs), nil
}

// MarshalJSON returns the JSON representation of a Profile.
func (p Profile) MarshalJSON() ([]byte, error) {
	var pubKey = ""
	if p.GetPubKey() != nil {
		pubKey = p.GetPubKey().String()
	}

	return json.Marshal(profilePretty{
		Address:       p.GetAddress(),
		PubKey:        pubKey,
		AccountNumber: p.GetAccountNumber(),
		Sequence:      p.GetSequence(),
		DTag:          p.DTag,
		Nickname:      p.Nickname,
		Bio:           p.Bio,
		Pictures:      p.Pictures,
		CreationDate:  p.CreationDate,
	})
}

// -------------------------------------------------------------------------------------------------------------------

// ProfileUpdate contains all the data that can be updated about a profile.
// When performing an update, if a field should not be edited then it must be set to types.DoNotModify
type ProfileUpdate struct {
	DTag     string
	Nickname string
	Bio      string
	Pictures Pictures
}

// NewProfileUpdate builds a new ProfileUpdate instance containing the given data
func NewProfileUpdate(dTag, nickname, bio string, pictures Pictures) *ProfileUpdate {
	return &ProfileUpdate{
		DTag:     dTag,
		Nickname: nickname,
		Bio:      bio,
		Pictures: pictures,
	}
}

// Update updates the fields of a given profile. An error is
// returned if the resulting profile contains invalid values.
func (p *Profile) Update(update *ProfileUpdate) (*Profile, error) {
	if update.DTag == DoNotModify {
		update.DTag = p.DTag
	}

	if update.Nickname == DoNotModify {
		update.Nickname = p.Nickname
	}

	if update.Bio == DoNotModify {
		update.Bio = p.Bio
	}

	if update.Pictures.Profile == DoNotModify {
		update.Pictures.Profile = p.Pictures.Profile
	}

	if update.Pictures.Cover == DoNotModify {
		update.Pictures.Cover = p.Pictures.Cover
	}

	newProfile, err := NewProfile(
		update.DTag,
		update.Nickname,
		update.Bio,
		update.Pictures,
		p.CreationDate,
		p.GetAccount(),
	)
	if err != nil {
		return nil, err
	}

	return newProfile, nil
}
