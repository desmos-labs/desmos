package types

import "regexp"

const (
	ModuleName = "profile"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateAccount = "create_account"
	ActionEditAccount   = "edit_account"
	ActionDeleteAccount = "delete_account"

	//Queries
	QuerierRoute  = ModuleName
	QueryAccount  = "account"
	QueryAccounts = "accounts"
)

var (
	TxHashRegEx = regexp.MustCompile("^[a-fA-F0-9]{64}$")

	AccountStorePrefix = []byte("accounts")
)

// AccountStoreKey turns a moniker to a key used to store an account into the accounts store
func AccountStoreKey(moniker string) []byte {
	return append(AccountStorePrefix, []byte(moniker)...)
}
