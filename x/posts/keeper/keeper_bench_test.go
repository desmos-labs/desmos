package keeper_test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
	"time"

	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/posts/simulation"
	"github.com/desmos-labs/desmos/x/posts/types"
)

// RandomPostIDOrSubspace returns a random PostID
func RandomPostIDOrSubspace() types.PostID {
	bytes := make([]byte, 128)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(bytes)
	return types.PostID(hex.EncodeToString(hash[:]))
}

// RandomMessage returns a random String with len <= 500
func RandomMessage(r *rand.Rand) string {
	bytes := make([]byte, r.Intn(100))
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
}

// RandomPost returns a post with a 50% chance to have random medias and random poll
func RandomPost() types.Post {
	r := rand.New(rand.NewSource(100))
	accounts := sim.RandomAccounts(r, r.Intn(20))

	post := types.NewPost(
		RandomPostIDOrSubspace(),
		RandomMessage(r),
		r.Intn(101) <= 50,
		RandomPostIDOrSubspace().String(),
		nil,
		time.Now(),
		accounts[r.Intn(len(accounts))].Address,
	)

	if r.Intn(101) <= 50 {
		post = post.WithAttachments(simulation.RandomAttachments(r, accounts))
	}

	if r.Intn(101) <= 50 {
		if pollData := simulation.RandomPollData(r); pollData != nil {
			post = post.WithPollData(*pollData)
		}
	}

	return post
}

//RandomQueryParams returns randomized QueryPostsParams
func RandomQueryParams(r *rand.Rand) types.QueryPostsParams {
	sortBy := types.PostSortByCreationDate
	sortOrder := types.PostSortOrderAscending
	allowsComments := r.Intn(101) <= 50

	if r.Intn(101) <= 50 {
		sortBy = types.PostSortByID
	}

	if r.Intn(101) <= 50 {
		sortOrder = types.PostSortOrderDescending
	}

	return types.QueryPostsParams{
		Page:           r.Intn(10),
		Limit:          r.Intn(100),
		SortBy:         sortBy,
		SortOrder:      sortOrder,
		ParentID:       nil,
		CreationTime:   nil,
		AllowsComments: &allowsComments,
		Subspace:       "",
		Creator:        nil,
		Hashtags:       nil,
	}
}

func (suite *KeeperTestSuite) BenchmarkKeeper_SavePost(b *testing.B) {
	fmt.Println("Benchmark: Save a post")
	post := RandomPost()

	b.SetParallelism(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		suite.keeper.SavePost(suite.ctx, post)
	}
}

func (suite *KeeperTestSuite) BenchmarkKeeper_GetPost(b *testing.B) {
	fmt.Println("Benchmark: Get a post")
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		suite.keeper.SavePost(suite.ctx, RandomPost())
	}

	posts := suite.keeper.GetPosts(suite.ctx)
	randomPost := posts[r.Intn(len(posts))]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		suite.keeper.GetPost(suite.ctx, randomPost.PostID)
	}

}

func (suite *KeeperTestSuite) BenchmarkKeeper_GetPosts(b *testing.B) {
	fmt.Println("Benchmark: GetPosts")

	for i := 0; i < b.N; i++ {
		suite.keeper.SavePost(suite.ctx, RandomPost())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		suite.keeper.GetPosts(suite.ctx)
	}
}

func (suite *KeeperTestSuite) BenchmarkKeeper_GetPostsFiltered(b *testing.B) {
	fmt.Println("Benchmark: Get posts filtered")
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		suite.keeper.SavePost(suite.ctx, RandomPost())
	}

	randomQueryParams := RandomQueryParams(r)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = suite.keeper.GetPostsFiltered(suite.ctx, randomQueryParams)
	}
}

func (suite *KeeperTestSuite) BenchmarkKeeper_SavePostReaction(b *testing.B) {
	fmt.Println("Benchmark Save a post reaction")
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		suite.keeper.SavePost(suite.ctx, RandomPost())
	}

	posts := suite.keeper.GetPosts(suite.ctx)
	post := posts[r.Intn(len(posts))]
	reaction := simulation.RandomEmojiPostReaction(r)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// nolint: errcheck
		suite.keeper.SavePostReaction(suite.ctx, post.PostID, reaction)
	}
}

func (suite *KeeperTestSuite) BenchmarkKeeper_GetPostReactions(b *testing.B) {
	fmt.Println("Benchmark Get a post reaction")
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		suite.keeper.SavePost(suite.ctx, RandomPost())
	}

	posts := suite.keeper.GetPosts(suite.ctx)
	post := posts[r.Intn(len(posts))]
	reaction := simulation.RandomEmojiPostReaction(r)

	for i := 0; i < b.N; i++ {
		// nolint: errcheck
		suite.keeper.SavePostReaction(suite.ctx, post.PostID, reaction)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		suite.keeper.GetPostReactions(suite.ctx, post.PostID)
	}
}
