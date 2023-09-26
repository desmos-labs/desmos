package v3

import (
	"fmt"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2 "github.com/desmos-labs/desmos/v6/x/posts/legacy/v2"
	v4 "github.com/desmos-labs/desmos/v6/x/posts/legacy/v4"
	"github.com/desmos-labs/desmos/v6/x/posts/types"
)

// MigrateStore performs in-place store migrations from v2 to v3.
// During the migration, the following operations are performed:
// - convert all the existing posts
// - convert all the existing attachments
// - convert all the existing user answers
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	// Migrate all the posts
	err := migratePosts(store, cdc)
	if err != nil {
		return err
	}

	// Migrate the attachments
	err = migrateAttachments(store, cdc)
	if err != nil {
		return err
	}

	// Migrate the user answers
	err = migrateUserAnswers(store, cdc)
	if err != nil {
		return err
	}

	return nil
}

// migratePosts migrates the posts preset inside the store from v2 to v3
func migratePosts(store storetypes.KVStore, cdc codec.BinaryCodec) error {
	prefixStore := prefix.NewStore(store, types.PostPrefix)
	iterator := prefixStore.Iterator(nil, nil)

	// Get all the posts
	var v2Posts []v2.Post
	for ; iterator.Valid(); iterator.Next() {
		var v2Post v2.Post
		cdc.MustUnmarshal(iterator.Value(), &v2Post)
		v2Posts = append(v2Posts, v2Post)
	}

	// Close the iterator
	err := iterator.Close()
	if err != nil {
		return err
	}

	// Convert the posts
	for _, v2Post := range v2Posts {
		v3Post := v4.NewPost(
			v2Post.SubspaceID,
			v2Post.SectionID,
			v2Post.ID,
			v2Post.ExternalID,
			v2Post.Text,
			v2Post.Author,
			v2Post.ConversationID,
			migrateEntities(v2Post.Entities),
			nil,
			migratePostReferences(v2Post.ReferencedPosts),
			migrateReplySettings(v2Post.ReplySettings),
			v2Post.CreationDate,
			v2Post.LastEditedDate,
		)

		// Save the post
		store.Set(types.PostStoreKey(v3Post.SubspaceID, v3Post.ID), cdc.MustMarshal(&v3Post))
	}

	return nil
}

// migrateEntities migrates the given entities from v2 to v3
func migrateEntities(v2Entities *v2.Entities) *v4.Entities {
	if v2Entities == nil {
		return nil
	}

	return v4.NewEntities(
		migrateTags(v2Entities.Hashtags),
		migrateTags(v2Entities.Mentions),
		migrateUrls(v2Entities.Urls),
	)
}

// migrateTags migrates the given tags from v2 to v3
func migrateTags(v2Tags []v2.Tag) []v4.TextTag {
	if v2Tags == nil {
		return nil
	}

	v3Tags := make([]v4.TextTag, len(v2Tags))
	for i, v2Tag := range v2Tags {
		v3Tags[i] = v4.NewTextTag(v2Tag.Start, v2Tag.End, v2Tag.Tag)
	}
	return v3Tags
}

// migrateUrls migrates the given urls from v2 to v3
func migrateUrls(v2Urls []v2.Url) []v4.Url {
	if v2Urls == nil {
		return nil
	}

	v3Urls := make([]v4.Url, len(v2Urls))
	for i, v2Url := range v2Urls {
		v3Urls[i] = v4.NewURL(v2Url.Start, v2Url.End, v2Url.Url, v2Url.DisplayUrl)
	}
	return v3Urls
}

// migratePostReferences migrates the given references from v2 to v3
func migratePostReferences(v2References []v2.PostReference) []v4.PostReference {
	if v2References == nil {
		return nil
	}

	v3References := make([]v4.PostReference, len(v2References))
	for i, v2Reference := range v2References {
		v3References[i] = v4.NewPostReference(
			migratePostReferenceType(v2Reference.Type),
			v2Reference.PostID,
			v2Reference.Position,
		)
	}
	return v3References
}

// migratePostReferenceType migrates the given post reference type from v2 to v3
func migratePostReferenceType(v2Type v2.PostReferenceType) v4.PostReferenceType {
	switch v2Type {
	case v2.POST_REFERENCE_TYPE_UNSPECIFIED:
		return v4.POST_REFERENCE_TYPE_UNSPECIFIED
	case v2.POST_REFERENCE_TYPE_REPLY:
		return v4.POST_REFERENCE_TYPE_REPLY
	case v2.POST_REFERENCE_TYPE_QUOTE:
		return v4.POST_REFERENCE_TYPE_QUOTE
	case v2.POST_REFERENCE_TYPE_REPOST:
		return v4.POST_REFERENCE_TYPE_REPOST
	default:
		panic(fmt.Errorf("invalid post reference type: %s", v2Type))
	}
}

