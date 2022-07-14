package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var msgCreateSubspace = types.NewMsgCreateSubspace(
	"Test subspace",
	"This is a test subspace",
	"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
	"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
	"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
)

func TestMsgCreateSubspace_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgCreateSubspace.Route())
}

func TestMsgCreateSubspace_Type(t *testing.T) {
	require.Equal(t, types.ActionCreateSubspace, msgCreateSubspace.Type())
}

func TestMsgCreateSubspace_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgCreateSubspace
		shouldErr bool
	}{
		{
			name: "invalid name returns error",
			msg: types.NewMsgCreateSubspace(
				"",
				"This is a test subspace",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			),
			shouldErr: true,
		},
		{
			name: "invalid treasury returns error",
			msg: types.NewMsgCreateSubspace(
				"Test subspace",
				"This is a test subspace",
				"cosmos1qzskhrcjnkdz2ln4",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			),
			shouldErr: true,
		},
		{
			name: "invalid owner returns error",
			msg: types.NewMsgCreateSubspace(
				"Test subspace",
				"This is a test subspace",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4ye",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			),
			shouldErr: true,
		},
		{
			name: "invalid creator returns error",
			msg: types.NewMsgCreateSubspace(
				"Test subspace",
				"This is a test subspace",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgCreateSubspace,
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

func TestMsgCreateSubspace_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCreateSubspace","value":{"creator":"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69","description":"This is a test subspace","name":"Test subspace","owner":"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez","treasury":"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0"}}`
	require.Equal(t, expected, string(msgCreateSubspace.GetSignBytes()))
}

func TestMsgCreateSubspace_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCreateSubspace.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msgCreateSubspace.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgEditSubspace = types.NewMsgEditSubspace(
	1,
	"This is a new name",
	"This is a new description",
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgEditSubspace_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgEditSubspace.Route())
}

func TestMsgEditSubspace_Type(t *testing.T) {
	require.Equal(t, types.ActionEditSubspace, msgEditSubspace.Type())
}

func TestMsgEditSubspace_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgEditSubspace
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgEditSubspace(
				0,
				"This is a new name",
				"This is a new description",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid signer returns error",
			msg: types.NewMsgEditSubspace(
				1,
				"This is a new name",
				"This is a new description",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				"cosmos1m0czrla04f7rp3z",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgEditSubspace,
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

func TestMsgEditSubspace_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgEditSubspace","value":{"description":"This is a new description","name":"This is a new name","owner":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","signer":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","subspace_id":"1","treasury":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"}}`
	require.Equal(t, expected, string(msgEditSubspace.GetSignBytes()))
}

func TestMsgEditSubspace_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgEditSubspace.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgEditSubspace.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgDeleteSubspace = types.NewMsgDeleteSubspace(
	1,
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgDeleteSubspace_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgDeleteSubspace.Route())
}

func TestMsgDeleteSubspace_Type(t *testing.T) {
	require.Equal(t, types.ActionDeleteSubspace, msgDeleteSubspace.Type())
}

func TestMsgDeleteSubspace_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgDeleteSubspace
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			msg:       types.NewMsgDeleteSubspace(0, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			shouldErr: true,
		},
		{
			name:      "invalid signer returns error",
			msg:       types.NewMsgDeleteSubspace(1, "cosmos1m0czrla04f7rp3z"),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgDeleteSubspace,
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

func TestMsgDeleteSubspace_GetSignBytes(t *testing.T) {
	expected := `{"signer":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","subspace_id":"1"}`
	require.Equal(t, expected, string(msgDeleteSubspace.GetSignBytes()))
}

func TestMsgDeleteSubspace_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgDeleteSubspace.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgDeleteSubspace.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgCreateSection = types.NewMsgCreateSection(
	1,
	"Test section",
	"This is a test section",
	0,
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgCreateSection_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgCreateSection.Route())
}

func TestMsgCreateSection_Type(t *testing.T) {
	require.Equal(t, types.ActionCreateSection, msgCreateSection.Type())
}

func TestMsgCreateSection_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgCreateSection
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgCreateSection(
				0,
				msgCreateSection.Name,
				msgCreateSection.Description,
				msgCreateSection.ParentID,
				msgCreateSection.Creator,
			),
			shouldErr: true,
		},
		{
			name: "invalid name returns error",
			msg: types.NewMsgCreateSection(
				msgCreateSection.SubspaceID,
				"",
				msgCreateSection.Description,
				msgCreateSection.ParentID,
				msgCreateSection.Creator,
			),
			shouldErr: true,
		},
		{
			name: "invalid signer returns error",
			msg: types.NewMsgCreateSection(
				msgCreateSection.SubspaceID,
				msgCreateSection.Name,
				msgCreateSection.Description,
				msgCreateSection.ParentID,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgCreateSection,
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

func TestMsgCreateSection_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCreateSection","value":{"creator":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","description":"This is a test section","name":"Test section","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgCreateSection.GetSignBytes()))
}

func TestMsgCreateSection_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCreateSection.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msgCreateSection.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgEditSection = types.NewMsgEditSection(
	1,
	1,
	"Test section",
	"This is a test section",
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgEditSection_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgEditSection.Route())
}

