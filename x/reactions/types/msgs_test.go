package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

var msgAddReaction = types.NewMsgAddReaction(
	1,
	1,
	types.NewRegisteredReactionValue(1),
	"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
)

func TestMsgAddReaction_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgAddReaction.Route())
}

func TestMsgAddReaction_Type(t *testing.T) {
	require.Equal(t, types.ActionAddReaction, msgAddReaction.Type())
}

func TestMsgAddReaction_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgAddReaction
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgAddReaction(
				0,
				msgAddReaction.PostID,
				msgAddReaction.Value.GetCachedValue().(types.ReactionValue),
				msgAddReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgAddReaction(
				msgAddReaction.SubspaceID,
				0,
				msgAddReaction.Value.GetCachedValue().(types.ReactionValue),
				msgAddReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid value returns error",
			msg: types.NewMsgAddReaction(
				msgAddReaction.SubspaceID,
				msgAddReaction.PostID,
				types.NewRegisteredReactionValue(0),
				msgAddReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid user returns error",
			msg: types.NewMsgAddReaction(
				msgAddReaction.SubspaceID,
				msgAddReaction.PostID,
				msgAddReaction.Value.GetCachedValue().(types.ReactionValue),
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgAddReaction,
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

func TestMsgAddReaction_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgAddReaction","value":{"post_id":"1","subspace_id":"1","user":"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc","value":{"type":"desmos/RegisteredReactionValue","value":{"registered_reaction_id":1}}}}`
	require.Equal(t, expected, string(msgAddReaction.GetSignBytes()))
}

func TestMsgAddReaction_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgAddReaction.User)
	require.Equal(t, []sdk.AccAddress{addr}, msgAddReaction.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgRemoveReaction = types.NewMsgRemoveReaction(
	1,
	1,
	1,
	"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
)

func TestMsgRemoveReaction_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgRemoveReaction.Route())
}

func TestMsgRemoveReaction_Type(t *testing.T) {
	require.Equal(t, types.ActionRemoveReaction, msgRemoveReaction.Type())
}

func TestMsgRemoveReaction_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRemoveReaction
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRemoveReaction(
				0,
				msgRemoveReaction.PostID,
				msgRemoveReaction.ReactionID,
				msgRemoveReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgRemoveReaction(
				msgRemoveReaction.SubspaceID,
				0,
				msgRemoveReaction.ReactionID,
				msgRemoveReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid reaction id returns error",
			msg: types.NewMsgRemoveReaction(
				msgRemoveReaction.SubspaceID,
				msgRemoveReaction.PostID,
				0,
				msgRemoveReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid user returns error",
			msg: types.NewMsgRemoveReaction(
				msgRemoveReaction.SubspaceID,
				msgRemoveReaction.PostID,
				msgRemoveReaction.ReactionID,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgRemoveReaction,
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

func TestMsgRemoveReaction_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRemoveReaction","value":{"post_id":"1","reaction_id":1,"subspace_id":"1","user":"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc"}}`
	require.Equal(t, expected, string(msgRemoveReaction.GetSignBytes()))
}

func TestMsgRemoveReaction_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRemoveReaction.User)
	require.Equal(t, []sdk.AccAddress{addr}, msgRemoveReaction.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgAddRegisteredReaction = types.NewMsgAddRegisteredReaction(
	1,
	":hello:",
	"https://example.com?images=hello.png",
	"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
)

func TestMsgAddRegisteredReaction_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgAddRegisteredReaction.Route())
}

func TestMsgAddRegisteredReaction_Type(t *testing.T) {
	require.Equal(t, types.ActionAddRegisteredReaction, msgAddRegisteredReaction.Type())
}

func TestMsgAddRegisteredReaction_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgAddRegisteredReaction
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgAddRegisteredReaction(
				0,
				msgAddRegisteredReaction.ShorthandCode,
				msgAddRegisteredReaction.DisplayValue,
				msgAddRegisteredReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid shorthand code returns error",
			msg: types.NewMsgAddRegisteredReaction(
				msgAddRegisteredReaction.SubspaceID,
				"",
				msgAddRegisteredReaction.DisplayValue,
				msgAddRegisteredReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid display value returns error",
			msg: types.NewMsgAddRegisteredReaction(
				msgAddRegisteredReaction.SubspaceID,
				msgAddRegisteredReaction.ShorthandCode,
				"",
				msgAddRegisteredReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid user returns error",
			msg: types.NewMsgAddRegisteredReaction(
				msgAddRegisteredReaction.SubspaceID,
				msgAddRegisteredReaction.ShorthandCode,
				msgAddRegisteredReaction.DisplayValue,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgAddRegisteredReaction,
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

func TestMsgAddRegisteredReaction_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgAddRegisteredReaction","value":{"display_value":"https://example.com?images=hello.png","shorthand_code":":hello:","subspace_id":"1","user":"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc"}}`
	require.Equal(t, expected, string(msgAddRegisteredReaction.GetSignBytes()))
}

func TestMsgAddRegisteredReaction_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgAddRegisteredReaction.User)
	require.Equal(t, []sdk.AccAddress{addr}, msgAddRegisteredReaction.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgEditRegisteredReaction = types.NewMsgEditRegisteredReaction(
	1,
	1,
	":hello:",
	"https://example.com?images=hello.png",
	"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
)

func TestMsgEditRegisteredReaction_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgEditRegisteredReaction.Route())
}

func TestMsgEditRegisteredReaction_Type(t *testing.T) {
	require.Equal(t, types.ActionEditRegisteredReaction, msgEditRegisteredReaction.Type())
}

func TestMsgEditRegisteredReaction_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgEditRegisteredReaction
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgEditRegisteredReaction(
				0,
				msgEditRegisteredReaction.RegisteredReactionID,
				msgEditRegisteredReaction.ShorthandCode,
				msgEditRegisteredReaction.DisplayValue,
				msgEditRegisteredReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid registered reaction id returns error",
			msg: types.NewMsgEditRegisteredReaction(
				msgEditRegisteredReaction.SubspaceID,
				0,
				msgEditRegisteredReaction.ShorthandCode,
				msgEditRegisteredReaction.DisplayValue,
				msgEditRegisteredReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid shorthand code returns error",
			msg: types.NewMsgEditRegisteredReaction(
				msgEditRegisteredReaction.SubspaceID,
				msgEditRegisteredReaction.RegisteredReactionID,
				"",
				msgEditRegisteredReaction.DisplayValue,
				msgEditRegisteredReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid display value returns error",
			msg: types.NewMsgEditRegisteredReaction(
				msgEditRegisteredReaction.SubspaceID,
				msgEditRegisteredReaction.RegisteredReactionID,
				msgEditRegisteredReaction.ShorthandCode,
				"",
				msgEditRegisteredReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid user returns error",
			msg: types.NewMsgEditRegisteredReaction(
				msgEditRegisteredReaction.SubspaceID,
				msgEditRegisteredReaction.RegisteredReactionID,
				msgEditRegisteredReaction.ShorthandCode,
				msgEditRegisteredReaction.DisplayValue,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgEditRegisteredReaction,
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

func TestMsgEditRegisteredReaction_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgEditRegisteredReaction","value":{"display_value":"https://example.com?images=hello.png","registered_reaction_id":1,"shorthand_code":":hello:","subspace_id":"1","user":"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc"}}`
	require.Equal(t, expected, string(msgEditRegisteredReaction.GetSignBytes()))
}

func TestMsgEditRegisteredReaction_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgEditRegisteredReaction.User)
	require.Equal(t, []sdk.AccAddress{addr}, msgEditRegisteredReaction.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgRemoveRegisteredReaction = types.NewMsgRemoveRegisteredReaction(
	1,
	1,
	"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
)

func TestMsgRemoveRegisteredReaction_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgRemoveRegisteredReaction.Route())
}

func TestMsgRemoveRegisteredReaction_Type(t *testing.T) {
	require.Equal(t, types.ActionRemoveRegisteredReaction, msgRemoveRegisteredReaction.Type())
}

func TestMsgRemoveRegisteredReaction_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRemoveRegisteredReaction
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRemoveRegisteredReaction(
				0,
				msgRemoveRegisteredReaction.RegisteredReactionID,
				msgRemoveRegisteredReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid registered reaction id returns error",
			msg: types.NewMsgRemoveRegisteredReaction(
				msgRemoveRegisteredReaction.SubspaceID,
				0,
				msgRemoveRegisteredReaction.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid user returns error",
			msg: types.NewMsgRemoveRegisteredReaction(
				msgRemoveRegisteredReaction.SubspaceID,
				msgRemoveRegisteredReaction.RegisteredReactionID,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgRemoveRegisteredReaction,
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

func TestMsgRemoveRegisteredReaction_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRemoveRegisteredReaction","value":{"registered_reaction_id":1,"subspace_id":"1","user":"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc"}}`
	require.Equal(t, expected, string(msgRemoveRegisteredReaction.GetSignBytes()))
}

func TestMsgRemoveRegisteredReaction_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRemoveRegisteredReaction.User)
	require.Equal(t, []sdk.AccAddress{addr}, msgRemoveRegisteredReaction.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgSetReactionsParams = types.NewMsgSetReactionsParams(
	1,
	types.NewRegisteredReactionValueParams(true),
	types.NewFreeTextValueParams(true, 100, "[a-zA-Z]"),
	"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
)

func TestMsgSetReactionsParams_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgSetReactionsParams.Route())
}

func TestMsgSetReactionsParams_Type(t *testing.T) {
	require.Equal(t, types.ActionSetReactionParams, msgSetReactionsParams.Type())
}

func TestMsgSetReactionsParams_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgSetReactionsParams
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgSetReactionsParams(
				0,
				msgSetReactionsParams.RegisteredReaction,
				msgSetReactionsParams.FreeText,
				msgSetReactionsParams.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid free text params returns error",
			msg: types.NewMsgSetReactionsParams(
				msgSetReactionsParams.SubspaceID,
				msgSetReactionsParams.RegisteredReaction,
				types.NewFreeTextValueParams(true, 0, ""),
				msgSetReactionsParams.User,
			),
			shouldErr: true,
		},
		{
			name: "invalid user returns error",
			msg: types.NewMsgSetReactionsParams(
				msgSetReactionsParams.SubspaceID,
				msgSetReactionsParams.RegisteredReaction,
				msgSetReactionsParams.FreeText,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgSetReactionsParams,
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

func TestMsgSetReactionsParams_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgSetReactionsParams","value":{"free_text":{"enabled":true,"max_length":100,"reg_ex":"[a-zA-Z]"},"registered_reaction":{"enabled":true},"subspace_id":"1","user":"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc"}}`
	require.Equal(t, expected, string(msgSetReactionsParams.GetSignBytes()))
}

func TestMsgSetReactionsParams_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgSetReactionsParams.User)
	require.Equal(t, []sdk.AccAddress{addr}, msgSetReactionsParams.GetSigners())
}
