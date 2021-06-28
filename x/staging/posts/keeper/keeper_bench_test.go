package keeper_test

import (
	"fmt"
	"math/rand"
	"testing"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	postssim "github.com/desmos-labs/desmos/x/staging/posts/simulation"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

// RandomSubspace returns a post with a 50% chance to have random medias and random poll
func RandomPost() types.Post {
	r := rand.New(rand.NewSource(100))
	accounts := simtypes.RandomAccounts(r, r.Intn(20))
	post := postssim.RandomPostData(r, accounts)
	return post.Post
}

func (suite *KeeperTestSuite) BenchmarkKeeper_SavePost(b *testing.B) {
	fmt.Println("Benchmark: Save a post")
	post := RandomPost()

	b.SetParallelism(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		suite.k.SavePost(suite.ctx, post)
	}
}

func (suite *KeeperTestSuite) BenchmarkKeeper_GetPost(b *testing.B) {
	fmt.Println("Benchmark: Get a post")
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		suite.k.SavePost(suite.ctx, RandomPost())
	}

	posts := suite.k.GetPosts(suite.ctx)
	randomPost := posts[r.Intn(len(posts))]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		suite.k.GetPost(suite.ctx, randomPost.PostID)
	}

}

func (suite *KeeperTestSuite) BenchmarkKeeper_GetPosts(b *testing.B) {
	fmt.Println("Benchmark: GetPosts")

	for i := 0; i < b.N; i++ {
		suite.k.SavePost(suite.ctx, RandomPost())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		suite.k.GetPosts(suite.ctx)
	}
}

func (suite *KeeperTestSuite) BenchmarkKeeper_SavePostReaction(b *testing.B) {
	fmt.Println("Benchmark Save Registered Reactions")
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		suite.k.SavePost(suite.ctx, RandomPost())
	}

	reaction := postssim.RandomEmojiPostReaction(r)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := suite.k.SavePostReaction(suite.ctx, reaction)
		suite.Require().NoError(err)
	}
}

func (suite *KeeperTestSuite) BenchmarkKeeper_GetPostReactions(b *testing.B) {
	fmt.Println("Benchmark Get Post Reactions")
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		suite.k.SavePost(suite.ctx, RandomPost())
	}

	posts := suite.k.GetPosts(suite.ctx)
	post := posts[r.Intn(len(posts))]
	reaction := postssim.RandomEmojiPostReaction(r)

	for i := 0; i < b.N; i++ {
		err := suite.k.SavePostReaction(suite.ctx, reaction)
		suite.Require().NoError(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var reactions []types.PostReaction
		suite.k.IteratePostReactionsByPost(suite.ctx, post.PostID, func(_ int64, reaction types.PostReaction) bool {
			reactions = append(reactions, reaction)
			return false
		})
	}
}
