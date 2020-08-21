package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/profiles/keeper"
)

// Simulation operation weights constants
const (
	OpWeightMsgSaveProfile         = "op_weight_msg_save_profile"
	OpWeightMsgDeleteProfile       = "op_weight_msg_delete_profile"
	OpWeightMsgCreateRelationship  = "op_weight_msg_create_relationship"
	OpWeightMsgRequestRelationship = "op_weight_msg_request_relationship"
	OpWeightMsgAcceptRelationship  = "op_weight_msg_request_relationship"
	OpWeightMsgDenyRelationship    = "op_weight_msg_deny_relationship"
	OpWeightMsgDeleteRelationship  = "op_weight_msg_delete_relationship"

	DefaultGasValue = 350000
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

	var weightMsgCreateRelationship int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreateRelationship, &weightMsgCreateRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRelationship = params.DefaultWeightMsgCreateRelationship
		},
	)

	var weightMsgRequestRelationship int
	appParams.GetOrGenerate(cdc, OpWeightMsgRequestRelationship, &weightMsgRequestRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgRequestRelationship = params.DefaultWeightMsgRequestRelationship
		},
	)

	var weightMsgAcceptRelationship int
	appParams.GetOrGenerate(cdc, OpWeightMsgAcceptRelationship, &weightMsgAcceptRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgAcceptRelationship = params.DefaultWeightMsgAcceptRelationship
		},
	)

	var weightMsgDenyRelationship int
	appParams.GetOrGenerate(cdc, OpWeightMsgDenyRelationship, &weightMsgDenyRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgDenyRelationship = params.DefaultWeightMsgDenyRelationship
		},
	)

	var weightMsgDeleteRelationship int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeleteRelationship, &weightMsgDeleteRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgDenyRelationship = params.DefaultWeightMsgDeleteRelationship
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
		sim.NewWeightedOperation(
			weightMsgCreateRelationship,
			SimulateMsgCreateMonoDirectionalRelationship(k, ak),
		),
		sim.NewWeightedOperation(
			weightMsgRequestRelationship,
			SimulateMsgRequestBidirectionalRelationship(k, ak),
		),
		sim.NewWeightedOperation(
			weightMsgAcceptRelationship,
			SimulateMsgAcceptBidirectionalRelationship(k, ak),
		),
		sim.NewWeightedOperation(
			weightMsgDenyRelationship,
			SimulateMsgDenyBidirectionalRelationship(k, ak),
		),
		sim.NewWeightedOperation(
			weightMsgDeleteRelationship,
			SimulateMsgDeleteRelationship(k, ak),
		),
	}
}
