# Query the stored posts
This query endpoint allows you to get all the stored posts that match one or more filters. 

**CLI**
```bash
desmoscli query posts [--flags]
```

Available flags: 
- `--parent-id` (e.g. `--parent-id=a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc`)
- `--creation-time` (e.g. `--creation-time=2020-01-01T12:00:00`)
- `--allows-comments` (e.g. `--allows-comments=true`)
- `--subspace` (e.g. `--subspace=desmos`)
- `--creator` (e.g. `--creator=desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax`)
- `--sort-by` (e.g. `--sort-by=created`)  
   Accepted values: 
   - `created` 
   - `id` (default)
- `--sort-order` (e.g. `--sort-order=descending`)  
   Accepted values:
   - `ascending`
   - `descending`

```bash
# Example
# desmoscli query posts --parent-id=a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc --allows-comments=true --subspace=desmos --sort=created --sort-order=descending
```

**REST**
```bash
/posts
```

Available parameters: 
- `parent_id` (e.g. `parent_id=a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc`)
- `creation_time` (e.g. `creation_time=2020-01-01T12:00:00`)
- `allows_comments` (e.g. `allows_comments=true`)
- `subspace` (e.g. `subspace=desmos`)
- `creator` (e.g. `creator=desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax`)
- `sort_by` (e.g. `sort_by=created`)
- `sort_order` (e.g. `sort_order=descending`)

```bash
# Example
# curl https://morpheus4000.desmos.network/posts?parent_id=a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc&allows_comments=true&subspace=desmos&sort_by=created&sort_order=descending
```
