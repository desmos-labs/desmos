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
- `--parent-id` (e.g. `--parent-id=a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc`)
- `--creation-time` (e.g. `--creation-time=2020-01-01T12:00:00`)
- `--subspace` (e.g. `--subspace=desmos`)
- `--creator` (e.g. `--creator=desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax`)
- `--hashtag` (e.g. `--hashtag=#desmos`)  
- `--sort-by` (e.g. `--sort-by=created`)  
   Accepted values: 
   - `created` 
   - `id` (default)
- `--sort-order` (e.g. `--sort-order=descending`)  
   Accepted values:
   - `ascending`
   - `descending`
- `--page` (e.g. `--page=1`)
- `--limit` (e.g `--limit=10`)

```bash
# Example
# desmos query posts --parent-id=a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc --disable-comments=false --subspace=desmos --sort=created --sort-order=descending
```
