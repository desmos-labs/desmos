package types

type SubspacesQueryRoutes struct {
	Subspaces SubspacesQueryRequest `json:"subspaces"`
}

type SubspacesQueryRequest struct {
	Subspaces *QuerySubspacesRequest `json:"subspaces"`
	Subspace  *QuerySubspaceRequest  `json:"subspace"`
}