// migrateReplySettings migrates the given reply setting from v2 to v3
func migrateReplySettings(settings v2.ReplySetting) v4.ReplySetting {
	switch settings {
	case v2.REPLY_SETTING_UNSPECIFIED:
		return v4.REPLY_SETTING_UNSPECIFIED
	case v2.REPLY_SETTING_EVERYONE:
		return v4.REPLY_SETTING_EVERYONE
	case v2.REPLY_SETTING_FOLLOWERS:
		return v4.REPLY_SETTING_FOLLOWERS
	case v2.REPLY_SETTING_MUTUAL:
		return v4.REPLY_SETTING_MUTUAL
	case v2.REPLY_SETTING_MENTIONS:
		return v4.REPLY_SETTING_MENTIONS
	default:
		panic(fmt.Errorf("invalid reply settings value: %s", settings))
	}
}

// migrateAttachments migrates the attachments present inside the store from v2 to v3
func migrateAttachments(store storetypes.KVStore, cdc codec.BinaryCodec) error {
	prefixStore := prefix.NewStore(store, types.AttachmentPrefix)
	iterator := prefixStore.Iterator(nil, nil)

	// Get all the attachments
	var v2Attachments []v2.Attachment
	for ; iterator.Valid(); iterator.Next() {
		var v2Attachment v2.Attachment
		cdc.MustUnmarshal(iterator.Value(), &v2Attachment)
		v2Attachments = append(v2Attachments, v2Attachment)
	}

	// Close the iterator
	err := iterator.Close()
	if err != nil {
		return err
	}

	// Convert the attachments
	v3Attachments := convertAttachments(v2Attachments)
	for i, v3Attachment := range v3Attachments {
		// Save the attachment
		store.Set(
			types.AttachmentStoreKey(v3Attachment.SubspaceID, v3Attachment.PostID, v3Attachment.ID),
			cdc.MustMarshal(&v3Attachments[i]),
		)
	}

	return nil
}

// convertAttachments converts the given attachments from v2 to v3
func convertAttachments(v2Attachment []v2.Attachment) []v4.Attachment {
	v3Attachments := make([]v4.Attachment, len(v2Attachment))
	for i, attachment := range v2Attachment {
		v3Attachments[i] = v4.NewAttachment(
			attachment.SubspaceID,
			attachment.PostID,
			attachment.ID,
			convertAttachmentContent(attachment.Content),
		)
	}
	return v3Attachments
}

// convertAttachmentContent converts the given attachment content from v2 to v3
func convertAttachmentContent(contentAny *codectypes.Any) v4.AttachmentContent {
	switch content := contentAny.GetCachedValue().(v2.AttachmentContent).(type) {
	case *v2.Media:
		return v4.NewMedia(content.Uri, content.MimeType)

	case *v2.Poll:
		return v4.NewPoll(content.Question,
			convertProvidedAnswers(content.ProvidedAnswers),
			content.EndDate,
			content.AllowsMultipleAnswers,
			content.AllowsAnswerEdits,
			convertTallyResults(content.FinalTallyResults),
		)

	default:
		panic(fmt.Errorf("invalid content type: %T", contentAny.GetCachedValue()))
	}
}

// convertTallyResults converts the given poll tally results from v2 to v3
func convertTallyResults(v2Results *v2.PollTallyResults) *v4.PollTallyResults {
	if v2Results == nil {
		return nil
	}

	v3Results := make([]v4.PollTallyResults_AnswerResult, len(v2Results.Results))
	for i, result := range v2Results.Results {
		v3Results[i] = v4.NewAnswerResult(result.AnswerIndex, result.Votes)
	}
	return v4.NewPollTallyResults(v3Results)
}

// convertProvidedAnswers converts the given poll provided answers from v2 to v3
func convertProvidedAnswers(v2Answers []v2.Poll_ProvidedAnswer) []v4.Poll_ProvidedAnswer {
	v3Answers := make([]v4.Poll_ProvidedAnswer, len(v2Answers))
	for i, answer := range v2Answers {
		v3Answers[i] = v4.NewProvidedAnswer(answer.Text, convertAttachments(answer.Attachments))
	}
	return v3Answers
}

// migrateUserAnswers migrates all the user answers present inside the store from v2 to v3
func migrateUserAnswers(store storetypes.KVStore, cdc codec.BinaryCodec) error {
	prefixStore := prefix.NewStore(store, types.UserAnswerPrefix)
	iterator := prefixStore.Iterator(nil, nil)

	// Get all the answers
	var v2Answers []v2.UserAnswer
	for ; iterator.Valid(); iterator.Next() {
		var v2Answer v2.UserAnswer
		cdc.MustUnmarshal(iterator.Value(), &v2Answer)
		v2Answers = append(v2Answers, v2Answer)
	}

	// Close the iterator
	err := iterator.Close()
	if err != nil {
		return err
	}

	// Convert the answers
	for _, v2Answer := range v2Answers {
		v3Answer := v4.NewUserAnswer(
			v2Answer.SubspaceID,
			v2Answer.PostID,
			v2Answer.PollID,
			v2Answer.AnswersIndexes,
			v2Answer.User,
		)

		// Save the attachment
		store.Set(types.PollAnswerStoreKey(v3Answer.SubspaceID, v3Answer.PostID, v3Answer.PollID, v3Answer.User), cdc.MustMarshal(&v3Answer))
	}

	return nil
}
