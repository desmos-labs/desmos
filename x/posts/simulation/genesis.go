package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

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

	postsGenesis := types.NewGenesisState(
		posts,
		nil,
		randomPostReactionsEntries(simState.Rand, posts, reactionsData),
		registeredReactions(reactionsData),
		randomParams(simState),
	)

	bz, err := json.MarshalIndent(&postsGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated posts parameters:\n%s\n", bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(postsGenesis)
}

// randomPosts returns randomly generated genesis accounts
func randomPosts(simState *module.SimulationState) (posts []types.Post) {
	postsNumber := simState.Rand.Intn(100)

	posts = make([]types.Post, postsNumber)
	for index := 0; index < postsNumber; index++ {
		postData := RandomPostData(simState.Rand, simState.Accounts)
		posts[index] = postData.Post
	}

	return posts
}

// randomPostReactionsEntries returns a randomly generated list of reactions entries
func randomPostReactionsEntries(r *rand.Rand, posts []types.Post, reactionsData []ReactionData) []types.PostReactionsEntry {
	if len(posts) == 0 {
		return nil
	}

	reactionsNumber := r.Intn(len(posts))

	entries := make([]types.PostReactionsEntry, reactionsNumber)
	for i := 0; i < reactionsNumber; i++ {
		reactionsLen := r.Intn(20)
		reactions := make([]types.PostReaction, reactionsLen)

		for j := 0; j < reactionsLen; j++ {
			privKey := ed25519.GenPrivKey().PubKey()
			data := reactionsData[r.Intn(len(reactionsData))]

			reactions[j] = types.NewPostReaction(data.ShortCode, data.Value, sdk.AccAddress(privKey.Address()).String())
		}

		id := RandomPostIDFromPosts(r, posts)
		entries[i] = types.NewPostReactionsEntry(id, reactions)
	}

	return entries
}

// registeredReactions returns all the possible registered reactions
func registeredReactions(reactionsData []ReactionData) []types.RegisteredReaction {
	regReactions := make([]types.RegisteredReaction, len(reactionsData))
	for index, reactionData := range reactionsData {
		regReactions[index] = types.NewRegisteredReaction(
			reactionData.Creator.Address.String(),
			reactionData.ShortCode,
			reactionData.Value,
			reactionData.Subspace,
		)
	}

	return regReactions
}

// randomParams returns randomly generated module parameters
func randomParams(simState *module.SimulationState) types.Params {
	return RandomParams(simState.Rand)
}
