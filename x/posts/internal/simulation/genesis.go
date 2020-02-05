package simulation

// DONTCOVER

import (
	"fmt"
	"strconv"
	"time"

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
	posts := RandomPosts(simState)
	reactions := randomReactions(simState, posts)
	postsGenesis := types.NewGenesisState(posts, reactions)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(postsGenesis)
}

// RandomPosts returns randomly generated genesis accounts
func RandomPosts(simState *module.SimulationState) (posts types.Posts) {
	postsNumber := simState.Rand.Intn(100)
	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	posts = make(types.Posts, postsNumber)
	for index := 0; index < postsNumber; index++ {
		id := index + 1
		privKey := ed25519.GenPrivKey().PubKey()

		posts[index] = types.NewPost(
			types.PostID(id),
			types.PostID(simState.Rand.Intn(id)),
			fmt.Sprintf("Post %d", id),
			simState.Rand.Int31n(101) < 50,
			"desmos",
			map[string]string{},
			time.Date(2020, 01, simState.Rand.Intn(27)+1, 12, 0, 0, 0, location),
			sdk.AccAddress(privKey.Address()),
			randomMedias(simState),
			randomPollData(simState),
		)
	}

	return posts
}

// randomReactions returns a randomly generated list of reactions
func randomReactions(simState *module.SimulationState, posts types.Posts) (reactionsMap map[string]types.Reactions) {
	reactionsNumber := simState.Rand.Intn(len(posts))

	reactionsMap = make(map[string]types.Reactions, reactionsNumber)
	for i := 0; i < reactionsNumber; i++ {
		reactionsLen := simState.Rand.Intn(20)
		reactions := make(types.Reactions, reactionsLen)
		for j := 0; j < reactionsLen; j++ {
			privKey := ed25519.GenPrivKey().PubKey()
			reactsValues := []string{"💙", "⬇️", "👎", "like"}

			reactions[j] = types.NewReaction(
				reactsValues[simState.Rand.Intn(len(reactsValues))],
				sdk.AccAddress(privKey.Address()),
			)
		}

		postIndex := simState.Rand.Intn(len(posts))
		reactionsMap[posts[postIndex].PostID.String()] = reactions
	}

	return reactionsMap
}

// randomMedias returns a randomly generated list of post medias
func randomMedias(simState *module.SimulationState) types.PostMedias {
	mediaNumber := simState.Rand.Intn(20)

	postMedias := make(types.PostMedias, mediaNumber)
	for i := 0; i < mediaNumber; i++ {
		host := RandomHosts[simState.Rand.Intn(len(RandomHosts))]
		mimeType := RandomMimeTypes[simState.Rand.Intn(len(RandomMimeTypes))]
		postMedias[i] = types.NewPostMedia(host+strconv.Itoa(i), mimeType)
	}

	return postMedias
}

// randomPollData returns a randomly generated poll data
func randomPollData(simState *module.SimulationState) *types.PollData {
	shouldBeNil := simState.Rand.Intn(100) < 50
	if shouldBeNil {
		return nil
	}

	answersLen := simState.Rand.Intn(10)
	answers := make(types.PollAnswers, answersLen)
	for i := 0; i < answersLen; i++ {
		answers[i] = types.NewPollAnswer(uint(i), RandomMessage(simState.Rand))
	}

	return types.NewPollData(
		RandomMessage(simState.Rand),
		RandomDate(simState.Rand),
		answers,
		simState.Rand.Intn(100) > 30, // 30% possibility of closed poll
		simState.Rand.Intn(100) > 50, // 50% possibility of multiple answers
		simState.Rand.Intn(100) > 50, // 50% possibility of allowing answers edits
	)
}
