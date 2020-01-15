# Query the stored posts
This query endpoint allows you to get all the stored posts that match one or more filters. 

**CLI**
```shell
desmoscli query posts [--flags]
```

Available flags: 
- `--parent-id` (e.g. `--parent-id 5`)
- `--creation-time` (e.g. `--creation-time 2020-01-01T12:00:00`)
- `--allows-comments` (e.g. `--allows-comments true`)
- `--subspace` (e.g. `--subspace desmos`)
- `--creator` (e.g. `--creator desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax`)

```shell
# Example
# desmoscli query posts --parent-id=1 --allows-comments=true --subspace=desmos
```

**REST**
```shell
/posts
```

Available parameters: 
- `parent_id` (e.g. `parent_id=5`)
- `creation_time` (e.g. `creation_time=2020-01-01T12:00:00`)
- `allows_comments` (e.g. `allows_comments=true`)
- `subspace` (e.g. `subspace=desmos`)
- `creator` (e.g. `creator=desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax`)

```shell
# Example
# curl https://morpheus1000.desmos.network/posts?parent_id=1&allows_comments=true&subspace=desmos
```