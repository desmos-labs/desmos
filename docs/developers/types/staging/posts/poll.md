# Poll
The `Poll` object allows to specify the details of a poll that should be associated to a post. Please note that it
is **not** necessary to associate a poll to each post. Instead, if you want to create a [`Post`](post.md) without any
poll associated to it, you simply have to use the `nil` value for this field.

Following you will find a description for all the contained field of the `Poll` object.

## `Question`
This field contains the question that should be associated with the poll. It currently has no checks associated to it a part from the non-empty check. 

## `ProvidedAnswers`
This field allows to specify a list of answers that are provided to the users willing to answer the  poll.

Each answer should be composed of two attributes: 

- `ID`, which identifies uniquely inside the answers' list that answer.
- `Text`, which contains the text of the answer itself. 

The minimum number of answers that a poll must have is 2. 

## `EndDate`
The `EndDate` field allows you to specify the date after which the poll should be considered closed and no longer accepting the answers. 

## `Open` 
This field tells whether the poll is still open and accepting new answers from users or not. Please note that the default value for this field is `true` and trying to create a poll with it set to `false` will result in an error. 

During the chain execution, this field will be automatically changed to `false` when the `EndDate` is passed.

## `AllowsMultipleAnswers`
This field allows to specify whether or not the poll allows multiple answers from the same user. If set to `true`, the users will be able to specify more than one answer to the same poll. Otherwise, if set to `false`, each user will be allowed to answer with only one option to the poll. 

## `AllowsAnswerEdits`
By setting this field to `true`, you will allow users to change their mind while the poll is still open, allowing them to change their answer(s). If set to `false`, they will not be able to do so and their final answer(s) will be the first one they give.      
