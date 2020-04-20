# Query poll answers made to a post's poll
This query endpoint allows you to retrieve the details of answers made to a post's poll'. 

**CLI**
 ```bash
desmoscli query posts poll-answers [id]

# Example
# desmoscli query posts poll-answers a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc
``` 

**REST**
```
/posts/{postId}/poll-answers

# Example
# curl http://lcd.morpheus.desmos.network:1317/posts/a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc/poll-answers
```
