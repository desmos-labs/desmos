package types

const (
	ModuleName = "fees"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionUpdateParams = "update_params"
)

var (
	ParamsKey = []byte{0x01}
)
