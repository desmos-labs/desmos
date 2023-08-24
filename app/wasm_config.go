package app

import (
	"fmt"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"

	wasmdesmos "github.com/desmos-labs/desmos/v6/cosmwasm"

	postskeeper "github.com/desmos-labs/desmos/v6/x/posts/keeper"
	poststypes "github.com/desmos-labs/desmos/v6/x/posts/types"
	postswasm "github.com/desmos-labs/desmos/v6/x/posts/wasm"

	profileskeeper "github.com/desmos-labs/desmos/v6/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v6/x/profiles/types"
	profileswasm "github.com/desmos-labs/desmos/v6/x/profiles/wasm"

	reactionskeeper "github.com/desmos-labs/desmos/v6/x/reactions/keeper"
	reactionstypes "github.com/desmos-labs/desmos/v6/x/reactions/types"
	reactionswasm "github.com/desmos-labs/desmos/v6/x/reactions/wasm"

	relationshipskeeper "github.com/desmos-labs/desmos/v6/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v6/x/relationships/types"
	relationshipswasm "github.com/desmos-labs/desmos/v6/x/relationships/wasm"

	reportskeeper "github.com/desmos-labs/desmos/v6/x/reports/keeper"
	reportstypes "github.com/desmos-labs/desmos/v6/x/reports/types"
	reportswasm "github.com/desmos-labs/desmos/v6/x/reports/wasm"

	subspaceskeeper "github.com/desmos-labs/desmos/v6/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
	subspaceswasm "github.com/desmos-labs/desmos/v6/x/subspaces/wasm"

	tokenfactorytypes "github.com/desmos-labs/desmos/v6/x/tokenfactory/types"

	desmoswasmtypes "github.com/desmos-labs/desmos/v6/x/wasm/types"
)

const (
	// DefaultDesmosInstanceCost is how much SDK gas we charge each time we load a WASM instance
	DefaultDesmosInstanceCost uint64 = 60_000
	// DefaultDesmosCompileCost is how much SDK gas is charged *per byte* for compiling WASM code
	DefaultDesmosCompileCost uint64 = 2
)

// DesmosWasmGasRegister is defaults plus a custom compile amount
func DesmosWasmGasRegister() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultDesmosInstanceCost
	gasConfig.CompileCost = DefaultDesmosCompileCost

	return gasConfig
}

func NewDesmosWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(DesmosWasmGasRegister())
}

// NewDesmosCustomQueryPlugin initialize the custom querier to handle desmos queries for contracts
func NewDesmosCustomQueryPlugin(
	cdc codec.Codec,
	grpcQueryRouter *baseapp.GRPCQueryRouter,
	profilesKeeper profileskeeper.Keeper,
	subspacesKeeper subspaceskeeper.Keeper,
	relationshipsKeeper relationshipskeeper.Keeper,
	postsKeeper postskeeper.Keeper,
	reportsKeeper reportskeeper.Keeper,
	reactionsKeeper reactionskeeper.Keeper,
) wasmkeeper.QueryPlugins {
	queriers := map[string]wasmdesmos.Querier{
		wasmdesmos.QueryRouteProfiles:      profileswasm.NewProfilesWasmQuerier(profilesKeeper, cdc),
		wasmdesmos.QueryRouteSubspaces:     subspaceswasm.NewSubspacesWasmQuerier(subspacesKeeper, cdc),
		wasmdesmos.QueryRouteRelationships: relationshipswasm.NewRelationshipsWasmQuerier(relationshipsKeeper, cdc),
		wasmdesmos.QueryRoutePosts:         postswasm.NewPostsWasmQuerier(postsKeeper, cdc),
		wasmdesmos.QueryRouteReports:       reportswasm.NewReportsWasmQuerier(reportsKeeper, cdc),
		wasmdesmos.QueryRouteReactions:     reactionswasm.NewReactionsWasmQuerier(reactionsKeeper, cdc),
		// add other modules querier here
	}

	querier := wasmdesmos.NewQuerier(queriers)

	protoCdc, ok := cdc.(*codec.ProtoCodec)
	if !ok {
		panic(fmt.Errorf("codec must be *codec.ProtoCodec type: actual: %T", cdc))
	}

	stargateCdc := codec.NewProtoCodec(desmoswasmtypes.NewWasmInterfaceRegistry(protoCdc.InterfaceRegistry()))
	return wasmkeeper.QueryPlugins{
		Stargate: wasmkeeper.AcceptListStargateQuerier(GetStargateAcceptedQueries(), grpcQueryRouter, stargateCdc),
		Custom:   querier.QueryCustom,
	}
}

