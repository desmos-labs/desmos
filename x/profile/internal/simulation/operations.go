package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateProfile = "op_weight_msg_create_profile"
	OpWeightMsgEditProfile   = "op_weight_msg_edit_profile"
	OpWeightMsgDeleteProfile = "op_weight_msg_delete_profile"

	DefaultGasValue = 5000000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(appParams sim.AppParams, cdc *codec.Codec, k keeper.Keeper, ak auth.AccountKeeper) sim.WeightedOperations {
	var weightMsgCreateProfile int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreateProfile, &weightMsgCreateProfile, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProfile = params.DefaultWeightMsgCreateAccount
		},
	)

	var weightMsgEditProfile int
	appParams.GetOrGenerate(cdc, OpWeightMsgEditProfile, &weightMsgEditProfile, nil,
		func(_ *rand.Rand) {
			weightMsgEditProfile = params.DefaultWeightMsgEditAccount
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
			weightMsgCreateProfile,
			SimulateMsgCreateProfile(ak),
		),
		sim.NewWeightedOperation(
			weightMsgEditProfile,
			SimulateMsgEditProfile(k, ak),
		),
		sim.NewWeightedOperation(
			weightMsgDeleteProfile,
			SimulateMsgDeleteProfile(k, ak),
		),
	}
}
