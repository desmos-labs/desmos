package v090

import (
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v080posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.8.0"
)

// Migrate accepts exported genesis state from v0.8.0 and migrates it to v0.9.0
// genesis state. This migration replace all the old post media structure
// with the new one renamed to attachment.
func Migrate(oldGenState v080posts.GenesisState) GenesisState {
	return GenesisState{
		Posts:               ConvertPosts(oldGenState.Posts),
		UsersPollAnswers:    oldGenState.UsersPollAnswers,
		PostReactions:       oldGenState.PostReactions,
		RegisteredReactions: oldGenState.RegisteredReactions,
		Params:              oldGenState.Params,
	}
}

// ConvertPosts v0.8.0 posts into v0.9.0 posts
func ConvertPosts(oldPosts []v040posts.Post) (posts []Post) {
	for _, post := range oldPosts {
		np := Post{
			PostID:         post.PostID,
			ParentID:       post.ParentID,
			Message:        post.Message,
			Created:        post.Created,
			LastEdited:     post.LastEdited,
			AllowsComments: post.AllowsComments,
			Subspace:       post.Subspace,
			OptionalData:   post.OptionalData,
			Creator:        post.Creator,
			Attachments:    ConvertMediasToAttachments(post.Medias),
			PollData:       post.PollData,
		}
		posts = append(posts, np)
	}
	return posts
}

// ConvertMediasToAttachments converts v0.8.0 post medias into v0.9.0 attachments
func ConvertMediasToAttachments(medias []v040posts.PostMedia) (atts []Attachment) {
	for _, media := range medias {
		attachment := Attachment{
			URI:      media.URI,
			MimeType: media.MimeType,
			Tags:     nil,
		}
		atts = append(atts, attachment)
	}
	return atts
}
