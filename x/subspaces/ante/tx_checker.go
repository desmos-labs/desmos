package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	antetypes "github.com/desmos-labs/desmos/v5/x/subspaces/ante/types"

	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

// GetSocialTxSubspaceID returns the valid subspace id, returns false if it is invalid
func GetSocialTxSubspaceID(tx sdk.Tx) (uint64, bool) {
	subspaceID := uint64(0)
	for _, msg := range tx.GetMsgs() {
		if socialMsg, ok := msg.(types.SocialMsg); ok {
			if subspaceID == 0 {
				subspaceID = socialMsg.GetSubspaceID()
			}

			if socialMsg.GetSubspaceID() == subspaceID {
				continue
			}
		}

		return 0, false
	}
	return subspaceID, true
}

// CheckTxFeeWithSubspaceMinPrices returns the tx checker that including the subspace allowed tokens into minimum prices list
func CheckTxFeeWithSubspaceMinPrices(txFeeChecker ante.TxFeeChecker, sk antetypes.SubspacesKeeper) ante.TxFeeChecker {
	return func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
		subspaceID, isSocialTx := GetSocialTxSubspaceID(tx)
		if !isSocialTx {
			return txFeeChecker(ctx, tx)
		}

		subspace, found := sk.GetSubspace(ctx, subspaceID)
		if !found {
			return txFeeChecker(ctx, tx)
		}

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
