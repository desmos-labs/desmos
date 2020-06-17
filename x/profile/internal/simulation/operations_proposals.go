package simulation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"math/rand"
)

const (
	OpWeightSubmitNameSurnameParamsEditProposal = "op_weight_submit_name_surname_proposal"
	OpWeightSubmitMonikerParamsEditProposal     = "op_weight_submit_moniker_proposal"
	OpWeightSubmitBiographyParamsEditProposal   = "op_weight_submit_biography_proposal"
)

// ProposalContents defines the module weighted proposals' contents
func ProposalContents() []simulation.WeightedProposalContent {
	return []simulation.WeightedProposalContent{
		{
			AppParamsKey:       OpWeightSubmitNameSurnameParamsEditProposal,
			DefaultWeight:      params.DefaultWeightEditProfileParamsProposal,
			ContentSimulatorFn: SimulateNameSurnameEditParamsProposal,
		},
		{
			AppParamsKey:       OpWeightSubmitMonikerParamsEditProposal,
			DefaultWeight:      params.DefaultWeightEditProfileParamsProposal,
			ContentSimulatorFn: SimulateMonikerEditParamsProposal,
		},
		{
			AppParamsKey:       OpWeightSubmitBiographyParamsEditProposal,
			DefaultWeight:      params.DefaultWeightEditProfileParamsProposal,
			ContentSimulatorFn: SimulateBiographyEditParamsProposal,
		},
	}
}

func SimulateNameSurnameEditParamsProposal(r *rand.Rand, _ sdk.Context, _ []simulation.Account) gov.Content {
	return types.NewNameSurnameParamsEditProposal(
		simulation.RandStringOfLength(r, 140),
		simulation.RandStringOfLength(r, 5000),
		RandomNameSurnameParams(r),
	)
}

func SimulateMonikerEditParamsProposal(r *rand.Rand, _ sdk.Context, _ []simulation.Account) gov.Content {
	return types.NewMonikerParamsEditProposal(
		simulation.RandStringOfLength(r, 140),
		simulation.RandStringOfLength(r, 5000),
		RandomMonikerParams(r),
	)
}

func SimulateBiographyEditParamsProposal(r *rand.Rand, _ sdk.Context, _ []simulation.Account) gov.Content {
	return types.NewBioParamsEditProposal(
		simulation.RandStringOfLength(r, 140),
		simulation.RandStringOfLength(r, 5000),
		RandomBioParams(r),
	)
}
