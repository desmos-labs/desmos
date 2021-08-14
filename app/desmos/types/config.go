package types

// Config contains the data of the configuration
type Config struct {
	Chains   []Chain
	FilePath string
}

// DefaultConfig returns the default config instance of the configuration
func DefaultConfig() Config {
	return Config{
		Chains: []Chain{
			NewChain("Desmos", "desmos", "desmos", "m/44'/852'/0'/0/0"),
			NewChain("Cosmos", "cosmos", "cosmos", "m/44'/118'/0'/0/0"),
			NewChain("Akash", "akash", "akash", "m/44'/118'/0'/0/0"),
			NewChain("Osmosis", "osmosis", "osmo", "m/44'/118'/0'/0/0"),
			NewChain("Other", "", "", ""),
		},
	}
}
