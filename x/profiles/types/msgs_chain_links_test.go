package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/v5/testutil/profilestesting"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/profiles/types"

	"github.com/stretchr/testify/require"
)

var msgChainLinkAccount = types.NewMsgLinkChainAccount(
	types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
	types.NewProof(
		profilestesting.PubKeyFromBech32("cosmospub1addwnpepq0j8zw4t6tg3v8gh7d2d799gjhue7ewwmpg2hwr77f9kuuyzgqtrw5r6wec"),
		profilestesting.SingleSignatureFromHex("ad112abb30e5240c7b9d21b4cc5421d76cfadfcd5977cca262523b5f5bc759457d4aa6d5c1eb6223db104b47aa1f222468be8eb5bb2762b971622ac5b96351b5"),
		"74657874",
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
	testCases := []struct {
		name      string
		msg       *types.MsgLinkChainAccount
		shouldErr bool
	}{
		{
			name: "empty chain address returns error",
			msg: &types.MsgLinkChainAccount{
				Proof:       msgChainLinkAccount.Proof,
				ChainConfig: msgChainLinkAccount.ChainConfig,
				Signer:      msgChainLinkAccount.Signer,
			},
			shouldErr: true,
		},
		{
			name: "invalid proof returns error",
			msg: types.NewMsgLinkChainAccount(
				types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), &types.SingleSignature{}, "wrong"),
				msgChainLinkAccount.ChainConfig,
				msgChainLinkAccount.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid chain config returns error",
			msg: types.NewMsgLinkChainAccount(
				types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
				msgChainLinkAccount.Proof,
				types.NewChainConfig(""),
				msgChainLinkAccount.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid signer returns error",
			msg: types.NewMsgLinkChainAccount(
				types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
				msgChainLinkAccount.Proof,
				msgChainLinkAccount.ChainConfig,
				"",
			),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgChainLinkAccount,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgLinkChainAccount_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgLinkChainAccount","value":{"chain_address":{"type":"desmos/Bech32Address","value":{"prefix":"cosmos","value":"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0"}},"chain_config":{"name":"cosmos"},"proof":{"plain_text":"74657874","pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A+RxOqvS0RYdF/NU3xSolfmfZc7YUKu4fvJLbnCCQBY3"},"signature":{"type":"desmos/SingleSignature","value":{"signature":"rREquzDlJAx7nSG0zFQh12z6381Zd8yiYlI7X1vHWUV9SqbVwetiI9sQS0eqHyIkaL6OtbsnYrlxYirFuWNRtQ==","value_type":1}}},"signer":"cosmos1u9hgsqfpe3snftr7p7fsyja3wtlmj2sgf2w9yl"}}`
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
	testCases := []struct {
		name      string
		msg       *types.MsgUnlinkChainAccount
		shouldErr bool
	}{
		{
			name: "invalid owner returns error",
			msg: types.NewMsgUnlinkChainAccount(
				"",
				"cosmos",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			shouldErr: true,
		},
		{
			name: "invalid chain name returns error",
			msg: types.NewMsgUnlinkChainAccount(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			shouldErr: true,
		},
		{
			name: "invalid target returns error",
			msg: types.NewMsgUnlinkChainAccount(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos",
				"",
			),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgUnlinkChainAccount,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
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

// ___________________________________________________________________________________________________________________

var msgSetDefaultExternalAddress = types.NewMsgSetDefaultExternalAddress(
	"cosmos",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
)

func TestMsgSetDefaultExternalAddress_Route(t *testing.T) {
	require.Equal(t, "profiles", msgSetDefaultExternalAddress.Route())
}

func TestMsgSetDefaultExternalAddress_Type(t *testing.T) {
	require.Equal(t, "set_default_external_address", msgSetDefaultExternalAddress.Type())
}

func TestMsgSetDefaultExternalAddress_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgSetDefaultExternalAddress
		shouldErr bool
	}{
		{
			name: "invalid chain name returns error",
			msg: types.NewMsgSetDefaultExternalAddress(
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			shouldErr: true,
		},
		{
			name: "invalid target returns error",
			msg: types.NewMsgSetDefaultExternalAddress(
				"cosmos",
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			shouldErr: true,
		},
		{
			name: "invalid owner returns error",
			msg: types.NewMsgSetDefaultExternalAddress(
				"cosmos",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgSetDefaultExternalAddress,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgSetDefaultExternalAddress_GetSignBytes(t *testing.T) {
	actual := msgSetDefaultExternalAddress.GetSignBytes()
	expected := `{"type":"desmos/MsgSetDefaultExternalAddress","value":{"chain_name":"cosmos","signer":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","target":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgSetDefaultExternalAddress_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgSetDefaultExternalAddress.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgSetDefaultExternalAddress.GetSigners())
}
