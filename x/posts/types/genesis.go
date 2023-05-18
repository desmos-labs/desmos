package types

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec/types"
)

// NewGenesisState returns a new GenesisState instance
func NewGenesisState(
	subspacesData []SubspaceDataEntry,
	posts []Post,
	postsData []PostDataEntry,
	attachments []Attachment,
	activePolls []ActivePollData,
	userAnswers []UserAnswer,
	params Params,
	transferRequests []PostOwnerTransferRequest,
) *GenesisState {
	return &GenesisState{
		SubspacesData:             subspacesData,
		Posts:                     posts,
		PostsData:                 postsData,
		Attachments:               attachments,
		ActivePolls:               activePolls,
		UserAnswers:               userAnswers,
		Params:                    params,
		PostOwnerTransferRequests: transferRequests,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, nil, nil, nil, nil, DefaultParams(), nil)
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (e GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, report := range e.Attachments {
		err := report.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, entry := range data.SubspacesData {
		if containsDuplicatedSubspaceDataEntry(data.SubspacesData, entry) {
			return fmt.Errorf("duplicated subspace data entry: %d", entry.SubspaceID)
		}

		err := entry.Validate()
		if err != nil {
			return err
		}
	}

	for _, post := range data.Posts {
		if containsDuplicatedPost(data.Posts, post) {
			return fmt.Errorf("duplicated post: subspace id %d, post id %d", post.SubspaceID, post.ID)
		}

		err := post.Validate()
		if err != nil {
			return err
		}
	}

	for _, entry := range data.PostsData {
		if containsDuplicatedPostDataEntry(data.PostsData, entry) {
			return fmt.Errorf("duplicated post data entry: subspace id %d, post id %d", entry.SubspaceID, entry.PostID)
		}

		err := entry.Validate()
		if err != nil {
			return err
		}
	}

	for _, attachment := range data.Attachments {
		if containsDuplicatedAttachment(data.Attachments, attachment) {
			return fmt.Errorf("duplicated attachment: subspace id %d, post id %d, attachment id %d",
				attachment.SubspaceID, attachment.PostID, attachment.ID)
		}

		err := attachment.Validate()
		if err != nil {
			return err
		}
	}

	for _, pollData := range data.ActivePolls {
		if containsDuplicatedPollData(data.ActivePolls, pollData) {
			return fmt.Errorf("duplicated poll data: subspace id %d, post id %d, poll id %d",
				pollData.SubspaceID, pollData.PostID, pollData.PollID)
		}

		err := pollData.Validate()
		if err != nil {
			return err
		}
	}

	for _, answer := range data.UserAnswers {
		if containsDuplicatedAnswer(data.UserAnswers, answer) {
			return fmt.Errorf("duplicated user answer: subspace id %d, post id %d, poll id %d, user: %s",
				answer.SubspaceID, answer.PostID, answer.PollID, answer.User)
		}

		err := answer.Validate()
		if err != nil {
			return err
		}
	}

	err := data.Params.Validate()
	if err != nil {
		return err
	}

	for _, request := range data.PostOwnerTransferRequests {
		if containDuplicatedRequest(data.PostOwnerTransferRequests, request) {
			return fmt.Errorf("duplicated post owner transfer request: subspace id %d, post id %d",
				request.SubspaceID, request.PostID)
		}

		err := request.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// containsDuplicatedSubspaceDataEntry tells whether the given entries slice contains
// two or more entries for the same subspace
func containsDuplicatedSubspaceDataEntry(entries []SubspaceDataEntry, entry SubspaceDataEntry) bool {
	var count = 0
	for _, s := range entries {
		if s.SubspaceID == entry.SubspaceID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedPostDataEntry tells whether the given entries slice contains
// two or more entries for the same post
func containsDuplicatedPostDataEntry(entries []PostDataEntry, entry PostDataEntry) bool {
	var count = 0
	for _, s := range entries {
		if s.SubspaceID == entry.SubspaceID && s.PostID == entry.PostID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedPost tells whether the given posts slice contains two or more posts
// having the same id of the given one
func containsDuplicatedPost(posts []Post, post Post) bool {
	var count = 0
	for _, s := range posts {
		if s.SubspaceID == post.SubspaceID && s.ID == post.ID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedAttachment tells whether the given attachments slice contains two or more attachments
// having the same id of the given one
func containsDuplicatedAttachment(attachments []Attachment, attachment Attachment) bool {
	var count = 0
	for _, s := range attachments {
		if s.SubspaceID == attachment.SubspaceID && s.PostID == attachment.PostID && s.ID == attachment.ID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedPollData tells whether the given polls data slice contains two or more data
// having the same id of the given one
func containsDuplicatedPollData(pollsData []ActivePollData, data ActivePollData) bool {
	var count = 0
	for _, d := range pollsData {
		if d.SubspaceID == data.SubspaceID && d.PostID == data.PostID && d.PollID == data.PollID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedAnswer tells whether the given user answers slice contains two or more answers
// by the same user as the given one
func containsDuplicatedAnswer(answers []UserAnswer, answer UserAnswer) bool {
	var count = 0
	for _, s := range answers {
		if s.SubspaceID == answer.SubspaceID && s.PostID == answer.PostID && s.PollID == answer.PollID && s.User == answer.User {
			count++
		}
	}
	return count > 1
}

// containDuplicatedRequest tells whether the given post owner transfer request slice contains two or more given
// post owner transfer request by the same id as the given one
func containDuplicatedRequest(requests []PostOwnerTransferRequest, request PostOwnerTransferRequest) bool {
	var count = 0
	for _, r := range requests {
		if r.SubspaceID == request.SubspaceID && r.PostID == request.PostID {
			count++
		}
	}
	return count > 1
}

// --------------------------------------------------------------------------------------------------------------------

// NewSubspaceDataEntry returns a new SubspaceDataEntry instance
func NewSubspaceDataEntry(subspaceID uint64, initialPostID uint64) SubspaceDataEntry {
	return SubspaceDataEntry{
		SubspaceID:    subspaceID,
		InitialPostID: initialPostID,
	}
}

// Validate returns an error if something is wrong within the entry data
func (e SubspaceDataEntry) Validate() error {
	if e.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", e.SubspaceID)
	}

	if e.InitialPostID == 0 {
		return fmt.Errorf("invalid initial post id: %d", e.InitialPostID)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewPostDataEntry returns a new PostDataEntry instance
func NewPostDataEntry(subspaceID uint64, postID uint64, attachmentID uint32) PostDataEntry {
	return PostDataEntry{
		SubspaceID:          subspaceID,
		PostID:              postID,
		InitialAttachmentID: attachmentID,
	}
}

// Validate returns an error if something is wrong within the entry
func (p PostDataEntry) Validate() error {
	if p.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", p.SubspaceID)
	}

	if p.PostID == 0 {
		return fmt.Errorf("invalid post id: %d", p.PostID)
	}

	if p.InitialAttachmentID == 0 {
		return fmt.Errorf("invalid initial attachment id: %d", p.InitialAttachmentID)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewActivePollData returns a new ActivePollData instance
func NewActivePollData(subspaceID uint64, postID uint64, pollID uint32, endTime time.Time) ActivePollData {
	return ActivePollData{
		SubspaceID: subspaceID,
		PostID:     postID,
		PollID:     pollID,
		EndDate:    endTime,
	}
}

// Validate returns an error if something is wrong with the poll data
func (d ActivePollData) Validate() error {
	if d.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id %d", d.SubspaceID)
	}

	if d.PostID == 0 {
		return fmt.Errorf("invalid post id %d", d.PostID)
	}

	if d.PollID == 0 {
		return fmt.Errorf("invalid poll id %d", d.PollID)
	}

	if d.EndDate.IsZero() {
		return fmt.Errorf("invalid end time: %s", d.EndDate)
	}

	return nil
}
