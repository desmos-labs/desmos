package models

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

const (
	ProposalTypeNameSurnameParamsEdit string = "EditNameSurnameParams"
	ProposalTypeMonikerParamsEdit     string = "EditMonikerParams"
	ProposalTypeBioParamsEdit         string = "EditBioParams"
)

func init() {
	gov.RegisterProposalType(ProposalTypeNameSurnameParamsEdit)
	gov.RegisterProposalTypeCodec(EditNameSurnameParamsProposal{}, "desmos/EditNameSurnameParamsProposal")
	gov.RegisterProposalType(ProposalTypeMonikerParamsEdit)
	gov.RegisterProposalTypeCodec(EditMonikerParamsProposal{}, "desmos/EditMonikerParamsProposal")
	gov.RegisterProposalType(ProposalTypeBioParamsEdit)
	gov.RegisterProposalTypeCodec(EditBioParamsProposal{}, "desmos/EditBioParamsProposal")

}

/////////////////////////////////////////////
/////////EditNameSurnameParamsProposal//////
///////////////////////////////////////////

// Implements Proposal Interface
var _ gov.Content = EditNameSurnameParamsProposal{}

type EditNameSurnameParamsProposal struct {
	Title             string             `json:"title" yaml:"title"`
	Description       string             `json:"description" yaml:"description"`
	NameSurnameParams NameSurnameLengths `json:"name_surname_params" yaml:"name_surname_params"`
}

func NewNameSurnameParamsEditProposal(title, description string, nsParams NameSurnameLengths) gov.Content {
	return EditNameSurnameParamsProposal{
		Title:             title,
		Description:       description,
		NameSurnameParams: nsParams,
	}
}

func (nsp EditNameSurnameParamsProposal) GetTitle() string {
	return nsp.Title
}

func (nsp EditNameSurnameParamsProposal) GetDescription() string {
	return nsp.Description
}

func (nsp EditNameSurnameParamsProposal) ProposalRoute() string {
	return RouterKey
}

func (nsp EditNameSurnameParamsProposal) ProposalType() string {
	return ProposalTypeNameSurnameParamsEdit
}

func (nsp EditNameSurnameParamsProposal) ValidateBasic() error {
	err := ValidateNameSurnameLenParams(nsp.NameSurnameParams)
	if err != nil {
		return err
	}
	return gov.ValidateAbstract(nsp)
}

func (nsp EditNameSurnameParamsProposal) String() string {
	return fmt.Sprintf(`Name/Surname Profiles' params edit proposal:
  Title:       %s
  Description: %s
  Proposed name/Surname params lengths:
  Min: %s
  Max: %s
`, nsp.Title, nsp.Description, nsp.NameSurnameParams.MinNameSurnameLen, nsp.NameSurnameParams.MaxNameSurnameLen)
}

//////////////////////////////////////////
/////////EditMonikerParamsProposal///////
////////////////////////////////////////

// Implements Proposal Interface
var _ gov.Content = EditMonikerParamsProposal{}

type EditMonikerParamsProposal struct {
	Title         string         `json:"title" yaml:"title"`
	Description   string         `json:"description" yaml:"description"`
	MonikerParams MonikerLengths `json:"moniker_params" yam:"moniker_params"`
}

func NewMonikerParamsEditProposal(title, description string, mParams MonikerLengths) gov.Content {
	return EditMonikerParamsProposal{
		Title:         title,
		Description:   description,
		MonikerParams: mParams,
	}
}

func (mp EditMonikerParamsProposal) GetTitle() string {
	return mp.Title
}

func (mp EditMonikerParamsProposal) GetDescription() string {
	return mp.Description
}

func (mp EditMonikerParamsProposal) ProposalRoute() string {
	return RouterKey
}

func (mp EditMonikerParamsProposal) ProposalType() string {
	return ProposalTypeMonikerParamsEdit
}

func (mp EditMonikerParamsProposal) ValidateBasic() error {
	err := ValidateMonikerLenParams(mp.MonikerParams)
	if err != nil {
		return err
	}
	return gov.ValidateAbstract(mp)
}

func (mp EditMonikerParamsProposal) String() string {
	return fmt.Sprintf(`Moniker Profiles' params edit proposal:
  Title:       %s
  Description: %s
  Proposed moniker params lengths:
  Min: %s
  Max: %s
`, mp.Title, mp.Description, mp.MonikerParams.MinMonikerLen, mp.MonikerParams.MaxMonikerLen)
}

//////////////////////////////////////////
/////////EditBioParamsProposal///////////
////////////////////////////////////////

// Implements Proposal Interface
var _ gov.Content = EditBioParamsProposal{}

type EditBioParamsProposal struct {
	Title       string           `json:"title" yaml:"title"`
	Description string           `json:"description" yaml:"description"`
	BioParams   BiographyLengths `json:"bio_params" yaml:"bio_params"`
}

func NewBioParamsEditProposal(title, description string, bParams BiographyLengths) gov.Content {
	return EditBioParamsProposal{
		Title:       title,
		Description: description,
		BioParams:   bParams,
	}
}

func (bp EditBioParamsProposal) GetTitle() string {
	return bp.Title
}

func (bp EditBioParamsProposal) GetDescription() string {
	return bp.Description
}

func (bp EditBioParamsProposal) ProposalRoute() string {
	return RouterKey
}

func (bp EditBioParamsProposal) ProposalType() string {
	return ProposalTypeBioParamsEdit
}

func (bp EditBioParamsProposal) ValidateBasic() error {
	err := ValidateBioLenParams(bp.BioParams)
	if err != nil {
		return err
	}
	return gov.ValidateAbstract(bp)
}

func (bp EditBioParamsProposal) String() string {
	return fmt.Sprintf(`Biography Profiles' params edit proposal:
  Title:       %s
  Description: %s
  Proposed biography params lengths:
  Max: %s
`, bp.Title, bp.Description, bp.BioParams.MaxBioLen)
}
