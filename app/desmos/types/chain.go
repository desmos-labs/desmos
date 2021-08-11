package types

// ChainType contains the data of the chain
type ChainType struct {
	ID             string
	Name           string
	Prefix         string
	DerivationPath string
}

// NewChainType returns a new ChainType instance
func NewChainType(id, name, prefix, derivationPath string) ChainType {
	return ChainType{
		id,
		name,
		prefix,
		derivationPath,
	}
}
