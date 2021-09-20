package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewOracleRequestPacketData constructs a new OracleRequestPacketData instance
func NewOracleRequestPacketData(
	clientID string, oracleScriptID OracleScriptID, calldata []byte, askCount uint64, minCount uint64, feeLimit sdk.Coins, prepareGas uint64, executeGas uint64,
) OracleRequestPacketData {
	return OracleRequestPacketData{
		ClientID:       clientID,
		OracleScriptID: oracleScriptID,
		Calldata:       calldata,
		AskCount:       askCount,
		MinCount:       minCount,
		FeeLimit:       feeLimit,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
	}
}

// GetBytes is a helper for serialising
func (p OracleRequestPacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}
