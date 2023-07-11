package types

import (
	"github.com/cosmos/gogoproto/proto"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var _ codectypes.InterfaceRegistry = &WasmInterfaceRegistry{}

// NewWasmInterfaceRegistry returns a new WasmInterfaceRegistry instance
func NewWasmInterfaceRegistry(registry codectypes.InterfaceRegistry) WasmInterfaceRegistry {
	return WasmInterfaceRegistry{registry}
}

// WasmInterfaceRegistry represents a interface registry with a custom any resolver
type WasmInterfaceRegistry struct {
	codectypes.InterfaceRegistry
}

// Resolve implements jsonpb.AnyResolver of codectypes.InterfaceRegistry
func (WasmInterfaceRegistry) Resolve(typeURL string) (proto.Message, error) {
	return new(WasmAny), nil
}
