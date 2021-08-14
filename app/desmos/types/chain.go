package types

// Chain contains the data of the chain
type Chain struct {
	ID             string
	Name           string
	Prefix         string
	DerivationPath string
}

// NewChain returns a new ChainType instance
func NewChain(id, name, prefix, derivationPath string) Chain {
	return Chain{
		id,
		name,
		prefix,
		derivationPath,
	}
}
