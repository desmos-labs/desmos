package profiles

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v2/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	// check for events connected to the DTag auctioneer smart contract
	events := ctx.EventManager().Events().ToABCIEvents()

	for _, event := range events {
		if event.Type == profilestypes.ActionAcceptDTagTransfer {
			k.IteratePermissionedContracts(ctx, func(index int64, contract profilestypes.PermissionedContract) bool {
				var userAddr string
				for _, attr := range event.Attributes {
					if string(attr.Key) == profilestypes.AttributeRequestReceiver {
						userAddr = string(attr.Value)
						break
					}
				}
				err := k.UpdateDtagAuctionStatus(ctx, contract.Address, userAddr)

				if err != nil {
					k.Logger(ctx).Error("ERROR", err)
					fmt.Println("[!] error: ", err.Error())
				}

				return false
			})
		}

	}
}
