package simulation

import (
	"fmt"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/kv"

	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	privKey         = ed25519.GenPrivKey().PubKey()
	postCreatorAddr = sdk.AccAddress(privKey.Address())

	timeZone, _ = time.LoadLocation("UTC")
	testPost    = types.NewPost(
		types.PostID(3257),
		types.PostID(0),
		"Post message",
		false,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone),
		postCreatorAddr,
	).WithMedias(types.NewPostMedias(
		types.NewPostMedia("https://uri.com", "text/plain"),
	)).WithPollData(types.NewPollData(
		"title",
		time.Date(2100, 1, 1, 10, 0, 0, 0, timeZone),
		types.NewPollAnswers(
			types.NewPollAnswer(0, "first"),
			types.NewPollAnswer(1, "second"),
		),
		true,
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

	lastPostID := types.PostID(1)
	comments := types.PostIDs{types.PostID(10), types.PostID(14), types.PostID(20)}
	reactions := types.Reactions{
		types.NewReaction("like", postCreatorAddr),
		types.NewReaction("💙", postCreatorAddr),
	}

	kvPairs := kv.Pairs{
		kv.Pair{Key: types.LastPostIDStoreKey, Value: cdc.MustMarshalBinaryBare(lastPostID)},
		kv.Pair{Key: types.PostStoreKey(testPost.PostID), Value: cdc.MustMarshalBinaryBare(&testPost)},
		kv.Pair{Key: types.PostCommentsStoreKey(testPost.PostID), Value: cdc.MustMarshalBinaryBare(&comments)},
		kv.Pair{Key: types.PostReactionsStoreKey(testPost.PostID), Value: cdc.MustMarshalBinaryBare(&reactions)},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"LastPostID", fmt.Sprintf("LastPostIDA: %s\nLastPostIDB: %s\n", lastPostID, lastPostID)},
		{"Post", fmt.Sprintf("PostA: %s\nPostB: %s\n", testPost, testPost)},
		{"Comments", fmt.Sprintf("CommentsA: %s\nCommentsB: %s\n", comments, comments)},
		{"Reactions", fmt.Sprintf("ReactionsA: %s\nReactionsB: %s\n", reactions, reactions)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { DecodeStore(cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, DecodeStore(cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
