package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetDTagTransferRequests(ctx),
		k.GetParams(ctx),
		k.GetPort(ctx),
		k.GetChainLinks(ctx),
		k.GetDefaultExternalAddressEntries(ctx),
		k.GetApplicationLinks(ctx),
	)
}

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) []abci.ValidatorUpdate {
	// Initialize the module params
	k.SetParams(ctx, data.Params)

	// Initialize the IBC settings
	k.SetPort(ctx, data.IBCPortID)

	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, data.IBCPortID) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, data.IBCPortID)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}

	// Initialize the Profiles
	k.ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {
		if profile, ok := (account).(*types.Profile); ok {
			err := k.SaveProfile(ctx, profile)
			if err != nil {
				panic(err)
			}
		}
		return false
	})

	// Store the transfer requests
	for _, request := range data.DTagTransferRequests {
		err := k.SaveDTagTransferRequest(ctx, request)
		if err != nil {
			panic(err)
		}
	}

	for _, link := range data.ChainLinks {
		err := k.SaveChainLink(ctx, link)
		if err != nil {
			panic(err)
		}
	}

	for _, entry := range data.DefaultExternalAddresses {
		k.SaveDefaultExternalAddress(ctx, entry.Owner, entry.ChainName, entry.Target)
	}

	// Store the application links
	for _, link := range data.ApplicationLinks {
		err := k.SaveApplicationLink(ctx, link)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
