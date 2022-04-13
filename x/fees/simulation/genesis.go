package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v3/x/fees/types"
)

var msgsTypes = []string{
	sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
	sdk.MsgTypeURL(&profilestypes.MsgDeleteProfile{}),
	sdk.MsgTypeURL(&profilestypes.MsgRequestDTagTransfer{}),
	sdk.MsgTypeURL(&profilestypes.MsgAcceptDTagTransferRequest{}),
	sdk.MsgTypeURL(&profilestypes.MsgRefuseDTagTransferRequest{}),
	sdk.MsgTypeURL(&profilestypes.MsgCancelDTagTransferRequest{}),
}

// RandomizedGenState generates a random GenesisState for the fees module
func RandomizedGenState(simState *module.SimulationState) {
	minFees := randomMinFees(simState.Rand)
	feesGenesis := types.NewGenesisState(types.NewParams(minFees))

	bz, err := simState.Cdc.MarshalJSON(feesGenesis)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated fees parameters:\n%s\n", bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(feesGenesis)
}

// randomMinFees returns a randomly generated types.MinFee slice
func randomMinFees(r *rand.Rand) []types.MinFee {
	// 50% chance of not having min fees
	randFixedFeeNum := r.Intn(101)
	if randFixedFeeNum <= 50 {
		return nil
	}

	feesLength := r.Intn(len(msgsTypes))
	fees := make([]types.MinFee, feesLength)
	for i := 0; i < feesLength; i++ {
		amt := simulation.RandIntBetween(r, 1, 100)
		fees[i] = types.NewMinFee(
			msgsTypes[r.Intn(len(msgsTypes))],
			sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(amt)))),
		)
	}

	return removeDuplicatedFees(fees)
}

// removeDuplicatedFees removes all the duplicated min fees checking for message types
func removeDuplicatedFees(minFees []types.MinFee) []types.MinFee {
	var result []types.MinFee
	for _, minFee := range minFees {
		if !types.ContainsMinFee(result, minFee.MessageType) {
			result = append(result, minFee)
		}
	}
	return result
}
