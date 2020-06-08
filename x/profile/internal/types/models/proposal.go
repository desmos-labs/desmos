package models

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
