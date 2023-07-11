package wasm

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
)

// MsgTestContractQueryRequest represents the request type of test contract query
type MsgTestContractQueryRequest struct {
	DesmosChain DesmosChainMsg `json:"desmos_chain"`
}

// DesmosChainMsg represents the request message of Desmos for test contract query
type DesmosChainMsg struct {
	Request wasmvmtypes.QueryRequest `json:"request"`
}

// MsgTestContractQueryResponse represents the response type of test contract query
type MsgTestContractQueryResponse struct {
	Data json.RawMessage `json:"data"`
}
