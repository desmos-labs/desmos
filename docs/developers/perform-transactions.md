# Performing transactions
## Introduction
As seeing inside the [FAQs](developer-faq.md#what-is-a-transaction), transactions are the way to alter the current chain state by providing it with the actions to take. Inside this page you will see all the messages that represents the available actions that can be used to edit the chain state.  

## Available messages
Here is the list of currently available [messages](developer-faq.md#what-is-a-message) that you can use while creating transactions for the Desmos chain. 

### Sessions
* [`MsgCreateSession`](msgs/create-session.md): allows you to create a new session binding an existing account on another chain to a Desmos account. 

### Posts
* [`MsgCreatePost`](msgs/create-post.md): allows you to create a new post or a comment for an existing post. 
* [`MsgEditPost`](msgs/edit-post.md): allows you to edit a previously created post message.
* [`MsgAddPostReaction`](msgs/add-reaction.md): allows you to add a reaction to an existing post. 
* [`MsgRemoveReaction`](msgs/remove-reaction.md): allows you to remove a reaction from a post.
* [`MsgAnswerPoll`](msgs/answer-poll.md): allows you to answer a post's poll.
