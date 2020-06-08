package models

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

const (
	ProposalTypeEditParams string = "ParamsEditProposal"
)

func init() {
	gov.RegisterProposalType(ProposalTypeEditParams)
	gov.RegisterProposalTypeCodec(EditParamsProposal{}, "desmos/EditParamsProposal")
}

/////////////////////////////////////////////
/////////EditParamsProposal///////
////////////////////////////////////////////

// Implements Proposal Interface
var _ gov.Content = EditParamsProposal{}

type EditParamsProposal struct {
	Title                string                `json:"title" yaml:"title"`
	Description          string                `json:"description" yaml:"description"`
	NameSurnameLenParams *NameSurnameLenParams `json:"name_surname_len_params" yaml:"name_surname_len_params"`
	MonikerLenParams     *MonikerLenParams     `json:"moniker_len_params" yaml:"moniker_len_params"`
	BioLenParams         *BioLenParams         `json:"bio_len_params" yaml:"bio_len_params"`
}

func (ep EditParamsProposal) GetTitle() string {
	return ep.Title
}

func (ep EditParamsProposal) GetDescription() string {
	return ep.Description
}

func (ep EditParamsProposal) ProposalRoute() string {
	return RouterKey
}

func (ep EditParamsProposal) ProposalType() string {
	return ProposalTypeEditParams
}

func (ep EditParamsProposal) ValidateBasic() error {
	err := ValidateNameSurnameLenParams(ep)
	if err != nil {
		return err
	}
	return nil
}

func (ep EditParamsProposal) String() string {
	out := fmt.Sprintf(`Edit Profiles' params proposal:
  Title:       %s
  Description: %s
  Proposed lengths:`, ep.Title, ep.Description)

	return fmt.Sprintf(`Edit Profiles' params proposal:
  Title:       %s
  Description: %s
  Proposed lengths:
  Name/Surname: Min %s, Max %s,
  Moniker: Min %s, Max %s,
  Biography: Max %s,
`, ep.Title, ep.Description, ep.NameSurnameLenParams.MinNameSurnameLen, ep.NameSurnameLenParams.MaxNameSurnameLen)
}
