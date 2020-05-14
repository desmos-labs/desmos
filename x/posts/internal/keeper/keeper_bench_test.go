package keeper_test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"testing"
	"time"

	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	emoji "github.com/desmos-labs/Go-Emoji-Utils"
	"github.com/desmos-labs/desmos/x/posts/internal/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
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
		RandomPostIDOrSubspace(),
		RandomMessage(r),
		r.Intn(101) <= 50,
		RandomPostIDOrSubspace().String(),
		map[string]string{},
		time.Now(),
		accounts[r.Intn(len(accounts))].Address,
	)

	if r.Intn(101) <= 50 {
		post = post.WithMedias(simulation.RandomMedias(r, accounts))
	}

	if r.Intn(101) <= 50 {
		if pollData := simulation.RandomPollData(r); pollData != nil {
			post = post.WithPollData(*pollData)
		}
	}

	return post
}

//RandomQueryParams returns randomized QueryPostsParams
//TODO: the nil parameters should have a randomize value based on existent posts values?
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

/*
// randomMapKey returns a random map key
func randomMapKey(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()
	return keys[rand.Intn(len(keys))].Interface()
}

//RandomPostReaction returns a random emoji from the Emojis Map
func RandomPostReaction(r *rand.Rand) types.PostReaction {
	accounts := sim.RandomAccounts(r, 20)
	em := emoji.Emojis[randomMapKey(emoji.Emojis).(string)]
	creator := accounts[r.Intn(len(accounts))].Address
	return types.NewPostReaction(em.Shortcodes[0], em.Value, creator)
}
*/

func BenchmarkKeeper_SavePost(b *testing.B) {
	fmt.Println("Benchmark: Save a post")
	ctx, k := SetupTestInput()
	post := RandomPost()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.SavePost(ctx, post)
	}
}

func BenchmarkKeeper_GetPost(b *testing.B) {
	fmt.Println("Benchmark: Get a post")
	ctx, k := SetupTestInput()
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		k.SavePost(ctx, RandomPost())
	}

	posts := k.GetPosts(ctx)
	randomPost := posts[r.Intn(len(posts))]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetPost(ctx, randomPost.PostID)
	}

}

func BenchmarkKeeper_GetPostsFiltered(b *testing.B) {
	fmt.Println("Benchmark: Get posts filtered")
	ctx, k := SetupTestInput()
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		k.SavePost(ctx, RandomPost())
	}

	randomQueryParams := RandomQueryParams(r)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = k.GetPostsFiltered(ctx, randomQueryParams)
	}
}

func BenchmarkKeeper_SavePostReaction(b *testing.B) {
	fmt.Println("Benchmark Save a post reaction")
	ctx, k := SetupTestInput()
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		k.SavePost(ctx, RandomPost())
	}

	posts := k.GetPosts(ctx)
	post := posts[r.Intn(len(posts))]
	//reaction := RandomPostReaction(r)

	accs := sim.RandomAccounts(r, 20)
	reactions := types.PostReactions{
		types.NewPostReaction(emoji.Emojis["1F36A"].Shortcodes[0], emoji.Emojis["1F36A"].Value, accs[r.Intn(len(accs))].Address),
		types.NewPostReaction(emoji.Emojis["1F600"].Shortcodes[0], emoji.Emojis["1F600"].Value, accs[r.Intn(len(accs))].Address),
		types.NewPostReaction(emoji.Emojis["1F630"].Shortcodes[0], emoji.Emojis["1F630"].Value, accs[r.Intn(len(accs))].Address),
		types.NewPostReaction(emoji.Emojis["1F919"].Shortcodes[0], emoji.Emojis["1F919"].Value, accs[r.Intn(len(accs))].Address),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// nolint: errcheck
		k.SavePostReaction(ctx, post.PostID, reactions[r.Intn(len(reactions))])
	}
}

func BenchmarkKeeper_GetPostReactions(b *testing.B) {
	fmt.Println("Benchmark Get a post reaction")
	ctx, k := SetupTestInput()
	r := rand.New(rand.NewSource(100))

	for i := 0; i < b.N; i++ {
		k.SavePost(ctx, RandomPost())
	}

	posts := k.GetPosts(ctx)
	post := posts[r.Intn(len(posts))]
	//reaction := RandomPostReaction(r)

	accs := sim.RandomAccounts(r, 20)
	reactions := types.PostReactions{
		types.NewPostReaction(emoji.Emojis["1F36A"].Shortcodes[0], emoji.Emojis["1F36A"].Value, accs[r.Intn(len(accs))].Address),
		types.NewPostReaction(emoji.Emojis["1F600"].Shortcodes[0], emoji.Emojis["1F600"].Value, accs[r.Intn(len(accs))].Address),
		types.NewPostReaction(emoji.Emojis["1F630"].Shortcodes[0], emoji.Emojis["1F630"].Value, accs[r.Intn(len(accs))].Address),
		types.NewPostReaction(emoji.Emojis["1F919"].Shortcodes[0], emoji.Emojis["1F919"].Value, accs[r.Intn(len(accs))].Address),
	}

	for i := 0; i < b.N; i++ {
		// nolint: errcheck
		k.SavePostReaction(ctx, post.PostID, reactions[r.Intn(len(reactions))])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetPostReactions(ctx, post.PostID)
	}
}