func TestMsgEditSection_Type(t *testing.T) {
	require.Equal(t, types.ActionEditSection, msgEditSection.Type())
}

func TestMsgEditSection_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgEditSection
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgEditSection(
				0,
				msgEditSection.SectionID,
				msgEditSection.Name,
				msgEditSection.Description,
				msgEditSection.Editor,
			),
			shouldErr: true,
		},
		{
			name: "invalid name returns error",
			msg: types.NewMsgEditSection(
				msgEditSection.SubspaceID,
				msgEditSection.SectionID,
				"",
				msgEditSection.Description,
				msgEditSection.Editor,
			),
			shouldErr: true,
		},
		{
			name: "invalid editor returns error",
			msg: types.NewMsgEditSection(
				msgEditSection.SubspaceID,
				msgEditSection.SectionID,
				msgEditSection.Name,
				msgEditSection.Description,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgEditSection,
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

func TestMsgEditSection_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgEditSection","value":{"description":"This is a test section","editor":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","name":"Test section","section_id":1,"subspace_id":"1"}}`
	require.Equal(t, expected, string(msgEditSection.GetSignBytes()))
}

func TestMsgEditSection_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgEditSection.Editor)
	require.Equal(t, []sdk.AccAddress{addr}, msgEditSection.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgMoveSection = types.NewMsgMoveSection(
	1,
	1,
	1,
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgMoveSection_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgMoveSection.Route())
}

func TestMsgMoveSection_Type(t *testing.T) {
	require.Equal(t, types.ActionMoveSection, msgMoveSection.Type())
}

func TestMsgMoveSection_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgMoveSection
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgMoveSection(
				0,
				msgMoveSection.SectionID,
				msgMoveSection.NewParentID,
				msgMoveSection.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid section id returns error",
			msg: types.NewMsgMoveSection(
				msgMoveSection.SubspaceID,
				0,
				msgMoveSection.NewParentID,
				msgMoveSection.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid editor returns error",
			msg: types.NewMsgMoveSection(
				msgMoveSection.SubspaceID,
				msgMoveSection.SectionID,
				msgMoveSection.NewParentID,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgMoveSection,
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

func TestMsgMoveSection_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgMoveSection","value":{"new_parent_id":1,"section_id":1,"signer":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgMoveSection.GetSignBytes()))
}

func TestMsgMoveSection_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgMoveSection.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgMoveSection.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgDeleteSection = types.NewMsgDeleteSection(
	1,
	1,
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgDeleteSection_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgDeleteSection.Route())
}

func TestMsgDeleteSection_Type(t *testing.T) {
	require.Equal(t, types.ActionDeleteSection, msgDeleteSection.Type())
}

func TestMsgDeleteSection_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgDeleteSection
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgDeleteSection(
				0,
				msgDeleteSection.SectionID,
				msgDeleteSection.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid section id returns error",
			msg: types.NewMsgDeleteSection(
				msgDeleteSection.SubspaceID,
				0,
				msgDeleteSection.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid editor returns error",
			msg: types.NewMsgDeleteSection(
				msgDeleteSection.SubspaceID,
				msgDeleteSection.SectionID,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgDeleteSection,
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

func TestMsgDeleteSection_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgDeleteSection","value":{"section_id":1,"signer":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgDeleteSection.GetSignBytes()))
}

func TestMsgDeleteSection_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgDeleteSection.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgDeleteSection.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgCreateUserGroup = types.NewMsgCreateUserGroup(
	1,
	0,
	"Group",
	"Description",
	types.NewPermissions(types.PermissionEditSubspace),
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgCreateUserGroup_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgCreateUserGroup.Route())
}

func TestMsgCreateUserGroup_Type(t *testing.T) {
	require.Equal(t, types.ActionCreateUserGroup, msgCreateUserGroup.Type())
}

func TestMsgCreateUserGroup_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgCreateUserGroup
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgCreateUserGroup(
				0,
				1,
				"group",
				"description",
				types.NewPermissions(types.PermissionEditSubspace),
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid group name returns error",
			msg: types.NewMsgCreateUserGroup(
				1,
				1,
				"",
				"description",
				types.NewPermissions(types.PermissionEditSubspace),
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid permissions return error",
			msg: types.NewMsgCreateUserGroup(
				1,
				1,
				"group",
				"description",
				types.NewPermissions("INVALID"),
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid creator returns error",
			msg: types.NewMsgCreateUserGroup(
				1,
				1,
				"group",
				"description",
				types.NewPermissions(types.PermissionEditSubspace),
				"cosmos1m0czrla04f7rp3zg7dsgc4kl",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgCreateUserGroup,
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

func TestMsgCreateUserGroup_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCreateUserGroup","value":{"creator":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","default_permissions":["EDIT_SUBSPACE"],"description":"Description","name":"Group","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgCreateUserGroup.GetSignBytes()))
}

func TestMsgCreateUserGroup_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCreateUserGroup.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msgCreateUserGroup.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgEditUserGroup = types.NewMsgEditUserGroup(
	1,
	1,
	"Group",
	"Description",
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgEditUserGroup_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgEditUserGroup.Route())
}

func TestMsgEditUserGroup_Type(t *testing.T) {
	require.Equal(t, types.ActionEditUserGroup, msgEditUserGroup.Type())
}

func TestMsgEditUserGroup_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgEditUserGroup
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgEditUserGroup(
				0,
				1,
				"group",
				"description",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid creator returns error",
			msg: types.NewMsgEditUserGroup(
				1,
				1,
				"group",
				"description",
				"cosmos1m0czrla04f7rp3zg7dsgc4kl",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgEditUserGroup,
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

func TestMsgEditUserGroup_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgEditUserGroup","value":{"description":"Description","group_id":1,"name":"Group","signer":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgEditUserGroup.GetSignBytes()))
}

func TestMsgEditUserGroup_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgEditUserGroup.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgEditUserGroup.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgMoveUserGroup = types.NewMsgMoveUserGroup(
	1,
	1,
	2,
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgMoveUserGroup_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgMoveUserGroup.Route())
}

func TestMsgMoveUserGroup_Type(t *testing.T) {
	require.Equal(t, types.ActionMoveUserGroup, msgMoveUserGroup.Type())
}

func TestMsgMoveUserGroup_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgMoveUserGroup
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgMoveUserGroup(
				0,
				msgMoveUserGroup.GroupID,
				msgMoveUserGroup.NewSectionID,
				msgMoveUserGroup.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid creator returns error",
			msg: types.NewMsgMoveUserGroup(
				msgMoveUserGroup.SubspaceID,
				msgMoveUserGroup.GroupID,
				msgMoveUserGroup.NewSectionID,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgMoveUserGroup,
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

func TestMsgMoveUserGroup_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgMoveUserGroup","value":{"group_id":1,"new_section_id":2,"signer":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgMoveUserGroup.GetSignBytes()))
}

func TestMsgMoveUserGroup_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgMoveUserGroup.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgMoveUserGroup.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgSetUserGroupPermissions = types.NewMsgSetUserGroupPermissions(
	1,
	1,
	types.NewPermissions(types.PermissionEditSubspace),
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgSetUserGroupPermissions_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgSetUserGroupPermissions.Route())
}

func TestMsgSetUserGroupPermissions_Type(t *testing.T) {
	require.Equal(t, types.ActionSetUserGroupPermissions, msgSetUserGroupPermissions.Type())
}

func TestMsgSetUserGroupPermissions_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgSetUserGroupPermissions
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgSetUserGroupPermissions(
				0,
				1,
				types.NewPermissions(types.PermissionEditSubspace),
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid permissions return error",
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.NewPermissions("INVALID"),
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid creator returns error",
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.NewPermissions(types.PermissionEditSubspace),
				"cosmos1m0czrla04f7rp3zg7dsgc4kl",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgSetUserGroupPermissions,
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

func TestMsgSetUserGroupPermissions_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgSetUserGroupPermissions","value":{"group_id":1,"permissions":["EDIT_SUBSPACE"],"signer":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgSetUserGroupPermissions.GetSignBytes()))
}

func TestMsgSetUserGroupPermissions_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgSetUserGroupPermissions.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgSetUserGroupPermissions.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgDeleteUserGroup = types.NewMsgDeleteUserGroup(
	1,
	1,
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgDeleteUserGroup_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgDeleteUserGroup.Route())
}

func TestMsgDeleteUserGroup_Type(t *testing.T) {
	require.Equal(t, types.ActionDeleteUserGroup, msgDeleteUserGroup.Type())
}

func TestMsgDeleteUserGroup_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgDeleteUserGroup
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgDeleteUserGroup(
				0,
				1,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid group id returns error",
			msg: types.NewMsgDeleteUserGroup(
				1,
				0,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid signer returns error",
			msg: types.NewMsgDeleteUserGroup(
				1,
				1,
				"cosmos1m0czrla04f7rp3zg7dsgc4kl",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgDeleteUserGroup,
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

func TestMsgDeleteUserGroup_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgDeleteUserGroup","value":{"group_id":1,"signer":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgDeleteUserGroup.GetSignBytes()))
}

func TestMsgDeleteUserGroup_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgDeleteUserGroup.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgDeleteUserGroup.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgAddUserToGroup = types.NewMsgAddUserToUserGroup(
	1,
	1,
	"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgAddUserToUserGroup_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgAddUserToGroup.Route())
}

func TestMsgAddUserToUserGroup_Type(t *testing.T) {
	require.Equal(t, types.ActionAddUserToUserGroup, msgAddUserToGroup.Type())
}

func TestMsgAddUserToUserGroup_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgAddUserToUserGroup
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgAddUserToUserGroup(
				0,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid group id returns error",
			msg: types.NewMsgAddUserToUserGroup(
				1,
				0,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid user returns error",
			msg: types.NewMsgAddUserToUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znn",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid signer returns error",
			msg: types.NewMsgAddUserToUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7d",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgAddUserToGroup,
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

func TestMsgAddUserToUserGroup_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgAddUserToUserGroup","value":{"group_id":1,"signer":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","subspace_id":"1","user":"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"}}`
	require.Equal(t, expected, string(msgAddUserToGroup.GetSignBytes()))
}

