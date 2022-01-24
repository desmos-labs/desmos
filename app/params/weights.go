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
	DefaultWeightMsgCreateUserGroup         int = 10
	DefaultWeightMsgDeleteUserGroup         int = 5
	DefaultWeightMsgAddUserToUserGroup      int = 7
	DefaultWeightMsgRemoveUserFromUserGroup int = 3
	DefaultWeightMsgSetPermissions          int = 85
)
