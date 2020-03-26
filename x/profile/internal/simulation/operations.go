package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"math/rand"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateAccount = "op_weight_msg_create_account"

	DefaultGasValue = 5000000
)

func WeightedOperations(appParams sim.AppParams, cdc *codec.Codec, k keeper.Keeper, ak auth.AccountKeeper) sim.WeightedOperations {
	var weightMsgCreateAccount int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreateAccount, &weightMsgCreateAccount, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAccount = params.DefaultWeightMsgCreatePost
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgCreateAccount,
			SimulateMsgCreateAccount(k, ak),
		),
	}
}
