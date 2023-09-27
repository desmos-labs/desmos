package builder

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/desmos-labs/desmos/v6/app/desmos/cmd/chainlink/types"
	"github.com/desmos-labs/desmos/v6/x/profiles/client/utils"
)

// ChainLinkJSONBuilder allows to build a ChainLinkJSON instance
type ChainLinkJSONBuilder interface {
	BuildChainLinkJSON(ctx client.Context, chain types.Chain) (utils.ChainLinkJSON, error)
}

// ChainLinkJSONBuilderProvider allows to provide the provider ChainLinkJSONBuilder implementation based on whether
// it should create the JSON chain link for single or
type ChainLinkJSONBuilderProvider func(owner string, isSingleAccount bool) ChainLinkJSONBuilder
