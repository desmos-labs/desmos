package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"github.com/stretchr/testify/assert"
)

// ----------------------
// --- MsgCreateSession
// ----------------------

var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var timeZone, _ = time.LoadLocation("UTC")
var msgShareDocumentSchema = types.MsgCreateSession{
	Owner:         testOwner,
	Created:       time.Date(2019, 10, 31, 9, 42, 0, 0, timeZone),
	Namespace:     "cosmos",
	ExternalOwner: "cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge",
	PubKey:        "cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs",
	Signature:     "QmZh...===",
}

func TestMsgCreateSession_Route(t *testing.T) {
	actual := msgShareDocumentSchema.Route()
	assert.Equal(t, "magpie", actual)
}

func TestMsgCreateSession_Type(t *testing.T) {
	actual := msgShareDocumentSchema.Type()
	assert.Equal(t, "create_session", actual)
}

func TestMsgCreateSession_ValidateBasic_Schema_valid(t *testing.T) {
	err := msgShareDocumentSchema.ValidateBasic()
	assert.Nil(t, err)
}

func TestMsgCreateSession_GetSignBytes(t *testing.T) {
	actual := msgShareDocumentSchema.GetSignBytes()
	expected := `{"type":"desmos/MsgCreateSession","value":{"created":"2019-10-31T09:42:00Z","external_owner":"cosmos1njrqah832yfdv8yhxnrskerzxhj5zj9e563uge","namespace":"cosmos","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","pub_key":"cosmospub1addwnpepqf06gxm8tf4u9af99zsuphr2jmqvr2t956me5rcx9kywmrtg6jewy8gjtcs","signature":"QmZh...==="}}`
	assert.Equal(t, expected, string(actual))
}

func TestMsgCreateSession_GetSigners(t *testing.T) {
	actual := msgShareDocumentSchema.GetSigners()
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, msgShareDocumentSchema.Owner, actual[0])
}
