package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

const (
	DenomParam           = "denom"
	DividerExponentParam = "divider-exponent"
)

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	r := rtr.NewRoute().Subrouter()
	registerQueryRoutes(clientCtx, r)
}
