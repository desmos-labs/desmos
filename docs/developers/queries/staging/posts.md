# Query a post
This query allows you to retrieve the details of a single post having its id.

**CLI**
 ```bash
desmos query posts post [id]

# Example
# desmos query posts post a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc
```

# Query the stored posts with filters
This query allows you to get all the stored posts that match one or more filters. 

**CLI**
```bash
desmos query posts [--flags]
```

Available flags: 
- `--subspace` (e.g. `--subspace=4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e`)
- `--page` (e.g. `--page=1`)
- `--limit` (e.g `--limit=10`)

```bash
# Example
# desmos query posts --subspace=4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e
```
