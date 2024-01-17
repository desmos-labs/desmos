package wasm_test

import (
	"encoding/base64"
	"strings"
	"testing"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"

	ibctesting "github.com/cosmos/ibc-go/v8/testing"

	"github.com/desmos-labs/desmos/v6/tests/wasm"
	desmosibctesting "github.com/desmos-labs/desmos/v6/testutil/ibctesting"
	profilestypes "github.com/desmos-labs/desmos/v6/x/profiles/types"
	wasmtypes "github.com/desmos-labs/desmos/v6/x/wasm/types"
)

func TestAnyResolverByProfile(t *testing.T) {
	// Create a test chain
	coord := desmosibctesting.NewCoordinator(t, 1)
	chain := &wasm.TestChain{coord.GetChain(ibctesting.GetChainID(1))}

	// Store and instantiate test contract
	contractAddr := wasm.InstantiateTestContract(t, chain)

	// Save a profile
	saveProfileMsg := profilestypes.NewMsgSaveProfile(
		"test_user", "test_user", "test bio", "https://profile.pic", "https://cover.pic", chain.SenderAccount.GetAddress().String(),
	)
	_, err := chain.SendMsgs(saveProfileMsg)
	require.NoError(t, err)

	// Request profile via test contract
	profileReq := profilestypes.QueryProfileRequest{
		User: "test_user",
	}

	var wasmRes wasm.MsgTestContractQueryResponse
	err = chain.SmartQuery(contractAddr.String(), wasm.MsgTestContractQueryRequest{
		DesmosChain: wasm.DesmosChainMsg{
			Request: wasmvmtypes.QueryRequest{
				Stargate: &wasmvmtypes.StargateQuery{
					Path: "/desmos.profiles.v3.Query/Profile",
					Data: chain.Codec.MustMarshal(&profileReq),
				},
			},
		},
	}, &wasmRes)
	require.NoError(t, err)

	base64Encoded, err := wasmRes.Data.MarshalJSON()
	require.NoError(t, err)

	// Decode base64-encoded res into JSON res
	res, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(string(base64Encoded), "\"", ""))
	require.NoError(t, err)

	// Build stargate codec with custom any resolver
	protoCdc, ok := chain.Codec.(*codec.ProtoCodec)
	require.True(t, ok)
	stargateCdc := codec.NewProtoCodec(wasmtypes.NewWasmInterfaceRegistry(protoCdc.InterfaceRegistry()))

	// Build expected profile response
	expProfile, err := profilestypes.NewProfile(
		"test_user",
		"test_user",
		"test bio",
		profilestypes.NewPictures(
			"https://profile.pic", "https://cover.pic",
		),
		chain.LastHeader.GetTime(), // creation time is the latest block time
		chain.SenderAccount,
	)
	require.NoError(t, err)

	// Encode expected profile response by stargate codec
	expRes := stargateCdc.MustMarshalJSON(&profilestypes.QueryProfileResponse{Profile: codectypes.UnsafePackAny(expProfile)})

	var r profilestypes.QueryProfileResponse
	stargateCdc.MustUnmarshalJSON(expRes, &r)

	require.Equal(t, expRes, res)
}
