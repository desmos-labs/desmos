---
id: concepts
title: Concepts
sidebar_label: Concepts
slug: concepts
---

# Concepts 

## Post
Inside Desmos, a post represents a single piece of content within a subspace. This can be used to represent things such as a blog post, a tweet, or anything else. A post can also be used to represents a quote of another post, a comment or a re-post of another content.

Aside from the inner text, developers can put on chain some metadata that can be useful during the visualization of the post, such as (but not limited to): links, users mentions or hashtags.

### Subspace ID
The subspace id specifies the id of the subspace inside which the post is placed.

### Section ID 
The [section] ID indicates the ID of the subspace's section where the post lives.

### ID
The unique ID that identifies the post itself. This ID is automatically assigned to the post at the moment of its 
creation in a sequential way (e.g. if there's 4 posts in the chain, the one we are creating will have id equal to 5).

### External ID (Optional)
External ID indicates and external ID attached to the post. It is represented by a string.

### Text (Optional)
The text is the actual textual content of the post. It has a fixed max length that is determined by an on-chain governance parameter.

### Entities (Optional)
Entities are particular parts of the text that can be parsed out of it in order to be displayed in custom ways.
Entities are divided in 3 different categories:
- Hashtags (i.e. #desmos)
- Mentions (i.e. @desmos1xcfui...., @Forbole)
- Urls (i.e. https://desmos.network, ftp://user:password@example.com/image.png)

#### TextTag
Both hashtags and mentions are represented as a `TextTag`. The `TextTag` structure contains the necessary fields that ease the process of 
text's parsing.

##### Start
It is the text's index that indicates the exact point where the tag start and the text should be parsed.

##### End
It is the text's index that indicates the exact point where the tag ends and the last character where to parse the text.

##### Tag
The actual value of the tag. It can be a user address, a hashtag value or anything else you want to emphasize in your post.

#### Url
This field contains the details of a generic URL. There is no validation on it, so you can implement your own, deciding which
one to accept and which not.

##### Start
It is the text's index that indicates the exact point where the url start and the text should be parsed.

##### End
It is the text's index that indicates the exact point where the url ends and the last character where to parse the text.

##### Url
The actual value of the url where the user should be redirected to.

##### Display url (Optional)
The visual url that should be displayed in the app. Ideally something that shorten or make an url more human-readable.

### Author
The address of the author of the post.

### Conversation ID (optional)
The ID of the original post of the conversation. This ID identifies the parent post (if existent) where the conversation
started.

### Referenced Posts
Referenced posts are represented by a `PostReference` structure.

#### PostReference
A reference to an external post, usable for reply, repost or quote purposes.

##### Type
This is the type of the reference. It can be one of the following values:

| **Name**                          | **Description**                                          |  
|:----------------------------------|:---------------------------------------------------------|
| `POST_REFERENCE_TYPE_UNSPECIFIED` | No reference specified                                   |
| `POST_REFERENCE_TYPE_REPLY`       | This reference represents a reply to the specified post  |
| `POST_REFERENCE_TYPE_QUOTE`       | This reference represents a quote of the specified post  |
| `POST_REFERENCE_TYPE_REPOST`      | This reference represents a repost of the specified post |

##### Post ID
The ID of the referenced post.

##### Position
This field should be used only when the `Type` field is equal to `TYPE_QUOTE` to indicates where the `PostReference`
start in the post's text.

### Reply Setting
This field contains the possible reply settings that a post can have. It can be one of the following values:

| **Name**                    | **Description**                                              |  
|:----------------------------|:-------------------------------------------------------------|
| `REPLY_SETTING_UNSPECIFIED` | No reference specified                                       |
| `REPLY_SETTING_EVERYONE`    | This reference represents a reply to the specified post      |
| `REPLY_SETTING_FOLLOWERS`   | This reference represents a quote of the specified post      |
| `REPLY_SETTING_MUTUAL`      | This reference represents a repost of the specified post     |
| `REPLY_SETTING_MENTIONS`    | Only people mentioned inside this post will be able to reply |

### Creation Date
The creation date of the post.

### Last Edited Date
The las time the post has been edited.

## Attachment
An attachment represent something that can be added to a post in order to enrich it or give it some additional use cases.

### Subspace ID
The [subspace] ID indicates the ID of the Dapp where the attachment is hosted and lives.

### Post ID
The [post](#Post) ID to which the attachment is linked.

### ID
The unique ID that identifies the attachment. This ID is automatically assigned the same way the post one is.

### Content
The content of the attachment. It can be either:
- A media content (pictures, clips, videos, GIFs);
- A poll.

## Media
The Media structure represent a media content of any kind from pictures, to clips, videos, GIFs, etc... attached to 
a post.

### URI
The URI address referencing the media.

### Mime type
The [mime type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types) of the media.

## Poll
The Poll structure represents a poll attached to a post.

### Question
The question field define the question of the poll.

### Provided answers
The possible answers choices the users have to reply the poll. They are represented by a `ProvidedAnswer` structure.

#### Provided Answer
The representation of a provided answer for a poll.

##### Text (Optional)
The text of the answer.

##### Attachments (Optional)
The [attachments](#attachment) of the answer. If not provided, a text answer has to be specified.

### End Date
The date when the poll will close.

### Allow Multiple Answers
This field tells if the poll allows multiple answers or not.

### Allow Answer Edits
This field tells if the poll allows users to edit their answers or not.

### Final Tally Results 
This fields contains the final results of the poll.

#### Results
The answers' results represent by the `AnswerResult` structure.

##### Answer Result
This field contains the result of a single poll provided answer

###### Answer Index
The index of the answer within the post's `ProvidedAnswer` field associated to which these results.

###### Votes
The number of votes received by an answer.

## User Answer
The user answer represent an answer given by a user to a poll.

#### Subspace ID
The subspace id represents the id of the subspace inside which is placed the post that contains the poll to which this answer refers. 

#### Post ID
Id of the post that contains the poll to which this answer refers.

#### Poll ID
Id of the poll to which this answer refers. 

#### Answer Indexes
The answer indexes contains a list of user answers, each one identified by the index of the chosen option within the poll's `ProvidedAnswer` array.

#### User 
The address of the user answering the poll.