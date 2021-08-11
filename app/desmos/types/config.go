package types

// Config contains the data of the configuration
type Config struct {
	Chains []ChainType
}

// DefaultConfig returns the default config instance of the configuration
func DefaultConfig() Config {
	return Config{
		Chains: []ChainType{
			NewChainType("Desmos", "desmos", "desmos", "m/44'/852'/0'/0/0"),
			NewChainType("Cosmos", "cosmos", "cosmos", "m/44'/118'/0'/0/0"),
			NewChainType("Akash", "akash", "akash", "m/44'/118'/0'/0/0"),
			NewChainType("Osmosis", "osmosis", "osmo", "m/44'/118'/0'/0/0"),
			NewChainType("Other", "", "", ""),
		},
	}
}
