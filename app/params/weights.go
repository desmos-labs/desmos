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
	DefaultWeightMsgCreateUserGroup         int = 10
	DefaultWeightMsgEditUserGroup           int = 30
	DefaultWeightMsgSetUserGroupPermissions int = 50
	DefaultWeightMsgDeleteUserGroup         int = 5
	DefaultWeightMsgAddUserToUserGroup      int = 7
	DefaultWeightMsgRemoveUserFromUserGroup int = 3
	DefaultWeightMsgSetUserPermissions      int = 85

	DefaultWeightMsgCreatePost           int = 80
	DefaultWeightMsgEditPost             int = 20
	DefaultWeightMsgDeletePost           int = 15
	DefaultWeightMsgAddPostAttachment    int = 20
	DefaultWeightMsgRemovePostAttachment int = 20
	DefaultWeightMsgAnswerPoll           int = 20
)
