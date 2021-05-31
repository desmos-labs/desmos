package keeper_test

import (
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"

	"time"
)

func (suite *KeeperTestsuite) TestKeeper_IterateSubspace() {
	date, err := time.Parse(time.RFC3339, "2010-10-02T12:10:00.000Z")
	suite.NoError(err)

	subspaces := []*types.Subspace{
		{
			ID:           "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			Name:         "mooncake",
			Owner:        "cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			Creator:      "cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			CreationTime: date,
			Type:         types.Open,
		},
		{
			ID:           "dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
			Name:         "star",
			Owner:        "cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			CreationTime: date,
			Type:         types.Open,
		},
		{
			ID:           "e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
			Name:         "bad",
			Owner:        "cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			CreationTime: date,
			Type:         types.Open,
		},
	}

	expSubspaces := []*types.Subspace{
		subspaces[0],
		subspaces[1],
	}

	for _, subspace := range subspaces {
		suite.k.SaveSubspace(suite.ctx, *subspace)
	}

	var validSubspaces []*types.Subspace
	suite.k.IterateSubspaces(suite.ctx, func(_ int64, subspace types.Subspace) (stop bool) {
		if subspace.Name == "bad" {
			return false
		}
		validSubspaces = append(validSubspaces, &subspace)
		return false
	})

	suite.Len(expSubspaces, len(validSubspaces))
	for _, subspace := range validSubspaces {
		suite.Contains(expSubspaces, subspace)
	}
}
