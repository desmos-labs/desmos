package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	antetypes "github.com/desmos-labs/desmos/v7/x/subspaces/ante/types"

	"github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

// hasManageSubspaceMsg checks if tx has manage subspace msg or not
func hasManageSubspaceMsg(tx sdk.Tx) bool {
	for _, msg := range tx.GetMsgs() {
		if _, ok := msg.(types.ManageSubspaceMsg); ok {
			return true
		}
	}
	return false
}

// CheckTxFeeWithSubspaceMinPrices returns the tx checker that including the subspace allowed tokens into minimum prices list
func CheckTxFeeWithSubspaceMinPrices(txFeeChecker ante.TxFeeChecker, sk antetypes.SubspacesKeeper) ante.TxFeeChecker {
	return func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
		// Perform default tx fee checker if tx is not subspace tx
		subspaceID, isSubspaceTx := GetTxSubspaceID(tx)
		if !isSubspaceTx {
			return txFeeChecker(ctx, tx)
		}

		// Perform default tx fee checker if tx includes manage subspace msg
		if hasManageSubspaceMsg(tx) {
			return txFeeChecker(ctx, tx)
		}

		// Perform default tx fee checker if subspace not found
		subspace, found := sk.GetSubspace(ctx, subspaceID)
		if !found {
			return txFeeChecker(ctx, tx)
		}

		// Perform custom tx fee checker with subspace additional fee tokens
		newMinPrices := MergeMinPrices(ctx.MinGasPrices(), sdk.NewDecCoinsFromCoins(subspace.AdditionalFeeTokens...))
		newCtx := ctx.WithMinGasPrices(newMinPrices)
		return txFeeChecker(newCtx, tx)
	}
}

// MergeMinPrices adds the other coins to the original if it does not exist inside the original
func MergeMinPrices(original sdk.DecCoins, other sdk.DecCoins) sdk.DecCoins {
	for _, coin := range other {
		if !contains(original, coin) {
			original = append(original, coin)
		}
	}

	return original.Sort()
}

// contains checks the coins slice has the target's denom
func contains(slice sdk.DecCoins, target sdk.DecCoin) bool {
	for _, v := range slice {
		if v.Denom == target.Denom {
			return true
		}
	}

	return false
}
