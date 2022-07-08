package types

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/desmos-labs/desmos/v4/x/commons"

	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ghodss/yaml"
	"github.com/gogo/protobuf/proto"
)

var (
	_ authtypes.AccountI      = (*Profile)(nil)
	_ exported.VestingAccount = (*Profile)(nil)
)

// NewProfile builds a new profile having the given DTag, creator and creation date
func NewProfile(dTag string, nickname, bio string, pictures Pictures, creationDate time.Time, account authtypes.AccountI) (*Profile, error) {
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

// getVestingAccount returns the underlying account as an exported.VestingAccount instance
func (p *Profile) getVestingAccount() exported.VestingAccount {
	acc, ok := p.Account.GetCachedValue().(exported.VestingAccount)
	if !ok {
		return nil
	}
	return acc
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

// LockedCoins implements exported.VestingAccount
func (p *Profile) LockedCoins(blockTime time.Time) sdk.Coins {
	acc := p.getVestingAccount()
	if acc == nil {
		return sdk.NewCoins()
	}
	return acc.LockedCoins(blockTime)
}

// TrackDelegation implements exported.VestingAccount
func (p *Profile) TrackDelegation(blockTime time.Time, balance, amount sdk.Coins) {
	acc := p.getVestingAccount()
	if acc == nil {
		return
	}

	acc.TrackDelegation(blockTime, balance, amount)
	err := p.setAccount(acc)
	if err != nil {
		panic(err)
	}
}

// TrackUndelegation implements exported.VestingAccount
func (p *Profile) TrackUndelegation(amount sdk.Coins) {
	acc := p.getVestingAccount()
	if acc == nil {
		return
	}

	acc.TrackUndelegation(amount)
	err := p.setAccount(acc)
	if err != nil {
		panic(err)
	}
}

// GetVestedCoins implements exported.VestingAccount
func (p *Profile) GetVestedCoins(blockTime time.Time) sdk.Coins {
	acc := p.getVestingAccount()
	if acc == nil {
		return sdk.NewCoins()
	}
	return acc.GetVestedCoins(blockTime)
}

// GetVestingCoins implements exported.VestingAccount
func (p *Profile) GetVestingCoins(blockTime time.Time) sdk.Coins {
	acc := p.getVestingAccount()
	if acc == nil {
		return sdk.NewCoins()
	}
	return acc.GetVestingCoins(blockTime)
}

// GetStartTime implements exported.VestingAccount
func (p *Profile) GetStartTime() int64 {
	acc := p.getVestingAccount()
	if acc == nil {
		return -1
	}
	return acc.GetStartTime()
}

// GetEndTime implements exported.VestingAccount
func (p *Profile) GetEndTime() int64 {
	acc := p.getVestingAccount()
	if acc == nil {
		return -1
	}
	return acc.GetEndTime()
}

// GetOriginalVesting implements exported.VestingAccount
func (p *Profile) GetOriginalVesting() sdk.Coins {
	acc := p.getVestingAccount()
	if acc == nil {
		return sdk.NewCoins()
	}
	return acc.GetOriginalVesting()
}

// GetDelegatedFree implements exported.VestingAccount
func (p *Profile) GetDelegatedFree() sdk.Coins {
	acc := p.getVestingAccount()
	if acc == nil {
		return sdk.NewCoins()
	}
	return acc.GetDelegatedFree()
}

// GetDelegatedVesting implements exported.VestingAccount
func (p *Profile) GetDelegatedVesting() sdk.Coins {
	acc := p.getVestingAccount()
	if acc == nil {
		return sdk.NewCoins()
	}
	return acc.GetDelegatedVesting()
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
func (p *Profile) MarshalJSON() ([]byte, error) {
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

// NewPictures is a constructor function for Pictures
func NewPictures(profile, cover string) Pictures {
	return Pictures{
		Profile: profile,
		Cover:   cover,
	}
}

// Validate check the validity of the Pictures
func (pic Pictures) Validate() error {
	if pic.Profile != "" {
		valid := commons.IsURIValid(pic.Profile)
		if !valid {
			return fmt.Errorf("invalid profile picture uri provided")
		}
	}

	if pic.Cover != "" {
		valid := commons.IsURIValid(pic.Cover)
		if !valid {
			return fmt.Errorf("invalid profile cover uri provided")
		}
	}

	return nil
}
