package provider

import (
	"github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink/builder"
	multibuilder "github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink/builder/multi"
	singlebuilder "github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink/builder/single"
	multigetter "github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink/getter/multi"
	singlegetter "github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink/getter/single"
)

// DefaultChainLinkJSONBuilderProvider returns the default ChainLinkJSONBuilder provider implementation
func DefaultChainLinkJSONBuilderProvider(owner string, isSingleAccount bool) builder.ChainLinkJSONBuilder {
	if isSingleAccount {
		return singlebuilder.NewAccountChainLinkJSONBuilder(owner, singlegetter.NewChainLinkJSONInfoGetter())
	}
	return multibuilder.NewAccountChainLinkJSONBuilder(multigetter.NewChainLinkJSONInfoGetter())
}
