package single

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	"github.com/desmos-labs/desmos/v6/app/desmos/cmd/chainlink/getter"
	"github.com/desmos-labs/desmos/v6/app/desmos/cmd/chainlink/types"
	"github.com/desmos-labs/desmos/v6/x/profiles/client/utils"
	profilestypes "github.com/desmos-labs/desmos/v6/x/profiles/types"
)

const (
	KeyName = "desmos_chain_link_account"
)

// AccountChainLinkJSONBuilder implements the ChainLinkJSONBuilder for single signature accounts
type AccountChainLinkJSONBuilder struct {
	owner  string
	getter getter.SingleSignatureAccountReferenceGetter
}

// NewAccountChainLinkJSONBuilder returns a new AccountChainLinkJSONBuilder instance
func NewAccountChainLinkJSONBuilder(owner string, getter getter.SingleSignatureAccountReferenceGetter) *AccountChainLinkJSONBuilder {
	return &AccountChainLinkJSONBuilder{
		owner:  owner,
		getter: getter,
	}
}

// BuildChainLinkJSON implements ChainLinkJSONBuilder
func (b *AccountChainLinkJSONBuilder) BuildChainLinkJSON(ctx client.Context, chain types.Chain) (utils.ChainLinkJSON, error) {
	mnemonic, err := b.getter.GetMnemonic()
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	// Create an in-memory keybase for signing
	keyBase := keyring.NewInMemory(ctx.Codec)
	_, err = keyBase.NewAccount(KeyName, mnemonic, "", chain.DerivationPath, hd.Secp256k1)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	// Generate the proof signing it with the key
	key, _ := keyBase.Key(KeyName)
	accAddr, _ := key.GetAddress()
	addr, _ := sdk.Bech32ifyAddressBytes(chain.Prefix, accAddr)
	value := []byte(b.owner)
	sig, pubkey, err := keyBase.Sign(KeyName, value, signing.SignMode_SIGN_MODE_TEXTUAL)
	if err != nil {
		return utils.ChainLinkJSON{}, err
	}

	return utils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(
			pubkey,
			profilestypes.NewSingleSignature(profilestypes.SIGNATURE_VALUE_TYPE_RAW, sig),
			hex.EncodeToString(value),
		),
		profilestypes.NewChainConfig(chain.Name),
	), nil
}
