package keeper_test

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/desmos-labs/desmos/x/posts/internal/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"math/rand"
	"testing"
)

// RandomPostIDOrSubspace
func RandomPostIDOrSubspace() types.PostID {
	bytes := make([]byte, 128)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(bytes)
	return types.PostID(hex.EncodeToString(hash[:]))
}

func RandomPosts() ([]types.Post, rand.Rand) {
	r := rand.New(rand.NewSource(100))
	return []types.Post{
		types.NewPost(
			RandomPostIDOrSubspace(),
			RandomPostIDOrSubspace(),
			"Message",
			false,
			RandomPostIDOrSubspace().String(),
			map[string]string{},
			testPost.Created,
			testPost.Creator,
		),
		types.NewPost(
			RandomPostIDOrSubspace(),
			RandomPostIDOrSubspace(),
			"Message",
			false,
			RandomPostIDOrSubspace().String(),
			map[string]string{},
			testPost.Created,
			testPost.Creator,
		).WithMedias(testPost.Medias),
		types.NewPost(
			RandomPostIDOrSubspace(),
			RandomPostIDOrSubspace(),
			"Message",
			false,
			RandomPostIDOrSubspace().String(),
			map[string]string{},
			testPost.Created,
			testPost.Creator,
		).WithMedias(simulation.RandomMedias(r)),
		types.NewPost(
			RandomPostIDOrSubspace(),
			RandomPostIDOrSubspace(),
			"Message",
			false,
			RandomPostIDOrSubspace().String(),
			map[string]string{},
			testPost.Created,
			testPost.Creator,
		).WithPollData(*simulation.RandomPollData(r)),
		types.NewPost(
			RandomPostIDOrSubspace(),
			RandomPostIDOrSubspace(),
			"Message",
			false,
			RandomPostIDOrSubspace().String(),
			map[string]string{},
			testPost.Created,
			testPost.Creator,
		).WithMedias(testPost.Medias).WithPollData(*simulation.RandomPollData(r)),
		types.NewPost(
			RandomPostIDOrSubspace(),
			RandomPostIDOrSubspace(),
			"Message",
			false,
			RandomPostIDOrSubspace().String(),
			map[string]string{},
			testPost.Created,
			testPost.Creator,
		).WithMedias(simulation.RandomMedias(r)).WithPollData(*simulation.RandomPollData(r)),
	}, *r
}

func BenchmarkKeeper_SavePost(b *testing.B) {
	ctx, k := SetupTestInput()
	posts, r := RandomPosts()

	r.Intn(len(posts))

	for i := 0; i < b.N; i++ {

	}
}

func BenchmarkKeeper_GetPost(b *testing.B) {

}

func BenchmarkKeeper_GetPostsFiltered(b *testing.B) {

}

func BenchmarkKeeper_SavePostReaction(b *testing.B) {

}

func BenchmarkKeeper_GetPostReactions(b *testing.B) {

}
