package types

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/desmos-labs/desmos/v2/app"
	"github.com/desmos-labs/desmos/v2/app/desmos/cmd/chainlink/types/multi"
	"github.com/desmos-labs/desmos/v2/app/desmos/cmd/chainlink/types/single"
	"github.com/desmos-labs/desmos/v2/x/profiles/client/utils"
	profilescliutils "github.com/desmos-labs/desmos/v2/x/profiles/client/utils"
	profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"
)

const (
	KeyName = "desmos_chain_link_account"
)

type ChainLinkJSONBuilder interface {
	BuildChainLinkJSON(chain Chain) (utils.ChainLinkJSON, error)
}

type ChainLinkJSONBuilderProvider func(isSingleAccount bool) ChainLinkJSONBuilder

func DefaultChainLinkJSONBuilderProvider(isSingleAccount bool) ChainLinkJSONBuilder {
	if isSingleAccount {
		return NewSingleAccountChainLinkJSONBuilder(single.NewChainLinkJSONInfoGetter())
	}
	return NewMultisigAccountChainLinkJSONBuilder(multi.NewChainLinkJSONInfoGetter())
}

// -------------------------------------------------------------------------------------------------------------------

type SingleAccountChainLinkJSONBuilder struct {
	getter SingleSignatureAccountReferenceGetter
}

func NewSingleAccountChainLinkJSONBuilder(getter SingleSignatureAccountReferenceGetter) *SingleAccountChainLinkJSONBuilder {
	return &SingleAccountChainLinkJSONBuilder{
		getter: getter,
	}
}

// BuildChainLinkJSON implements ChainLinkJSONBuilder
func (b *SingleAccountChainLinkJSONBuilder) BuildChainLinkJSON(chain Chain) (utils.ChainLinkJSON, error) {
	mnemonic, err := b.getter.GetMnemonic()
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	// Create an in-memory keybase for signing
	keyBase := keyring.NewInMemory()
	_, err = keyBase.NewAccount(KeyName, mnemonic, "", chain.DerivationPath, hd.Secp256k1)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	// Generate the proof signing it with the key
	key, _ := keyBase.Key(KeyName)
	addr, _ := sdk.Bech32ifyAddressBytes(chain.Prefix, key.GetAddress())
	value := []byte(addr)
	sig, pubkey, err := keyBase.Sign(KeyName, value)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}
	sigData := &profilestypes.SingleSignatureData{
		Mode:      signing.SignMode_SIGN_MODE_DIRECT,
		Signature: sig,
	}

	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(pubkey, sigData, hex.EncodeToString(value)),
		profilestypes.NewChainConfig(chain.Name),
	), nil
}

// -------------------------------------------------------------------------------------------------------------------

type MultisigAccountChainLinkJSONBuilder struct {
	getter MultiSignatureAccountReferenceGetter
}

func NewMultisigAccountChainLinkJSONBuilder(getter MultiSignatureAccountReferenceGetter) *MultisigAccountChainLinkJSONBuilder {
	return &MultisigAccountChainLinkJSONBuilder{
		getter: getter,
	}
}

// BuildChainLinkJSON implements ChainLinkJSONBuilder
func (b *MultisigAccountChainLinkJSONBuilder) BuildChainLinkJSON(chain Chain) (utils.ChainLinkJSON, error) {
	txFilePath, err := b.getter.GetMultiSignedTxFilePath()
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	signedChainID, err := b.getter.GetSignedChainID()
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	encodingConfig := app.MakeTestEncodingConfig()
	txCfg := encodingConfig.TxConfig

	// Read the transaction file
	bytes, err := ioutil.ReadFile(txFilePath)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	// Parse the transaction
	parsedTx, err := txCfg.TxJSONDecoder()(bytes)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	// Get the sign mode
	signMode := signing.SignMode_SIGN_MODE_DIRECT
	if _, ok := parsedTx.(legacytx.StdTx); ok {
		signMode = signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
	}

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

	// Re-create the bytes that have been signed in order to produce the signature
	signingData := authsigning.SignerData{AccountNumber: 0, Sequence: 0, ChainID: signedChainID}
	value, err := txCfg.SignModeHandler().GetSignBytes(signMode, signingData, parsedTx)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	addr, err := sdk.Bech32ifyAddressBytes(chain.Prefix, sigs[0].PubKey.Address().Bytes())
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	sigData, err := profilestypes.SignatureDataFromCosmosSignatureData(sigs[0].Data)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(sigs[0].PubKey, sigData, hex.EncodeToString(value)),
		profilestypes.NewChainConfig(chain.Name),
	), err
}
