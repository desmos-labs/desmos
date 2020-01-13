package types

import (
	"encoding/json"
)

// PostQueryResponse represents the data of a post
// that is returned to user upon a query
type PostQueryResponse struct {
	Type      string    `json:"type"`
	Post      Post      `json:"post"`
	Reactions Reactions `json:"reactions"`
	Children  PostIDs   `json:"children"`
}

func (response PostQueryResponse) String() string {
	bytes, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func NewPostResponse(post Post, reactions Reactions, children PostIDs) PostQueryResponse {
	var ttype string

	if _, ok := post.(TextPost); ok {
		ttype = "TextPost"
	}

	if _, ok := post.(MediaPost); ok {
		ttype = "MediaPost"
	}

	return PostQueryResponse{
		Type:      ttype,
		Post:      post,
		Reactions: reactions,
		Children:  children,
	}
}

// MarshalJSON implements json.Marshaler as Amino does
// not respect default json composition
func (response PostQueryResponse) MarshalJSON() ([]byte, error) {
	type temp PostQueryResponse
	return json.Marshal(temp(response))
}

// UnmarshalJSON implements json.Unmarshaler as Amino does
// not respect default json composition
func (response *PostQueryResponse) UnmarshalJSON(data []byte) error {
	type postResponse PostQueryResponse
	var temp postResponse
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*response = PostQueryResponse(temp)
	return nil
}
