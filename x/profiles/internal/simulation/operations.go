package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/profiles/internal/keeper"
)

// Simulation operation weights constants
const (
	OpWeightMsgSaveProfile   = "op_weight_msg_save_profile"
	OpWeightMsgDeleteProfile = "op_weight_msg_delete_profile"

	DefaultGasValue = 200000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(appParams sim.AppParams, cdc *codec.Codec, k keeper.Keeper, ak auth.AccountKeeper) sim.WeightedOperations {
	var weightMsgSaveProfile int
	appParams.GetOrGenerate(cdc, OpWeightMsgSaveProfile, &weightMsgSaveProfile, nil,
		func(_ *rand.Rand) {
			weightMsgSaveProfile = params.DefaultWeightMsgSaveAccount
		},
	)

	var weightMsgDeleteProfile int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeleteProfile, &weightMsgDeleteProfile, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProfile = params.DefaultWeightMsgDeleteAccount
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgSaveProfile,
			SimulateMsgSaveProfile(k, ak),
		),
		sim.NewWeightedOperation(
			weightMsgDeleteProfile,
			SimulateMsgDeleteProfile(k, ak),
		),
	}
}
