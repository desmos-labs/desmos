package simulation

// DONTCOVER

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/desmos/x/posts/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	RandomMimeTypes = []string{"audio/aac", "application/x-bzip2", "audio/ogg", "image/webp", "image/png"}
	RandomHosts     = []string{"https://example.com/", "https://ipfs.ink/"}
)

// RandomizedGenState generates a random GenesisState for auth
func RandomizedGenState(simState *module.SimulationState) {
	posts := randomPosts(simState)
	postReactions := randomPostReactions(simState, posts)
	registeredReactions := registeredReactions(simState)
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
			types.ComputeID(postData.CreationDate, postData.Creator.Address, postData.Subspace),
			"",
			postData.Message,
			postData.AllowsComments,
			postData.Subspace,
			postData.OptionalData,
			postData.CreationDate,
			postData.Creator.Address,
		).WithMedias(postData.Medias)

		if postData.PollData != nil {
			posts[index] = posts[index].WithPollData(*postData.PollData)
		}
	}

	return posts
}

// randomPostReactions returns a randomly generated list of reactions
func randomPostReactions(simState *module.SimulationState, posts types.Posts) (reactionsMap map[string]types.PostReactions) {
	reactionsNumber := simState.Rand.Intn(len(posts))

	reactionsMap = make(map[string]types.PostReactions, reactionsNumber)
	for i := 0; i < reactionsNumber; i++ {
		reactionsLen := simState.Rand.Intn(20)
		reactions := make(types.PostReactions, reactionsLen)
		for j := 0; j < reactionsLen; j++ {
			privKey := ed25519.GenPrivKey().PubKey()
			reactions[j] = types.NewPostReaction(RandomReactionShortCode(simState.Rand), RandomPostReactionValue(simState.Rand), sdk.AccAddress(privKey.Address()))
		}

		reactionsMap[RandomPostID(simState.Rand, posts).String()] = reactions
	}

	return reactionsMap
}

// registeredReactions returns all the possible registered reactions
func registeredReactions(simState *module.SimulationState) types.Reactions {
	reactionsData := RegisteredReactionsData(simState.Rand, simState.Accounts)

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
