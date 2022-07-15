package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/x/posts/simulation"
	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	decoder := simulation.NewDecodeStore(cdc)

	post := types.NewPost(
		1,
		0,
		2,
		"External id",
		"This is a post text that does not contain any useful information",
		"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
		1,
		types.NewEntities(
			[]types.TextTag{
				types.NewTextTag(1, 3, "tag"),
			},
			[]types.TextTag{
				types.NewTextTag(4, 6, "tag"),
			},
			[]types.Url{
				types.NewURL(7, 9, "URL", "Display URL"),
			},
		),
		[]string{"general"},
		[]types.PostReference{
			types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
		},
		types.REPLY_SETTING_EVERYONE,
		time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
		nil,
	)
	attachment := types.NewAttachment(1, 1, 1, types.NewMedia(
		"ftp://user:password@example.com/image.png",
		"image/png",
	))
	userAnswer := types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos19r59nc7wfgc5gjnu5ga5yztkvr5qssj24krx2f")

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.NextPostIDStoreKey(1),
			Value: types.GetPostIDBytes(1),
		},
		{
			Key:   types.PostStoreKey(1, 1),
			Value: cdc.MustMarshal(&post),
		},
		{
			Key:   types.NextAttachmentIDStoreKey(1, 1),
			Value: types.GetAttachmentIDBytes(1),
		},
		{
			Key:   types.AttachmentStoreKey(1, 1, 1),
			Value: cdc.MustMarshal(&attachment),
		},
		{
			Key:   types.PollAnswerStoreKey(1, 1, 1, "cosmos19r59nc7wfgc5gjnu5ga5yztkvr5qssj24krx2f"),
			Value: cdc.MustMarshal(&userAnswer),
		},
		{
			Key:   types.ActivePollQueueKey(1, 1, 1, time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)),
			Value: types.GetPollIDBytes(1, 1, 1),
		},
		{
			Key:   []byte("Unknown key"),
			Value: nil,
		},
	}}

	testCases := []struct {
		name        string
		expectedLog string
	}{
		{"Post ID", fmt.Sprintf("PostIDA: %d\nPostIDB: %d\n",
			1, 1)},
		{"Post", fmt.Sprintf("PostA: %s\nPostB: %s\n",
			&post, &post)},
		{"Attachment ID", fmt.Sprintf("AttachmentIDA: %d\nAttachmentIDB: %d\n",
			1, 1)},
		{"Attachment", fmt.Sprintf("AttachmentA: %s\nAttachmentB: %s\n",
			&attachment, &attachment)},
		{"User answer", fmt.Sprintf("UserAnswerA: %s\nUserAnswerB: %s\n",
			&userAnswer, &userAnswer)},
		{"Active poll queue", fmt.Sprintf("SubspaceIDA: %d, PostIDA: %d, PollIDA: %d\nSubspaceIDB: %d, PostIDB: %d, PollIDB: %d\n",
			1, 1, 1, 1, 1, 1)},
		{"other", ""},
	}

	for i, tc := range testCases {
		i, tc := i, tc
		t.Run(tc.name, func(t *testing.T) {
			switch i {
			case len(testCases) - 1:
				require.Panics(t, func() { decoder(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tc.name)
			default:
				require.Equal(t, tc.expectedLog, decoder(kvPairs.Pairs[i], kvPairs.Pairs[i]), tc.name)
			}
		})
	}
}
