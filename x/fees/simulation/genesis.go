package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	magpieTypes "github.com/desmos-labs/desmos/x/magpie/types"
	postsTypes "github.com/desmos-labs/desmos/x/posts/types"
	profilesTypes "github.com/desmos-labs/desmos/x/profiles/types"
	relationshipsTypes "github.com/desmos-labs/desmos/x/relationships/types"
	reportsTypes "github.com/desmos-labs/desmos/x/reports/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/desmos-labs/desmos/x/fees/types"
)

const (
	MinFees = "min_fees"
)

var msgsTypes = []string{
	magpieTypes.ActionCreationSession,
	postsTypes.ActionCreatePost,
	postsTypes.ActionEditPost,
	postsTypes.ActionAnswerPoll,
	postsTypes.ActionAddPostReaction,
	postsTypes.ActionRemovePostReaction,
	postsTypes.ActionRegisterReaction,
	profilesTypes.ActionSaveProfile,
	profilesTypes.ActionDeleteProfile,
	profilesTypes.ActionRequestDtag,
	profilesTypes.ActionAcceptDtagTransfer,
	profilesTypes.ActionRefuseDTagTransferRequest,
	profilesTypes.ActionCancelDTagTransferRequest,
	relationshipsTypes.ActionCreateRelationship,
	relationshipsTypes.ActionDeleteRelationship,
	relationshipsTypes.ActionBlockUser,
	relationshipsTypes.ActionUnblockUser,
	reportsTypes.ActionReportPost,
}

// RandomizedGenState generates a random GenesisState for fees
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

// GenMinFees randomized MinFees
func GenMinFees(r *rand.Rand) (fees []types.MinFee) {
	// 50% chance to have min fee
	if randFixedFeeNum := r.Intn(101); randFixedFeeNum <= 50 {
		for i := 0; i < randFixedFeeNum; i++ {
			amt := simulation.RandIntBetween(r, 1, 100)
			fees = append(fees, types.NewMinFee(msgsTypes[i], sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(amt))))))
		}
	}

	return fees
}
