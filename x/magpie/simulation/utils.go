package simulation

// DONTCOVER

import (
	"encoding/base64"
	"math/rand"

	secp256k1 "github.com/btcsuite/btcd/btcec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tendermint/tendermint/libs/bech32"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

var (
	RandomNamespaces = []string{"cosmos", "iris", "kava", "regen"}
)

type SessionData struct {
	Owner         simulation.Account
	Namespace     string
	ExternalOwner string
	PubKey        string
	Signature     string
}

func RandomSessionData(simAccount simulation.Account, r *rand.Rand) SessionData {
	namespace := RandomNamespaces[r.Intn(len(RandomNamespaces))]

	extOwner, err := bech32.ConvertAndEncode(namespace, simAccount.Address.Bytes())
	if err != nil {
		panic(err) // Shouldn't happen
	}

	_, pubkeyObject := secp256k1.PrivKeyFromBytes(secp256k1.S256(), simAccount.PrivKey.Bytes())
	bytes := pubkeyObject.SerializeCompressed()
	extPubKey := base64.StdEncoding.EncodeToString(bytes)

	// Create the signature data
	msg := types.NewMsgCreateSession(simAccount.Address, namespace, extOwner, extPubKey, "")
	signBytes := auth.StdSignBytes(namespace, 0, 0, auth.NewStdFee(200000, nil), []sdk.Msg{msg}, "")

	// Create the signature
	signedBytes, err := simAccount.PrivKey.Sign(signBytes)
	if err != nil {
		panic(err)
	}

	// Create the session data
	return SessionData{
		Owner:         simAccount,
		Namespace:     namespace,
		ExternalOwner: extOwner,
		PubKey:        extPubKey,
		Signature:     base64.StdEncoding.EncodeToString(signedBytes),
	}
}
