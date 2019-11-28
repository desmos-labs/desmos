package keeper

import (
	"encoding/json"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// PostQueryResponse represents the data of a post
// that is returned to user upon a query
type PostQueryResponse struct {
	types.Post
	Likes    types.Likes   `json:"likes"`
	Children types.PostIDs `json:"children"`
}

func NewPostResponse(post types.Post, likes types.Likes, children types.PostIDs) PostQueryResponse {
	return PostQueryResponse{
		Post:     post,
		Likes:    likes,
		Children: children,
	}
}

// MarshalJSON implements json.Marshaler as Amino does
// not respect default json composition
func (response PostQueryResponse) MarshalJSON() ([]byte, error) {
	type temp PostQueryResponse
	return json.Marshal(temp(response))
}
