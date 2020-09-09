package params

// Default simulation operation weights for messages
const (
	DefaultWeightMsgCreatePost          int = 100
	DefaultWeightMsgEditPost            int = 100
	DefaultWeightMsgAddReaction         int = 100
	DefaultWeightMsgRemoveReaction      int = 100
	DefaultWeightMsgAnswerPoll          int = 100
	DefaultWeightMsgRegisterReaction    int = 100
	DefaultWeightMsgSaveProfile         int = 100
	DefaultWeightMsgDeleteProfile       int = 100
	DefaultWeightMsgRequestDTagTransfer int = 100
	DefaultWeightMsgAcceptDTagTransfer  int = 100
	DefaultWeightMsgReportPost          int = 100
	DefaultWeightMsgCreateRelationship  int = 100
	DefaultWeightMsgDeleteRelationship  int = 100
	DefaultWeightMsgBlockUser           int = 100
	DefaultWeightMsgUnblockUser         int = 100
)
