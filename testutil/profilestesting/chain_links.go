package profilestesting

import (
	"encoding/hex"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/profiles/types"
)

type ChainLinkAccount struct {
	privKey      cryptotypes.PrivKey
	pubKey       cryptotypes.PubKey
	chainName    string
	bech32Prefix string
}

func GetChainLinkAccount(chainName string, bech32Prefix string) ChainLinkAccount {
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	return ChainLinkAccount{
		privKey:      privKey,
		pubKey:       pubKey,
		chainName:    chainName,
		bech32Prefix: bech32Prefix,
	}
}

func (a ChainLinkAccount) ChainName() string {
	return a.chainName
}

func (a ChainLinkAccount) PubKey() cryptotypes.PubKey {
	return a.pubKey
}

func (a ChainLinkAccount) PubKeyAny() *codectypes.Any {
	return NewAny(a.pubKey)
}

func (a ChainLinkAccount) Sign(value string) []byte {
	bech32Sig, _ := a.privKey.Sign([]byte(value))
	return bech32Sig
}

func (a ChainLinkAccount) Bech32Address() *types.Bech32Address {
	addr, _ := sdk.Bech32ifyAddressBytes(a.bech32Prefix, a.pubKey.Address())
	return types.NewBech32Address(addr, a.bech32Prefix)
}

func (a ChainLinkAccount) Bech32SignatureData(signedValue string) types.Signature {
	return types.NewSingleSignature(types.SIGNATURE_VALUE_TYPE_RAW, a.Sign(signedValue))
}

func (a ChainLinkAccount) Bech32Proof(user string) types.Proof {
	return types.NewProof(
		a.pubKey,
		a.Bech32SignatureData(user),
		hex.EncodeToString([]byte(user)),
	)
}

func (a ChainLinkAccount) GetBech32ChainLink(user string, date time.Time) types.ChainLink {
	return types.NewChainLink(
		user,
		a.Bech32Address(),
		a.Bech32Proof(user),
		types.NewChainConfig(a.chainName),
		date,
	)
}
