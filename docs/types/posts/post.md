# Post
Inside Desmos, posts are the way that users have to share publicly whatever they want. 

You can see posts much like tweets on Twitter, as they have the same functionality: they allow you to write what you wish (without any length limitation) and they are publicly visible to all the Desmos users.  

The only difference with tweets is that once you've created a Desmos post you will **not** be able to delete it! This is due to the blockchain's intrinsic characteristic of being immutable: each transaction that is performed cannot be undone.

## Contained data
A post is made of different parts. Following you will find out what are those and how they can be used. 

### `PostID`
The `PostID` identifies uniquely any single post. It cannot be specified in any way but instead it is given to posts when they are stored inside the chain. 

### `ParentID`
The `ParentID` identifies the ID of the post to which the current post is a comment of. If this post is not a comment to any other post, then it's `ParentID` will be `0`. 

### `Message`
The `Message` represents the field that should be used to specify the textual content of a post. It must always have a length not exceeding 500 characters, or an error will be thrown while inserting it inside the chain.

### `Created`
The `Created`field must be used in order to specify the creation date of the post. It must be an [RCF3339]()-formatted date.  

### `LastEdited`
The `LastEdited` field should be specified only when editing posts. It is used in order to tell in which date and time the post has been edited for the last time. 

### `AllowsComments`
Inside Desmos we allow users to decide on their own whether their posts should accept comments or not. In order to do so, we have created the `AllowsComments` field. If this field is set to `false` on a specific post, an error will be thrown when trying to comment on such post. On the other hand, if it is set to `true` (default value) then the post will always accept comments. 

### `Subspace`
As Desmos is thought to be a protocol on top of which many applications can be developed, the `Subspace` field identifies the application inside which the message should be seen. Currently the subspace must be a SHA256 hash of the previously plain-text value.

A common value that is used when you don't want to crete your own is `4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e` which corresponds to the SHA256 hash of the plain-text `desmos`. 

If you instead prefer having a custom subspace, you can create your own by hashing any plain-text using any online SHA256 calculator such as [this one](https://emn178.github.io/online-tools/sha256.html).

### `OptionalData`
In order to allow developers to specify any arbitrary data they want inside a post, we've introduced the `OptionalData` field. This field is an array of objects containing a key and a value, and it allows inserting up to 10 fields containing any value you prefer.  

Please note that this field should be used only when strictly necessary as it might cause an unexpected chain state dimensions increment. Also, each value must be no longer than 200 characters, or an error will be thrown.

### `Creator`
The `Creator` field is used to specify the Bech32 address of the creator of the post. In order for a creator address to be valid, it must begin with the `desmos` Bech32 human-readable part. 

### `Attachments`
Starting from version `v0.3.0`, we've introduced the `Attachments` (previously called `Medias`) field. This contains a (possibly empty) array of attachment files that can be associated to a post. 

In order to know how an attachment object must be created, please refer to the [`Attachment` type documentation](./attachment.md)

### `PollData`
Along with the [`Attachments`](#attachments) field, with `v0.3.0` we've introduced the `PollData` field as well. This field allows to specify an optional poll that should be associated with the post itself. 

In order to better understand how the value of this field should be created, please refer to the [`PollData` type documentation](./post-poll-data.md) 
