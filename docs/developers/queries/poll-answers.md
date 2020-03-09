# Query poll answers made to a post's poll
This query endpoint allows you to retrieve the details of answers made to a post's poll'. 

**CLI**
 ```bash
desmoscli query posts poll-answers [id]

# Example
# desmoscli query posts poll-answers 1
``` 

**REST**
```
/posts/{postId}/poll-answers

# Example
# curl https://morpheus1000.desmos.network/posts/10/poll-answers
```
