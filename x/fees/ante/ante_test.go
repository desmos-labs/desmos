package ante_test

import (
	"errors"
	"github.com/desmos-labs/desmos/x/fees"
	"github.com/desmos-labs/desmos/x/fees/ante"
	"testing"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	cosmosante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	desmos "github.com/desmos-labs/desmos/app"
	feesTypes "github.com/desmos-labs/desmos/x/fees/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/desmos-labs/desmos/x/posts/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// returns context and app with params set on account keeper
func createTestApp(isCheckTx bool, isBlockZero bool) (*desmos.DesmosApp, sdk.Context) {
	db := dbm.NewMemDB()
	app := desmos.NewDesmosApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, 0)

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

	anteHandler := ante.NewAnteHandler(
		app.AccountKeeper,
		app.SupplyKeeper,
		cosmosante.DefaultSigVerificationGasConsumer,
		app.FeesKeeper,
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

	feesParams := fees.NewParams([]feesTypes.MinFee{
		feesTypes.NewMinFee("create_post", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
	})

	app.FeesKeeper.SetParams(ctx, feesParams)

	// Signer has not specified the fees
	var tx sdk.Tx
	feez := sdk.NewCoins()
	tx = authtypes.NewTestTx(ctx, msgs, privs, accnums, seqs, auth.NewStdFee(200000, feez))
	checkInvalidTx(t, anteHandler, ctx, tx, false, sdkerrors.ErrInsufficientFee)

	// Signer has not specified enough fee
	feez = sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 9999))
	seqs = []uint64{0}
	tx = authtypes.NewTestTx(ctx, msgs, privs, accnums, seqs, auth.NewStdFee(200000, feez))
	checkInvalidTx(t, anteHandler, ctx, tx, false, sdkerrors.ErrInsufficientFee)

	// Signer has specified enough fee
	feez = sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000))
	_ = app.BankKeeper.SetCoins(ctx, addr, feez)
	seqs = []uint64{1}
	tx = authtypes.NewTestTx(ctx, msgs, privs, accnums, seqs, auth.NewStdFee(200000, feez))
	checkValidTx(t, anteHandler, ctx, tx, true)
}
