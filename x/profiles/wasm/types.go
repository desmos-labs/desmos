package wasm

import "github.com/desmos-labs/desmos/v2/x/profiles/types"

type ProfilesQueryRoutes struct {
	Profile *ProfileQuery `json:"profile"`
}

type ProfileQuery struct {
	Request *types.QueryProfileRequest `json:"request"`
}
