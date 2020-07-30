package simulation

// DONTCOVER

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/desmos-labs/desmos/x/posts/types"
)

var (
	RandomMimeTypes = []string{"audio/aac", "application/x-bzip2", "audio/ogg", "image/webp", "image/png"}
	RandomHosts     = []string{"https://example.com/", "https://ipfs.ink/"}
)

// RandomizedGenState generates a random GenesisState for auth
func RandomizedGenState(simState *module.SimulationState) {
	posts := randomPosts(simState)
	reactionsData := RandomReactionsData(simState.Rand, simState.Accounts)

	postReactions := randomPostReactions(simState, posts, reactionsData[simState.Rand.Intn(len(reactionsData))])
	registeredReactions := registeredReactions(reactionsData)
	params := randomParams(simState)
	postsGenesis := types.NewGenesisState(posts, postReactions, registeredReactions, params)

	fmt.Printf("Selected randomly generated posts parameters:\n%s\n%s\n%s\n",
		codec.MustMarshalJSONIndent(simState.Cdc, postsGenesis.Params.MaxPostMessageLength),
		codec.MustMarshalJSONIndent(simState.Cdc, postsGenesis.Params.MaxOptionalDataFieldsNumber),
		codec.MustMarshalJSONIndent(simState.Cdc, postsGenesis.Params.MaxOptionalDataFieldValueLength),
	)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(postsGenesis)
}

// randomPosts returns randomly generated genesis accounts
func randomPosts(simState *module.SimulationState) (posts types.Posts) {
	postsNumber := simState.Rand.Intn(100)

	posts = make(types.Posts, postsNumber)
	for index := 0; index < postsNumber; index++ {
		postData := RandomPostData(simState.Rand, simState.Accounts)
		posts[index] = types.NewPost(
			RandomPostID(simState.Rand),
			"",
			postData.Message,
			postData.AllowsComments,
			postData.Subspace,
			postData.OptionalData,
			time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC),
			postData.Creator.Address,
		).WithAttachments(postData.Attachments)

		if postData.PollData != nil {
			posts[index] = posts[index].WithPollData(*postData.PollData)
		}
	}

	return posts
}

// randomPostReactions returns a randomly generated list of reactions
func randomPostReactions(simState *module.SimulationState, posts types.Posts, reactionData ReactionData) (reactionsMap map[string]types.PostReactions) {
	reactionsNumber := simState.Rand.Intn(len(posts))

	reactionsMap = make(map[string]types.PostReactions, reactionsNumber)
	for i := 0; i < reactionsNumber; i++ {
		reactionsLen := simState.Rand.Intn(20)
		reactions := make(types.PostReactions, reactionsLen)
		for j := 0; j < reactionsLen; j++ {
			privKey := ed25519.GenPrivKey().PubKey()
			reactions[j] = types.NewPostReaction(reactionData.ShortCode, reactionData.Value, sdk.AccAddress(privKey.Address()))
		}

		reactionsMap[RandomPostIDFromPosts(simState.Rand, posts).String()] = reactions
	}

	return reactionsMap
}

// registeredReactions returns all the possible registered reactions
func registeredReactions(reactionsData []ReactionData) types.Reactions {
	regReactions := types.Reactions{}

	for _, reactionData := range reactionsData {
		reaction := types.NewReaction(
			reactionData.Creator.Address,
			reactionData.ShortCode,
			reactionData.Value,
			reactionData.Subspace,
		)
		regReactions = append(regReactions, reaction)
	}
	return regReactions
}

func randomParams(simState *module.SimulationState) types.Params {
	return RandomParams(simState.Rand)
}
