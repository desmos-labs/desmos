package models

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type RelationshipsResponse struct {
	Relationships []sdk.AccAddress `json:"relationships,omitempty" yaml:"relationships,omitempty"`
}

func NewRelationshipResponse(relationships []sdk.AccAddress) RelationshipsResponse {
	return RelationshipsResponse{relationships}
}

func (response RelationshipsResponse) String() string {
	out := fmt.Sprintf(`
Relationships: %s`, response.Relationships)
	return strings.TrimSpace(out)
}
