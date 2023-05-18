package params

// Default simulation operation weights for messages
const (
	DefaultWeightMsgSaveProfile               int = 80
	DefaultWeightMsgDeleteProfile             int = 20
	DefaultWeightMsgRequestDTagTransfer       int = 85
	DefaultWeightMsgCancelDTagTransfer        int = 25
	DefaultWeightMsgAcceptDTagTransfer        int = 75
	DefaultWeightMsgRefuseDTagTransfer        int = 25
	DefaultWeightMsgLinkChainAccount          int = 75
	DefaultWeightMsgUnlinkChainAccount        int = 25
	DefaultWeightMsgSetDefaultExternalAddress int = 15

	DefaultWeightMsgCreateRelationship int = 80
	DefaultWeightMsgDeleteRelationship int = 30
	DefaultWeightMsgBlockUser          int = 50
	DefaultWeightMsgUnblockUser        int = 50

	DefaultWeightMsgCreateSubspace              int = 80
	DefaultWeightMsgEditSubspace                int = 30
	DefaultWeightMsgDeleteSubspace              int = 5
	DefaultWeightMsgCreateSection               int = 20
	DefaultWeightMsgEditSection                 int = 12
	DefaultWeightMsgMoveSection                 int = 10
	DefaultWeightMsgDeleteSection               int = 5
	DefaultWeightMsgCreateUserGroup             int = 10
	DefaultWeightMsgEditUserGroup               int = 30
	DefaultWeightMsgMoveUserGroup               int = 30
	DefaultWeightMsgSetUserGroupPermissions     int = 50
	DefaultWeightMsgDeleteUserGroup             int = 5
	DefaultWeightMsgAddUserToUserGroup          int = 7
	DefaultWeightMsgRemoveUserFromUserGroup     int = 3
	DefaultWeightMsgSetUserPermissions          int = 85
	DefaultWeightMsgGrantTreasuryAuthorization  int = 20
	DefaultWeightMsgRevokeTreasuryAuthorization int = 5
	DefaultWeightMsgGrantAllowance              int = 20
	DefaultWeightMsgRevokeAllowance             int = 5

	DefaultWeightMsgCreateReport          int = 50
	DefaultWeightMsgDeleteReport          int = 35
	DefaultWeightMsgSupportStandardReason int = 20
	DefaultWeightMsgAddReason             int = 10
	DefaultWeightMsgRemoveReason          int = 10

	DefaultWeightMsgAddReaction              int = 40
	DefaultWeightMsgRemoveReaction           int = 30
	DefaultWeightMsgAddRegisteredReaction    int = 25
	DefaultWeightMsgEditRegisteredReaction   int = 25
	DefaultWeightMsgRemoveRegisteredReaction int = 15
	DefaultWeightMsgSetReactionsParams       int = 10
)
