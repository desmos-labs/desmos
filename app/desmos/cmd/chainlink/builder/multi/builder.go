package multi

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/desmos-labs/desmos/v7/app"
	"github.com/desmos-labs/desmos/v7/app/desmos/cmd/chainlink/getter"
	"github.com/desmos-labs/desmos/v7/app/desmos/cmd/chainlink/types"
	"github.com/desmos-labs/desmos/v7/x/profiles/client/utils"
	profilestypes "github.com/desmos-labs/desmos/v7/x/profiles/types"
)

// AccountChainLinkJSONBuilder implements the ChainLinkJSONBuilder for multi signature accounts
type AccountChainLinkJSONBuilder struct {
	getter getter.MultiSignatureAccountReferenceGetter
}

// NewAccountChainLinkJSONBuilder returns a new AccountChainLinkJSONBuilder instance
func NewAccountChainLinkJSONBuilder(getter getter.MultiSignatureAccountReferenceGetter) *AccountChainLinkJSONBuilder {
	return &AccountChainLinkJSONBuilder{
		getter: getter,
	}
}

// BuildChainLinkJSON implements ChainLinkJSONBuilder
func (b *AccountChainLinkJSONBuilder) BuildChainLinkJSON(_ codec.Codec, chain types.Chain) (utils.ChainLinkJSON, error) {
	txFilePath, err := b.getter.GetMultiSignedTxFilePath()
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	signedChainID, err := b.getter.GetSignedChainID()
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	encodingConfig := app.MakeEncodingConfig()
	txCfg := encodingConfig.TxConfig

	// Read the transaction file
	bytes, err := os.ReadFile(txFilePath)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	// Parse the transaction
	parsedTx, err := txCfg.TxJSONDecoder()(bytes)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	// Get the sign mode
	signMode := signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON

	// Wrap the transaction inside a builder to make it easier to get the signatures
	txBuilder, err := txCfg.WrapTxBuilder(parsedTx)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	sigs, err := txBuilder.GetTx().GetSignaturesV2()
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	// Make sure there is only one signature for the multisig account
	if len(sigs) != 1 {
		return utils.ChainLinkJSON{}, fmt.Errorf("invalid number of signatures")
	}

	addr, err := sdk.Bech32ifyAddressBytes(chain.Prefix, sigs[0].PubKey.Address().Bytes())
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	// Re-create the bytes that have been signed in order to produce the signature
	signingData := authsigning.SignerData{
		AccountNumber: 0,
		Sequence:      0,
		ChainID:       signedChainID,
		Address:       addr,
		PubKey:        sigs[0].PubKey,
	}
	value, err := txCfg.SignModeHandler().GetSignBytes(signMode, signingData, parsedTx)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	sigData, err := profilestypes.CosmosSignatureDataToSignature(sigs[0].Data)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	return utils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(sigs[0].PubKey, sigData, hex.EncodeToString(value)),
		profilestypes.NewChainConfig(chain.Name),
	), err
}
