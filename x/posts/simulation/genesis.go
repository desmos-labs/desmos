package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
	subspacessim "github.com/desmos-labs/desmos/v3/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// RandomizeGenState generates a random GenesisState for posts
func RandomizeGenState(simState *module.SimulationState) {
	// Read the subspaces data
	subspacesGenesisBz := simState.GenState[subspacestypes.ModuleName]
	var subspacesGenesis subspacestypes.GenesisState
	simState.Cdc.MustUnmarshalJSON(subspacesGenesisBz, &subspacesGenesis)

	params := types.NewParams(RandomMaxTextLength(simState.Rand))
	posts := randomPosts(simState.Rand, subspacesGenesis.Subspaces, simState.Accounts, params)
	subspacesDataEntries := getSubspacesData(posts)
	attachments := randomAttachments(simState.Rand, posts)
	genesisPosts := getGenesisPosts(posts, attachments)
	userAnswers := randomUserAnswers(simState.Rand, attachments, simState.Accounts)

	// Save the genesis
	postsGenesis := types.NewGenesisState(subspacesDataEntries, genesisPosts, attachments, userAnswers, params)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(postsGenesis)
}

// randomPosts returns randomly generated genesis posts
func randomPosts(r *rand.Rand, subspaces []subspacestypes.GenesisSubspace, accs []simtypes.Account, params types.Params) (posts []types.Post) {
	postsNumber := uint64(r.Intn(100))
	posts = make([]types.Post, postsNumber)
	for index := uint64(0); index < postsNumber; index++ {
		subspace := subspacessim.RandomGenesisSubspace(r, subspaces)
		posts[index] = GenerateRandomPost(r, accs, subspace.Subspace.ID, index+1, params)
	}
	return posts
}

// getSubspacesData returns the
func getSubspacesData(posts []types.Post) (entries []types.SubspaceDataEntry) {
	if len(posts) == 0 {
		return nil
	}

	subspacesMaxPostID := map[uint64]uint64{}
	for _, post := range posts {
		maxPostID, ok := subspacesMaxPostID[post.SubspaceID]
		if !ok || post.ID > maxPostID {
			subspacesMaxPostID[post.SubspaceID] = post.ID
		}
	}

	entries = make([]types.SubspaceDataEntry, len(subspacesMaxPostID))
	var i = 0
	for subspaceID, maxPostID := range subspacesMaxPostID {
		entries[i] = types.NewSubspaceDataEntry(subspaceID, maxPostID+1)
		i++
	}

	return entries
}

// randomAttachments returns randomly generated attachments
func randomAttachments(r *rand.Rand, posts []types.Post) (attachments []types.Attachment) {
	if len(posts) == 0 {
		return nil
	}

	attachmentsNumber := uint32(r.Intn(50))
	attachments = make([]types.Attachment, attachmentsNumber)
	for index := uint32(0); index < attachmentsNumber; index++ {
		post := RandomPost(r, posts)
		attachments[index] = GenerateRandomAttachment(r, post, index+1)
	}
	return attachments
}

// randomUserAnswers returns randomly generated user answers
func randomUserAnswers(r *rand.Rand, attachments []types.Attachment, accs []simtypes.Account) (answers []types.UserAnswer) {
	if len(attachments) == 0 {
		return nil
	}

	// Get only the polls
	var polls []types.Attachment
	for _, attachment := range attachments {
		if types.IsPoll(attachment) {
			polls = append(polls, attachment)
		}
	}

	answersNumber := r.Intn(50)
	for index := 0; index < answersNumber; index++ {
		attachment := RandomAttachment(r, polls)
		answersIndexes := RandomAnswersIndexes(r, attachment.Content.GetCachedValue().(*types.Poll))
		user, _ := simtypes.RandomAcc(r, accs)
		answer := types.NewUserAnswer(attachment.SubspaceID, attachment.PostID, attachment.ID, answersIndexes, user.Address.String())

		// Make sure there are no duplicated answers
		if !containsAnswer(answers, answer) {
			answers = append(answers, answer)
		}
	}
	return answers
}

// containsAnswer tells whether the given answers slice contains an answer from the same user of the given one
func containsAnswer(answers []types.UserAnswer, answer types.UserAnswer) bool {
	for _, item := range answers {
		if item.SubspaceID == answer.SubspaceID && item.PostID == answer.PostID && item.PollID == answer.PollID && item.User == answer.User {
			return true
		}
	}
	return false
}

// getGenesisPosts uses the given posts and attachments to return a genesis posts slice
func getGenesisPosts(posts []types.Post, attachments []types.Attachment) (genesisPosts []types.GenesisPost) {
	if len(posts) == 0 {
		return nil
	}

	type postReference struct {
		SubspaceID uint64
		PostID     uint64
	}

	// Get the max attachment id for each post that has an attachment
	maxAttachmentIDs := map[postReference]uint32{}
	for _, attachment := range attachments {
		key := postReference{SubspaceID: attachment.SubspaceID, PostID: attachment.PostID}
		maxAttachmentID, ok := maxAttachmentIDs[key]
		if !ok || maxAttachmentID < attachment.ID {
			maxAttachmentIDs[key] = attachment.ID
		}
	}

	genesisPosts = make([]types.GenesisPost, len(posts))
	for i, post := range posts {
		key := postReference{SubspaceID: post.SubspaceID, PostID: post.ID}
		maxAttachmentID, ok := maxAttachmentIDs[key]
		if !ok {
			maxAttachmentID = 0
		}
		genesisPosts[i] = types.NewGenesisPost(maxAttachmentID+1, post)
	}
	return genesisPosts
}
