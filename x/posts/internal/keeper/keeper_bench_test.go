package keeper_test

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"testing"
	"time"

	sim "github.com/cosmos/cosmos-sdk/x/simulation"
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

func RandomMessage(r *rand.Rand) string {
	bytes := make([]byte, r.Intn(500))
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
}

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
		accounts[r.Intn(20)].Address,
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

func BenchmarkKeeper_SavePost(b *testing.B) {
	ctx, k := SetupTestInput()

	for i := 0; i < b.N; i++ {
		k.SavePost(ctx, RandomPost())
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
