package simulation

import (
	"math/rand"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

var (
	randomMessages = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Vestibulum a nulla sed purus pellentesque euismod quis ut risus.",
		"Morbi nec magna interdum, rhoncus nisl ac, posuere sapien.",
		"Duis vitae nisi efficitur, lobortis neque at, bibendum ipsum.",
		"Donec semper nisi vel mollis cursus.",
		"Curabitur quis massa id libero posuere venenatis ac ac erat.",
		"Morbi dictum elit vitae libero lobortis luctus.",
		"Nam sit amet velit venenatis est scelerisque egestas vitae nec turpis.",
		"Duis commodo sapien id ligula volutpat tincidunt in et est.",
		"Cras et magna cursus, volutpat purus at, dictum diam.",
		"Phasellus in arcu euismod, accumsan urna quis, consectetur enim.",
		"Morbi tincidunt urna sit amet vulputate bibendum.",
		"Etiam vehicula eros vel libero sollicitudin elementum.",
		"Pellentesque at nunc ac orci consequat varius.",
		"Donec aliquam libero eu purus cursus, in congue magna tempor.",
		"Vivamus a dolor scelerisque, posuere justo quis, pharetra nibh.",
	}
	subspaces = []string{
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"2bdf5932925584b9a86470bea60adce69041608a447f84a3317723aa5678ec88",
		"3d59f7548e1af2151b64135003ce63c0a484c26b9b8b166a7b1c1805ec34b00a",
		"ec8202b6f9fb16f9e26b66367afa4e037752f3c09a18cefab426165e06a424b1",
		"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
		"3f40462915a3e6026a4d790127b95ded4d870f6ab18d9af2fcbc454168255237",
	}

	hashtags = []string{"#desmos", "#mooncake", "#test", "#cosmos", "#terra", "#bidDipper"}

	shortCodes  = []string{":blue_heart:", ":arrow_down:", ":thumbsdown:", ":thumbsup:", ":dog:", ":cat:"}
	reactValues = []string{"http://earth.jpg", "U+1F600", "U+1F605", "U+1F610"}
)

// RandomPost picks and returns a random post from an array and returns its
// position in the array.
func RandomPost(r *rand.Rand, posts types.Posts) (types.Post, int) {
	idx := r.Intn(len(posts))
	return posts[idx], idx
}

// PostData contains the randomly generated data of a post
type PostData struct {
	Creator        sim.Account
	ParentID       types.PostID
	Message        string
	AllowsComments bool
	Subspace       string
	CreationDate   time.Time
	OptionalData   map[string]string
	Medias         types.PostMedias
	PollData       *types.PollData
}

// RandomPostData returns a randomly generated PostData based on the given random and accounts list
func RandomPostData(r *rand.Rand, accs []sim.Account) PostData {
	simAccount, _ := sim.RandomAcc(r, accs)
	return PostData{
		Creator:        simAccount,
		ParentID:       types.PostID(0),
		Message:        RandomMessage(r) + RandomHashtag(r),
		AllowsComments: r.Intn(101) <= 50, // 50% chance of allowing comments
		Subspace:       RandomSubspace(r),
		CreationDate:   time.Now().UTC(),
		Medias:         RandomMedias(r),
		PollData:       RandomPollData(r),
	}
}

// PostReactionData contains all the data needed for a post reaction to be properly added or removed from a post
type PostReactionData struct {
	Value  string
	User   sim.Account
	PostID types.PostID
}

// RandomPostReactionData returns a randomly generated post reaction data object
func RandomPostReactionData(r *rand.Rand, accs []sim.Account, postID types.PostID, shortCode string) PostReactionData {
	return PostReactionData{
		Value:  shortCode,
		User:   accs[r.Intn(len(accs))],
		PostID: postID,
	}
}

// RandomPostReactionValue returns a random reaction value
func RandomPostReactionValue(r *rand.Rand) string {
	return shortCodes[r.Intn(len(shortCodes))]
}

