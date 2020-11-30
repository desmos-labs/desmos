package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/fees/types"
)

const (
	MinFees = "min_fees"
)

var msgsTypes = []string{
	"create_session",
	"create_post",
	"edit_post",
	"answer_poll",
	"add_post_reaction",
	"remove_post_reaction",
	"register_reaction",
	"save_profile",
	"delete_profile",
	"request_dtag",
	"accept_dtag_request",
	"refuse_dtag_request",
	"cancel_dtag_request",
	"create_relationship",
	"delete_relationship",
	"block_user",
	"unblock_user",
	"report_post",
}

// RandomizedGenState generates a random GenesisState for fees
func RandomizedGenState(simState *module.SimulationState) {
	var minFees []types.MinFee
	simState.AppParams.GetOrGenerate(
		simState.Cdc, MinFees, &minFees, simState.Rand, func(r *rand.Rand) {
			minFees = GenMinFees(r)
		})

	feesGenesis := types.NewGenesisState(types.NewParams(minFees))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(feesGenesis)

	fmt.Printf("Selected randomly generated fees parameters:\n%s\n",
		codec.MustMarshalJSONIndent(simState.Cdc, feesGenesis.Params.MinFees),
	)
}

// GenMinFees randomized MinFees
func GenMinFees(r *rand.Rand) (fees []types.MinFee) {
	randFixedFeeNum := simulation.RandIntBetween(r, 0, len(msgsTypes))
	for i := 0; i < randFixedFeeNum; i++ {
		amt := simulation.RandIntBetween(r, 1, 100)
		fees = append(fees, types.NewMinFee(msgsTypes[i], sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(amt))))))
	}
	return fees
}
