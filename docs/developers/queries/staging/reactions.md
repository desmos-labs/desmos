# Query registered reactions
This query allows you to retrieve the list of registered reactions inside an optional subspace.

**CLI**
 ```bash
desmos query posts registered-reactions [[subspace-id]]

# Example
# desmos query posts registered-reactions 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e
```

# Query post reactions
This query allows you to retrieve the list of reactions that have been added to a post.


**CLI**
```bash
desmos query posts reactions [post-id]

# Example
# desmos query posts reactions 301921ac3c8e623d8f35aef1886fea20849e49f08ec8ddfdd9b96feaf0c4fd15
```