// RandomPostID returns a randomly extracted post id from the list of posts given
func RandomPostID(r *rand.Rand, posts []types.Post) types.PostID {
	p, _ := RandomPost(r, posts)
	return p.PostID
}

// RandomMessage returns a random post message from the above random lorem phrases
func RandomMessage(r *rand.Rand) string {
	idx := r.Intn(len(randomMessages))
	return randomMessages[idx]
}

// RandomSubspace returns a random post subspace from the above random subspaces
func RandomSubspace(r *rand.Rand) string {
	idx := r.Intn(len(subspaces))
	return subspaces[idx]
}

// RandomHashtag returns a random hashtag from the above random hashtags
func RandomHashtag(r *rand.Rand) string {
	idx := r.Intn(len(hashtags))
	return hashtags[idx]
}

// RandomMedias returns a randomly generated list of post medias
func RandomMedias(r *rand.Rand) types.PostMedias {
	mediaNumber := r.Intn(20)

	postMedias := make(types.PostMedias, mediaNumber)
	for i := 0; i < mediaNumber; i++ {
		host := RandomHosts[r.Intn(len(RandomHosts))]
		mimeType := RandomMimeTypes[r.Intn(len(RandomMimeTypes))]
		postMedias[i] = types.NewPostMedia(host+strconv.Itoa(i), mimeType)
	}

	return postMedias
}

// RandomPollData returns a randomly generated poll data
func RandomPollData(r *rand.Rand) *types.PollData {
	shouldBeNil := r.Intn(100) < 50
	if shouldBeNil {
		return nil
	}

	answersLen := r.Intn(10) + 2 // Answers must be at least two
	answers := make(types.PollAnswers, answersLen)
	for i := 0; i < answersLen; i++ {
		answers[i] = types.NewPollAnswer(types.AnswerID(i), RandomMessage(r))
	}

	closingDate := time.Now().UTC()

	// 30% possibility of closed poll
	open := r.Intn(100) > 70
	if open {
		closingDate = time.Now().UTC().AddDate(1, 0, 0)
	}

	poll := types.NewPollData(
		RandomMessage(r),
		closingDate,
		answers,
		open,
		r.Intn(100) > 50, // 50% possibility of multiple answers
		r.Intn(100) > 50, // 50% possibility of allowing answers edits
	)
	return &poll
}

// GetAccount gets the account having the given address from the accs list
func GetAccount(address sdk.Address, accs []sim.Account) *sim.Account {
	for _, acc := range accs {
		if acc.Address.Equals(address) {
			return &acc
		}
	}
	return nil
}

// RegisteredReactionData contains all the data needed for a registered reaction to be properly registered
type ReactionData struct {
	Creator   sim.Account
	ShortCode string
	Value     string
	Subspace  string
}

// RandomReactionValue returns a random reaction value
func RandomReactionValue(r *rand.Rand) string {
	return reactValues[r.Intn(len(reactValues))]
}

// RandomReactionShortCode return a random reaction shortCode
func RandomReactionShortCode(r *rand.Rand) string {
	return shortCodes[r.Intn(len(reactValues))]
}

// RandomReactionData returns a randomly generated reaction data object
func RandomReactionData(r *rand.Rand, accs []sim.Account) ReactionData {
	return ReactionData{
		Creator:   accs[r.Intn(len(accs))],
		ShortCode: RandomReactionShortCode(r),
		Value:     RandomReactionValue(r),
		Subspace:  RandomSubspace(r),
	}
}

// RegisteredReactionsData returns all the possible registered reactions with given data
func RegisteredReactionsData(r *rand.Rand, accs []sim.Account) []ReactionData {
	reactionsData := []ReactionData{}

	for _, subspace := range subspaces {
		for _, shortCode := range shortCodes {
			reactionData := ReactionData{
				Creator:   accs[r.Intn(len(accs))],
				ShortCode: shortCode,
				Value:     RandomReactionValue(r),
				Subspace:  subspace,
			}
			reactionsData = append(reactionsData, reactionData)
		}
	}

	return reactionsData
}
