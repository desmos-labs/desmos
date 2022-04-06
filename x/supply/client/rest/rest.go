package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

const (
	DenomParam = "denom"
)

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	r := clientrest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
}

type SupplyReq struct {
	BaseReq         rest.BaseReq `json:"base_req" yaml:"base_req"`
	DividerExponent string       `json:"divider_exponent" yaml:"divider_exponent"` // DividerExponent is a factor used to power the divider used to convert the supply to the desired representation
}
