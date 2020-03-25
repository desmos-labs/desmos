package types

import "regexp"

const (
	ModuleName = "profile"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateAccount = "create_account"
	ActionEditAccount   = "edit_account"

	//Queries
	QuerierRoute = ModuleName
)

var (
	TxHashRegEx = regexp.MustCompile("^[a-fA-F0-9]{64}$")
)