// GetStargateAcceptedQueries returns the stargate accepted queries
func GetStargateAcceptedQueries() wasmkeeper.AcceptedStargateQueries {
	return wasmkeeper.AcceptedStargateQueries{
		// Register x/profiles queries
		"/desmos.profiles.v3.Query/Profile":                      &profilestypes.QueryProfileResponse{},
		"/desmos.profiles.v3.Query/IncomingDTagTransferRequests": &profilestypes.QueryIncomingDTagTransferRequestsResponse{},
		"/desmos.profiles.v3.Query/ChainLinks":                   &profilestypes.QueryChainLinksResponse{},
		"/desmos.profiles.v3.Query/ChainLinkOwners":              &profilestypes.QueryChainLinkOwnersResponse{},
		"/desmos.profiles.v3.Query/DefaultExternalAddresses":     &profilestypes.QueryDefaultExternalAddressesResponse{},
		"/desmos.profiles.v3.Query/ApplicationLinks":             &profilestypes.QueryApplicationLinksResponse{},
		"/desmos.profiles.v3.Query/ApplicationLinkByClientID":    &profilestypes.QueryApplicationLinkByClientIDResponse{},
		"/desmos.profiles.v3.Query/ApplicationLinkOwners":        &profilestypes.QueryApplicationLinkOwnersResponse{},

		// Register x/relationships queries
		"/desmos.relationships.v1.Query/Relationships": &relationshipstypes.QueryRelationshipsResponse{},
		"/desmos.relationships.v1.Query/Blocks":        &relationshipstypes.QueryBlocksResponse{},

		// Register x/subspaces queries
		"/desmos.subspaces.v3.Query/Subspaces":        &subspacestypes.QuerySubspacesResponse{},
		"/desmos.subspaces.v3.Query/Subspace":         &subspacestypes.QuerySubspaceResponse{},
		"/desmos.subspaces.v3.Query/Sections":         &subspacestypes.QuerySectionsResponse{},
		"/desmos.subspaces.v3.Query/Section":          &subspacestypes.QuerySectionResponse{},
		"/desmos.subspaces.v3.Query/UserGroups":       &subspacestypes.QueryUserGroupsResponse{},
		"/desmos.subspaces.v3.Query/UserGroup":        &subspacestypes.QueryUserGroupResponse{},
		"/desmos.subspaces.v3.Query/UserGroupMembers": &subspacestypes.QueryUserGroupMembersResponse{},
		"/desmos.subspaces.v3.Query/UserPermissions":  &subspacestypes.QueryUserPermissionsResponse{},
		"/desmos.subspaces.v3.Query/UserAllowances":   &subspacestypes.QueryUserAllowancesResponse{},
		"/desmos.subspaces.v3.Query/GroupAllowances":  &subspacestypes.QueryGroupAllowancesResponse{},

		// Register x/posts queries
		"/desmos.posts.v3.Query/SubspacePosts":                          &poststypes.QuerySubspacePostsResponse{},
		"/desmos.posts.v3.Query/SectionPosts":                           &poststypes.QuerySectionPostsResponse{},
		"/desmos.posts.v3.Query/Post":                                   &poststypes.QueryPostResponse{},
		"/desmos.posts.v3.Query/PostAttachments":                        &poststypes.QueryPostAttachmentsResponse{},
		"/desmos.posts.v3.Query/PollAnswers":                            &poststypes.QueryPollAnswersResponse{},
		"/desmos.posts.v3.Query/Params":                                 &poststypes.QueryParamsResponse{},
		"/desmos.posts.v3.Query/QueryIncomingPostOwnerTransferRequests": &poststypes.QueryIncomingPostOwnerTransferRequestsResponse{},

		// Register x/reports queries
		"/desmos.reports.v1.Query/Reports": &reportstypes.QueryReportsResponse{},
		"/desmos.reports.v1.Query/Report":  &reportstypes.QueryReportResponse{},
		"/desmos.reports.v1.Query/Reasons": &reportstypes.QueryReasonsResponse{},
		"/desmos.reports.v1.Query/Reason":  &reportstypes.QueryReasonResponse{},
		"/desmos.reports.v1.Query/Params":  &reportstypes.QueryParamsResponse{},

		// Register x/reactions queries
		"/desmos.reactions.v1.Query/Reactions":           &reactionstypes.QueryReactionsResponse{},
		"/desmos.reactions.v1.Query/Reaction":            &reactionstypes.QueryReactionResponse{},
		"/desmos.reactions.v1.Query/RegisteredReactions": &reactionstypes.QueryRegisteredReactionsResponse{},
		"/desmos.reactions.v1.Query/RegisteredReaction":  &reactionstypes.QueryRegisteredReactionResponse{},
		"/desmos.reactions.v1.Query/ReactionsParams":     &reactionstypes.QueryReactionsParamsResponse{},

		// Register x/tokenfactory queries
		"/desmos.tokenfactory.v1.Query/Params":         &tokenfactorytypes.QueryParamsResponse{},
		"/desmos.tokenfactory.v1.Query/SubspaceDenoms": &tokenfactorytypes.QuerySubspaceDenomsResponse{},
	}
}

// NewDesmosCustomMessageEncoder initialize the custom message encoder to desmos app for contracts
func NewDesmosCustomMessageEncoder(cdc codec.Codec) wasmkeeper.MessageEncoders {
	// Initialization of custom Desmos messages for contracts
	parserRouter := wasmdesmos.NewParserRouter()
	parsers := map[string]wasmdesmos.MsgParserInterface{
		wasmdesmos.WasmMsgParserRouteProfiles:      profileswasm.NewWasmMsgParser(cdc),
		wasmdesmos.WasmMsgParserRouteSubspaces:     subspaceswasm.NewWasmMsgParser(cdc),
		wasmdesmos.WasmMsgParserRouteRelationships: relationshipswasm.NewWasmMsgParser(cdc),
		wasmdesmos.WasmMsgParserRoutePosts:         postswasm.NewWasmMsgParser(cdc),
		wasmdesmos.WasmMsgParserRouteReports:       reportswasm.NewWasmMsgParser(cdc),
		wasmdesmos.WasmMsgParserRouteReactions:     reactionswasm.NewWasmMsgParser(cdc),
		// add other modules parsers here
	}

	parserRouter.Parsers = parsers
	return wasmkeeper.MessageEncoders{
		Custom: parserRouter.ParseCustom,
	}
}
