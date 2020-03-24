package simulation

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
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
	registeredReactions := randomRegisteredReactions(simState, postReactions, posts)
	postsGenesis := types.NewGenesisState(posts, postReactions, registeredReactions)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(postsGenesis)
}

// randomPosts returns randomly generated genesis accounts
func randomPosts(simState *module.SimulationState) (posts types.Posts) {
	postsNumber := simState.Rand.Intn(100)

	posts = make(types.Posts, postsNumber)
	for index := 0; index < postsNumber; index++ {
		id := index + 1

		postData := RandomPostData(simState.Rand, simState.Accounts)

		posts[index] = types.NewPost(
			types.PostID(id),
			types.PostID(simState.Rand.Intn(id)),
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
			reactions[j] = types.NewPostReaction(RandomPostReactionValue(simState.Rand), sdk.AccAddress(privKey.Address()))
		}

		reactionsMap[RandomPostID(simState.Rand, posts).String()] = reactions
	}

	return reactionsMap
}

// randomRegisteredReactions returns all the possible registered reactions based on given postReactions
func randomRegisteredReactions(simState *module.SimulationState, postReactionsMap map[string]types.PostReactions, posts types.Posts) types.Reactions {

	regReactions := types.Reactions{}

	for index, postReactions := range postReactionsMap {
		postID, err := types.ParsePostID(index)
		if err != nil {
			panic(err)
		}
		subspace := getPostSubspace(postID, posts)
		for _, postReaction := range postReactions {
			reaction := types.NewReaction(
				postReaction.Owner,
				postReaction.Value,
				RandomReactionValue(simState.Rand),
				subspace,
			)
			regReactions, _ = regReactions.AppendIfMissing(reaction)
		}
	}

	return regReactions
}

// getPostSubspace returns the post subspace from the given postID
func getPostSubspace(id types.PostID, posts types.Posts) string {
	for _, post := range posts {
		if post.PostID == id {
			return post.Subspace
		}
	}
	return ""
}
