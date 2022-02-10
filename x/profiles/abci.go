package profiles

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v2/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) {
	// check for events connected to the DTag auctioneer smart contract
	events := ctx.EventManager().Events()
	k.Logger(ctx).Info("Events length:", "length", len(events))
	for _, event := range events {
		k.Logger(ctx).Info("Event", "type", event.Type)
		if event.Type == profilestypes.EventTypeDTagTransferAccept || event.Type == profilestypes.EventTypeDTagTransferRefuse {
			k.IteratePermissionedContracts(ctx, func(index int64, contract profilestypes.PermissionedContract) bool {
				k.Logger(ctx).Info("Iterating permissioned contract: ", "contract address", contract.Address)
				var userAddr string
				for _, attr := range event.Attributes {
					k.Logger(ctx).Info("Iterating attributes contract: ", "attribute:", string(attr.Key))
					if string(attr.Key) == profilestypes.AttributeRequestReceiver {
						userAddr = string(attr.Value)
					}
					if string(attr.Key) == profilestypes.AttributeRequestSender {
						if string(attr.Value) == contract.Address {
							k.Logger(ctx).Info("Updating dtag auction status...")
							err := k.UpdateDtagAuctionStatus(ctx, contract.Address, userAddr, event.Type)
							if err != nil {
								k.Logger(ctx).Error("ERROR", err)
								fmt.Println("[!] error: ", err.Error())
							}
							k.Logger(ctx).Info("Updating dtag auction status SUCCESSFUL")
						}
					}
				}
				return false
			})
		}

	}
}