func TestMsgAddUserToUserGroup_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgAddUserToGroup.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgAddUserToGroup.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgRemoveUserFromUserGroup = types.NewMsgRemoveUserFromUserGroup(
	1,
	1,
	"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgRemoveUserFromUserGroup_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgRemoveUserFromUserGroup.Route())
}

func TestMsgRemoveUserFromUserGroup_Type(t *testing.T) {
	require.Equal(t, types.ActionRemoveUserFromUserGroup, msgRemoveUserFromUserGroup.Type())
}

func TestMsgRemoveUserFromUserGroup_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRemoveUserFromUserGroup
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRemoveUserFromUserGroup(
				0,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid group id returns error",
			msg: types.NewMsgRemoveUserFromUserGroup(
				1,
				0,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid user returns error",
			msg: types.NewMsgRemoveUserFromUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znn",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid signer returns error",
			msg: types.NewMsgRemoveUserFromUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7d",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgRemoveUserFromUserGroup,
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

func TestMsgRemoveUserFromUserGroup_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRemoveUserFromUserGroup","value":{"group_id":1,"signer":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","subspace_id":"1","user":"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"}}`
	require.Equal(t, expected, string(msgRemoveUserFromUserGroup.GetSignBytes()))
}

func TestMsgRemoveUserFromUserGroup_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRemoveUserFromUserGroup.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgRemoveUserFromUserGroup.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------
var msgSetUserPermissions = types.NewMsgSetUserPermissions(
	1,
	0,
	"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
	types.NewPermissions(types.PermissionEditSubspace),
	"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
)

func TestMsgSetUserPermissions_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgSetUserPermissions.Route())
}

func TestMsgSetUserPermissions_Type(t *testing.T) {
	require.Equal(t, types.ActionSetUserPermissions, msgSetUserPermissions.Type())
}

func TestMsgSetUserPermissions_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgSetUserPermissions
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgSetUserPermissions(
				0,
				1,
				"group",
				types.NewPermissions(types.PermissionEditSubspace),
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid target returns error",
			msg: types.NewMsgSetUserPermissions(
				1,
				1,
				"",
				types.NewPermissions(types.PermissionEditSubspace),
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "invalid signer returns error",
			msg: types.NewMsgSetUserPermissions(
				1,
				1,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				types.NewPermissions(types.PermissionEditSubspace),
				"cosmos1m0czrla04f7rp3zg7d",
			),
			shouldErr: true,
		},
		{
			name: "same user and signer returns error",
			msg: types.NewMsgSetUserPermissions(
				1,
				1,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				types.NewPermissions(types.PermissionEditSubspace),
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgSetUserPermissions,
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

func TestMsgSetUserPermissions_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgSetUserPermissions","value":{"permissions":["EDIT_SUBSPACE"],"signer":"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5","subspace_id":"1","user":"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"}}`
	require.Equal(t, expected, string(msgSetUserPermissions.GetSignBytes()))
}

func TestMsgSetUserPermissions_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgSetUserPermissions.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgSetUserPermissions.GetSigners())
}
