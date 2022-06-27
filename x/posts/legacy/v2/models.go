package v2

// DONTCOVER

import "time"

// NewPost allows to build a new Post instance
func NewPost(
	subspaceID uint64,
	sectionID uint32,
	id uint64,
	externalID string,
	text string,
	author string,
	conversationID uint64,
	entities *Entities,
	referencedPosts []PostReference,
	replySetting ReplySetting,
	creationDate time.Time,
	lastEditedDate *time.Time,
) Post {
	return Post{
		SubspaceID:      subspaceID,
		SectionID:       sectionID,
		ID:              id,
		ExternalID:      externalID,
		Text:            text,
		Entities:        entities,
		Author:          author,
		ConversationID:  conversationID,
		ReferencedPosts: referencedPosts,
		ReplySettings:   replySetting,
		CreationDate:    creationDate,
		LastEditedDate:  lastEditedDate,
	}
}

// NewEntities returns a new Entities instance
func NewEntities(hashtags []Tag, mentions []Tag, urls []Url) *Entities {
	return &Entities{
		Hashtags: hashtags,
		Mentions: mentions,
		Urls:     urls,
	}
}

// NewTag returns a new Tag instance
func NewTag(start, end uint64, tag string) Tag {
	return Tag{
		Start: start,
		End:   end,
		Tag:   tag,
	}
}

// NewURL returns a new Url instance
func NewURL(start, end uint64, url, displayURL string) Url {
	return Url{
		Start:      start,
		End:        end,
		Url:        url,
		DisplayUrl: displayURL,
	}
}

// NewPostReference returns a new PostReference instance
func NewPostReference(referenceType PostReferenceType, postID uint64, position uint64) PostReference {
	return PostReference{
		Type:     referenceType,
		PostID:   postID,
		Position: position,
	}
}
