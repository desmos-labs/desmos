package types

// Chain contains the data of a single chain
type Chain struct {
	ID             string
	Name           string
	Prefix         string
	DerivationPath string
}

// NewChain returns a new Chain instance
func NewChain(id, name, prefix, derivationPath string) Chain {
	return Chain{
		id,
		name,
		prefix,
		derivationPath,
	}
}
