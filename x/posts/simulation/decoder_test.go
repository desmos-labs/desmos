package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/kv"

	sim "github.com/desmos-labs/desmos/x/posts/simulation"
	"github.com/desmos-labs/desmos/x/posts/types"
)

var (
	id              = types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2             = types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	id3             = types.PostID("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
	id4             = types.PostID("a33e173b6b96129f74acf41b5219a6bbc9f90e9e41f37115f1ce7f1f5860211c")
	privKey         = ed25519.GenPrivKey().PubKey()
	postCreatorAddr = sdk.AccAddress(privKey.Address())

	timeZone, _ = time.LoadLocation("UTC")
	testPost    = types.NewPost(
		id4,
		"",
		"Post message",
		false,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone),
		postCreatorAddr,
	).WithAttachments(types.NewAttachments(
		types.NewAttachment("https://uri.com", "text/plain", []sdk.AccAddress{postCreatorAddr}),
	)).WithPollData(types.NewPollData(
		"title",
		time.Date(2100, 1, 1, 10, 0, 0, 0, timeZone),
		types.NewPollAnswers(
			types.NewPollAnswer(0, "first"),
			types.NewPollAnswer(1, "second"),
		),
		true,
		true,
	))
)

func makeTestCodec() (cdc *codec.Codec) {
	cdc = codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	return
}

func TestDecodeStore(t *testing.T) {
	cdc := makeTestCodec()
	comments := types.PostIDs{id, id2, id3}
	postReactions := types.PostReactions{
		types.NewPostReaction(":thumbsup:", "👍", postCreatorAddr),
		types.NewPostReaction("blue_heart:", "💙", postCreatorAddr),
	}

	reaction := types.NewReaction(
		postCreatorAddr,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	totalPosts := sdk.NewInt(10)

	kvPairs := kv.Pairs{
		kv.Pair{Key: types.PostStoreKey(testPost.PostID), Value: cdc.MustMarshalBinaryBare(&testPost)},
		kv.Pair{Key: types.PostCommentsStoreKey(testPost.PostID), Value: cdc.MustMarshalBinaryBare(&comments)},
		kv.Pair{Key: types.PostReactionsStoreKey(testPost.PostID), Value: cdc.MustMarshalBinaryBare(&postReactions)},
		kv.Pair{Key: types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace), Value: cdc.MustMarshalBinaryBare(&reaction)},
		kv.Pair{Key: types.PostIndexedIDStoreKey(testPost.PostID), Value: cdc.MustMarshalBinaryBare(&totalPosts)},
		kv.Pair{Key: types.PostTotalNumberPrefix, Value: cdc.MustMarshalBinaryBare(&totalPosts)},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Post", fmt.Sprintf("PostA: %s\nPostB: %s\n", testPost, testPost)},
		{"Comments", fmt.Sprintf("CommentsA: %s\nCommentsB: %s\n", comments, comments)},
		{"PostReactions", fmt.Sprintf("PostReactionsA: %s\nPostReactionsB: %s\n", postReactions, postReactions)},
		{"Reactions", fmt.Sprintf("ReactionA: %s\nReactionB: %s\n", reaction, reaction)},
		{"PostID", fmt.Sprintf("IndexedIDA: %s\nIndexedIDB: %s\n", totalPosts, totalPosts)},
		{"TotalPots", fmt.Sprintf("TotalPostsA: %s\nTotalPostsB: %s\n", totalPosts, totalPosts)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { sim.DecodeStore(cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, sim.DecodeStore(cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
