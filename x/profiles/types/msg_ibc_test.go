package types_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func generateMsgLinkChainAccount(t *testing.T) (*types.MsgLinkChainAccount, types.AddressData) {
	srcPrivKey := &secp256k1.PrivKey{Key: []byte{26, 15, 45, 205, 181, 29, 11, 13, 171, 161, 135, 61, 94, 174, 82, 9, 220, 10, 66, 180, 9, 49, 96, 179, 16, 189, 143, 132, 152, 111, 59, 30}}
	srcPubKey := srcPrivKey.PubKey()
	srcAddr, err := bech32.ConvertAndEncode("cosmos", srcPubKey.Address().Bytes())
	require.NoError(t, err)

	srcPlainText := srcAddr
	srcSig, err := srcPrivKey.Sign([]byte(srcPlainText))
	require.NoError(t, err)
	srcSigHex := hex.EncodeToString(srcSig)

	destPrivKey := &secp256k1.PrivKey{Key: []byte{25, 15, 45, 205, 181, 29, 11, 13, 171, 161, 135, 61, 94, 174, 82, 9, 220, 10, 66, 180, 9, 49, 96, 179, 16, 189, 143, 132, 152, 111, 59, 30}}
	destPubKey := destPrivKey.PubKey()
	destAddr, err := bech32.ConvertAndEncode("cosmos", destPubKey.Address().Bytes())

	return types.NewMsgLinkChainAccount(
		types.NewBech32Address(srcAddr, "cosmos"),
		types.NewProof(srcPubKey, srcSigHex, srcPlainText),
		types.NewChainConfig("cosmos"),
		destAddr,
	), types.NewBech32Address(srcAddr, "cosmos")
}

func TestMsgLinkChainAccount_Route(t *testing.T) {
	msg, _ := generateMsgLinkChainAccount(t)
	require.Equal(t, "profiles", msg.Route())
}

func TestMsgLinkChainAccount_Type(t *testing.T) {
	msg, _ := generateMsgLinkChainAccount(t)
	require.Equal(t, "link_chain_account", msg.Type())
}

func TestMsgLinkChainAccount_ValidateBasic(t *testing.T) {
	validMsg, srcAddr := generateMsgLinkChainAccount(t)
	tests := []struct {
		name     string
		msg      *types.MsgLinkChainAccount
		expError error
	}{
		{
			name: "Empty source address returns error",
			msg: &types.MsgLinkChainAccount{
				nil,
				validMsg.SourceProof,
				validMsg.SourceChainConfig,
				validMsg.DestinationAddress,
			},
			expError: fmt.Errorf("source address cannot be nil"),
		},
		{
			name: "Invalid source proof returns error",
			msg: types.NewMsgLinkChainAccount(
				srcAddr,
				types.NewProof(secp256k1.GenPrivKey().PubKey(), "=", "wrong"),
				validMsg.SourceChainConfig,
				validMsg.DestinationAddress,
			),
			expError: fmt.Errorf("failed to decode hex string of signature"),
		},
		{
			name: "Invalid chain config returns error",
			msg: types.NewMsgLinkChainAccount(
				srcAddr,
				validMsg.SourceProof,
				types.NewChainConfig(""),
				validMsg.DestinationAddress,
			),
			expError: fmt.Errorf("chain name cannot be empty or blank"),
		},
		{
			name: "Invalid destination address returns error",
			msg: types.NewMsgLinkChainAccount(
				srcAddr,
				validMsg.SourceProof,
				validMsg.SourceChainConfig,
				"",
			),
			expError: fmt.Errorf("invalid destination address: %s", ""),
		},
		{
			name:     "No error message",
			msg:      validMsg,
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

func TestMsgLinkChainAccount_GetSignBytes(t *testing.T) {
	msg, _ := generateMsgLinkChainAccount(t)
	actual := msg.GetSignBytes()
	expected := `{"type":"desmos/MsgLinkChainAccount","value":{"destination_address":"cosmos1u9hgsqfpe3snftr7p7fsyja3wtlmj2sgf2w9yl","source_address":{"prefix":"cosmos","value":"cosmos1ma346arwsqpmjmkctwxa5uxdx66le3nty0jeax"},"source_chain_config":{"name":"cosmos"},"source_proof":{"plain_text":"cosmos1ma346arwsqpmjmkctwxa5uxdx66le3nty0jeax","pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A7v3HEjiNO2jXJA+2gcBtO2VQ6Vsirs7GODz7dN39H7Q"},"signature":"ad112abb30e5240c7b9d21b4cc5421d76cfadfcd5977cca262523b5f5bc759457d4aa6d5c1eb6223db104b47aa1f222468be8eb5bb2762b971622ac5b96351b5"}}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgLinkChainAccount_GetSigners(t *testing.T) {
	msg, _ := generateMsgLinkChainAccount(t)
	addr, _ := sdk.AccAddressFromBech32(msg.DestinationAddress)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
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
