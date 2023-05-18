package simulation

// DONTCOVER

import (
	"math/rand"
	"time"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

var (
	mediaTypes = []string{
		"image/png",
		"image/jpeg",
		"image/gif",
		"video/mp4",
	}
	allowedReplySettings = []types.ReplySetting{
		types.REPLY_SETTING_EVERYONE,
		types.REPLY_SETTING_MENTIONS,
		types.REPLY_SETTING_FOLLOWERS,
		types.REPLY_SETTING_MUTUAL,
	}
)

// GenerateRandomPost generates a random post
func GenerateRandomPost(r *rand.Rand, accounts []simtypes.Account, subspaceID uint64, sectionID uint32, postID uint64, params types.Params) types.Post {
	author, _ := simtypes.RandomAcc(r, accounts)
	owner, _ := simtypes.RandomAcc(r, accounts)
	return types.NewPost(
		subspaceID,
		sectionID,
		postID,
		GenerateRandomText(r, 20),
		GenerateRandomText(r, params.MaxTextLength),
		author.Address.String(),
		0,
		nil,
		GenerateRandomTags(r, 10),
		nil,
		RandomReplySettings(r),
		time.Now(),
		nil,
		owner.Address.String(),
	)
}

// GenerateRandomText returns a random text that does not exceed the given max length
func GenerateRandomText(r *rand.Rand, maxLength uint32) string {
	return simtypes.RandStringOfLength(r, int(maxLength))
}

// GenerateRandomTags returns a slice containing tagsNumber randomly generated tags
func GenerateRandomTags(r *rand.Rand, tagsNumber int) []string {
	tags := make([]string, tagsNumber)
	for i := 0; i < tagsNumber; i++ {
		tags[i] = simtypes.RandStringOfLength(r, 10)
	}
	return tags
}

// RandomReplySettings returns a random valid ReplySetting for the post
func RandomReplySettings(r *rand.Rand) types.ReplySetting {
	return allowedReplySettings[r.Intn(len(allowedReplySettings))]
}

// RandomPost returns a random post from the slice given
func RandomPost(r *rand.Rand, posts []types.Post) types.Post {
	return posts[r.Intn(len(posts))]
}

// GenerateRandomAttachment generates a random attachment
func GenerateRandomAttachment(r *rand.Rand, post types.Post, id uint32) types.Attachment {
	return types.NewAttachment(
		post.SubspaceID,
		post.ID,
		id,
		GenerateRandomAttachmentContent(r, time.Now()),
	)
}

// GenerateRandomAttachmentContent returns a randomly generated attachment content
func GenerateRandomAttachmentContent(r *rand.Rand, currentTime time.Time) types.AttachmentContent {
	// 50% of being a poll
	if r.Intn(101) < 50 {
		return GenerateRandomPoll(r, currentTime)
	}

	return GenerateRandomMedia(r)
}

// GenerateRandomMedia returns a randomly generated media content
func GenerateRandomMedia(r *rand.Rand) *types.Media {
	return types.NewMedia(
		GenerateRandomText(r, 50),
		mediaTypes[r.Intn(len(mediaTypes))],
	)
}

// GenerateRandomPoll returns a randomly generated poll content
func GenerateRandomPoll(r *rand.Rand, currentTime time.Time) *types.Poll {
	answersNumber := r.Intn(5) + 2 // At least 2 answers are required to make sense
	answers := make([]types.Poll_ProvidedAnswer, answersNumber)
	for index := 0; index < answersNumber; index++ {
		answers[index] = types.NewProvidedAnswer(GenerateRandomText(r, 10), nil)
	}

	// 50% of accepting multiple answer edits
	acceptsMultipleAnswers := r.Intn(101) < 50

	// 50% of allowing answers edits
	allowAnswerEdits := r.Intn(101) < 50

	return types.NewPoll(
		GenerateRandomText(r, 30),
		answers,
		currentTime.Add(30*24*time.Hour),
		acceptsMultipleAnswers,
		allowAnswerEdits,
		nil,
	)
}

// RandomAttachment returns a random attachment from the ones given
func RandomAttachment(r *rand.Rand, attachments []types.Attachment) types.Attachment {
	return attachments[r.Intn(len(attachments))]
}

// RandomAnswersIndexes returns a random answers indexes slice based on the given poll
func RandomAnswersIndexes(r *rand.Rand, poll *types.Poll) (answersIndexes []uint32) {
	maxAnswersNumber := 1
	if poll.AllowsMultipleAnswers {
		maxAnswersNumber = r.Intn(len(poll.ProvidedAnswers)) + 1
	}

	// Generate some answer indexes
	indexes := make([]uint32, maxAnswersNumber)
	for index := 0; index < maxAnswersNumber; index++ {
		indexes[index] = uint32(r.Intn(len(poll.ProvidedAnswers)))
	}

	// Eliminate duplicated generated indexes
	uniqueIndexes := map[uint32]bool{}
	for _, index := range indexes {
		if _, inserted := uniqueIndexes[index]; !inserted {
			answersIndexes = append(answersIndexes, index)
		}
		uniqueIndexes[index] = true
	}

	return answersIndexes
}

// RandomMaxTextLength returns a random max text length in the [30, 999] range
func RandomMaxTextLength(r *rand.Rand) uint32 {
	return uint32(r.Intn(950)) + 30
}

// RandomPostOwnerTransferRequest returns a random post owner transfer request from the ones given
func RandomPostOwnerTransferRequest(r *rand.Rand, requests []types.PostOwnerTransferRequest) types.PostOwnerTransferRequest {
	return requests[r.Intn(len(requests))]
}
