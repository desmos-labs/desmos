package ante_test

import (
	"errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	cosmosante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	app2 "github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/ante"
	feesTypes "github.com/desmos-labs/desmos/x/fees/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/desmos-labs/desmos/x/posts/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// returns context and app with params set on account keeper
func createTestApp(isCheckTx bool, isBlockZero bool) (*app2.DesmosApp, sdk.Context) {
	db := dbm.NewMemDB()
	app := app2.NewDesmosApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, 0)

	header := abci.Header{}

	if !isBlockZero {
		header.Height = 1
	}

	ctx := app.BaseApp.NewContext(isCheckTx, header)
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.FeesKeeper.SetParams(ctx, feesTypes.DefaultParams())

	return app, ctx
}

// run the tx through the anteHandler and ensure its valid
func checkValidTx(t *testing.T, anteHandler sdk.AnteHandler, ctx sdk.Context, tx sdk.Tx, simulate bool) {
	_, err := anteHandler(ctx, tx, simulate)
	require.Nil(t, err)
}

// run the tx through the anteHandler and ensure it fails with the given code
func checkInvalidTx(t *testing.T, anteHandler sdk.AnteHandler, ctx sdk.Context, tx sdk.Tx, simulate bool, code error) {
	_, err := anteHandler(ctx, tx, simulate)
	require.NotNil(t, err)

	require.True(t, errors.Is(sdkerrors.ErrInsufficientFee, code))
}

func TestAnteHandlerFees_MsgCreatePost(t *testing.T) {
	// variables for later usage
	timeZone, _ := time.LoadLocation("UTC")
	pollData := types.NewPollData(
		"poll?",
		time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
		types.NewPollAnswers(
			types.NewPollAnswer(types.AnswerID(1), "Yes"),
			types.NewPollAnswer(types.AnswerID(2), "No"),
		),
		false,
		true,
	)
	attachments := types.NewAttachments(types.NewAttachment("https://uri.com", "text/plain", nil))
	id := types.PostID("dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1")

	app, ctx := createTestApp(true, false)
	feeTokenDenom := "udaric"
	defaultBondDenom := "desmos"

	anteHandler := ante.NewAnteHandler(
		app.AccountKeeper,
		app.SupplyKeeper,
		cosmosante.DefaultSigVerificationGasConsumer,
		app.FeesKeeper,
		defaultBondDenom,
	)

	// keys and addresses
	priv, _, addr := authtypes.KeyTestPubAddr()

	// Set the accounts
	acc := app.AccountKeeper.NewAccountWithAddress(ctx, addr)
	_ = acc.SetCoins(sdk.NewCoins(sdk.NewInt64Coin("desmos", 100000000000)))
	app.AccountKeeper.SetAccount(ctx, acc)

	// Prepare the msg
	msgCreatePost := types.NewMsgCreatePost(
		"My new post",
		id,
		false,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		nil,
		acc.GetAddress(),
		attachments,
		&pollData,
	)

	privs, accnums, seqs := []crypto.PrivKey{priv}, []uint64{0}, []uint64{0}
	msgs := []sdk.Msg{msgCreatePost}

	// Signer has not specified the fees
	var tx sdk.Tx
	fees := sdk.NewCoins()
	tx = authtypes.NewTestTx(ctx, msgs, privs, accnums, seqs, auth.NewStdFee(200000, fees))
	checkInvalidTx(t, anteHandler, ctx, tx, false, sdkerrors.ErrInsufficientFee)

	// Signer has not specified enough fee
	fees = sdk.NewCoins(sdk.NewInt64Coin(feeTokenDenom, 9999))
	seqs = []uint64{0}
	tx = authtypes.NewTestTx(ctx, msgs, privs, accnums, seqs, auth.NewStdFee(200000, fees))
	checkInvalidTx(t, anteHandler, ctx, tx, false, sdkerrors.ErrInsufficientFee)

	// Signer has not specified enough fee and uses default bond instead
	fees = sdk.NewCoins(sdk.NewInt64Coin(defaultBondDenom, 2))
	seqs = []uint64{1}
	tx = authtypes.NewTestTx(ctx, msgs, privs, accnums, seqs, auth.NewStdFee(200000, fees))
	checkValidTx(t, anteHandler, ctx, tx, true)

	// Signer has specified enough fee
	fees = sdk.NewCoins(sdk.NewInt64Coin(feeTokenDenom, 10000))
	_ = app.BankKeeper.SetCoins(ctx, addr, fees)
	seqs = []uint64{1}
	tx = authtypes.NewTestTx(ctx, msgs, privs, accnums, seqs, auth.NewStdFee(200000, fees))
	checkValidTx(t, anteHandler, ctx, tx, true)
}
