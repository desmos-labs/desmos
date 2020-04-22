package v040

import (
	"strings"

	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

// Migrate accepts exported genesis state from v0.3.0 and migrates it to v0.4.0
// genesis state. This migration changes the way posts IDs are specified from a simple
// uint64 to a sha-256 hashed string
func Migrate(oldGenState v030posts.GenesisState) GenesisState {

	posts := migratePosts(oldGenState.Posts)

	answers, err := MigrateUsersAnswers(oldGenState.PollAnswers, oldGenState.Posts)
	if err != nil {
		panic(err)
	}

	reactions, err := MigratePostReactions(oldGenState.Reactions, oldGenState.Posts)
	if err != nil {
		panic(err)
	}

	registeredReactions, err := GetReactionsToRegister(posts, reactions)
	if err != nil {
		panic(err)
	}

	return GenesisState{
		Posts:               posts,
		UsersPollAnswers:    answers,
		PostReactions:       reactions,
		RegisteredReactions: registeredReactions,
	}
}

// ComputeParentID get the post related to the given parentID if exists and returns it computed ID.
// Returns "" otherwise
func ComputeParentID(posts []v030posts.Post, parentID v030posts.PostID) PostID {
	if parentID == v030posts.PostID(uint64(0)) {
		return ""
	}

	for _, post := range posts {
		if post.PostID == parentID {
			return ComputeID(post.Created, post.Creator, post.Subspace)
		}
	}

	//it should never reach this
	return ""
}

// migratePosts takes a slice of v0.3.0 Post object and migrates them to v0.4.0 Post.
// For each post, its id is converted from an uint64 representation to a SHA-256 string representation.
func migratePosts(posts []v030posts.Post) []Post {
	migratedPosts := make([]Post, len(posts))

	// Migrate the posts
	for index, oldPost := range posts {
		migratedPosts[index] = Post{
			PostID:         ComputeID(oldPost.Created, oldPost.Creator, oldPost.Subspace),
			ParentID:       ComputeParentID(posts, oldPost.ParentID),
			Message:        oldPost.Message,
			Created:        oldPost.Created,
			LastEdited:     oldPost.LastEdited,
			AllowsComments: oldPost.AllowsComments,
			Subspace:       oldPost.Subspace,
			OptionalData:   OptionalData(oldPost.OptionalData),
			Creator:        oldPost.Creator,
		}
	}

	return migratedPosts
}

// ConvertID take the given v030 post ID and convert it to a v040 post ID
func ConvertID(id string, posts []v030posts.Post) (postID PostID, error error) {
	for _, post := range posts {
		convertedID, err := v030posts.ParsePostID(id)
		if err != nil {
			return "", err
		}

		if post.PostID == convertedID {
			postID = ComputeID(post.Created, post.Creator, post.Subspace)
		}
	}
	return postID, nil
}

// MigrateUsersAnswers takes a slice of v0.3.0 UsersAnswers object and migrates them to v0.4.0 UserAnswers
func MigrateUsersAnswers(
	usersAnswersMap map[string][]v030posts.UserAnswer, posts []v030posts.Post,
) (map[string][]UserAnswer, error) {
	migratedUsersAnswers := make(map[string][]UserAnswer, len(usersAnswersMap))

	//Migrate the users answers
	for key, value := range usersAnswersMap {

		newUserAnswers := make([]UserAnswer, len(value))
		for index, userAnswers := range value {
			migratedAnswersIDs := make([]AnswerID, len(userAnswers.Answers))

			for index, answerID := range userAnswers.Answers {
				migratedAnswersIDs[index] = AnswerID(answerID)
			}
			newUserAnswers[index] = UserAnswer{
				Answers: migratedAnswersIDs,
				User:    userAnswers.User,
			}
		}

		postID, err := ConvertID(key, posts)
		if err != nil {
			return nil, err
		}

		migratedUsersAnswers[string(postID)] = newUserAnswers
	}

	return migratedUsersAnswers, nil
}

// GetReactionShortCodeFromValue retrieves the shortcode that should
// be associated to the reaction having the given value
func GetReactionShortCodeFromValue(originalValue string) (string, error) {
	// Make the value lowercase and replace - dividers with _
	value := strings.ToLower(originalValue)
	value = strings.ReplaceAll(value, "️", "")
	value = strings.ReplaceAll(value, "-", "_")

	// Try to get any present emoji by considering the value as the emoji itself
	if presentEmojis := emoji.FindAll(value); len(presentEmojis) > 0 {
		return presentEmojis[0].Match.(emoji.Emoji).Shortcodes[0], nil
	}

	value = strings.Split(value, " ")[0]
	if value == "like" {
		value = ":heart:"
	} else if value == "true" || value == "q" || strings.Contains(value, "nice") || strings.Contains(value, "well") {
		value = ":+1:"
	} else if value == ":grinning_face_with_star_eyes:" {
		value = ":star_struck:"
	} else if value == ":grinning_face_with_one_large_and_one_small_eye:" {
		value = ":zany_face:"
	} else if value == ":lion_face:" {
		value = ":lion:"
	}

	// Try to get the emoji by considering the value as the shortcode
	foundEmoji, err := emoji.LookupEmojiByCode(value)
	if err != nil {
		return "", err
	}

	return foundEmoji.Shortcodes[0], nil
}
