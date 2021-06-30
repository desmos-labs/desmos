package keeper_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/simulation"
	types2 "github.com/desmos-labs/desmos/x/posts/types"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

// RandomSubspace returns a post with a 50% chance to have random medias and random poll
func RandomPost() types2.Post {
	r := rand.New(rand.NewSource(100))
	accounts := simtypes.RandomAccounts(r, r.Intn(20))
	post := simulation.RandomPostData(r, accounts)
	return post.Post
}

//RandomQueryParams returns randomized QueryPostsParams
func RandomQueryParams(r *rand.Rand) types2.QueryPostsParams {
	sortBy := types2.PostSortByCreationDate
	sortOrder := types2.PostSortOrderAscending

	if r.Intn(101) <= 50 {
		sortBy = types2.PostSortByID
	}

	if r.Intn(101) <= 50 {
		sortOrder = types2.PostSortOrderDescending
	}

	return types2.QueryPostsParams{
		Page:         r.Uint64(),
		Limit:        r.Uint64(),
		SortBy:       sortBy,
		SortOrder:    sortOrder,
		ParentID:     "",
		CreationTime: nil,
		Subspace:     "",
		Creator:      "",
		Hashtags:     nil,
	}
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

func (suite *KeeperTestSuite) BenchmarkKeeper_GetPostsFiltered(b *testing.B) {
	fmt.Println("Benchmark: Get posts filtered")
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		suite.k.SavePost(suite.ctx, RandomPost())
	}

	randomQueryParams := RandomQueryParams(r)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = suite.k.GetPostsFiltered(suite.ctx, randomQueryParams)
	}
}

func (suite *KeeperTestSuite) BenchmarkKeeper_SavePostReaction(b *testing.B) {
	fmt.Println("Benchmark Save a post registeredReactions")
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		suite.k.SavePost(suite.ctx, RandomPost())
	}

	posts := suite.k.GetPosts(suite.ctx)
	post := posts[r.Intn(len(posts))]
	reaction := simulation.RandomEmojiPostReaction(r)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := suite.k.SavePostReaction(suite.ctx, post.PostID, reaction)
		suite.Require().NoError(err)
	}
}

func (suite *KeeperTestSuite) BenchmarkKeeper_GetPostReactions(b *testing.B) {
	fmt.Println("Benchmark Get a post registeredReactions")
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		suite.k.SavePost(suite.ctx, RandomPost())
	}

	posts := suite.k.GetPosts(suite.ctx)
	post := posts[r.Intn(len(posts))]
	reaction := simulation.RandomEmojiPostReaction(r)

	for i := 0; i < b.N; i++ {
		err := suite.k.SavePostReaction(suite.ctx, post.PostID, reaction)
		suite.Require().NoError(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		suite.k.GetPostReactions(suite.ctx, post.PostID)
	}
}
