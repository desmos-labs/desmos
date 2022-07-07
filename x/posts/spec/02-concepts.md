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
The section id specifies the id of the subspace section where the post is placed.

### ID
A post id is the unique identifier of that post within the subspace. Each post id is generated automatically when handling a `MsgCreatePost` and can be used along with the subspace id itself to uniquely identify a post within Desmos.

### External ID (Optional)
A post external id is an optional text field that can be useful to developers that want to link this post to an external data storage. As an example, if a developer wants to store the post content on their own data storage what they can do is use the external id field to tell how the content should be retrieved.

### Text (Optional)
A post text is the actual textual content of the post. It has a fixed max length that is determined by an on-chain governance parameter. Any post with a text length greater than the current max length allowed will be considered invalid and not stored on the chain (an error will be returned during the saving). To store large text posts we recommend storing the post text contents on an external storage and then using either the text or external id fields to specify how to retrieve them.

### Entities (Optional)
Entities represent part of the post's text content that should be rendered in a particular way. These include hashtags, mentions to other users or links.

#### TextTag
Both hashtags and mentions are represented as a `TextTag`. The `TextTag` structure contains the necessary fields that ease the process of 
text's parsing.

##### Start
Index within the post text at which the tag starts.

##### End
Index within the post text at which the tag ends.

##### Tag
The actual value of the tag. Usually this is going to be either the hashtag value, or the address of the mentioned user (if within the post text the DTag is used to reference the user).

#### Url
Within Desmos there is no URL validation, so this type can be used to represent any kind of URL that might be present inside the post text.

##### Start
Index within the post text at which the URL starts.

##### End
Index within the post text at which the URL ends.

##### Url
The actual value of the url where the user should be redirected to.

##### Display url (Optional)
Value to be displayed within the post text instead of the URL value itself. This might be useful if you want to shorten a link or if you want to represent a clickable text.

### Author
Address of the post author.

### Conversation ID (optional)
In order to be able to display a conversation properly, you can use the conversation id to specify the id of the post that should be considered as the first post of the conversation itself. As an example, if a thread is started then all posts referring to that thread should have the same conversation, even if they are answering other posts. This is to make sure they can be connected to the original conversation and displayed within it.

### Referenced Posts
In order to represent various interactions with a post, inside Desmos we use post references. All the references that a post has to other posts should be placed inside the referenced posts array.

#### PostReference
A reference to an external post. This can be used to represent a reply (comment), repost or quote of another post.

##### Type
Type of the post reference. It can be one of the following values:

| **Name**                          | **Description**                                          |  
|:----------------------------------|:---------------------------------------------------------|
| `POST_REFERENCE_TYPE_REPLY`       | This reference represents a reply to the specified post  |
| `POST_REFERENCE_TYPE_QUOTE`       | This reference represents a quote of the specified post  |
| `POST_REFERENCE_TYPE_REPOST`      | This reference represents a repost of the specified post |

##### Post ID
The id of the referenced post. Note that it's possible to reference only posts within the same subspace. Trying to reference a post within another subspace will return an error or might end up in unwanted behavior. 

##### Position
If the post reference represents a quote of a post, it might be useful to be able to specify the index within the text at which the quoted post should be rendered. This is particularly useful if a post should quote more than one post: with different positions application will be able to render the quoted post properly. 

Note that trying to set this field to a non-zero value with any kind of reference aside `TYPE_QUOTE` will end up in the post being considered invalid.

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
An attachment represents any kind of media that can be attached to a post, such as (but not limited to): an image, a GIF, or a poll.

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