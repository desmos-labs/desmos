# Query a subspace
This query endpoint allows you to get the details of a single subspace with the given id

**CLI**
```bash
desmos query subspaces subspace [id]
```

**REST**
```bash
/subspaces/{subspace_id}

# Example
# curl http://lcd.morpheus.desmos.network:1317/subspaces/a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc
```

# Query all the subspaces
This query endpoint allows you to get all the stored subspaces

**CLI**
```bash
desmos query subspaces subspaces [--flags]
```

Available flags:
- `--page` (e.g `--page=1`)
- `--limit` (e.g `--limit=10`)

**REST**
```bash
/subspaces

# Example
# curl http://lcd.morpheus.desmos.network:1317/subspaces
```