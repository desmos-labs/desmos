package params

// Default simulation operation weights for messages
const (
	DefaultWeightMsgCreatePost       int = 100
	DefaultWeightMsgEditPost         int = 75
	DefaultWeightMsgAddReaction      int = 80
	DefaultWeightMsgRemoveReaction   int = 40
	DefaultWeightMsgAnswerPoll       int = 20
	DefaultWeightMsgRegisterReaction int = 50

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

	DefaultWeightMsgReportPost int = 50
)
