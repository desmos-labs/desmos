package profiles

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v2/x/profiles/keeper"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) {
	// check for events connected to the DTag auctioneer smart contract
	// TODO this need to be entirely reviewed
	/*
		k.IteratePermissionedContracts(ctx, func(index int64, contract profilestypes.PermissionedContract) bool {
			k.Logger(ctx).Info("Iterating permissioned contract: ", "contract address", contract.Address)
			msg, err := contract.GetMessage()
			k.Logger(ctx).Info("User", "user", msg.UpdateDtagAuctionStatus.User)
			k.Logger(ctx).Info("Transfer_status", "status", msg.UpdateDtagAuctionStatus.TransferStatus)
			err = k.UpdateDtagAuctionStatus(ctx, contract.Address, msg)
			if err != nil {
				k.Logger(ctx).Error("ERROR", err)
				fmt.Println("[!] error: ", err.Error())
			}
			return false
		})
	*/
}
