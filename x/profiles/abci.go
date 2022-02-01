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
		if event.Type == profilestypes.EventTypeDTagTransferAccept || event.Type == profilestypes.EventTypeDTagTransferRefuse {
			k.IteratePermissionedContracts(ctx, func(index int64, contract profilestypes.PermissionedContract) bool {
				var userAddr string
				for _, attr := range event.Attributes {
					if string(attr.Key) == profilestypes.AttributeRequestReceiver {
						userAddr = string(attr.Value)
					}
					if string(attr.Key) == profilestypes.AttributeRequestSender {
						if string(attr.Value) == contract.Address {
							err := k.UpdateDtagAuctionStatus(ctx, contract.Address, userAddr, event.Type)
							if err != nil {
								k.Logger(ctx).Error("ERROR", err)
								fmt.Println("[!] error: ", err.Error())
							}
						}
					}
				}
				return false
			})
		}

	}
}
