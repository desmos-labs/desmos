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
	OpWeightMsgCreateAccount = "op_weight_msg_create_account"
	OpWeightMsgEditAccount   = "op_weight_msg_edit_account"
	OpWeightMsgDeleteAccount = "op_weight_msg_delete_account"

	DefaultGasValue = 5000000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(appParams sim.AppParams, cdc *codec.Codec, k keeper.Keeper, ak auth.AccountKeeper) sim.WeightedOperations {
	var weightMsgCreateAccount int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreateAccount, &weightMsgCreateAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAccount = params.DefaultWeightMsgCreateAccount
		},
	)

	var weightMsgEditAccount int
	appParams.GetOrGenerate(cdc, OpWeightMsgEditAccount, &weightMsgEditAccount, nil,
		func(_ *rand.Rand) {
			weightMsgEditAccount = params.DefaultWeightMsgEditAccount
		},
	)

	var weightMsgDeleteAccount int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeleteAccount, &weightMsgDeleteAccount, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAccount = params.DefaultWeightMsgDeleteAccount
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgCreateAccount,
			SimulateMsgCreateAccount(ak),
		),
		sim.NewWeightedOperation(
			weightMsgEditAccount,
			SimulateMsgEditAccount(k, ak),
		),
		sim.NewWeightedOperation(
			weightMsgDeleteAccount,
			SimulateMsgDeleteAccount(k, ak),
		),
	}
}
