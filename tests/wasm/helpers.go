package wasm

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm/types"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"

	chaintesting "github.com/cosmos/ibc-go/v8/testing"
)

type TestChain struct {
	*chaintesting.TestChain
}

// StoreCodeFile compresses code file then stores its byte code on chain
func (chain *TestChain) StoreCodeFile(filename string) types.MsgStoreCodeResponse {
	wasmCode, err := os.ReadFile(filename)
	require.NoError(chain, err)
	if strings.HasSuffix(filename, "wasm") { // compress for gas limit
		var buf bytes.Buffer
		gz := gzip.NewWriter(&buf)
		_, err := gz.Write(wasmCode)
		require.NoError(chain, err)
		err = gz.Close()
		require.NoError(chain, err)
		wasmCode = buf.Bytes()
	}
	return chain.StoreCode(wasmCode)
}

// StoreCode stores byte code on chain
func (chain *TestChain) StoreCode(byteCode []byte) types.MsgStoreCodeResponse {
	storeMsg := &types.MsgStoreCode{
		Sender:       chain.SenderAccount.GetAddress().String(),
		WASMByteCode: byteCode,
	}
	r, err := chain.SendMsgs(storeMsg)
	require.NoError(chain, err)

	var pInstResp types.MsgStoreCodeResponse
	chain.UnwrapExecTXResult(r, &pInstResp)

	require.NotEmpty(chain, pInstResp.CodeID)
	require.NotEmpty(chain, pInstResp.Checksum)
	return pInstResp
}

// InstantiateContract instantiates contract by the given code ID
func (chain *TestChain) InstantiateContract(codeID uint64, initMsg []byte) sdk.AccAddress {
	instantiateMsg := &types.MsgInstantiateContract{
		Sender: chain.SenderAccount.GetAddress().String(),
		Admin:  chain.SenderAccount.GetAddress().String(),
		CodeID: codeID,
		Label:  "any-resolver-test",
		Msg:    initMsg,
		Funds:  sdk.Coins{chaintesting.TestCoin},
	}

	r, err := chain.SendMsgs(instantiateMsg)
	require.NoError(chain, err)

	var pExecResp types.MsgInstantiateContractResponse
	chain.UnwrapExecTXResult(r, &pExecResp)

	a, err := sdk.AccAddressFromBech32(pExecResp.Address)
	require.NoError(chain, err)
	return a
}

// UnwrapExecTXResult is a helper to unpack execution result from proto any type
func (chain *TestChain) UnwrapExecTXResult(r *abci.ExecTxResult, target proto.Message) {
	var wrappedRsp sdk.TxMsgData
	require.NoError(chain, chain.Codec.Unmarshal(r.Data, &wrappedRsp))

	// unmarshal protobuf response from data
	require.Len(chain, wrappedRsp.MsgResponses, 1)
	require.NoError(chain, proto.Unmarshal(wrappedRsp.MsgResponses[0].Value, target))
}

// SmartQuery This will serialize the query message and submit it to the contract.
// The response is parsed into the provided interface.
// Usage: SmartQuery(addr, QueryMsg{Foo: 1}, &response)
func (chain *TestChain) SmartQuery(contractAddr string, queryMsg interface{}, response interface{}) error {
	msg, err := json.Marshal(queryMsg)
	if err != nil {
		return err
	}

	req := types.QuerySmartContractStateRequest{
		Address:   contractAddr,
		QueryData: msg,
	}
	reqBin := chain.Codec.MustMarshal(&req)

	res, err := chain.App.Query(chain.GetContext(), &abci.RequestQuery{
		Path: "/cosmwasm.wasm.v1.Query/SmartContractState",
		Data: reqBin,
	})
	if err != nil {
		return err
	}

	if res.Code != 0 {
		return fmt.Errorf("query failed: (%d) %s", res.Code, res.Log)
	}

	// unpack protobuf
	var resp types.QuerySmartContractStateResponse
	chain.Codec.MustUnmarshal(res.Value, &resp)

	// unpack json content
	return json.Unmarshal(resp.Data, response)
}

// InstantiateTestContract store and instantiate a test contract instance
func InstantiateTestContract(t *testing.T, chain *TestChain) sdk.AccAddress {
	codeID := chain.StoreCodeFile("./testdata/test_contract.wasm").CodeID
	contractAddr := chain.InstantiateContract(codeID, []byte(`{}`))
	require.NotEmpty(t, contractAddr)
	return contractAddr
}
