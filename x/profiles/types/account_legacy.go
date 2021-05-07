package types

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ghodss/yaml"
)

// Ensure that acc
//// String implements authtypes.AccountIount implements stringer
func (p *Profile016) String() string {
	out, _ := p.MarshalYAML()
	return out.(string)
}

// MarshalYAML returns the YAML representation of a Profile.
func (p *Profile016) MarshalYAML() (interface{}, error) {
	bs, err := yaml.Marshal(profilePretty{
		Address:       p.GetAccount().GetAddress(),
		PubKey:        p.GetPubKey().String(),
		AccountNumber: p.GetAccountNumber(),
		Sequence:      p.GetSequence(),
		DTag:          p.DTag,
		Nickname:      p.Moniker,
		Bio:           p.Bio,
		Pictures:      p.Pictures,
		CreationDate:  p.CreationDate,
	})

	if err != nil {
		return nil, err
	}

	return string(bs), nil
}

// GetAccount returns the underlying account as an authtypes.AccountI instance
func (p *Profile016) GetAccount() authtypes.AccountI {
	return p.Account.GetCachedValue().(authtypes.AccountI)
}

// GetAddress implements authtypes.AccountI
func (p *Profile016) GetAddress() sdk.AccAddress {
	return p.GetAccount().GetAddress()
}

// GetPubKey implements authtypes.AccountI
func (p *Profile016) GetPubKey() cryptotypes.PubKey {
	return p.GetAccount().GetPubKey()
}

// GetAccountNumber implements authtypes.AccountI
func (p *Profile016) GetAccountNumber() uint64 {
	return p.GetAccount().GetAccountNumber()
}

// GetSequence implements authtypes.AccountI
func (p *Profile016) GetSequence() uint64 {
	return p.GetAccount().GetSequence()
}
