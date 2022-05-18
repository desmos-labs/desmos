package params

// Default simulation operation weights for messages
const (
	DefaultWeightMsgSaveProfile         int = 80
	DefaultWeightMsgDeleteProfile       int = 20
	DefaultWeightMsgRequestDTagTransfer int = 85
	DefaultWeightMsgCancelDTagTransfer  int = 25
	DefaultWeightMsgAcceptDTagTransfer  int = 75
	DefaultWeightMsgRefuseDTagTransfer  int = 25

	DefaultWeightMsgCreateRelationship int = 80
	DefaultWeightMsgDeleteRelationship int = 30
	DefaultWeightMsgBlockUser          int = 50
	DefaultWeightMsgUnblockUser        int = 50

	DefaultWeightMsgCreateSubspace          int = 80
	DefaultWeightMsgEditSubspace            int = 30
	DefaultWeightMsgDeleteSubspace          int = 5
	DefaultWeightMsgCreateSection           int = 20
	DefaultWeightMsgEditSection             int = 12
	DefaultWeightMsgMoveSection             int = 10
	DefaultWeightMsgDeleteSection           int = 5
	DefaultWeightMsgCreateUserGroup         int = 10
	DefaultWeightMsgEditUserGroup           int = 30
	DefaultWeightMsgMoveUserGroup           int = 30
	DefaultWeightMsgSetUserGroupPermissions int = 50
	DefaultWeightMsgDeleteUserGroup         int = 5
	DefaultWeightMsgAddUserToUserGroup      int = 7
	DefaultWeightMsgRemoveUserFromUserGroup int = 3
	DefaultWeightMsgSetUserPermissions      int = 85
)
