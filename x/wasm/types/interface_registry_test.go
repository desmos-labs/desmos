package types_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/desmos-labs/desmos/v5/app"
	"github.com/desmos-labs/desmos/v5/x/wasm/types"
	"github.com/stretchr/testify/require"
)

func TestWasmInterfaceRegistry_Resolve(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	protoCdc := cdc.(*codec.ProtoCodec)
	registry := types.NewWasmInterfaceRegistry(protoCdc.InterfaceRegistry())
	message, err := registry.Resolve("/cosmos.auth.v1beta1.BaseAccount")
	require.NoError(t, err)
	require.Equal(t, new(types.WasmAny), message)
}

func TestWasmInterfaceRegistry_MarshalJSON(t *testing.T) {
	account := authtypes.NewBaseAccountWithAddress(sdk.MustAccAddressFromBech32("cosmos1f8uxultn8sqzhznrsz3q77xwaquhgrsg6jyvfy"))

	// testData must have a field required *codectypes.Any type
	testData := authtypes.QueryAccountResponse{
		Account: codectypes.UnsafePackAny(account),
	}

	cdc, _ := app.MakeCodecs()
	protoCdc := cdc.(*codec.ProtoCodec)
	registry := types.NewWasmInterfaceRegistry(protoCdc.InterfaceRegistry())
	bz, err := codec.ProtoMarshalJSON(&testData, registry)
	require.NoError(t, err)

	expected := "{\"account\":{\"@type\":\"/cosmos.auth.v1beta1.BaseAccount\",\"value\":\"Ci1jb3Ntb3MxZjh1eHVsdG44c3F6aHpucnN6M3E3N3h3YXF1aGdyc2c2anl2Znk=\"}}"
	require.JSONEq(t, expected, string(bz))
}
