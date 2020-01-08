package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerTxRoutes(cliCtx, r)
	registerQueryRoutes(cliCtx, r)
}

// CreatePostReq defines the properties of a post creation request's body.
type CreatePostReq struct {
	BaseReq        rest.BaseReq      `json:"base_req"`
	Message        string            `json:"message"`
	ParentID       string            `json:"parent_id"`
	AllowsComments bool              `json:"allows_comments"`
	Subspace       string            `json:"subspace"`
	OptionalData   map[string]string `json:"optional_data"`
}

// AddReactionReq defines the properties of a reaction adding request's body.
type AddReactionReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	PostID   string       `json:"post_id"`
	Reaction string       `json:"reaction"`
}

// RemoveReactionReq defines the properties of a reaction removal request's body.
type RemoveReactionReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	PostID   string       `json:"post_id"`
	Reaction string       `json:"reaction"`
}
