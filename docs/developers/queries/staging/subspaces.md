# Query a subspace
This query allows you to get the details of a single subspace with the given id

**CLI**
```bash
desmos query subspaces subspace [id]
```

# Query all the subspaces
This query allows you to get all the stored subspaces

**CLI**
```bash
desmos query subspaces subspaces [--flags]
```

# Query the admins of a subspace
This query allows to get all the admins of a subspace

**CLI**  
```bash
desmos query subspaces admins [subspace-id] [--flags]
```

# Query the registered users of a subspace
This query allows to get all the registered users of a subspace

**CLI**
```bash
desmos query subspaces registered-users [subspace-id] [--flags]
```

# Query the banned users of a subspace
This query allows to get all the banned users of a subspace

**CLI**
```bash
desmos query subspaces banned-users [subspace-id] [--flags]
```