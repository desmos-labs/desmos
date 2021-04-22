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
	dTag string, moniker, bio string, pictures Pictures, creationDate time.Time, account authtypes.AccountI,
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
		Dtag:         dTag,
		Moniker:      moniker,
		Bio:          bio,
		Pictures:     pictures,
		CreationDate: creationDate,
		Account:      myAccountAny,
	}, nil
}

// GetAccount returns the underlying account as an authtypes.AccountI instance
func (p *Profile) GetAccount() authtypes.AccountI {
	return p.Account.GetCachedValue().(authtypes.AccountI)
}

// GetAddress implements authtypes.AccountI
func (p *Profile) GetAddress() sdk.AccAddress {
	return p.GetAccount().GetAddress()
}

// SetAddress implements authtypes.AccountI
func (p *Profile) SetAddress(addr sdk.AccAddress) error {
	return p.GetAccount().SetAddress(addr)
}

// GetPubKey implements authtypes.AccountI
func (p *Profile) GetPubKey() cryptotypes.PubKey {
	return p.GetAccount().GetPubKey()
}

// SetPubKey implements authtypes.AccountI
func (p *Profile) SetPubKey(pubKey cryptotypes.PubKey) error {
	return p.GetAccount().SetPubKey(pubKey)
}

// GetAccountNumber implements authtypes.AccountI
func (p *Profile) GetAccountNumber() uint64 {
	return p.GetAccount().GetAccountNumber()
}

// SetAccountNumber implements authtypes.AccountI
func (p *Profile) SetAccountNumber(accountNumber uint64) error {
	return p.GetAccount().SetAccountNumber(accountNumber)
}

// GetSequence implements authtypes.AccountI
func (p *Profile) GetSequence() uint64 {
	return p.GetAccount().GetSequence()
}

// SetSequence implements authtypes.AccountI
func (p *Profile) SetSequence(sequence uint64) error {
	return p.GetAccount().SetSequence(sequence)
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (p *Profile) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if p.Account != nil {
		var account authtypes.AccountI
		return unpacker.UnpackAny(p.Account, &account)
	}
	return nil
}

// Validate check the validity of the Profile
func (p *Profile) Validate() error {
	if strings.TrimSpace(p.Dtag) == "" || p.Dtag == DoNotModify {
		return fmt.Errorf("invalid profile DTag: %s", p.Dtag)
	}

	if p.Moniker == DoNotModify {
		return fmt.Errorf("invalid profile moniker: %s", p.Moniker)
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
	Moniker       string         `json:"moniker" yaml:"moniker"`
	Bio           string         `json:"bio" yaml:"bio"`
	Pictures      Pictures       `json:"pictures" yaml:"pictures"`
	CreationDate  time.Time      `json:"creation_date" yaml:"creation_date"`
}

// Ensure that acc
//// String implements authtypes.AccountIount implements stringer
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
		DTag:          p.Dtag,
		Moniker:       p.Moniker,
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
		DTag:          p.Dtag,
		Moniker:       p.Moniker,
		Bio:           p.Bio,
		Pictures:      p.Pictures,
		CreationDate:  p.CreationDate,
	})
}

// -------------------------------------------------------------------------------------------------------------------

// ProfileUpdate contains all the data that can be updated about a profile.
// When performing an update, if a field should not be edited then it must be set to types.DoNotModify
type ProfileUpdate struct {
	Dtag     string
	Moniker  string
	Bio      string
	Pictures Pictures
}

// NewProfileUpdate builds a new ProfileUpdate instance containing the given data
func NewProfileUpdate(dTag, moniker, bio string, pictures Pictures) *ProfileUpdate {
	return &ProfileUpdate{
		Dtag:     dTag,
		Moniker:  moniker,
		Bio:      bio,
		Pictures: pictures,
	}
}

// Update updates the fields of a given profile. An error is
// returned if the resulting profile contains invalid values.
func (p *Profile) Update(update *ProfileUpdate) (*Profile, error) {
	if update.Dtag == DoNotModify {
		update.Dtag = p.Dtag
	}

	if update.Moniker == DoNotModify {
		update.Moniker = p.Moniker
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

	newProfile, err := NewProfile(update.Dtag, update.Moniker, update.Bio, update.Pictures, p.CreationDate, p.GetAccount())
	if err != nil {
		return nil, err
	}

	err = newProfile.Validate()
	if err != nil {
		return nil, err
	}

	return newProfile, nil
}
