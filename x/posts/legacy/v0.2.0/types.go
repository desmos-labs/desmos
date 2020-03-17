package v020

import (
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "posts"
)

// GenesisState contains the data of a v0.2.0 genesis state for the posts module
type GenesisState struct {
	Posts     []Post                `json:"posts"`
	Reactions map[string][]Reaction `json:"reactions"`
}

// PostID represents a unique post id
type PostID uint64

// String implements fmt.Stringer
func (id PostID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

type OptionalData map[string]string

// Post is a struct of a post
type Post struct {
	PostID         PostID         `json:"id"`                      // Unique id
	ParentID       PostID         `json:"parent_id"`               // Post of which this one is a comment
	Message        string         `json:"message"`                 // Message contained inside the post
	Created        time.Time      `json:"created"`                 // RFC3339 date at which the post has been created
	LastEdited     time.Time      `json:"last_edited"`             // RFC3339 date at which the post has been edited the last time
	AllowsComments bool           `json:"allows_comments"`         // Tells if users can reference this PostID as the parent
	Subspace       string         `json:"subspace"`                // Identifies the application that has posted the message
	OptionalData   OptionalData   `json:"optional_data,omitempty"` // Arbitrary data that can be used from the developers
	Creator        sdk.AccAddress `json:"creator"`                 // Creator of the Post
}

func (p Post) ConflictsWith(other Post) bool {
	return p.Created.Equal(other.Created) &&
		p.Subspace == other.Subspace &&
		p.Creator.Equals(other.Creator)
}

// ContentsEquals returns true if and only if p and other contain the same data, without considering the ID
func (p Post) ContentsEquals(other Post) bool {
	equalsOptionalData := len(p.OptionalData) == len(other.OptionalData)
	if equalsOptionalData {
		for key := range p.OptionalData {
			equalsOptionalData = equalsOptionalData && p.OptionalData[key] == other.OptionalData[key]
		}
	}

	return p.ParentID == other.ParentID &&
		p.Message == other.Message &&
		p.Created.Equal(other.Created) &&
		p.LastEdited.Equal(other.LastEdited) &&
		p.AllowsComments == other.AllowsComments &&
		p.Subspace == other.Subspace &&
		equalsOptionalData &&
		p.Creator.Equals(other.Creator)
}

// PostReaction is a struct of a user reaction to a post
type Reaction struct {
	Owner sdk.AccAddress `json:"owner"` // User that has created the reaction
	Value string         `json:"value"` // PostReaction of the reaction
}
