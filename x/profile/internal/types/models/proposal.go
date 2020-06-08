package models

import "github.com/cosmos/cosmos-sdk/x/gov"

const (
	ProposalTypeNameSurnameParamsEdit string = "NameSurnameParamsEdit"
	ProposalTypeMonikerParamsEdit     string = "MonikerParamsEdit"
	ProposalTypeBioParamsEdit         string = "BioParamsEdit"
)

type NameSurnameParamsEditProposal struct {
	Title             string               `json:"title" yaml:"title"`
	Description       string               `json:"description" yaml:"description"`
	NameSurnameParams NameSurnameLenParams `json:"name_surname_params" yaml:"name_surname_params"`
}

var _ gov.Content = NameSurnameParamsEditProposal{}

func init() {
	gov.RegisterProposalType(ProposalTypeNameSurnameParamsEdit)
	gov.RegisterProposalTypeCodec(NameSurnameParamsEditProposal{}, "desmos/NameSurnameProfileParamsEditProposal")
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

func (nsp NameSurnameParamsEditProposal) ValidateBasic() string {

}
