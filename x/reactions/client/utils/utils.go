package utils

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
)

// ParseSetReactionsParamsJSON reads and parses a SetReactionsParamsJSON from file.
func ParseSetReactionsParamsJSON(cdc codec.Codec, dataFile string) (SetReactionsParamsJSON, error) {
	var data SetReactionsParamsJSON

	contents, err := os.ReadFile(dataFile)
	if err != nil {
		return data, err
	}

	if err := cdc.UnmarshalJSON(contents, &data); err != nil {
		return data, err
	}

	return data, nil
}
