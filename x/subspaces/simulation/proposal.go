package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v5/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

// DONTCOVER

// Simulation operation weights constants
const (
	DefaultWeightMsgUpdateParams int = 100

	OpWeightMsgUpdateSubspaceFeeTokens = "op_weight_msg_update_subspace_fee_tokens" //nolint:gosec
)

// ProposalMsgs defines the module weighted proposals' contents
func ProposalMsgs(k keeper.Keeper) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			OpWeightMsgUpdateSubspaceFeeTokens,
			DefaultWeightMsgUpdateParams,
			SimulateMsgUpdateSubspaceFeeTokens(k),
		),
	}
}

// SimulateMsgUpdateParams returns a random MsgUpdateParams
func SimulateMsgUpdateSubspaceFeeTokens(k keeper.Keeper) simtypes.MsgSimulatorFn {
	return func(r *rand.Rand, ctx sdk.Context, _ []simtypes.Account) sdk.Msg {
		// use the default gov module account address as authority
		var authority sdk.AccAddress = address.Module("gov")

		var subspaceID uint64

		subspaces := k.GetAllSubspaces(ctx)
		if len(subspaces) > 0 {
			subspace := RandomSubspace(r, subspaces)
			subspaceID = subspace.ID
		}

		return types.NewMsgUpdateSubspaceFeeTokens(
			subspaceID,
			GenerateRandomFeeTokens(r),
			authority.String(),
		)
	}
}
