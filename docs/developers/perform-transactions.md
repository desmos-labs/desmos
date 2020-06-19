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
* [`MsgAddPostReaction`](msgs/add-post-reaction.md): allows you to add a reaction to an existing post. 
* [`MsgRemovePostReaction`](msgs/remove-post-reaction.md): allows you to remove a reaction from a post.
* [`MsgAnswerPoll`](msgs/answer-poll.md): allows you to answer a post's poll.
* [`MsgRegisterReaction`](msgs/register-reaction.md): allows you to register a reaction.

### Profile
* [`MsgSaveProfile`](docs/developers/msgs/save-profile.md): allows you to create or edit an existing profile.
* [`MsgDeleteProfile`](msgs/delete-profile.md): allows you to delete an existing profile.
* [`EditNameSurnameParamsProposal`](msgs/name-surname-params-edit-proposal.md): allows you to create a proposal to change
profile name and surname parameters.
* [`EditMonikerParamsProposal`](msgs/moniker-params-edit-proposal.md): allows you to create a proposal to change
profile moniker parameters.
* [`EditBioParamsProposal`](msgs/bio-params-edit-proposal.md): allows you to create a proposal to change
profile biography parameters.

### Reports
* [`MsgReportPost`](docs/developers/msgs/report-post.md): allows you to report an existing post.