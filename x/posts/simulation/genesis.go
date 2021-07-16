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
		randomPostReactions(simState.Rand, posts, reactionsData),
		registeredReactions(reactionsData),
		randomReports(simState),
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

// randomPostReactions returns a randomly generated list of reactions
func randomPostReactions(r *rand.Rand, posts []types.Post, reactionsData []ReactionData) []types.PostReaction {
	if len(posts) == 0 {
		return nil
	}

	postsNumber := r.Intn(len(posts))

	reactions := make([]types.PostReaction, postsNumber)
	for i := 0; i < postsNumber; i++ {
		id := RandomPostIDFromPosts(r, posts)
		privKey := ed25519.GenPrivKey().PubKey()
		data := reactionsData[r.Intn(len(reactionsData))]
		reactions[i] = types.NewPostReaction(id, data.ShortCode, data.Value, sdk.AccAddress(privKey.Address()).String())
	}
	return reactions
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

func randomReports(simState *module.SimulationState) (reportsMap []types.Report) {
	reportsMapLen := simState.Rand.Intn(50)

	reports := make([]types.Report, reportsMapLen)
	for i := 0; i < reportsMapLen; i++ {
		privKey := ed25519.GenPrivKey().PubKey()
		reports[i] = types.NewReport(
			RandomPostID(simState.Rand),
			RandomReportTypes(simState.Rand),
			RandomReportMessage(simState.Rand),
			sdk.AccAddress(privKey.Address()).String(),
		)
	}

	return reports
}
