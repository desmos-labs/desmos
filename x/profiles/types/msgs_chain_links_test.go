package types_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

var chainAddress = types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos")
var pubKey, _ = sdk.GetPubKeyFromBech32(
	sdk.Bech32PubKeyTypeAccPub,
	"cosmospub1addwnpepq0j8zw4t6tg3v8gh7d2d799gjhue7ewwmpg2hwr77f9kuuyzgqtrw5r6wec",
)
var msgChainLinkAccount = types.NewMsgLinkChainAccount(
	chainAddress,
	types.NewProof(
		pubKey,
		"ad112abb30e5240c7b9d21b4cc5421d76cfadfcd5977cca262523b5f5bc759457d4aa6d5c1eb6223db104b47aa1f222468be8eb5bb2762b971622ac5b96351b5",
		"text",
	),
	types.NewChainConfig("cosmos"),
	"cosmos1u9hgsqfpe3snftr7p7fsyja3wtlmj2sgf2w9yl",
)

func TestMsgLinkChainAccount_Route(t *testing.T) {
	require.Equal(t, "profiles", msgChainLinkAccount.Route())
}

func TestMsgLinkChainAccount_Type(t *testing.T) {
	require.Equal(t, "link_chain_account", msgChainLinkAccount.Type())
}

func TestMsgLinkChainAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name      string
		msg       *types.MsgLinkChainAccount
		shouldErr bool
	}{
		{
			name: "Empty chain address returns error",
			msg: &types.MsgLinkChainAccount{
				Proof:       msgChainLinkAccount.Proof,
				ChainConfig: msgChainLinkAccount.ChainConfig,
				Signer:      msgChainLinkAccount.Signer,
			},
			shouldErr: true,
		},
		{
			name: "Invalid proof returns error",
			msg: types.NewMsgLinkChainAccount(
				chainAddress,
				types.NewProof(secp256k1.GenPrivKey().PubKey(), "=", "wrong"),
				msgChainLinkAccount.ChainConfig,
				msgChainLinkAccount.Signer,
			),
			shouldErr: true,
		},
		{
			name: "Invalid chain config returns error",
			msg: types.NewMsgLinkChainAccount(
				chainAddress,
				msgChainLinkAccount.Proof,
				types.NewChainConfig(""),
				msgChainLinkAccount.Signer,
			),
			shouldErr: true,
		},
		{
			name: "Invalid signer returns error",
			msg: types.NewMsgLinkChainAccount(
				chainAddress,
				msgChainLinkAccount.Proof,
				msgChainLinkAccount.ChainConfig,
				"",
			),
			shouldErr: true,
		},
		{
			name:      "Valid message returns no error",
			msg:       msgChainLinkAccount,
			shouldErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.msg.ValidateBasic()
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgLinkChainAccount_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgLinkChainAccount","value":{"chain_address":{"prefix":"cosmos","value":"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0"},"chain_config":{"name":"cosmos"},"proof":{"plain_text":"text","pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A+RxOqvS0RYdF/NU3xSolfmfZc7YUKu4fvJLbnCCQBY3"},"signature":"ad112abb30e5240c7b9d21b4cc5421d76cfadfcd5977cca262523b5f5bc759457d4aa6d5c1eb6223db104b47aa1f222468be8eb5bb2762b971622ac5b96351b5"},"signer":"cosmos1u9hgsqfpe3snftr7p7fsyja3wtlmj2sgf2w9yl"}}`
	require.Equal(t, expected, string(msgChainLinkAccount.GetSignBytes()))
}

func TestMsgLinkChainAccount_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgChainLinkAccount.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgChainLinkAccount.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgUnlinkChainAccount = types.NewMsgUnlinkChainAccount(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
)

func TestMsgUnlinkChainAccount_Route(t *testing.T) {
	require.Equal(t, "profiles", msgUnlinkChainAccount.Route())
}

func TestMsgUnlinkChainAccount_Type(t *testing.T) {
	require.Equal(t, "unlink_chain_account", msgUnlinkChainAccount.Type())
}

func TestMsgUnlinkChainAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name     string
		msg      *types.MsgUnlinkChainAccount
		expError error
	}{
		{
			name: "Invalid owner returns error",
			msg: types.NewMsgUnlinkChainAccount(
				"",
				"cosmos",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner"),
		},
		{
			name: "Invalid chain name returns error",
			msg: types.NewMsgUnlinkChainAccount(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "chain name cannot be empty or blank"),
		},
		{
			name: "Invalid target returns error",
			msg: types.NewMsgUnlinkChainAccount(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos",
				"",
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid target"),
		},
		{
			name:     "No error message",
			msg:      msgUnlinkChainAccount,
			expError: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.expError == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.expError.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgUnlinkChainAccount_GetSignBytes(t *testing.T) {
	actual := msgUnlinkChainAccount.GetSignBytes()
	expected := `{"type":"desmos/MsgUnlinkChainAccount","value":{"chain_name":"cosmos","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","target":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgUnlinkChainAccount_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgUnlinkChainAccount.Owner)
	require.Equal(t, []sdk.AccAddress{addr}, msgUnlinkChainAccount.GetSigners())
}
