# Performing transactions
## Introduction
As seeing inside the [FAQs](developer-faq.md#what-is-a-transaction), transactions are the way to alter the current chain state by providing it with the actions to take. Inside this page you will see all the messages that represents the available actions that can be used to edit the chain state.  

## Available messages
Here is the list of currently available [messages](developer-faq.md#what-is-a-message) that you can use while creating transactions for the Desmos chain. 

### Sessions
* [`MsgCreateSession`](msgs/create-session.md): allows you to create a new session binding an existing account on another chain to a Desmos account. 

### Posts
* [`MsgCreatePost`](msgs/create-post.md): allows you to create a new post or a comment for an existing post. 
* [`MsgLikePost`](msgs/like-post.md): allows you to like an existing post. 
* [`MsgUnlikePost`](msgs/unlike-post.md): allows you to remove the link from a previously liked post. 