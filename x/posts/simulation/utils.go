package simulation

// DONTCOVER

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	sdk "github.com/cosmos/cosmos-sdk/types"
	emoji "github.com/desmos-labs/Go-Emoji-Utils"

	"github.com/desmos-labs/desmos/x/posts/types"
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
)

// RandomPost picks and returns a random post from an array and returns its
// position in the array.
func RandomPost(r *rand.Rand, posts []types.Post) (types.Post, int) {
	idx := r.Intn(len(posts))
	return posts[idx], idx
}

// PostData contains the randomly generated data of a post
type PostData struct {
	types.Post
	CreatorAccount simtypes.Account
}

// RandomPostData returns a randomly generated PostData based on the given random and accounts list
func RandomPostData(r *rand.Rand, accs []simtypes.Account) PostData {
	simAccount, _ := simtypes.RandomAcc(r, accs)

	// Create a random post
	post := types.NewPost(
		"",
		RandomPostID(r),
		RandomMessage(r)+RandomHashtag(r),
		r.Intn(101) <= 50, // 50% chance of allowing comments
		RandomSubspace(r),
		nil,
		RandomAttachments(r, accs),
		RandomPollData(r),
		time.Time{},
		RandomDate(r),
		simAccount.Address.String(),
	)

	// Get the post id
	bytes, _ := post.Marshal()
	hash := sha256.Sum256(bytes)
	post.PostID = hex.EncodeToString(hash[:])

	return PostData{
		Post:           post,
		CreatorAccount: simAccount,
	}
}

// PostReactionData contains all the data needed for a post reaction to be properly added or removed from a post
type PostReactionData struct {
	Shortcode string
	Value     string
	User      simtypes.Account
	PostID    string
}

// RandomPostReactionData returns a randomly generated post reaction data object
func RandomPostReactionData(
	r *rand.Rand, accs []simtypes.Account, postID string, shortCode, value string,
) PostReactionData {
	return PostReactionData{
		Shortcode: shortCode,
		Value:     value,
		User:      accs[r.Intn(len(accs))],
		PostID:    postID,
	}
}

// RandomPostID returns a randomly generated postID
func RandomPostID(r *rand.Rand) string {
	randBytes := make([]byte, 4)
	_, err := r.Read(randBytes)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(randBytes)
	return hex.EncodeToString(hash[:])
}

// RandomPostIDFromPosts returns a randomly extracted post id from the list of posts given
func RandomPostIDFromPosts(r *rand.Rand, posts []types.Post) string {
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

// RandomDate returns a random post creation date
func RandomDate(r *rand.Rand) time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := r.Int63n(delta) + min
	return time.Unix(sec, 0).Truncate(time.Millisecond)
}

// RandomHashtag returns a random hashtag from the above random hashtags
func RandomHashtag(r *rand.Rand) string {
	idx := r.Intn(len(hashtags))
	return hashtags[idx]
}

// RandomAttachments returns a randomly generated list of post attachments
func RandomAttachments(r *rand.Rand, accs []simtypes.Account) types.Attachments {
	attNumber := r.Intn(20)

	tagsLen := r.Intn(50)
	tags := make([]string, tagsLen)
	for i := 0; i < tagsLen; i++ {
		acc, _ := simtypes.RandomAcc(r, accs)
		tags[i] = acc.Address.String()
	}

	postAttachments := make(types.Attachments, attNumber)
	for i := 0; i < attNumber; i++ {
		host := RandomHosts[r.Intn(len(RandomHosts))]
		mimeType := RandomMimeTypes[r.Intn(len(RandomMimeTypes))]
		postAttachments[i] = types.NewAttachment(host+strconv.Itoa(i), mimeType, tags)
	}

	return postAttachments
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
		answers[i] = types.NewPollAnswer(fmt.Sprint(i), RandomMessage(r))
	}

	return types.NewPollData(
		RandomMessage(r),
		RandomDate(r),
		answers,
		r.Intn(100) > 50, // 50% possibility of multiple answers
		r.Intn(100) > 50, // 50% possibility of allowing answers edits
	)
}

// GetAccount gets the account having the given address from the accs list
func GetAccount(address sdk.Address, accs []simtypes.Account) *simtypes.Account {
	for _, acc := range accs {
		if acc.Address.Equals(address) {
			return &acc
		}
	}
	return nil
}

// RegisteredReactionData contains all the data needed for a registered reaction to be properly registered
type ReactionData struct {
	Creator   simtypes.Account
	ShortCode string
	Value     string
	Subspace  string
}

// RandomReactionData returns a randomly generated reaction data object
func RandomReactionData(r *rand.Rand, accs []simtypes.Account) ReactionData {
	return ReactionData{
		Creator:   accs[r.Intn(len(accs))],
		ShortCode: fmt.Sprintf(":%s:", "x"+strings.ToLower(simtypes.RandStringOfLength(r, 5))),
		Value:     fmt.Sprintf("http://%s.jpg", simtypes.RandStringOfLength(r, 5)),
		Subspace:  RandomSubspace(r),
	}
}

func RandomReactionsData(r *rand.Rand, accs []simtypes.Account) []ReactionData {
	limit := simtypes.RandIntBetween(r, 5, 20)
	reactionsData := make([]ReactionData, limit)
	for index := range reactionsData {
		reactionsData[index] = RandomReactionData(r, accs)
	}
	return reactionsData
}

// RandomEmojiPostReaction returns a random post reaction representing an emoji reaction
func RandomEmojiPostReaction(r *rand.Rand) types.PostReaction {
	accounts := simtypes.RandomAccounts(r, 20)
	creator := accounts[r.Intn(len(accounts))].Address

	rEmoji := emoji.EmojisList[r.Intn(len(emoji.EmojisList))]
	return types.NewPostReaction(rEmoji.Shortcodes[0], rEmoji.Value, creator.String())
}

func RandomParams(r *rand.Rand) types.Params {
	return types.Params{
		MaxPostMessageLength:            sdk.NewInt(int64(simtypes.RandIntBetween(r, 500, 1000))),
		MaxOptionalDataFieldsNumber:     sdk.NewInt(int64(simtypes.RandIntBetween(r, 10, 20))),
		MaxOptionalDataFieldValueLength: sdk.NewInt(int64(simtypes.RandIntBetween(r, 200, 500))),
	}
}
