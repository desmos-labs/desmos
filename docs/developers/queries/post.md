# Query a post
This query endpoint allows you to retrieve the details of a single post having its id. 

**CLI**
 ```bash
desmoscli query posts post [id]

# Example
# desmoscli query posts post a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc
``` 

**REST**
```
/posts/{postId}

# Example
# curl https://morpheus4000.desmos.network/posts/a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc
```
