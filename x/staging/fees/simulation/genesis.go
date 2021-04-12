package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	magpietypes "github.com/desmos-labs/desmos/x/magpie/types"
	poststypes "github.com/desmos-labs/desmos/x/posts/types"
	profilestypes "github.com/desmos-labs/desmos/x/profiles/types"
	relationshipstypes "github.com/desmos-labs/desmos/x/relationships/types"
	reportstypes "github.com/desmos-labs/desmos/x/reports/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/x/fees/types"
)

const (
	MinFees = "min_fees"
)

var msgsTypes = []string{
	magpietypes.ActionCreationSession,
	poststypes.ActionCreatePost,
	poststypes.ActionEditPost,
	poststypes.ActionAnswerPoll,
	poststypes.ActionAddPostReaction,
	poststypes.ActionRemovePostReaction,
	poststypes.ActionRegisterReaction,
	profilestypes.ActionSaveProfile,
	profilestypes.ActionDeleteProfile,
	profilestypes.ActionRequestDtag,
	profilestypes.ActionAcceptDtagTransfer,
	profilestypes.ActionRefuseDTagTransferRequest,
	profilestypes.ActionCancelDTagTransferRequest,
	relationshipstypes.ActionCreateRelationship,
	relationshipstypes.ActionDeleteRelationship,
	relationshipstypes.ActionBlockUser,
	relationshipstypes.ActionUnblockUser,
	reportstypes.ActionReportPost,
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
