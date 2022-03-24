package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v2/x/fees/types"
)

const (
	MinFees = "min_fees"
)

var msgsTypes = []string{
	profilestypes.ActionSaveProfile,
	profilestypes.ActionDeleteProfile,
	profilestypes.ActionRequestDTag,
	profilestypes.ActionAcceptDTagTransfer,
	profilestypes.ActionRefuseDTagTransferRequest,
	profilestypes.ActionCancelDTagTransferRequest,
	profilestypes.ActionCreateRelationship,
	profilestypes.ActionDeleteRelationship,
	profilestypes.ActionBlockUser,
	profilestypes.ActionUnblockUser,
}

// RandomizedGenState generates a random GenesisState for the fees module
func RandomizedGenState(simState *module.SimulationState) {
	var minFees []types.MinFee
	simState.AppParams.GetOrGenerate(
		simState.Cdc, MinFees, &minFees, simState.Rand, func(r *rand.Rand) {
			minFees = GenMinFees(r)
		})

	feesGenesis := types.NewGenesisState(types.NewParams(minFees))

	bz, err := json.MarshalIndent(&feesGenesis, "", "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Selected randomly generated fees parameters:\n%s\n", bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(feesGenesis)
}

// GenMinFees returns a randomly generated types.MinFee slice
func GenMinFees(r *rand.Rand) []types.MinFee {
	// 50% chance of not having min fees
	randFixedFeeNum := r.Intn(101)
	if randFixedFeeNum <= 50 {
		return nil
	}

	feesLength := r.Intn(20)
	fees := make([]types.MinFee, feesLength)
	for i := 0; i < feesLength; i++ {
		amt := simulation.RandIntBetween(r, 1, 100)

		msgIndex := r.Intn(len(msgsTypes))
		fees[i] = types.NewMinFee(
			msgsTypes[msgIndex],
			sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(amt)))),
		)
	}

	return fees
}
