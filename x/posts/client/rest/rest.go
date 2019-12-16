package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

// REST Variable names
// nolint
const (
	RestCreator      = "creator"
	RestParentID     = "parent-id"
	RestCreationTime = "creation-time"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerTxRoutes(cliCtx, r)
	registerQueryRoutes(cliCtx, r)
}

// CreatePostReq defines the properties of a post request's body.
type CreatePostReq struct {
	BaseReq           rest.BaseReq `json:"base_req"`
	Message           string       `json:"message"`
	ParentID          string       `json:"parent_id"`
	AllowsComments    bool         `json:"allows_comments"`
	ExternalReference string       `json:"external_reference"`
}

// CreatePostReq defines the properties of a like request's body.
type AddLikeReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	PostID  string       `json:"post_id"`
}

// CreatePostReq defines the properties of a unlike request's body.
type RemoveLikeReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	PostID  string       `json:"post_id"`
}
