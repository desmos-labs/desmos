package rest

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/desmos-labs/desmos/x/posts/types"
)

const (
	ParamPostId = "post_id"

	ParamSortBy       = "sort_by"
	ParamSortOrder    = "sort_order"
	ParamParentId     = "parent_id"
	ParamCreationTime = "creation_time"
	ParamSubspace     = "subspace"
	ParamCreator      = "creator"
	ParamHashtags     = "hashtags"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}

// CreatePostReq defines the properties of a post creation request's body.
type CreatePostReq struct {
	BaseReq        rest.BaseReq       `json:"base_req"`
	Message        string             `json:"message"`
	ParentId       string             `json:"parent_id"`
	AllowsComments bool               `json:"allows_comments"`
	Subspace       string             `json:"subspace"`
	OptionalData   types.OptionalData `json:"optional_data"`
	CreationTime   time.Time          `json:"creation_time"`
	Medias         types.Attachments  `json:"attachments,omitempty"`
	PollData       *types.PollData    `json:"poll_data,omitempty"`
}

// AddReactionReq defines the properties of a reaction adding request's body.
type AddReactionReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	PostId   string       `json:"post_id"`
	Reaction string       `json:"reaction"`
}

// RemoveReactionReq defines the properties of a reaction removal request's body.
type RemoveReactionReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	PostId   string       `json:"post_id"`
	Reaction string       `json:"reaction"`
}

type AnswerPollPostReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Answers []string     `json:"answers"`
}

type RegisterReactionReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Shortcode string       `json:"shortcode"`
	Value     string       `json:"value"`
	Subspace  string       `json:"subspace"`
}
