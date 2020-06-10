package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/desmos-labs/desmos/x/profile/client/cli"
	"github.com/desmos-labs/desmos/x/profile/client/rest"
)

var NSParamsEditProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitNameSurnameParamsEditProposal, rest.NameSurnameParamsEditProposalRESTHandler)
var MonikerParamsEditProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitMonikerParamsEditProposal, rest.MonikerParamsEditProposalRESTHandler)
var BioParamsEditProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitBioParamsEditProposal, rest.BioParamsEditProposalRESTHandler)
