package simulation

// DONTCOVER

import (
	"encoding/base64"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	legacyauth "github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	secp256k1 "github.com/btcsuite/btcd/btcec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

var (
	RandomNamespaces = []string{"cosmos", "iris", "kava", "regen"}
)

type SessionData struct {
	Owner         simtypes.Account
	Namespace     string
	ExternalOwner string
	PubKey        string
	Signature     string
}

func RandomSessionData(simAccount simtypes.Account, r *rand.Rand) SessionData {
	namespace := RandomNamespaces[r.Intn(len(RandomNamespaces))]

	extOwner, err := bech32.ConvertAndEncode(namespace, simAccount.Address.Bytes())
	if err != nil {
		panic(err) // Shouldn't happen
	}

	_, pubkeyObject := secp256k1.PrivKeyFromBytes(secp256k1.S256(), simAccount.PrivKey.Bytes())
	bytes := pubkeyObject.SerializeCompressed()
	extPubKey := base64.StdEncoding.EncodeToString(bytes)

	// Create the signature data
	msg := types.NewMsgCreateSession(simAccount.Address.String(), namespace, extOwner, extPubKey, "")

	//nolint:staticcheck
	signBytes := legacyauth.StdSignBytes(
		namespace,
		0,
		0,
		0,
		legacyauth.NewStdFee(200000, nil),
		[]sdk.Msg{msg},
		"",
	)

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
