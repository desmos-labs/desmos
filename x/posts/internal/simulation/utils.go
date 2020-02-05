package simulation

import (
	"math/rand"
	"strconv"
	"time"

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
)

// RandomAcc picks and returns a random post from an array and returns its
// position in the array.
func RandomPost(r *rand.Rand, posts types.Posts) (types.Post, int) {
	idx := r.Intn(len(posts))
	return posts[idx], idx
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

// RandomDate returns a randomly generated date
func RandomDate(r *rand.Rand) time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := r.Int63n(delta) + min
	return time.Unix(sec, 0)
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

	answersLen := r.Intn(10) + 1 // Answers should never be empty
	answers := make(types.PollAnswers, answersLen)
	for i := 0; i < answersLen; i++ {
		answers[i] = types.NewPollAnswer(uint(i), RandomMessage(r))
	}

	return types.NewPollData(
		RandomMessage(r),
		RandomDate(r),
		answers,
		r.Intn(100) > 30, // 30% possibility of closed poll
		r.Intn(100) > 50, // 50% possibility of multiple answers
		r.Intn(100) > 50, // 50% possibility of allowing answers edits
	)
}
