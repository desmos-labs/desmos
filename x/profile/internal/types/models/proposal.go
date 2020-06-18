package models

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

const (
	ProposalTypeNameSurnameParamsEdit string = "NameSurnameParamsEdit"
	ProposalTypeMonikerParamsEdit     string = "MonikerParamsEdit"
	ProposalTypeBioParamsEdit         string = "BioParamsEdit"
)

func init() {
	gov.RegisterProposalType(ProposalTypeNameSurnameParamsEdit)
	gov.RegisterProposalTypeCodec(NameSurnameParamsEditProposal{}, "desmos/NameSurnameProfileParamsEditProposal")
	gov.RegisterProposalType(ProposalTypeMonikerParamsEdit)
	gov.RegisterProposalTypeCodec(MonikerParamsEditProposal{}, "desmos/MonikerProfileParamsEditProposal")
	gov.RegisterProposalType(ProposalTypeBioParamsEdit)
	gov.RegisterProposalTypeCodec(BioParamsEditProposal{}, "desmos/BioParamsEditProposal")

}

/////////////////////////////////////////////
/////////NameSurnameParamsEditProposal//////
///////////////////////////////////////////

// Implements Proposal Interface
var _ gov.Content = NameSurnameParamsEditProposal{}

type NameSurnameParamsEditProposal struct {
	Title             string             `json:"title" yaml:"title"`
	Description       string             `json:"description" yaml:"description"`
	NameSurnameParams NameSurnameLengths `json:"name_surname_params" yaml:"name_surname_params"`
}

func NewNameSurnameParamsEditProposal(title, description string, nsParams NameSurnameLengths) gov.Content {
	return NameSurnameParamsEditProposal{
		Title:             title,
		Description:       description,
		NameSurnameParams: nsParams,
	}
}

func (nsp NameSurnameParamsEditProposal) GetTitle() string {
	return nsp.Title
}

func (nsp NameSurnameParamsEditProposal) GetDescription() string {
	return nsp.Description
}

func (nsp NameSurnameParamsEditProposal) ProposalRoute() string {
	return RouterKey
}

func (nsp NameSurnameParamsEditProposal) ProposalType() string {
	return ProposalTypeNameSurnameParamsEdit
}

func (nsp NameSurnameParamsEditProposal) ValidateBasic() error {
	err := ValidateNameSurnameLenParams(nsp.NameSurnameParams)
	if err != nil {
		return err
	}
	return gov.ValidateAbstract(nsp)
}

func (nsp NameSurnameParamsEditProposal) String() string {
	return fmt.Sprintf(`Name/Surname Profiles' params edit proposal:
  Title:       %s
  Description: %s
  Proposed name/Surname params lengths:
  Min: %s
  Max: %s
`, nsp.Title, nsp.Description, nsp.NameSurnameParams.MinNameSurnameLen, nsp.NameSurnameParams.MaxNameSurnameLen)
}

//////////////////////////////////////////
/////////MonikerParamsEditProposal///////
////////////////////////////////////////

// Implements Proposal Interface
var _ gov.Content = MonikerParamsEditProposal{}

type MonikerParamsEditProposal struct {
	Title         string         `json:"title" yaml:"title"`
	Description   string         `json:"description" yaml:"description"`
	MonikerParams MonikerLengths `json:"moniker_params" yam:"moniker_params"`
}

func NewMonikerParamsEditProposal(title, description string, mParams MonikerLengths) gov.Content {
	return MonikerParamsEditProposal{
		Title:         title,
		Description:   description,
		MonikerParams: mParams,
	}
}

func (mp MonikerParamsEditProposal) GetTitle() string {
	return mp.Title
}

func (mp MonikerParamsEditProposal) GetDescription() string {
	return mp.Description
}

func (mp MonikerParamsEditProposal) ProposalRoute() string {
	return RouterKey
}

func (mp MonikerParamsEditProposal) ProposalType() string {
	return ProposalTypeMonikerParamsEdit
}

func (mp MonikerParamsEditProposal) ValidateBasic() error {
	err := ValidateMonikerLenParams(mp.MonikerParams)
	if err != nil {
		return err
	}
	return gov.ValidateAbstract(mp)
}

func (mp MonikerParamsEditProposal) String() string {
	return fmt.Sprintf(`Moniker Profiles' params edit proposal:
  Title:       %s
  Description: %s
  Proposed moniker params lengths:
  Min: %s
  Max: %s
`, mp.Title, mp.Description, mp.MonikerParams.MinMonikerLen, mp.MonikerParams.MaxMonikerLen)
}

//////////////////////////////////////////
/////////BioParamsEditProposal///////////
////////////////////////////////////////

// Implements Proposal Interface
var _ gov.Content = BioParamsEditProposal{}

type BioParamsEditProposal struct {
	Title       string           `json:"title" yaml:"title"`
	Description string           `json:"description" yaml:"description"`
	BioParams   BiographyLengths `json:"bio_params" yaml:"bio_params"`
}

func NewBioParamsEditProposal(title, description string, bParams BiographyLengths) gov.Content {
	return BioParamsEditProposal{
		Title:       title,
		Description: description,
		BioParams:   bParams,
	}
}

func (bp BioParamsEditProposal) GetTitle() string {
	return bp.Title
}

func (bp BioParamsEditProposal) GetDescription() string {
	return bp.Description
}

func (bp BioParamsEditProposal) ProposalRoute() string {
	return RouterKey
}

func (bp BioParamsEditProposal) ProposalType() string {
	return ProposalTypeBioParamsEdit
}

func (bp BioParamsEditProposal) ValidateBasic() error {
	err := ValidateBioLenParams(bp.BioParams)
	if err != nil {
		return err
	}
	return gov.ValidateAbstract(bp)
}

func (bp BioParamsEditProposal) String() string {
	return fmt.Sprintf(`Biography Profiles' params edit proposal:
  Title:       %s
  Description: %s
  Proposed biography params lengths:
  Max: %s
`, bp.Title, bp.Description, bp.BioParams.MaxBioLen)
}
